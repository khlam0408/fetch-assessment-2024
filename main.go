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
	newUUID := uuid.New()
	var currReceipt models.Receipt

	err := c.BindJSON(&currReceipt)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The receipt is invalid"})
		return
	}

	points := internal.CalculatePoints(0, currReceipt)

	idsToPoints[newUUID.String()] = points
	
	c.IndentedJSON(http.StatusCreated, gin.H{"id": newUUID.String()})
}


func getPoints(c *gin.Context){
	id := c.Param("id")
	for i, p := range idsToPoints{
        if i == id {
            c.IndentedJSON(http.StatusOK, gin.H{"points": p})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No receipt found for that ID"})
	
}


func SetUpRouter() *gin.Engine{
	router := gin.Default()
	router.POST("/receipts/process", postID)
	router.GET("/receipts/:id/points", getPoints)
	return router
}
func main(){
	r := SetUpRouter()
	r.Run(":8080")
}