# go-exercise
Go程序设计语言练习
<br>

1.1 [完成](exercise-1.1/main.go)<br>
修改echo程序输出os.Args[0]，即命令的名字。<br>
Modify the echo program to also print os.Args[0], the name of the command that invoked it.<br>

1.2 [完成](exercise-1.2/main.go)<br>
修改echo程序，输出参数的索引和值，每行一个。<br>
Modify the echo program to print the index and value of each of its arguments, one per line.<br>

1.3 [未完成](exercise-1.3/main.go)<br>
尝试测量可能低效的程序和使用strings.Join的程序在执行时间上的差异。（1.6节有time包，11.4节展示如何撰写系统性的性能测试评估测试）<br>
Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)<br>

1.4 [未完成](exercise-1.4/main.go)<br>
修改dup2程序，输出出现重复行的文件的名称。<br>
Modify dup2 to print the names of all files in which each duplicated line occurs.<br>

1.5 [未完成](exercise-1.5/main.go)<br>

1.6 [未完成](exercise-1.6/main.go)<br>

1.7 [未完成](exercise-1.7/main.go)<br>

1.8 [未完成](exercise-1.8/main.go)<br>

1.9 [未完成](exercise-1.9/main.go)<br>

1.10 [未完成](exercise-1.10/main.go)<br>

1.11 [未完成](exercise-1.11/main.go)<br>

1.12 [未完成](exercise-1.12/main.go)<br>

2.1 [完成](exercise-2.1/main.go)<br>
添加类型、常量和函数到tempconv包中，处理以开尔文为单位（K）的温度值，0K＝－273.5°C，变化1K和变化1°C是等价的。<br>
Add types, constants, and functions to tempconv for pro cessing temperatures in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.<br>

2.1 [完成](exercise-2.1/main.go)<br>
添加类型、常量和函数到tempconv包中，处理以开尔文为单位（K）的温度值，0K＝－273.5°C，变化1K和变化1°C是等价的。<br>
Add types, constants, and functions to tempconv for pro cessing temperatures in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.<br>

2.2 [完成](exercise-2.2/main.go)<br>
写一个类似于cf的通用的单位转换程序，从命令行参数或者标准输入（如果没有参数）获取数字，然后将每一个数字转换为以摄氏温度和华氏温度表示的温度，以英寸和米表示的长度单位，以磅和千克表示的重量，等等。<br>
Write a general-purpose unit-conversion program analogous to cf that reads
numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahren heit, length in feet and meters, weight in pounds and kilograms, and the like.<br>

2.3 [未完成](exercise-2.3/main.go)<br>
使用循环重写PopCount，来替代单个表达式。对比两个版本的效率（11.4节会展示如何系统性的对比不同实现的性能。）<br>
Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions. (Section 11.4 shows how to compare the performance of different implementations systematically.)<br>

2.4 [未完成](exercise-2.4/main.go)<br>

2.5 [未完成](exercise-2.5/main.go)<br>

3.1 [未完成](exercise-3.1/main.go)<br>
假如函数f返回一个float64型的无穷大值，就会导致SVG文件含有无效的<polygon>元素（尽管很多SVG绘图程序对此处理得当）。修改本程序以避免无效的多边形。<br>
If the function f returns a non-finite float64 value, the SVG file will contain invalid <polygon> elements (although many SVG renderers handle this gracefully). Modify the program to skip invalid polygons.

3.2 [未完成](exercise-3.2/main.go)<br>

3.3 [未完成](exercise-3.3/main.go)<br>

3.4 [未完成](exercise-3.4/main.go)<br>

3.5 [未完成](exercise-3.5/main.go)<br>

3.6 [未完成](exercise-3.6/main.go)<br>

3.7 [未完成](exercise-3.7/main.go)<br>

3.8 [未完成](exercise-3.8/main.go)<br>

3.9 [未完成](exercise-3.9/main.go)<br>

3.10 [完成](exercise-3.10/main.go)<br>
编写一个非递归的comma函数，运用bytes.Buffer，而不是简单的字符串拼接。<br>
Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.

3.11 [未完成](exercise-3.11/main.go)<br>
增强comma函数的功能，让其正确处理浮点数，以及带有可选正负号的数字。<br>
Enhance comma so that it deals correctly with floating-point numbers and an option alsign.

3.12 [未完成](exercise-3.12/main.go)<br>
编写一个函数判断两个字符串是否同文异构，也就是，它们都含有相同的字符但是排列顺序不同。<br>
Write a function that reports whether two strings are anagrams of each other,that is, they contain the same letters in a different order.

4.3 [完成](exercise-4.3/main.go)<br>
重写函数reverse，使用数组指针作为参数而不是slice。<br>
Rewrite reverse to use an array pointer instead of a slice.<br>
Example: [1,2,3] to [3,2,1]

4.4 [完成](exercise-4.4/main.go)<br>
编写一个函数rotate，实现一次遍历就可以完成元素旋转。<br>
Write a version of rotate that operates in a single pass.<br>
Example: [1,2,3,4,5] to [3,4,5,1,2]

4.5 [完成](exercise-4.5/main.go)<br>
编写一个就地处理函数，用于去除[]sting slice中相邻的重复字符串元素。<br>
Write an in-place function to eliminate adjacent duplicates in a []string slice.<br>
Example: []string{"a", "a", "b", "a", "c", "b", "b", "b"} to []string{"a", "a", "b", "a", "c", "b"} in-place funcition 不生成中间变量，类似a[i],a[j]=a[j],a[i]

4.6 [未完成](exercise-4.6/main.go)<br>
编写一个就地处理函数，用于将一个UTF-8编码的字节slice中的所有相邻Unicode空白字符（查看unicode.IsSpace）缩减为一个ASCII空白字符。<br>
Write an in-place function that squashes each run of adjacent Unicode spaces(see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.

4.7 [未完成](exercise-4.7/main.go)<br>
Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string , in place. Can you do it without allocating new memory?