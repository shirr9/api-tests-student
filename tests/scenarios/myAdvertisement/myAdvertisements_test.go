package myAdvertisement

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/assertions"
	apiResponse "api-tests-template/internal/helpers/api-response"
	"api-tests-template/internal/managers/myAdvertisements"

	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

const pageSize = 50

type TestSuite struct {
	suite.Suite
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

// Positive: Проверяем получение собственных объявлений в профиле
func (s *TestSuite) TestGetMyAdvertisementsPositive() {
	t := s.T()
	base.Precondition("Авторизуем владельца и создаём объявление для этого теста")
	owner := fixtures.LoginOwner(t)
	request := fixtures.NewAdvertisementRequest(t)
	ad := fixtures.CreateAdvertisement(t, owner.Token, request)

	t.Log("Получаем список собственных объявлений")
	body := myAdvertisements.GetMyAdvertisements(t, owner.Token, http.StatusOK, pageSize)
	var actual gjson.Result
	t.Log("Проверяем владельца и обязательные поля каждого объявления")
	for _, item := range gjson.Get(body, "items").Array() {
		require.Equal(t, owner.User.Id, item.Get("user_id").String(), "У объявления другой владелец")
		require.NotEmpty(t, item.Get("id").String(), "Нет ID объявления")
		require.NotEmpty(t, item.Get("created_at").String(), "Нет даты создания")
		require.NotEmpty(t, item.Get("updated_at").String(), "Нет даты обновления")
		require.NotEmpty(t, item.Get("title").String(), "Нет названия")
		require.NotEmpty(t, item.Get("description").String(), "Нет описания")
		require.NotEmpty(t, item.Get("photos").Array(), "Нет фотографий")
		if item.Get("id").String() == ad.ID {
			actual = item
		}
	}
	require.True(t, actual.Exists(), "В списке собственных объявлений нет %s", ad.ID)
	require.Equal(t, ad.Title, actual.Get("title").String(), "Название не совпало")
	require.Equal(t, ad.Description, actual.Get("description").String(), "Описание не совпало")
	require.Len(t, actual.Get("photos").Array(), len(request.PhotoPaths), "Количество фотографий не совпало")
	t.Log("Объявление владельца найдено, поля совпали")
}

// Negative: Проверяем 401 при запросе собственных объявлений с неверным токеном.
func (s *TestSuite) TestGetMyAdvertisementsIncorrectToken() {
	t := s.T()
	t.Log("Запрашиваем список собственных объявлений с неверным токеном")
	body := myAdvertisements.GetMyAdvertisements(t, fixtures.InvalidToken, http.StatusUnauthorized, pageSize)
	actual := apiResponse.Error{
		Code:    gjson.Get(body, "error").String(),
		Message: gjson.Get(body, "message").String(),
	}
	assertions.APIError(t, actual, apiResponse.Unauthorized, fixtures.InvalidBearerMessage)
}
