package campaign

import (
	contract "emailsender/internal/contracts"
)

type Service struct {
	Repositery Repositery
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	s.Repositery.Save(campaign)
	return campaign.ID, nil

}
