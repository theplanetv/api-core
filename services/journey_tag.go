package services

import (
	"api-core/models"
)

type JourneyTagService struct {
	DatabaseService
}

func (s *JourneyTagService) GetAll() []models.JourneyTag {
	s.Open()

	tags := []models.JourneyTag{}

	s.DB.Table("journey_tag").Scan(&tags)

	return tags
}

