package main

type DailyStatisticsTotals struct {
	Date    string `json:"date"`
	RunTime int64  `json:"runtime"`
	Points  int64  `json:"points"`
	Results int64  `json:"results"`
}

type TeamList struct {
	Teams []TeamInfo `json:"teams"`
}

type TeamInfo struct {
	TeamID string `json:"teamid"`
}
