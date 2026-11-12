package advertisementSearch

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/utils"

	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

const (
	existingAdTitle = "QA Practice — тестовое объявление 1"
	searchBeginning = "QA Practice"
	searchMiddle    = "Practice — тестовое"
	searchEnd       = "объявление 1"
)

type TestSuite struct {
	suite.Suite
	token string
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	t := s.T()
	require.NotEmpty(t, existingAdTitle, "Задайте полное название существующего объявления")
	require.NotEmpty(t, searchBeginning, "Задайте подстроку из начала названия")
	require.NotEmpty(t, searchMiddle, "Задайте подстроку из середины названия")
	require.NotEmpty(t, searchEnd, "Задайте подстроку из конца названия")

	s.token = fixtures.LoginOwner(t).Token
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

func (s *TestSuite) TestSearchByTitle() {
	t := s.T()
	list := advertisements.GetAdvertisementsList(t, s.token, searchMiddle, "")

	require.NotEmpty(t, list.Items, "Поиск по названию вернул пустой список")
	assertions.ContainsAdvertisementByTitle(t, list.Items, existingAdTitle)
}

func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name   string
		search string
	}{
		{name: "beginning", search: searchBeginning},
		{name: "middle", search: searchMiddle},
		{name: "end", search: searchEnd},
		{name: "lowercase", search: strings.ToLower(searchMiddle)},
		{name: "uppercase", search: strings.ToUpper(searchMiddle)},
		{name: "mixed_case", search: utils.MixCase(searchMiddle)},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			t := s.T()
			list := advertisements.GetAdvertisementsList(t, s.token, tc.search, "")
			assertions.ContainsAdvertisementByTitle(t, list.Items, existingAdTitle)
		})
	}
}
