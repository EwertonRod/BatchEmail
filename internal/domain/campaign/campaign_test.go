package campaign

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign"
	content  = "Body"
	contacts = []string{"email1@gmail.com", "email2@gmail.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {

	// Arrange
	assert := assert.New(t)

	//Action
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.Name, name, "expected %s not %s", name, campaign.Name)
	assert.Equal(campaign.Content, content, "expected %s not %s ", content, campaign.Content)
	assert.Equal(len(campaign.Contact), len(contacts), "is missing emails we expect %s and receive %s", strconv.Itoa(len(contacts)), strconv.Itoa(len(campaign.Contact)))

	/*
		if campaign.ID != "1" {
			t.Errorf("expected %s", campaign.ID)
		} else if campaign.Name != name {
			t.Errorf("expected %s not %s", name, campaign.Name)
		} else if campaign.Content != content {
			t.Errorf("expected %s not %s ", content, campaign.Content)
		} else if len(campaign.Contact) != len(contacts) {
			t.Errorf("is missing emails we expect %s and receive %s", strconv.Itoa(len(contacts)), strconv.Itoa(len(campaign.Contact)))
		}]
	*/
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {

	// Arrange
	assert := assert.New(t)
	//Action
	campaign, _ := NewCampaign(name, content, contacts)
	assert.NotNil(campaign.ID)

}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {

	// Arrange
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	//Action
	campaign, _ := NewCampaign(name, content, contacts)
	assert.Greater(campaign.CreatedOn, now)

}

func Test_NewCampaign_MustValidateName(t *testing.T) {

	// Arrange
	assert := assert.New(t)

	//Action
	_, error := NewCampaign("", content, contacts)
	assert.Equal("name is required", error.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {

	// Arrange
	assert := assert.New(t)

	//Action
	_, error := NewCampaign(name, content, contacts)
	assert.Equal("content is required", error.Error())
}

func Test_NewCampaign_MustValidateContact(t *testing.T) {

	// Arrange
	assert := assert.New(t)

	//Action
	_, error := NewCampaign(name, content, []string{})
	assert.Equal("contact is required", error.Error())
}
