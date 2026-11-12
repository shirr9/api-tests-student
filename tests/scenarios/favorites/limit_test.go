package favorites

import (
	"fmt"
	"net/http"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorites"
)

func TestCannotExceedFavoritesLimit(t *testing.T) {
	t.Log("Подготовка: создаём объявление и заполняем свободные места в избранном покупателя")
	data := newFavoriteFixture(t)
	fillFavorites(t, data)
	before := favorites.List(t, data.buyer.Token)
	require.Equal(t, favorites.MaxPerUser, before.Total, "Избранное не заполнено перед проверкой лимита")
	require.Len(t, before.Items, before.Total, "Получили неполный список избранного перед проверкой лимита")

	t.Log("Пробуем добавить ещё одно объявление после достижения лимита")
	actual := favorites.RejectAdd(t, data.buyer.Token, data.ad.ID, http.StatusBadRequest)
	assertions.APIError(t, actual, "limit_reached", fmt.Sprintf("Favorites limit of %d reached", favorites.MaxPerUser))

	t.Log("Проверяем число записей и отсутствие нового объявления после отказа")
	after := favorites.List(t, data.buyer.Token)
	require.Equal(t, favorites.MaxPerUser, after.Total, "Число записей в избранном изменилось после отказа")
	require.Len(t, after.Items, after.Total, "Получили неполный список избранного после отказа")
	for _, item := range after.Items {
		require.NotEqual(t, data.ad.ID, item.Advertisement.ID, "Объявление %s попало в избранное после отказа", data.ad.ID)
	}
}

// fillFavorites заполняет свободные места чужими объявлениями и удаляет только добавленные записи.
func fillFavorites(t *testing.T, data favoriteFixture) {
	t.Helper()
	existing := favorites.List(t, data.buyer.Token)
	require.Len(t, existing.Items, existing.Total, "Получили неполный список избранного перед заполнением")
	require.LessOrEqual(t, existing.Total, favorites.MaxPerUser, "Избранное уже превышает лимит")
	slots := favorites.MaxPerUser - existing.Total
	if slots == 0 {
		return
	}

	excluded := map[string]bool{data.ad.ID: true}
	for _, item := range existing.Items {
		excluded[item.Advertisement.ID] = true
	}

	ids := make([]string, 0, slots)
	for cursor := ""; len(ids) < slots; {
		page := advertisements.GetAdvertisementsList(t, "", "", cursor)
		for _, ad := range page.Items {
			// TestBoard не возвращает список избранного, если у продавца нет фамилии.
			if ad.ID == "" || ad.User.LastName == "" || ad.UserID == data.buyer.User.Id || excluded[ad.ID] {
				continue
			}
			excluded[ad.ID] = true
			ids = append(ids, ad.ID)
			if len(ids) == slots {
				break
			}
		}
		if len(ids) == slots || !page.HasNext {
			break
		}
		require.NotEmpty(t, page.NextCursor, "Каталог сообщил о следующей странице без курсора")
		require.NotEqual(t, cursor, page.NextCursor, "Каталог вернул тот же курсор")
		cursor = page.NextCursor
	}
	require.Len(t, ids, slots, "В каталоге недостаточно чужих объявлений от продавцов с фамилией")

	responseReceived := false
	t.Cleanup(func() {
		if !responseReceived {
			t.Logf("Нет подтверждения результата массового добавления. Сверьте эти ID с избранным перед ручной очисткой: %v", ids)
		}
	})
	t.Logf("Добавляем в избранное %d объявлений из каталога", len(ids))
	added := favorites.BulkAdd(t, data.buyer.Token, ids)
	responseReceived = true

	confirmed := 0
	for _, id := range ids {
		if !slices.Contains(added, id) {
			continue
		}
		confirmed++
		t.Cleanup(func() { favorites.RemoveIfExists(t, data.buyer.Token, id) })
	}
	if confirmed > 0 {
		t.Cleanup(func() {
			t.Logf("Очистка: удаляем %d подтверждённых записей избранного", confirmed)
		})
	}
	require.ElementsMatch(t, ids, added, "Массовое добавление сохранило не все выбранные объявления")
}
