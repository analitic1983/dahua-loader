-- +goose Up
DELETE FROM `goose_db_version`;
ALTER TABLE `goose_db_version` CHANGE COLUMN `tstamp` `tstamp` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ;
INSERT INTO `goose_db_version` (id,version_id,is_applied,tstamp) VALUES (1,0,1, '2025-01-01 10:20:10');

-- +goose Down

