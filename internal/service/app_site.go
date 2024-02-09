package service

import (
	app "rest_api"
	"rest_api/internal/repository"
)

type SiteService struct {
	repo repository.Site
}

func NewSiteService(repo repository.Site) *SiteService {
	return &SiteService{
		repo: repo,
	}
}

func (s *SiteService) Create(site app.Site) (int, error) {
	return s.repo.Create(site)
}

func (s *SiteService) GetAll() ([]app.Site, error) {
	return s.repo.GetAll()
}

func (s *SiteService) GetById(siteId int) (app.Site, error) {
	return s.repo.GetById(siteId)
}

func (s *SiteService) Delete(siteId int) error {
	return s.repo.DeleteSite(siteId)
}

func (s *SiteService) Update(siteId int, input app.SiteUpdate) error {
	return s.repo.Update(siteId, input)
}
