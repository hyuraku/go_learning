package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:@/test_mysql_database?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    // { // Create a new table
    //     query := `
    //         create table users (
    //             id INT AUTO_INCREMENT,
    //             username TEXT NOT NULL,
    //             password TEXT NOT NULL,
    //             created_at DATETIME,
    //             PRIMARY KEY (id)
    //         );`
    //
    //     if _, err := db.Exec(query); err != nil {
    //         log.Fatal(err)
    //     }
    // }

    { // Insert a new user
      username := "Jone"
      password := "password"
      createdAt := time.Now()

      result, err := db.Exec(`insert into users (username, password, created_at) values (?,?,?)`,username,password, createdAt)
      if err != nil{
        log.Fatal(err)
      }

      id, err := result.LastInsertId()
      fmt.Println(id)
    }

    { // Query all users
      type user struct {
        id  int
        username string
        password string
        createdAt time.Time
      }

      rows,err := db.Query(`select id, username, password, created_at from users`)

      if err != nil{
        log.Fatal(err)
      }
      defer rows.Close()
      var users []user
      for rows.Next(){
        var u user
        err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
        if err != nil {
          log.Fatal(err)
         }
        users = append(users, u)
      }

       fmt.Printf("%#v", users)
    }
}
