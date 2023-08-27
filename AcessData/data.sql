CREATE DATABASE if not exists TikTok;

use TikTok;
-- ----------------------------
-- Table structure for video
-- ----------------------------
CREATE TABLE if not exists video(
    `id` int NOT NULL auto_increment primary key ,-- 自增序列
    `title` varchar(255) ,
    `play_url` varchar(255),
    `cover_url` varchar(255),
    `publish_time` timestamp(0)
 );
-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (44, 'D', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308191519', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308191519', '2023-08-19 15:19:36');
INSERT INTO `video` VALUES (45, 'D', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308191520', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308191520', '2023-08-19 15:21:09');
INSERT INTO `video` VALUES (46, 'D', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308191521', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308191521', '2023-08-19 15:21:31');
INSERT INTO `video` VALUES (25, 'B', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308181647', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308181647', '2023-08-18 16:47:34');
INSERT INTO `video` VALUES (23, 'A', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308181643', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308181643', '2023-08-18 16:43:05');
INSERT INTO `video` VALUES (47, 'D', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308191521', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308191521', '2023-08-19 15:21:49');
INSERT INTO `video` VALUES (26, 'C', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308191419', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308191419', '2023-08-19 14:19:31');
INSERT INTO `video` VALUES (48, 'D', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308191523', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308191523', '2023-08-19 15:23:41');
INSERT INTO `video` VALUES (49, 'D', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308191527', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308191527', '2023-08-19 15:27:33');
INSERT INTO `video` VALUES (50, 'D', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/video/202308191528', 'http://rynw10bh8.hn-bkt.clouddn.com/test-TikTok/cover/202308191528', '2023-08-19 15:28:11');

CREATE TABLE if not exists user_video (
    `user_id` int,
    `video_id` int
 );
-- ----------------------------
-- Records of user_video
-- ----------------------------
INSERT INTO `user_video` VALUES (1, 23);
INSERT INTO `user_video` VALUES (1, 25);
INSERT INTO `user_video` VALUES (1, 26);
INSERT INTO `user_video` VALUES (1, 44);
INSERT INTO `user_video` VALUES (1, 45);
INSERT INTO `user_video` VALUES (1, 46);
INSERT INTO `user_video` VALUES (1, 47);
INSERT INTO `user_video` VALUES (1, 48);
INSERT INTO `user_video` VALUES (1, 49);
INSERT INTO `user_video` VALUES (1, 50);

-- ----------------------------
-- Table structure for user_video_favorite
-- ----------------------------
CREATE TABLE if not exists user_video_favorite (
    `user_id` int,
    `video_id` int
);
-- ----------------------------
-- Records of user_video_favorite
-- ----------------------------
INSERT INTO `user_video_favorite` VALUES (1, 59);
INSERT INTO `user_video_favorite` VALUES (1, 60);
INSERT INTO `user_video_favorite` VALUES (1, 56);
INSERT INTO `user_video_favorite` VALUES (1, 56);
INSERT INTO `user_video_favorite` VALUES (1, 59);
INSERT INTO `user_video_favorite` VALUES (1, 60);
INSERT INTO `user_video_favorite` VALUES (1, 1);


-- ----------------------------
-- Table structure for user_follow
-- ----------------------------
CREATE TABLE if not exists user_follow (
    `user_id` int,
    `follower_id` int
);

-- ----------------------------
-- Records of user_follow
-- ----------------------------
INSERT INTO `user_follow` VALUES (1, 1);
INSERT INTO `user_follow` VALUES (1, 2);
INSERT INTO `user_follow` VALUES (2, 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
CREATE TABLE if not exists user (
    `id` int4,
    `name` varchar(255) ,
    `password` varchar(255),
    `avatar` varchar(255),
    `background_image` varchar(255),
    `signature` varchar(255)
);

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'ithwind', '$2a$10$6QbKR8wBjwkvfkEc14aT0e3oe4zEugGqiSfQcYOI8vHznH8Mx/mTa', 'www.baidu.com', 'www.baidu.com', '加油！');
INSERT INTO `user` VALUES (0, 'AI', '$2a$10$4Hg03ATgAvGHPq7z2WZLfOBH9808Tn2lG3VtZt32CmoaEzIRikMMu
', 'https://th.bing.com/th/id/OIP.YeYI4ACa3020bl89ARJSywHaHa?pid=ImgDet&rs=1', 'https://th.bing.com/th/id/OIP.YeYI4ACa3020bl89ARJSywHaHa?pid=ImgDet&rs=1', '文心一言');



-- ----------------------------
-- Table structure for friend
-- ----------------------------
CREATE TABLE if not exists friend (
    `user_id` int,
    `friend_id` int
 );

-- ----------------------------
-- Table structure for comment
-- ----------------------------
CREATE TABLE if not exists comment (
    `user_id` int,
    `video_id` int,
    `comment_text` varchar(255)
)
;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, 1, '哈哈');


-- ----------------------------
-- Table structure for chat_messages
-- ----------------------------
CREATE TABLE if not exists chat_messages (
    `id` int auto_increment primary key ,
    `from_user_id` int NOT NULL,
    `to_user_id` int NOT NULL,
    `content` text NOT NULL,
    `created_at` timestamp(6)
);

-- ----------------------------
-- Records of chat_messages
-- ----------------------------
INSERT INTO `chat_messages` VALUES (54, 1, 0, '你好', '2023-08-27 07:48:43.071132');
INSERT INTO `chat_messages` VALUES (55, 0, 1, '您好！我是百度研发的知识增强大语言模型，中文名是文心一言，英文名是ERNIE Bot。我能够与人对话互动，回答问题，协助创作，高效便捷地帮助人们获取信息、知识和灵感。', '2023-08-27 07:48:43.073317');
INSERT INTO `chat_messages` VALUES (79, 1, 0, '1+2', '2023-08-27 16:17:45.139407');
INSERT INTO `chat_messages` VALUES (80, 0, 1, '1+2=3。这是加法运算中的基本规则，即加数1和加数2相加，其结果为3。', '2023-08-27 16:17:45.141701');