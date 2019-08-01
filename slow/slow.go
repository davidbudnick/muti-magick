package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func slowworker(filename string) {
	cmd := "magick"
	args := []string{"convert", "../doge.jpg", "-resize", "50%", filename}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	start := time.Now()
	os.MkdirAll("./tmp", os.ModePerm)

	for i := 1; i < 100; i++ {
		slowworker(fmt.Sprintf("tmp/half_doge_%d.jpg", i))
	}
	elapsed := time.Since(start)
	log.Printf("Processsing took %s", elapsed)
}
