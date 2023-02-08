package main

import (
	"fmt"
	"go_ecommerce/api"

	"github.com/anthdm/weavebox"
)

func main() {
	app := weavebox.New()

	adminRoute := app.Box("/admin")

	productHandler := &api.ProductHandler{}

	adminRoute.Get("/api/v1/product", productHandler.HandleGetProduct)

	app.Serve(3550)
	fmt.Println("ECOMMERCE")
}