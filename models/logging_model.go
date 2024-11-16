package models

import "time"

type Logging struct {
	Id            int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Log_token     string `json:"log_token" gorm:"column:log_token"`
	User_id       string `json:"user_id" gorm:"column:user_id"`
	Action_code   string `json:"action_code" gorm:"column:action_code"`
	Desc_action   string `json:"desc_action" gorm:"column:desc_action"`
	Subject       string `json:"subject" gorm:"column:subject"`
	Result_status string `json:"result_status" gorm:"column:result_status"`
	Time          time.Time
}

func (Logging) TableName() string {
	return "logging.user"
}