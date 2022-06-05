package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	listner, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("FTP server listening at 8080...")
	for {
		client, err := listner.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(client)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(c, err)
		return
	}
	curDir := homeDir

	for input.Scan() {

		cmd := NewCommandObj(input.Text())
		cmdArgs := cmd.Args

		switch cmd.Type {

		case "cd":
			newPath := filepath.Join(curDir, cmdArgs[0])
			err = os.Chdir(newPath)
			if err != nil {
				fmt.Fprintln(c, err.Error())
			} else {
				curDir = newPath
				fmt.Fprintln(c, curDir)
			}

		case "cwd":
			fmt.Fprintln(c, curDir)

		case "ls":

			dirs, err := os.ReadDir(curDir)
			if err != nil {
				fmt.Fprintln(c, err.Error())
			} else {
				for _, dir := range dirs {
					fmt.Fprintln(c, dir.Name())
				}
			}

		case "get":
			file := cmdArgs[0]
			dat, err := os.ReadFile(file)
			if err != nil {
				fmt.Fprintln(c, err.Error())
			} else {
				fmt.Fprintln(c, string(dat))
			}

		case "close":
			fmt.Fprintln(c, "close")
			c.Close()

		default:
			fmt.Fprintln(c, "Command not found!")
		}
	}
}
