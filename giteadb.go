package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

//GiteaDB holds a database and the neeeded functions
type GiteaDB struct {
	path string
}

//NewGiteaDB returns a new GiteaDB
func NewGiteaDB(path string) *GiteaDB {
	db := &GiteaDB{path}
	db.Init()
	return db
}

//Init initializes the db, if it exists in the path it will be overwritten
func (dbg *GiteaDB) Init() {
	os.Remove(dbg.path)

	dbtmp, err := sql.Open("sqlite3", dbg.path)
	if err != nil {
		log.Fatal(err)
	}
	defer dbtmp.Close()

	sqlStmt := `
	create table tokens (room text not null primary key, token text);
	delete from tokens;
	`
	_, err = dbtmp.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}

func (dbg *GiteaDB) GetToken(room string) string {

	db, err := sql.Open("sqlite3", dbg.path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select token from tokens where room = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var token string
	err = stmt.QueryRow(room).Scan(&token)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

//GetAll returns all existing rooms with token
func (dbg *GiteaDB) GetAll() map[string]string {
	tokens := make(map[string]string)

	db, err := sql.Open("sqlite3", dbg.path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select room, token from tokens")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var room string
		var token string
		err = rows.Scan(&room, &token)
		if err != nil {
			log.Fatal(err)
		}
		tokens[room] = token
		fmt.Println(room, token)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return tokens

}

//Unset deletes a room and it's token from the database, if it exists
func (dbg *GiteaDB) Unset(room, token string) {

	db, err := sql.Open("sqlite3", dbg.path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from tokens where room = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(room)
	if err != nil {
		log.Fatal(err)
	}
}

//Set sets a token for a room and saves it to the db
func (dbg *GiteaDB) Set(room, token string) {

	dbg.Unset(room, token)

	db, err := sql.Open("sqlite3", dbg.path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserting")

	stmt, err := tx.Prepare("insert into tokens(room, token) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(room, token)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	dbg.GetAll()

}
