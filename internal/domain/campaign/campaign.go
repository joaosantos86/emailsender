package campaign

import (
	"errors"
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

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("Name is Required")
	} else if content == "" {
		return nil, errors.New("Content is Required")
	} else if len(emails) == 0 {
		return nil, errors.New("Emails is Required")
	}
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
	}, nil
}
