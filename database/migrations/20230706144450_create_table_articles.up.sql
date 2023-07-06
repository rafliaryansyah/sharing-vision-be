CREATE TABLE IF NOT EXISTS articles
(
    id           INT AUTO_INCREMENT PRIMARY KEY,
    title        VARCHAR(255)                        not null,
    slug         VARCHAR(300)                        not null,
    content      TEXT                                not null,
    category     VARCHAR(255)                        not null,
    created_date TIMESTAMP                           not null,
    updated_date TIMESTAMP                           not null,
    status       ENUM ('Publish', 'Draft', 'Thrash') null,
    constraint slug
        unique (slug)
)ENGINE = InnoDB;

