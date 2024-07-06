package main

import "fmt"

type EngineButton struct { // Button to execute a command
	cmd Command
}

func (eb EngineButton) Press() {
	eb.cmd.execute()
}

type Command interface {
	execute()
}

type StartEngine struct {
	en Engine
}

func (s StartEngine) execute() {
	s.en.Start()
}

type StopEngine struct {
	en Engine
}

func (s StopEngine) execute() {
	s.en.Stop()
}

type Engine struct {
	isRunning bool
}

func (e *Engine) Start() {
	e.isRunning = true
	fmt.Println("Engine started")
}

func (e *Engine) Stop() {
	e.isRunning = false
	fmt.Println("Engine stoped")
}

func main() {
	Engine := Engine{isRunning: false}
	start := StartEngine{en: Engine} // Command 1
	stop := StopEngine{en: Engine}   // Command 2

	button := EngineButton{cmd: start}
	button.Press()
	button = EngineButton{cmd: stop}
	button.Press()
}
