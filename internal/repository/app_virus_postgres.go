package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	app "rest_api"
	"strings"
)

type VirusPostgres struct {
	db *sqlx.DB
}

func NewVirusPostgres(db *sqlx.DB) *VirusPostgres {
	return &VirusPostgres{db: db}
}

func (r *VirusPostgres) Create(vir app.Virus) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var virusID int
	createVirusQuery := fmt.Sprintf("INSERT INTO %s (virus_name, virus_type, infection_method, severity) VALUES ($1, $2, $3, $4) RETURNING virus_id", virusesTable)
	row := tx.QueryRow(createVirusQuery, vir.VirusName, vir.VirusType, vir.InfectionMethod, vir.Severity)
	if err := row.Scan(&virusID); err != nil {
		tx.Rollback()
		return 0, nil
	}

	return virusID, tx.Commit()
}

func (r *VirusPostgres) GetAll() ([]app.Virus, error) {
	var viruses []app.Virus

	query := fmt.Sprintf("SELECT * FROM %s", virusesTable)
	err := r.db.Select(&viruses, query)

	return viruses, err
}

func (r *VirusPostgres) GetVirusById(virusId int) (app.Virus, error) {
	var virus app.Virus

	query := fmt.Sprintf("SELECT * FROM %s WHERE virus_id = $1", virusesTable)
	err := r.db.Get(&virus, query, virusId)

	return virus, err
}

func (r *VirusPostgres) DeleteVirus(virusId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE virus_id = $1", virusesTable)
	_, err := r.db.Exec(query, virusId)

	return err
}

func (r *VirusPostgres) Update(virusId int, input app.VirusUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.VirusName != nil {
		setValues = append(setValues, fmt.Sprintf("virus_name = $%d", argId))
		args = append(args, *input.VirusName)
		argId++
	}

	if input.VirusType != nil {
		setValues = append(setValues, fmt.Sprintf("virus_type = $%d", argId))
		args = append(args, *input.VirusType)
		argId++
	}

	if input.Severity != nil {
		setValues = append(setValues, fmt.Sprintf("severity = $%d", argId))
		args = append(args, *input.Severity)
		argId++
	}

	if input.InfectionMethod != nil {
		setValues = append(setValues, fmt.Sprintf("infection_method = $%d", argId))
		args = append(args, *input.InfectionMethod)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE virus_id = $%d",
		virusesTable, setQuery, argId)
	args = append(args, virusId)

	logrus.Debugf("Update query: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
