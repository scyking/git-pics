# git-pics

## 项目说明
> 基于git实现图床功能。

### 开发环境
- go 1.16

## 使用说明

### 前置安装
```
go get github.com/akavel/rsrc
go get github.com/lxn/win
go get github.com/lxn/walk
```

### 编译使用
1. 使用` rsrc tool`，生成`*.syso`文件
```
rsrc -manifest main.manifest -o main.syso
```
1. 编译
```
// 调试
go build
// 正式构建
go build -ldflags="-H windowsgui"
```
1. 运行
```
git-pics.exe
```