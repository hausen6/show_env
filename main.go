package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var (
		envVar string
		sepVar string
	)
	flag.StringVar(&envVar, "v", "PATH", "search for environment variable")
	flag.StringVar(&sepVar, "s", "\n", "separate char")
	flag.Parse()

	env := os.Getenv(envVar)
	fmt.Fprintf(os.Stderr, "%s=\n", envVar)
	for _, v := range filepath.SplitList(env) {
		fmt.Print(v, sepVar)
	}
}
