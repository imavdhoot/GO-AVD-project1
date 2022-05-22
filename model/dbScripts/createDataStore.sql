DROP DATABASE IF EXISTS goPaceDb;
CREATE DATABASE goPaceDb;


DROP TABLE IF EXISTS `merchants`;
CREATE TABLE `merchants` (
    `id` VARCHAR(64) NOT NULL,
    `name` VARCHAR(64) NOT NULL,
    `address` VARCHAR(256),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idxName` (`name`)
) ENGINE=InnoDB;

INSERT INTO merchants (`id`, `name`, `address`)
VALUES
('merc1', 'merchant_one', 'Bedok, 083664'),
('merc2', 'merchant_two', 'Outram, 083664');



DROP TABLE IF EXISTS `members`;
CREATE TABLE `members` (
    `id` INT(10) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL,
    `email` VARCHAR(64) NOT NULL,
    `merchant_id` VARCHAR(64) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idxEmail` (`email`),
    INDEX `idxMerchant` (`merchant_id`),
    FOREIGN KEY (`merchant_id`) REFERENCES merchants (`id`)
) ENGINE=InnoDB;

INSERT INTO members (`id`, `name`, `email`, `merchant_id`)
VALUES
    (1, 'avd1', 'avd1@gmail.com', 'merc1'),
    (2, 'avd2', 'avd2@gmail.com', 'merc1'),
    (3, 'avd3', 'avd3@gmail.com', 'merc1'),
    (4, 'avd4', 'avd4@gmail.com', 'merc1'),
    (5, 'avd5', 'avd5@gmail.com', 'merc1'),
    (6, 'avd6', 'avd6@gmail.com', 'merc1'),
    (7, 'avd7', 'avd7@gmail.com', 'merc1'),
    (8, 'avd8', 'avd8@gmail.com', 'merc1'),
    (9, 'avd9', 'avd9@gmail.com', 'merc1'),
    (10, 'avd10', 'avd10@gmail.com', 'merc1'),
    (11, 'avd11', 'avd11@gmail.com', 'merc2');