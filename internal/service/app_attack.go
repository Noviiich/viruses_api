package service

import (
	app "rest_api"
	"rest_api/internal/repository"
)

type AttackService struct {
	repo repository.Attack
}

func NewAttackService(repo repository.Attack) *AttackService {
	return &AttackService{
		repo: repo,
	}
}

func (s *AttackService) Create(attack app.Attack) (int, error) {
	return s.repo.Create(attack)
}

func (s *AttackService) GetAll() ([]app.Attack, error) {
	return s.repo.GetAll()
}

func (s *AttackService) GetById(attackId int) (app.Attack, error) {
	return s.repo.GetById(attackId)
}

func (s *AttackService) Delete(attackId int) error {
	return s.repo.Delete(attackId)
}

func (s *AttackService) Update(attackId int, input app.AttackUpdate) error {
	return s.repo.Update(attackId, input)
}
