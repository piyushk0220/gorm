package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Movie struct {
	ID    uint `gorm:"primary_key"`
	Title string
}

type Artist struct {
	ID     uint `gorm:"primary_key"`
	Name   string
	Movies []Movie `gorm:"many2many:artist_movies"`
}

var input int
var artists []Artist
var movies []Movie

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Movie{}, &Artist{})
	//db.Create(&Movie{ID: 22, Title: "KGF"})

	for {

		fmt.Println("Welcome to Artists management app")
		fmt.Println("Press 1. Create")
		fmt.Println("Press 2. Read")
		fmt.Println("Press 3. Update")
		fmt.Println("Press 4. Delete")
		fmt.Println("Press 5. Join")
		fmt.Println("Press 6. Exit")

		fmt.Scanln(&input)

		if input == 1 {
			create()
		} else if input == 2 {
			art := read()
			fmt.Println(art)
		} else if input == 3 {
			update()
		} else if input == 4 {
			delete()
		} else if input == 5 {
			join()
		} else {
			break
		}

	}

}

func create() {

	dsn := "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	movies := []Movie{
		{ID: 1, Title: "Nayagan"},
		{ID: 2, Title: "Anbe sivam"},
		{ID: 3, Title: "3 idiots"},
		{ID: 4, Title: "Shamithab"},
		{ID: 5, Title: "Dark Knight"},
		{ID: 6, Title: "310 to Yuma"},
	}
	for i := range movies {
		if err := db.Create(&movies[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	artists := []Artist{{Name: "Madhavan", Movies: []Movie{movies[1], movies[2]}},
		{Name: "Kamal Hassan", Movies: []Movie{movies[0], movies[1]}},
		{Name: "Dhanush", Movies: []Movie{movies[3]}},
		{Name: "Aamir Khan", Movies: []Movie{movies[2]}},
		{Name: "Amitabh Bachchan", Movies: []Movie{movies[3]}},
		{Name: "Christian Bale", Movies: []Movie{movies[4], movies[5]}},
		{Name: "Russell Crowe", Movies: []Movie{movies[5]}},
	}

	for i := range artists {
		if err := db.Create(&artists[i]).Error; err != nil {
			log.Fatal(err)
		}
	}
}

func read() []Artist {
	dsn := "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	db.Preload("Movies").Find(&artists)

	return artists

}

// updating multiple columns (id & name)
func update() {

	dsn := "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	qry := "SET FOREIGN_KEY_CHECKS=0"
	db.Exec(qry)
	db.Model(Artist{}).Where("ID = ?", 1).Updates(Artist{ID: 111, Name: "laksksk"})

}

func delete() {
	dsn := "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	qry := "DELETE FROM Artists WHERE Name='laksksk';"
	db.Exec(qry)
}
func join() {
	dsn := "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	// Get the list the artists for movie "3 idiots"

	artists = []Artist{}
	if err = db.Joins("JOIN artist_movies on artist_movies.artist_id=artists.id").
		Joins("JOIN movies on artist_movies.movie_id=movies.id").Where("movies.title=?", "3 idiots").
		Group("artists.id").Find(&artists).Error; err != nil {
		log.Fatal(err)
	}

	for _, ar := range artists {
		fmt.Println(ar.Name)
	}
}
