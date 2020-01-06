package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	os.MkdirAll("./tmp", os.ModePerm)

	sem := make(chan struct{}, 100)
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i < 100; i++ {
		go worker(fmt.Sprintf("tmp/half_gopher_%d.png", i), &wg, sem)
	}

	wg.Wait()
	close(sem)

	elapsed := time.Since(start)
	log.Printf("Processsing took %s", elapsed)
}

func worker(filename string, wg *sync.WaitGroup, sem chan struct{}) {
	sem <- struct{}{}
	defer func() { <-sem }()
	defer wg.Done()

	cmd := "magick"
	args := []string{"convert", "../gopher.png", "-resize", "50%", filename}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
