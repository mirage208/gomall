-- 购物车表
CREATE TABLE `cart` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0 COMMENT '删除状态：0-未删除；1-已删除',
    `user_id` bigint NOT NULL DEFAULT '0',
    `product_id` bigint NOT NULL DEFAULT '0',
    `count` int NOT NULL DEFAULT 1 COMMENT '数量',
    `checked` tinyint (1) NOT NULL DEFAULT '0' COMMENT '是否勾选：1-已勾选，0-未勾选',
    PRIMARY KEY (`id`),
    KEY `ix_user_id` (`user_id`),
    KEY `ix_product_id` (`product_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;