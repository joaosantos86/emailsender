package campaign

import (
	contract "emailsender/internal/contracts"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositeryMock struct {
	mock.Mock
}

func (r *repositeryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var newCampaignDto = contract.NewCampaign{
	Name:    "Test Y",
	Content: "Body",
	Emails:  []string{"teste@gmail.com"},
}

func Test_Create_Campaign(t *testing.T) {
	repository := new(repositeryMock)
	repository.On("Save", mock.Anything).Return(nil)
	service := Service{Repositery: repository}
	assert := assert.New(t)

	id, err := service.Create(newCampaignDto)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repository := new(repositeryMock) //
	service := Service{Repositery: repository}

	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == newCampaignDto.Name &&
			campaign.Content == newCampaignDto.Content &&
			len(campaign.Contacts) == len(newCampaignDto.Emails)
	})).Return(nil)

	service.Create(newCampaignDto)

	repository.AssertExpectations(t)
}
func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositeryMock) 
	service := Service{Repositery: repository}

	repository.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repositery = repository
	service.Create(newCampaignDto)
	_, err := service.Create(newCampaignDto)
	assert.Equal("error to save on database", err.Error())
}
