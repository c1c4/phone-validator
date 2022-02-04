package integration

import (
	"api/app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *SuiteTest) TestGetPhones_Success() {
	s.seedMultipleCustomers()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/phones", baseURL), nil)
	s.NoError(err)

	client := http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var paginationPhone models.PaginationPhone
	errUnmarshal := json.Unmarshal(byteBody, &paginationPhone)
	s.Nil(errUnmarshal)
	s.NotNil(paginationPhone)
	s.NotEqual(paginationPhone.Phones[0], models.Phone{})
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(10, len(paginationPhone.Phones))
}

func (s *SuiteTest) TestGetPhones_WithStateFilter() {
	s.seedMultipleCustomers()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/phones", baseURL), nil)
	s.NoError(err)

	queryParam := req.URL.Query()
	queryParam.Add("state", "OK")
	req.URL.RawQuery = queryParam.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var paginationPhone models.PaginationPhone
	errUnmarshal := json.Unmarshal(byteBody, &paginationPhone)
	s.Nil(errUnmarshal)
	s.NotNil(paginationPhone)
	s.NotEqual(paginationPhone.Phones[0], models.Phone{})
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(10, len(paginationPhone.Phones))
}

func (s *SuiteTest) TestGetPhones_WithCountryFilter() {
	s.seedMultipleCustomers()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/phones", baseURL), nil)
	s.NoError(err)

	queryParam := req.URL.Query()
	queryParam.Add("country", "Cameroon")
	req.URL.RawQuery = queryParam.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var paginationPhone models.PaginationPhone
	errUnmarshal := json.Unmarshal(byteBody, &paginationPhone)
	s.Nil(errUnmarshal)
	s.NotNil(paginationPhone)
	s.NotEqual(paginationPhone.Phones[0], models.Phone{})
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(10, len(paginationPhone.Phones))
}

func (s *SuiteTest) TestGetPhones_WithPageFilter() {
	s.seedMultipleCustomers()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/phones", baseURL), nil)
	s.NoError(err)

	queryParam := req.URL.Query()
	queryParam.Add("page", "4")
	req.URL.RawQuery = queryParam.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var paginationPhone models.PaginationPhone
	errUnmarshal := json.Unmarshal(byteBody, &paginationPhone)
	s.Nil(errUnmarshal)
	s.NotNil(paginationPhone)
	s.Equal(http.StatusOK, resp.StatusCode)
	s.NotEqual(true, paginationPhone.Pagination.HasNext)
	s.Equal(true, paginationPhone.Pagination.HasPrev)
	s.Equal(41, paginationPhone.Pagination.Total)
	s.Equal(4, paginationPhone.Pagination.TotalPages)
	s.Equal(1, len(paginationPhone.Phones))
}

func (s *SuiteTest) TestGetPhones_WithAllFilters() {
	s.seedMultipleCustomers()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/phones", baseURL), nil)
	s.NoError(err)

	queryParam := req.URL.Query()
	queryParam.Add("state", "OK")
	queryParam.Add("country", "Cameroon")
	queryParam.Add("page", "1")
	req.URL.RawQuery = queryParam.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var paginationPhone models.PaginationPhone
	errUnmarshal := json.Unmarshal(byteBody, &paginationPhone)
	s.Nil(errUnmarshal)
	s.NotNil(paginationPhone)
	s.NotEqual(paginationPhone.Phones[0], models.Phone{})
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(7, len(paginationPhone.Phones))
}

func (s *SuiteTest) TestGetPhones_EmptyList() {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/phones", baseURL), nil)
	s.NoError(err)

	client := http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var paginationPhone models.PaginationPhone
	errUnmarshal := json.Unmarshal(byteBody, &paginationPhone)
	s.Nil(errUnmarshal)
	s.NotNil(paginationPhone)
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(0, len(paginationPhone.Phones))
}
