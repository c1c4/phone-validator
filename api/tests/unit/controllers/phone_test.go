package controllers

import (
	"api/app/controllers"
	"api/app/models"
	"api/app/repositories"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	getCustomersRepository func() []models.Customer
	handlerGetAllPhones    = controllers.GetAllPhones
)

type customerRepoMock struct{}

func (customerRepo *customerRepoMock) GetAll() []models.Customer {
	return getCustomersRepository()
}

func (customerRepo *customerRepoMock) Init() {}

func TestGetCustomer_Success(t *testing.T) {
	repositories.CustomerRepo = &customerRepoMock{}

	getCustomersRepository = func() []models.Customer {
		return []models.Customer{
			{
				ID:    1,
				Name:  "Test",
				Phone: "(237) 696443597",
			},
			{
				ID:    2,
				Name:  "Test 2",
				Phone: "(212) 6007989253",
			},
			{
				ID:    3,
				Name:  "Test 3",
				Phone: "(212) 698054317",
			},
		}
	}

	r := gin.Default()
	req, errRequest := http.NewRequest(http.MethodGet, "/phones", nil)

	if errRequest != nil {
		t.Errorf("this is the error: %v\n", errRequest)
	}

	rr := httptest.NewRecorder()
	r.GET("/phones", handlerGetAllPhones)
	r.ServeHTTP(rr, req)

	var phones []models.Phone
	err := json.Unmarshal(rr.Body.Bytes(), &phones)
	assert.Nil(t, err)
	assert.NotNil(t, phones)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, 3, len(phones))
}

func TestGetCustomer_EmptyList(t *testing.T) {
	repositories.CustomerRepo = &customerRepoMock{}

	getCustomersRepository = func() []models.Customer {
		return []models.Customer{}
	}

	r := gin.Default()
	req, errRequest := http.NewRequest(http.MethodGet, "/phones", nil)

	if errRequest != nil {
		t.Errorf("this is the error: %v\n", errRequest)
	}

	rr := httptest.NewRecorder()
	r.GET("/phones", handlerGetAllPhones)
	r.ServeHTTP(rr, req)

	var phones []models.Phone
	err := json.Unmarshal(rr.Body.Bytes(), &phones)
	assert.Nil(t, err)
	assert.NotNil(t, phones)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, 0, len(phones))
}
