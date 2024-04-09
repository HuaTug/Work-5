
use Hertz;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
SET global log_bin_trust_function_creators = 1;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
use Hertz;
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
                             `comment_id` bigint NOT NULL AUTO_INCREMENT,
                             `user_id` bigint NOT NULL,
                             `video_id` bigint NOT NULL,
                             `comment` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL,
                             `time` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL,
                             `index_id` bigint NOT NULL ,
                             `favorite_count` bigint NOT NULL default 0,
                             PRIMARY KEY (`comment_id`) USING BTREE,
                             INDEX `commentUser`(`user_id`) USING BTREE,
                             INDEX `commentVideo`(`video_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 0;

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
use Hertz;
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`  (
                              `favorite_id` bigint NOT NULL AUTO_INCREMENT,
                              `user_id` bigint NOT NULL,
                              `video_id` bigint NOT NULL,
                              `comment_id` bigint NOT NULL default 0,
                              `video_type` bigint NOT NULL ,
                              PRIMARY KEY (`favorite_id`) USING BTREE,
                              INDEX `favoriteUser`(`user_id`) USING BTREE,
                              INDEX `favoriteVideo`(`video_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 0;

-- ----------------------------
-- Table structure for relations
-- ----------------------------
DROP TABLE IF EXISTS `relations`;
use Hertz;
CREATE TABLE `relations`  (
                              `relation_id` bigint NOT NULL AUTO_INCREMENT,
                              `follow_id` bigint NOT NULL,
                              `follower_id` bigint NOT NULL,
                              `user_id`  bigint NOT NULL ,
                              PRIMARY KEY (`relation_id`) USING BTREE,
                              INDEX `FollowId`(`follow_id`) USING BTREE,
                              INDEX `FollowerId`(`follower_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `user_id` bigint NOT NULL AUTO_INCREMENT,
                          `user_name` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
                          `password` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
                          `follow_count` bigint NULL DEFAULT NULL,
                          `follower_count` bigint NULL DEFAULT NULL,
                          `favorite_count` bigint NULL DEFAULT NULL,
                          PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 0;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
                           `video_id` bigint NOT NULL AUTO_INCREMENT,
                           `author_id` bigint NOT NULL,
                           `play_url` varchar(255) CHARACTER SET utf8mb4  DEFAULT NULL,
                           `cover_url` varchar(255) CHARACTER SET utf8mb4  DEFAULT NULL,
                           `favorite_count` bigint NULL DEFAULT NULL,
                           `comment_count` bigint NULL DEFAULT NULL,
                           `publish_time` varchar(255) NULL DEFAULT NULL,
                           `title` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL,
                           PRIMARY KEY (`video_id`) USING BTREE,
                           INDEX `user`(`author_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 0;

use Hertz;
Drop TABLE IF EXISTS  `messages`;
CREATE TABLE `messages`(
                           `message_id` bigint primary key auto_increment,
                           `sender_id` bigint,
                           `receiver_id` bigint,
                           `message_text` varchar(255),
                           `send_time` timestamp,
                           `state` bigint
)ENGINE = InnoDB AUTO_INCREMENT = 0;
-- ----------------------------
-- Triggers structure for table comments
-- ----------------------------
DROP TRIGGER IF EXISTS `comment_action`;
delimiter ;;
CREATE TRIGGER `comment_action` AFTER INSERT ON `comments` FOR EACH ROW update videos set comment_count = comment_count + 1 where videos.video_id = new.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table comments
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_comment`;
delimiter ;;
CREATE TRIGGER `delete_comment` AFTER DELETE ON `comments` FOR EACH ROW update videos set comment_count = comment_count - 1 where videos.video_id = old.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `like_action`;
delimiter ;;
CREATE TRIGGER `like_action` AFTER INSERT ON `favorites` FOR EACH ROW update videos set favorite_count = favorite_count + 1 where videos.video_id = new.video_id and new.video_type=1
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `unlike_action`;
delimiter ;;
CREATE TRIGGER `unlike_action` AFTER DELETE ON `favorites` FOR EACH ROW update videos set favorite_count = favorite_count - 1 where videos.video_id = old.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `fav_count`;
delimiter ;;
CREATE TRIGGER `fav_count` AFTER INSERT ON `favorites` FOR EACH ROW
BEGIN update users set users.favorite_count = users.favorite_count + 1 where users.user_id = new.user_id;
update comments set comments.favorite_count=comments.favorite_count+1 where comments.comment_id=new.comment_id;
END;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `unfav_count`;
delimiter ;;
CREATE TRIGGER `unfav_count` AFTER DELETE ON `favorites` FOR EACH ROW update users set users.favorite_count = users.favorite_count - 1 where users.user_id = old.user_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table relations
-- ----------------------------
DROP TRIGGER IF EXISTS `false_follow_action`;
delimiter ;;
CREATE TRIGGER `false_follow_action` AFTER DELETE ON `relations` FOR EACH ROW update users set follow_count = follow_count - 1 where users.user_id = old.follow_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table relations
-- ----------------------------
DROP TRIGGER IF EXISTS `false_follower_action`;
delimiter ;;
CREATE TRIGGER `false_follower_action` AFTER DELETE ON `relations` FOR EACH ROW update users set follower_count = follower_count - 1 where users.user_id = old.follower_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table relations
-- ----------------------------
DROP TRIGGER IF EXISTS `follow_action`;
delimiter ;;
CREATE TRIGGER `follow_action` AFTER INSERT ON `relations` FOR EACH ROW
BEGIN
    update users set follow_count = follow_count + 1 where users.user_id = new.user_id;
    update users set follower_count=follower_count+1 where users.user_id=new.follow_id;
END
;;
delimiter ;

