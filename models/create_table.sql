# create database `db_bbs`;
use `db_bbs`;
CREATE TABLE `user`
(
    `id`          bigint(20)                             NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(20)                             NOT NULL,
    `username`    varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password`    varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email`       varchar(64) COLLATE utf8mb4_general_ci,
    `gender`      tinyint(4)                             NOT NULL DEFAULT '0',
    `create_time` timestamp                              NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp                              NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`
(
    `id`             int(11)                                 NOT NULL AUTO_INCREMENT,
    `community_id`   int(10) unsigned                        NOT NULL,
    `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction`   varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`community_id`),
    UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;


INSERT INTO `community`
VALUES ('1', '1', 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO `community`
VALUES ('2', '2', 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00');
INSERT INTO `community`
VALUES ('3', '3', 'CS:GO', 'Rush B。。。', '2018-08-07 08:30:00', '2018-08-07 08:30:00');
INSERT INTO `community`
VALUES ('4', '4', 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');
