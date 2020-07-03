package main

import (
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

func main() {
	// println("starting...")
	//fmt.Println(os.Args)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage: gotestfile.exe path/file_test.go . . . ")
		return
	}
	err := process(args)
	if err != nil {
		log.Fatal(err)
	}

}

func process(files []string) error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "os.getwd")
	}
	//fmt.Printf("dir:   %s\n", wd)

	m, err := parse.Module()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("mod:   %s\n", m)

	for _, f := range files {
		rel := filepath.Dir(f)
		//fmt.Printf("rel:   %s\n", rel)
		fq := filepath.Join(wd, f)
		//fmt.Printf("file:  %s\n", fq)
		//pkg := filepath.Dir(fq)
		//fmt.Printf("pkg:   %s\n", pkg)

		names, err := parse.Tests(fq, nil)
		if err != nil {
			return errors.Wrap(err, "tests")
		}

		//fmt.Printf("tests: %v\n", names)

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
