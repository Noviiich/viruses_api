package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	app "rest_api"
	"strings"
)

type SitePostgres struct {
	db *sqlx.DB
}

func NewSitePostgres(db *sqlx.DB) *SitePostgres {
	return &SitePostgres{db: db}
}

func (r *SitePostgres) Create(site app.Site) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var siteId int
	createSiteQuery := fmt.Sprintf("INSERT INTO %s (site_name, hack_date, virus_id) VALUES ($1, $2, $3) RETURNING site_id", sitesTable)
	row := tx.QueryRow(createSiteQuery, site.SiteName, site.HackDate, site.VirusID)
	if err := row.Scan(&siteId); err != nil {
		tx.Rollback()
		return 0, nil
	}

	return siteId, tx.Commit()
}

func (r *SitePostgres) GetAll() ([]app.Site, error) {
	var sites []app.Site

	query := fmt.Sprintf("SELECT * FROM %s", sitesTable)
	err := r.db.Select(&sites, query)

	return sites, err
}

func (r *SitePostgres) GetById(siteId int) (app.Site, error) {
	var site app.Site

	query := fmt.Sprintf("SELECT * FROM %s WHERE site_id = $1", sitesTable)
	err := r.db.Get(&site, query, siteId)

	return site, err
}

func (r *SitePostgres) DeleteSite(siteId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE site_id = $1", sitesTable)
	_, err := r.db.Exec(query, siteId)

	return err
}

func (r *SitePostgres) Update(siteId int, input app.SiteUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.SiteName != nil {
		setValues = append(setValues, fmt.Sprintf("site_name = $%d", argId))
		args = append(args, *input.SiteName)
		argId++
	}

	if input.HackDate != nil {
		setValues = append(setValues, fmt.Sprintf("hack_date = $%d", argId))
		args = append(args, *input.HackDate)
		argId++
	}

	if input.VirusID != nil {
		setValues = append(setValues, fmt.Sprintf("virus_id = $%d", argId))
		args = append(args, *input.VirusID)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE site_id = $%d",
		sitesTable, setQuery, argId)
	args = append(args, siteId)

	logrus.Debugf("Update query: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
