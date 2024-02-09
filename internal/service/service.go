package service

import (
	app "rest_api"
	"rest_api/internal/repository"
)

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
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Virus: NewVirusService(repo.Virus),
		Site:  NewSiteService(repo.Site),
	}
}
