CREATE TABLE `users`
(
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    name       varchar(255)        not null,
    email      varchar(255) Unique not null,
    password   varchar(255)        not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp default null

)