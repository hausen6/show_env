package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func showEnvironmentVal(env string, sep string) int {
	environ := os.Getenv(env)
	fmt.Fprintf(os.Stderr, "%s=\n", env)
	for _, v := range filepath.SplitList(environ) {
		fmt.Print(v, sep)
	}
	return 0
}

func showEnvironmentList() int {
	envs := os.Environ()
	for _, e := range envs {
		p := strings.Split(e, "=")
		fmt.Println(p[0])
	}
	return 0
}

func main() {
	var (
		envVar  string
		sepVar  string
		listVar bool
	)
	flag.StringVar(&envVar, "v", "PATH", "search for environment variable")
	flag.StringVar(&sepVar, "s", "\n", "separate char")
	flag.BoolVar(&listVar, "l", false, "show list of environment variables")
	flag.Parse()

	var exitCode int = 0
	if listVar {
		exitCode = showEnvironmentList()
	} else {
		exitCode = showEnvironmentVal(envVar, sepVar)
	}
	os.Exit(exitCode)
}
