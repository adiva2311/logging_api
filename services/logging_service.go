package services

import (
	"logging-api/helpers"
	"logging-api/models"
	"logging-api/repositories"

	"gorm.io/gorm"
)

type loggingService struct {
	loggingRepo repositories.LoggingRepository
}

// Create implements LoggingService.
func (service *loggingService) Create(logging models.Logging) helpers.Response {
	var response helpers.Response
	if err := service.loggingRepo.Create(logging); err != nil{
		response.Status = 500
		response.Message = "Failed to Create New Log"
	}
	response.Status = 200
	response.Message = "Success to Create New Log"

	return response
}

// Delete implements LoggingService.
func (service *loggingService) Delete(id int) helpers.Response {
	var response helpers.Response
	if err := service.loggingRepo.Delete(id); err != nil{
		response.Status = 500
		response.Message = "Failed to Delete Log"
	}
	response.Status = 200
	response.Message = "Success to Delete Log"

	return response
}

// GetAll implements LoggingService.
func (service *loggingService) GetAll() helpers.Response {
	var response helpers.Response

	data, err := service.loggingRepo.GetAll()
	if err != nil{
		response.Status = 500
		response.Message = "Failed to Get Log"
	} else {
		response.Status = 200
		response.Message = "Success to Get Log"
		response.Data = data
	}

	return response
}

// Update implements LoggingService.
func (service *loggingService) Update(id int, logging models.Logging) helpers.Response {
	var response helpers.Response
	if err := service.loggingRepo.Update(id, logging); err != nil{
		response.Status = 500
		response.Message = "Failed to Update New Log"
	}
	response.Status = 200
	response.Message = "Success to Update New Log"

	return response
}

type LoggingService interface {
	Create(logging models.Logging) helpers.Response
	Update(id int, logging models.Logging) helpers.Response
	Delete(id int) helpers.Response
	GetAll() helpers.Response
}

func NewLoggingService(db *gorm.DB) LoggingService {
	return &loggingService{loggingRepo: repositories.NewLoggingRepository(db)}
}
