# NETFLIX-VERIFY

最新版本: `v1.0-stable`

流媒体NetFlix解锁检测脚本，使用Go语言编写。

在VPS网络正常的情况下，哪怕是双栈网络也可在几秒内快速完成IPv4/IPv6的NF解锁情况判断。

## 其他常见流媒体脚本链接
~~DisneyPlus 解锁检测： https://github.com/sjlleo/VerifyDisneyPlus~~

~~Youtube 缓存节点、地域信息检测：https://github.com/sjlleo/TubeCheck~~

## 新特性

**2024/04/19**

添加`完整的IP顯示`，发布 `v1.0` 第一个稳定版

## 编译脚本

编译脚本来自于 [@utpinfo](https://github.com/utpinfo) 同学的修改版，在从特表感谢~

## 指定网卡出口测试

感谢 @caobug

有些时候，我们使用了`Warp`或者是其他隧道网卡工具的时候，设置的路由表默认不通过他们上网

在这种情况下，如果我们依旧想要测试该网卡出口是否支持解锁NetFlix，可以选择指定网卡IP进行测试

在终端输入`ip a`，查看您想测试的网卡IP，然后加入参数 `-address` + 您的网卡IP地址即可，具体可看下图

![图片](https://user-images.githubusercontent.com/13616352/169537511-8a10c9d2-3f4c-438a-b20a-8e5e31464259.png)


## 鸣谢

1. 感谢 [@CoiaPrant](https://github.com/CoiaPrant) 指出对于地域检测更简便的方法
2. 感谢 [@XmJwit](https://github.com/XmJwit) 解决了IPV6 Only VPS无法下载脚本的问题
3. 感谢 [@samleong123](https://github.com/samleong123) 指出了文档的解释缺陷
4. 感谢 [@ymcoming](https://github.com/ymcoming) 指出了IPv6的检测Bug

## 功能实现

- [X] 解锁情况判断
- [X] 地域信息显示
- [X] 双栈网络测试
- [X] 代理检测 (Experiment)

~~半解锁检测（Deprecated）~~

## 如何食用

使用前，如果您不知道您使用的是什么架构的CPU，请先使用`uname -m`查看

**如果提示`Exec format error`是因为您下载了与您系统架构不对应的二进制文件**

对于`amd64`（`x86_64`），请使用如下命令下载运行
```bash
wget -O nf https://github.com/utpinfo/GoGoLang/releases/download/v1.0.0/nf-utp && chmod +x nf && ./nf
```

## 相关名词解释

1. **不提供服务** - 所在的地区NF没开通，连自制剧也看不了
2. **宽松版权** - 有些NF拍摄的影片不是特别注重版权，所以限制放的很开
3. **解锁自制剧** - 代表可以看由NF自己拍摄的影片
4. **解锁非自制剧** - 代表可以看NF买下的第三方版权影片
5. **地域解锁** - NF在不同的地区可以看的片源都是不同的，有些影片只能在特定区观看

一般来说，需要能看非自制剧才算真正意义上的NF解锁
## Stargazers over time

[![Stargazers over time](https://starchart.cc/sjlleo/netflix-verify.svg)](https://starchart.cc/sjlleo/netflix-verify)

