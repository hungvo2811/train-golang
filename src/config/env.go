package config

//import (
//	"fmt"
//	"github.com/joho/godotenv"
//	"os"
//	_ "os"
//)
//
//var ProcessEnv Env
//
//func InitDotEnv() {
//	if err := godotenv.Load(); err != nil {
//		fmt.Println("Error loading env file")
//	}
//	portApp := GetEnv("PORT", "8080")
//	database := Database{Url: GetEnv("MONGODB_URL", ""), Name: GetEnv("Name_BD", "test")}
//	ProcessEnv = Env{PortApp: portApp, Database: database}
//
//}
//
//func GetEnv(key, defaultValue string) string {
//	value := os.Getenv(key)
//	if len(value) == 0 {
//		return defaultValue
//	}
//	return value
//}
