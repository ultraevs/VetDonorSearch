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

type TopStat struct {
	Email     string `json:"email"`
	Blood     int    `json:"blood"`
	Plasma    int    `json:"plasma"`
	Platelets int    `json:"platelets"`
}
