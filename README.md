# 是什么

github520cli是一个使用GO编写,依赖[github520](https://github.com/521xueweihan/GitHub520)的命令行工具,用于修改hosts文件,解决github访问慢的问题。
- 支持MAC, Linux和Windows系统

# 为什么

手动修改hosts文件太麻烦了,而且存在出错的可能,所以写了这个工具。

# 怎么安装

## GO

需要先安装GO,安装方法请参考[GO官网](https://golang.org/dl/)

```shell
go install github.com/xuhe2/github520cli@latest
```

# 怎么用

## Linux和mac系统

```shell
sudo ~/go/bin/github520cli
```
> 注意: 在mac和linux系统上需要使用sudo命令,否则会报权限错误.

> 使用`go install`把github520cli安装到了`~/go/bin/`目录下,这个目录在`$PATH`中,为什么不能直接使用`github520cli`命令呢?因为root的`$PATH`和普通用户的`$PATH`不一样,所以需要使用绝对路径来运行。

> 或者先使用`su root`命令切换到root用户,然后再使用`github520cli`命令也是可以的。因为`su root`命令会切换到root用户,会保留原先的`$PATH`环境变量。

![use-case-linux](./docs/imgs/use-case-linux.png)

## Windows系统

- 在Windows系统上需要使用管理员权限运行

![use-case-windows](./docs/imgs/use-case-windows.png)