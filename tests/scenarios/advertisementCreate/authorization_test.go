package advertisementCreate

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	apiResponse "api-tests-template/internal/helpers/api-response"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/tests/fixtures"
)

func TestCreateAdvertisementRequiresAuthorization(t *testing.T) {
	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			t.Log("Подготовка: собираем корректное объявление")
			request := fixtures.NewAdvertisementRequest(t)

			t.Log("Отправляем POST /advertisement без действительного токена")
			actual := advertisements.RejectCreation(t, tc.Token, request, http.StatusUnauthorized, fixtures.LoginOwner)
			assertions.APIError(t, actual, apiResponse.Unauthorized, tc.Message)
		})
	}
}
