package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Friend struct {
	ID         int64  `json:"id"`
	PersonName string `json:"personName"`
	Age        int64  `json:"age"`
	Career     string `json:"career"`
}

func getFriends(c *gin.Context) {
	var friends []Friend

	rows, err := db.Query("SELECT * FROM friends")

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "Error")
	}

	defer rows.Close()

	for rows.Next() {
		var friend Friend
		if err := rows.Scan(&friend.ID, &friend.PersonName, &friend.Age, &friend.Career); err != nil {
			c.IndentedJSON(http.StatusUnauthorized, "Error")
		}
		friends = append(friends, friend)
	}

	c.IndentedJSON(http.StatusOK, friends)

	// log.Fatal(err)
	// hacer lista de amigos
	// c.IndentedJSON(http.StatusOK, friends)
}

func main() {

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()
	router.GET("/friends", getFriends)
	// router.GET("/friends/:id", getFriendById)
	// router.POST("/friends", postFriends)

	router.Run("localhost:8080")
}
