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
 * Complete the 'queensAttack' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER r_q
 *  4. INTEGER c_q
 *  5. 2D_INTEGER_ARRAY obstacles
 */

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
    // Write your code here
    // var ret int32

    var up int32
    var down int32
    var left int32
    var right int32
    var up_left int32
    var up_right int32
    var down_left int32
    var down_right int32

    // Just setting up all possible maximum values for the board for queen
    up = n - r_q
    down = r_q - 1
    right = n - c_q
    left = c_q -1
    if up > left{
      up_left = left
    } else {
      up_left = up
    }
    if up > right{
      up_right = right
    } else {
      up_right = up
    }
    if down > left{
      down_left = left
    } else {
      down_left = down
    }
    if down > right{
      down_right = right
    } else {
      down_right = down
    }

    // Now going through obstacles to figure out if they block
    for i := 0; i < len(obstacles); i++{
      if obstacles[i][0]==r_q{
        // Same row, possibly relelvant

        // Checking columns to see if it parses down possible left or right
        if (obstacles[i][1] < c_q) && (c_q-obstacles[i][1]-1 < left){
          left = c_q - obstacles[i][1] - 1
        } else if obstacles[i][1]-c_q-1 < right{
          right = obstacles[i][1] - c_q - 1
        }
      } else if obstacles[i][1]==c_q{
        // Same column, could block up or down movement
        if (obstacles[i][0] < r_q) && (r_q-obstacles[i][0]-1 < down){
          down = r_q - obstacles[i][0] - 1
        } else if obstacles[i][0]-r_q-1 < up{
          up = obstacles[i][0] - r_q - 1
        }
      } else if (obstacles[i][0] - obstacles[i][1] == r_q - c_q){
        // Now checking if it blocks the dignal movement in up right direction of board
        // If they are subtracted with row and column and get teh same number they are in the same diagonal
        if (obstacles[i][1] < c_q) && (int32(math.Abs(float64(c_q-obstacles[i][1])))-1 < down_left){
          // Now just gotta see if it limits movement more than already there
          down_left = int32(math.Abs(float64(c_q-obstacles[i][1])))-1
        } else if int32(math.Abs(float64(obstacles[i][1]-c_q)))-1 < up_right{
          fmt.Println(obstacles[i][1], c_q)
          up_right = int32(math.Abs(float64(obstacles[i][1]-c_q)))-1
        }
      } else if (obstacles[i][0] + obstacles[i][1] == r_q + c_q){
        // Now checking the other direction of diagonals
        if (obstacles[i][1] < c_q) && (int32(math.Abs(float64(c_q-obstacles[i][1])))-1 < up_left){
          up_left = int32(math.Abs(float64(c_q-obstacles[i][1])))-1
        } else if int32(math.Abs(float64(obstacles[i][1]-c_q)))-1 < down_right{
          fmt.Println(obstacles[i][1], c_q)
          down_right = int32(math.Abs(float64(obstacles[i][1]-c_q)))-1
        }
      }
    }
    return (up+down+left+right+up_right+up_left+down_right+down_left)
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

    secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    r_qTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
    checkError(err)
    r_q := int32(r_qTemp)

    c_qTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
    checkError(err)
    c_q := int32(c_qTemp)

    var obstacles [][]int32
    for i := 0; i < int(k); i++ {
        obstaclesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var obstaclesRow []int32
        for _, obstaclesRowItem := range obstaclesRowTemp {
            obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
            checkError(err)
            obstaclesItem := int32(obstaclesItemTemp)
            obstaclesRow = append(obstaclesRow, obstaclesItem)
        }

        if len(obstaclesRow) != 2 {
            panic("Bad input")
        }

        obstacles = append(obstacles, obstaclesRow)
    }

    result := queensAttack(n, k, r_q, c_q, obstacles)

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
