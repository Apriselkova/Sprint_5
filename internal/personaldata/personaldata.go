package personaldata

import (
	"fmt"
)

// структура Personal
type Personal struct {
	Name   string  // имя пользователя
	Weight float64 // вес пользователя
	Height float64 // рост пользователя
}

// метод Print() выводит данные структуры на экран
func (p Personal) Print() {
	fmt.Printf(`Имя: %s
Вес: %f
Рост: %f
`, p.Name, p.Weight, p.Height)
}
