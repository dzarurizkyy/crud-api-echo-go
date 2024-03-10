package controller

import (
	"crud_product/pkg/dto"
	"crud_product/pkg/model"
	"crud_product/shared/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Struct
type ProductController struct {
	ProductUsecase model.ProductUseCase;
};

// Contract or dependency between two layer
func NewProductController(pu model.ProductUseCase) ProductController {
	return ProductController {
		ProductUsecase: pu,
	};
};

// Communicate between client on system
func (pc ProductController) GetAllProduct(c echo.Context) (error) {
	response, err := pc.ProductUsecase.GetAllProduct();
	
	if (err != nil) {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil);
	};

	return utils.SetResponse(c, http.StatusOK, "success", response);
};

func (pc ProductController) GetProductById(c echo.Context) (error) {
	id, _ := strconv.Atoi(c.Param("id"));
	response, err := pc.ProductUsecase.GetProductById(id);

	if (err != nil) {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil);
	};

	return utils.SetResponse(c, http.StatusOK, "success", response);
};

func (pc ProductController) CreateProduct(c echo.Context) (error) {
	var request dto.ProductRequest;
	c.Bind(&request);

	err := request.Validation();

	if (err != nil) {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil);
	};

	// -------------------------------------------------
	
	err = pc.ProductUsecase.CreateProduct(request);

	if (err != nil) {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil);	
	};

	return utils.SetResponse(c, http.StatusOK, "success", nil);
};

func (pc ProductController) UpdateProduct(c echo.Context) (error) {
	id, _ := strconv.Atoi(c.Param("id"));
	var request dto.ProductRequest;
	c.Bind(&request);

	err := request.Validation();

	if (err != nil) {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil);
	};

	// -------------------------------------------------

	err = pc.ProductUsecase.UpdateProduct(id, request);
	
	if (err != nil) {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil);
	};

	return utils.SetResponse(c, http.StatusOK, "success", nil);
};

func (pc ProductController) DeleteProduct(c echo.Context) (error) {
	id, _ := strconv.Atoi(c.Param("id"));
	
	err := pc.ProductUsecase.DeleteProduct(id);

	if (err != nil) {
		return utils.SetResponse(c, http.StatusInternalServerError, err.Error(), nil);
	};

	return utils.SetResponse(c, http.StatusOK, "success", nil);
};