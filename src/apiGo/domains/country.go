package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"sync"
)

type Country struct {
	CurrencyID       string `json:"currency_id"`
	DecimalSeparator string `json:"decimal_separator"`
	GeoInformation struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	ID     string `json:"id"`
	Locale string `json:"locale"`
	Name   string `json:"name"`
	States []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"states"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
}

func (country *Country) Get() *utils.ApiError {
	if country.ID == "" {
		return &utils.ApiError{
			Messege: "Site ID is empty",
			Status:  http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s", utils.UrlCountries, country.ID)

	response, err := http.Get(url)
	if err != nil {
		return &utils.ApiError{
			Messege: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Messege: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil {
		return &utils.ApiError{
			Messege: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (country *Country) GetWg(wg *sync.WaitGroup, apiError *utils.ApiError){
	if country.ID == "" {
		wg.Done()
		apiError.Messege = "Site ID is empty"
		apiError.Status = http.StatusBadRequest
		return
	}
	url := fmt.Sprintf("%s%s", utils.UrlCountries, country.ID)

	response, err := http.Get(url)
	if err != nil {
		wg.Done()
		apiError.Messege = err.Error()
		apiError.Status = http.StatusInternalServerError
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		wg.Done()
		apiError.Messege = err.Error()
		apiError.Status = http.StatusInternalServerError
		return
	}

	if err := json.Unmarshal(data, &country); err != nil {
		wg.Done()
		apiError.Messege = err.Error()
		apiError.Status = http.StatusInternalServerError
		return
	}
	wg.Done()
	return
}

func (country *Country) GetCH(ch chan Result) {
	if country.ID == "" {
		ch <- Result{
			Error: &utils.ApiError{
				Messege: "Site ID is empty",
				Status:  http.StatusBadRequest,
			},
		}
		return
	}
	url := fmt.Sprintf("%s%s", utils.UrlCountries, country.ID)

	response, err := http.Get(url)
	if err != nil {
		ch <- Result{
			Error: &utils.ApiError{
				Messege: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ch <- Result{
			Error: &utils.ApiError{
				Messege: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
		return
	}

	if err := json.Unmarshal(data, &country); err != nil {
		ch <- Result{
			Error: &utils.ApiError{
				Messege: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
		return
	}

	ch <- Result{
		Country: country,
	}
	return

}
