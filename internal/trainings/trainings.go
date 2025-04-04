package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return errors.New("Invalid slice length")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(slice[0]))
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("The number of steps must be greater than 0")
	}
	t.Steps = steps

	trainingType := strings.TrimSpace(slice[1])
	if trainingType != "Бег" && trainingType != "Ходьба" {
		return errors.New("Unknown training type")
	}
	t.TrainingType = trainingType

	duration, err := time.ParseDuration(strings.TrimSpace(slice[2]))
	if err != nil {
		return err
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	if t.Duration <= 0 {
		return "", errors.New("duration must be greater than 0")
	}

	distance := spentenergy.Distance(t.Steps)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64

	switch t.TrainingType {
	case "Бег":
		calories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	case "Ходьба":
		calories = spentenergy.WalkingSpentCalories(t.Steps, float64(t.Weight), float64(t.Height), t.Duration)
	default:
		return "", errors.New("unknown training type")
	}

	// Форматируем строку с результатами
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)

	return result, nil
}
