# go-exercise
Go程序设计语言练习
<br>

1.1 [完成](exercise-1.1/main.go)<br>
1.1 修改echo程序输出os.Args[0]，即命令的名字。<br>
1.1 Modify the echo program to also print os.Args[0], the name of the command that invoked it.<br>

1.2 [完成](exercise-1.2/main.go)<br>
1.2 修改echo程序，输出参数的索引和值，每行一个。<br>
1.2 Modify the echo program to print the index and value of each of its arguments, one per line.<br>

1.3 [未完成](exercise-1.3/main.go)<br>
1.3 尝试测量可能低效的程序和使用strings.Join的程序在执行时间上的差异。（1.6节有time包，11.4节展示如何撰写系统性的性能测试评估测试）<br>
1.3 Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)<br>

1.4 [未完成](exercise-1.4/main.go)<br>
1.4 修改dup2程序，输出出现重复行的文件的名称。<br>
1.4 Modify dup2 to print the names of all files in which each duplicated line occurs.<br>

2.1 [完成](exercise-2.1/main.go)<br>
添加类型、常量和函数到tempconv包中，处理以开尔文为单位（K）的温度值，0K＝－273.5°C，变化1K和变化1°C是等价的。<br>
Add types, constants, and functions to tempconv for pro cessing temperatures in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.<br>

4.3 [完成](exercise-4.3/main.go)<br>
4.3 重写函数reverse，使用数组指针作为参数而不是slice。<br>
4.3 Rewrite reverse to use an array pointer instead of a slice.<br>
4.3 Example: [1,2,3] to [3,2,1]

4.4 [完成](exercise-4.4/main.go)<br>
4.4 编写一个函数rotate，实现一次遍历就可以完成元素旋转。<br>
4.4 Write a version of rotate that operates in a single pass.<br>
4.4 Example: [1,2,3,4,5] to [3,4,5,1,2]

4.5 [完成](exercise-4.5/main.go)<br>
4.5 编写一个就地处理函数，用于去除[]sting slice中相邻的重复字符串元素。<br>
4.5 Write an in-place function to eliminate adjacent duplicates in a []string slice.<br>
4.5 Example: []string{"a", "a", "b", "a", "c", "b", "b", "b"} to []string{"a", "a", "b", "a", "c", "b"} in-place funcition 不生成中间变量，类似a[i],a[j]=a[j],a[i]

4.6 [未完成](exercise-4.6/main.go)<br>
4.6 编写一个就地处理函数，用于将一个UTF-8编码的字节slice中的所有相邻Unicode空白字符（查看unicode.IsSpace）缩减为一个ASCII空白字符。<br>
4.6 Write an in-place function that squashes each run of adjacent Unicode spaces(see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.

4.7 [未完成](exercise-4.7/main.go)<br>
4.7 Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string , in place. Can you do it without allocating new memory?