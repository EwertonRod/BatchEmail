package campaign

import "batchemail/internal/contract"

type Service struct {
	Repository Repository
}

func (s Service) Create(newcampaign contract.NewCampaign) (string, error) {

	campaign, _ := NewCampaign(newcampaign.Name, newcampaign.Content, newcampaign.Emails)

	return campaign.ID, nil

}
