package domains

import (
	"../utils"
)

type Result struct {
	User 		*User
	Country 	*Country
	Site 		*Site
	Error		*utils.ApiError
}



