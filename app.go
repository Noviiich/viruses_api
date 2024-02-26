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
	SiteID        int    `json:"site_id" db:"site_id"`
	SiteName      string `json:"site_name" db:"site_name"`
	SecurityLevel string `json:"security_level" db:"security_level"`
	OwnerContact  string `json:"owner_contact" db:"owner_contact"`
}

type VirusUpdate struct {
	VirusName       *string `json:"virus_name"`
	VirusType       *string `json:"virus_type"`
	InfectionMethod *string `json:"infection_method"`
	Severity        *string `json:"severity"`
}

type SiteUpdate struct {
	SiteName      *string `json:"site_name"`
	SecurityLevel *string `json:"security_level"`
	OwnerContact  *string `json:"owner_contact"`
}

type Attack struct {
	AttackId int       `db:"id" json:"id"`
	VirusID  int       `db:"virus_id" json:"virus_id"`
	SiteID   int       `json:"site_id" db:"site_id"`
	HackDate time.Time `json:"hack_date" db:"hack_date"`
}

type AttackUpdate struct {
	VirusID  *int       `db:"virus_id" json:"virus_id"`
	SiteID   *int       `json:"site_id" db:"site_id"`
	HackDate *time.Time `json:"hack_date" db:"hack_date"`
}
