-- 商品店铺表
CREATE TABLE `product_store` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态，0-未删除，1-已删除',
  `title` varchar(32) NOT NULL DEFAULT '' COMMENT '店铺名称',
  `boss_id` bigint NOT NULL DEFAULT '0' COMMENT '店主用户id',
  `introduction` varchar(255) NOT NULL DEFAULT '' COMMENT '店铺介绍',
  `state` tinyint (1) NOT NULL DEFAULT '1' COMMENT '0-禁止营业，1-正常营业',
  `collect_num` bigint NOT NULL DEFAULT '0' COMMENT '店铺收藏数',
  `cover` varchar(1024) NOT NULL DEFAULT '' COMMENT '封面图',
  PRIMARY KEY (`id`),
  KEY `ix_boss_id` (`boss_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
