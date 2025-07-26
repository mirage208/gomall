-- 收藏商品表
CREATE TABLE `favorite_product` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态 0:未删除 1:已删除',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户id',
  `product_id` bigint NOT NULL DEFAULT '0',
  `product_title` varchar(32) NOT NULL DEFAULT '' COMMENT '商品标题',
  `store_id` bigint NOT NULL DEFAULT '0' COMMENT '收藏商品店铺id',
  PRIMARY KEY (`id`),
  KEY `ix_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
