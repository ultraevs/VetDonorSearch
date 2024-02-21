package model

type ResponseStat struct {
	Blood     int `json:"blood"`
	Plasma    int `json:"plasma"`
	Platelets int `json:"platelets"`
}

type RequestUpdateStat struct {
	DonationType string `json:"donationType"`
	Email        string `json:"email"`
}
