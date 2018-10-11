# selpg

### 任务详情
使用golang开发cli命令行工具selpg，具体如下：[开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)

参考如下：

[Golang之使用Flag和Pflag](https://o-my-chenjian.com/2017/09/20/Using-Flag-And-Pflag-With-Golang/)

[Go学习笔记：flag库的使用](https://studygolang.com/articles/5608)

[作业详情](https://pmlpml.github.io/ServiceComputingOnCloud/ex-cli-basic)

### Usage:
```
	.\selpg -s=num -e=num [options] [filename]

	-s=num          Start of Page <pageLength>.
	-e=num          End of Page <pageLength>.
	-l=num          [options]Specify the number of line per page.Default is 72.
	-d=lp number    [options]Using cat to test.\n")
	-f              [options]Specify that the pages are sperated by '\f'.
	[filename][options]  Read input from the file.
```


### 测试

1.`./selpg -s=0 -e=1 test.txt`
结果：

```
[lostking@localhost selpg]$ ./selpg -s=0 -e=1 test.txt
Line-1
Line-2
Line-3
Line-4
Line-5
Line-6
Line-7
···
···
···
Line-72
```

2.测试错误反馈

`./selpg -s=2 -e=1 -l=5 test.txt`

结果：

```
[lostking@localhost selpg]$ ./selpg -s=3 -e=2 test.txt

 Error: 
 ./selpg: Invalid arguments

Usage:
······
······
```

3. -l测试
`./selpg -s=0 -e=1 -l=4 test.txt`

结果

```
[lostking@localhost selpg]$ ./selpg -s=0 -e=1 -l=4 test.txt
Line-1
Line-2
Line-3
Line-4
```
4. -f测试
先创建带有'\f'换行符的测试文件

```
echo -e aaaaaa'\n'bbbbbb'\f'cccccc'\n'dddddd'\f'eeeeee'\n'ffffff'\f'gggggg'\n'hhhhhh >delimTest.txt
```
然后测试：

`./selpg -s=0 -e=4 -f delimTest.txt`

结果：

```
[lostking@localhost selpg]$ ./selpg -s=0 -e=4 -f delimTest.txt 
aaaaaa
bbbbbb

cccccc
dddddd

eeeeee
ffffff

```

5. -d测试
`./selpg -s=0 -e=3 -l=5 -d=lp1 test.txt`

结果：

```
[lostking@localhost selpg]$ ./selpg -s=0 -e=1 -l=4 -d=lp1 test.txt
     1	Line-1
     2	Line-2
     3	Line-3
     4	Line-4
```

6.测试输出到文件
`./selpg -s=0 -e=1 -l=4 test.txt >out.txt`

结果：

```
[lostking@localhost selpg]$ cat out.txt 
Line-1
Line-2
Line-3
Line-4
```

7.测试报错输出到文件
`./selpg -s=1 2>error.txt`

结果：

```
[lostking@localhost selpg]$ cat error.txt 

 Error: 
 ./selpg: arguments are not enough

```




