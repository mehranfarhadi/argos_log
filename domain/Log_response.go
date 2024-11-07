package domain

import (
	"time"
)

type LogResponse struct {
	CameraID      string    `json:"camera_id"`
	Ram           float32   `json:"fan"`
	BoardTemp     float32   `json:"board_temp"`
	Cpu           float32   `json:"cpu"`
	PowerMode     float32   `json:"power_mode"`
	BattV         float32   `json:"batt_V"`
	PanelV        float32   `json:"panel_V"`
	LoadI         float32   `json:"load_I"`
	LoadV         float32   `json:"load_v"`
	ChargeI       float32   `json:"charge_I"`
	ChargeStage   string    `json:"charge_stage"`
	Error         string    `json:"error"`
	BattType      string    `json:"batt_type"`
	ChargerStatus string    `json:"charger_status"`
	temp          float32   `json:"temp"`
	time          time.Time `json:"time"`
	CreatedAt     time.Time `json:"created_at"`
	fps           float32   `json:"fps"`
}

type LogResponseServer struct {
	CameraID      string    `json:"camera_id"`
	Ram           float32   `json:"fan"`
	BoardTemp     float32   `json:"board_temp"`
	Cpu           float32   `json:"cpu"`
	PowerMode     float32   `json:"power_mode"`
	BattV         float32   `json:"batt_V"`
	PanelV        float32   `json:"panel_V"`
	LoadI         float32   `json:"load_I"`
	LoadV         float32   `json:"load_v"`
	ChargeI       float32   `json:"charge_I"`
	ChargeStage   string    `json:"charge_stage"`
	Error         string    `json:"error"`
	BattType      string    `json:"batt_type"`
	ChargerStatus string    `json:"charger_status"`
	temp          float32   `json:"temp"`
	time          time.Time `json:"time"`
	CreatedAt     time.Time `json:"created_at"`
	fps           float32   `json:"fps"`
}
