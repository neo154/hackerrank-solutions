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
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func nonDivisibleSubset(k int32, s []int32) int32 {
    ret := int32(0)
    tmp_slice := make([]int32, k)
    for i:=0; i<len(s); i++{
        tmp_slice[s[i]%k] ++
    }
    fmt.Println(tmp_slice)
    // Just getting first entry and seeing if it is not 0
    if tmp_slice[0] > 0{
        ret ++
    }
    var tmp_index int
    // Now stepping through from front and back starting at 1 to see which is greater
    for i:=1; i*2<=int(k); i++{
        // Easier reference here
        tmp_index = (int(k) - i) % int(k)
        // If slice length is even and last entry is greater than 0, increment
        if (i==tmp_index)&&(tmp_slice[tmp_index]>0){
            ret ++
        } else {
            // Otherwise just add which ever is greater
            if tmp_slice[i] >= tmp_slice[tmp_index]{
                ret += tmp_slice[i]
            }else{
                ret += tmp_slice[tmp_index]
            }
        }
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

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var s []int32

    for i := 0; i < int(n); i++ {
        sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
        checkError(err)
        sItem := int32(sItemTemp)
        s = append(s, sItem)
    }

    result := nonDivisibleSubset(k, s)

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
