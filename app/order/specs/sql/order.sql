-- 支付表
CREATE TABLE `payment` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0 COMMENT '删除状态：0-未删除；1-已删除',
    `version` bigint NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
    `payment_sn` varchar(25) NOT NULL DEFAULT '' COMMENT '支付流水单号',
    `order_sn` varchar(25) NOT NULL DEFAULT '' COMMENT '订单号',
    `user_id` bigint NOT NULL DEFAULT '0',
    `pay_mode` tinyint (1) NOT NULL DEFAULT '1' COMMENT '支付方式：1-平台钱包支付，2-微信支付，3-支付宝支付',
    `trade_type` varchar(20) NOT NULL DEFAULT '' COMMENT '第三方支付类型',
    `trade_state` varchar(20) NOT NULL DEFAULT '' COMMENT '第三方支付状态',
    `pay_total` bigint NOT NULL DEFAULT '0' COMMENT '支付总金额（分）',
    `transaction_id` char(32) NOT NULL DEFAULT '' COMMENT '第三方支付单号',
    `trade_state_desc` varchar(256) NOT NULL DEFAULT '' COMMENT '支付状态描述',
    `pay_status` tinyint (1) NOT NULL DEFAULT '0' COMMENT '交易状态：-1-支付失败；0-未支付；1-支付成功；2-已退款',
    `pay_time` datetime NOT NULL DEFAULT '1970-01-01 08:00:00' COMMENT '支付成功时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `ix_payment_sn` (`payment_sn`),
    UNIQUE KEY `ix_order_sn` (`order_sn`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;