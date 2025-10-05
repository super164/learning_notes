# GOPROXY的配置

Go1.13之后GOPROXY默认值为https://proxy.golang.org,国内无法直接访问,推荐使用goproxy.cn

```命令行
go env -w GOPROXY=https://goproxy.cn,direct
```





# go mod命令

常用指令如下

```
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```