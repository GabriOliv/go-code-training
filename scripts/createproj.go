package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func verify_os_arg(array []string) (string, error) {
	if len(array) != 2 {
		return "", errors.New("wrong argument")
	}
	return array[1], nil
}

func verify_err(err error) {
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
		// panic(err)
	}
}

func main() {

	// verify arg number
	arg, err := verify_os_arg(os.Args)
	verify_err(err)

	// print arg passed
	fmt.Println(arg)

	// create project dir
	cmd := exec.Command("mkdir", arg)
	verify_err(cmd.Run())

	// change workdir
	verify_err(os.Chdir(arg))

	// get filepath
	wd, err := os.Getwd()
	verify_err(err)

	// run go init
	cmd = exec.Command("go", "mod", "init", filepath.Base(filepath.Dir(wd))+"/"+arg)
	verify_err(cmd.Run())

	// create main.go
	cmd = exec.Command("touch", "main.go", "README.md")
	verify_err(cmd.Run())
}
