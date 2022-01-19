package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'largestRectangle' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY h as parameter.
 */

func largestRectangle(h []int32) int64 {
    var newHeight int32
    var tmp_area int64
    var incrementer int64
    var left int
    var right int
    currentHeight := int32(0)
    max := int64(0)
    l := len(h) // No length recalc required
    for i:=0; i < l; i++{
        newHeight = h[i]
        // Checking new possible height and currentHeight, otherwise just calc the same results
        if newHeight!=currentHeight{
            // Setting up for new calculations of area, including new left/right pos
            currentHeight = newHeight
            incrementer = int64(newHeight)
            tmp_area = incrementer
            left = i - 1
            right = i + 1
            // Go left and right as far as possible adding incrementer for each valid entry
            for (left>=0 && currentHeight<=h[left]){
                left--
                tmp_area += incrementer
            }
            for (right<l && currentHeight<=h[right]){
                right++
                tmp_area += incrementer
            }
            // Check new max
            if tmp_area > max{
                max = tmp_area
            }
        }
    }
    return max
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    hTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var h []int32

    for i := 0; i < int(n); i++ {
        hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
        checkError(err)
        hItem := int32(hItemTemp)
        h = append(h, hItem)
    }

    result := largestRectangle(h)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
