package handlers

import (
	"Food_shop/db"
	"Food_shop/types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DecodeRequest(r *http.Request, req interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("error decoding json")
	}
	return nil
}
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	response := types.Response{}
	w.Header().Set("Content-Type", "application/json")
	req := &types.CreateRequest{}

	err := DecodeRequest(r, req)
	if err != nil {
		response.Code = http.StatusInternalServerError
		return
	}

	err = db.CreateProduct(req.Name, req.Price)
	if err != nil {
		response.Code = http.StatusInternalServerError
		return
	}

	response.Code = http.StatusOK
	response.Message = "УСПЕШНО"

}

//type

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// get filters, parse json
	response := types.Response{}
	w.Header().Set("Content-Type", "application/json")

	res, err := database.Query("SELECT * FROM products where price between 1 AND 60")
	//(Where name=$1", name)
	if err != nil {
		fmt.Println(err)
	}
	products := []types.Products{}

	for res.Next() {
		p := types.Products{}
		err = res.Scan(&p.ID, &p.Name, &p.Price, &p.Created)
		if err != nil {
			response.Code = http.StatusInternalServerError
			continue
		}
		products = append(products, p)
	}
	//fmt.Println(products)

	response.Code = http.StatusOK
	response.Message = "УСПЕШНО"
	response.Payload = products
	json.NewEncoder(w).Encode(response)

}

func GetProductsbyID(w http.ResponseWriter, r *http.Request) {
	response := types.Response{}
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	row := database.QueryRow("SELECT * FROM products where id=$1", id)
	prod := types.Products{}
	err := row.Scan(&prod.ID, &prod.Name, &prod.Price, &prod.Created)
	if err != nil {
		response.Code = http.StatusInternalServerError

	}
	response.Code = http.StatusOK
	response.Message = "Успешно"
	response.Payload = prod

	json.NewEncoder(w).Encode(response)
}

type UpdateRequest struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func UpdatePruducts(w http.ResponseWriter, r *http.Request) {
	response := types.Response{}
	w.Header().Set("Content-Type", "application/json")
	req := &UpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec("UPDATE products set name = $1, price = $2 where id = $3", req.Name, req.Price, req.ID)
	if err != nil {
		response.Code = http.StatusInternalServerError
	}
	response.Code = http.StatusOK
	response.Message = "Успешно"
	response.Payload = req

	json.NewEncoder(w).Encode(response)

}

type DelateRequest struct {
	ID string `json:"id"`
}

func DeletePruducts(w http.ResponseWriter, r *http.Request) {
	response := types.Response{}
	w.Header().Set("Content-Type", "application/json")
	req := &DelateRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec("Delete FROM products where id = $1", req.ID)
	if err != nil {
		response.Code = http.StatusInternalServerError
	}
	response.Code = http.StatusOK
	response.Message = "Успешно"

}
