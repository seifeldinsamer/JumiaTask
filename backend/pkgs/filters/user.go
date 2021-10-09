package filters

import (
	"backend/entity"
	"backend/pkgs/phoneNumberValidator"
	"github.com/pkg/errors"
)

type Filter func(user *entity.User) error

func CountryHandler() Filter {
	return func(user *entity.User) error {
		if user == nil {
			return errors.Errorf("Empty record")
		}
		country, err := phoneNumberValidator.FindPhoneNumberCountry(user.PhoneNumber)

		if err != nil {
			return err
		}
		user.Country = country
		return nil
	}
}
