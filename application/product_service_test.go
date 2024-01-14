/**
 * file: application/product_service_test.go
 * description: file responsible for testing the application layer of the application.
 * data: 14/01/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package application_test

import (
	"testing"

	"github.com/glaucia86/fc-studies-hexagonal-architecture-go/application"
	mock_application "github.com/glaucia86/fc-studies-hexagonal-architecture-go/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
