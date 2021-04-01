
## 应用概述
> 基于使用**github图床**，开发的辅助工具。

#### 开发背景

> 在编写Markdown文档时，常常会使用图床存放Image文件。
> 综合考虑，选用**github**。
> 在使用过程中，添加图片、提交、复制链接，重复性工作比较多。
> 再者，刚刚学习`go`，为了解其语法特性，写点简单的东西来练练手。

#### 基本功能

- 图片展示。点击图片，自动生成对应格式链接地址到粘贴板
- 图片上传。支持拖动自动上传

#### [版本说明](VERSION.md)

## 项目构建
> 其他环境下，不保证是否能够正常构建。

#### 开发环境
- win 10
- go 1.16

#### 前置安装
> 因网络原因，`go get`可能会失败。
> 将对应源码，拷贝到`GOPATH`中即可。（[参考](https://github.com/scyking/my-notes/blob/master/%E5%BC%80%E5%8F%91%E8%AF%AD%E8%A8%80/go/go%20get%E5%A4%B1%E8%B4%A5%E9%97%AE%E9%A2%98.md)）

```
go get github.com/akavel/rsrc
go get github.com/lxn/win
go get github.com/lxn/walk
```

#### 构建过程
> 可查看[walk](https://github.com/lxn/walk/README.md)具体说明

1. 使用`rsrc tool`，生成`*.syso`文件（不需重复生成）
```
rsrc -manifest main.manifest -o main.syso
```
1. 构建
```
// 调试(运行会有调试窗口)
go build
// 正式构建
go build -ldflags="-H windowsgui"
```
1. 运行生成的`.exe`文件