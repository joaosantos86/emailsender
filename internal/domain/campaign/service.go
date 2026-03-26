package campaign

import (
	contract "emailsender/internal/contracts"
)

type Service struct {
	Repositery Repositery
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
    campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
    if err != nil {
        return "", err
    }
	err = s.Repositery.Save(campaign)
    if err := s.Repositery.Save(campaign); err != nil {
        return "", err
    }

    return campaign.ID, nil
}