package handlers

import (
	"gocom/database"
	"gocom/util"
	"net/http"
	"strconv"
)

// http://localhost:8080/products/5  ---> 5 is called as request "path parameter"
func GetProductByID(w http.ResponseWriter, r  *http.Request){  // GET and OPTIONS
	productID := r.PathValue("productId")
	pId, err := strconv.Atoi(productID)
	if err != nil{
		http.Error(w, "Please give me a valid product ID", 400)
		return
	}

	for _, product := range database.ProductList{
       if product.ID == pId{
			util.SendData(w, product, 200)
			return
	   }
	//    log.Println("productid", pId, product)
	}
	util.SendData(w, "No product found", 404)
}