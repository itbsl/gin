-- 创建blog数据库
CREATE DATABASE `blog` CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;

-- 标签表
CREATE TABLE `blog_tag` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL DEFAULT "" COMMENT "标签名称",
    `state` tinyint unsigned DEFAULT 1 COMMENT "状态:0为禁用、1为启用",
    `create_time` int(11) unsigned DEFAULT 0 COMMENT "创建时间",
    `update_time` int(11) unsigned DEFAULT 0 COMMENT "更新时间",
    `delete_time` int(11) unsigned DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ENGINE=InnoDB COMMENT "标签表";

-- 文章表
CREATE TABLE `blog_article` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL DEFAULT "" COMMENT "标题",
    `desc` varchar(500) NOT NULL DEFAULT  "" COMMENT "简述",
    `cover` varchar(255) NOT NULL DEFAULT "" COMMENT "封面",
    `content` text COMMENT "内容",
    `state` tinyint unsigned NOT NULL DEFAULT 1 COMMENT "状态:0为禁用、1为启用",
    `create_time` int(11) unsigned DEFAULT 0 COMMENT "创建时间",
    `update_time` int(11) unsigned DEFAULT 0 COMMENT "更新时间",
    `delete_time` int(11) unsigned DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ENGINE=InnoDB COMMENT "文章表";

-- 文章标签关联表
CREATE TABLE `blog_article_tag` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `article_id` int(11) unsigned NOT NULL COMMENT "文章ID",
    `tag_id` int(11) unsigned NOT NULL COMMENT "标签ID",
    `create_time` int(11) unsigned DEFAULT 0 COMMENT "创建时间",
    `update_time` int(11) unsigned DEFAULT 0 COMMENT "更新时间",
    `delete_time` int(11) unsigned DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY (`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ENGINE=InnoDB COMMENT "文章标签关联表";