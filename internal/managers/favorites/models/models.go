package models

import adModels "api-tests-template/internal/managers/advertisements/models"

const MaxPerUser = 50

type AddRequest struct {
	AdvertisementID string `json:"advertisement_id,omitempty"`
}

type BulkAddRequest struct {
	AdvertisementIDs []string `json:"advertisement_ids"`
}

type BulkAddResponse struct {
	Added []string `json:"added"`
}

type Favorite struct {
	ID            string                 `json:"favorite_id"`
	Advertisement adModels.Advertisement `json:"advertisement"`
}

type FavoriteList struct {
	Items []Favorite `json:"items"`
	Total int        `json:"total"`
}
