CREATE TABLE IF NOT EXISTS `users` (
    `id`         BIGINT UNSIGNED AUTO_INCREMENT,
    `name`       VARCHAR(255) NOT NULL,
    `body`       VARCHAR(999),
    `password`   VARCHAR(999),
    `is_forbid`  TINYINT   DEFAULT 0,
    `created_at` TIMESTAMP DEFAULT NULL,
    `updated_at` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX users_is_forbid_index (`is_forbid`),
    UNIQUE INDEX users_name_unique (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
