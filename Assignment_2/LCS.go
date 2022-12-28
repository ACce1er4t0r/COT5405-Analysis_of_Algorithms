package main

import (
	"fmt"
	"strconv"
)

func lcs(str1 string, str2 string) (float64, string, string, [][]float64) {
	str1Len := len(str1)
	str2Len := len(str2)

	weight := map[rune]float64{
		'A': 0.08168, 'B': 0.01492, 'C': 0.02782, 'D': 0.04253, 'E': 0.12702, 'F': 0.02228,
		'G': 0.02015, 'H': 0.06094, 'I': 0.06966, 'J': 0.00153, 'K': 0.00772, 'L': 0.04025,
		'M': 0.02406, 'N': 0.06749, 'O': 0.07507, 'P': 0.01929, 'Q': 0.00095, 'R': 0.05987,
		'S': 0.06327, 'T': 0.09056, 'U': 0.02758, 'V': 0.00978, 'W': 0.02360, 'X': 0.00150,
		'Y': 0.01974, 'Z': 0.00074,
	}

	penalty, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(1)/26), 64)

	if str1Len > str1Len {
		str1, str2 = str2, str1
		str1Len, str2Len = str2Len, str1Len
	}
	if str1Len == 0 {
		return 0, "", "", nil
	}
	dp := make([][]float64, str1Len)
	for i := 0; i < str1Len; i++ {
		dp[i] = make([]float64, str2Len)
	}
	ret := 0.0
	x, y := 0, 0
	result1 := ""
	result2 := ""
	for i := 0; i < str1Len; i++ {
		if str1[i] == str2[0] {
			dp[i][0] = weight[rune(str1[i])]
		} else {
			dp[i][0] = -penalty
		}
	}
	for j := 1; j < str2Len; j++ {
		if str1[0] == str2[j] {
			dp[0][j] = weight[rune(str2[j])]
		} else {
			dp[0][j] = -penalty
		}
	}
	for i := 1; i < str1Len; i++ {
		for j := 1; j < str2Len; j++ {
			if str1[i] == str2[j] {
				dp[i][j] = weight[rune(str1[i])] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] - penalty
			}
			if ret < dp[i][j] {
				ret = dp[i][j]
				x, y = i, j
			}
		}
	}
	if x <= y {
		tmpx, tmpy, length := x, y, 0
		for {

			if tmpx != 0 {
				if dp[tmpx][tmpy] != penalty {
					length++
					tmpx--
					tmpy--
				} else {
					break
				}
			} else {
				length++
				break
			}
		}
		result1 = str1[x-length+1 : x+1]
		result2 = str2[y-length+1 : y+1]
	}
	if y < x {
		tmpx, tmpy, length := x, y, 0
		for {
			if tmpy != 0 {
				if dp[tmpx][tmpy] != penalty {
					length++
					tmpx--
					tmpy--
				} else {
					break
				}
			} else {
				length++
				break
			}
		}
		result1 = str1[x-length+1 : x+1]
		result2 = str2[y-length+1 : y+1]
	}

	return ret, result1, result2, dp
}
