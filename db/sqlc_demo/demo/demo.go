package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go_demo/db/sqlc_demo/db/sqlc"
	"log"
)

func main() {
	conn, err := sql.Open("postgres", "dbname=db01 host=127.0.0.1 port=5432 user=test password=welcome sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	queries := db.New(conn)

	//// select 多行记录
	//fmt.Println("------------------- select many rows -------------------")
	//authors, err := queries.ListAuthors(context.Background())
	//if err != nil {
	//	log.Fatal("listAuthors error:", err)
	//}
	//
	//for _, author := range authors {
	//	fmt.Println(author.Bio)
	//}
	//
	// select 单行记录
	fmt.Println("------------------- select one row -------------------")
	author, err := queries.GetAuthor(context.Background(), 7)
	if err != nil {
		log.Printf("query error, %s", err)
	}
	fmt.Println(author)

	//// select in array
	//fmt.Println("------------------- select many rows by IDS -------------------")
	//ids := []int32{1, 2, 3}
	//ds, err := queries.ListAuthorByIDs(context.Background(), ids)
	//if err != nil {
	//	log.Printf("query error, %s", err)
	//}
	//for _, d := range ds {
	//	fmt.Println(d.Name)
	//}
	//
	//// insert
	//fmt.Println("------------------- insert -------------------")
	//auth1 := db.CreateAuthorParams{
	//	Name: "东邪",
	//	Bio:  sql.NullString{"", true},
	//}
	//
	//auth2 := db.CreateGetAuthorParams{
	//	Name: "张山丰",
	//	Bio:  sql.NullString{"武当山", true},
	//}
	//
	//queries.CreateAuthor(context.Background(), auth1)
	//zsf, err := queries.CreateGetAuthor(context.Background(), auth2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(zsf)
	//
	//// delete
	//fmt.Println("------------------- delete -------------------")
	//queries.DeleteAuthor(context.Background(), 1)
	//getAuthor, err := queries.DeleteGetAuthor(context.Background(), 2)
	//if err != nil {
	//	log.Printf("delete error, %s", err)
	//}
	//fmt.Println(getAuthor)
	//
	//// update
	//fmt.Println("------------------- update -------------------")
	//updateAuthor := db.UpdateGetAuthorParams{
	//	ID:   3,
	//	Name: "洪七",
	//	Bio:  sql.NullString{"叫花子", true},
	//}
	//
	//updateGetAuthor, err := queries.UpdateGetAuthor(context.Background(), updateAuthor)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(updateGetAuthor)

	// prepar query
	fmt.Println("------------------- preparing queries --------------------------")
	prepare, err := db.Prepare(context.Background(), conn)
	if err != nil {
		log.Fatal("prepare error,", err)
	}

	auth7, err := prepare.GetAuthor(context.Background(), 7)
	if err != nil {
		log.Printf("preparing select error, %s", err)
	}
	fmt.Println(auth7)

	auth8, err := prepare.GetAuthor(context.Background(), 8)
	if err != nil {
		log.Printf("preparing select error, %s", err)
	}
	fmt.Println(auth8)
}
