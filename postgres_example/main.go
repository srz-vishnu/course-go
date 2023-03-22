package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("hello")

	connectionStr := "postgres://postgres:test123@localhost:5432/mydb?sslmode=disable"

	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		fmt.Println("error at connection", err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("error at ping", err)
	}

	var tabChoose int
	fmt.Println("Choose Your Table")
	fmt.Println("1 TO CHOOSE TABLE COMPANY")
	fmt.Println("2 TO CHOOSE TABLE EMPLOYEE")

	fmt.Scanln(&tabChoose)

	switch tabChoose {

	case 1:
		chooseCompany(db)

	case 2:
		chooseEmployee(db)
	}

}

///////////////////////////////////////////////////////////////

func chooseCompany(db *sql.DB) {

	var option int
	for {

		fmt.Println("sELECT your Choice")
		fmt.Println("Choose 1 to insert")
		fmt.Println("Choose 2 to read")
		fmt.Println("Choose 3 to update")
		fmt.Println("choose 4 to delete")

		fmt.Scanln(&option)

		switch option {
		case 1:
			insertCompany(db)
			//fmt.Println("testttttttttt")

		case 2:
			readCompany(db)
			//fmt.Println("testttttttttt 22222")

		case 3:
			upadteCompany(db)

		case 4:
			deleteCompany(db)
		}

	}

}

func insertCompany(db *sql.DB) {
	var cmpy_id int
	var comp_name string
	var ceo string
	var phone_num int

	fmt.Println("Enter id")
	fmt.Scan(&cmpy_id)
	fmt.Println("Enter company name")
	fmt.Scan(&comp_name)
	fmt.Println("Enter ceo name")
	fmt.Scan(&ceo)
	fmt.Println("Enter phone number")
	fmt.Scan(&phone_num)

	query := "INSERT INTO COMPANY VALUES ($1, $2, $3, $4)"
	_, err := db.Query(query, cmpy_id, comp_name, ceo, phone_num)

	//query := fmt.Sprintf("INSERT INTO COMPANY VALUES(%d,%s,%s,%d)", cmpy_id, comp_name, ceo, phone_num)
	//_, err := db.Query(query)
	if err != nil {
		fmt.Println("db query err:", err)
	}

}

func readCompany(db *sql.DB) {
	query := "select * from COMPANY"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("db query err:", err)
	}

	defer rows.Close()

	for rows.Next() { //print all vendi
		var cmpy_id int
		var comp_name string
		var ceo string
		var phone_num int

		if err := rows.Scan(&cmpy_id, &comp_name, &ceo, &phone_num); err != nil {

			fmt.Println("rows error:", err)

		}

		fmt.Printf("company id = %d, company_name = %s, ceo = %s, phone number = %d\n", cmpy_id, comp_name, ceo, phone_num)
	}
}

func upadteCompany(db *sql.DB) {

	query := "UPDATE COMPANY SET phone_num = 9115000 WHERE cmpy_id = 8"
	_, err := db.Query(query)
	if err != nil {
		fmt.Println("db query err update:", err)
	}
}

func deleteCompany(db *sql.DB) {

	query := "DELETE FROM COMPANY WHERE cmpy_id = 11"

	_, err := db.Query(query)
	if err != nil {
		fmt.Println("id not exist:", err)
	}

}

//   2nd function

func chooseEmployee(db *sql.DB) {

	var option int
	for {

		fmt.Println("Select your choice")
		fmt.Println("Choose 1 to insert")
		fmt.Println("Choose 2 to read")
		fmt.Println("Choose 3 to update")
		fmt.Println("choose 4 to delete")

		fmt.Scanln(&option)

		switch option {
		case 1:
			insertEmployee(db)

		case 2:
			readEmployee(db)

		case 3:
			upadteEmployee(db)

		case 4:
			deleteEmployee(db)
		}

	}

}

func insertEmployee(db *sql.DB) {
	var firstname string
	var lastname string
	var department string
	var company int

	fmt.Println("Enter first name")
	fmt.Scan(&firstname)
	fmt.Println("Enter last name")
	fmt.Scan(&lastname)
	fmt.Println("Enter department name")
	fmt.Scan(&department)
	fmt.Println("Enter cmpy id")
	fmt.Scan(&company)

	query := "INSERT INTO EMPLOYEE VALUES ($1, $2, $3, $4)"
	_, err := db.Query(query, firstname, lastname, department, company)

	if err != nil {
		fmt.Println("db query err:", err)
	}

}

func readEmployee(db *sql.DB) {
	query := "select * from employee"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("err at read employee:", err)
	}

	defer rows.Close()

	for rows.Next() { //print all vendi
		var firstname string
		var lastname string
		var department string
		var company int

		if err := rows.Scan(&firstname, &lastname, &department, &company); err != nil {

			fmt.Println("rows error:", err)

		}

		fmt.Printf("firstname = %s, lastname = %s, department = %s, company = %d\n", firstname, lastname, department, company)
	}

}

func upadteEmployee(db *sql.DB) {

	query := "UPDATE EMPLOYEE SET firstname = 'vishnu' WHERE company = 1 "
	_, err := db.Query(query)
	if err != nil {
		fmt.Println("db query err update:", err)
	}
}

func deleteEmployee(db *sql.DB) {

	query := "DELETE FROM EMPLOYEE WHERE company = 1 AND department = 'Hr' "

	_, err := db.Query(query)
	if err != nil {
		fmt.Println("id not exist:", err)
	}

}
