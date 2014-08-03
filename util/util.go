package util

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

var (
	Verbose     = false
	VeryVerbose = false
)

func Darwin() bool {
	b, _ := FileExists("/mach_kernel")
	return b
}

func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(data []byte) ([]string, error) {
	var lines []string
	scan := bufio.NewScanner(bytes.NewReader(data))
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines, scan.Err()
}

//
// Logging functions
//

// Uh oh, not good but not worthy of process death
func Warn(msg string, args ...interface{}) {
	log.Printf(msg+"\n", args...)
}

// Typical logging output, the default level
func Info(msg string, args ...interface{}) {
	log.Printf(msg+"\n", args...)
}

// -v: Verbosity level which helps track down production issues
func Debug(msg string, args ...interface{}) {
	if Verbose {
		log.Printf(msg+"\n", args...)
	}
}

// -V: Very verbose for development purposes
func DebugDebug(msg string, args ...interface{}) {
	if VeryVerbose {
		log.Printf(msg+"\n", args...)
	}
}
