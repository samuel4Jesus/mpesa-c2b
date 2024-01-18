package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CreatedAt time.Time
	TXNID     string `gorm:"unique"`
	MSISDN    string
	AMOUNT    int
	CUSTREFID string
	STATUS    string
	MNO       string
}
