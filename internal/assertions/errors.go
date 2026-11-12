package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	apiResponse "api-tests-template/internal/helpers/api-response"
)

// APIError проверяет машинный код и сообщение об ошибке.
func APIError(t *testing.T, actual apiResponse.Error, code, message string) {
	t.Helper()
	require.Equal(t, code, actual.Code, "Некорректное поле error: %s", actual.Message)
	require.Equal(t, message, actual.Message, "Некорректное поле message")
	t.Logf("Получили ожидаемую ошибку %s: %s", actual.Code, actual.Message)
}
