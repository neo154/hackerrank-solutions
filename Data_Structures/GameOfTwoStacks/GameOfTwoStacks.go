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
 * Complete the 'twoStacks' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER maxSum
 *  2. INTEGER_ARRAY a
 *  3. INTEGER_ARRAY b
 */


func twoStacks(maxSum int32, a []int32, b []int32) int32 {
    // Default case, no more possible moves
    counter := 0
    // Using index refs to avoid having to wait for slice operations
    a_index := 0
    b_index := 0
    a_len := len(a) //Really like max indices
    b_len := len(b)
    tmp_total := int32(0)
    // Getting initial totals using just one stack as the primary
    for (a_index < a_len && tmp_total + a[a_index] <= maxSum){
      tmp_total += a[a_index]
      a_index++
    }
    for (b_index < b_len && tmp_total + b[b_index] <= maxSum){
      tmp_total += b[b_index]
      b_index++
    }
    // Again, just initial run using primarily stack a
    counter = a_index + b_index
    // Now going to refine it
    for {
      // So remove a_index entries while we are not able to add next entry in b
      for (a_index > 0 && b_index < b_len && tmp_total + b[b_index] > maxSum) {
        // Decrement, a_index is in bounds and we are going to test tmp_totals using b
        a_index--
        tmp_total -= a[a_index]
      }
      // If we are at the end of all possible b entries after max number of a entries removed
      if (b_index==b_len || tmp_total + b[b_index] > maxSum){
        break;
      }
      // Now add b index entries
      tmp_total += b[b_index]
      b_index++
      if counter < a_index+b_index{
        counter = a_index+b_index
      }
    }
    return int32(counter)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    g := int32(gTemp)

    for gItr := 0; gItr < int(g); gItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        m := int32(mTemp)

        maxSumTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
        checkError(err)
        maxSum := int32(maxSumTemp)

        aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var a []int32

        for i := 0; i < int(n); i++ {
            aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
            checkError(err)
            aItem := int32(aItemTemp)
            a = append(a, aItem)
        }

        bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var b []int32

        for i := 0; i < int(m); i++ {
            bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
            checkError(err)
            bItem := int32(bItemTemp)
            b = append(b, bItem)
        }

        result := twoStacks(maxSum, a, b)

        fmt.Fprintf(writer, "%d\n", result)
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
