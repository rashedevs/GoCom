package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct{
	ID int               `json:"id"`           //add tag
	Title string          `json:"title"`
	Description string    `json:"description"`
	Price float64         `json:"price"`
	ImgURL string         `json:"imageUrl"`
}

// var a int8 = 10
// var b int = 5


var productList []Product

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello.! from go server")
}

func getProducts(w http.ResponseWriter, r  *http.Request){
	handleCors(w)
    handlePreflightRequest(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Please make a GET request.", 400)
		return
	}

	sendData(w, productList, 200)
    // encoder := json.NewEncoder(w)
	// encoder.Encode(productList)
}

func createProduct(w http.ResponseWriter, r *http.Request){
	handleCors(w)
    handlePreflightRequest(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Please make a POST request.", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "Please send a valid json body", 400)
		return
	}
    newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

    sendData(w, newProduct, 201)
}

func handleCors(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Habib")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflightRequest(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int){
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}


func main(){ 

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/create-products", createProduct)

	fmt.Println("Server is running on port :8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil{
		fmt.Println("Error starting the server", err)
	}
}

func init(){

	prd1 := Product{
		ID : 1,
		Title : "Orange",
		Description: "This is orange.",
		Price: 150.00,
		ImgURL: "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRwJf4xPOftZLLgkWjr2eMMumu9XuRdKiGc1eZXFku9OmA4lSymIPBm_vU0bFIM_BjpYOsI_pX7O9c9sRUxfkXxE0N1x_bWSERw7SXHP3A",
	}
	prd2 := Product{
		ID : 2,
		Title : "Apple",
		Description: "This is apple.",
		Price: 120.00,
		ImgURL: "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcSbWqLu3A4mBXzZNi1POVqPrGV4UAuPw3Bl9RwTDjJqjCQrMiid5VgbDGHbAIA0rcFkAWdpjba5sqFdLdIvRXZ27_H_dPxXDpebJuqHiHQ",
	}
	prd3 := Product{
		ID : 3,
		Title : "Banana",
		Description: "This is banana.",
		Price: 40.00,
		ImgURL: "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
	}
	prd4 := Product{
		ID : 4,
		Title : "Grapes",
		Description: "This is grapes.",
		Price: 140.00,
        ImgURL: "https://d2t8nl1y0ie1km.cloudfront.net/images/6706bbb12b01cd80b29a0488_Grapes-Red-N-500gm_1.webp",
	}
	prd5 := Product{
		ID : 5,
		Title : "Mango",
		Description: "This is mango.",
		Price: 80.00,
		ImgURL: "https://belescooverseas.com/wp-content/uploads/2024/04/Mango.jpeg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
}