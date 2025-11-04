package main

import (
    "log"
    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"os"
)

func main() {
    // Создаем логгер
    logger := log.New(os.Stderr, "app: ", log.LstdFlags)

    // Создаем сервер с помощью функции из пакета server
    srv := server.NewServer(logger)

    // Запускаем сервер
    err := srv.Server.ListenAndServe()
    if err != nil {
        // Если при запуске сервера возникают ошибки, выводим их с помощью логгера на уровне Fatal
        logger.Fatal("Ошибка при запуске сервера: ", err)
    }
}
