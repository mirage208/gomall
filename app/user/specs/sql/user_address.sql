CREATE TABLE `user_address` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0 COMMENT '删除状态：0-未删除，1-已删除',
    `user_id` bigint NOT NULL DEFAULT 0 COMMENT '用户id',
    `is_default` tinyint (1) NOT NULL DEFAULT 0 COMMENT '是否为默认地址',
    `province` varchar(100) NOT NULL DEFAULT '' COMMENT '收货省份',
    `city` varchar(100) NOT NULL DEFAULT '' COMMENT '收货城市',
    `region` varchar(100) NOT NULL DEFAULT '' COMMENT '收货区/县',
    `detail_address` varchar(255) NOT NULL DEFAULT '' COMMENT '收货详情地址',
    `name` varchar(64) NOT NULL DEFAULT '' COMMENT '收货人名称',
    `phone` varchar(11) NOT NULL DEFAULT '' COMMENT '收货人手机号',
    PRIMARY KEY (`id`),
    KEY `ix_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;