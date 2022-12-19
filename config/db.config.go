package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Gani-laboratory/go-crud/entities"
	"github.com/joho/godotenv" // package yang dipakai untuk membaca file .env
	_ "github.com/lib/pq"      // postgres golang driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresInstance *gorm.DB

// kita buat koneksi dgn db posgres
func CreateConnection() {
	// load .env file
	err := godotenv.Load(".env")
	

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Membuat koneksi ke database
	PostgresInstance, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("SSL_MODE"))), &gorm.Config{})
	// check the connection
  if err != nil {
    panic("Gagal terhubung ke database")
  }

	fmt.Println("Sukses Konek ke Db!")
}

func Migrate(){	
	PostgresInstance.AutoMigrate(&entities.Todo{})
	log.Println("Database Migration Completed...")
}

