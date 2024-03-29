DROP TABLE `users` IF EXISTS `users`

CREATE TABLE `users` (
  `user_id` bigint COLLATE utf8mb4_bin AUTO_INCREMENT NOT NULL COMMENT 'ユーザID',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザー名',
  `email` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'メールアドレスを暗号化した文字列',
  `birth_date` date COLLATE utf8mb4_bin NOT NULL COMMENT '生年月日',
  `password` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'パスワード',
  PRIMARY KEY (`user_id`)
  UNIQUE KEY `users_name_uindex` (`name`)
  UNIQUE KEY `users_email_uindex` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;