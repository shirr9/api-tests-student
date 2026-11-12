package favorites

import (
	"testing"

	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func TestMain(m *testing.M) {
	base.RunSuite(m)
}

func TestAddAndRemoveFavorite(t *testing.T) {
	t.Log("Подготовка: создаём объявление владельца и авторизуем покупателя")
	data := newFavoriteFixture(t)

	t.Log("Добавляем чужое объявление в избранное")
	favorites.Add(t, data.buyer.Token, data.ad.ID)
	requireFavorite(t, data.buyer.Token, data.ad.ID, 1)

	t.Log("Удаляем объявление из избранного")
	favorites.Remove(t, data.buyer.Token, data.ad.ID)
	requireFavorite(t, data.buyer.Token, data.ad.ID, 0)
}
