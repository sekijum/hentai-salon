-- Create "ads" table
CREATE TABLE `ads` (`id` bigint NOT NULL AUTO_INCREMENT, `content` longtext NOT NULL, `is_active` bigint NOT NULL DEFAULT 1, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
