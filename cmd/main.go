package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/nikhilsharma270027/API-Cart-GO/cmd/api"
	"github.com/nikhilsharma270027/API-Cart-GO/config"
	"github.com/nikhilsharma270027/API-Cart-GO/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}
	// server := api.NewAPIServer(":8080", nil) adding db
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal("Server Running on port ")
	}
}
