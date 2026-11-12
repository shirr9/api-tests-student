package fixtures

import (
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
)

const (
	InvalidToken         = "incorrect_token"
	InvalidBearerMessage = "Invalid or expired token" // #nosec G101 -- Ожидаемый текст ответа API.
)

type AuthorizationCase struct {
	Name    string
	Token   string
	Message string
}

type testAccount struct {
	loginKey    string
	passwordKey string
	mu          sync.Mutex
	data        models.LoginOkResponse
}

var (
	ownerAccount = testAccount{loginKey: "TEST_OWNER_LOGIN", passwordKey: "TEST_OWNER_PASSWORD"}
	buyerAccount = testAccount{loginKey: "TEST_BUYER_LOGIN", passwordKey: "TEST_BUYER_PASSWORD"}
)

// UnauthorizedCases задаёт запросы без заголовка и с неверным токеном.
func UnauthorizedCases() []AuthorizationCase {
	return []AuthorizationCase{
		{Name: "missing_token", Message: "Authorization header required"},
		{Name: "invalid_token", Token: InvalidToken, Message: InvalidBearerMessage},
	}
}

// LoginOwner входит в аккаунт, от имени которого создаются объявления.
func LoginOwner(t *testing.T) models.LoginOkResponse {
	t.Helper()
	return ownerAccount.login(t)
}

// LoginBuyer входит в отдельный аккаунт для операций с чужими объявлениями.
func LoginBuyer(t *testing.T, ownerID string) models.LoginOkResponse {
	t.Helper()
	buyer := buyerAccount.login(t)
	require.NotEqual(t, ownerID, buyer.User.Id, "Для владельца и покупателя нужны разные аккаунты TestBoard")
	return buyer
}

// login получает токен один раз на пакет, чтобы не упираться в лимит запросов входа.
func (account *testAccount) login(t *testing.T) models.LoginOkResponse {
	t.Helper()
	account.mu.Lock()
	defer account.mu.Unlock()
	if account.data.Token != "" {
		return account.data
	}

	login, password := os.Getenv(account.loginKey), os.Getenv(account.passwordKey)
	require.NotEmpty(t, login, "Укажите %s в .env.override", account.loginKey)
	require.NotEmpty(t, password, "Укажите %s в .env.override", account.passwordKey)
	result := auth.Login(t, login, password)
	require.NotEmpty(t, result.Token, "TestBoard не вернул токен для %s", account.loginKey)
	require.NotEmpty(t, result.User.Id, "TestBoard не вернул ID пользователя для %s", account.loginKey)
	account.data = result
	return result
}
