package transaction

import (
	"backend-crowd-funding/campaign"
	"backend-crowd-funding/users"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	Users      users.Users       `gorm:"foreignKey:UserID"`
	Campaign   campaign.Campaign `gorm:"foreignKey:CampaignID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
