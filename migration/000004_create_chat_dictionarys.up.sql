CREATE TABLE `chat_dictionarys` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `chat_dictionarys_content` VARCHAR(255) NOT NULL,
    `chat_dictionarys_status` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;