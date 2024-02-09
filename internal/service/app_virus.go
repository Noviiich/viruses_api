package service

import (
	app "rest_api"
	"rest_api/internal/repository"
)

type VirusService struct {
	repo repository.Virus
}

func NewVirusService(repo repository.Virus) *VirusService {
	return &VirusService{repo: repo}
}

func (s *VirusService) Create(vir app.Virus) (int, error) {
	return s.repo.Create(vir)
}

func (s *VirusService) GetAll() ([]app.Virus, error) {
	return s.repo.GetAll()
}

func (s *VirusService) GetVirusById(virusId int) (app.Virus, error) {
	return s.repo.GetVirusById(virusId)
}

func (s *VirusService) Delete(virusId int) error {
	return s.repo.DeleteVirus(virusId)
}

func (s *VirusService) Update(virusId int, input app.VirusUpdate) error {
	return s.repo.Update(virusId, input)
}
