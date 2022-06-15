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

INSERT INTO `users` (`user_id`, `name`) VALUES
  (1, 'yamada'), (2, 'tanaka'), (3, 'katou'), (4, 'satou'), (5, 'kobayashi'), (6, 'suzuki'), (7, 'takahashi'), (8, 'itou'), (9, 'watanabe'), (10, 'yamamoto'), (11, 'yoshida'), (12, 'yamada'), (13, 'sasaki'), (14, 'yamaguchi'), (15, 'matsumoto'), (16, 'inoue'), (17, 'kimura');

INSERT INTO `friend_link` (`user1_id`, `user2_id`) VALUES
  (1, 2), (1, 3), (1, 4), (1, 5), (1, 6),(1, 7), (1, 8), (1, 9), (1, 10), (1, 11), (1, 12), (1, 13), (1, 14), (1, 14), (2, 3), (3, 5), (8, 1), (9, 13), (3, 9), (10, 11), (7, 8);

INSERT INTO `block_list` (`user1_id`, `user2_id`) VALUES
  (2, 4), (1, 3);
