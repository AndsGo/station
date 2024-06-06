# 需求分析阶段

## 站群是什么

站群，即**一个人或一个团队操作多个网站，目的是通过搜索引擎获得大量流量，或者是将链接指向同一个网站，以提高搜索排名**。

## 站群包括哪些功能

- [x] 站点管理
- [x] 内容管理
- [ ] 数据统计和分析
- [ ] 用户管理
- [ ] 等其他功能

一个完善的站群系统还包含其他功能，我们这里先实现 站点管理、内容管理这两个核心功能。

## 需求整理

### 1.原型

https://98ughm.axshare.com/#id=817l34&p=%E7%AB%99%E7%BE%A4%E7%AE%A1%E7%90%86

### 2.模块

我们一期先简单实现站点管理和内容管理。实现一个内容推送的功能。大致功能模块如下：

```
|-内容管理
	|-增改查
	|-投放
	|-文章编辑
		|-媒体资源管理
		|-文章编辑器
|-站点管理
	|-增改查
	|-投放
|-日志
	|-查
```

### 3.资源

1.wordpress 站点 

​	5.6版本以下	需要安装https://github.com/WP-API/Basic-Auth  安装完成后使用账号密码 即可需要授权的 访问rest-api

​	5.6版本以上	5.6+已经支持授权只要在users 中配置 api key

2.wordpress api 文档 https://developer.wordpress.org/rest-api/reference/

## 系统设计

### 1.技术选型

- 后端 Go

  ​	go-wordpress 库 sdk

  ​    go-zore 库

- 前端 vue + elementUi

### 2.数据库设计

```
|-内容表
|-内容大字段表
|-站点表
|-站点内容关系表
|-日志表
```

### 3.wordpress api对接

文档地址:https://developer.wordpress.org/rest-api/reference/

```
|-Posts
|-Categories
|-Media
|-Users
```

### 4.项目模块设计

后端:

```
|-api	--网关
|-station rpc	--对外rpc 主要业务逻辑这里面
|-station job	--job  使用xxl-job实现
|-pkg	--公共模块、一般放工具
```

前端:

使用 vue-element-plus-admin

## 代码实现

[2.代码实现]("doc/2代码实现.md")

