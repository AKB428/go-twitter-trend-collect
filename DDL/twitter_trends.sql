CREATE TABLE `twitter_trends` (
  `id` int NOT NULL AUTO_INCREMENT,
  `rank` int DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_twitter_trends_on_name` (`name`),
  KEY `index_twitter_trends_on_created_at` (`created_at`),
  KEY `index_twitter_trend_on_rank` (`rank`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;