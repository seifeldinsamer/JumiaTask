package model

import (
	"backend/entity"
)

type ErrorOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type CountryUsersOutput struct {
	Data []entity.User `json:"data"`
}

type TotalCountryUsersOutput struct {
	Total int64 `json:"total"`
}
