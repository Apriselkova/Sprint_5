package personaldata

import (
	"fmt"
)

// Ниже создайте структуру Personal
type Personal struct {
	Name   string  // имя пользователя
	Weight float64 // вес пользователя
	Height float64 // рост пользователя
}

// Ниже создайте метод Print()
func (p Personal) Print() {
	fmt.Printf(`Имя: %s
Вес: %f
Рост: %f
`, p.Name, p.Weight, p.Height)
}
