package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-logs/loggers"
	productModel "golang-logs/models"
)

func main() {

	product := productModel.NewProduct()
	formattedProduct, _ := json.Marshal(product)
	stringProduct := string(formattedProduct)

	fmt.Println("====================")
	loggers.SetupLogrus()
	loggers.LogrusError(stringProduct)
	fmt.Println("====================")

	loggers.SetupZerolog() //unit formatterd time
	loggers.ZerologError("error shipping product", map[string]interface{}{
		"product":      product,
		"nil":          nil,
		"error":        "integration error message",
		"errorMessage": errors.New("error calculating shipping cost").Error(),
	})
	fmt.Println("====================")

}
