package utils

import "time"

var (
	/* url mercado libre
	UrlSite = "https://api.mercadolibre.com/sites/"
	UrlUsers = "https://api.mercadolibre.com/countries/"
	UrlCountries = "https://api.mercadolibre.com/users/"

	url mock */

	UrlSite = "http://localhost:8081/sites/"
	UrlUsers = "http://localhost:8081/users/"
	UrlCountries = "http://localhost:8081/countries/"

	MaxRequests = 3
	Interval    time.Duration
	Timeout     time.Duration

)

