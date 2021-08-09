-- 
-- 学习MySQL优化构造数据
-- 

-- 数据库: mysql_study
-- CREATE DATABASE mysql_study;

-- 用户表
DROP FUNCTION IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(36) NOT NULL,
  `user_name` varchar(12) DEFAULT NULL,
  `age` tinyint(3) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `province` varchar(10) DEFAULT NULL,
  `city` varchar(10) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--  订单表
DROP FUNCTION IF EXISTS `order`;
CREATE TABLE `order` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) DEFAULT NULL,
  `product_count` int(11) DEFAULT NULL,
  `price` decimal(10,0) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4