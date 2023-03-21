CREATE TABLE `user_video_like`  (
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  `is_like` int NOT NULL,
  PRIMARY KEY (`user_id`, `video_id`) USING BTREE,
  INDEX `video_id`(`video_id` ASC) USING BTREE,
  CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `video_id` FOREIGN KEY (`video_id`) REFERENCES `video` (`video_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;