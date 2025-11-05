package handlers

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    data, err := os.ReadFile("index.html") // TO DO перевести на относительный путь
    if err != nil {
        http.Error(w, "Failed to load index.html", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")
    w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(0)
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        http.Error(w, fmt.Sprintf("Ошибка парсинга: %v", err), http.StatusInternalServerError)
        return
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError)
        return
    }

    convertedString, err := service.Convert(string(data))
    if err != nil {
        http.Error(w, fmt.Sprintf("Ошибка конвертации данных: %v", err), http.StatusInternalServerError)
        return
    }

    fileName := time.Now().Format("2006-01-02 15-04-05") + filepath.Ext(handler.Filename)
    err = os.WriteFile(fileName, []byte(convertedString), 0644)
    if err != nil {
        http.Error(w, fmt.Sprintf("Ошибка записи файла: %v", err), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "%s", convertedString)
}

func main() {
    http.HandleFunc("/", IndexHandler)
    http.ListenAndServe(":8080", nil)
}
