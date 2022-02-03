package controllers

import (
	"api/app/models"
	"api/app/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPhones(c *gin.Context) {
	state := c.Query("state")
	country := c.Query("country")
	phones := []models.Phone{}
	dbCustomers := repositories.CustomerRepo.GetAll()
	for _, customer := range dbCustomers {
		phone := models.Phone{}
		phone.Prepare(&customer)

		if len(state) == 0 && len(country) == 0 {
			phones = append(phones, phone)
		} else if len(state) > 0 && len(country) > 0 {
			if state == phone.State && country == phone.Country {
				phones = append(phones, phone)
			}
		} else {
			if len(state) > 0 && state == phone.State {
				phones = append(phones, phone)
			}

			if len(country) > 0 && country == phone.Country {
				phones = append(phones, phone)
			}
		}
	}

	c.JSON(http.StatusOK, phones)
}
