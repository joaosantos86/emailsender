package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newCampaign(t *testing.T) *Campaign {

	t.Helper()
	campaign, _ := NewCampaign("Campaign X", "body", []string{"email1@gmail.com", "email2@gmail.com"})
	return campaign
}

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign := newCampaign(t)

	assert.Equal("Campaign X", campaign.Name)
	assert.Equal("body", campaign.Content)
	assert.Equal(2, len(campaign.Contacts))
}

func TestNewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	campaign := newCampaign(t)

	assert.NotNil(campaign.ID)
}
func TestNewCampaign_CreatedOnMusBeNow(t *testing.T) {
	assert := assert.New(t)
	campaign := newCampaign(t)

	now := time.Now().Add(-time.Minute)

	assert.Greater(campaign.CreatedOn, now)
	// assert.NotZero(campaign.CreatedOn)
}

func TestNewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", "body", []string{"email1@gmail.com"})

	assert.Equal("Name is Required", err.Error())
}

 func TestNewCampaign_MustValidateContent(t *testing.T) {
 	assert := assert.New(t)
 	_, err := NewCampaign("Campaign X", "", []string{"email1@gmail.com"})

 	assert.Equal("Content is Required", err.Error())
 }

 func TestNewCampaign_MustValidateContacts(t *testing.T) {
 	assert := assert.New(t)
 	_, err := NewCampaign("Campaign X", "body", []string{})

 	assert.Equal("Emails is Required", err.Error())
 }
