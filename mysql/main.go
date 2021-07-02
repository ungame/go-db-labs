package main

import (
	"context"
	"flag"
	"fmt"
	"go-db-labs/mysql/db"
	"go-db-labs/mysql/models"
	"go-db-labs/mysql/store"
	"log"
	"sync"
)

var mysqlConfig db.MysqlConfig

func init() {
	mysqlConfig = db.NewConfigFromFlags(flag.CommandLine)
	flag.Parse()
}

func main() {
	ctx := context.Background()

	mysql := db.New(mysqlConfig)
	defer db.Close(mysql)

	err := mysql.Ping()
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Mysql connected!")

	conn, err := mysql.Conn(ctx)
	if err != nil {
		log.Panicln(err)
	}

	userStore := store.NewUsersStore(conn)

	GenerateUsers(ctx, userStore)
}

func GenerateUsers(ctx context.Context, usersStore store.UsersStore) {
	total := 100000

	wg := &sync.WaitGroup{}

	div := 10000
	workers := 0

	for i := 0; i < total; {

		wg.Add(1)
		go func(c context.Context, worker int) {
			for i := 0; i < div; i++ {
				err := usersStore.Create(c, models.NewDummyUser())
				if err != nil {
					fmt.Println("insert error: ", err)
				} else {
					//fmt.Printf("#%002d Worker - Insert %d/%d success!\n", worker, i, div)
				}
			}
			fmt.Printf("%d Worker done!\n", worker)
			wg.Done()
		}(ctx, workers + 1)

		workers++

		fmt.Printf("#%d Worker started\n", workers)

		i += div
	}
	wg.Wait()
}
