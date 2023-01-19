package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var args = os.Args[1:]
	command := args[0]

	// todo: fix naive implementation, instead find the latest version of the app
	// and then launch that one
	// also fuck "run without debugger" and DLV in particular.
	if command == "d" {
		discordFile, err := filepath.Abs("C:/Users/keeb/AppData/Local/Discord/app-1.0.9006/Discord.exe")
		if err != nil {
			panic(err)
		}
		run(discordFile)

	}

	if command == "k" {
		run("k.exe")
	}
}

func run(file string) {
	p := exec.Command("cmd.exe", "/C", "start", "/b", file)
	p.Start()
}
