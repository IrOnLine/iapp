package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// User model
type User struct {
	ID    int
	Name  string
	Email string
}

// CreateUser creates a new user in the database
func CreateUser(user User) error {

	// connect to db
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/dbname")
	if err != nil {
		return err
	}
	defer db.Close()

	// insert user
	res, err := db.Exec("INSERT INTO users VALUES (?, ?, ?)", user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}

	// get ID of created user
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set ID on user model
	user.ID = int(id)

	return nil
}

// GetUsers gets all users from the database
func GetUsers() ([]User, error) {

	// connect to db
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/dbname")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// query db
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	// scan rows into user models
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUser gets a single user by ID from the database
func GetUser(id string) (User, error) {

	// connect to db
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/dbname")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	// query db
	var user User
	err = db.QueryRow("SELECT * FROM users WHERE id=?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Good model
type Good struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

// Counterparty model
type Counterparty struct {
	ID   int
	Name string
	Type string // supplier or vendor
}

// Order model
type Order struct {
	ID     int
	UserID int
	Items  []OrderItem
	Total  float64
}

// OrderItem model
type OrderItem struct {
	GoodID   int
	Quantity int
}

// Warehouse model
type Warehouse struct {
	ID       int
	Name     string
	Location string
}

// Invoice model
type Invoice struct {
	ID      int
	OrderID int
	Total   float64
}
