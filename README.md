#Mac上配置Go环境

## 1.查找当前最新版本
```
brew search go
```
![WeChat6480f0a52240f9e97e1a5368cceb8a26.png](https://upload-images.jianshu.io/upload_images/1745735-234dc0d47e9aa293.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 2.安装go
```
brew install go@1.9
```
## 3.配置环境变量
>3.1、打开文件
```
vim ~/.bashrc
```
```
#GOROOT
export GOROOT=/usr/local/opt/go\@1.9
#GOPATH
export GOPATH=$HOME/Documents/code/gopath
#GOPATH root bin
export PATH=\$PATH:$GOROOT/bin
```
![屏幕快照 2018-11-29 下午2.10.56.png](https://upload-images.jianshu.io/upload_images/1745735-8488a2a4db03f734.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

>3.2、让配置生效
```
source ~/.bashrc
```
## 4.验证是否成功
```
go env
```
![屏幕快照 2018-11-29 下午2.11.22.png](https://upload-images.jianshu.io/upload_images/1745735-99b55f6f041ce0c3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

