package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	apiResponse "api-tests-template/internal/helpers/api-response"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

func TestFavoritesRequireAuthorization(t *testing.T) {
	operations := []struct {
		name string
		call func(*testing.T, string) apiResponse.Error
	}{
		{"list", favorites.RejectList},
		{"add", func(t *testing.T, token string) apiResponse.Error {
			return favorites.RejectAdd(t, token, absentAdvertisementID, http.StatusUnauthorized)
		}},
		{"remove", func(t *testing.T, token string) apiResponse.Error {
			return favorites.RejectRemove(t, token, absentAdvertisementID, http.StatusUnauthorized)
		}},
	}

	for _, operation := range operations {
		t.Run(operation.name, func(t *testing.T) {
			for _, tc := range fixtures.UnauthorizedCases() {
				t.Run(tc.Name, func(t *testing.T) {
					t.Log("Подготовка: используем отсутствующий или неверный токен")
					t.Logf("Вызываем операцию избранного: %s", operation.name)
					assertions.APIError(t, operation.call(t, tc.Token), apiResponse.Unauthorized, tc.Message)
				})
			}
		})
	}
}
