// Package printh is a helper for printing standardized messages to stdout.
package printh

import (
	"fmt"
	"log"
	"runtime"
)

// Info prints to stdout using log.Printf the string INFO, the package and the function who
// called it, and the messages m...
func Info(m ...interface{}) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.Printf("INFO  [%v] %v", frame.Function, fmt.Sprintln(m...))
}

// Debug checks if v == true and if it is, prints to stdout using log.Printf the string
// DEBUG, the package and the function who called it, and the messages m...
func Debug(v bool, m ...interface{}) {
	if v {
		pc := make([]uintptr, 15)
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		log.Printf("DEBUG [%v] %v", frame.Function, fmt.Sprintln(m...))
	}
}

// Err checks if err != nil and if it is, prints to stdout using log.Printf the string
// ERR, the package and the function who called it, err, and the messages m...
func Err(err error, m ...interface{}) {
	if err != nil {
		pc := make([]uintptr, 15)
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		log.Printf("ERR   [%v] (%v) %v", frame.Function, err, fmt.Sprint(m...))
	}
}

// ErrFatal checks if err != nil and if it is, prints to stdout using log.Fatalf the string
// ERR, the package and the function who called it, err, and the messages m...
func ErrFatal(err error, m ...interface{}) {
	if err != nil {
		pc := make([]uintptr, 15)
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		log.Fatalf("ERR   [%v] (%v) %v", frame.Function, err, fmt.Sprint(m...))
	}
}
