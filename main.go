package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/billgraziano/gotestfile/parse"
	"github.com/pkg/errors"
)

var debug bool

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage: gotestfile.exe path/file_test.go . . . ")
		return
	}

	var prefix = flag.String("prefix", "", "list environment variables with this prefix")
	flag.BoolVar(&debug, "debug", false, "enable debug printing")
	flag.Parse()
	if prefix != nil {
		printenv(*prefix)
	}
	err := process(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
}

func process(files []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "os.getwd")
	}
	if debug {
		fmt.Printf("dir:   %s\n", wd)
	}

	m, err := parse.Module()
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		fmt.Printf("mod:   %s\n", m)
	}

	for _, f := range files {
		rel := filepath.Dir(f)
		if debug {
			fmt.Printf("rel:   %s\n", rel)
		}
		fq := filepath.Join(wd, f)
		if debug {
			fmt.Printf("file:  %s\n", fq)
		}

		names, err := parse.Tests(fq, nil)
		if err != nil {
			return errors.Wrap(err, "tests")
		}

		if debug {
			fmt.Printf("tests: %v\n", names)
		}

		testRegex := fmt.Sprintf("^(%s)$", strings.Join(names, "|"))
		parms := []string{"test", "-timeout", "30s"}
		parms = append(parms, path.Join(m, rel))
		parms = append(parms, "-count=1")
		parms = append(parms, "-p=1")
		parms = append(parms, "-run")
		parms = append(parms, testRegex)
		fmt.Printf("cmd: go.exe %s\n", strings.Join(parms, " "))

		out, err := exec.Command("go.exe", parms...).Output()
		if err != nil {
			log.Fatal(err)
		}
		str := strings.TrimSpace(string(out))
		fmt.Println(str)

	}
	return nil
}

// print env variables that match a prefix
func printenv(prefix string) {
	if prefix == "" {
		return
	}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.HasPrefix(pair[0], prefix) {
			fmt.Println(e)
		}
	}
}
