package advertisements

import (
	"net/http"
	"testing"

	"api-tests-template/internal/client/http/advertisements"
	apiResponse "api-tests-template/internal/helpers/api-response"
	"api-tests-template/internal/managers/advertisements/models"
	authModels "api-tests-template/internal/managers/auth/models"
)

// CreateAdvertisement проверяет успешное создание и возвращает объявление.
func CreateAdvertisement(t *testing.T, token string, request models.CreateRequest) models.Advertisement {
	t.Helper()
	return apiResponse.JSON[models.Advertisement](t, advertisements.HttpPostAdvertisement(t, token, request), http.StatusCreated)
}

// RejectCreation проверяет ошибку создания и при неожиданном успехе удаляет объявление через аккаунт владельца.
func RejectCreation(t *testing.T, token string, request models.CreateRequest, status int, loginOwner func(*testing.T) authModels.LoginOkResponse) apiResponse.Error {
	t.Helper()
	response := advertisements.HttpPostAdvertisement(t, token, request)
	if response.StatusCode == http.StatusCreated {
		ad := apiResponse.JSON[struct {
			ID     string `json:"id"`
			UserID string `json:"user_id"`
		}](t, response, http.StatusCreated)
		if ad.ID == "" {
			t.Fatalf("POST /advertisement вернул HTTP 201 без ID для очистки, ожидался HTTP %d", status)
		}
		t.Cleanup(func() {
			owner := loginOwner(t)
			if ad.UserID == "" || ad.UserID != owner.User.Id {
				t.Errorf("Владелец объявления %s не подтверждён. Проверьте объявление перед ручной очисткой", ad.ID)
				return
			}
			t.Logf("Очистка: удаляем объявление %s через аккаунт владельца", ad.ID)
			DeleteIfExists(t, owner.Token, ad.ID)
		})
		t.Fatalf("POST /advertisement создал объявление %q, ожидался HTTP %d", ad.ID, status)
	}
	return apiResponse.JSON[apiResponse.Error](t, response, status)
}

// GetAdvertisement возвращает поля объявления по ID.
func GetAdvertisement(t *testing.T, id string) models.Advertisement {
	t.Helper()
	return apiResponse.JSON[models.Advertisement](t, advertisements.HttpGetAdvertisement(t, id), http.StatusOK)
}

// GetAdvertisementsList возвращает страницу каталога с результатами поиска или без фильтра.
func GetAdvertisementsList(t *testing.T, token, search, cursor string) models.AdvertisementsList {
	t.Helper()
	return apiResponse.JSON[models.AdvertisementsList](t, advertisements.HttpGetAdvertisements(t, token, search, cursor), http.StatusOK)
}

// GetPhotos возвращает фотографии объявления.
func GetPhotos(t *testing.T, token, id string) []models.Photo {
	t.Helper()
	return apiResponse.JSON[[]models.Photo](t, advertisements.HttpGetAdvertisementPhotos(t, token, id), http.StatusOK)
}

// DeleteAdvertisement проверяет успешное удаление объявления.
func DeleteAdvertisement(t *testing.T, token, id string) {
	t.Helper()
	apiResponse.Body(t, advertisements.HttpDeleteAdvertisement(t, token, id), http.StatusNoContent)
}

// DeleteIfExists убирает объявление после теста, даже если сценарий уже удалил его.
func DeleteIfExists(t *testing.T, token, id string) {
	t.Helper()
	apiResponse.Body(t, advertisements.HttpDeleteAdvertisement(t, token, id), http.StatusNoContent, http.StatusNotFound)
}
