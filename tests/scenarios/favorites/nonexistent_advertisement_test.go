package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	apiResponse "api-tests-template/internal/helpers/api-response"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

func TestAddNonexistentAdvertisement(t *testing.T) {
	t.Log("Подготовка: авторизуемся и берём UUID отсутствующего объявления")
	user := fixtures.LoginOwner(t)

	t.Log("Пробуем добавить несуществующее объявление в избранное")
	actual := favorites.RejectAdd(t, user.Token, absentAdvertisementID, http.StatusNotFound)
	assertions.APIError(t, actual, apiResponse.NotFound, "Advertisement not found")
}
