-- 秒杀商品表
CREATE TABLE `seckill_product` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态，0-未删除，1-已删除',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
  `product_id` bigint NOT NULL DEFAULT '0' COMMENT '参与秒杀的商品id',
  `store_id` bigint NOT NULL DEFAULT '0' COMMENT '秒杀商品的店铺id',
  `seckill_price` bigint NOT NULL DEFAULT '0' COMMENT '商品秒杀价格（分）',
  `stock_count` int NOT NULL DEFAULT '0' COMMENT '秒杀商品的库存',
  `start_time` date NOT NULL DEFAULT '1970-01-01' COMMENT '秒杀开始日期（不包括具体时间）',
  `time` int NOT NULL DEFAULT '0' COMMENT '秒杀开始的整点时间（时间为10，12，14）',
  PRIMARY KEY (`id`),
  KEY `ix_time` (`start_time`, `time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
