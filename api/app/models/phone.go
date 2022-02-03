package models

import (
	"log"
	"regexp"
	"strings"
)

type Phone struct {
	Country                string `json:"country"`
	State                  string `json:"state"`
	CountryCode            string `json:"countryCode"`
	PhoneNumberWithoutCode string `json:"phoneNumberFormated"`
}

func (p *Phone) Prepare(customer *Customer) {
	country := returnMapOfCountries(customer.Phone)
	p.Country = country.Name
	p.CountryCode = country.CountryCode
	p.validateNumber(customer.Phone, country.RegexToValidate)
	p.formatPhoneNumber(customer.Phone)
}

func (p *Phone) validateNumber(phone string, pattern string) {
	if ok, err := regexp.MatchString(pattern, phone); err != nil || !ok {
		if err != nil {
			log.Print(err.Error())
		}
		p.State = "NOK"
	} else {
		p.State = "OK"
	}
}

func (p *Phone) formatPhoneNumber(phone string) {
	p.PhoneNumberWithoutCode = strings.Split(phone, " ")[1]
}
