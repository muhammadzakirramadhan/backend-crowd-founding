package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(ID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImages) (CampaignImages, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) FindAll() ([]Campaign, error) {
	var campaign []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r repository) FindByUserID(ID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", ID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r repository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Preload("Users").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil

}

func (r repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r repository) CreateImage(campaignImage CampaignImages) (CampaignImages, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r repository) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(&CampaignImages{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
