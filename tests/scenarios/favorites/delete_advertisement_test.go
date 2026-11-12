package favorites

import (
	"testing"

	"api-tests-template/internal/managers/favorites"
)

func TestDeletingAdvertisementRemovesFavorite(t *testing.T) {
	t.Log("Подготовка: создаём объявление и добавляем его в избранное покупателя")
	data := newFavoriteFixture(t)
	favorites.Add(t, data.buyer.Token, data.ad.ID)
	requireFavorite(t, data.buyer.Token, data.ad.ID, 1)

	t.Log("Владелец удаляет само объявление")
	data.ad.Delete(t)

	t.Log("Повторно читаем избранное покупателя")
	requireFavorite(t, data.buyer.Token, data.ad.ID, 0)
}
