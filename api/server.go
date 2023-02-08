package api

import (
	"encoding/json"
	"go_ecommerce/types"
	"net/http"

	"github.com/anthdm/weavebox"
)

type CreateProductRequest struct {
	SKU         string `json:"sku"`
	ProductName string `json:"product_name"`
}

type ProductStore interface {
	InsertProduct(*types.Product) error
	GetOneProductByID(id string) (*types.Product, error)
}

type ProductHandler struct {
	store ProductStore
}

func NewProductHandler(product_store ProductStore) *ProductHandler {
	return &ProductHandler{
		store: product_store,
	}
}

func (p *ProductHandler) HandleCreateProduct(c *weavebox.Context) error {
	req := &CreateProductRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, req)
}

func (p *ProductHandler) HandleGetProduct(c *weavebox.Context) error {
	return c.JSON(http.StatusOK, &types.Product{SKU: "MAC M1"})
}
