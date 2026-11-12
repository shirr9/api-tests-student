package advertisements

import (
	"fmt"
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/advertisements/models"
)

const catalogPageSize = "50"

// HttpPostAdvertisement отправляет поля и фотографии в multipart/form-data.
func HttpPostAdvertisement(t *testing.T, token string, data models.CreateRequest) *http.Response {
	t.Helper()
	request := apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath)
	response := request.Expect(t)
	for field, value := range data.Fields {
		request.MultipartFormData(field, value)
	}
	if len(data.PhotoPaths) > 0 {
		request.MultipartFile(models.PhotosField, data.PhotoPaths...)
	}
	return response.End().Response
}

// HttpGetAdvertisement читает объявление по ID.
func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Create().Get(path.AdvertisementPath).
		Query("id", id).Expect(t).End().Response
}

// HttpGetAdvertisements читает страницу каталога с учётом поиска и курсора.
func HttpGetAdvertisements(t *testing.T, token, search, cursor string) *http.Response {
	t.Helper()
	request := apiRunner.GetRunner().Auth(token).Create().Get(path.AdvertisementsPath).Query("limit", catalogPageSize)
	if search != "" {
		request.Query("search", search)
	}
	if cursor != "" {
		request.Query("cursor", cursor)
	}
	return request.Expect(t).End().Response
}

// HttpGetAdvertisementPhotos читает фотографии объявления.
func HttpGetAdvertisementPhotos(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(fmt.Sprintf(path.AdvertisementPhotosPath, id)).
		Expect(t).End().Response
}

// HttpDeleteAdvertisement удаляет объявление владельца по ID.
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.AdvertisementPath).
		Query("id", id).Expect(t).End().Response
}
