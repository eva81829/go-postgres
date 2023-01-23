package controller

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
	"os"
)

const (
    host     = "myweb-db-1"
    port     = 5432
    user     = "psqladmin"
    password = "secret"
    dbname   = "go_database"
)

type User struct{
	Id   int     `json:"id"`
	Name string  `json:"name"`
}

func GetUsers(c *gin.Context) {
	db := ConnectDatabase()

	var users []User
	var uid int
	var uname string	
	sqlStmt := fmt.Sprintf("SELECT * FROM users;")
	rows, err := db.Query(sqlStmt)
	checkErr(err)
	
    for rows.Next() {
        err = rows.Scan(&uid, &uname)
		users = append(users, User{Id: uid, Name: uname})
    }

	c.JSON(http.StatusOK, users)

	db.Close()
}

func GetUserById(c *gin.Context) {	
	db := ConnectDatabase()
	id := c.Param("user_id")
	
	var user User
	var uid int
	var uname string
	sqlStmt := `SELECT userid, username FROM users WHERE userid=$1`
	err := db.QueryRow(sqlStmt, id).Scan(&uid, &uname)
	checkErr(err)

	user = User{Id: uid, Name: uname}

	c.JSON(http.StatusOK, user)

	db.Close()
}

func AddUser(c *gin.Context) {
	db := ConnectDatabase()

	var user User
	c.BindJSON(&user)

	sqlStmt := `INSERT INTO users(username) VALUES($1);`
	_, err := db.Exec(sqlStmt, user.Name)
    checkErr(err)

	c.JSON(http.StatusOK, gin.H{"message": "AddUser!"})

	db.Close()		
}

func ModUserById(c *gin.Context) {
	db := ConnectDatabase()
	id := c.Param("user_id")

	var user User
	c.BindJSON(&user)

	sqlStmt := `UPDATE users SET username=$1 WHERE userid=$2;`
	_, err := db.Exec(sqlStmt, user.Name, id)
    checkErr(err)
	
	c.JSON(http.StatusOK, gin.H{"message": "ModUserById"})

	db.Close()	
}

func DelUserById(c *gin.Context) {
	db := ConnectDatabase()
	id := c.Param("user_id")
	
	sqlStmt := `DELETE FROM users WHERE userid=$1`
	_, err := db.Exec(sqlStmt, id)
	checkErr(err)

	c.JSON(http.StatusOK, gin.H{"message": "DelUserById"})

	db.Close()	
}

func ConnectDatabase() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkErr(err)
	return db
}

func checkErr(err error) {
	fileName := "go_debug.log"
    logFile,err  := os.Create(fileName)
    defer logFile.Close()

	debugLog := log.New(logFile,"[Debug]",log.Llongfile)
	if err != nil {
		debugLog.Println("Connect Database Error!")
    } else {
		debugLog.Println("Connect Database Success!")
	}
}