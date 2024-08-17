-- Modify "thread_comments" table
ALTER TABLE `thread_comments` DROP FOREIGN KEY `thread_comments_threads_comments`;
-- Modify "thread_comments" table
ALTER TABLE `thread_comments` ADD CONSTRAINT `thread_comments_threads_comments` FOREIGN KEY (`thread_id`) REFERENCES `threads` (`id`) ON DELETE CASCADE;
