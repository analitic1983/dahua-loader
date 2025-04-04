-- +goose Up
CREATE TABLE `user_login_history` (
                                      `user_uuid` char(36) NOT NULL,
                                      `date` datetime NOT NULL,
                                      `ip` varchar(45) NOT NULL,
                                      PRIMARY KEY (`user_uuid`)
) ENGINE=InnoDB;


-- +goose Down
DROP TABLE `user_login_history`
