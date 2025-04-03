package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Printf("Ошибка парсинга данных '%s': %v\n", data, err)
			continue // Переход к следующей итерации в случае ошибки
		}

		// Получение информации об активности
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка получения информации для данных '%s': %v\n", data, err)
			continue // Переход к следующей итерации в случае ошибки
		}

		// Вывод информации
		fmt.Println(info)
	}
}
