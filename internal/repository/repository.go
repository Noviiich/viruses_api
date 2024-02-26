package repository

import (
	"github.com/jmoiron/sqlx"
	app "rest_api"
)

type Repository struct {
	Virus
	Site
	Attack
	Authorization
}

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GetUser(username, password string) (app.User, error)
}

type Attack interface {
	Create(attack app.Attack) (int, error)
	GetAll() ([]app.Attack, error)
	GetById(attackId int) (app.Attack, error)
	Delete(attackId int) error
	Update(attackId int, input app.AttackUpdate) error
}

type Virus interface {
	Create(vir app.Virus) (int, error)
	GetAll() ([]app.Virus, error)
	GetVirusById(virusId int) (app.Virus, error)
	DeleteVirus(virusId int) error
	Update(virusId int, input app.VirusUpdate) error
}

type Site interface {
	Create(site app.Site) (int, error)
	GetAll() ([]app.Site, error)
	GetById(siteId int) (app.Site, error)
	DeleteSite(siteId int) error
	Update(siteId int, input app.SiteUpdate) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Virus:         NewVirusPostgres(db),
		Site:          NewSitePostgres(db),
		Attack:        NewAttackPostgres(db),
		Authorization: NewAuthPostgres(db),
	}
}
