CREATE TABLE `station_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64)  DEFAULT NULL COMMENT '姓名 页面显示',
  `email` varchar(64)  DEFAULT NULL COMMENT '邮箱',
  `slug` varchar(64)  DEFAULT NULL COMMENT '用户标识符',
  `username` varchar(64)  DEFAULT NULL COMMENT '登录名称',
  `password` varchar(64)  DEFAULT NULL COMMENT '登录密码',
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_slug` (`slug`) USING BTREE,
  UNIQUE KEY `uk_username` (`username`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点用户表';
