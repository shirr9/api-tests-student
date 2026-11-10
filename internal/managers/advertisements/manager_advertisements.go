package advertisements

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
)

// GetAdvertisementsList возвращает список объявлений из результатов поиска.
func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	response := advertisements.HttpGetAdvertisements(t, token, text)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err, "Не удалось прочитать ответ поиска объявлений")

	var list models.AdvertisementsList
	require.NoError(t, json.Unmarshal(body, &list), "Не удалось разобрать JSON ответа поиска объявлений")
	return list
}
