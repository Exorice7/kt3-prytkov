package main

import (
	"fmt"
	"sync"
	"time"
)

func makeSound(animal string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("%s издает звук!\n", animal)
}

func main() {
	var wg sync.WaitGroup
	animals := []string{"Собака", "Кошка", "Корова", "Овца", "Утка"}

	for _, animal := range animals {
		wg.Add(1)
		go makeSound(animal, &wg)
	}

	wg.Wait()
}
