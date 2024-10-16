package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

//Exec // when the query doesn't return anything
//QueryRow // when the query returns exactly one row
//Query // when the query returns multiple rows

type Conf struct {
	db *pgxpool.Pool
}

func Open() (*Conf, error) {
	const (
		host     = "localhost"
		port     = "5433"
		user     = "postgres"
		password = "postgres"
		dbname   = "postgres"
	)

	//sql.Open(psqlInfo)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	config, err := pgxpool.ParseConfig(psqlInfo)
	if err != nil {
		return nil, err
	}
	// MaxConns is the maximum number of connections that can be opened to PostgreSQL.
	// This limit can be used to prevent overwhelming the PostgreSQL server with too many concurrent connections.
	config.MaxConns = 30

	// MinConns is the minimum number of connections kept open by the pool.
	// The pool will not proactively create this many connections, but once this many have been established,
	// it will not close idle connections unless the total number exceeds MaxConns.
	config.MinConns = 10
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	//By setting the HealthCheckPeriod, your application proactively verifies that idle connections in the pool are still usable.
	//This can help detect problems earlier and ensure that new queries are given healthy connections.
	config.HealthCheckPeriod = time.Minute

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	return &Conf{db: db}, nil
}

func main() {
	conf, err := Open()
	if err != nil {
		log.Fatal(err)
	}
	err = conf.db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	//conf.createTable(context.Background())
	conf.insertUser(context.Background(), "Dev", "dev@email.co", 35)
	conf.queryUsers(context.Background())
	conf.updateUserEmail(context.Background(), 1, "john123@email.co")
}

// createTable creates the users table if it doesn't exist.
func (c *Conf) createTable(ctx context.Context) {
	// SQL query to create the users table
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE NOT NULL,
        age INT
    );`

	// Execute the query to create the table
	_, err := c.db.Exec(ctx, query)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err) // Log and terminate if table creation fails
	}
}

// insertUser inserts a new user into the users table and returns the new user's ID.
func (c *Conf) insertUser(ctx context.Context, name, email string, age int) int {
	// SQL query to insert a user and return the new user's ID
	// don't hardcode the values, or use the string in construction, sql injection can happen
	query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`
	var id int

	// Execute the query to insert the user and get the new user's ID
	//QueryRow returns one row as output
	err := c.db.QueryRow(ctx, query, name, email, age).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to insert user: %v\n", err) // Log and terminate if user insertion fails
	}
	return id // Return the new user's ID
}

// queryUsers retrieves and displays all users from the users table.
func (c *Conf) queryUsers(ctx context.Context) {
	// SQL query to retrieve all users
	query := `SELECT id, name, email, age FROM users`

	// Execute the query to retrieve all users
	rows, err := c.db.Query(ctx, query)
	if err != nil {
		log.Fatalf("Unable to query users: %v\n", err) // Log and terminate if query fails
	}
	defer rows.Close() // Ensure the rows are closed when done

	fmt.Println("Users:")

	// this loop would run until there are rows to scan
	for rows.Next() {
		var id, age int
		var name, email string

		// Scan each row into variables
		err := rows.Scan(&id, &name, &email, &age)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err) // Log and terminate if scanning fails
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", id, name, email, age) // Print user details
	}
}

// updateUserEmail updates a user's email based on their ID.
func (c *Conf) updateUserEmail(ctx context.Context, userID int, newEmail string) {
	// SQL query to update a user's email
	query := `UPDATE users SET email = $1 WHERE id = $2`

	// Execute the query to update the user's email
	_, err := c.db.Exec(context.Background(), query, newEmail, userID)
	if err != nil {
		log.Fatalf("Unable to update user: %v\n", err) // Log and terminate if update fails
	}
	fmt.Println("User email updated")
}

// deleteUser deletes a user from the users table based on their ID.
func (c *Conf) deleteUser(ctx context.Context, userID int) {
	// SQL query to delete a user by ID
	query := `DELETE FROM users WHERE id = $1`

	// Execute the query to delete the user
	_, err := c.db.Exec(context.Background(), query, userID)
	if err != nil {
		log.Fatalf("Unable to delete user: %v\n", err) // Log and terminate if deletion fails
	}
	fmt.Println("User deleted")
}
