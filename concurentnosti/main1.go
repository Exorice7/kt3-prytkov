package main

import (
	"fmt"
	"sync"
)

type Animal struct {
	Name              string
	HighSpeed         int
	Size              string
	ClimbTree         bool
	RecognizeDiseases bool
}

func boolToRussian(b bool) string {
	if b {
		return "Да"
	}
	return "Нет"
}

func animalInfo(animal Animal, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf(
		"Животное: %s, Скорость: %d км/ч, Размер: %s, Лазает по деревьям: %s, Распознаёт болезни: %s\n",
		animal.Name,
		animal.HighSpeed,
		animal.Size,
		boolToRussian(animal.ClimbTree),
		boolToRussian(animal.RecognizeDiseases),
	)
}

func main() {
	var wg sync.WaitGroup
	animals := []Animal{
		{"Гепард", 120, "Средний", false, false},
		{"Слон", 25, "Большой", false, true},
		{"Обезьяна", 30, "Средний", true, false},
		{"Собака", 45, "Средний", false, true},
		{"Кошка", 30, "Маленький", true, false},
	}

	for _, animal := range animals {
		wg.Add(1)
		go animalInfo(animal, &wg)
	}

	wg.Wait()
}
