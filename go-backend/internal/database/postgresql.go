package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var Db *sql.DB

func ConnectDatabase() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	//we read our .env file
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")
	psqlSetup := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user, pass, host, port, dbname)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", err)
		panic(err)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}

	createTablesQuery := `
	CREATE TABLE IF NOT EXISTS vetdonor_users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
	    pets VARCHAR(255)[]
	);

	CREATE TABLE IF NOT EXISTS vetdonor_clinic (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		address VARCHAR(255) NOT NULL,
		bloodtypesincluded VARCHAR(255)[],
	    bloodtypesnotincluded VARCHAR(255)[],
	    workhours VARCHAR(255) NOT NULL,
	    contacts VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS vetdonor_pets (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		breed VARCHAR(255) NOT NULL,
		petname VARCHAR(255) NOT NULL,
		pettype VARCHAR(255) NOT NULL,
		bloodtype VARCHAR(255) NOT NULL,
		age VARCHAR(255) NOT NULL,
		weight VARCHAR(255) NOT NULL,
		vaccinations VARCHAR(255) NOT NULL,
		photopath VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS vetdonor_need (
	  	id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
	    name VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS vetdonor_user_other_info (
	  	id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
	    city VARCHAR(255) NOT NULL,
	    phone VARCHAR(255) NOT NULL,
	    telegram VARCHAR(255) NOT NULL,
	    path VARCHAR(255) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS vetdonor_donation_stat (
	  	id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
    	blood INTEGER,
    	plasma INTEGER,
    	platelets INTEGER
	);

	CREATE TABLE IF NOT EXISTS vetdonor_pet_donors (
	  	id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		petname VARCHAR(255) NOT NULL,
    	blood INTEGER,
    	breed VARCHAR(255) NOT NULL,
    	contacts VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS vetdonor_donation (
	    id SERIAL PRIMARY KEY,
	    email VARCHAR(255) NOT NULL,
	    photo_path VARCHAR(255) NOT NULL,
	    donationType VARCHAR(255) NOT NULL
	)
`
	_, err = Db.Exec(createTablesQuery)
	if err != nil {
		fmt.Println("An error occurred while creating the table:", err)
		panic(err)
	} else {
		fmt.Println("Tables have been created successfully or already exist")
	}

}
