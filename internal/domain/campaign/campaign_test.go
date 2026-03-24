package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newCampaign(t *testing.T) *Campaign {
	t.Helper()
	return NewCampaign("Campaign X", "body", []string{"email1@gmail.com", "email2@gmail.com"})
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
