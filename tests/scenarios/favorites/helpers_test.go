package favorites

import (
	"testing"

	"github.com/stretchr/testify/require"

	authModels "api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

const absentAdvertisementID = "2a62eb60-a742-4a84-869f-d8719ea99bcb"

type favoriteFixture struct {
	ad    *fixtures.AdvertisementFixture
	buyer authModels.LoginOkResponse
}

// newFavoriteFixture создаёт объявление одного пользователя для избранного другого.
func newFavoriteFixture(t *testing.T) favoriteFixture {
	t.Helper()
	owner := fixtures.LoginOwner(t)
	require.NotEmpty(t, owner.User.LastName, "Укажите фамилию владельца: иначе TestBoard не читает избранное с его объявлениями")
	buyer := fixtures.LoginBuyer(t, owner.User.Id)
	ad := fixtures.CreateAdvertisement(t, owner.Token, fixtures.NewAdvertisementRequest(t))
	t.Cleanup(func() { favorites.RemoveIfExists(t, buyer.Token, ad.ID) })

	return favoriteFixture{ad: ad, buyer: buyer}
}

// requireFavorite проверяет, сколько раз объявление встречается в избранном.
func requireFavorite(t *testing.T, token, adID string, expectedCount int) {
	t.Helper()
	t.Logf("Проверяем GET /favorites для объявления %s", adID)
	list := favorites.List(t, token)
	count := 0
	for _, item := range list.Items {
		if item.Advertisement.ID == adID {
			count++
		}
	}
	require.Equal(t, expectedCount, count, "Число записей с объявлением %s в избранном", adID)
	t.Logf("В избранном найдено записей с объявлением: %d", count)
}
