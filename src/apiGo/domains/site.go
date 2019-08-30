package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"sync"
)

type Site struct {
	Categories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
	CountryID string `json:"country_id"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	DefaultCurrencyID  string   `json:"default_currency_id"`
	ID                 string   `json:"id"`
	ImmediatePayment   string   `json:"immediate_payment"`
	MercadopagoVersion int      `json:"mercadopago_version"`
	Name               string   `json:"name"`
	PaymentMethodIds   []string `json:"payment_method_ids"`
	SaleFeesMode       string   `json:"sale_fees_mode"`
	Settings struct {
		IdentificationTypes []string `json:"identification_types"`
		IdentificationTypesRules []struct {
			IdentificationType string `json:"identification_type"`
			Rules []struct {
				BeginsWith           string   `json:"begins_with"`
				EnabledTaxpayerTypes []string `json:"enabled_taxpayer_types"`
				MaxLength            int      `json:"max_length"`
				MinLength            int      `json:"min_length"`
				Type                 string   `json:"type"`
			} `json:"rules"`
		} `json:"identification_types_rules"`
		TaxpayerTypes []string `json:"taxpayer_types"`
	} `json:"settings"`
}

//resiver

func (site *Site) Get() *utils.ApiError {
	if site.ID == "" {
		return &utils.ApiError{
			Messege: "Site ID is empty",
			Status:  http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

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

	if err := json.Unmarshal(data, &site); err != nil {
		return &utils.ApiError{
			Messege: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (site *Site) GetWG(wg *sync.WaitGroup, apiError *utils.ApiError) {
	if site.ID == "" {
		wg.Done()
		apiError.Messege = "Site ID is empty"
		apiError.Status = http.StatusBadRequest
		return
	}
	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

	response, err := http.Get(url)
	if err != nil {
		wg.Done()
		apiError.Messege = err.Error()
		apiError.Status = http.StatusBadRequest
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		wg.Done()
		apiError.Messege = err.Error()
		apiError.Status = http.StatusBadRequest
		return
	}

	if err := json.Unmarshal(data, &site); err != nil {
		wg.Done()
		apiError.Messege = err.Error()
		apiError.Status = http.StatusBadRequest
		return
	}
	wg.Done()
	return
}

func (site *Site) GetCH(ch chan Result) {
	if site.ID == "" {
		ch <- Result{
			Error: &utils.ApiError{
				Messege: "Site ID is empty",
				Status:  http.StatusBadRequest,
			},
		}
		return
	}
	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

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

	if err := json.Unmarshal(data, &site); err != nil {
		ch <- Result{
			Error: &utils.ApiError{
				Messege: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
		return
	}
	ch <- Result{
		Site: site,
	}
	return

}
