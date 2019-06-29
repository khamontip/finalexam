package main
import (
	"fmt"
	"net/http"
    "github.com/gin-gonic/gin"
   _"github.com/lib/pq"
   "github.com/Khamontip/finalexam/database"
   "github.com/Khamontip/finalexam/customer"
    
)


func autMiddleware(c * gin.Context){
	//r.Use(func(c *gin.Context) {
		fmt.Println("Hello")
		token := c.GetHeader("Authorization")
		fmt.Println("token:", token)
		if token != "token2019"{
			//c.JSON(http.StatusUnauthorized , gin.H{"error": http.StatusText(http.StatusUnauthuthorized)})
			c.JSON(http.StatusUnauthorized , gin.H{"error": "Use"})
		c.Abort()
		return
		}
		c.Next()
		fmt.Println("Goodbye!")
}
func main() {

database.CreatDB()
r := gin.Default()

r.Use(autMiddleware)
r.POST("/customers", customer.PostCustomersHandler)
r.GET ("/customers/:id", customer.GetCustomersByIdHandler)
r.GET ("/customers", customer.GetListCustomersHandler)
r.PUT ("/customers/:id", customer.PutUpdateCustomersHandler)
r.DELETE ("/customers/:id", customer.DeleteCustomersHandler)
r.Run(":2019")

}