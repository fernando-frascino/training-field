package main

import (
	"golang-logs/loggers"
	productModel "golang-logs/models"
)

func main() {

	product := productModel.NewProduct()
	// formattedProduct, _ := json.Marshal(product)
	// stringProduct := string(formattedProduct)

	// loggers.SetupLogrus()
	// loggers.LogrusError(stringProduct)
	// fmt.Println("====================")

	loggers.ZerologError("error shipping product", product, product)
	// fmt.Println("====================")
}
