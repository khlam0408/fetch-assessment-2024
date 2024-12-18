package main
import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"fetch-assessment/internal"
	"fetch-assessment/models"
)

var idsToPoints = make(map[string]int)

func postID(c *gin.Context){
	newUUID := uuid.New() // the ID for the receipt
	var currReceipt models.Receipt

	err := c.BindJSON(&currReceipt)

	// If it cannot bind to the receipt, a error message will be sent
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The receipt is invalid"})
		return
	}

	points := internal.CalculatePoints(0, currReceipt)

	idsToPoints[newUUID.String()] = points
	
	c.IndentedJSON(http.StatusCreated, gin.H{"id": newUUID.String()}) // response with a JSON object
}


func getPoints(c *gin.Context){
	id := c.Param("id")
	// goes through a list of all ids that has been inputed via the POST route
	for i, p := range idsToPoints{
        if i == id {
            c.IndentedJSON(http.StatusOK, gin.H{"points": p})
            return
        }
    }
	// If there is no such ID (not being placed via the POST), will respond with an error message 
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No receipt found for that ID"})
	
}

// Router set up for all the necessary routes
func SetUpRouter() *gin.Engine{
	router := gin.Default()
	router.POST("/receipts/process", postID)
	router.GET("/receipts/:id/points", getPoints)
	return router
}

// Runs the router
func main(){
	r := SetUpRouter()
	r.Run(":8080")
}