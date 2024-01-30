package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])
	if err != nil {
		return nil, err
	}

	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, err
	}

	sortTasks(tasks)
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	date, err := time.Parse("02.01.2006", src)
	if err != nil {
		return time.Time{}, errors.New("failed to parse date")
	}
	return date, nil
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	var tasks []Task
	taskMap := make(map[string]time.Duration)

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 4 {
			return nil, errors.New("invalid task format")
		}
		startTime, err := time.Parse("15:04", parts[0])
		if err != nil {
			return nil, errors.New("failed to parse task start time")
		}
		endTime, err := time.Parse("15:04", parts[2])
		if err != nil {
			return nil, errors.New("failed to parse task end time")
		}
		if endTime.Before(startTime) || endTime.Equal(startTime) {
			return nil, errors.New("end time is before or equal to start time")
		}
		duration := endTime.Sub(startTime)
		title := strings.Join(parts[3:], " ")

		if existingDur, ok := taskMap[title]; ok {
			taskMap[title] = duration + existingDur
		} else {
			taskMap[title] = duration
		}
	}

	for title, duration := range taskMap {
		task := Task{
			Date:  date,
			Dur:   duration,
			Title: title,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Dur > tasks[j].Dur
	})
}

// конец решения
// ::footer

func MainFormat() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}

/*
- Напряженная работа: 4h0m0s
- Напряженная работа: 4h0m0s
- Безудержное веселье: 3h0m0s
- Оглаживание кота: 1h0m0s
- Обед: 45m0s
- Завтрак: 30m0s
- Интернеты: 30m0s
- Интернеты: 30m0s
- Оглаживание кота: 30m0s
- Оглаживание кота: 15m0s
*/
