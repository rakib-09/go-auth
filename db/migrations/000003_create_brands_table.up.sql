CREATE TABLE `brands`
(
    id            BIGINT PRIMARY KEY AUTO_INCREMENT,
    title         varchar(255)  not null,
    company_id    BIGINT not null,
    phone         varchar(20)   not null,
    address       varchar(255),
    created_at    timestamp default CURRENT_TIMESTAMP,
    updated_at    timestamp default CURRENT_TIMESTAMP,
    deleted_at    timestamp default null
)