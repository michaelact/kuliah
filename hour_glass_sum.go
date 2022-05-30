// Hacker Rank

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
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY[6][6] arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {
    // Find maksimal hour glass sum
    count := 0
    var maks int32
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            var result int32
            result += arr[i][j] + arr[i][j+1] + arr[i][j+2]
            result += arr[i+1][j+1]
            result += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
            
            if count == 0 || result > maks {
                maks = result
            }

            count++
        }
    }
    
    return maks
}
