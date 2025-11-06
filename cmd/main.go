package main

import (
    "log"
    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"os"
)

func main() {
    logger := log.New(os.Stderr, "app: ", log.LstdFlags)
    srv := server.NewServer(logger)
    err := srv.Server.ListenAndServe()
    if err != nil {
        logger.Fatal("Ошибка при запуске сервера: ", err)
    }
}
