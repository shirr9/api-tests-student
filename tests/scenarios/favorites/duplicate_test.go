package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
)

func TestCannotAddFavoriteTwice(t *testing.T) {
	t.Log("Подготовка: создаём чужое объявление и добавляем его в избранное")
	data := newFavoriteFixture(t)
	favorites.Add(t, data.buyer.Token, data.ad.ID)

	t.Log("Пробуем добавить это же объявление повторно")
	actual := favorites.RejectAdd(t, data.buyer.Token, data.ad.ID, http.StatusConflict)
	assertions.APIError(t, actual, "already_exists", "Advertisement is already in favorites")
}
