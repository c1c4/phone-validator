package models

import (
	"strings"
)

type Country struct {
	Name            string
	RegexToValidate string
	CountryCode     string
}

func returnMapOfCountries(phoneNumber string) Country {
	code := strings.Split(phoneNumber, " ")[0][1:4]
	countriesMapByCountryCode := map[string]Country{
		"237": {
			Name:            "Cameroon",
			RegexToValidate: `\(237\)\ ?[2368]\d{7,8}$`,
			CountryCode:     "+237",
		},
		"251": {
			Name:            "Ethiopia",
			RegexToValidate: `\(251\)\ ?[1-59]\d{8}$`,
			CountryCode:     "+251",
		},
		"212": {
			Name:            "Morocco",
			RegexToValidate: `\(212\)\ ?[5-9]\d{8}$`,
			CountryCode:     "+212",
		},
		"258": {
			Name:            "Mozambique",
			RegexToValidate: `\(258\)\ ?[28]\d{7,8}$`,
			CountryCode:     "+258",
		},
		"256": {
			Name:            "Uganda",
			RegexToValidate: `\(256\)\ ?\d{9}$`,
			CountryCode:     "+256",
		},
	}
	return countriesMapByCountryCode[code]
}
