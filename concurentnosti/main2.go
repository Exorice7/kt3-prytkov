package main

import (
	"fmt"
	"sync"
	"time"
)

type Animal struct {
	Name              string
	HighSpeed         int
	Size              string
	ClimbTree         bool
	RecognizeDiseases bool
}

func loadAnimalInfo(animalName string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	animal := Animal{
		Name:              animalName,
		HighSpeed:         60,
		Size:              "Средний",
		ClimbTree:         true,
		RecognizeDiseases: false,
	}
	fmt.Printf("Загружена информация о животном: %s\n", animal.Name)
}

func main() {
	var wg sync.WaitGroup
	animals := []string{"Лев", "Тигр", "Медведь", "Волк", "Лиса"}

	for _, animal := range animals {
		wg.Add(1)
		go loadAnimalInfo(animal, &wg)
	}

	wg.Wait()
}
