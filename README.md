# 项目名称

简要描述你的项目。解释项目的目标、功能和用途。

## 目录

- [简介](#简介)
- [功能](#功能)
- [安装](#安装)
- [使用](#使用)
- [贡献](#贡献)
- [许可证](#许可证)

## 简介

提供有关项目的详细信息。解释为什么这个项目是有用的，或者它解决了什么问题。

## 功能

- 列出项目的主要功能或特点。
- 例如：
  - 功能1
  - 功能2
  - 功能3

## 安装

提供如何安装和配置项目的步骤。包括必要的前提条件和依赖项。

### 先决条件

- 依赖项1
- 依赖项2

### 安装步骤

待完善




# dootask-workspace
## 接口使用说明
请求地址：http://127.0.0.1:5555
### 一、同步用户ID (GET)
 ```
http://127.0.0.1:5555/sync
 ```
### 二、设置创建工作区权限 (POST)
 ```
http://127.0.0.1:5555/set
 ```
 ```
headers
 "Content-Type: application/json" 
body
 {
  "user_id": 1,           // 用户ID
  "is_create": true       // true:允许创建工作区，false:不允许创建工作区
 }
 ```
### 三、创建工作区 (POST)
 ```
http://127.0.0.1:5555/create
 ```
 ```
headers
 "Content-Type: application/json" 
body
 {
  "user_id": 1           // 用户ID
 }
 ```
 ### 四、删除工作区 (DELETE)
 ```
http://127.0.0.1:5555/delete-ws
 ```
 ```
headers
 "Content-Type: application/json" 
body
 {
  "user_id": 1           // 用户ID
 }
 ```
### 五、检查已创建的工作区数量 (GET)
 ```
http://127.0.0.1:5555/check
 ```
### 六、新建对话窗口 (POST)
```
http://127.0.0.1:5555/new
```
```
headers
 "Content-Type: application/json" 
body
{
  "slug": "workspace-for-user-1",
  "model": "ChatGPT",
  "avatar": "sk123"
}
```
### 七、查询已获授权用户 (POST)
```
http://127.0.0.1:5555/get-user
```
```
headers
"Content-Type: application/json"
body
{
  "user_id": 1
}
```
### 八、更新最后一条对话 (POST)
```
http://127.0.0.1:5555/update-last
```
```
headers
"Content-Type: application/json"
body
{
  "workspaceSlug": "workspace-for-user-1",              // 工作区 Slug
  "threadSlug": "d4c12455-92cc-442b-b701-58c4972dfcd0"  // 对话 Slug
}
```
### 九、获取对话列表 (POST)
```
http://127.0.0.1:5555/get-list
```
```
headers
"Content-Type: application/json"
body
{
  "user_id": 1
}
```