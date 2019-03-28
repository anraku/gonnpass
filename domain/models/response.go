package models

import "time"

// generate by https://mholt.github.io/json-to-go/
type ConnpassResponse struct {
	ResultsReturned int `json:"results_returned"`
	Events          []struct {
		EventURL      string `json:"event_url"`
		EventType     string `json:"event_type"`
		OwnerNickname string `json:"owner_nickname"`
		Series        struct {
			URL   string `json:"url"`
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"series"`
		UpdatedAt        time.Time `json:"updated_at"`
		Lat              string    `json:"lat"`
		StartedAt        time.Time `json:"started_at"`
		HashTag          string    `json:"hash_tag"`
		Title            string    `json:"title"`
		EventID          int       `json:"event_id"`
		Lon              string    `json:"lon"`
		Waiting          int       `json:"waiting"`
		Limit            int       `json:"limit"`
		OwnerID          int       `json:"owner_id"`
		OwnerDisplayName string    `json:"owner_display_name"`
		Description      string    `json:"description"`
		Address          string    `json:"address"`
		Catch            string    `json:"catch"`
		Accepted         int       `json:"accepted"`
		EndedAt          time.Time `json:"ended_at"`
		Place            string    `json:"place"`
	} `json:"events"`
	ResultsStart     int `json:"results_start"`
	ResultsAvailable int `json:"results_available"`
}
