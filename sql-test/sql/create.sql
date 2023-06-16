CREATE TABLE `llm_manage_state` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
    `resource_type` VARCHAR(50) NOT NULL DEFAULT "" COMMENT '资源类型, 1-专栏, 4-视频',
    `resource_aid` BIGINT NOT NULL DEFAULT 0 COMMENT '资源id',
    `abstract` VARCHAR(2048) NOT NULL DEFAULT "" COMMENT '摘要信息',
    `state` TINYINT NOT NULL DEFAULT 0 COMMENT '状态, 1-已上线, 0-已下线',
    `c_uname` VARCHAR(50) NOT NULL DEFAULT "" COMMENT '创建人',
    `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_mtime` (`mtime`),
    INDEX `manuscript_ix` (`resource_type`, `resource_aid`);
) ENGINE=InnoDB COMMENT='更新稿件state';