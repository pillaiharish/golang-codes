package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func getMemInfo() (map[string]int64, error) {
	meminfo := make(map[string]int64)
	if runtime.GOOS == "linux" {
		file, err := os.Open("/proc/meminfo")
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(file)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		for scanner.Scan() {
			split_text := strings.Split(scanner.Text(), ": ")
			text_int := strings.Trim(strings.Split(split_text[1], "kB")[0], " ")
			meminfo[split_text[0]], err = strconv.ParseInt(text_int, 10, 64)
			if err != nil {
				fmt.Println("String to integer parsing failed", err)
			}
		}
	} else if runtime.GOOS == "darwin" {
		outFile, err := os.Create("vm_stat_output.txt") // file pointer
		if err != nil {
			panic(err)
		}
		cmd := exec.Command("vm_stat")
		cmd.Stdout = outFile // file pointer will point to the stdout data
		fmt.Println(outFile)
		error := cmd.Run()
		if error != nil {
			panic(error)
		}
		file, err := os.Open("vm_stat_output.txt")
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(file)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		isFirstLine := true // to detect first line
		for scanner.Scan() {
			if isFirstLine {
				isFirstLine = false
				continue // skip the first line
			}

			split_text := strings.Split(scanner.Text(), ": ")
			trim_space := strings.TrimSpace(split_text[1])
			trim_dot := strings.Trim(trim_space, ".")
			meminfo[split_text[0]], err = strconv.ParseInt(trim_dot, 10, 64)
			if err != nil {
				fmt.Println("String to integer parsing failed", err)
			}
		}

	}
	// fmt.Println(meminfo)
	return meminfo, nil
}

func main() {
	memInfo, err := getMemInfo()
	if err != nil {
		fmt.Println("Error in meminfo", err)
	}
	// fmt.Println(meminfo)
	if runtime.GOOS == "linux" {
		memFree := memInfo["MemFree"]
		buffers := memInfo["Buffers"]
		cached := memInfo["Cached"]
		totalFree := memFree + buffers + cached

		fmt.Printf("MemFree: %d kB\n", memFree)
		fmt.Printf("Buffers: %d kB\n", buffers)
		fmt.Printf("Cached: %d kB\n", cached)
		fmt.Printf("Total free memory: %d kB\n", totalFree)
	} else if runtime.GOOS == "darwin" {
		var memSizePerPage int64 = 4096 // 4 KiB (kibibytes)
		pagesFree := memInfo["Pages free"]
		fileBackedPages := memInfo["File-backed pages"]                      // buffers
		pagesOccupiedByCompressor := memInfo["Pages occupied by compressor"] // cached
		pagesInactive := memInfo["Pages inactive"]
		pagesSpeculative := memInfo["Pages speculative"]

		//Total Free Memory = (Pages free + Pages inactive + Pages speculative) * Memory size per page
		totalFreeMem := (pagesFree + pagesInactive + pagesSpeculative) * memSizePerPage
		fmt.Printf("MemFree: %d B\n", pagesFree)
		fmt.Printf("Buffers: %d B\n", fileBackedPages)
		fmt.Printf("Cached: %d B\n", pagesOccupiedByCompressor)
		fmt.Printf("Total free memory: %d B\n", totalFreeMem)

	}
}
