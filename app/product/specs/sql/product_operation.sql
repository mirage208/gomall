-- 商品运营表
CREATE TABLE `product_operation` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT 0 COMMENT '删除状态：0-未删除，1-已删除',
  `product_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `status` int NOT NULL DEFAULT 1 COMMENT '运营商品状态 0-下线 1-上线',
  PRIMARY KEY (`id`),
  KEY `ix_update_time` (`update_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
