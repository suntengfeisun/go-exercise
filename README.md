# go-exercise
Go程序设计语言练习
<br>

1.1 [完成](blob/master/exercise-1.1.go)<br>
1.1 修改echo程序输出os.Args[0]，即命令的名字。<br>
1.1 Modify the echo program to also print os.Args[0], the name of the command that invoked it.<br>

4.3 [完成](blob/master/exercise-4.3.go)<br>
4.3 重写函数reverse，使用数组指针作为参数而不是slice。<br>
4.3 Rewrite reverse to use an array pointer instead of a slice.<br>
4.3 Example: [1,2,3] to [3,2,1]

4.4 [完成](blob/master/exercise-4.4.go)<br>
4.4 编写一个函数rotate，实现一次遍历就可以完成元素旋转。<br>
4.4 Write a version of rotate that operates in a single pass.<br>
4.4 Example: [1,2,3,4,5] to [3,4,5,1,2]

4.5 [完成](blob/master/exercise-4.5.go)<br>
4.5 编写一个就地处理函数，用于去除[]sting slice中相邻的重复字符串元素。<br>
4.5 Write an in-place function to eliminate adjacent duplicates in a []string slice.<br>
4.5 Example: []string{"a", "a", "b", "a", "c", "b", "b", "b"} to []string{"a", "a", "b", "a", "c", "b"} in-place funcition 不生成中间变量，类似a[i],a[j]=a[j],a[i]

4.6 [未完成](blob/master/exercise-4.6.go)<br>
4.6 编写一个就地处理函数，用于将一个UTF-8编码的字节slice中的所有相邻Unicode空白字符（查看unicode.IsSpace）缩减为一个ASCII空白字符。<br>
4.6 Write an in-place function that squashes each run of adjacent Unicode spaces(see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.

4.7 [未完成](https://github.com/suntengfeisun/go-exercise/blob/master/exercise-4.7.go)<br>
4.7 Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string , in place. Can you do it without allocating new memory?