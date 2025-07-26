-- 商品评论表
CREATE TABLE `product_comment` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态 0:未删除 1:已删除',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户id',
  `product_id` bigint NOT NULL DEFAULT '0' COMMENT '商品id',
  `is_good` tinyint (1) NOT NULL DEFAULT '0' COMMENT '-1-差评，0-中性，1-好评',
  `content` text NOT NULL DEFAULT '' COMMENT '评论内容',
  `add_content` varchar(1024) NOT NULL DEFAULT '' COMMENT '追加评论内容',
  PRIMARY KEY (`id`),
  KEY `ix_user_id` (`user_id`),
  KEY `ix_product_id` (`product_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
