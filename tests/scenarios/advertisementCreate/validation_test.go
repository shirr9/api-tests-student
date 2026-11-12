package advertisementCreate

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"api-tests-template/internal/assertions"
	apiResponse "api-tests-template/internal/helpers/api-response"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/tests/fixtures"
)

const (
	maxTitleLength       = 50
	maxDescriptionLength = 500
	titleRequired        = "title is required"
	negativeNumber       = "-1"
	nonNumericValue      = "abc"
)

func TestCreateAdvertisementValidatesInput(t *testing.T) {
	cases := []struct {
		name    string
		field   string
		value   string
		omit    bool
		message string
	}{
		{name: "missing_title", field: models.TitleField, omit: true, message: titleRequired},
		{name: "empty_title", field: models.TitleField, message: titleRequired},
		{name: "missing_description", field: models.DescriptionField, omit: true, message: "description is required"},
		{name: "missing_photo", field: models.PhotosField, omit: true, message: "at least 1 photo is required"},
		{
			name: "long_title", field: models.TitleField, value: strings.Repeat("x", maxTitleLength+1),
			message: fmt.Sprintf("title must not exceed %d characters provided %d", maxTitleLength, maxTitleLength+1),
		},
		{
			name: "long_description", field: models.DescriptionField, value: strings.Repeat("x", maxDescriptionLength+1),
			message: fmt.Sprintf("description must be between 1 and %d characters", maxDescriptionLength),
		},
		{
			name: "negative_price", field: models.PriceField, value: negativeNumber,
			message: "price must be between 0 and 1000000 inclusive",
		},
		{
			name: "non_numeric_price", field: models.PriceField, value: nonNumericValue,
			message: fmt.Sprintf("cannot convert string value %s to float", nonNumericValue),
		},
		{
			name: "negative_quantity", field: models.QuantityField, value: negativeNumber,
			message: "quantity must be between 1 and 100 inclusive",
		},
		{
			name: "non_numeric_quantity", field: models.QuantityField, value: nonNumericValue,
			message: "cannot convert string value to int",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Подготовка: авторизуем владельца и меняем одно поле корректного запроса")
			owner := fixtures.LoginOwner(t)
			request := fixtures.NewAdvertisementRequest(t)
			switch {
			case tc.field == models.PhotosField:
				request.PhotoPaths = nil
			case tc.omit:
				delete(request.Fields, tc.field)
			default:
				request.Fields[tc.field] = tc.value
			}

			t.Logf("Отправляем POST /advertisement с невалидными данными: %s", tc.name)
			actual := advertisements.RejectCreation(t, owner.Token, request, http.StatusBadRequest, fixtures.LoginOwner)
			assertions.APIError(t, actual, apiResponse.InvalidRequest, tc.message)
		})
	}
}
