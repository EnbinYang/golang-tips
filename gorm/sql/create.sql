CREATE TABLE `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `uid` int(11) NOT NULL COMMENT '用户id',
    `keywords` text NOT NULL COMMENT '索引词',
    `degree` char(2) NOT NULL COMMENT '学历',
    `gender` char(2) NOT NULL COMMENT '性别',
    `city` char(2) NOT NULL COMMENT '城市',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=1707919 DEFAULT CHARSET=utf8 COMMENT='用户信息表';