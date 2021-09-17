package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "learn-pipeline-temp")

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFile()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func generateFile() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	//pipeline 1: job distribution
	chanFileIndex := generateFileIndexes()

	//pipeline 2:= the main logic (creating file)
	createFileWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFileWorker)

	counterTotal := 0
	counterSucces := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stacktrace : %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSucces++
		}
		counterTotal++
	}
	log.Printf("%d/%d of total files created", counterSucces, counterTotal)
}

func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()
	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	//wait group to control worker
	wg := new(sync.WaitGroup)

	//allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {
		//dispatch N worker
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				//listen to chanIn channel for incoming jobs
				for job := range chanIn {
					//do the job
					filepath := filepath.Join(tempPath, job.FileName)
					content := randomString(contentLength)
					err := ioutil.WriteFile(filepath, []byte(content), os.ModePerm)

					log.Println("Worker", workerIndex, "working on", job.FileName, "file generation")

					//construct the job result and send it to chanOut
					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}
				wg.Done()
			}(workerIndex)
		}
	}()

	//wait until chanIn close and then all workers are done
	//because after right that we need to close chaneOut channel
	go func() {
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}
