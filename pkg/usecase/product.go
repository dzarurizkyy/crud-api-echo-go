package usecase

import (
	"crud_product/pkg/dto"
	"crud_product/pkg/model"
	"errors"

	"github.com/mitchellh/mapstructure"
)

// Struct
type ProductUsecase struct {
	ProductRepository model.ProductRepository;
};

// Contract or dependency between two layer
func NewProductUsecase(pr model.ProductRepository) model.ProductUseCase {
	return &ProductUsecase{
		ProductRepository: pr,
	};
};

// Business Logic
func (pu ProductUsecase) GetAllProduct() ([]model.Product, error) {
	return pu.ProductRepository.GetAllProduct();
};  

func (pu ProductUsecase) GetProductById(id int) (model.Product, error) {
	response, err := pu.ProductRepository.GetProductById(id);

	if (err != nil) {
		return model.Product{}, errors.New("product id not found");
	};

	return response, nil;
};

func (pu ProductUsecase) CreateProduct(req dto.ProductRequest) (error) {
	var product model.Product;
	mapstructure.Decode(req, &product);
	
	return pu.ProductRepository.CreateProduct(product);
};

func (pu ProductUsecase) UpdateProduct(id int, req dto.ProductRequest) (error) {
	_, err := pu.GetProductById(id);

	if (err != nil) {
		return errors.New("product id not found");
	};

	// ---------------------------------------
	
	var product model.Product;
	mapstructure.Decode(req, &product);	

	return pu.ProductRepository.UpdateProduct(id, product);
};

func (pu ProductUsecase) DeleteProduct(id int) (error) {
	_, err := pu.GetProductById(id);

	if (err != nil) {
		return errors.New("product id not found");
	};

	// ---------------------------------------

	return pu.ProductRepository.DeleteProduct(id);
};