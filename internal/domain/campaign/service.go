package campaign

import contract "emailsender/internal/contracts"

type Service struct {
	Repositery Repositery
}

func (s *Service) Create(newCampaign contract.NewCampaign) error {
	return  nil
}
