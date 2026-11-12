package favorites

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/favorites"
	apiResponse "api-tests-template/internal/helpers/api-response"
	"api-tests-template/internal/managers/favorites/models"
)

const MaxPerUser = models.MaxPerUser

// List возвращает избранное авторизованного пользователя.
func List(t *testing.T, token string) models.FavoriteList {
	t.Helper()
	return apiResponse.JSON[models.FavoriteList](t, favorites.HttpGetFavorites(t, token), http.StatusOK)
}

// Add проверяет, что добавилось нужное объявление и у записи есть ID.
func Add(t *testing.T, token, id string) {
	t.Helper()
	item := apiResponse.JSON[models.Favorite](t, favorites.HttpPostFavorite(t, token, id), http.StatusCreated)
	require.NotEmpty(t, item.ID, "У записи избранного нет favorite_id")
	require.Equal(t, id, item.Advertisement.ID, "В ответе вернулся другой ID объявления")
}

// Remove проверяет успешное удаление записи из избранного.
func Remove(t *testing.T, token, id string) {
	t.Helper()
	apiResponse.Body(t, favorites.HttpDeleteFavorite(t, token, id), http.StatusNoContent)
}

// RemoveIfExists очищает тестовую запись, если сценарий не удалил её раньше.
func RemoveIfExists(t *testing.T, token, id string) {
	t.Helper()
	apiResponse.Body(t, favorites.HttpDeleteFavorite(t, token, id), http.StatusNoContent, http.StatusNotFound)
}

// BulkAdd возвращает ID объявлений, добавленных в избранное одним запросом.
func BulkAdd(t *testing.T, token string, ids []string) []string {
	t.Helper()
	return apiResponse.JSON[models.BulkAddResponse](t, favorites.HttpPostFavoritesBulk(t, token, ids), http.StatusOK).Added
}

// RejectList возвращает ошибку чтения избранного без авторизации.
func RejectList(t *testing.T, token string) apiResponse.Error {
	t.Helper()
	return apiResponse.JSON[apiResponse.Error](t, favorites.HttpGetFavorites(t, token), http.StatusUnauthorized)
}

// RejectAdd возвращает ошибку добавления в избранное.
func RejectAdd(t *testing.T, token, id string, status int) apiResponse.Error {
	t.Helper()
	return apiResponse.JSON[apiResponse.Error](t, favorites.HttpPostFavorite(t, token, id), status)
}

// RejectRemove возвращает ошибку удаления из избранного.
func RejectRemove(t *testing.T, token, id string, status int) apiResponse.Error {
	t.Helper()
	return apiResponse.JSON[apiResponse.Error](t, favorites.HttpDeleteFavorite(t, token, id), status)
}
