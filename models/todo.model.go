package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/Gani-laboratory/go-crud/config"
	"github.com/Gani-laboratory/go-crud/entities"
	_ "github.com/lib/pq" // postgres golang driver
)

func TambahTodo(todo entities.Todo) uint {
	err := config.PostgresInstance.Create(&todo).Error

	if err != nil {
		log.Fatalf("Tidak Bisa menambahkan data todo. %v", err)
	}

	fmt.Printf("Insert data single record %v", todo.ID)

	// return insert id
	return todo.ID
}

// memberikan semua todo
func AmbilSemuaTodo() ([]entities.Todo) {
	var todo_list []entities.Todo
	config.PostgresInstance.Find(&todo_list)
	return todo_list
}

func cekJikaTodoDitemukan(todoId int64) bool {
	var product entities.Todo
	config.PostgresInstance.First(&product, todoId)
	return product.ID != 0
}

// mengambil satu todo
func AmbilSatuTodo(id int64) (entities.Todo, error) {
	var todo entities.Todo

	config.PostgresInstance.First(&todo, id)

	if !cekJikaTodoDitemukan(id) {
		return todo, errors.New("to-do tidak ditemukan")
	}
	
	return todo, nil
}

// update user dari database
func UpdateTodo(id int64, todo entities.Todo) (int64, error) {
	if !cekJikaTodoDitemukan(id) {
		return 0, errors.New("to-do tidak ditemukan")
	}

	config.PostgresInstance.First(&todo, id)
	// cek berapa banyak row/data yang diupdate
	rowsAffected := config.PostgresInstance.Save(&todo).RowsAffected

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected, nil
}

func HapusTodo(id int64) (int64, error) {
	if !cekJikaTodoDitemukan(id) {
		return 0, errors.New("to-do tidak ditemukan")
	}

	var todo entities.Todo
	rowsAffected := config.PostgresInstance.Delete(&todo, id).RowsAffected

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected, nil
}