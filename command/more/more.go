package main

import (
	"bufio"
	"fmt"
	"os"
)

const PAGELEN = 24

func main() {
	if len(os.Args) == 1 {
		doMore(os.Stdin)
	} else {
		f, err := os.Open(os.Args[1])
		if err != nil {
			os.Exit(1)
		}
		defer f.Close()

		doMore(f)
	}
}

// 处理stdin传过来的数据
func doMore(file *os.File) {
	scanner := bufio.NewScanner(file)

	tty, err := os.Open("/dev/tty")
	if err != nil {
		fmt.Print(err)
		return
	}

	lineNum := 0
	for {
		if !scanner.Scan() {
			break
		}

		if lineNum == PAGELEN {
			reply := seeMore(tty)
			if reply == 0 {
				break
			}

			lineNum -= reply
		}

		os.Stdout.Write(scanner.Bytes())
		os.Stdout.Write([]byte("\n"))

		lineNum++
	}
}

func seeMore(file *os.File) int {
	reader := bufio.NewReader(file)
	fmt.Print("more?")
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return 0
		}
		s = string(s[0])

		switch s {
		case "q":
			return 0
		case "n":
			return PAGELEN
		case "\n":
			return 1
		default:
			return 0
		}

		return 0
	}
}
