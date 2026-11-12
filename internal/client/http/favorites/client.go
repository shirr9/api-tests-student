package favorites

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/favorites/models"
)

// HttpGetFavorites читает избранное одной страницей размером с лимит.
func HttpGetFavorites(t *testing.T, token string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(path.FavoritesPath).
		Query("page_size", strconv.Itoa(models.MaxPerUser)).Expect(t).End().Response
}

// HttpPostFavorite отправляет ID объявления для добавления в избранное.
func HttpPostFavorite(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesPath).
		JSON(models.AddRequest{AdvertisementID: id}).Expect(t).End().Response
}

// HttpDeleteFavorite удаляет запись из избранного по ID объявления.
func HttpDeleteFavorite(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(fmt.Sprintf(path.FavoritePath, id)).
		Expect(t).End().Response
}

// HttpPostFavoritesBulk добавляет объявления одним запросом для подготовки теста лимита.
func HttpPostFavoritesBulk(t *testing.T, token string, ids []string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesBulkPath).
		JSON(models.BulkAddRequest{AdvertisementIDs: ids}).Expect(t).End().Response
}
