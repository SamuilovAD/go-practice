package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

type LogStats struct {
	InfoCount  int
	ErrorCount int
	UserLogins map[string]int
	syncMutex  sync.Mutex
}

/*
*
Task:
Build a concurrent log files parser that:
1. Reads one or more log files line-by-line concurrently and streams lines through a channel.
2. Processes lines with a pool of worker goroutines.
3. Tracks statistics in a shared structure protected by a mutex:
  - Total INFO lines.
  - Total ERROR lines.
  - Per-user login counts, where a user is identified by `user_id=<digits>` in a line.

4. Properly coordinates goroutines with WaitGroups and closes channels when reading is finished.
5. Prints a final report with the collected statistics.

Input example:
- Use the provided sample file `files/log1.txt` (adjust the path if you run from a different working directory).

Hints:
- Use `bufio.Scanner` for efficient reading.
- Use `regexp` to extract `user_id` values.
- Guard shared state with a `sync.Mutex`.

Expected output:
- Totals for INFO and ERROR.
- A list of `user_id=<id>: <count>` entries for users seen in the logs.
*/
func main() {
	filesNames := []string{getFilePath("files/log1.txt")}
	logLines := make(chan string, 100)
	var syncWaitGroup sync.WaitGroup
	stats := &LogStats{
		UserLogins: make(map[string]int),
	}
	for _, file := range filesNames {
		syncWaitGroup.Add(1)
		go func(filename string) {
			defer syncWaitGroup.Done()
			readLogFile(filename, logLines)
		}(file)
	}
	go func() {
		syncWaitGroup.Wait()
		close(logLines)
	}()
	numWorkers := 4
	var workersWg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			processLogs(logLines, stats)
		}()
	}

	workersWg.Wait()
	printStats(stats)
}

func readLogFile(filename string, out chan<- string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out <- scanner.Text()
	}
}

func processLogs(in <-chan string, stats *LogStats) {
	userIDRegex := regexp.MustCompile(`user_id=(\d+)`)
	for line := range in {
		stats.syncMutex.Lock()
		switch {
		case strings.Contains(line, "INFO"):
			stats.InfoCount++
			match := userIDRegex.FindStringSubmatch(line)
			if len(match) == 2 {
				stats.UserLogins[match[1]]++
			}
		case strings.Contains(line, "ERROR"):
			stats.ErrorCount++
		}
		stats.syncMutex.Unlock()
	}
}

func printStats(stats *LogStats) {
	fmt.Println("Log stats:")
	fmt.Println("INFO:", stats.InfoCount)
	fmt.Println("ERROR:", stats.ErrorCount)
	fmt.Println("User logins:")
	for userID, count := range stats.UserLogins {
		fmt.Printf("user_id=%s: %d\n", userID, count)
	}
}

func getFilePath(rel string) string {
	// Try path relative to this source file directory (works both in IDE run and `go run` from repo root)
	_, thisFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(thisFile)
	p := filepath.Join(baseDir, rel)
	if _, err := os.Stat(p); err == nil {
		return p
	}
	// Fallback: if the relative path exists from current working directory
	if _, err := os.Stat(rel); err == nil {
		return rel
	}
	// Last resort: join with current working directory
	cwd, err := os.Getwd()
	if err == nil {
		return filepath.Join(cwd, rel)
	}
	return rel
}
