package usecase

import (
	"golang-api/model"
	"golang-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(product model.Product) (*model.Product, error) {
	updatedProduct, err := pu.repository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (pu *ProductUsecase) DeleteProduct(id int) (*model.Product, error) {
	product, err := pu.repository.DeleteProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
