package main

import (
	"encoding/json"
	"fmt"
)

type MymarketProduct struct {
	Name  string `json:"name"`
	Qty   string `json:"qty"`
	Owner string `json:"owner"`
}

type ProductObject struct {
	Key    string          `json:"key"`
	Record MymarketProduct `json:"record"`
}

type ProductList struct {
	ProductList []ProductObject `json:"productList"`
}

type ProductKey struct {
	Key string `json:"key"`
	Idx int    `json:"idx"`
}

func main() {


	var marketProduct = MymarketProduct{Name:"MJ100",Qty:"300",Owner:"MJ"}
	var productObject = ProductObject{Key:"PD01", Record: marketProduct}

	var marketProduct2 = MymarketProduct{Name:"MJ101",Qty:"300",Owner:"MJ"}
	var productObject2 = ProductObject{Key:"PD02", Record: marketProduct2}

	x := make([]ProductObject,0)
	x = append(x, productObject)
	x = append(x, productObject2)

	var productList = ProductList{ProductList: x}

	productListAsBytes, _ := json.Marshal(productList)

	fmt.Println(string(productListAsBytes))


	var p = "{\"name\":\"MJ100\",\"qty\":\"300\",\"owner\":\"MJ\"}"

	var mP = MymarketProduct{}
	json.Unmarshal([]byte(p), &mP)


}

