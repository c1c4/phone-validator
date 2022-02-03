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

	var phones []models.Phone
	errUnmarshal := json.Unmarshal(byteBody, &phones)
	s.Nil(errUnmarshal)
	s.NotNil(phones)
	s.NotEqual(phones[0], models.Phone{})
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(41, len(phones))
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

	var phones []models.Phone
	errUnmarshal := json.Unmarshal(byteBody, &phones)
	s.Nil(errUnmarshal)
	s.NotNil(phones)
	s.NotEqual(phones[0], models.Phone{})
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(27, len(phones))
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

	var phones []models.Phone
	errUnmarshal := json.Unmarshal(byteBody, &phones)
	s.Nil(errUnmarshal)
	s.NotNil(phones)
	s.NotEqual(phones[0], models.Phone{})
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(10, len(phones))
}

func (s *SuiteTest) TestGetPhones_WithAllFilters() {
	s.seedMultipleCustomers()
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/phones", baseURL), nil)
	s.NoError(err)

	queryParam := req.URL.Query()
	queryParam.Add("state", "OK")
	queryParam.Add("country", "Cameroon")
	req.URL.RawQuery = queryParam.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var phones []models.Phone
	errUnmarshal := json.Unmarshal(byteBody, &phones)
	s.Nil(errUnmarshal)
	s.NotNil(phones)
	s.NotEqual(phones[0], models.Phone{})
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(7, len(phones))
}

func (s *SuiteTest) TestGetPhones_EmptyList() {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/phones", baseURL), nil)
	s.NoError(err)

	client := http.Client{}
	resp, err := client.Do(req)
	s.NoError(err)

	byteBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var phones []models.Phone
	errUnmarshal := json.Unmarshal(byteBody, &phones)
	s.Nil(errUnmarshal)
	s.NotNil(phones)
	s.Equal(http.StatusOK, resp.StatusCode)
	s.Equal(0, len(phones))
}
