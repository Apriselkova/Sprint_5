package actioninfo

import (
	"fmt"
)

// интерфейс DataParser, который реализует методы Parse() и ActionInfo()
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info() принимает слайс строк с данными о тренировках или прогулках и экземпляр одной из ваших структур Training или DaySteps
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Printf("ошибка парсинга данных '%s': %v\n", data, err)
			continue
		}

		// Получение информации об активности
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("ошибка получения информации для данных '%s': %v\n", data, err)
			continue
		}

		// Вывод информации
		fmt.Println(info)
	}
}
