CREATE TABLE `twitter_trends` (
  `id` int NOT NULL AUTO_INCREMENT,
  `rank` int DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_twitter_trends_on_name` (`name`),
  KEY `index_twitter_trends_on_created_at` (`created_at`),
  KEY `index_twitter_trend_on_rank` (`rank`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;