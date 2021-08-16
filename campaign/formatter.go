package campaign

import "strings"

// Struct Formatter
type CampaignForamtter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type CampaignDetailFormatter struct {
	ID               int                         `json:"id"`
	Name             string                      `json:"name"`
	ShortDescription string                      `json:"short_description"`
	Description      string                      `json:"description"`
	ImageURL         string                      `json:"image_url"`
	GoalAmount       int                         `json:"goal_amount"`
	CurrentAmount    int                         `json:"current_amount"`
	UserID           int                         `json:"user_id"`
	Slug             string                      `json:"slug"`
	Perks            []string                    `json:"perks"`
	User             CampaignUserDetailFormatter `json:"user"`
	Images           []CampaignImageFormatter    `json:"images"`
}

type CampaignUserDetailFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

// Fucn Formatter
func FormatCampaign(campaign Campaign) CampaignForamtter {
	campaignFormatter := CampaignForamtter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatterCampaings(campaigns []Campaign) []CampaignForamtter {
	campaignsFormatter := []CampaignForamtter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetail := CampaignDetailFormatter{}
	campaignDetail.ID = campaign.ID
	campaignDetail.Name = campaign.Name
	campaignDetail.ShortDescription = campaign.ShortDescription
	campaignDetail.Description = campaign.Description
	campaignDetail.GoalAmount = campaign.GoalAmount
	campaignDetail.CurrentAmount = campaign.CurrentAmount
	campaignDetail.UserID = campaign.UserID
	campaignDetail.Slug = campaign.Slug
	campaignDetail.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetail.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetail.Perks = perks

	user := campaign.Users
	campaignUser := CampaignUserDetailFormatter{}
	campaignUser.Name = user.Name
	campaignUser.ImageURL = user.AvatarFileName

	campaignDetail.User = campaignUser

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		imageFormatter := CampaignImageFormatter{}
		imageFormatter.ImageURL = image.FileName
		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}

		imageFormatter.IsPrimary = isPrimary
		images = append(images, imageFormatter)
	}

	campaignDetail.Images = images
	return campaignDetail
}
