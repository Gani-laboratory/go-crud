package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api
	"strconv"  // package yang digunakan untuk mengubah string menjadi tipe int

	"github.com/Gani-laboratory/go-crud/entities"
	"github.com/Gani-laboratory/go-crud/models" // models package dimana tabel todo didefinisikan
	"github.com/gorilla/mux"                    // digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq"                       // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []entities.Todo `json:"data"`
}

// Tambah Todo
func TmbhTodo(w http.ResponseWriter, r *http.Request) {

	// create an empty user of type models.User
	// kita buat empty todo dengan tipe entities.Todo
	var todo entities.Todo

	// decode data json request ke todo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	// panggil modelnya lalu insert todo
	insertID := models.TambahTodo(todo)

	// format response objectnya
	res := response{
		ID:      int64(insertID),
		Message: "Todo baru telah ditambahkan",
	}

	// kirim response
	json.NewEncoder(w).Encode(res)
}

// AmbilTodo mengambil single data dengan parameter id
func AmbilTodo(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// dapatkan id todo dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	// konversi id dari string ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// memanggil models AmbilSatuTodo dengan parameter id yg nantinya akan mengambil single data
	todo, err := models.AmbilSatuTodo(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil todo. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(todo)
}

// Ambil semua list todo
func AmbilSemuaTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// memanggil models AmbilSemuaTodo
	todoList := models.AmbilSemuaTodo()

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = todoList

	// kirim semua response
	json.NewEncoder(w).Encode(response)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// buat variable todo dengan type entities.Todo
	var todo entities.Todo

	// decode json request ke variable todo
	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	// panggil UpdateTodo untuk mengupdate data
	updatedRows, err := models.UpdateTodo(int64(id), todo)

	if err != nil {
		log.Fatalf("Gagal mengupdate data. %v",err)
	}

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Todo telah berhasil diupdate. Jumlah yang diupdate %v rows/record", updatedRows)

	// ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}

func HapusTodo(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// panggil fungsi HapusTodo , dan convert int ke int64
	deletedRows, err := models.HapusTodo(int64(id))

	if err != nil {
		log.Fatalf("Gagal menghapus data. %v",err)
	}

	// ini adalah format message berupa string
	msg := fmt.Sprintf("todo berhasil di hapus. Total data yang dihapus %v", deletedRows)

	// ini adalah format reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}