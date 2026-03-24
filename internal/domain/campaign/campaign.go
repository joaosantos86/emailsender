package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	Content   string
	Contacts  []Contact
	CreatedOn time.Time
}

func NewCampaign(name string, content string, emails []string) *Campaign {

	contacts := make([]Contact, 0, len(emails))
	for _, email := range emails {
		contacts = append(contacts, Contact{Email: email})
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}
}
