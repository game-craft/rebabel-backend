CREATE TABLE `chat_classifications` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `worlds_id` INT NOT NULL,
    `chat_classifications_content` VARCHAR(255) NOT NULL,
    `chat_classifications_status` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

ALTER TABLE chat_classifications
ADD CONSTRAINT fk_chat_classifications_worlds
FOREIGN KEY (worlds_id)
REFERENCES worlds(id);