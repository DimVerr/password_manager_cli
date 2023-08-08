package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
	Password string
	Credential []Credential 
  }
  
  type Credential struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	UserID      uint `gorm:"foreignKey"`
	CredName     string
	Domain       string
	Login        string
	Password     string
  }
  
var UtilsCmd = &cobra.Command{
	Use:   "utils",
	Short: "All  helpful structs, variables, functions are set here",
	Long: `All  helpful structs, variables, functions are set here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Utils has no command for use")
	},
}

func ConnectToDB() *gorm.DB{
	var err error
	errEnv := godotenv.Load()
	if errEnv != nil {
	  panic("Failed to load .env file")
	}
	
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD :=  os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	if DB_HOST == "" || DB_USER == "" || DB_PASSWORD == "" || DB_NAME =="" || DB_PORT =="" {
		fmt.Println("Please add your db data to .env file")
		os.Exit(1)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Credential{})
	return db
}

