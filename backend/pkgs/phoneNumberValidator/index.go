package phoneNumberValidator

import (
	"github.com/pkg/errors"
	"regexp"
)

var (
	countryToCodeExpression = map[string]string{
		"Cameroon":   "237\\ ?[2368]\\d{7,8}$",
		"Ethiopia":   "251\\ ?[1-59]\\d{8}$",
		"Morocco":    "212\\ ?[5-9]\\d{8}$",
		"Mozambique": "258\\ ?[28]\\d{7,8}$",
		"Uganda":     "256\\ ?\\d{9}$",
	}
)

func FindPhoneNumberCountry(number string) (string, error) {
	for country, codeExp := range countryToCodeExpression {
		if match, _ := regexp.MatchString(codeExp, number); match {
			return country, nil
		}
	}
	return "", errors.Errorf("Missing country")
}
