/*
正三角形的外接圆面积
时间限制：1000 ms  |  内存限制：65535 KB
难度：0
描述
给你正三角形的边长，pi=3.1415926 ,求正三角形的外接圆面积。
输入
只有一组测试数据 第一行输入一个整数n(1<n<1000)表示接下来要输入n个边长m(1.0<=m<1000.0)
输出
输出每个正三角形的外接圆面积，保留两位小数，每个面积单独占一行。
样例输入
5
1
13
22
62
155
样例输出
1.05
176.98
506.84
4025.43
25158.92

PI * R * R / 3
 */
package main

import "fmt"

func main() {
	var n int
	var m float64
	pi := 3.1415926
	fmt.Scan(&n)
	for i := 0;i < n;i++ {
		fmt.Scan(&m)
		fmt.Printf("%.2f\n", pi * m * m / 3)
	}
}