package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"

	"github.com/harukitosa/goaster/ast"

	"github.com/mattn/go-tty"
)

func initFlags() string {
	flag.Parse()
	args := flag.Args()
	filepath := ""
	if len(args) == 1 {
		filepath = args[0]
	} else {
		fmt.Printf("Err:%s\n", errors.New("Please enter a file name"))
		os.Exit(1)
	}
	return filepath
}

func openEditor(fileName string) error {
	editor := os.Getenv("EDITOR")
	log.Println(editor)
	if editor == "" {
		editor = "vi"
	}
	tty, err := tty.Open()
	if err != nil {
		return err
	}
	defer tty.Close()
	cmd := exec.Command(editor, fileName)
	cmd.Stdin = tty.Input()
	cmd.Stdout = tty.Output()
	cmd.Stderr = tty.Output()
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	filepath := initFlags()
	fset := token.NewFileSet()
	expr, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		panic(err)
	}
	bf := bytes.NewBufferString("")
	// ast.Fprint(bf, fset, expr, nil)
	ast.Fprint(bf, fset, expr, nil)
	f, err := os.Create(filepath + "ast")
	defer func() {
		f.Close()
	}()
	f.WriteString(string(bf.Bytes()))
	f.WriteString("\n")
	openEditor(f.Name())
}
