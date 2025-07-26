CREATE TABLE `user_header_image` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `del_state` tinyint NOT NULL DEFAULT 0 COMMENT '删除状态：0-未删除，1-已删除',
    `path` varchar(255) NOT NULL DEFAULT '' COMMENT '图片路径',
    `type` tinyint (1) NOT NULL DEFAULT 2 COMMENT '图片存储类型：1-本地存储，2-七牛云存储',
    `hash` longtext NOT NULL DEFAULT '' COMMENT '图片hash值，防止重复图片上传',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;