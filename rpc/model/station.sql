CREATE TABLE `station` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `domain_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '域名',
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip',
  `domain_year` int(11) DEFAULT NULL COMMENT '域名年份',
  `google_weight` decimal(10,0) DEFAULT NULL COMMENT '谷歌权重',
  `type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网站类型',
  `industry` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网站行业',
  `articles_num` int(11) DEFAULT NULL COMMENT '文章数量',
  `user_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '账号名',
  `pass_word` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '密码',
  `status` tinyint(2) DEFAULT NULL COMMENT '状态 0 正常 1 失效 2 删除',
  `creater` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '创建人',
  `create_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `modifier` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '修改人',
  `update_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点信息';

CREATE TABLE `posts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标题',
  `source` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '来源',
  `author` int(11) DEFAULT NULL COMMENT '作者',
  `thrown_num` int(11) DEFAULT NULL COMMENT '投放数量',
  `categories` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '分类',
  `creater` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '创建人',
  `create_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `modifier` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '修改人',
  `update_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点博客内容';

CREATE TABLE `station_posts_relation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `station_id` int(10) unsigned NOT NULL COMMENT '站点id',
  `posts_id` int(10) unsigned NOT NULL COMMENT '内容id',
  `wp_id` int(10) unsigned DEFAULT '0' COMMENT 'wp id',
  PRIMARY KEY (`id`),
  KEY `uk_posts_station` (`posts_id`,`station_id`),
  KEY `idx_station_id` (`station_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点与post关系表';

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

CREATE TABLE `station_posts_text` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `posts_id` int(11) DEFAULT NULL COMMENT ' posts 主键',
  `content` mediumtext COLLATE utf8mb4_unicode_ci COMMENT 'html内容',
  `create_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `update_at` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_posts_id` (`posts_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='站点博客内容大字段';
