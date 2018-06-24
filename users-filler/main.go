package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var (
	serverAddr = flag.String("servername", "localhost:26257", "address of cocroachdb node")
	userName   = flag.String("username", "username", "prefix of username")
	from       = flag.Int("from", 0, "initial value of N")
	to         = flag.Int("to", 0, "upper bound of N")
)

const queryTemplate = `INSERT INTO users(username, node) VALUES  %s ON CONFLICT(username) DO NOTHING;`

func main() {
	flag.Parse()

	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://root@%s/messcore?sslmode=disable", *serverAddr))
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	for i := *from / 100; i < *to/100; i++ {
		q := makeBatchQuery(i*100, "dn1.cn1.mellody.cloud")

		if _, err := db.Exec(q); err != nil {
			log.Fatal(err)
		}

		// fmt.Println(q)
		fmt.Println("insert pack: ", i+1, " ok")
	}
}

func makeBatchQuery(batchn int, node string) string {
	var buf [100]string
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(batchn)
		buf[i] = fmt.Sprintf("('%s', '%s')", *userName+"_"+s, node)
		batchn++
	}

	return fmt.Sprintf(queryTemplate, strings.Join(buf[:], ","))
}
