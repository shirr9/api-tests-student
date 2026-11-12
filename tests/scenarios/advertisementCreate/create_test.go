package advertisementCreate

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

func TestMain(m *testing.M) {
	base.RunSuite(m)
}

func TestCreateAdvertisementWithPhoto(t *testing.T) {
	t.Log("Подготовка: авторизуем владельца и собираем объявление со всеми полями")
	owner := fixtures.LoginOwner(t)
	request := fixtures.NewAdvertisementRequest(t)

	t.Log("Создаём объявление с фотографией")
	created := fixtures.CreateAdvertisement(t, owner.Token, request)

	t.Log("Получаем объявление по ID и сравниваем поля с запросом")
	actual := advertisements.GetAdvertisement(t, created.ID)
	require.Equal(t, created.ID, actual.ID, "GET вернул объявление с другим ID")
	require.Equal(t, owner.User.Id, actual.UserID, "У объявления другой владелец")
	require.Equal(t, request.Fields[models.TitleField], actual.Title, "Название не сохранилось")
	require.Equal(t, request.Fields[models.DescriptionField], actual.Description, "Описание не сохранилось")
	require.NotNil(t, actual.Price, "В ответе нет цены")
	require.Equal(t, fixtures.AdvertisementPrice, *actual.Price, "Цена не сохранилась")
	require.NotNil(t, actual.Quantity, "В ответе нет количества")
	require.Equal(t, fixtures.AdvertisementQuantity, *actual.Quantity, "Количество не сохранилось")

	t.Log("Ищем объявление по полному названию и проверяем его ID")
	list := advertisements.GetAdvertisementsList(t, "", request.Fields[models.TitleField], "")
	found := false
	for _, item := range list.Items {
		if item.ID == created.ID {
			require.Equal(t, request.Fields[models.TitleField], item.Title, "Поиск вернул другое название")
			found = true
			break
		}
	}
	require.True(t, found, "Поиск по полному названию не вернул объявление %s", created.ID)

	t.Log("Получаем фотографии и сравниваем их с ответами создания и чтения")
	photos := advertisements.GetPhotos(t, owner.Token, created.ID)
	require.Len(t, photos, len(request.PhotoPaths), "Количество фотографий отличается от запроса")
	ids := make(map[string]bool, len(photos))
	for _, photo := range photos {
		require.NotEmpty(t, photo.ID, "У фотографии нет ID")
		require.NotEmpty(t, photo.URL, "У фотографии нет URL")
		require.NotContains(t, ids, photo.ID, "В ответе повторяется ID фотографии %s", photo.ID)
		ids[photo.ID] = true
	}
	require.ElementsMatch(t, created.Photos, photos, "GET /photos вернул другие фотографии, чем POST")
	require.ElementsMatch(t, actual.Photos, photos, "GET /photos вернул другие фотографии, чем GET объявления")
	t.Log("Поля, результат поиска и фотографии совпали")
}
