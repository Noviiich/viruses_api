package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	app "rest_api"
	"strings"
)

type AttackPostgres struct {
	db *sqlx.DB
}

func NewAttackPostgres(db *sqlx.DB) *AttackPostgres {
	return &AttackPostgres{db: db}
}

func (r *AttackPostgres) Create(attack app.Attack) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var attackId int
	createAttackQuery := fmt.Sprintf("INSERT INTO %s (site_id, virus_id, hack_date) VALUES ($1, $2, $3) RETURNING id", attackListTable)
	row := tx.QueryRow(createAttackQuery, attack.SiteID, attack.VirusID, attack.HackDate)
	if err := row.Scan(&attackId); err != nil {
		tx.Rollback()
		return 0, nil
	}

	return attackId, tx.Commit()
}

func (r *AttackPostgres) GetAll() ([]app.Attack, error) {
	var attacks []app.Attack
	query := fmt.Sprintf("SELECT * FROM %s", attackListTable)
	err := r.db.Select(&attacks, query)

	return attacks, err
}

func (r *AttackPostgres) GetById(attackId int) (app.Attack, error) {
	var attack app.Attack

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", attackListTable)
	err := r.db.Get(&attack, query, attackId)

	return attack, err

}

func (r *AttackPostgres) Delete(attackId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", attackListTable)
	_, err := r.db.Exec(query, attackId)

	return err
}

func (r *AttackPostgres) Update(attackId int, input app.AttackUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.SiteID != nil {
		setValues = append(setValues, fmt.Sprintf("site_id = $%d", argId))
		args = append(args, *input.SiteID)
		argId++
	}

	if input.VirusID != nil {
		setValues = append(setValues, fmt.Sprintf("virus_id = $%d", argId))
		args = append(args, *input.VirusID)
		argId++
	}

	if input.HackDate != nil {
		setValues = append(setValues, fmt.Sprintf("hack_date = $%d", argId))
		args = append(args, *input.HackDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", attackListTable, setQuery, argId)
	args = append(args, attackId)

	logrus.Debugf("Update query: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
