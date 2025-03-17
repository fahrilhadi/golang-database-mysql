package repository

import (
	"context"
	"fmt"
	"testing"

	golang_database_mysql "github.com/fahrilhadi/golang-database-mysql"
	"github.com/fahrilhadi/golang-database-mysql/entity"
	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T)  {
	commentRepository := NewCommentRepository(golang_database_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T)  {
	commentRepository := NewCommentRepository(golang_database_mysql.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T)  {
	commentRepository := NewCommentRepository(golang_database_mysql.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}