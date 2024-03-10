package repository

import (
	"crud_product/pkg/model"
	"database/sql"
	"time"
)

// Struct
type ProductRepository struct {
	db *sql.DB;
};

// Contract or dependency between two layer
func NewProductRepository(db *sql.DB) model.ProductRepository {
	return &ProductRepository{
		db: db,
	};
};

// Get All Product
func (pr *ProductRepository) GetAllProduct() ([]model.Product, error) {
	var products []model.Product;

	sql := "SELECT * FROM product ORDER BY id ASC";
  rows, err := pr.db.Query(sql); 
	
	if (err != nil) {
		return nil, err;
	};

	for rows.Next() {
		var product model.Product
		err = rows.Scan(
			&product.Id,
			&product.ImageUrl,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt)
		
		if (err != nil) {
			return nil, err;
		};

		products = append(products, product);
	};

	return products, nil;
};

// Get Product By Id
func(pr ProductRepository) GetProductById(id int) (model.Product, error) {
	var product model.Product;

	sql := `SELECT * FROM product WHERE id = $1`;
	err := pr.db.QueryRow(sql, id).Scan(
		&product.Id,
		&product.ImageUrl,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt);

	if (err != nil) {
		return product, err;
	};

	return product, nil;
};

// Create Product
func (pr *ProductRepository) CreateProduct(req model.Product) (error) {
	sql := `INSERT INTO product (
					image_url,
					name,
					description,
					price,
					stock
					) VALUES ($1, $2, $3, $4, $5)`;

	_, err := pr.db.Exec(
		sql, 
		req.ImageUrl, 
		req.Name,
		req.Description,
		req.Price,
		req.Stock,
	);
	
	if (err != nil) {
		return err;
	};

	return nil;
};

// Update Product By Id
func (pr *ProductRepository) UpdateProduct(id int, req model.Product) (error) {
	sql := `UPDATE product SET 
				  image_url = $1,
					name = $2,
					description = $3,
					price = $4,
					stock = $5,
					updated_at = $6 
					WHERE id = $7`;

	_, err := pr.db.Exec(
		sql,
		req.ImageUrl,
		req.Name,
		req.Description,
		req.Price,
		req.Stock,
		time.Now(),
		id,
	);

	if (err != nil) {
		return err;
	};
	
	return nil;
};

// Delete Product By Id
func (pr *ProductRepository) DeleteProduct(id int) (error) {
	sql := `DELETE FROM product WHERE id = $1`;

	_, err := pr.db.Exec(sql, id);

	if (err != nil) {
		return err;	
	};

	return nil;
};