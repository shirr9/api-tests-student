package apiresponse

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	Unauthorized   = "unauthorized"
	InvalidRequest = "invalid_request"
	NotFound       = "not_found"
)

type Error struct {
	Code    string `json:"error"`
	Message string `json:"message"`
}

// Body читает и закрывает тело ответа, проверяя допустимые статусы.
func Body(t *testing.T, response *http.Response, statuses ...int) []byte {
	t.Helper()
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err, "Не удалось прочитать тело ответа HTTP %d", response.StatusCode)
	require.Contains(t, statuses, response.StatusCode, "Неожиданный статус HTTP %d: %.400s", response.StatusCode, body)
	return body
}

// JSON проверяет статус и разбирает тело ответа в заданную модель.
func JSON[T any](t *testing.T, response *http.Response, status int) T {
	t.Helper()
	body := Body(t, response, status)

	var value T
	require.NoError(t, json.Unmarshal(body, &value), "Не удалось разобрать JSON ответа: %s", body)
	return value
}
