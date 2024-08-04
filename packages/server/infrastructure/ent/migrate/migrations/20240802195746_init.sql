-- Modify "thread_tags" table
ALTER TABLE `thread_tags` DROP FOREIGN KEY `thread_tags_tags_tag`;
-- Modify "thread_tags" table
ALTER TABLE `thread_tags` ADD CONSTRAINT `thread_tags_tags_tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE;
