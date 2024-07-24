package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/shirou/gopsutil/process"
)

func main() {
	for {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		cmd := ""
		args := ""
		fmt.Printf("%s@%s#: ", pwd, usr.Username)
		fmt.Scan(&cmd)

		switch {
		case cmd == "pwd":
			fmt.Println(pwd)
			continue
		case cmd == "cd":
			fmt.Scanln(&args)
			os.Chdir(args)
			continue
		case cmd == "echo":
			input := bufio.NewReader(os.Stdin)
			line, _ := input.ReadString('\n')
			fmt.Printf("%v", line)
			continue
		case cmd == "ps":
			processes, err := process.Processes()
			if err != nil {
				log.Fatalf("Error getting processes: %v", err)
			}
			fmt.Printf("%-10s %-25s %-10s %-10s\n", "PID", "Name", "CPU%", "Memory%")
			fmt.Println("---------------------------------------------------------------")
			for _, proc := range processes {
				pid := proc.Pid
				name, err := proc.Name()
				if err != nil {
					name = "Unknown"
				}

				cpuPercent, err := proc.CPUPercent()
				if err != nil {
					cpuPercent = 0
				}

				memPercent, err := proc.MemoryPercent()
				if err != nil {
					memPercent = 0
				}

				fmt.Printf("%-10d %-25s %-10.2f %-10.2f\n", pid, name, cpuPercent, memPercent)
			}
		case cmd == "kill":
			var pid int
			fmt.Scan(&pid)
			proc, err := process.NewProcess(int32(pid))
			if err != nil {
				log.Fatal(err)
			}
			err = proc.Terminate()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Some errors while killing process pid: %v", pid)
			}
			continue
		case cmd == "quit":
			return
		default:
			continue
		}
	}
}
