package checker

import (
	"net/http"
)

var (
	client *http.Client
	//ipinfo   IPInfo
	ipinfo Ifconfig

	//endpoint = "https://ipinfo.io/json"
	endpoint = "https://ifconfig.me/all.json"
)
