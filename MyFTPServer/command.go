package main

import "strings"

type Command struct {
	Type string
	Args []string
}

func NewCommandObj(cmd string) Command {
	cmds := strings.Split(cmd, " ")
	coommand := Command{
		cmds[0],
		cmds[1:],
	}
	return coommand
}
