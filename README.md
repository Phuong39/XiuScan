# XiuScan

## 介绍

一个纯Golang编写基于命令行的Java框架漏洞扫描工具

致力于参考xray打造一款高效方便的漏扫神器

计划支持Fastjson、Shiro、Struts2、Spring、WebLogic等框架

PS: 取名为XiuScan因为带我入安全的大哥是修君

## 特点

- 类似xray，直接提供一个可执行文件（例如exe）无需配置环境或下载依赖

- 使用ceye.io平台做无回显漏洞验证，提供api token即可自动化检测

- 将会用golang实现简化版ysoserial，不借助java生成payload

## Struts2

Struts2系列漏洞扫描，无害化命令检测，使用命令**struts2**

支持：S2-001 S2-005 S2-007 S2-008 S2-009 S2-012 
S2-013 S2-015 S2-016 S2-032 S2-045 S2-046 S2-048 
S2-052 S2-053 S2-057

```shell
./xiuscan struts2 -t http://127.0.0.1/ -o result.json
```

## TODO