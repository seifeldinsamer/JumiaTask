package users

import (
	"backend/entity"
)

type UseCase interface {
	CountryUsers(countryName string) []entity.User
	TotalCountryUsers(countryName string) int64
}
