package core

import (
	"log"
	"os"
)

var Log = log.New(os.Stdout, "", log.Ldate|log.Ltime)