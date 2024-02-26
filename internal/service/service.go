package service

import (
	app "rest_api"
	"rest_api/internal/repository"
)

type Authorization interface {
	CreateUser(user app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
	Delete(virusId int) error
	Update(virusId int, input app.VirusUpdate) error
}

type Site interface {
	Create(site app.Site) (int, error)
	GetAll() ([]app.Site, error)
	GetById(siteId int) (app.Site, error)
	Delete(virusId int) error
	Update(siteId int, input app.SiteUpdate) error
}

type Service struct {
	Virus
	Site
	Attack
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Virus:         NewVirusService(repo.Virus),
		Site:          NewSiteService(repo.Site),
		Attack:        NewAttackService(repo.Attack),
		Authorization: NewAuthService(repo.Authorization),
	}
}
