package logger

import (
	"log"
	"os"
)

// InitLogger initializes logging system
func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Logger initialized")
}
