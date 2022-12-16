package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Gani-laboratory/go-crud/config"
	_ "github.com/lib/pq" // postgres golang driver
)

// schema dari tabel Todo
// kita coba dengan jika datanya null
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`
type Todo struct {
	ID            int64  `json:"id"`
	Title    string `json:"title"`
	Penulis       string `json:"penulis"`
	Tgl_publikasi string `json:"created_at"`
}

func TambahTodo(todo Todo) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat insert query
	// mengembalikan nilai id akan mengembalikan id dari todo yang dimasukkan ke db
	sqlStatement := `INSERT INTO todo (title, penulis, created_at) VALUES ($1, $2, $3) RETURNING id`

	// id yang dimasukkan akan disimpan di id ini
	var id int64

	// Scan function akan menyimpan insert id didalam id id
	err := db.QueryRow(sqlStatement, todo.Title, todo.Penulis, todo.Tgl_publikasi).Scan(&id)

	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("Insert data single record %v", id)

	// return insert id
	return id
}

// ambil satu todo
func AmbilSemuaTodo() ([]Todo, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var todo_list []Todo

	// kita buat select query
	sqlStatement := `SELECT * FROM todo`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var todo Todo

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Penulis, &todo.Tgl_publikasi)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		// masukkan kedalam slice todo_list
		todo_list = append(todo_list, todo)

	}

	// return empty todo atau jika error
	return todo_list, err
}

// mengambil satu todo
func AmbilSatuTodo(id int64) (Todo, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var todo Todo

	// buat sql query
	sqlStatement := `SELECT * FROM todo WHERE id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&todo.ID, &todo.Title, &todo.Penulis, &todo.Tgl_publikasi)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return todo, nil
	case nil:
		return todo, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return todo, err
}

// update user in the DB
func UpdateTodo(id int64, todo Todo) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat sql query create
	sqlStatement := `UPDATE todo SET title=$2, penulis=$3, created_at=$4 WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id, todo.Title, todo.Penulis, todo.Tgl_publikasi)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa banyak row/data yang diupdate
	rowsAffected, err := res.RowsAffected()

	//kita cek
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func HapusTodo(id int64) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// buat sql query
	sqlStatement := `DELETE FROM todo WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa jumlah data/row yang di hapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}