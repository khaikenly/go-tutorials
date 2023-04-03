// Staircase detail

// This is a staircase of size n=4 :

//    #
//   ##
//  ###
// ####
// Its base and height are both equal to . It is drawn using # symbols and spaces. The last line is not preceded by any spaces.

// Write a program that prints a staircase of size .

// Function Description

// Complete the staircase function in the editor below.

// staircase has the following parameter(s):

// int n: an integer
// Print

// Print a staircase as described above.

// Input Format

// A single integer, , denoting the size of the staircase.

// Constraints

//  .

// Output Format

// Print a staircase of size  using # symbols and spaces.

// Note: The last line must have  spaces in it.

// Sample Input

// 6
// Sample Output

//      #
//     ##
//    ###
//   ####
//  #####
// ######
// Explanation

// The staircase is right-aligned, composed of # symbols and spaces, and has a height and width of .

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// func draw(s string, count int) string {
// 	if count == 0 {
// 		return ""
// 	}
// 	n := len(s) * count
// 	var b strings.Builder
// 	b.Grow(n)
// 	b.WriteString(s)
// 	for b.Len() < n {
// 		if b.Len() <= n/2 {
// 			b.WriteString(b.String())
// 		} else {
// 			b.WriteString(b.String()[:n-b.Len()])
// 			break
// 		}
// 	}
// 	return b.String()
// }

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */
func staircase(n int32) {
	// Write your code here
	var i int32
	var str string
	var j int32
	// for i = 1; i <= n; i++ {
	// 	// fmt.Println(draw(" ", int(n-i)) + draw("#", int(i)))
	// 	fmt.Println(strings.Repeat(" ", int(n-i)) + strings.Repeat("#", int(i)))
	// }
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			if j > (n - i) {
				str += "#"
			} else {
				str += " "
			}
		}
		str += "\n"
	}
	fmt.Println(str)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
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
