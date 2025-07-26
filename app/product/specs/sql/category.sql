-- 商品分类表
CREATE TABLE `category` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态 0:未删除 1:已删除',
  `name` varchar(12) NOT NULL DEFAULT '' COMMENT '分类名称',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
