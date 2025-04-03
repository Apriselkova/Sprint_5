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

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return errors.New("Invalid slice length")
	}

	// Парсинг количества шагов
	steps, err := strconv.Atoi(strings.TrimSpace(slice[0]))
	if err != nil || steps < 0 {
		return errors.New("Invalid number of steps")
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

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("Duration must be greater than 0")
	}

	// Вычисляем дистанцию
	distance := float64(ds.Steps) * StepLength / 1000 // в километрах
	// Вычисляем количество сожжённых калорий
	calories := spentenergy.WalkingSpentCalories(ds.Steps, float64(ds.Weight), float64(ds.Height), ds.Duration)
	if calories == 0 {
		return "", errors.New("Calories must be greater than 0")
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps, distance, calories)

	return result, nil
}
