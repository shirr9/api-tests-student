package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

func TestCannotFavoriteOwnAdvertisement(t *testing.T) {
	t.Log("Подготовка: создаём объявление от имени владельца")
	owner := fixtures.LoginOwner(t)
	ad := fixtures.CreateAdvertisement(t, owner.Token, fixtures.NewAdvertisementRequest(t))

	t.Log("Пробуем добавить своё объявление в избранное")
	actual := favorites.RejectAdd(t, owner.Token, ad.ID, http.StatusForbidden)
	assertions.APIError(t, actual, "forbidden", "Cannot add own advertisement to favorites")
}
