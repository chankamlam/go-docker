# go-docker



# Version
#### V0.9
+ 实现基础Namespace隔离
+ 使用Cobra分割指令
+ 创建container package

#### V0.10
+ 改造ContainerName产生次序
+ 增加容器文件系统目录为rootfs
+ 实现docker run --name=xx 指定容器名称
+ 实现docker run -d
+ 实现docker logs
+ 使用busybox作为base image
+ 完善错误提示
```
docker pull busybox
docker run -d busybox top -b
docker export -o busybox.tar xxxxxxxxx
mkdir base
tar -xvf busybox.tar -C base
```
