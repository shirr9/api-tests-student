package models

import authModels "api-tests-template/internal/managers/auth/models"

type Advertisement struct {
	ID          string          `json:"id"`
	UserID      string          `json:"user_id"`
	User        authModels.User `json:"user"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Price       *float64        `json:"price"`
	Quantity    *int            `json:"quantity"`
	Photos      []Photo         `json:"photos"`
}

type Photo struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	SortOrder int    `json:"sort_order"`
}

type AdvertisementsList struct {
	Items      []Advertisement `json:"items"`
	HasNext    bool            `json:"has_next"`
	NextCursor string          `json:"next_cursor"`
}
