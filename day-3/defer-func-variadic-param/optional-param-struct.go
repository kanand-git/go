package main

type Config struct {
	Host       string
	Port       int
	SSl        bool
	MaxConn    int
	MaxRetries int
}

func main() {

	// when creating a struct, we don't have to assign values to every field
	c := Config{Host: "localhost", Port: 8080, MaxRetries: 3}
	openConn(c)
}
func openConn(c Config) {
	if c.Host == "" {
		// checking if host is empty then we can initialize with a default value
		c.Host = "localhost"
	}
}
