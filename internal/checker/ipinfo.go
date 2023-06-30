package checker

type IPInfo struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Readme   string `json:"readme"`
	Region   string `json:"region"`
	Timezone string `json:"timezone"`
}

type Ifconfig struct {
	IP        string `json:"ip_addr"`
	Country   string `json:"remote_host"`
	UserAgent string `json:"user_agent"`
	Port      int    `json:"port"`
	Language  string `json:"language"`
	Method    string `json:"method"`
	Encoding  string `json:"encoding"`
	Mime      string `json:"mime"`
	Via       string `json:"via"`
	Forwarded string `json:"forwarded"`
}
