/**
 * file: adapters/dto/product.go
 * description: file responsible for the dto layer of the application.
 * data: 01/29/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package dto

import "github.com/glaucia86/fc-studies-hexagonal-architecture-go/application"

type Product struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.Id != "" {
		product.ID = p.Id
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}