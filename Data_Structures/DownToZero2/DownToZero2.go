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
 * Complete the 'downToZero' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

// Builds Cache reference from 0 to the MAX using contstraints on the page
func buildRefCache() []int32{
    var tmp_score int32
    var tmp_limit int32
    MAX := int32(1000000)
    SQRT_MAX := int32(1000)
    queue := make([]int32, MAX+1)
    // Worst case scenarios initially
    for i:=int32(0); i < MAX+1; i++{
        queue[i]=i
    }
    // Now going to max
    for i:=int32(2); i < MAX; i++{
        // Getting new score and seeing if it is better than the next entrie's
        tmp_score = queue[i] + 1
        if queue[i+1] > tmp_score{
            queue[i+1] = tmp_score
        }
        // Setting limits for inner loop searching i+(i*x) space
        if i > SQRT_MAX{
            tmp_limit = MAX
        } else {
            tmp_limit = i * i
        }
        // Now go in increments of i, which means its a multiple and swap scores if it is more effecient
        for j := i+i; j <= tmp_limit; j+=i{
            if queue[j] > tmp_score{
                queue[j] = tmp_score
            }
        }
    }
    return queue
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)
    // Rewrote just tad what was here to use caching to avoid timeouts
    // felt stupid but put this cache build in the wrong place initially
    mapping := buildRefCache()
    for qItr := 0; qItr < int(q); qItr++ {
        nTemp, _ := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        fmt.Fprintf(writer, "%d\n", mapping[n])
    }
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
