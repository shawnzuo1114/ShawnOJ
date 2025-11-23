# 🧭 Online Judge Backend

一个用于在线评测系统的后端服务，实现用户管理、题目管理、代码提交与评测等核心能力。项目采用模块化设计，接口清晰、可扩展性强，可作为搭建完整 Online Judge 平台的基础工程。

------

## 📌 项目简介

本项目聚焦 Online Judge 系统的核心后端功能，包括用户体系、题目管理、提交记录以及后续可扩展的评测模块。项目接口风格遵循 RESTful 设计规范，适合用于个人学习、考核项目或完整 OJ 系统的后端架构基础。

统一 API 前缀：`/api/v1`

------

## 🗂 功能概览

以待办清单形式展示当前与预期功能的进度。

### ✅ 用户模块（User & Auth）

- [x] 用户注册  
- [x] 用户登录  
- [x] 获取个人信息  
- [x] 修改密码  
- [ ] 用户注销  
- [ ] 验证码服务  

---

### 📘 题目模块（Problem）
- [ ] 获取题目列表  
- [ ] 获取题目详情  
- [ ] 发布题目  
- [ ] 修改题目  
- [ ] 删除题目  

---

### ⚙️ 提交与评测模块（Submission & Judge）
- [ ] 提交代码  
- [ ] 存储提交记录  
- [ ] 异步评测调度  
- [ ] 获取评测结果  
- [ ] Docker 沙盒执行（推荐）  
- [ ] Go 语言评测  
- [ ] 多语言评测（扩展）  

---

### 🏆 比赛模块（Contest）
- [ ] 获取比赛列表  
- [ ] 报名参赛  
- [ ] 比赛题目访问  
- [ ] 比赛排行榜  

---

### 🔐 安全与增强项（Enhancements）
- [ ] 密码加盐加密  
- [ ] 多种登录方式（邮箱 / 短信 / 第三方）  
- [ ] 认证方式：JWT 或 Session  
- [ ] 前后端完整部署上线  
- [ ] HTTPS 支持  
- [ ] 数据缓存  
- [ ] 安全增强（XSS / SQL 注入 / CORS / CSRF）  
- [ ] 自研沙盒（cgroup / namespace / qemu）  

---

## 📡 API 概述

以下为项目接口的核心结构。

### 用户模块

```
POST    /api1/v1/users/register
POST    /api1/v1/users/login
GET     /api2/v1/users/me
PUT     /api2/v1/users/password
POST    /api2/v1/users/logout
GET     /api2/v1/users/captcha
```

### 题目模块

```
GET     /api/v1/problems
GET     /api/v1/problems/{id}
POST    /api/v1/problems
PUT     /api/v1/problems/{id}
DELETE  /api/v1/problems/{id}
```

### 提交与评测

```
POST    /api/v1/submissions
GET     /api/v1/submissions
GET     /api/v1/submissions/{id}
```

------

## 🏗 技术架构概述

- **后端框架**：可自由选择（Go / Node.js / Java / Python）
- **数据库**：MySQL 或 PostgreSQL
- **认证方式**：JWT 或 Session
- **评测执行**：基于 Docker 的隔离沙盒（推荐）
- **任务模型**：提交后入队，异步评测，结果写入数据库

------

## 🔄 评测流程示意

```mermaid
flowchart LR
    A[提交代码] --> B[记录 Submission]
    B --> C[加入评测队列]
    C --> D[评测 Worker 拉取任务]
    D --> E[沙盒运行代码]
    E --> F[写入评测结果]
    F --> G[用户查询]
```

------

## 📦 部署说明（可后续补充）

建议包含以下内容：

- 环境要求
- 后端运行方式
- 数据库初始化
- Docker 部署
- 配置文件说明

------

## 🧪 开发与测试

可逐步补充：

- 单元测试
- 评测模块调试方法
- API 调试（如 Postman、Swagger）

------

## 📝 未来扩展方向

- 多语言编译执行
- 分布式评测集群
- 比赛系统完善
- 前端管理界面与 OJ 提交页面
- 完整监控体系