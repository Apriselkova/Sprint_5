package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// структура DaySteps содержит все необходимые данные о дневных прогулках: количество шагов, длительность, а также данные из структуры personaldata.Personal
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// сметод Parse() парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return errors.New("Invalid slice length")
	}

	// Парсинг количества шагов
	steps, err := strconv.Atoi(strings.TrimSpace(slice[0]))
	if err != nil || steps < 0 {
		return errors.New("invalid number of steps")
	}
	ds.Steps = steps

	// Парсинг длительности
	durationString := strings.TrimSpace(slice[1])
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return err
	}
	ds.Duration = duration

	return nil
}

// метод ActionInfo() формирует и возвращает строку с данными о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("duration must be greater than 0")
	}

	// Вычисляем дистанцию
	distance := spentenergy.Distance(ds.Steps)

	// Вычисляем количество сожжённых калорий
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if calories == 0 {
		return "", errors.New("calories must be greater than 0")
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps, distance, calories)

	return result, nil
}
