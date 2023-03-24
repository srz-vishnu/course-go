package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

//defining struct

type Laptops struct {
	Id        int    `json: "laptopid"`
	Brandname string `json: "brandname"`
	Price     int    `json: "price" `
}

// type JsonResponse struct {
// 	Type string    `json:"type"`
// 	Data []Laptops `json:"data"`
// }

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "test123"
	DB_NAME     = "mydb"
)

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println("error at connection", err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("error at ping", err)
	}

	return db
}

func main() {

	// connection to DB
	// connectionStr := "postgres://postgres:test123@localhost:5432/mydb?sslmode=disable"

	// db, err := sql.Open("postgres", connectionStr)

	// if err != nil {
	// 	fmt.Println("error at connection", err)
	// }

	// if err := db.Ping(); err != nil {
	// 	fmt.Println("error at ping", err)
	// }

	fmt.Println("hello")

	router := mux.NewRouter() //initialized mux router

	router.HandleFunc("/laptops", GetLaptops).Methods("GET")

	router.HandleFunc("/laptops/view", PostLaptops).Methods("POST")

	router.HandleFunc("/laptops/delete", DeleteLaptops).Methods("DELETE")

	router.HandleFunc("/laptops/update", UpdateLaptops).Methods("PUT")

	fmt.Println("Server at 8082")
	log.Fatal(http.ListenAndServe(":8082", router))

}

func GetLaptops(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	query := "select * from laptop"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("db query err:", err)
	}

	//defer rows.Close()

	var laptop []Laptops

	for rows.Next() { //print all vendi
		var id int
		var brandname string
		var price int

		if err := rows.Scan(&id, &brandname, &price); err != nil {

			fmt.Println("rows error:", err)

		}

		fmt.Printf("Laptop id = %d, Brand name = %s, price = %d\n", id, brandname, price)
		laptop = append(laptop, Laptops{Id: id, Brandname: brandname, Price: price})
	}

	// var response = JsonResponse{Type: "success", Data: laptop}
	// json.NewEncoder(w).Encode(response)

	json.NewEncoder(w).Encode(laptop)

	// data, err := json.Marshal(laptop)
	// if er != nil {
	// 	fmt.Println("Error:", err)
	// }
	// w.Write(data)

}

func PostLaptops(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	var id int
	var brandname string
	var price int

	var laptop []Laptops

	fmt.Println("Enter id")
	fmt.Scan(&id)
	fmt.Println("Enter brand name")
	fmt.Scan(&brandname)
	fmt.Println("Enter price")
	fmt.Scan(&price)
	fmt.Println("Done")

	query := "INSERT INTO laptop VALUES ($1, $2, $3)"
	_, err := db.Query(query, id, brandname, price)
	laptop = append(laptop, Laptops{Id: id, Brandname: brandname, Price: price})

	if err != nil {
		fmt.Println("db query err at post:", err)
	}

	json.NewEncoder(w).Encode(laptop)

}

func DeleteLaptops(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	var id int

	//var laptop []Laptops

	fmt.Println("Enter the id you want to delete")
	fmt.Scan(&id)

	query := "DELETE FROM laptop WHERE id = $1"

	_, err := db.Query(query, id)
	if err != nil {
		fmt.Println("id not exist:", err)
	}
	//json.NewEncoder(w).Encode(laptop)
	w.Write([]byte("deleted")) // to get notification on postman ,like println
}

func UpdateLaptops(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	query := "UPDATE laptop set PRICE = 200000 WHERE ID = 2"

	_, err := db.Query(query)
	if err != nil {
		fmt.Println("update error:", err)
	}

}
