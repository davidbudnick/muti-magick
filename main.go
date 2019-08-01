package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func worker(filename string, wg *sync.WaitGroup, sem chan struct{}) {
	sem <- struct{}{}
	defer func() { <-sem }()
	defer wg.Done()

	cmd := "magick"
	args := []string{"convert", "doge.jpg", "-resize", "50%", filename}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	start := time.Now()
	os.MkdirAll("./tmp", os.ModePerm)

	sem := make(chan struct{}, 100)
	l := 10
	var wg sync.WaitGroup
	wg.Add(l)

	for i := 1; i < 100; i++ {
		go worker(fmt.Sprintf("tmp/half_doge_%d.jpg", i), &wg, sem)
	}
	wg.Wait()
	close(sem)

	elapsed := time.Since(start)

	log.Printf("Processsing took %s", elapsed)
}
