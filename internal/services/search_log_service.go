package services

import (
	"github.com/gabriel-hawerroth/capitech-back/internal/dto"
	"github.com/gabriel-hawerroth/capitech-back/internal/infra/database/repositories"
)

type SearchLogService struct {
	SearchLogRepository repositories.SearchLogRepository
}

func NewSearchLogService(repository repositories.SearchLogRepository) *SearchLogService {
	return &SearchLogService{
		SearchLogRepository: repository,
	}
}

func (s *SearchLogService) Save(dto dto.SaveSearchLogDTO) error {
	return s.SearchLogRepository.Save(dto)
}

func (s *SearchLogService) SaveWithUser(dto dto.SaveSearchLogWithUserDTO) error {
	return s.SearchLogRepository.SaveWithUser(dto)
}
