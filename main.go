package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing paths of images to cat")
		os.Exit(2)
	}

	for _, path := range os.Args[1:] {
		if err := cat(path); err != nil {
			fmt.Fprintf(os.Stderr, "could not cat %s: %v\n", path, err)
		}
	}
}

func cat(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "could not open image")
	}
	defer f.Close()

	fmt.Printf("\033]1337;File=inline=1:")
	wc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	_, err = io.Copy(wc, f)
	if err != nil {
		return errors.Wrap(err, "could not encode image")
	}
	if err := wc.Close(); err != nil {
		return errors.Wrap(err, "could not close base64 encoder")
	}
	fmt.Printf("\a\n")

	return nil
}
