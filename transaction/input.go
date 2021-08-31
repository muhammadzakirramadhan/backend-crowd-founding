package transaction

import "backend-crowd-funding/users"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User users.Users
}
