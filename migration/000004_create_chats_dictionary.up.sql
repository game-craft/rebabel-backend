CREATE TABLE `chats_dictionary` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `chats_dictionary_content` VARCHAR(255) NOT NULL,
    `chats_dictionary_status` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;