CREATE TABLE `delivery_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标题',
  `source` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '来源',
  `domain_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '域名',
  `delivery_date` date DEFAULT NULL COMMENT '投放日期',
  `author` int(11) DEFAULT NULL COMMENT '作者',
  `deliverer` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '投放人',
  `status` tinyint(2) DEFAULT NULL COMMENT '投放状态 0待投放 1已投放 2 投放失败',
  `result` text COLLATE utf8mb4_unicode_ci COMMENT '投放执行结果',
  `posts_id` int(11) DEFAULT NULL COMMENT '文章id',
  `station_id` int(11) DEFAULT NULL COMMENT '站点id',
  `wp_cate_ids` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'wp分类',
  `create_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `update_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_posts_station` (`posts_id`,`station_id`),
  KEY `idx_status` (`status`) USING BTREE COMMENT '用来获取为0的数据初始化'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='投放日志';