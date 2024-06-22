package createaccount

import (
	"testing"

	"github.com/deirofelippe/curso-fullcycle/internal/entity"
	"github.com/deirofelippe/curso-fullcycle/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUsecase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "john@doe")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUsecase(accountMock, clientMock)
	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := uc.Execute(inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
