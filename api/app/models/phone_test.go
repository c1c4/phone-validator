package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepare(t *testing.T) {
	customer := Customer {
		ID: 1,
		Name: "Test",
		Phone: "(237) 696443597",
	}

	phone := Phone{}

	phone.Prepare(&customer)

	assert.Equal(t, "696443597", phone.PhoneNumberWithoutCode)
	assert.Equal(t, "OK", phone.State)
	assert.Equal(t, "+237", phone.CountryCode)
	assert.Equal(t, "Cameroon", phone.Country )
}

func TestPrepareError(t *testing.T) {
	customer := Customer {
		ID: 1,
		Name: "Test",
		Phone: "(999) 696443597",
	}

	phone := Phone{}

	phone.Prepare(&customer)

	assert.Equal(t, "696443597", phone.PhoneNumberWithoutCode)
	assert.Equal(t, "OK", phone.State)
	assert.Equal(t, "", phone.CountryCode)
	assert.Equal(t, "", phone.Country )
}
