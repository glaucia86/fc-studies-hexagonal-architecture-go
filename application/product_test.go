/**
 * file: application/product_test.go
 * description: file responsible for the
 * data: 08/01/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package application_test

import (
	"testing"

	"github.com/glaucia86/fc-studies-hexagonal-architecture-go/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test-Product"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}
