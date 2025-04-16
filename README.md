# Pulse 可视化告警通知引擎

**Pulse** 是一款面向企业级监控场景的 **智能化告警管理平台**，深度整合了 Prometheus、阿里云 SLS 等主流时序数据库，提供从告警规则配置、路由、分派到响应的全生命周期管理，助力企业实现监控告警的统一化、智能化和高效化。

---

## 核心功能

### 1. **多源时序数据集成**
- **支持对接**：  
  Prometheus、阿里云 SLS等主流时序数据库。
- **统一查询**：  
  通过标准化接口屏蔽底层存储差异，实现跨数据源的告警规则配置与数据拉取。
- **动态适配**：  
  自动解析不同时序库的指标格式，无需手动转换。

---

### 2. **统一告警规则管理**
- **双向同步**：  
  支持与 Prometheus 等系统的 **规则双向同步**，保持配置一致性。
- **可视化配置**：  
  通过 GUI 或 YAML 编辑器快速定义告警规则（如阈值、持续时间）。

---

### 3. **告警分析与总览**
- **全局告警仪表盘**：  
  实时展示告警数量、等级分布、触发趋势等关键指标。
- **历史回溯**：  
  支持按时间、服务、等级等多维度检索历史告警。

---

### 4. **多租户与数据隔离**
- **自定义工作区**：  
  按部门、项目或团队划分独立工作区，实现数据与告警的物理/逻辑隔离。
- **权限继承**：  
  支持工作区内的子资源（如告警规则、通知策略）权限继承管理。

---

### 5. **动态告警分派与升级**
- **分派策略引擎**：  
  - 按告警等级、服务类型、标签等条件动态路由至指定人员或组。  
  - 支持分派到企业微信、短信、电话、邮件等渠道。  
- **告警升级机制**：  
  若超时未确认，自动升级至更高优先级人员或值班组。
- **告警静默机制**：  
  支持按标签、服务等条件设置告警静默期，专注于重要的告警信息。
- **智能的路由机制**：  
支持按标签、服务等条件设置告警静默期，专注于重要的告警信息。
---

### 6. **智能值班排班**
- **排班日历**：  
  可视化配置值班组、轮班周期（按天/周/月）及交接规则。  
- **自动触达**：  
  值班期间自动将告警推送至当值人员，支持“一键找人”应急响应。  

---

## 开始使用

Pulse是一个前后端分离的应用，需要分别启动前端和后端服务。

### 1. 前端服务

进入前端frontend目录，安装依赖：
```bash
cd frontend
yarn install
```
初始账号：
```bash
# 用户名
admin
# 密码
admin
```

启动前端服务：   
```bash 
yarn dev
```

### 2. 后端服务

前置依赖服务：
- [mysql]
- [redis]
- [umdp](https://gitee.com/florenbai_1_floren_bai/umdp)


配置信息：  

```bash
# 配置文件路径
config/config.yaml
```

注意：
- 告警通知服务强依赖[umdp](https://gitee.com/florenbai_1_floren_bai/umdp)消息发送平台，如需接入自己的告警通知，请进行定制化开发。
- 如需要实时同步Prometheus的告警规则，请使用[agent](https://gitee.com/florenbai_1_floren_bai/pulse/tree/master/agent)，并开启remote_prometheus参数
- 如果开启remote_prometheus参数（自动同步Prometheus告警规则），需要确保agent和pulse的配置文件中prometheus的配置项一致。默认使用nacos作为服务发现。
- 告警通知功能使用[umdp](https://gitee.com/florenbai_1_floren_bai/umdp)，请配置umdp中的内容。系统已经集成了：

```bash
email_channel:邮件
wechat_channel:企业微信
phone_channel:电话
template:模板编号
```

3个通知渠道，并配置您在umdp的template即可.


启动后端服务：  
```bash
go run main.go
```

## 相关项目:
- [umdp](https://gitee.com/florenbai_1_floren_bai/umdp)：一个集成了企业微信、飞书、电话、邮件等多渠道的消息推送平台。
- [hertz](https://github.com/cloudwego/hertz)：字节的高性能微服务 HTTP 框架。
- [kitex](https://github.com/cloudwego/kitex)：字节的高性能Go 微服务 RPC 框架。

## 赞赏
如果您觉得这个项目对您有帮助，欢迎赞赏支持：

<img src="./assets/wechat.jpg" alt="赞赏码" width="405" height="550" />