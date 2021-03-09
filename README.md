# git-pics

## 开发环境
- go 1.16

## 前置安装
```
go get github.com/akavel/rsrc
go get github.com/lxn/win
go get github.com/lxn/walk
```

## 编译
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

## 文件说明

> `main.manifest` 是应用程序配置元数据的清单文件。
> `main.manifest` 保存在 `main.go` 同级目录下。