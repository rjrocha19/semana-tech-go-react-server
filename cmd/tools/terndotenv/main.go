package main

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        fmt.Printf("Error loading .env file: %v\n", err)
        return
    }
    fmt.Println("DATABASE_NAME:", os.Getenv("PGDATABASE"))
    fmt.Println("DATABASE_USER:", os.Getenv("PGUSER"))
    fmt.Println("DATABASE_PASSWD:", os.Getenv("PGPASSWORD"))
    fmt.Println("DATABASE_HOST:", os.Getenv("PGHOST"))

    cmd := exec.Command(
        "tern",
        "migrate",
        "--migrations",
        "./internal/store/pgstore/migrations",
        "--config",
        "./internal/store/pgstore/migrations/tern.conf",
    )

    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Command execution failed with error: %v\n", err)
        fmt.Printf("Output: %s\n", string(output))
        return
    }

    fmt.Printf("Command executed successfully: %s\n", string(output))
}