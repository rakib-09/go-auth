CREATE TABLE `companies`
(
    id            BIGINT PRIMARY KEY AUTO_INCREMENT,
    title         varchar(255) not null,
    owner_user_id BIGINT       not null,
    phone         varchar(20)  not null,
    address       varchar(255),
    created_at    timestamp default CURRENT_TIMESTAMP
)