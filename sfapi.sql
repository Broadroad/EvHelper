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
DROP TABLE IF EXISTS `sf_order`;
CREATE TABLE `sf_order` (
    `orderid` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `userid` int(10) NOT NULL COMMENT '接单用户id',
    `senderid` int(10) unsigned NOT NULL COMMENT '发单用户id',
    `sendertime` datetime DEFAULT NULL COMMENT '发单时间',
    `ordertime` datetime DEFAULT NULL COMMENT '接单时间',
    `usergetpackagetime` datetime DEFAULT NULL COMMENT '接单用户获得包裹时间',
    `sendergetpackagetime` datetime DEFAULT NULL COMMENT '发单用户收到包裹时间',
    `expresslocation` varchar(100) NOT NULL COMMENT '接单用户取包裹地点',
    `senderlocation` varchar(100) NOT NULL COMMENT '发单用户接收包裹地点',
    `expresslatitude` varchar(100) DEFAULT NULL COMMENT '接单用户取包裹的纬度信息',
    `expresslogitude` varchar(100) DEFAULT NULL COMMENT '接单用户取包裹的纬度信息',
    `senderlatitude` varchar(100) DEFAULT NULL COMMENT '发单用户取包裹的纬度信息',
    `senderlogitude` varchar(100) DEFAULT NULL COMMENT '发单用户取包裹的纬度信息',
    `price` DECIMAL(5,2) DEFAULT NULL COMMENT '发单用户出价', 
    `telephone` varchar(11) DEFAULT NULL COMMENT '发单用户接收包裹时候的联系人',
    `is_deleted` int(11) DEFAULT NULL COMMENT '该订单是否被用户删了',
    `description` varchar(100) DEFAULT NULL COMMENT '发单人的备注',
    `is_ordered` int(11) DEFAULT NULL COMMENT '该订单是否被接单',
    `is_completed` int(11) DEFAULT NULL COMMENT '该订单是否完成',
    PRIMARY KEY (`orderid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
SET FOREIGN_KEY_CHECKS = 1;
