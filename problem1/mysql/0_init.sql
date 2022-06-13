CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `name` varchar(64) DEFAULT '' NOT NULL,
  PRIMARY KEY (`id`)
);
-- user1 user2
CREATE TABLE `friend_link` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user1_id` int(11) NOT NULL,
  `user2_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
);
-- user1 user2 block
CREATE TABLE `block_list` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user1_id` int(11) NOT NULL,
  `user2_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `users` VALUES
  (1, 1, 'yamada'), (2, 2, 'tanaka'), (3, 3, 'katou'),
  (4, 4, 'satou'), (5, 5, 'kobayashi'), (6, 6, 'suzuki'), (7, 7, 'takahashi'), (8, 8, 'itou'), (9, 9, 'watanabe'), (10, 10, 'yamamoto'), (11, 11, 'yoshida'), (12, 12, 'yamada'), (13, 13, 'sasaki'), (14, 14, 'yamaguchi'), (15, 15, 'matsumoto'), (16, 16, 'inoue'), (17, 17, 'kimura');

INSERT INTO `friend_link` VALUES
  (1, 1, 2), (2, 1, 3), (3, 1, 4), (4, 1, 5), (5, 1, 6),(6, 1, 7), (7, 1, 8), (8, 1, 9), (9, 1, 10), (10, 1, 11), (11, 1, 12), (12, 1, 13), (13, 1, 14), (14, 1, 14), (15, 2, 3), (16, 3, 5), (17, 8, 1), (18, 9, 13), (19, 3, 9), (20, 10, 11), (21, 7, 8);

INSERT INTO `block_list` VALUES
  (1, 2, 4), (2, 1, 3);
-- INSERT INTO `users` (`user_id`, `name`) VALUES (1, 'yamda')
-- INSERT INTO `users` (`user_id`, `name`) VALUES (2, 'tanaka')

