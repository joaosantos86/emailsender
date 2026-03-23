package campaign

import "time"

type Contact struct {
	Email string
}

type Campaing struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaing(name string, content string, emails []string) *Campaing {

	contacts := make([]Contact, len(emails))
	for _, email := range emails {
		contacts = append(contacts, Contact{Email: email})
	}


	return &Campaing {
		ID: "1",
		Name: name,
		Content: content,
		CreatedOn: time.Now(),
		Contacts: contacts,
	}
}
