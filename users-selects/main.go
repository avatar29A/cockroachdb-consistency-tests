package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/lib/pq"
)

var (
	serverAddr = flag.String("servername", "localhost:26257", "address of cocroachdb node")
	times      = flag.Int("times", 100, "")
)

func main() {
	flag.Parse()

	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://root@%s/messcore?sslmode=disable", *serverAddr))
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	for i := 0; i < *times; i++ {
		n := rand.Intn(50000)
		if _, err := db.Query(fmt.Sprintf(`SELECT * FROM users WHERE username = 'username_%d' LIMIT 1`, n)); err != nil {
			if err == sql.ErrNoRows {
				fmt.Println(i, " pass (rows not found)")
				continue
			}

			log.Fatal(err)
		}

		fmt.Println(i, n, " pass")
	}
}
