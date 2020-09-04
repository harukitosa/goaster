package goaster

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
	"path/filepath"

	"github.com/harukitosa/goaster/ast"
	"github.com/mattn/go-tty"
)

func initFlags() (string, bool) {
	flag.Parse()
	args := flag.Args()
	writeCmd := false
	filepath := ""
	var cmd string
	if len(args) == 1 {
		filepath = args[0]
	} else if len(args) == 2 {
		filepath = args[0]
		cmd = args[1]
	} else {
		fmt.Printf("Err:%s\n", errors.New("Please enter a file name"))
		os.Exit(1)
	}

	if cmd == "-w" {
		writeCmd = true
	}
	return filepath, writeCmd
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

func fileName(filePath string) (string, error) {
	if filePath == "" {
		return "", errors.New("no value")
	}
	e := filepath.Base(filePath)
	return e[0 : len(e)-3], nil
}

// Run starts goaster
// 抽象構文木からいじって直せたらよさそう
func Run() {
	filepath, writeCmd := initFlags()
	fset := token.NewFileSet()
	expr, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		panic(err)
	}
	bf := bytes.NewBufferString("")
	ast.Fprint(bf, fset, expr, nil)

	var f *os.File
	fname, err := fileName(filepath)
	f, err = os.Create("ast-" + fname)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		f.Close()
		if !writeCmd {
			os.Remove(f.Name())
		}
	}()
	f.WriteString(string(bf.Bytes()))
	f.WriteString("\n")

	if !writeCmd {
		openEditor(f.Name())
	}
}
