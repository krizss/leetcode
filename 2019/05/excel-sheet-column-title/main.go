package main

import "fmt"

// 给定一个正整数，返回它在 Excel 表中相对应的列名称。

// 例如，

//     1 -> A
//     2 -> B
//     3 -> C
//     ...
//     26 -> Z
//     27 -> AA
//     28 -> AB
//     ...
// 示例 1:

// 输入: 1
// 输出: "A"
// 示例 2:

// 输入: 28
// 输出: "AB"
// 示例 3:

// 输入: 701
// 输出: "ZY"

func main() {
	fmt.Print(convertToTitle(701))
}

func convertToTitle(n int) string {
	if n < 1 {
		return ""
	}
	result := ""
	for n > 0 {
		n-- // 26进制数,但是这里映射是从1开始,所以要-1来对应计算机
		r := n % 26
		result = string('A'+r) + result
		n = n / 26
	}
	return result
}
