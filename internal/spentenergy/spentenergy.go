package spentenergy

import (
	"errors"
	"time"
)

// Константы из шаблона. Если в твоём шаблоне значения другие — оставь свои.
const (
	stepLengthCoefficient      = 0.45 // коэффициент длины шага относительно роста
	walkingCaloriesCoefficient = 0.5  // корректирующий коэффициент калорий при ходьбе
	mInKm                      = 1000 // метров в километре
	minInH                     = 60   // минут в часе
)

// Distance возвращает дистанцию в километрах по количеству шагов и росту (м).
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

// MeanSpeed возвращает среднюю скорость в км/ч.
// При некорректных шагах или нулевой длительности возвращает 0.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// RunningSpentCalories возвращает калории, потраченные при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validate(steps, weight, height, duration); err != nil {
		return 0, err
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

// WalkingSpentCalories возвращает калории, потраченные при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validate(steps, weight, height, duration); err != nil {
		return 0, err
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories * walkingCaloriesCoefficient, nil
}

// validate проверяет корректность входных параметров для расчёта калорий.
func validate(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return errors.New("количество шагов должно быть больше 0")
	case weight <= 0:
		return errors.New("вес должен быть больше 0")
	case height <= 0:
		return errors.New("рост должен быть больше 0")
	case duration <= 0:
		return errors.New("продолжительность должна быть больше 0")
	}
	return nil
}
