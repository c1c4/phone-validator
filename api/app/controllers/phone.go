package controllers

import (
	"api/app/models"
	"api/app/repositories"
	"api/app/utils/error_utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPhones(c *gin.Context) {
	var err error
	page := 1
	state := c.Query("state")
	country := c.Query("country")
	pageString := c.Query("page")
	phones := []models.Phone{}

	if len(pageString) > 0 {
		page, err = strconv.Atoi(pageString)

		if err != nil {
			errBadRequest := error_utils.NewBadRequestError(fmt.Sprintf("not possible to convert %s into a number", pageString))
			c.JSON(errBadRequest.Status(), errBadRequest)
			return
		}
	}

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

	pagination := models.Pagination{}
	pagination.Prepare(phones, page)

	paginationPhones := models.PaginationPhone{}

	paginationPhones.Prepare(phones, pagination, page)

	c.JSON(http.StatusOK, paginationPhones)
}
