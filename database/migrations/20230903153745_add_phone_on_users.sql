-- migrate:up
ALTER TABLE `users` ADD `phone` CHAR(36) NOT NULL;

-- migrate:down
ALTER TABLE `users`` DROP COLUMN `phone`;
