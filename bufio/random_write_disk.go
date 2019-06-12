package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	GoroutineCount = 10

	KiloByte = 1024
	MByte = 1024 * 1024
)

func main() {
	// test case
	testCases := [][]int {
		{100, 10},      // TotalSize: 100MB, BlockSize: 10KB
		{100, 100},     // TotalSize: 100MB, BlockSize: 100KB
		{100, 512},     // TotalSize: 100MB, BlockSize: 512KB

		{500, 10},
		{500, 100},
		{500, 512},

		{1000, 10},
		{1000, 100},
		{1000, 512},

		{5000, 10},
		{5000, 100},
		{5000, 512},
		//
		//{10000, 10},
		//{10000, 100},
		//{10000, 512},
	}

	// print Table Header
	fmt.Println("| GoroutineIndex | TotalSize | BlockSize | TimeElapsed | WriteSpeed | TimeStamp |")


	var waitGroup sync.WaitGroup
	waitGroup.Add(GoroutineCount)

	startTime := time.Now()
	// execute case
	for index := 0; index < GoroutineCount; index++ {
		go func(index int, waitGroup *sync.WaitGroup) {
			// mark goroutine done
			defer waitGroup.Done()
			//
			for _, testCase := range testCases {
				totalSize := testCase[0]
				blockSize := testCase[1]
				randomWriteDisk(index, blockSize, totalSize)
			}
		}(index, &waitGroup)
	}

	// wait
	waitGroup.Wait()

	timeElapsed := time.Now().Sub(startTime).Seconds()
	fmt.Printf("Goroutine Count: %v, Total Time Elapsed: %vs\n", GoroutineCount, timeElapsed)
}

// blockSize uint as KB
// totalSize uint as MB
func randomWriteDisk(index, blockSize, totalSize int) {
	var message string
	for i := 0; i < blockSize * KiloByte / 10; i++ {
		message += "1234567890"
	}
	message += "\n"

	// create file
	fileName := fmt.Sprintf("file-%v.data", index)
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	// clean file
	defer f.Close()
	defer os.Remove(fileName)

	writer := bufio.NewWriter(f)
	round := totalSize * MByte / (blockSize * KiloByte)
	startTime := time.Now()
	for i := 0; i < round; i++ {
		writer.WriteString(message)
	}
	writer.Flush()

	// print statistic info
	timeElapsed := float64(time.Now().Sub(startTime).Nanoseconds()) / (1000 * 1000)
	writeSpeed := int64(float64(totalSize) / timeElapsed * 1000)
	fmt.Printf("| %v | %vMB \t| %vKB \t | %vms \t| %vMB/s \t| %v |\n", index,
		totalSize, blockSize, timeElapsed, writeSpeed, time.Now().Format("15:04:05"))
}


//| GoroutineIndex | TotalSize | BlockSize | TimeElapsed | WriteSpeed | TimeStamp |
//| 1 | 100MB 	| 10KB 	 | 77.795175ms 	| 1285MB/s 	| 22:52:23 |
//| 0 | 100MB 	| 10KB 	 | 80.249603ms 	| 1246MB/s 	| 22:52:23 |
//| 0 | 100MB 	| 100KB 	 | 76.684079ms 	| 1304MB/s 	| 22:52:23 |
//| 1 | 100MB 	| 100KB 	 | 74.998802ms 	| 1333MB/s 	| 22:52:23 |
//| 1 | 100MB 	| 512KB 	 | 67.737146ms 	| 1476MB/s 	| 22:52:27 |
//| 0 | 100MB 	| 512KB 	 | 62.327387ms 	| 1604MB/s 	| 22:52:27 |
//| 1 | 500MB 	| 10KB 	 | 342.947887ms 	| 1457MB/s 	| 22:52:27 |
//| 0 | 500MB 	| 10KB 	 | 345.553479ms 	| 1446MB/s 	| 22:52:27 |
//| 0 | 500MB 	| 100KB 	 | 326.370769ms 	| 1531MB/s 	| 22:52:28 |
//| 1 | 500MB 	| 100KB 	 | 418.584823ms 	| 1194MB/s 	| 22:52:29 |
//| 0 | 500MB 	| 512KB 	 | 584.172603ms 	| 855MB/s 	| 22:52:32 |
//| 1 | 500MB 	| 512KB 	 | 937.220374ms 	| 533MB/s 	| 22:52:33 |
//| 0 | 1000MB 	| 10KB 	 | 3064.351885ms 	| 326MB/s 	| 22:52:35 |
//| 1 | 1000MB 	| 10KB 	 | 4383.121417ms 	| 228MB/s 	| 22:52:43 |
//| 0 | 1000MB 	| 100KB 	 | 5293.352532ms 	| 188MB/s 	| 22:52:50 |
//| 1 | 1000MB 	| 100KB 	 | 7013.746168ms 	| 142MB/s 	| 22:52:57 |
//| 0 | 1000MB 	| 512KB 	 | 4920.002599ms 	| 203MB/s 	| 22:53:03 |
//| 1 | 1000MB 	| 512KB 	 | 761.772764ms 	| 1312MB/s 	| 22:53:08 |

//Goroutine Count: 2, Total Time Elapsed: 46.375617809s


//| GoroutineIndex | TotalSize | BlockSize | TimeElapsed | WriteSpeed | TimeStamp |
//| 3 | 100MB 	| 10KB 	 | 98.938239ms 	| 1010MB/s 	| 22:59:07 |
//| 2 | 100MB 	| 10KB 	 | 102.555209ms 	| 975MB/s 	| 22:59:07 |
//| 4 | 100MB 	| 10KB 	 | 104.169153ms 	| 959MB/s 	| 22:59:07 |
//| 0 | 100MB 	| 10KB 	 | 101.356851ms 	| 986MB/s 	| 22:59:07 |
//| 1 | 100MB 	| 10KB 	 | 102.127638ms 	| 979MB/s 	| 22:59:07 |
//| 0 | 100MB 	| 100KB 	 | 87.182837ms 	| 1147MB/s 	| 22:59:07 |
//| 2 | 100MB 	| 100KB 	 | 81.265069ms 	| 1230MB/s 	| 22:59:07 |
//| 3 | 100MB 	| 100KB 	 | 91.222511ms 	| 1096MB/s 	| 22:59:07 |
//| 1 | 100MB 	| 100KB 	 | 104.150638ms 	| 960MB/s 	| 22:59:07 |
//| 4 | 100MB 	| 100KB 	 | 102.058351ms 	| 979MB/s 	| 22:59:07 |
//| 0 | 100MB 	| 512KB 	 | 141.363577ms 	| 707MB/s 	| 22:59:16 |
//| 4 | 100MB 	| 512KB 	 | 127.911394ms 	| 781MB/s 	| 22:59:16 |
//| 3 | 100MB 	| 512KB 	 | 119.091218ms 	| 839MB/s 	| 22:59:16 |
//| 1 | 100MB 	| 512KB 	 | 121.978627ms 	| 819MB/s 	| 22:59:16 |
//| 2 | 100MB 	| 512KB 	 | 107.739003ms 	| 928MB/s 	| 22:59:16 |
//| 0 | 500MB 	| 10KB 	 | 6199.210773ms 	| 80MB/s 	| 22:59:22 |
//| 4 | 500MB 	| 10KB 	 | 7147.27286ms 	| 69MB/s 	| 22:59:23 |
//| 3 | 500MB 	| 10KB 	 | 7353.510043ms 	| 67MB/s 	| 22:59:23 |
//| 1 | 500MB 	| 10KB 	 | 7347.856216ms 	| 68MB/s 	| 22:59:23 |
//| 2 | 500MB 	| 10KB 	 | 10278.125139ms 	| 48MB/s 	| 22:59:27 |
//| 0 | 500MB 	| 100KB 	 | 5395.890947ms 	| 92MB/s 	| 22:59:29 |
//| 1 | 500MB 	| 100KB 	 | 3329.318946ms 	| 150MB/s 	| 22:59:36 |
//| 3 | 500MB 	| 100KB 	 | 3456.728744ms 	| 144MB/s 	| 22:59:37 |
//| 4 | 500MB 	| 100KB 	 | 3593.093279ms 	| 139MB/s 	| 22:59:37 |
//| 2 | 500MB 	| 100KB 	 | 3467.735176ms 	| 144MB/s 	| 22:59:37 |
//| 0 | 500MB 	| 512KB 	 | 3008.043239ms 	| 166MB/s 	| 22:59:38 |
//| 0 | 1000MB 	| 10KB 	 | 1529.426353ms 	| 653MB/s 	| 22:59:40 |
//| 0 | 1000MB 	| 100KB 	 | 1751.541394ms 	| 570MB/s 	| 22:59:44 |
//| 3 | 500MB 	| 512KB 	 | 878.865542ms 	| 568MB/s 	| 22:59:45 |
//| 1 | 500MB 	| 512KB 	 | 780.426795ms 	| 640MB/s 	| 22:59:45 |
//| 4 | 500MB 	| 512KB 	 | 778.928452ms 	| 641MB/s 	| 22:59:46 |
//| 2 | 500MB 	| 512KB 	 | 10833.286732ms 	| 46MB/s 	| 22:59:57 |
//| 3 | 1000MB 	| 10KB 	 | 19418.883656ms 	| 51MB/s 	| 23:00:05 |
//| 4 | 1000MB 	| 10KB 	 | 19401.104842ms 	| 51MB/s 	| 23:00:05 |
//| 0 | 1000MB 	| 512KB 	 | 26252.2851ms 	| 38MB/s 	| 23:00:14 |
//| 1 | 1000MB 	| 10KB 	 | 27264.302515ms 	| 36MB/s 	| 23:00:14 |
//| 2 | 1000MB 	| 10KB 	 | 24624.512467ms 	| 40MB/s 	| 23:00:21 |
//| 3 | 1000MB 	| 100KB 	 | 28970.649766ms 	| 34MB/s 	| 23:00:34 |
//| 4 | 1000MB 	| 100KB 	 | 28990.762506ms 	| 34MB/s 	| 23:00:34 |
//| 1 | 1000MB 	| 100KB 	 | 23077.029372ms 	| 43MB/s 	| 23:00:43 |
//| 2 | 1000MB 	| 100KB 	 | 19141.153417ms 	| 52MB/s 	| 23:00:44 |
//| 4 | 1000MB 	| 512KB 	 | 20294.278843ms 	| 49MB/s 	| 23:01:07 |
//| 3 | 1000MB 	| 512KB 	 | 20397.917654ms 	| 49MB/s 	| 23:01:07 |
//| 1 | 1000MB 	| 512KB 	 | 17486.899633ms 	| 57MB/s 	| 23:01:13 |
//| 2 | 1000MB 	| 512KB 	 | 15752.290922ms 	| 63MB/s 	| 23:01:15 |
//Goroutine Count: 5, Total Time Elapsed: 133.605472638s


// 3 * 100MB + 3 * 500MB + 3 * 1000MB
// Goroutine Count: 2, Total Time Elapsed: 46.375617809s
// Goroutine Count: 5, Total Time Elapsed: 133.605472638s	 => 180MB/s

// 3 * 100MB + 3 * 500MB + 3 * 1000MB + 3 * 5000MB
// Goroutine Count: 10, Total Time Elapsed: 1403.835882847s  => 140MB/s