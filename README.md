# XiuScan

## 不完善，正在开发中

## 介绍

一个纯Golang编写基于命令行的Java框架漏洞扫描工具

致力于参考xray打造一款高效方便的漏扫神器

计划支持Fastjson、Shiro、Struts2、Spring、WebLogic等框架

PS: 取名为XiuScan因为带我入安全的大哥是修君

## 特点

- 类似xray，直接提供一个可执行文件（例如exe）无需配置环境或下载依赖

- 使用ceye.io平台做无回显漏洞验证，提供api和token即可自动化检测

- 将会用golang实现简化版ysoserial，不借助java生成payload（已完成CC1）

## Shiro

Shiro一把梭检测，无害化命令，使用命令**shiro**

使用ceye.io：`./xiuscan shiro -ci xxxxxx.ceye.io -ct your_ceye_token`

当然如果不给出ci和ct参数，将会直接进行TomcatEcho检测

![](https://github.com/EmYiQing/XiuScan/blob/master/img/2.png)

## Struts2

Struts2系列漏洞扫描，无害化命令检测，使用命令**struts2**

查看已支持模块：`./xiuscan struts2 --list`

![](https://github.com/EmYiQing/XiuScan/blob/master/img/1.png)

目前还不能扫描，只完成了所有测试用例，预计下月完善扫描逻辑


