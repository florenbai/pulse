/*
SQLyog Community v13.1.9 (64 bit)
MySQL - 8.0.40 : Database - pulse
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`pulse` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `pulse`;

/*Table structure for table `agents` */

DROP TABLE IF EXISTS `agents`;

CREATE TABLE `agents` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_ip` varchar(100) DEFAULT NULL COMMENT 'agent地址',
  `heartbeat_time` datetime DEFAULT NULL COMMENT '上次心跳时间',
  `error_message` text COMMENT '异常信息',
  `status` tinyint(1) DEFAULT '1' COMMENT 'agent状态',
  `rule_source_id` int DEFAULT NULL COMMENT '规则数据源编号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `client_ip` (`client_ip`,`rule_source_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1538 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `agents` */

insert  into `agents`(`id`,`client_ip`,`heartbeat_time`,`error_message`,`status`,`rule_source_id`) values 
(1529,'172.24.0.1','2024-12-26 15:54:37','',1,5),
(1530,'169.254.27.126','2024-12-26 09:41:41','',1,5);

/*Table structure for table `alert_aggregation` */

DROP TABLE IF EXISTS `alert_aggregation`;

CREATE TABLE `alert_aggregation` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `aggregation_name` varchar(100) DEFAULT NULL COMMENT '聚合名称',
  `aggregation_desc` text COMMENT '聚合描述',
  `workspace_id` int NOT NULL COMMENT '工作区编号',
  `level_dimension` tinyint(1) DEFAULT '0' COMMENT '等级维度',
  `tags_dimension` json DEFAULT NULL COMMENT '标签维度',
  `title_dimension` tinyint(1) DEFAULT '0' COMMENT '标题维度',
  `windows` int DEFAULT '0' COMMENT '聚合窗口',
  `storm` int DEFAULT '0' COMMENT '风暴预警',
  `uid` int DEFAULT NULL COMMENT '用户编号',
  `status` enum('enabled','disabled','deleted') DEFAULT 'enabled' COMMENT '状态',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_aggregation` */

insert  into `alert_aggregation`(`id`,`aggregation_name`,`aggregation_desc`,`workspace_id`,`level_dimension`,`tags_dimension`,`title_dimension`,`windows`,`storm`,`uid`,`status`,`created_at`,`updated_at`) values 
(8,'dasdasd','',7,1,'[]',1,100,0,33,'deleted','2024-08-13 15:03:47','2024-08-13 15:07:19'),
(9,'ddddd','',7,1,'[{\"values\": [{\"tag\": \"region\", \"value\": [\"aliyun\"], \"operation\": \"NOT IN\"}]}]',1,20,0,33,'deleted','2024-08-13 15:06:48','2024-08-15 10:56:52');

/*Table structure for table `alert_channel` */

DROP TABLE IF EXISTS `alert_channel`;

CREATE TABLE `alert_channel` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `channel_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `channel_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '类型',
  `channel_sign` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标识',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `channel_group` varchar(100) DEFAULT NULL COMMENT '分组',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_channel` */

insert  into `alert_channel`(`id`,`channel_name`,`channel_type`,`channel_sign`,`created_at`,`updated_at`,`channel_group`) values 
(1,'企业微信','企业微信','wechat-darkhawk','2024-08-22 10:37:06','2024-08-22 16:47:34','统一消息平台'),
(2,'邮箱','邮箱','email-devops','2024-08-22 11:14:07','2024-08-22 16:47:45','统一消息平台'),
(3,'电话','电话','ymrt','2024-08-22 11:18:42','2024-08-22 16:47:53','统一消息平台');

/*Table structure for table `alert_event` */

DROP TABLE IF EXISTS `alert_event`;

CREATE TABLE `alert_event` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `alert_title` varchar(200) DEFAULT NULL COMMENT '告警标题',
  `source_id` int NOT NULL COMMENT '数据源编号',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '描述',
  `level` int NOT NULL COMMENT '告警等级',
  `first_trigger_time` datetime DEFAULT NULL COMMENT '首次告警时间',
  `first_ack_time` datetime DEFAULT NULL COMMENT '首次确认时间',
  `trigger_time` datetime DEFAULT NULL COMMENT '最新告警时间',
  `recover_time` datetime DEFAULT NULL COMMENT '恢复时间',
  `annotations` json DEFAULT NULL COMMENT '附加信息，json格式',
  `is_recovered` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否恢复',
  `notify_cur_number` int DEFAULT NULL COMMENT '告警发送数量',
  `progress` tinyint unsigned DEFAULT NULL COMMENT '处理进度 1 未认领 2 已认领 3 已关闭',
  `uid` int DEFAULT NULL COMMENT '最后处理人',
  `tags` json DEFAULT NULL COMMENT '告警标签，json格式',
  `workspace_id` int DEFAULT NULL COMMENT '所属工作区',
  `finger_print` varchar(100) DEFAULT NULL COMMENT '指纹标识',
  `source_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '来源IP',
  `integration_id` int DEFAULT NULL COMMENT '所属集成',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `workspace_id` (`workspace_id`),
  KEY `finger_print` (`finger_print`)
) ENGINE=InnoDB AUTO_INCREMENT=130 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_event` */

/*Table structure for table `alert_level` */

DROP TABLE IF EXISTS `alert_level`;

CREATE TABLE `alert_level` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `level_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `color` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '颜色',
  `is_default` tinyint(1) DEFAULT '1' COMMENT '是否系统默认',
  `level_desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '描述',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_level` */

insert  into `alert_level`(`id`,`level_name`,`color`,`is_default`,`level_desc`,`created_at`,`updated_at`) values 
(1,'A0','#f53f3f',1,'aaa',NULL,'2024-08-26 17:34:25'),
(2,'A1','#f77234',1,'aaa',NULL,'2024-08-26 17:34:29'),
(3,'A2','#ff7d00',1,NULL,NULL,NULL),
(4,'A3','#f7ba1e',1,NULL,NULL,NULL),
(5,'A4','#9fdb1d',1,NULL,NULL,NULL),
(6,'A5','#14c9c9',1,NULL,NULL,NULL);

/*Table structure for table `alert_log` */

DROP TABLE IF EXISTS `alert_log`;

CREATE TABLE `alert_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `alert_id` int NOT NULL COMMENT '告警事件编号',
  `action` enum('claim','closed','cancel_claim','opened') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '操作 认领、打开、关闭、取消认领',
  `uid` int DEFAULT NULL COMMENT '操作人',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_log` */

/*Table structure for table `alert_restrain` */

DROP TABLE IF EXISTS `alert_restrain`;

CREATE TABLE `alert_restrain` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `restrain_type` varchar(20) DEFAULT NULL COMMENT '抑制类型',
  `fields` json DEFAULT NULL COMMENT '匹配字段',
  `cumulative_time` int DEFAULT NULL COMMENT '抑制时长',
  `uid` int DEFAULT NULL COMMENT '编辑用户',
  `workspace_id` int DEFAULT NULL COMMENT '所属工作区',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_restrain` */

/*Table structure for table `alert_rule` */

DROP TABLE IF EXISTS `alert_rule`;

CREATE TABLE `alert_rule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(500) DEFAULT NULL COMMENT '告警标题',
  `health` varchar(100) DEFAULT NULL COMMENT '健康状态',
  `labels` json DEFAULT NULL COMMENT '标签',
  `annotations` json DEFAULT NULL COMMENT '注解',
  `duration` int DEFAULT NULL COMMENT '异常持续时间',
  `query` text COMMENT '查询语句',
  `gid` int DEFAULT NULL COMMENT '组编号',
  `source_id` int DEFAULT NULL COMMENT '数据源编号',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `gid_name` (`name`,`gid`)
) ENGINE=InnoDB AUTO_INCREMENT=56161 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_rule` */

/*Table structure for table `alert_rule_group` */

DROP TABLE IF EXISTS `alert_rule_group`;

CREATE TABLE `alert_rule_group` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `rule_source` int NOT NULL COMMENT '源编号',
  `group_name` varchar(100) DEFAULT NULL COMMENT '组名称',
  `file` varchar(500) DEFAULT NULL COMMENT '文件地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=266 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_rule_group` */

/*Table structure for table `alert_rule_source` */

DROP TABLE IF EXISTS `alert_rule_source`;

CREATE TABLE `alert_rule_source` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `source_name` varchar(100) NOT NULL COMMENT '数据源名称',
  `source_type` varchar(100) NOT NULL COMMENT '数据源类型',
  `address` varchar(500) NOT NULL COMMENT '连接地址',
  `mark` varchar(500) DEFAULT NULL COMMENT '备注',
  `auto_sync` tinyint(1) DEFAULT '1' COMMENT '自动同步',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `sign` varchar(100) DEFAULT NULL COMMENT '数据源标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_rule_source` */

/*Table structure for table `alert_silence` */

DROP TABLE IF EXISTS `alert_silence`;

CREATE TABLE `alert_silence` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `workspace_id` int NOT NULL COMMENT '工作区编号',
  `silence_name` varchar(100) DEFAULT NULL COMMENT '告警静默名称',
  `silence_desc` text COMMENT '告警静默描述',
  `silence_type` enum('once','period') DEFAULT NULL COMMENT '静默时间类型',
  `silence_time` json DEFAULT NULL COMMENT '静默时间',
  `filters` json DEFAULT NULL COMMENT '静默条件',
  `uid` int DEFAULT NULL COMMENT '设置用户',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_silence` */

insert  into `alert_silence`(`id`,`workspace_id`,`silence_name`,`silence_desc`,`silence_type`,`silence_time`,`filters`,`uid`,`created_at`,`updated_at`) values 
(1,8,'dasdasd','dsadasd','period','{\"times\": [\"02:02:04\", \"15:07:08\"], \"weeks\": [1, 3], \"rangeTime\": []}','[{\"tag\": \"aaaa\", \"value\": [\"dsadasd\"], \"operation\": \"IN\"}]',0,'2024-05-27 10:12:50','2024-05-27 13:33:23');

/*Table structure for table `alert_source` */

DROP TABLE IF EXISTS `alert_source`;

CREATE TABLE `alert_source` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `source_name` varchar(100) DEFAULT NULL COMMENT '告警源名称',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标',
  `category` varchar(100) DEFAULT NULL COMMENT '分类',
  `status` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `default_level_field` varchar(100) DEFAULT NULL COMMENT '默认告警等级字段',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_source` */

insert  into `alert_source`(`id`,`source_name`,`icon`,`category`,`status`,`default_level_field`,`created_at`,`updated_at`) values 
(1,'Prometheus','icon-prometheus','事件告警',1,'labels.severity','2024-01-02 09:42:58','2024-04-17 11:32:33'),
(4,'阿里云 SLS','icon-aliyun-sls','事件告警',1,'severity','2024-01-02 09:46:19','2024-04-17 11:32:53'),
(6,'自定义事件','icon-custom','事件告警',1,NULL,'2025-04-15 15:17:57','2025-04-15 15:18:00'),
(7,'阿里云监控 事件','icon-aliyun','事件告警',1,'level','2024-12-27 15:05:03','2024-12-27 15:06:43');

/*Table structure for table `alert_strategy` */

DROP TABLE IF EXISTS `alert_strategy`;

CREATE TABLE `alert_strategy` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `strategy_name` varchar(100) DEFAULT NULL COMMENT '分派策略名称',
  `strategy_desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '策略描述',
  `workspace_id` int DEFAULT NULL COMMENT '工作区编号',
  `template_id` int DEFAULT NULL COMMENT '模板编号',
  `status` enum('enabled','disabled','deleted') DEFAULT NULL COMMENT '状态',
  `system_strategy` tinyint(1) DEFAULT '0' COMMENT '是否执行全局策略',
  `continuous` tinyint(1) DEFAULT '0' COMMENT '是否接续匹配',
  `delay` int DEFAULT NULL COMMENT '延迟通知',
  `weight` int DEFAULT '0' COMMENT '权重',
  `time_slot` json DEFAULT NULL COMMENT '通知时间段策略',
  `filters` json DEFAULT NULL COMMENT '标签匹配策略',
  `strategy_set` json DEFAULT NULL COMMENT '策略详情',
  `uid` int DEFAULT NULL COMMENT '修改人编号',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '变更时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_strategy` */

insert  into `alert_strategy`(`id`,`strategy_name`,`strategy_desc`,`workspace_id`,`template_id`,`status`,`system_strategy`,`continuous`,`delay`,`weight`,`time_slot`,`filters`,`strategy_set`,`uid`,`created_at`,`updated_at`) values 
(4,'策略测试','',7,1,'deleted',0,0,0,1,'{\"type\": null, \"times\": null, \"weeks\": null, \"enable\": false, \"calendar\": null}','[{\"tag\": \"\", \"value\": [], \"operation\": \"IN\"}]','[{\"teams\": [], \"members\": [33], \"alertLoop\": {\"count\": 1, \"enable\": false, \"interval\": 0}, \"schedules\": [], \"alertLevel\": [\"A0\", \"A1\", \"A2\", \"A3\", \"A4\", \"A5\"], \"progression\": {\"enable\": false, \"duration\": 1, \"condition\": 0}, \"alertChannels\": [1, 2, 3]}]',33,'2024-05-13 15:26:53','2024-08-26 16:01:58'),
(5,'大撒大撒','',7,1,'deleted',0,0,0,1,'{\"type\": null, \"times\": null, \"weeks\": null, \"enable\": false, \"calendar\": null}','[]','[{\"teams\": [], \"members\": [2], \"alertLoop\": {\"count\": 1, \"enable\": false, \"interval\": 0}, \"schedules\": [], \"progression\": {\"enable\": false, \"duration\": 1, \"condition\": 0}, \"alertChannels\": [[1], [2], [3], [4], [2], [1]]}]',33,'2024-05-16 15:49:50','2024-08-26 16:01:48'),
(6,'asdasdasd','',7,5,'deleted',0,0,0,2,'{\"type\": null, \"times\": null, \"weeks\": null, \"enable\": false, \"calendar\": null}','[{\"values\": null}, {\"values\": null}]','[{\"teams\": [], \"members\": [33], \"alertLoop\": {\"count\": 1, \"enable\": false, \"interval\": 0}, \"schedules\": [], \"progression\": {\"enable\": false, \"duration\": 1, \"condition\": 0}, \"alertChannels\": [[1], [2], [2, 1], [3], [3], [3]]}]',0,'2024-05-21 10:51:32','2024-08-26 16:02:21'),
(7,'全局测试策略','sdasdasd',7,5,'enabled',0,0,0,1,'{\"type\": null, \"times\": null, \"weeks\": null, \"enable\": false, \"calendar\": null}','[]','[{\"teams\": [], \"members\": [33], \"alertLoop\": {\"count\": 1, \"enable\": false, \"interval\": 0}, \"schedules\": [], \"progression\": {\"enable\": false, \"duration\": 1, \"condition\": 0}, \"alertChannels\": [[1], [2], [2], [2], [1], [2]]}]',33,'2024-06-20 14:51:20','2024-08-28 10:07:01'),
(8,'策略2','',7,1,'disabled',0,0,0,2,'{\"type\": null, \"times\": null, \"weeks\": [2], \"enable\": true, \"calendar\": null}','[{\"values\": [{\"tag\": \"env\", \"value\": [\"devops\"], \"operation\": \"IN\"}, {\"tag\": \"aaa\", \"value\": [\"aaaa\"], \"operation\": \"IN\"}]}, {\"values\": [{\"tag\": \"cluster\", \"value\": [\"devops-nginx\"], \"operation\": \"IN\"}]}]','[{\"teams\": [], \"members\": [39], \"alertLoop\": {\"count\": 1, \"enable\": false, \"interval\": 0}, \"schedules\": [], \"progression\": {\"enable\": false, \"duration\": 1, \"condition\": 0}, \"alertChannels\": [[1], [1], [1], [1], [1], [1]]}]',33,'2024-07-16 11:50:24','2024-08-26 16:03:10'),
(9,'aaa','',8,1,'enabled',0,0,0,1,'{\"type\": null, \"times\": null, \"weeks\": null, \"enable\": false, \"calendar\": null}','[]','[{\"teams\": [], \"members\": [1], \"alertLoop\": {\"count\": 1, \"enable\": false, \"interval\": 0}, \"schedules\": [], \"progression\": {\"enable\": false, \"duration\": 1, \"condition\": 0}, \"alertChannels\": [[1], [1], [1], [1], [1], [1]]}]',1,'2024-12-27 14:09:55','2025-04-15 15:39:04');

/*Table structure for table `alert_template` */

DROP TABLE IF EXISTS `alert_template`;

CREATE TABLE `alert_template` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `template_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '告警模板名称',
  `template_desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '告警模板描述',
  `channels` json DEFAULT NULL COMMENT '绑定的通知渠道',
  `enable` tinyint(1) DEFAULT '1' COMMENT '是否开启',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_template` */

insert  into `alert_template`(`id`,`template_name`,`template_desc`,`channels`,`enable`,`created_at`,`updated_at`) values 
(1,'默认模板','如果没有模板，将默认使用本模板，','[{\"name\": \"企业微信\", \"content\": \"\", \"finished\": false}, {\"name\": \"邮件\", \"content\": \"\", \"finished\": false}]',0,'2024-01-08 10:56:46','2024-07-01 15:25:13'),
(5,'测试模板1','阿萨啊啊啊','[{\"name\": \"企业微信\", \"content\": \"您有一个新的告警待处理\\n**事项详情**\\n告警标题：{{ AlertTitle }}\\n告 警 源：{{ SourceName }}\\n告警等级：`{{ LevelName }}`\\n告警描述：{{ Description }}\\n告警标签：\\n{% for k,v in Tags %}  \\n`{{ k}} : {{ v}}`\\n {% endfor %}\\n告警时间：{{ TriggerTime }}\", \"finished\": true}, {\"name\": \"邮箱\", \"content\": \"\", \"finished\": false}]',1,'2024-06-14 10:36:43','2024-08-26 15:37:04');

/*Table structure for table `alert_workspace` */

DROP TABLE IF EXISTS `alert_workspace`;

CREATE TABLE `alert_workspace` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `workspace_id` int NOT NULL,
  `alert_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `alert_workspace` */

insert  into `alert_workspace`(`id`,`workspace_id`,`alert_id`) values 
(5,7,91),
(6,7,92),
(7,8,92),
(8,7,93),
(9,7,94),
(10,7,95),
(11,7,96),
(12,7,97),
(13,8,97),
(14,8,98),
(15,8,99),
(16,7,100),
(17,8,100),
(18,8,102),
(19,8,103),
(20,7,104),
(21,7,105),
(22,7,106),
(23,7,107),
(24,7,108),
(25,8,108),
(26,7,109),
(27,8,109),
(28,7,110),
(29,8,110),
(30,8,111),
(31,8,112),
(32,8,113),
(33,7,115),
(34,7,116),
(35,7,117),
(36,7,118),
(37,7,119),
(38,7,124),
(39,7,125),
(40,7,126),
(41,7,127),
(42,7,128),
(43,7,129);

/*Table structure for table `audit_log` */

DROP TABLE IF EXISTS `audit_log`;

CREATE TABLE `audit_log` (
  `user_id` bigint DEFAULT NULL COMMENT '关联用户',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '显示名',
  `user` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名',
  `action` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '具体方法描述',
  `business` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '业务',
  `method` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '请求方法',
  `handler` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '处理程序/接口',
  `user_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户源IP',
  `http_code` bigint DEFAULT NULL COMMENT '返回状态',
  `body` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '请求Body',
  `resp_body` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '响应Body',
  `dom` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'default' COMMENT '所属域',
  `obj_id` bigint DEFAULT NULL COMMENT '关联变更对象ID',
  `cost_time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '耗时',
  `status` bigint DEFAULT NULL COMMENT '状态[0: error, 1: success, 2: forbidden, 3: exception]',
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=346 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `audit_log` */

/*Table structure for table `channel_template` */

DROP TABLE IF EXISTS `channel_template`;

CREATE TABLE `channel_template` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `template_id` int DEFAULT NULL,
  `channel_id` int DEFAULT NULL,
  `content` text,
  `finished` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `channel_template` */

/*Table structure for table `integration_router` */

DROP TABLE IF EXISTS `integration_router`;

CREATE TABLE `integration_router` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `integration_id` int DEFAULT NULL COMMENT '集成编号',
  `filters` json DEFAULT NULL COMMENT '匹配条件',
  `next` tinyint(1) DEFAULT '0' COMMENT '是否继续匹配',
  `uid` int DEFAULT NULL COMMENT '编辑的用户',
  `sort` int DEFAULT NULL COMMENT '排序',
  `workspaces` json DEFAULT NULL COMMENT '路由工作区',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `integration_router` */

/*Table structure for table `level_integration_mapping` */

DROP TABLE IF EXISTS `level_integration_mapping`;

CREATE TABLE `level_integration_mapping` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `level_id` int NOT NULL,
  `integration_id` int NOT NULL,
  `alert_level` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `level_integration_mapping` */

/*Table structure for table `level_source_mapping` */

DROP TABLE IF EXISTS `level_source_mapping`;

CREATE TABLE `level_source_mapping` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `level_id` int DEFAULT NULL COMMENT '系统告警等级编号',
  `ws_id` int DEFAULT NULL COMMENT '告警源编号编号',
  `alert_source_level` varchar(50) DEFAULT NULL COMMENT '告警源等级',
  `workspace_id` int DEFAULT NULL COMMENT '工作区编号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `level_source_mapping` */

/*Table structure for table `menu` */

DROP TABLE IF EXISTS `menu`;

CREATE TABLE `menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `name` varchar(158) NOT NULL COMMENT '名称',
  `desc` text COMMENT '描述',
  `type` bigint DEFAULT '1' COMMENT '菜单类型[0:目录,1:菜单,2:按钮]',
  `pid` bigint DEFAULT '0' COMMENT '父级',
  `sort` bigint DEFAULT '100' COMMENT '排序',
  `icon` varchar(158) DEFAULT NULL COMMENT '模块',
  `route_path` varchar(255) DEFAULT NULL COMMENT '路由路径',
  `route_name` varchar(158) DEFAULT NULL COMMENT '路由名',
  `redirect` varchar(255) DEFAULT NULL COMMENT '重定向地址',
  `permission` varchar(255) DEFAULT NULL COMMENT '权限标识',
  `requires_auth` tinyint(1) DEFAULT '1' COMMENT '需要认证',
  `locale` varchar(255) DEFAULT NULL COMMENT '多语言标识',
  `hide_menu` tinyint(1) DEFAULT '0' COMMENT '隐藏菜单',
  `hide_child_menu` tinyint(1) DEFAULT '0' COMMENT '隐藏子菜单',
  `is_link` tinyint(1) DEFAULT '0' COMMENT '外部链接',
  `cache` tinyint(1) DEFAULT '0' COMMENT '启用缓存',
  `affix` tinyint(1) DEFAULT '0' COMMENT '固定标签',
  `enabled` tinyint(1) DEFAULT '1' COMMENT '启用',
  `active_menu` tinyint(1) DEFAULT '1' COMMENT '高亮菜单',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `menu` */

insert  into `menu`(`id`,`created_at`,`updated_at`,`name`,`desc`,`type`,`pid`,`sort`,`icon`,`route_path`,`route_name`,`redirect`,`permission`,`requires_auth`,`locale`,`hide_menu`,`hide_child_menu`,`is_link`,`cache`,`affix`,`enabled`,`active_menu`) values 
(1,'2023-12-27 21:03:06','2025-03-11 15:00:37','仪表盘','',1,0,1,'icon-dashboard','/dashboard','dashboard','','',1,'menu.dashboard',0,0,0,0,0,1,1),
(23,'2024-09-27 14:01:15','2025-03-11 15:00:38','告警总览','',1,1,2,'','/workplace','Workplace','','',1,'menu.alert.dashboard',0,0,0,0,0,1,1),
(32,'2024-10-11 16:18:45','2025-03-11 15:00:38','系统设置',NULL,1,0,21,'icon-settings','/system','system',NULL,NULL,1,'menu.system',0,0,0,0,0,1,1),
(33,'2024-10-11 16:25:41','2025-03-11 15:00:38','菜单设置',NULL,1,32,24,NULL,'/menu','system/menu',NULL,NULL,1,'menu.system.menu',0,0,0,0,0,1,1),
(34,'2024-10-12 14:10:15','2025-03-11 15:00:38','数据分析','',1,1,3,'','/data-analysis','DataAnalysis','','',0,'menu.dataAnalysis',0,0,0,0,0,1,0),
(35,'2024-10-12 14:17:59','2025-03-11 15:00:38','告警中心','',1,0,5,'icon-notification','/alert','alert','','',1,'menu.alert',0,0,1,0,0,1,1),
(36,'2024-10-12 14:25:57','2025-03-11 15:00:38','工作区','',1,35,6,'','/workspace','alert/workspace','','',1,'menu.alert.workspace',0,0,0,0,0,1,0),
(37,'2024-10-12 14:27:00','2025-03-11 15:00:38','告警集成','',1,35,7,'','/integration','alert/integration','','',0,'menu.source.integration',0,0,0,0,0,1,0),
(38,'2024-10-12 15:08:53','2025-03-11 15:00:38','告警通知管理','',1,0,13,'icon-bookmark','/template','notify','','',1,'menu.notify',0,0,0,0,0,1,0),
(39,'2024-10-12 15:12:21','2025-03-11 15:00:38','通知渠道管理','',1,38,14,'','/channels','channel/list','','',1,'menu.channel',0,0,0,0,0,1,1),
(40,'2024-10-12 15:13:16','2025-03-11 15:00:38','告警模板管理','',1,38,16,'','/list','template/list','','',1,'menu.template.list',0,0,0,0,0,1,0),
(41,'2024-10-12 15:20:24','2025-03-11 15:00:38','值班管理','',1,0,18,'icon-calendar','/schedule','schedule','','',1,'menu.schedule',0,0,0,0,0,1,1),
(42,'2024-10-12 15:22:29','2025-03-11 15:00:38','值班配置','',1,41,19,'','/index','schedule/index','','',1,'menu.schedule',0,0,0,0,0,1,1),
(43,'2024-10-12 15:26:17','2025-03-11 15:00:38','用户管理','',1,32,22,'','/user','system/user','','',1,'menu.system.user',0,0,0,0,0,1,0),
(44,'2024-10-12 15:27:04','2025-03-11 15:00:38','团队管理','',1,32,23,'','/team','system/team','','',1,'menu.system.team',0,0,0,0,0,1,0),
(45,'2024-10-12 15:27:58','2025-03-11 15:00:38','告警策略管理','',1,32,25,'','/strategy','system/strategy','','',1,'menu.system.strategy',0,0,0,0,0,1,0),
(46,'2024-10-12 15:28:50','2025-03-11 15:00:38','告警等级管理','',1,32,26,'','/level','system/level','','',1,'menu.system.level',0,0,0,0,0,1,0),
(47,'2024-10-14 11:03:24','2025-03-11 15:00:38','工作区详情','',1,35,8,'','workspace/detail/:id','alert/workspace/detail','','',1,'menu.alert.workspace',1,0,0,0,0,1,0),
(48,'2024-10-14 11:06:58','2025-03-11 15:00:38','告警集成详情','',1,35,9,'','integration/detail/:id','alert/integration/detail','','',1,'menu.source.integration',1,0,0,0,0,1,0),
(49,'2024-10-17 11:02:01','2025-03-11 15:00:38','告警模板详情','',1,38,17,'','detail/:id','template/detail','','',1,'menu.template.list',1,0,0,0,0,1,0),
(50,'2024-10-22 10:56:40','2025-03-11 15:00:38','告警规则管理','',1,0,10,'icon-layers','/alert-rule','alert-rule','','',1,'menu.alert.rule',0,0,0,0,0,1,1),
(51,'2024-10-22 11:00:17','2025-03-11 15:00:38','数据源','',1,50,11,'','/source','alert-rule-source','','',1,'menu.alert.rule.dataSource',0,1,0,0,0,1,0),
(52,'2024-10-22 19:15:28','2025-03-11 15:00:38','告警规则','',1,50,12,'','/list','alert-rule-list','','',0,'menu.alert.rule',0,0,0,0,0,1,0),
(53,'2024-12-25 15:32:11','2025-03-11 15:00:38','电视大盘','',1,1,4,'','/tv','TV','','',0,'menu.dashboard.tv',0,0,0,0,0,1,0),
(54,'2025-01-16 15:21:03','2025-03-11 15:00:38','排班','',1,41,20,'','detail/:id','schedule/detail','','',1,'menu.schedule.detail',1,0,0,1,0,1,0);

/*Table structure for table `schedule` */

DROP TABLE IF EXISTS `schedule`;

CREATE TABLE `schedule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '值班',
  `schedule_name` varchar(100) DEFAULT NULL COMMENT '值班名称',
  `schedule_desc` mediumtext COMMENT '值班描述',
  `team_id` int DEFAULT NULL COMMENT '团队编号',
  `start_range` time DEFAULT NULL COMMENT '值班开始时间',
  `end_range` time DEFAULT NULL COMMENT '值班结束时间',
  `period_type` enum('hour','day','week','month') DEFAULT NULL COMMENT '轮询类型',
  `period` tinyint DEFAULT NULL COMMENT '轮询时长',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '编辑时间',
  `users` json DEFAULT NULL COMMENT '值班用户',
  `schedule_user_count` int DEFAULT NULL COMMENT '值班人数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `schedule` */

/*Table structure for table `schedule_plan` */

DROP TABLE IF EXISTS `schedule_plan`;

CREATE TABLE `schedule_plan` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `schedule_id` int DEFAULT NULL COMMENT '值班编号',
  `month` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '周期',
  `plan` json DEFAULT NULL COMMENT '值班计划',
  `uid` int DEFAULT NULL COMMENT '用户编号',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=221 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `schedule_plan` */

/*Table structure for table `strategy_log` */

DROP TABLE IF EXISTS `strategy_log`;

CREATE TABLE `strategy_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `alert_id` int DEFAULT NULL COMMENT '告警事件编号',
  `uid` int DEFAULT NULL COMMENT '通知人编号',
  `strategy_content` json DEFAULT NULL COMMENT '告警策略信息',
  `strategy_id` int DEFAULT NULL COMMENT '告警策略编号',
  `channels` json DEFAULT NULL COMMENT '告警渠道',
  `is_notify` tinyint(1) DEFAULT '0' COMMENT '是否通知',
  `err_message` text COMMENT '通知错误信息',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `strategy_type` tinyint(1) DEFAULT '1' COMMENT '策略类型 1 默认策略 2 自定义策略',
  `notify_type` tinyint(1) DEFAULT '1' COMMENT '通知类型 1 告警 2 恢复',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=305 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `strategy_log` */

/*Table structure for table `system_integration` */

DROP TABLE IF EXISTS `system_integration`;

CREATE TABLE `system_integration` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `source_id` int DEFAULT NULL COMMENT '数据源编号',
  `description` text COMMENT '描述',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '集成名称',
  `token` varchar(100) DEFAULT NULL COMMENT '验签token',
  `level_field` varchar(100) DEFAULT NULL COMMENT '告警等级字段',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `disabled` tinyint(1) DEFAULT '0' COMMENT '是否禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `system_integration` */

insert  into `system_integration`(`id`,`source_id`,`description`,`name`,`token`,`level_field`,`created_at`,`updated_at`,`disabled`) values 
(3,1,'aaaa','Prometheus','f871c667-21ac-4aae-b99f-3eecde6a2897','labels.responseLevel','2024-08-08 14:06:56','2024-08-20 16:27:00',0),
(4,6,'','自定义事件','4753c308-5179-4dd9-838c-482b94eac35c','','2024-08-14 11:11:09','2024-08-14 11:11:09',0),
(8,7,'','阿里云监控 事件','6fcb31ca-8613-4c23-90e2-01bfccb99cdb','alert.meta.sysEventMeta.level','2024-12-27 15:08:09','2024-12-27 18:00:54',0);

/*Table structure for table `system_strategy` */

DROP TABLE IF EXISTS `system_strategy`;

CREATE TABLE `system_strategy` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `strategy_name` varchar(100) DEFAULT NULL COMMENT '分派策略名称',
  `strategy_desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '策略描述',
  `template_id` int DEFAULT NULL COMMENT '模板编号',
  `delay` int DEFAULT NULL COMMENT '延迟通知',
  `strategy_set` json DEFAULT NULL COMMENT '策略详情',
  `uid` int DEFAULT NULL COMMENT '修改人编号',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '变更时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `system_strategy` */

insert  into `system_strategy`(`id`,`strategy_name`,`strategy_desc`,`template_id`,`delay`,`strategy_set`,`uid`,`created_at`,`updated_at`) values 
(2,'默认策略','',5,0,'[{\"teams\": [], \"members\": [], \"alertLoop\": {\"count\": 2, \"enable\": false, \"interval\": 1}, \"schedules\": [], \"progression\": {\"enable\": false, \"duration\": 1, \"condition\": 0}, \"alertChannels\": [[], [], [], [], [], []]}]',1,'2024-05-28 11:53:10','2025-04-16 14:59:25');

/*Table structure for table `tag_rewrite` */

DROP TABLE IF EXISTS `tag_rewrite`;

CREATE TABLE `tag_rewrite` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `rewrite_type` varchar(20) NOT NULL COMMENT '重写类型',
  `integration_id` int NOT NULL COMMENT '集成编号',
  `filters` json DEFAULT NULL COMMENT '匹配条件',
  `rewrite_item` json DEFAULT NULL COMMENT '重写详情',
  `uid` int DEFAULT NULL COMMENT '修改用户',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `tag_rewrite` */

insert  into `tag_rewrite`(`id`,`rewrite_type`,`integration_id`,`filters`,`rewrite_item`,`uid`,`created_at`,`updated_at`) values 
(1,'extract',1,NULL,'{\"value\": \"\", \"newTag\": \"aaa\", \"oldTag\": \"core\", \"expression\": \"[a-z]\"}',33,'2024-08-07 14:07:46','2024-08-07 14:50:02'),
(2,'rewrite',1,NULL,'{\"value\": \"122\", \"newTag\": \"1233\", \"oldTag\": \"wad\", \"deleteTag\": null, \"expression\": \"qwqw\"}',33,'2024-08-07 14:40:26','2024-08-07 15:42:54'),
(4,'extract',1,NULL,'{\"value\": \"\", \"newTag\": \"dasd\", \"oldTag\": \"dsadas\", \"expression\": \"asdas\"}',33,'2024-08-07 15:07:42','2024-08-07 15:07:42'),
(5,'delete',1,NULL,'{\"value\": \"\", \"newTag\": \"\", \"oldTag\": \"\", \"deleteTag\": [\"aaa\", \"sss\"], \"expression\": \"\"}',33,'2024-08-07 15:24:03','2024-08-07 15:25:19');

/*Table structure for table `team_rule` */

DROP TABLE IF EXISTS `team_rule`;

CREATE TABLE `team_rule` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `menu_id` int DEFAULT NULL,
  `team_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=235 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `team_rule` */

insert  into `team_rule`(`id`,`menu_id`,`team_id`) values 
(20,35,4),
(21,36,4),
(22,37,4),
(23,47,4),
(24,48,4),
(209,35,1),
(210,32,1),
(211,39,1),
(212,40,1),
(213,49,1),
(214,50,1),
(215,23,1),
(216,34,1),
(217,53,1),
(218,1,1),
(219,41,1),
(220,36,1),
(221,37,1),
(222,47,1),
(223,48,1),
(224,43,1),
(225,44,1),
(226,33,1),
(227,45,1),
(228,46,1),
(229,51,1),
(230,52,1),
(231,42,1),
(232,54,1),
(233,55,1),
(234,38,1);

/*Table structure for table `teams` */

DROP TABLE IF EXISTS `teams`;

CREATE TABLE `teams` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `team_name` varchar(100) DEFAULT NULL COMMENT '团队名称',
  `team_desc` text COMMENT '团队描述',
  `updated_by` int DEFAULT NULL COMMENT '更新人编号',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态 1 启用  0 禁用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `data_permission` tinyint(1) DEFAULT '1' COMMENT '数据权限 1. 只看自己的 2. 看全部的',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `teams` */

insert  into `teams`(`id`,`team_name`,`team_desc`,`updated_by`,`status`,`created_at`,`updated_at`,`data_permission`) values 
(1,'测试团队','11111',33,1,'2023-12-28 15:27:49','2025-03-11 15:34:52',1),
(4,'主管团队','管理员团队',1,1,'2024-10-16 15:34:18','2024-10-16 15:34:56',1);

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '显示名',
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名',
  `password` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '密码',
  `empno` varchar(156) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '员工号',
  `userid` varchar(156) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'UserId',
  `email` varchar(156) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '电话',
  `user_type` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'local' COMMENT '用户类型:local,ldap,sso',
  `path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '节点路径',
  `avatar` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'default' COMMENT '用户头像',
  `otp_secret` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'OTP双因子认证密钥',
  `otp_enabled` tinyint(1) DEFAULT '0' COMMENT '启用OTP双因子认证',
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '1' COMMENT '状态',
  `by_update` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '更新人',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_name_code` (`nickname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `user` */

insert  into `user`(`id`,`nickname`,`username`,`password`,`empno`,`userid`,`email`,`phone`,`user_type`,`path`,`avatar`,`otp_secret`,`otp_enabled`,`status`,`by_update`,`created_at`,`updated_at`) values 
(1,'root','admin','21232f297a57a5a743894a0e4a801fc3',NULL,'00079','xxx@xx.cn','13800000000','local','[[1]]','head_lufei_2','QAZFJJEWQQZ3GVSCHTMIKOA2NM',1,'1','root','2022-07-22 12:41:10','2025-04-15 14:10:44');

/*Table structure for table `user_roles` */

DROP TABLE IF EXISTS `user_roles`;

CREATE TABLE `user_roles` (
  `user_id` bigint NOT NULL COMMENT '主键id',
  `role_id` bigint NOT NULL COMMENT '主键id',
  PRIMARY KEY (`user_id`,`role_id`) USING BTREE,
  CONSTRAINT `fk_user_roles_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `user_roles` */

insert  into `user_roles`(`user_id`,`role_id`) values 
(1,1);

/*Table structure for table `user_teams` */

DROP TABLE IF EXISTS `user_teams`;

CREATE TABLE `user_teams` (
  `user_id` int unsigned NOT NULL COMMENT '用户编号',
  `team_id` int unsigned NOT NULL COMMENT '团队编号'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `user_teams` */

insert  into `user_teams`(`user_id`,`team_id`) values 
(1,1);

/*Table structure for table `workspace` */

DROP TABLE IF EXISTS `workspace`;

CREATE TABLE `workspace` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `workspace_name` varchar(100) NOT NULL COMMENT '工作区名称',
  `workspace_desc` varchar(500) DEFAULT NULL COMMENT '工作区描述',
  `system_strategy` int DEFAULT NULL COMMENT '全局告警策略',
  `enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `workspace` */

insert  into `workspace`(`id`,`workspace_name`,`workspace_desc`,`system_strategy`,`enable`,`created_at`,`updated_at`) values 
(7,'运维','',2,1,'2024-05-13 14:41:03','2024-05-28 14:47:41'),
(8,'研发','',2,1,'2024-05-21 10:54:05','2024-08-20 16:46:05');

/*Table structure for table `workspace_router` */

DROP TABLE IF EXISTS `workspace_router`;

CREATE TABLE `workspace_router` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `workspace_id` int NOT NULL,
  `router_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `workspace_router` */

insert  into `workspace_router`(`id`,`workspace_id`,`router_id`) values 
(48,7,21),
(61,7,22),
(62,7,29);

/*Table structure for table `workspace_source` */

DROP TABLE IF EXISTS `workspace_source`;

CREATE TABLE `workspace_source` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `workspace_id` int NOT NULL COMMENT '告警源编号',
  `description` text COMMENT '告警源描述',
  `source_id` int NOT NULL COMMENT '告警源编号',
  `verified` tinyint(1) DEFAULT '0' COMMENT '是否验证',
  `configuration` json DEFAULT NULL COMMENT '配置信息',
  `source_name` varchar(100) DEFAULT NULL COMMENT '告警源名称',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `token` varchar(100) DEFAULT NULL COMMENT '请求token',
  `level_field` varchar(100) DEFAULT NULL COMMENT '告警等级字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `workspace_source` */

insert  into `workspace_source`(`id`,`workspace_id`,`description`,`source_id`,`verified`,`configuration`,`source_name`,`created_at`,`updated_at`,`token`,`level_field`) values 
(13,7,NULL,1,1,'{}','Prometheus','2024-05-13 14:41:03','2024-06-28 13:44:23','94e5e56a-980a-4655-a361-2de343d7d9ef','labels.responseLevel'),
(14,8,NULL,1,1,'{}','Prometheus','2024-05-21 10:54:05','2024-05-21 10:54:05','414e3ddd-a5f0-404b-b7dc-889c118c0060','labels.severity'),
(15,7,'asdasdad',6,1,'{}','自定义事件','2024-05-24 17:19:39','2024-06-25 11:26:55','e3ec998f-3f7c-4920-a6cf-6404d8227dfd',''),
(19,7,'',1,1,'{}','Prometheus','2024-08-08 11:08:13','2024-08-08 11:08:13','a5aecac8-7104-445f-857b-e915f70d27eb','labels.severity');

/*Table structure for table `workspace_team` */

DROP TABLE IF EXISTS `workspace_team`;

CREATE TABLE `workspace_team` (
  `team_id` int NOT NULL COMMENT '团队编号',
  `workspace_id` int DEFAULT NULL COMMENT '工作区编号'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `workspace_team` */

insert  into `workspace_team`(`team_id`,`workspace_id`) values 
(1,8),
(1,7),
(4,8);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
