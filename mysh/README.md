# MyShell

shell
- 内部命令
- 外部命令


Linux系统编程 李慧琴 55讲

## 基本目标
用go实现一个shell


shell脚本基本逻辑

```c
for (;;) {
    print_prompt(prompt_str);
    handle_input(arg_buf, arg_list);
    run_cmd(arg_list);
}
```

需要注意的点
- shell命令分为两种: 内建命令(built-in)和外部命令
内建命令为类似cd, exit之类的命令
外部命令为安装在PATH中的二进制命令


## 区分内建命令和外部命令
type和enable
```
enable -a #查看内部命令
```

## C语言版本

fork(2) + exec(3) + getty(1)

基本Shell: https://panqiincs.me/2017/02/26/write-a-shell-basic-functionality/
Shell重定向: https://panqiincs.me/2017/04/19/write-a-shell-redirect-and-pipeline/

c语言的执行逻辑:

- 处理用户输入
getline获取用户输入, strsep/strtok拆分输入字符串为多个命令, 对拆分后的字符串需要去除多余的空格
- 运行命令
父进程会fork一个子进程，在子进程中使用exec函数运行用户输入的命令
execvp运行命令, 好处是不用输入完整的命令路径, 会自动在PATH中寻找对应命令 (execv则必须传入完整路径)
父进程wait等待子进程退出


glob函数匹配命令

## go语言版本

go执行内部命令

```go
if runtime.GOOS != "windows" {
    cmd = exec.Command("bash", "-c", c)
} else {
    cmd = exec.Command("cmd.exe", "/c", c)
}
```

fmt.Fprint函数族因为是往文件输出, 所以可以用来重定向
```go
fmt.Fprintln(os.Stderr, err)
```


go shell教程 https://studygolang.com/articles/13884

go os/exec使用教程 https://colobu.com/2017/06/19/advanced-command-execution-in-Go-with-os-exec/


docker shell https://github.com/moby/moby/blob/master/pkg/reexec/reexec.go
docker可能要考虑到容器内的需求
go实现传统的fork https://jiajunhuang.com/articles/2018_03_08-golang_fork.md.html

os包用法汇总 https://cloud.tencent.com/developer/article/1342799