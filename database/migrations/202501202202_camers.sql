-- +goose Up
CREATE TABLE `cameras`
(
    `uuid`      CHAR(36)     NOT NULL,
    `title`     VARCHAR(200) NOT NULL,
    `base_url`  VARCHAR(250) NOT NULL,
    `user`      VARCHAR(250) NOT NULL,
    `password`  VARCHAR(250) NOT NULL,
    `last_connection_status` enum('not_connected_yet', 'online', 'offline', 'invalid_auth') NOT NULL,
    `status`    enum('active', 'inactive') NOT NULL,
    PRIMARY KEY (`uuid`)
);

CREATE TABLE `camera_files`
(
    `uuid`          CHAR(36)     NOT NULL,
    `camera_uuid`   CHAR(36)     NOT NULL,
    `path`          VARCHAR(255) NOT NULL,
    PRIMARY KEY (`uuid`)
);

ALTER TABLE `dahua`.`camera_files`
    ADD INDEX `fk_camera_files_1_idx` (`camera_uuid` ASC) VISIBLE;

ALTER TABLE `dahua`.`camera_files`
    ADD CONSTRAINT `fk_camera_files_1`
        FOREIGN KEY (`camera_uuid`)
            REFERENCES `dahua`.`cameras` (`uuid`)
            ON DELETE CASCADE
            ON UPDATE CASCADE;


-- +goose Down
DROP TABLE `camera_files`;
DROP TABLE `cameras`;

