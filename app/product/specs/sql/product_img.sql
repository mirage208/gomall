-- 商品图片表
CREATE TABLE `product_img` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT 0 COMMENT '删除状态：0-未删除，1-已删除',
  `product_id` bigint NOT NULL DEFAULT '0',
  `path` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片路径',
  `type` tinyint (1) NOT NULL DEFAULT 1 COMMENT '图片存储类型：1-本地存储，2-七牛云存储',
  `is_cover` tinyint (1) NOT NULL DEFAULT '0' COMMENT '是否为商品封面图片',
  `img_order` tinyint (1) NOT NULL DEFAULT '0' COMMENT '图片排序',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
