# 功能说明

> 出现未知问题，可使用git工具自行维护。

## 准备工作
> 安装`git`，`clone`对应图床项目到指定文件夹

## 主要功能

#### 主界面说明

<img src="https://github.com/scyking/my-pics/blob/mastergpics/2021471036659.png" width="50%">

#### 添加工作空间
> 工作空间需要是图床项目根目录（即包含`.git`文件夹目录），否则无法正常使用。

- 初次启动（或者配置文件更改），通过弹出对话框添加

    <img src="https://github.com/scyking/my-pics/blob/mastergpics/20214710212181.png" width="50%">
    
- 通过菜单栏配置项对话框添加

    <img src="https://github.com/scyking/my-pics/blob/mastergpics/2021471024087.png" width="50%">
    
#### 添加图片

- 可配置项
    1. 快捷上传：将图片上传到指定文件夹中
    1. 非快捷上传：点击`TreeView`中文件夹，将图片上传至该文件夹

- 上传方式
    1. 通过菜单栏添加图片项上传
    1. 通过将图片拖动到主界面上传
    1. 通过菜单栏截图项上传（待实现）
    
#### 复制地址

1. 选择图片展示区中图片
1. 选择复制项中`Radio Button`，会将对应地址复制到粘贴板

- 复制项说明
    1. Markdown：通过标准`md`语法添加图片文本格式
    1. HTML：通过`<img>`标签添加图片文本格式
    1. URL：资源`github`远程链接地址
    1. FilePath：文件资源绝对路径
    
#### 添加文件夹
> 右键`TreeView`中文件夹，添加对应子文件夹

#### 远程提交
> 点击菜单栏远程提交项，后台会执行GIT `pull`、`push`命令

## 配置项说明
> 配置文件默认地址，`%UserDir%\AppData\Roaming\scyking\GPics\settings.ini`。

#### Git信息
> 目前地址通过配置项拼接而成，所以需要正确填写。 

- Repository：仓库名称，形式如`scyking/gpics`
- Server：server，形式如`github.com`
- UserName：用戶名，暂未使用
- Password：密碼，暂未使用
- Token：token，暂未使用

#### 快捷上传

- 开启状态：是否开启快捷上传。开启则直接将图片上传到指定文件夹中。
- 上传目录：需选择图床项目中文件夹，否则上传图片无法使用。

#### 其他配置

- 自动提交：开启则在上传图片后，自动执行远程提交命令。（建议关闭）
- 超时时间：执行远程提交命令的超时时间。（执行远程提交命令会阻塞程序，所以设置了等待超时。超时后，命令仍在后台执行，可能会执行成功）
- 工作空间：需选择图床项目根目录
