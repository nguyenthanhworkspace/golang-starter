CREATE TABLE IF NOT EXISTS users(
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    email VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    first_name VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    last_name VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    PRIMARY KEY (id)
);