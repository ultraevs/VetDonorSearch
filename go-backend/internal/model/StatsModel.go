package model

type ResponseStat struct {
	Blood     int `json:"blood"`
	Plasma    int `json:"plasma"`
	Platelets int `json:"platelets"`
}

type RequestUpdateStat struct {
	Email     string `json:"email"`
	Blood     int    `json:"blood"`
	Plasma    int    `json:"plasma"`
	Platelets int    `json:"platelets"`
}
