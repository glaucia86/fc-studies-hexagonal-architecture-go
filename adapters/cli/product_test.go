/**
 * file: adapters/cli/product_test.go
 * description: file responsible for the implementation of unit tests
 * data: 01/23/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package cli_test

import (
	"fmt"
	"testing"

	"github.com/glaucia86/fc-studies-hexagonal-architecture-go/adapters/cli"
	mock_application "github.com/glaucia86/fc-studies-hexagonal-architecture-go/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	// Create
	resultExpected := fmt.Sprintf("Product ID %s with the name %s ha been created with the price %f and status %s",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// Enable
	resultExpected = fmt.Sprintf("Product %s has been enabled.", productName)

	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// Disable
	resultExpected = fmt.Sprintf("Product %s has been disabled.", productName)

	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// Get
	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId, productName, productPrice, productStatus)

	result, err = cli.Run(service, "get", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
