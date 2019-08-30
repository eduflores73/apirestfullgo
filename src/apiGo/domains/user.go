package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type User struct {
	Address struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
	BuyerReputation struct {
		Tags []interface{} `json:"tags"`
	} `json:"buyer_reputation"`
	CountryID        string      `json:"country_id"`
	ID               int         `json:"id"`
	Logo             interface{} `json:"logo"`
	Nickname         string      `json:"nickname"`
	Permalink        string      `json:"permalink"`
	Points           int         `json:"points"`
	RegistrationDate string      `json:"registration_date"`
	SellerReputation struct {
		LevelID           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Canceled  int    `json:"canceled"`
			Completed int    `json:"completed"`
			Period    string `json:"period"`
			Ratings   struct {
				Negative int `json:"negative"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"ratings"`
			Total int `json:"total"`
		} `json:"transactions"`
	} `json:"seller_reputation"`
	SiteID string `json:"site_id"`
	Status struct {
		SiteStatus string `json:"site_status"`
	} `json:"status"`
	Tags     []string `json:"tags"`
	UserType string   `json:"user_type"`
}


//resiver

func (user *User) Get() *utils.ApiError {
	if user.ID <= 0{
		return &utils.ApiError{
			Messege: "Site ID is empty",
			Status: http.StatusBadRequest,
		}
	}
	url := fmt.Sprintf("%s%d",utils.UrlUsers,user.ID)

	response, err := http.Get(url)
	if err != nil{
		return &utils.ApiError{
			Messege: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil{
		return &utils.ApiError{
			Messege: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data,&user); err != nil{
		return &utils.ApiError{
			Messege: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	return nil
}

