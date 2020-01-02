// channelsixtytwo
// Create ad-hoc VLC playlists

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	orderPtr := flag.Bool("o", false, "play episodes in order")
	flag.Parse()
	ordered := *orderPtr
	arg := flag.Args()[0]

	if ordered {
		watchDir(arg)
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
		paths := os.Args[1 : len(os.Args)-1]
		num := os.Args[len(os.Args)-1:]
		programs := randomPrograms(paths, num)
		for i := 0; i < len(programs); i++ {
			start(programs[i])
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isExt(path string, ext string) bool {
	if filepath.Ext(path) == ext {
		return true
	}
	return false
}

func viable(path string) bool {
	if isExt(path, ".mkv") || isExt(path, ".avi") || isExt(path, ".mp4") || isExt(path, ".m4v") {
		return true
	}
	return false
}

func randomPrograms(bases []string, num []string) []string {
	var programs []string
	x, err := strconv.Atoi(num[0])
	check(err)
	for i := x; i > 0; i-- {
		// Randomly pick a base
		base := bases[rand.Intn(len(bases))]
		programs = append(programs, randomProgram(base))
	}
	return programs
}

func randomProgram(base string) string {
	var program string
	suitors := make([]string, 0)
	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		check(err)
		if viable(path) {
			suitors = append(suitors, path)
		}
		return nil
	})
	check(err)
	// Randomly pluck a suitor, set it as program, and return it
	if len(suitors) > 1 {
		program = suitors[rand.Intn(len(suitors))]
	} else if len(suitors) == 1 {
		program = suitors[0]
	} else {
		fmt.Println("No playable files found in the given path")
		os.Exit(1)
	}
	return program
}

func watchDir(base string) {
	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		check(err)
		if viable(path) {
			start(path)
		}
		return nil
	})
	check(err)
}

func start(programs string) {
	vlc := filepath.Join("/Applications", "VLC.app", "Contents", "MacOS", "VLC")
	cmd := exec.Command(vlc, programs)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}
