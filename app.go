package app

import "time"

type Virus struct {
	VirusID         int    `db:"virus_id" json:"virus_id"`
	VirusName       string `json:"virus_name" db:"virus_name"`
	VirusType       string `json:"virus_type" db:"virus_type"`
	InfectionMethod string `json:"infection_method" db:"infection_method"`
	Severity        string `json:"severity" db:"severity"`
}

type Site struct {
	SiteID   int       `json:"site_id" db:"site_id"`
	SiteName string    `json:"site_name" db:"site_name"`
	HackDate time.Time `json:"hack_date" db:"hack_date"`
	VirusID  int       `json:"virus_id" db:"virus_id"`
}

type VirusUpdate struct {
	VirusName       *string `json:"virus_name"`
	VirusType       *string `json:"virus_type"`
	InfectionMethod *string `json:"infection_method"`
	Severity        *string `json:"severity"`
}

type SiteUpdate struct {
	VirusID  *int       `json:"virus_id"`
	SiteName *string    `json:"site_name"`
	HackDate *time.Time `json:"hack_date"`
}
