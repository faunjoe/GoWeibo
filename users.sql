/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : localhost
 Source Database       : zhihu

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : utf-8

 Date: 01/11/2018 20:32:56 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `rememberToken` varchar(255) DEFAULT NULL,
  `timestamps` bigint(20) DEFAULT NULL,
  `questions_count` int(11) DEFAULT '0',
  `answers_count` int(11) DEFAULT '0',
  `comments_count` int(11) DEFAULT '0',
  `favorites_count` int(11) DEFAULT '0',
  `likes_count` int(11) DEFAULT '0',
  `follows_count` int(11) DEFAULT NULL,
  `followins_count` int(11) DEFAULT '0',
  `setting` longtext,
  `is_activity` smallint(6) DEFAULT NULL,
  `confirmation_token` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
