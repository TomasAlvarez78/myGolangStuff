package main

import (
    "net/http"
    "fmt"
    "strconv"
    "github.com/gin-gonic/gin"
)

type Friend struct {
    ID     		int64	`json:"id"`
    PersonName  string	`json:"personName"`
    Age 		int64	`json:"age"`
    Career  	string	`json:"career"`
}

var friends = []Friend{
	{ID: 1, PersonName: "Tomas", Age:22, Career:"Software Engineering"},
	{ID: 2, PersonName: "Emma", Age:19, Career:"Environmental Management"},
	{ID: 3, PersonName: "Romanito", Age:22, Career:"Telecommunications Engineering"},
	{ID: 4, PersonName: "Facundito", Age:20, Career:"Software Engineering"},
	{ID: 5, PersonName: "Chaqueno", Age:19, Career:"Software Engineering"},
	{ID: 6, PersonName: "Benja", Age:20, Career:"Software Engineering"},
	{ID: 7, PersonName: "Lucas", Age:22, Career:"Software Engineering"},
}

func getFriends(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, friends)
}

func postFriends(c *gin.Context) {
    var newFriend Friend
    if err := c.BindJSON(&newFriend); err != nil {
        return
    }

    friends = append(friends, newFriend)
    c.IndentedJSON(http.StatusCreated, newFriend)
}

func getFriendById(c *gin.Context){
    
    id_param := c.Param("id")
    id,err := strconv.ParseInt(id_param, 10, 64)
    if err != nil {
        fmt.Println("hehe error")
    }

    fmt.Println("holiwi: id parameter: %v", id)

    for _, a := range friends {

        fmt.Println("actual friend in loop %v", a)

        if (id == a.ID){
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message":"friend not found"})
}

func main() {
    router := gin.Default()
    router.GET("/friends", getFriends)
    router.GET("/friends/:id", getFriendById)
    router.POST("/friends", postFriends)


    router.Run("localhost:8080")
}

