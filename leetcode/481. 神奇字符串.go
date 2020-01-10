package main

import (
	"strings"
	"fmt"
)

/*
网上看了很多篇博客，都是没有详细规律，让人看得云里雾里的，越看越😵

规律：初始的为 122，从初始的开始生成后面的字符串，先字符1，2，1，2这种交替添加
之后一位是数组下标3，下标2的值是2，所以添加两个1，之后是下标3是1，后面添加一个2

详细的规律生成如下：

122              要添加的值  要添加的个数
12211            1          下标2的值为2
122112           2          下标3的值为1
1221121          1          下标4的值为1
122112122        2          下标5的值为2
1221121221       1          下标6的值为1
...

*/
func magicalString(n int) int {
	s := "122"
	for i, j, c := 2, 0, '1';n > len(s); {
		for j = 0;j < int(s[i] - '0');j++ {
			s += string(c)
		}
		if c == '1' {
			c = '2'
		} else {
			c = '1'
		}
		i++
	}
	return strings.Count(s[:n], "1")
}

func main() {
	fmt.Println(magicalString(5))
}