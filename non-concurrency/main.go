package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	start := time.Now()
	os.MkdirAll("./tmp", os.ModePerm)

	for i := 1; i < 100; i++ {
		slowworker(fmt.Sprintf("tmp/half_gopher_%d.png", i))
	}
	elapsed := time.Since(start)
	log.Printf("Processsing took %s", elapsed)
}

func slowworker(filename string) {
	cmd := "magick"
	args := []string{"convert", "../gopher.png", "-resize", "50%", filename}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
