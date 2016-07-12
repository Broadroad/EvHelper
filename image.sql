/*
 Navicat Premium Data Transfer
 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50548
 Source Host           : localhost
 Source Database       : scheduler
 Target Server Type    : MySQL
 Target Server Version : 50548
 File Encoding         : utf-8
 Date: 05/24/2016 10:29:45 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  le structure for `Order`
-- ----------------------------
DROP TABLE IF EXISTS `images`;
CREATE TABLE `images` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uploadtime` datetime DEFAULT NULL COMMENT '上传时间',
    `userid` int(10) unsigned NOT NULL COMMENT '该图片所属用户的id',
    `picname` varchar(100) NOT NULL COMMENT '图片名',
    `key` varchar(100) DEFAULT NULL COMMENT '七牛对应的key',
    `point` int(11) DEFAULT NULL COMMENT '阅读该图片所需要的point', 
    `is_deleted` int(11) DEFAULT NULL COMMENT '是否删除',
    `description` varchar(100) DEFAULT NULL COMMENT '对于图片的描述',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
SET FOREIGN_KEY_CHECKS = 1;
