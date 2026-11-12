package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	apiResponse "api-tests-template/internal/helpers/api-response"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

func TestAddFavoriteRequiresValidAdvertisementID(t *testing.T) {
	cases := []struct {
		name    string
		adID    string
		message string
	}{
		{"missing_id", "", "Key: 'AddFavoriteRequest.AdvertisementID' Error:Field validation for 'AdvertisementID' failed on the 'required' tag"},
		{"malformed_id", "not-a-uuid", "invalid UUID length: 10"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Подготовка: авторизуемся")
			user := fixtures.LoginOwner(t)

			t.Logf("Пробуем добавить в избранное с невалидным ID: %s", tc.name)
			actual := favorites.RejectAdd(t, user.Token, tc.adID, http.StatusBadRequest)
			assertions.APIError(t, actual, apiResponse.InvalidRequest, tc.message)
		})
	}
}
