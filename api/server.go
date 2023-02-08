package api

import (
	"go_ecommerce/types"
	"net/http"

	"github.com/anthdm/weavebox"
)

type ProductHandler struct {
	// handleGetProduct(c *weavebox.Context) error
}

func (p *ProductHandler) HandleGetProduct(c *weavebox.Context) error {
	return c.JSON(http.StatusOK, &types.Product{SKU: "MAC M1"})
}
