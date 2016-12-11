SET foreign_key_checks = 1;
SET time_zone = '+05:30';

CREATE DATABASE IF NOT EXISTS atmdis DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

USE atmdis;

CREATE TABLE Atms (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    location VARCHAR(25) NOT NULL,
    city VARCHAR(25) NOT NULL,
    pincode VARCHAR(7) NOT NULL,
    bank VARCHAR(25) NOT NULL,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO `Atms` (`id`, `location`, `bank`, `city` ,`pincode`) VALUES
(1, 'Raheja Arcade, Koramangla', 'Citibank', 'Bangalore','560095'),
(2, '80Feet Road, Koramangla', 'Axis', 'Bangalore','560034');

