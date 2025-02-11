drop database if exists bootcamp_db;
create database bootcamp_db;

use bootcamp_db;

DROP TABLE IF EXISTS `countries`;

create table `countries` (
	`id` int primary key auto_increment,
    `country_name` varchar(255) UNIQUE NOT NULL 
);

DROP TABLE IF EXISTS `provinces`;
create table `provinces` (
	`id` int primary key auto_increment,
    `province_name` varchar(255) NOT NULL,
	`country_id` int,
    FOREIGN KEY (`country_id`) REFERENCES `countries`(`id`)
);

DROP TABLE IF EXISTS `localities`;
create table `localities` (
	`id` int primary key auto_increment,
    `locality_name` varchar(255) NOT NULL,
	`province_id` int,
	FOREIGN KEY (`province_id`) REFERENCES `provinces`(`id`)
    
);

DROP TABLE IF EXISTS `order_status`;
create table `order_status` (
	`id` int primary key auto_increment,
    `description` varchar(255) UNIQUE NOT NULL  
);

DROP TABLE IF EXISTS `buyers`;
create table `buyers` (
	`id` int primary key auto_increment,
    `id_card_number` varchar(255) UNIQUE NOT NULL,
	`first_name` varchar(255) NOT NULL,
	`last_name` varchar(255) NOT NULL
);

DROP TABLE IF EXISTS `sellers`;
create table `sellers` (
	`id` int primary key auto_increment,
    `cid` varchar(255) UNIQUE NOT NULL,
	`company_name` varchar(255) NOT NULL,
	`address` varchar(255) NOT NULL,
    `telephone` varchar(255) NOT NULL,
	`locality_id` int,
	FOREIGN KEY (`locality_id`) REFERENCES `localities`(`id`)
);

DROP TABLE IF EXISTS `carriers`;
create table `carriers` (
	`id` int primary key auto_increment,
    `cid` varchar(255) UNIQUE NOT NULL,
	`company_name` varchar(255) NOT NULL,
	`address` varchar(255) NOT NULL,
    `telephone` varchar(255) NOT NULL,
	`locality_id` int,
	FOREIGN KEY (`locality_id`) REFERENCES `localities`(`id`)
);

DROP TABLE IF EXISTS `warehouses`;
create table `warehouses` (
	`id` int primary key auto_increment,
    `warehouse_code` varchar(255) UNIQUE NOT NULL,
	`address` varchar(255) NOT NULL,
    `telephone` varchar(255) NOT NULL,
	`locality_id` int,
	FOREIGN KEY (`locality_id`) REFERENCES `localities`(`id`)
);

DROP TABLE IF EXISTS `products_types`;
create table `products_types` (
	`id` int primary key auto_increment,
    `description` varchar(255) UNIQUE NOT NULL  
);

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
   `id` INT PRIMARY KEY auto_increment,
   `description` VARCHAR(255),
   `expiration_rate` DECIMAL(19,2) not null ,
   `freezing_rate` DECIMAL(19,2) not null,
   `height` DECIMAL(19,2) not null ,
   `length` DECIMAL(19,2) not null ,
   `net_weight` DECIMAL(19,2) not null ,
   `product_code` VARCHAR(255) unique,
   `recommended_freezing_temperature` DECIMAL(19,2),
   `width` DECIMAL(19,2) not null ,
   `product_type_id` INT,
   `seller_id` INT,
   FOREIGN KEY (`product_type_id`) REFERENCES products_types(`id`),
   FOREIGN KEY (`seller_id`) REFERENCES sellers(`id`)
);

DROP TABLE IF EXISTS `sections`;
create table `sections` (
	`id` int primary key auto_increment,
    `section_number` varchar(255) UNIQUE NOT NULL,
	`current_capacity` int NOT NULL,
    `current_temperature`decimal(19,2),
    `maximum_capacity` int,
	`minimum_capacity` int,
    `minimum_temperature` decimal(19,2),
    `product_type_id` int,
    `warehouse_id` int,
	FOREIGN KEY (`product_type_id`) REFERENCES products_types(`id`),
    FOREIGN KEY (`warehouse_id`) REFERENCES warehouses(`id`)
);

DROP TABLE IF EXISTS `employees`;
create table `employees` (
	`id` int primary key auto_increment,
    `id_card_number` varchar(255) UNIQUE NOT NULL,
    `first_name` varchar(255) NOT NULL,
	`last_name` varchar(255) NOT NULL,
    `warehouse_id` int,
    FOREIGN KEY (`warehouse_id`) REFERENCES warehouses(`id`)
);

DROP TABLE IF EXISTS `purchase_orders`;
create table `purchase_orders` (
	`id` int primary key auto_increment,
    `order_number` varchar(255) UNIQUE NOT NULL,
    `order_date` datetime(6) NOT NULL,
	`tracking_code` varchar(255) NOT NULL,
    `buyer_id` int,
    `carrier_id` int,
    `order_status_id` int,
    `warehouse_id` int,
    FOREIGN KEY (`buyer_id`) REFERENCES buyers(`id`),
    FOREIGN KEY (`carrier_id`) REFERENCES carriers(`id`),
    FOREIGN KEY (`order_status_id`) REFERENCES order_status(`id`),
    FOREIGN KEY (`warehouse_id`) REFERENCES warehouses(`id`)
);
DROP TABLE IF EXISTS `product_batches`;
CREATE TABLE product_batches (
   `id` INT PRIMARY KEY auto_increment,
   `batch_number` VARCHAR(255) unique not null,
   `current_quantity` INT not null,
   `current_temperature` DECIMAL(19,2) not null,
   `due_date` DATETIME(6) not null,
   `initial_quantity` INT not null,
   `manufacturing_date` DATETIME(6) not null,
   `manufacturing_hour` DATETIME(6) not null,
   `minimum_temperature` DECIMAL(19,2) not null,
   `product_id` INT,
   `section_id` INT,
   FOREIGN KEY (`product_id`) REFERENCES products(`id`),
   FOREIGN KEY (`section_id`) REFERENCES sections(`id`)
);

DROP TABLE IF EXISTS `inbound_orders`;
CREATE TABLE inbound_orders (
   `id` INT PRIMARY KEY auto_increment,
   `order_date` DATETIME(6),
   `order_number` VARCHAR(255),
	`employee_id` int,
	`product_batch_id` int,
   `warehouses_id` int,
   FOREIGN KEY (`employee_id`) REFERENCES employees(`id`),
   FOREIGN KEY (`product_batch_id`) REFERENCES product_batches(`id`),
   FOREIGN KEY (`warehouses_id`) REFERENCES warehouses(`id`)
);



DROP TABLE IF EXISTS `product_records`;
CREATE TABLE `product_records` (
   `id` INT PRIMARY KEY auto_increment,
   `last_update_date` DATETIME(6) not null,
   `purchase_price` DECIMAL(19,2) not null,
   `sale_price` DECIMAL(19,2) not null,
   `product_id` INT,
   FOREIGN KEY (`product_id`) REFERENCES products(`id`)
);

DROP TABLE IF EXISTS `order_details`;
CREATE TABLE `order_details` (
   `id` INT PRIMARY KEY auto_increment,
   `clean_liness_status` varchar(255),
   `quantity` INT,
	`temperature` decimal(19,2),
	`product_record_id` INT,
   `purchase_order_id` INT,
   FOREIGN KEY (`product_record_id`) REFERENCES product_records(`id`),
	FOREIGN KEY (`purchase_order_id`) REFERENCES purchase_orders(`id`)

);
