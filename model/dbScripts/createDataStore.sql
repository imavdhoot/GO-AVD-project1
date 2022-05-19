DROP DATABASE IF EXISTS goPaceDb;
CREATE DATABASE goPaceDb;


DROP TABLE IF EXISTS `merchant`;
CREATE TABLE `merchant` (
    `id` VARCHAR(64) NOT NULL,
    `name` VARCHAR(64) NOT NULL,
    `address` VARCHAR(256),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idxName` (`name`)
) ENGINE=InnoDB;

INSERT INTO merchant (`id`, `name`, `address`)
VALUES
('merc1', 'merchant_one', 'Bedok, 083664'),
('merc2', 'merchant_two', 'Outram, 083664');