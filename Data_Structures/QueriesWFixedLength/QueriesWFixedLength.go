package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

/*
 * Complete the 'solve' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY arr
 *  2. INTEGER_ARRAY queries
 */


func runQuery(arr []int32, d int32) int32{
    // Setup some basic and reusable ariables
    var tmp_max int32
    var max_index int
    min := int32(math.MaxInt32)
    int_d := int(d)
    max_index = -1
    tmp_max = 0
    // For each array entry
    for i:=0; i < len(arr)-int_d+1; i++{
       // See if the max index is still ahead
       if (max_index < i ){
           // If its no longer in the window/infront, then reset tmp_max and get new one
           tmp_max = 0
           for j:=i; j < i+int_d; j++{
               if arr[j] > tmp_max {
                   // Again marking index ref for us to reuse
                   max_index = j
                   tmp_max = arr[j]
               }
           }
           // Check tmp_max
           if tmp_max < min {
               min = tmp_max
           }
       } else if tmp_max <= arr[i+int_d-1]{
           // If max_index is still in window, but the next entry is greater or equal
           // Reset index and values
           max_index = i+int_d-1
           tmp_max = arr[i+int_d-1]
           // Check tmp_max again
           if tmp_max < min {
               min = tmp_max
           }
       }
   }
   return min
}

func solve(arr []int32, queries []int32) []int32 {
    ret := make([]int32, len(queries))
    // Run query function for each one
    for i := 0; i < len(queries); i++{
        ret[i] = runQuery(arr, queries[i])
    }
    return ret
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    q := int32(qTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    var queries []int32

    for i := 0; i < int(q); i++ {
        queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        queriesItem := int32(queriesItemTemp)
        queries = append(queries, queriesItem)
    }

    result := solve(arr, queries)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
