package api

import (
	"encoding/json"
	"go_ecommerce/types"
	"net/http"

	"github.com/anthdm/weavebox"
)

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
	var product *types.Product
	if err := json.NewDecoder(c.Request().Body()).Decode(product); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &types.Product{SKU: "MAC M1"})
}

func (p *ProductHandler) HandleGetProduct(c *weavebox.Context) error {
	return c.JSON(http.StatusOK, &types.Product{SKU: "MAC M1"})
}
