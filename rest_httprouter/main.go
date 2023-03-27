package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

type Companys struct {
	Compny_id    int    `json:"companyid"`
	Company_name string `json:"companyname"`
	Ceo          string `json:"ceo"`
	Phonenum     int    `json:"phonenum" `
}

type JsonResponse struct {
	Type    string     `json:"type"`
	Data    []Companys `json:"data"`
	Message string     `json:"message"`
}

// DB connection
func setupDB() *sql.DB {

	connectionStr := "postgres://postgres:test123@localhost:5432/mydb?sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		fmt.Println("error at connection", err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("error at ping", err)
	}

	return db
}

func main() {

	fmt.Println("hello")

	router := httprouter.New() // initialized httprouter

	router.GET("/company", GetCompany)

	router.POST("/company/view", PostCompany)

	router.DELETE("/company/delete", DeleteCompany)

	router.PUT("/company/update", UpdateCompany)

	fmt.Println("Server at 8888")
	log.Fatal(http.ListenAndServe(":8888", router))

}

func GetCompany(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := setupDB()

	query := "select * from company"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("db query err:", err)
	}

	var company []Companys

	for rows.Next() {
		var companyid int
		var companyname string
		var ceo string
		var phonenum int

		if err := rows.Scan(&companyid, &companyname, &ceo, &phonenum); err != nil {

			fmt.Println("rows error:", err)

		}

		fmt.Printf("Company id = %d, Company name = %s,Ceo = %s, Phone num = %d\n", companyid, companyname, ceo, phonenum)
		company = append(company, Companys{Compny_id: companyid, Company_name: companyname, Ceo: ceo, Phonenum: phonenum})
	}

	var response = JsonResponse{Type: "success", Data: company, Message: "All data received"}
	json.NewEncoder(w).Encode(response)

}

func PostCompany(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := setupDB()

	var companyid int
	var companyname string
	var ceo string
	var phonenum int

	var company []Companys

	fmt.Println("Enter company id")
	fmt.Scan(&companyid)
	fmt.Println("Enter brand name")
	fmt.Scan(&companyname)
	fmt.Println("Enter CEO name")
	fmt.Scan(&ceo)
	fmt.Println("Enter phonenumber")
	fmt.Scan(&phonenum)
	fmt.Println("Done")

	query := "INSERT INTO company VALUES ($1, $2, $3,$4)"
	_, err := db.Query(query, companyid, companyname, ceo, phonenum)
	company = append(company, Companys{Compny_id: companyid, Company_name: companyname, Ceo: ceo, Phonenum: phonenum})

	if err != nil {
		fmt.Println("db query err at post:", err)
	}

	var response = JsonResponse{Type: "success", Data: company, Message: "One data Posted succesfully"}
	json.NewEncoder(w).Encode(response)

	//json.NewEncoder(w).Encode(company)

}

func DeleteCompany(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := setupDB()

	var companyid int

	//var laptop []Laptops

	fmt.Println("Enter the id you want to delete")
	fmt.Scan(&companyid)

	query := "DELETE FROM company WHERE cmpy_id = $1"

	_, err := db.Query(query, companyid)
	if err != nil {
		fmt.Println("id not exist:", err)
	}

	w.Write([]byte("deleted")) // to get notification on postman ,like println
}

func UpdateCompany(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := setupDB()

	query := "UPDATE company set ceo = 'vk' WHERE cmpy_id = 8"

	_, err := db.Query(query)
	if err != nil {
		fmt.Println("update error:", err)
	}

}
