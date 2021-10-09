package users

import (
	"backend/entity"
)

type Repository interface {
	GetTotalCountryUsers(country string) int64
	GetCountryUsers(country string) ([]entity.User, error)
}
