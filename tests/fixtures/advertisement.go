package fixtures

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/utils"
)

const (
	AdvertisementPrice    = 1290.50
	AdvertisementQuantity = 2
	titleSuffixLength     = 10
)

type AdvertisementFixture struct {
	models.Advertisement
	token   string
	deleted bool
}

// NewAdvertisementRequest собирает новый запрос с уникальным названием и локальными фотографиями.
func NewAdvertisementRequest(t *testing.T) models.CreateRequest {
	t.Helper()
	title := strings.TrimSpace(os.Getenv("TEST_AD_TITLE"))
	require.NotEmpty(t, title, "Укажите TEST_AD_TITLE в .env.override")

	photos := []string{photoPath(t, "TEST_AD_PHOTO")}
	if os.Getenv("TEST_AD_SECOND_PHOTO") != "" {
		photos = append(photos, photoPath(t, "TEST_AD_SECOND_PHOTO"))
	}

	return models.CreateRequest{
		Fields: map[string]string{
			models.TitleField:       title + " " + utils.RandomString(titleSuffixLength),
			models.DescriptionField: title,
			models.PriceField:       strconv.FormatFloat(AdvertisementPrice, 'f', 2, 64),
			models.QuantityField:    strconv.Itoa(AdvertisementQuantity),
		},
		PhotoPaths: photos,
	}
}

// CreateAdvertisement создаёт объявление и регистрирует удаление после теста.
func CreateAdvertisement(t *testing.T, token string, request models.CreateRequest) *AdvertisementFixture {
	t.Helper()
	ad := advertisements.CreateAdvertisement(t, token, request)
	require.NotEmpty(t, ad.ID, "В ответе создания нет ID объявления")
	fixture := &AdvertisementFixture{Advertisement: ad, token: token}
	t.Logf("Создали объявление %s", ad.ID)
	t.Cleanup(func() {
		if fixture.deleted {
			return
		}
		t.Logf("Очистка: удаляем тестовое объявление %s", ad.ID)
		advertisements.DeleteIfExists(t, token, ad.ID)
	})
	return fixture
}

// Delete удаляет объявление во время сценария и отмечает, что повторная очистка не нужна.
func (ad *AdvertisementFixture) Delete(t *testing.T) {
	t.Helper()
	advertisements.DeleteAdvertisement(t, ad.token, ad.ID)
	ad.deleted = true
}

// photoPath проверяет, что указан абсолютный путь к существующей фотографии.
func photoPath(t *testing.T, key string) string {
	t.Helper()
	path := os.Getenv(key)
	require.True(t, filepath.IsAbs(path), "Укажите абсолютный путь в %s в .env.override", key)
	require.FileExists(t, path, "Фотография из %s не найдена", key)
	return path
}
