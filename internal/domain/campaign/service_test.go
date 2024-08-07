package campaign

import (
	"batchemail/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Campaign(t *testing.T) {

	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Teste Campaign",
		Content: "Body",
		Emails:  []string{"email@gmail.com"},
	}

	id, err := service.Create(newCampaign)

	assert.Nil(err)
	assert.NotNil(id)
}
