package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable
	"strconv"  // package used to covert string into int type
	"strings"

	"github.com/ivalrivall/golang_crud1/models" // models package where User schema is defined
	"github.com/samber/lo"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"   // used to get the params from the route
	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// create user and store to db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.Customer

	err := decoder.Decode(&input)
	//Could not decode json
	if err != nil {
		ErrorResponse(http.StatusUnprocessableEntity, "Invalid JSON", w)
		return
	}

	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		var message string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				message = fmt.Sprintf("%s is not valid email",
					err.Field())
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			case "alpha":
				message = fmt.Sprintf("%s value must be alphanumeric",
					err.Field())
			}
			break
		}
		ErrorResponse(400, message, w)
		return
	}

	insertID := insertUser(input)

	res := make(map[string]interface{})
	res["message"] = "User created successfully with id " + strconv.Itoa(int(insertID))
	SuccessRespond(res, w)
}

// create brand and store to db
func CreateBrand(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.Brand

	err := decoder.Decode(&input)
	//Could not decode json
	if err != nil {
		ErrorResponse(http.StatusUnprocessableEntity, "Invalid JSON", w)
		return
	}

	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		var message string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required",
					err.Field())
			case "alpha":
				message = fmt.Sprintf("%s value must be alphanumeric",
					err.Field())
			}
			break
		}
		ErrorResponse(400, message, w)
		return
	}

	insertID := insertBrand(input)

	res := make(map[string]interface{})
	res["message"] = "Brand created successfully with id " + strconv.Itoa(int(insertID))
	SuccessRespond(res, w)
}

// create product and store to db
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input models.Product

	err := decoder.Decode(&input)
	//Could not decode json
	if err != nil {
		fmt.Println(err)
		ErrorResponse(http.StatusUnprocessableEntity, "Invalid JSON", w)
		return
	}

	validate := validator.New()
	err = validate.Struct(&input)
	if err != nil {
		var message string
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("err => ", err)
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				message = fmt.Sprintf("%s is not valid email",
					err.Field())
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			case "numeric":
				message = fmt.Sprintf("%s value must be numeric",
					err.Field())
			case "alpha":
				message = fmt.Sprintf("%s value must be alphanumeric",
					err.Field())
			}
			break
		}
		ErrorResponse(400, message, w)
		return
	}

	brand, err := getBrand(input.BrandId)

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
		ErrorResponse(400, "Unable to get brand. %v", w)
		return
	}

	fmt.Println("brand", brand)

	insertID := insertProduct(input)

	res := make(map[string]interface{})
	res["message"] = "Product created successfully with id " + strconv.Itoa(int(insertID))
	SuccessRespond(res, w)
}

// GetProduct will return a single product by its id
func GetProduct(w http.ResponseWriter, r *http.Request) {
	// get the id from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
		ErrorResponse(500, "Unable to convert the string into int. %v", w)
		return
	}

	// call the getProduct function with product id to retrieve a single product
	product, err := getProduct(int64(id))

	if err != nil {
		log.Fatalf("Unable to get product. %v", err)
		ErrorResponse(400, "Unable to get product. %v", w)
		return
	}

	// send the response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProductByBrand will return products by brand id
func GetProductByBrand(w http.ResponseWriter, r *http.Request) {
	// get the id from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
		ErrorResponse(500, "Unable to convert the string into int. %v", w)
		return
	}

	// call the getProductByBrand function by brand_id
	products, err := getProductByBrand(int64(id))

	if err != nil {
		log.Fatalf("Unable to get products. %v", err)
		ErrorResponse(400, "Unable to get products. %v", w)
		return
	}

	// send the response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(products)
}

// GetAllUser will return all the users
func GetAllUser(w http.ResponseWriter, r *http.Request) {

	// get all the users in the db
	users, err := getAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
		ErrorResponse(400, "Unable to get user. %v", w)
		return
	}

	// send all the users as response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
}

// UpdateUser update user's detail in the postgres db
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
		ErrorResponse(500, "Unable to convert the string into int.  %v", w)
		return
	}

	// create an empty user of type models.User
	var user models.Customer

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
		ErrorResponse(400, "Unable to decode the request body.  %v", w)
		return
	}

	// call update user to update the user
	updatedRows := updateUser(int64(id), user)

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// DeleteUser delete user's detail in the postgres db
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
		ErrorResponse(500, "Unable to convert the string into int.  %v", w)
		return
	}

	// call the deleteUser, convert the int to int64
	deletedRows := deleteUser(int64(id))

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// CreateOrder create order
type Products struct {
	ProductId int64 `json:"productId"`
}
type Orders struct {
	Products   []Products `json:"products"`
	CustomerId int64      `json:"customerId"`
}

type Value []interface{} // defined in the sql package

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var o Orders

	errDecode := json.NewDecoder(r.Body).Decode(&o)
	if errDecode != nil {
		http.Error(w, errDecode.Error(), http.StatusBadRequest)
		return
	}

	transactionSql := `INSERT INTO transactions (customer_id, amount) VALUES ($1, $2) RETURNING id`

	var transactionId int64

	db := createConnection()

	defer db.Close()

	errTrans := db.QueryRow(transactionSql, o.CustomerId, 0).Scan(&transactionId)

	if errTrans != nil {
		log.Fatalf("Unable to execute the query. %v", errTrans)
	}

	querySaveOrder := `INSERT INTO orders (product_id, transaction_id) VALUES ($1, $2) RETURNING id`

	for j := 0; j < len(o.Products); j++ {
		let := o.Products[j]
		orderID := 0
		errOrder := db.QueryRow(querySaveOrder, let.ProductId, transactionId).Scan(&orderID)
		if errOrder != nil {
			log.Fatalf("Unable to execute the query save order. %v", errOrder)
		}
	}

	var acc []string

	for _, b := range o.Products {
		acc = append(acc, fmt.Sprint(b.ProductId))
	}
	queryGetPrice := "SELECT price FROM products where id IN (" + strings.Join(acc, ",") + ")"

	stmt, errPrice := db.Query(queryGetPrice)
	if errPrice != nil {
		log.Fatalf("Unable to execute the query get price. %v", errPrice)
	}

	var totalPrice = []int64{}
	for stmt.Next() {
		var (
			price int64
		)
		if err := stmt.Scan(&price); err != nil {
			log.Fatal(err)
		}
		totalPrice = append(totalPrice, price)
	}

	sum := lo.SumBy(totalPrice, func(item int64) int64 {
		return item
	})

	queryUpdatePrice := "UPDATE transactions SET amount = " + fmt.Sprint(sum) + " WHERE id = " + fmt.Sprint(transactionId)

	resultUpdatePrice, errQueryUpdatePrice := db.Exec(queryUpdatePrice)

	if errQueryUpdatePrice != nil {
		log.Fatal(errQueryUpdatePrice)
	}

	row, errResultUpdatePrice := resultUpdatePrice.RowsAffected()
	if errResultUpdatePrice != nil {
		log.Fatal(errResultUpdatePrice)
	}

	fmt.Printf("query affected %d rows", row)

	res := make(map[string]interface{})
	res["message"] = "Order created successfully with id " + strconv.Itoa(int(transactionId))
	SuccessRespond(res, w)
}

/* =================== handler functions ==================== */
// insert one user in the DB
func insertUser(user models.Customer) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO customers (name) VALUES ($1) RETURNING id`

	// the inserted id will store in this id
	var lastInsertId int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, user.Name, user.CreatedAt).Scan(&lastInsertId)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", lastInsertId)

	// return the inserted id
	return lastInsertId
}

// insert one brand
func insertBrand(b models.Brand) int64 {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO brands (name) VALUES ($1) RETURNING id`

	// the inserted id will store in this id
	var lastInsertId int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, b.Name).Scan(&lastInsertId)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", lastInsertId)

	// return the inserted id
	return lastInsertId
}

// insert one product
func insertProduct(p models.Product) int64 {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO products (name, price, brand_id) VALUES ($1, $2, $3) RETURNING id`

	// the inserted id will store in this id
	var lastInsertId int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, p.Name, p.Price, p.BrandId).Scan(&lastInsertId)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", lastInsertId)

	// return the inserted id
	return lastInsertId
}

// get one brand from the DB by its id
func getBrand(id int64) (models.Brand, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a brand of models.Brand type
	var brand models.Brand

	// create the select sql query
	sqlStatement := `SELECT * FROM brands WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to brand
	err := row.Scan(&brand.ID, &brand.Name, &brand.CreatedAt)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return brand, nil
	case nil:
		return brand, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty brand on error
	return brand, err
}

// get one product from the DB by its id
func getProduct(id int64) (models.Product, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a product of models.Product type
	var product models.Product

	// create the select sql query
	sqlStatement := `SELECT * FROM products WHERE id=$1 LIMIT 1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to product
	err := row.Scan(&product.ID, &product.BrandId, &product.Name, &product.Price, &product.CreatedAt)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty product on error
	return product, err
}

func getProductByBrand(id int64) ([]models.Product, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a product of models.Product type
	var products []models.Product

	// create the select sql query
	sqlStatement := `SELECT * FROM products WHERE brand_id=$1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var product models.Product

		// unmarshal the row object to user
		err = rows.Scan(&product.ID, &product.BrandId, &product.Name, &product.Price, &product.CreatedAt)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		products = append(products, product)

	}

	// return empty product on error
	return products, err
}

// get one user from the DB by its userid
func getAllUsers() ([]models.Customer, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var users []models.Customer

	// create the select sql query
	sqlStatement := `SELECT * FROM customers`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user models.Customer

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, user)

	}

	// return empty user on error
	return users, err
}

// update user in the DB
func updateUser(id int64, user models.Customer) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE customers SET name=$2, location=$3, age=$4 WHERE userid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, user.Name)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete user in the DB
func deleteUser(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM customers WHERE userid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// return error response
func ErrorResponse(statusCode int, error string, writer http.ResponseWriter) {
	//Create a new map and fill it
	fields := make(map[string]interface{})
	fields["success"] = false
	fields["message"] = error
	message, err := json.Marshal(fields)

	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(message)
}

// return success response
func SuccessRespond(fields map[string]interface{}, writer http.ResponseWriter) {
	fields["success"] = true
	message, err := json.Marshal(fields)
	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(message)
}
