package main

// Dont' know why, but couldn't import normally in batch in the UI this time
import "fmt"
import "bufio"
import "strconv"
import "strings"
import "os"
import "io"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var raw_query string
    var command int64
    var element int64

    // Basic reader and setup for the queues
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    q, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    // Setup slice to act as the queue and then looping through queries in stream
    queue := make([]int64, 0)
    for i:=int64(0); i < q; i++{
        // Get raw command and see if there is actually supposed to be another number
        raw_query = strings.TrimSpace(readLine(reader))
        command, err = strconv.ParseInt(raw_query[0:1], 10, 64)
        // Only the enqueue command (1) needs to read another int
        if command==1{
            // Enqueue
            element, err = strconv.ParseInt(strings.TrimSpace(raw_query[1:]), 10,64)
            checkError((err))
            queue = append(queue, element)
        } else if command==2{
            // Pop first element off, just resetting slice references
            queue = queue[1:]
        } else if command==3{
            // Print first element
            fmt.Println(queue[0])
        }
    }
}

// These helper functions are from other challenges, just too lazy to rewrite what works simply

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
