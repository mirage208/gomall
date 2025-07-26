-- 商品推荐表
CREATE TABLE `product_recommend` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态 0:未删除 1:已删除',
  `product_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  PRIMARY KEY (`id`),
  KEY `ix_update_time` (`update_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
