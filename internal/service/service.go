package service
import (
	"strings"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
    "fmt"
)

func Convert(secret string) (otvet string, err error) {
    prefix1 := "."
    prefix2 := "-"
    if strings.HasPrefix(secret, prefix1) || strings.HasPrefix(secret, prefix2) {
        otvet = morse.ToText(secret)
        if otvet == "" { // Простая проверка на пустую строку как признак ошибки
            err = fmt.Errorf("не удалось преобразовать текст из азбуки Морзе")
        }
    } else {
        otvet = morse.ToMorse(secret)
        if otvet == "" { // Аналогичная проверка для преобразования в азбуку Морзе
            err = fmt.Errorf("не удалось преобразовать текст в азбуку Морзе")
        }
    }
    return otvet, err
}