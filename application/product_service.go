/**
 * file: application/product_service.go
 * description: file responsible for the application layer of the application.
 * data: 14/01/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package application

// Aqui é uma injecão de dependência
type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
