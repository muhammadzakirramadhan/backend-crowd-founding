package transaction

import "backend-crowd-funding/users"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User users.Users
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       users.Users
}
