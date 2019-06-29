package customer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"github.com/Khamontip/finalexam/database"
	
)
type Customers struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
	
func PostCustomersHandler(c *gin.Context){
	t := Customers{}
	fmt.Println(" 1) post")
	if err := c.ShouldBindJSON(&t); err != nil {
		log.Fatal("Open Error")
	}
	db,err := database.ConnDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
		
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal ("Open error")
	// }
	defer db.Close()
	query := `INSERT INTO customers (name,email,status) VALUES ($1,$2,$3) RETURNING id;`
	var id int
	row := db.QueryRow(query ,t.Name,t.Email,t.Status)
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("Scan id fail: " , id)
		return
	}
	t.ID  = id
	c.JSON(201,t)
 }
 
 func GetListCustomersHandler(c *gin.Context) {
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal("Open error: ", err.Error)
		fmt.Println("1 get list error ")
	}

	defer db.Close()
	fmt.Println("2.1 get list error ")
	//stmt ; err := db.Prepare("SELECT id, title ,status FROM todos")
	stmt, err := db.Prepare("SELECT id, name , email ,status FROM customers")
	if err != nil {
		log.Fatal("Prepare SQL error", err.Error)
	}

	rows, err := stmt.Query()
	fmt.Println("3 get list error ")
	if err != nil {
		log.Fatal("SQL error", err.Error)
	}

	customers := []Customers{}
	fmt.Println("4 get list error ")
	for rows.Next() {
		t := Customers{}
		fmt.Println("5 get list error ")
		err := rows.Scan(&t.ID, &t.Name, &t.Email  ,&t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, t)
	}
	c.JSON(200, customers)
}

func GetCustomersByIdHandler(c *gin.Context) {
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal(err.Error)
		fmt.Println("2 get error ")
	}
	defer db.Close()
	//stmt, _ := db.Prepare("SELECT id, title ,status from todos where id = $1")
	query := `SELECT id,name ,email ,status FROM customers WHERE id = $1;`
	if err != nil {
		log.Fatal("Sql error")
	}
	id := c.Param("id")
	fmt.Println("get --Pid-", id)
	row := db.QueryRow(query, id)
	t := Customers{}
	fmt.Println("List b4 scan ", (t))
	if err := row.Scan(&t.ID, &t.Name,&t.Email, &t.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errr": err.Error()})
		return
	}
	fmt.Println("After  scan ", (t))
	c.JSON(200, t)
}
func PutUpdateCustomersHandler(c *gin.Context) {
	fmt.Println("1 put error ")
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal("Open Error ", err.Error)
		fmt.Println("4 put error ")
	}
	defer db.Close()
	//query := `UPDATE todos SET title = $2, status = S3 WHERE id = $1 `
	//row := db.QueryRow(query, Pid, t.Title, t.Status)
	fmt.Println("2. put ")
	stmt, err := db.Prepare("UPDATE customers SET name=$2, email=$3 ,status=$4 WHERE id=$1;")
	if err != nil {
		log.Fatal("SQL eror", err.Error)
	}
	fmt.Println("1. put ")
	Pid := c.Param("id")
	fmt.Println("2. put Pid", Pid)
	t := Customers{}
	fmt.Println("3. put t", t)
	if err := (c.ShouldBindJSON(&t)); err != nil {
		c.JSON(http.StatusBadRequest, err.Error)
		return
	}

	t.ID, err = strconv.Atoi(Pid)
	fmt.Println("4. put t.ID ", t.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if _, err := stmt.Exec(Pid, t.Name, t.Email ,t.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Exec error": err.Error()})
		return
	}
	fmt.Println("Update successed ", t.ID, t.Name ,t.Email, t.Status)

	c.JSON(http.StatusOK, t)

}

func  DeleteCustomersHandler(c *gin.Context) {
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal("Can not open database", err.Error)
	}
	defer db.Close()
	//todos := []Todo{}
	query := `DELETE FROM customers  WHERE id=$1`
	var id int
	db.QueryRow(query, id)
	fmt.Println("Record Deleted ")
	c.JSON(200, gin.H{"message": "customer deleted"})

	
}