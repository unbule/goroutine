package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	findflag string
)

func init() {
	flag.StringVar(&findflag, "find", "1111", "find file")
}
func main() {
	// flag.Parse()
	files := directoryList()
	// fmt.Println(len(files))
	// fmt.Printf("%v", files)
	// findfileBynomal(files)
	// fmt.Println()
	// time.Sleep(1 * time.Second)
	// findfileBygo(files)
	// time.Sleep(1 * time.Second)
	// fmt.Println()
	// findfileBygotask(files)
	// time.Sleep(1 * time.Second)
	// fmt.Println()

	// findfileBygoGroup(files)
	// time.Sleep(1 * time.Second)

	// findfileMutex(files)
	// findfileBygotask(files)
	// findfileBygo(files)
	findfileBygoGroup(files)
}

func findfileBygo(files []string) {
	ch := make(chan struct{})
	for i := 0; i < len(files); i++ {
		go func(path string) {
			ch <- struct{}{}
			// fmt.Println(path)
			filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				fmt.Println(info.Name())
				return nil
			})
		}(files[i])
		<-ch
	}
}
func findfileBygoGroup(files []string) {
	wg := sync.WaitGroup{}
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(path string) {
			// fmt.Println(path)
			filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				// fmt.Println(info.Name())
				return nil
			})
			wg.Done()
		}(files[i])
	}
	wg.Wait()
}
func findfileBynomal(files []string) {
	for i := 0; i < len(files); i++ {
		// fmt.Println(files[i])
		filepath.Walk(files[i], func(path string, info os.FileInfo, err error) error {

			// fmt.Println(info.Name())
			return nil
		})

	}
}
func findfileBygotask(files []string) {
	ch := make(chan string, len(files))
	for i := 0; i < len(files); i++ {
		ch <- files[i]
		// fmt.Println("............")
	}

	for i := 0; i < len(files); i++ {
		// fmt.Println("11111111111")
		go func(h chan string) {
			path := <-ch
			// fmt.Println(path)
			// fmt.Println(path)
			filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				fmt.Println(info.Name())
				return nil
			})
		}(ch)
	}
}
func directoryList() []string {
	var files []string

	err := filepath.Walk("toys", func(path string, info os.FileInfo, err error) error {
		// fmt.Println(info.Name())
		if info.Name() == ".git" {
			return filepath.SkipDir
		}
		if info.IsDir() && info.Name() != "toys" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	// fmt.Println(files)
	return files
}

func findfileMutex(files []string) {
	// mu := &sync.Mutex{}
	for i := 0; i < len(files); i++ {
		go func(path string) {
			// mu.Lock()
			// fmt.Println(path)
			filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				fmt.Println(info.Name())
				return nil
			})
			// mu.Unlock()
		}(files[i])
	}
}
