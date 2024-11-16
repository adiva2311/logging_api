package handlers

import (
	"logging-api/models"
	"logging-api/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Logging API for Automobi")
}

type LoggingHandler struct{
	loggingService services.LoggingService
}

func (handler LoggingHandler) Create(c echo.Context) error{
	type payload struct{
		Log_token     string `json:"log_token"`
		User_id       string `json:"user_id"`
		Action_code   string `json:"action_code"`
		Desc_action   string `json:"desc_action"`
		Subject       string `json:"subject"`
		Result_status string `json:"result_status"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil{
		return err
	}

	result := handler.loggingService.Create(models.Logging{
		Log_token: payloadValidator.Log_token,
		User_id: payloadValidator.User_id,
		Action_code: payloadValidator.Action_code,
		Desc_action: payloadValidator.Desc_action,
		Subject: payloadValidator.Subject,
		Result_status: payloadValidator.Result_status,
	})

	return c.JSON(http.StatusOK, result)
}

func (handler LoggingHandler) Update(c echo.Context) error{
	type payload struct{
		Log_token     string `json:"log_token"`
		User_id       string `json:"user_id"`
		Action_code   string `json:"action_code"`
		Desc_action   string `json:"desc_action"`
		Subject       string `json:"subject"`
		Result_status string `json:"result_status"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil{
		return err
	}

	idString,_ := strconv.Atoi(c.Param("id"))
	result := handler.loggingService.Update(idString, models.Logging{
		Log_token: payloadValidator.Log_token,
		User_id: payloadValidator.User_id,
		Action_code: payloadValidator.Action_code,
		Desc_action: payloadValidator.Desc_action,
		Subject: payloadValidator.Subject,
		Result_status: payloadValidator.Result_status,
	})

	return c.JSON(http.StatusOK, result)
}

func (handler LoggingHandler) Delete(c echo.Context) error{
	idString,_ := strconv.Atoi(c.Param("id"))
	result := handler.loggingService.Delete(idString)

	return c.JSON(http.StatusOK, result)
}

func (handler LoggingHandler) GetAll(c echo.Context) error{
	result := handler.loggingService.GetAll()
	return c.JSON(http.StatusOK, result)
}

func NewLoggingHandler(db *gorm.DB) LoggingHandler{
	service := services.NewLoggingService(db)
	handler := LoggingHandler{
		loggingService: service,
	}
	return handler
}