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
 * Complete the 'isBalanced' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isBalanced(s string) string {
    // Workingwith bytes is considerably faster than comparing strings at each phase
    var char byte
    var stack_end int
    // String isn't divisible by 2, then not balanced
    if len(s)%2!=0{
        return "NO"
    }
    l := len(s)
    stack := make([]byte, 0)
    // Now for each character
    for i:=0; i < l; i++{
        // Get character and see if it just needs to be put on the stack
        char = s[i]
        if char=='{' || char=='[' ||char=='('{
            stack = append(stack, char)
        } else{
            // Or if we need to do comparissons, if the stack is empty here though it is not balanced
            if len(stack)==0{
                return "NO"
            }
            stack_end = len(stack) - 1
            // If all of these different scenarios fail then it is not balanced
            if char=='}' && stack[stack_end]=='{'{
                stack = stack[:stack_end]
            } else if char==']' && stack[stack_end]=='['{
                stack = stack[:stack_end]
            } else if char==')' && stack[stack_end]=='('{
                stack = stack[:stack_end]
            } else {
                return "NO"
            }
        }
    }
    // If stack isn't empty, we have unbalanced brackets
    if len(stack)!=0{
        return "NO"
    }
    return "YES"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        s := readLine(reader)

        result := isBalanced(s)

        fmt.Fprintf(writer, "%s\n", result)
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
