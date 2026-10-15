CREATE DATABASE IF NOT EXISTS `store_db` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `store_db`;

SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `orders`;

DROP TABLE IF EXISTS `products`;

DROP TABLE IF EXISTS `customers`;

DROP TABLE IF EXISTS `categories`;

SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `categories` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(3) NULL,
    `updated_at` DATETIME(3) NULL,
    `deleted_at` DATETIME(3) NULL,
    `name` VARCHAR(100) NOT NULL,
    `description` VARCHAR(255) NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_categories_name` (`name`),
    KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `products` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(3) NULL,
    `updated_at` DATETIME(3) NULL,
    `deleted_at` DATETIME(3) NULL,
    `name` VARCHAR(150) NOT NULL,
    `price` DECIMAL(12, 2) NOT NULL,
    `quantity` BIGINT NOT NULL,
    `category_id` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_products_deleted_at` (`deleted_at`),
    KEY `idx_products_name` (`name`),
    KEY `idx_products_category_id` (`category_id`),
    CONSTRAINT `fk_categories_products` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT `chk_products_price` CHECK (`price` >= 0),
    CONSTRAINT `chk_products_quantity` CHECK (`quantity` >= 0)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `customers` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(3) NULL,
    `updated_at` DATETIME(3) NULL,
    `deleted_at` DATETIME(3) NULL,
    `name` VARCHAR(150) NOT NULL,
    `email` VARCHAR(150) NOT NULL,
    `phone` VARCHAR(30) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_customers_email` (`email`),
    KEY `idx_customers_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `orders` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME(3) NULL,
    `updated_at` DATETIME(3) NULL,
    `deleted_at` DATETIME(3) NULL,
    `customer_id` BIGINT UNSIGNED NOT NULL,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `quantity` BIGINT NOT NULL,
    `total_price` DECIMAL(12, 2) NOT NULL,
    `status` VARCHAR(20) NOT NULL DEFAULT 'Pending',
    PRIMARY KEY (`id`),
    KEY `idx_orders_deleted_at` (`deleted_at`),
    KEY `idx_orders_customer_id` (`customer_id`),
    KEY `idx_orders_product_id` (`product_id`),
    KEY `idx_orders_status` (`status`),
    CONSTRAINT `fk_customers_orders` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT `fk_products_orders` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT `chk_orders_quantity` CHECK (`quantity` > 0),
    CONSTRAINT `chk_orders_total_price` CHECK (`total_price` >= 0),
    CONSTRAINT `chk_orders_status` CHECK (
        `status` IN (
            'Pending',
            'Completed',
            'Cancelled'
        )
    )
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

INSERT INTO
    `categories` (
        `id`,
        `created_at`,
        `updated_at`,
        `name`,
        `description`
    )
VALUES (
        1,
        NOW(3),
        NOW(3),
        'Workstations',
        'Portable and desktop work computers'
    ),
    (
        2,
        NOW(3),
        NOW(3),
        'Mobiles',
        'Smartphones and mobile devices'
    ),
    (
        3,
        NOW(3),
        NOW(3),
        'Peripherals',
        'Input and output accessories'
    ),
    (
        4,
        NOW(3),
        NOW(3),
        'Network',
        'Connectivity and routing equipment'
    );

INSERT INTO
    `products` (
        `id`,
        `created_at`,
        `updated_at`,
        `name`,
        `price`,
        `quantity`,
        `category_id`
    )
VALUES (
        1,
        NOW(3),
        NOW(3),
        'NovaBook Air S9',
        39900.00,
        8,
        1
    ),
    (
        2,
        NOW(3),
        NOW(3),
        'Aster Laptop Pro 15',
        22990.00,
        14,
        1
    ),
    (
        3,
        NOW(3),
        NOW(3),
        'Vertex Slim 3',
        18990.00,
        10,
        1
    ),
    (
        4,
        NOW(3),
        NOW(3),
        'Orbit Laptop 14',
        24990.00,
        6,
        1
    ),
    (
        5,
        NOW(3),
        NOW(3),
        'PulsePhone X1',
        29900.00,
        18,
        2
    ),
    (
        6,
        NOW(3),
        NOW(3),
        'Signal Galaxy Z9',
        27900.00,
        12,
        2
    ),
    (
        7,
        NOW(3),
        NOW(3),
        'North Pixel 9',
        26900.00,
        7,
        2
    ),
    (
        8,
        NOW(3),
        NOW(3),
        'Cursor Pro Mouse',
        690.00,
        43,
        3
    ),
    (
        9,
        NOW(3),
        NOW(3),
        'Echo Keypad K8',
        3290.00,
        21,
        3
    ),
    (
        10,
        NOW(3),
        NOW(3),
        'LinkHub Dock 8',
        1490.00,
        32,
        3
    ),
    (
        11,
        NOW(3),
        NOW(3),
        'Vision Display 27',
        7990.00,
        8,
        3
    ),
    (
        12,
        NOW(3),
        NOW(3),
        'RapidDrive SSD 1TB',
        3590.00,
        19,
        3
    ),
    (
        13,
        NOW(3),
        NOW(3),
        'AirWave Router AX3000',
        2590.00,
        13,
        4
    ),
    (
        14,
        NOW(3),
        NOW(3),
        'Metro Switch 8',
        1290.00,
        16,
        4
    ),
    (
        15,
        NOW(3),
        NOW(3),
        'CableFlow LAN 10m',
        290.00,
        48,
        4
    );

INSERT INTO
    `customers` (
        `id`,
        `created_at`,
        `updated_at`,
        `name`,
        `email`,
        `phone`
    )
VALUES (
        1,
        NOW(3),
        NOW(3),
        'Ari Ploy',
        'ari@example.com',
        '0811111111'
    ),
    (
        2,
        NOW(3),
        NOW(3),
        'Boonchai Niran',
        'boonchai@example.com',
        '0822222222'
    ),
    (
        3,
        NOW(3),
        NOW(3),
        'Chalerm Sai',
        'chalerm@example.com',
        '0833333333'
    ),
    (
        4,
        NOW(3),
        NOW(3),
        'Dara Wut',
        'dara@example.com',
        '0844444444'
    ),
    (
        5,
        NOW(3),
        NOW(3),
        'Ekkarat Sorn',
        'ekkarat@example.com',
        '0855555555'
    ),
    (
        6,
        NOW(3),
        NOW(3),
        'Fah Suda',
        'fah@example.com',
        '0866666666'
    ),
    (
        7,
        NOW(3),
        NOW(3),
        'Kittisak Nuan',
        'kittisak@example.com',
        '0877777777'
    ),
    (
        8,
        NOW(3),
        NOW(3),
        'Lada Chai',
        'lada@example.com',
        '0888888888'
    );

INSERT INTO
    `orders` (
        `id`,
        `created_at`,
        `updated_at`,
        `customer_id`,
        `product_id`,
        `quantity`,
        `total_price`,
        `status`
    )
VALUES (
        1,
        '2026-09-01 09:15:00.000',
        '2026-09-01 09:15:00.000',
        1,
        1,
        1,
        39900.00,
        'Completed'
    ),
    (
        2,
        '2026-09-02 10:20:00.000',
        '2026-09-02 10:20:00.000',
        2,
        5,
        1,
        29900.00,
        'Completed'
    ),
    (
        3,
        '2026-09-03 11:00:00.000',
        '2026-09-03 11:00:00.000',
        3,
        8,
        2,
        1380.00,
        'Completed'
    ),
    (
        4,
        '2026-09-04 13:45:00.000',
        '2026-09-04 13:45:00.000',
        4,
        9,
        1,
        3290.00,
        'Pending'
    ),
    (
        5,
        '2026-09-05 14:10:00.000',
        '2026-09-05 14:10:00.000',
        5,
        10,
        2,
        2980.00,
        'Completed'
    ),
    (
        6,
        '2026-09-06 15:30:00.000',
        '2026-09-06 15:30:00.000',
        6,
        13,
        1,
        2590.00,
        'Pending'
    ),
    (
        7,
        '2026-09-07 16:25:00.000',
        '2026-09-07 16:25:00.000',
        7,
        2,
        1,
        22990.00,
        'Cancelled'
    ),
    (
        8,
        '2026-09-08 17:05:00.000',
        '2026-09-08 17:05:00.000',
        8,
        15,
        2,
        580.00,
        'Completed'
    ),
    (
        9,
        '2026-09-09 09:40:00.000',
        '2026-09-09 09:40:00.000',
        1,
        12,
        1,
        3590.00,
        'Pending'
    ),
    (
        10,
        '2026-09-10 10:50:00.000',
        '2026-09-10 10:50:00.000',
        2,
        6,
        1,
        27900.00,
        'Completed'
    ),
    (
        11,
        '2026-09-11 12:35:00.000',
        '2026-09-11 12:35:00.000',
        3,
        14,
        1,
        1290.00,
        'Pending'
    ),
    (
        12,
        '2026-09-12 18:00:00.000',
        '2026-09-12 18:00:00.000',
        4,
        11,
        1,
        7990.00,
        'Completed'
    );