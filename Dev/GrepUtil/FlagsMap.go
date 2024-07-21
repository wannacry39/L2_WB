package main

type GrepFlags struct {
	flagsMap map[string]struct{}
}

var flagsTable = GrepFlags{flagsMap: map[string]struct{}{
	"-A": {},
	"-B": {},
	"-C": {},
	"-c": {},
	"-i": {},
	"-v": {},
	"-F": {},
	"-n": {},
}}
