package dto

import validation "github.com/go-ozzo/ozzo-validation"

// Struct
type ProductRequest struct {
	ImageUrl    string  `json:"image_url"`;
	Name        string  `json:"name"`;
	Description string  `json:"description"`;
	Price       float64 `json:"price"`;
	Stock       int     `json:"stock"`;
};

// Request Validator 
func (pr ProductRequest) Validation() error {
	err := validation.ValidateStruct(
		&pr, 
		validation.Field(&pr.Name, validation.Required),
		validation.Field(&pr.Description, validation.Required),
		validation.Field(&pr.Price, validation.Required),
	);

	if err != nil {
		return err;
	};

	return nil;
};