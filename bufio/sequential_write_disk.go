package main

import (
    "fmt"
)
import (
    "bufio"
    "os"
    "time"
)


const (
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
    fmt.Println("| TotalSize | BlockSize | TimeElapsed | WriteSpeed |")

    startTime := time.Now()
    // execute case
    for times := 0; times < 5; times++ {
        for _, testCase := range testCases {
            totalSize := testCase[0]
            blockSize := testCase[1]
            sequentialWriteDisk(blockSize, totalSize)
        }
    }

    timeElapsed := time.Now().Sub(startTime).Seconds()
    fmt.Printf("Repeate Count: %v, Total Time Elapsed: %vs\n", 5, timeElapsed)
}

// blockSize uint as KB
// totalSize uint as MB
func sequentialWriteDisk(blockSize, totalSize int) {
    var message string
    for i := 0; i < blockSize * KiloByte / 10; i++ {
        message += "1234567890"
    }
    message += "\n"

    // create file
    f, err := os.OpenFile("server", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    // clean file
    defer f.Close()
    defer os.Remove("server")

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
    fmt.Printf("| %vMB \t| %vKB \t | %vms \t| %vMB/s \t|\n", totalSize, blockSize, timeElapsed, writeSpeed)
}


//| TotalSize | BlockSize | TimeElapsed | WriteSpeed |
//| 100MB 	| 10KB 	 | 68.517974ms 	| 1459MB/s 	|
//| 100MB 	| 100KB 	 | 59.41526ms 	| 1683MB/s 	|
//| 100MB 	| 512KB 	 | 59.024536ms 	| 1694MB/s 	|
//| 500MB 	| 10KB 	 | 287.749706ms 	| 1737MB/s 	|
//| 500MB 	| 100KB 	 | 286.903992ms 	| 1742MB/s 	|
//| 500MB 	| 512KB 	 | 289.942947ms 	| 1724MB/s 	|
//| 1000MB 	| 10KB 	 | 609.795497ms 	| 1639MB/s 	|
//| 1000MB 	| 100KB 	 | 614.495982ms 	| 1627MB/s 	|
//| 1000MB 	| 512KB 	 | 1056.387148ms 	| 946MB/s 	|
//| 5000MB 	| 10KB 	 | 27639.138991ms 	| 180MB/s 	|
//| 5000MB 	| 100KB 	 | 26957.288659ms 	| 185MB/s 	|
//| 5000MB 	| 512KB 	 | 26577.727921ms 	| 188MB/s 	|


//| TotalSize | BlockSize | TimeElapsed | WriteSpeed |
//| 100MB 	| 10KB 	 | 74.425431ms 	| 1343MB/s 	|
//| 100MB 	| 100KB 	 | 57.636418ms 	| 1735MB/s 	|
//| 100MB 	| 512KB 	 | 58.644574ms 	| 1705MB/s 	|
//| 500MB 	| 10KB 	 | 285.357847ms 	| 1752MB/s 	|
//| 500MB 	| 100KB 	 | 283.075808ms 	| 1766MB/s 	|
//| 500MB 	| 512KB 	 | 286.254005ms 	| 1746MB/s 	|
//| 1000MB 	| 10KB 	 | 647.823166ms 	| 1543MB/s 	|
//| 1000MB 	| 100KB 	 | 882.224232ms 	| 1133MB/s 	|
//| 1000MB 	| 512KB 	 | 597.492461ms 	| 1673MB/s 	|
//| 5000MB 	| 10KB 	 | 27559.732989ms 	| 181MB/s 	|
//| 5000MB 	| 100KB 	 | 28170.354979ms 	| 177MB/s 	|
//| 5000MB 	| 512KB 	 | 27372.650611ms 	| 182MB/s 	|
//| 10000MB 	| 10KB 	 | 70397.766169ms 	| 142MB/s 	|
//| 10000MB 	| 100KB 	 | 71313.955531ms 	| 140MB/s 	|
//| 10000MB 	| 512KB 	 | 63221.576177ms 	| 158MB/s 	|


// Repeat Times: 2, Total Time Elapsed: 25.84726539s    => 372MB/s
// Repeat Times: 5, Total Time Elapsed: 75.342777636s   => 320MB/s