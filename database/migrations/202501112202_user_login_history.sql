-- +goose Up
CREATE TABLE `user_login_histories` (
                                      `uuid` char(36) NOT NULL,
                                      `user_uuid` char(36) NOT NULL,
                                      `date` datetime NOT NULL,
                                      `ip` varchar(45) NOT NULL,
                                      PRIMARY KEY (`uuid`)
) ENGINE=InnoDB;

-- +goose Down
DROP TABLE `user_login_histories`
