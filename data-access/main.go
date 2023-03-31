package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

type Friend struct {
    ID     		int64
    personName  string
    age 		int64
    career  	string
}

func main() {
    // Capture connection properties.
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

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	friendsAge, err := friendsByAge(20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Friends found: %v\n", friendsAge)

	friendsCareer, err := friendsByCareer("Plants")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Friends found: %v\n", friendsCareer)

    friendById, err := friendById(4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Friend found: %v\n", friendById)

    friend := Friend{
        personName: "Lucas",
        age: 22,
        career: "Software Engineering",
    }

    friendId, err := addFriend(friend)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Friend add with id: %v\n", friendId)

}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album

    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

func friendsByAge(age int64) ([]Friend, error) {

	var friends []Friend

    rows, err := db.Query("SELECT * FROM friends WHERE age >= ?", age)
    
	if err != nil {
        return nil, fmt.Errorf("friendsByAge %q: %v", age, err)
    }

    defer rows.Close()

	for rows.Next() {
        var friend Friend
        if err := rows.Scan(&friend.ID, &friend.personName, &friend.age, &friend.career); err != nil {
            return nil, fmt.Errorf("friendsByAge %q: %v", age, err)
        }
        friends = append(friends, friend)
    }
	
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("friendsByAge %q: %v", age, err)
    }

    return friends, nil
}

func friendsByCareer(career string) ([]Friend, error) {

	var friends []Friend

    rows, err := db.Query("SELECT * FROM friends WHERE career = ?", career)
    
	if err != nil {
        return nil, fmt.Errorf("friendsByCareer %q: %v", career, err)
    }

    defer rows.Close()

	for rows.Next() {
        var friend Friend
        if err := rows.Scan(&friend.ID, &friend.personName, &friend.age, &friend.career); err != nil {
            return nil, fmt.Errorf("friendsByCareer %q: %v", career, err)
        }
        friends = append(friends, friend)
    }
	
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("friendsByCareer %q: %v", career, err)
    }

    return friends, nil
}

func friendById(id int64) (Friend, error) {

	var friend Friend

    row := db.QueryRow("SELECT * FROM friends WHERE id = ?", id)
    if err := row.Scan(&friend.ID, &friend.personName, &friend.age, &friend.career); err != nil {
        if err == sql.ErrNoRows {
            return friend, fmt.Errorf("friendById %d: no such friend", id)
        }
        return friend, fmt.Errorf("friendById %d: %v", id, err)
    }
    return friend, nil
}

func addFriend(friend Friend) (int64, error) {

    result, err := db.Exec("INSERT INTO friends (personName, age, career) VALUES (?, ?, ?)", friend.personName, friend.age, friend.career)
    
    if err != nil {
        return 0, fmt.Errorf("addFriend %v",err)
    }
    
    id, err := result.LastInsertId()
    
    if err != nil {
        return 0, fmt.Errorf("addFriend %v",err)
    }

    return id, nil
}