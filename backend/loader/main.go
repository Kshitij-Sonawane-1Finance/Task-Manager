package main

import (
    "fmt"
    "io"
    "os"

    "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/kshitij/taskManager/models"
)

func main() {
    stmts, err := gormschema.New("postgres").Load(&models.User{}, &models.Admin{}, &models.Task{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
        os.Exit(1)
    }
    io.WriteString(os.Stdout, stmts)
}