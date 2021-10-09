package users

import (
	"backend/entity"
)

const NoUsersAvailable = -1

func (service *Service) CountryUsers(countryName string) []entity.User {
	users, err := service.Repository.GetCountryUsers(countryName)

	if err != nil {
		service.Logger.Error().Msgf("failed to get users with error: %s", err)
		return nil
	}

	return users
}

func (service *Service) TotalCountryUsers(countryName string) int64 {
	total := service.Repository.GetTotalCountryUsers(countryName)

	if total == 0 {
		service.Logger.Error().Msgf("failed to get users with error")
		return NoUsersAvailable
	}

	return total
}
