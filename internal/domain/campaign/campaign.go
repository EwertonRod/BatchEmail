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
	CreatedOn time.Time
	Content   string
	Contact   []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) != 0 {
		return nil, errors.New("contact is required")
	}

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email

	}

	return &Campaign{ID: xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contact:   contacts,
	}, nil

}
