package main

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/gen2brain/beeep"
)

func notifyAnimal(animal string, wg *sync.WaitGroup) {
	defer wg.Done()
	iconPath, _ := filepath.Abs("icon.png")
	err := beeep.Notify("Уведомление о животном", fmt.Sprintf("%s  обработан!", animal), iconPath)
	if err != nil {
		fmt.Println("Ошибка при отправке уведомления:", err)
	}
}

func main() {
	var wg sync.WaitGroup
	animals := []string{"Лев", "Тигр", "Медведь"}

	for _, animal := range animals {
		wg.Add(1)
		go notifyAnimal(animal, &wg)
	}

	wg.Wait()
}
