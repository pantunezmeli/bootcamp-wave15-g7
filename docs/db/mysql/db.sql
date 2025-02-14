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
   `order_number` VARCHAR(255) unique,
	`employee_id` int,
	`product_batch_id` int,
   `warehouse_id` int,
   FOREIGN KEY (`employee_id`) REFERENCES employees(`id`),
   FOREIGN KEY (`product_batch_id`) REFERENCES product_batches(`id`),
   FOREIGN KEY (`warehouse_id`) REFERENCES warehouses(`id`)
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

INSERT INTO `countries` (`country_name`) VALUES
('Argentina'),
('Brazil'),
('Chile'),
('Uruguay'),
('Paraguay'),
('Bolivia'),
('Peru'),
('Ecuador'),
('Colombia'),
('Venezuela'),
('Guyana'),
('Suriname'),
('French Guiana'),
('Mexico'),
('United States'),
('Canada'),
('Guatemala'),
('Honduras'),
('El Salvador'),
('Nicaragua'),
('Costa Rica'),
('Panama'),
('Cuba'),
('Dominican Republic'),
('Haiti'),
('Jamaica'),
('Trinidad and Tobago'),
('Bahamas'),
('Barbados'),
('Belize'),
('Saint Lucia'),
('Grenada'),
('Saint Vincent and the Grenadines'),
('Antigua and Barbuda'),
('Saint Kitts and Nevis'),
('United Kingdom'),
('France'),
('Germany'),
('Spain'),
('Italy'),
('Portugal'),
('Netherlands'),
('Belgium'),
('Switzerland'),
('Austria'),
('Denmark'),
('Sweden'),
('Norway'),
('Finland'),
('Iceland'),
('Ireland');


INSERT INTO `provinces` (`province_name`, `country_id`) VALUES
('Buenos Aires', 1),
('Córdoba', 1),
('Santa Fe', 1),
('Río de Janeiro', 2),
('São Paulo', 2),
('Minas Gerais', 2),
('Santiago', 3),
('Valparaíso', 3),
('Biobío', 3),
('Montevideo', 4),
('Canelones', 4),
('Maldonado', 4),
('Central', 5),
('Alto Paraná', 5),
('Itapúa', 5),
('La Paz', 6),
('Santa Cruz', 6),
('Cochabamba', 6),
('Lima', 7),
('Arequipa', 7),
('Cusco', 7),
('Pichincha', 8),
('Guayas', 8),
('Manabí', 8),
('Cundinamarca', 9),
('Antioquia', 9),
('Valle del Cauca', 9),
('Distrito Capital', 10),
('Miranda', 10),
('Zulia', 10),
('Demerara-Mahaica', 11),
('Essequibo Islands-West Demerara', 11),
('Pomeroon-Supenaam', 11),
('Paramaribo', 12),
('Wanica', 12),
('Nickerie', 12),
('Île-de-France', 13),
('Cayenne', 13),
('Saint-Laurent-du-Maroni', 13),
('Mexico City', 14),
('Jalisco', 14),
('Nuevo León', 14),
('California', 15),
('Texas', 15),
('Florida', 15),
('Ontario', 16),
('Quebec', 16),
('British Columbia', 16);

INSERT INTO `localities` (`locality_name`, `province_id`) VALUES
('La Plata', 1),
('Mar del Plata', 1),
('Bahía Blanca', 1),
('Córdoba Capital', 2),
('Villa María', 2),
('Río Cuarto', 2),
('Rosario', 3),
('Santa Fe Capital', 3),
('Rafaela', 3),
('Rio de Janeiro', 4),
('Niterói', 4),
('Campos dos Goytacazes', 4),
('São Paulo', 5),
('Campinas', 5),
('Santos', 5),
('Belo Horizonte', 6),
('Uberlândia', 6),
('Juiz de Fora', 6),
('Santiago de Chile', 7),
('Viña del Mar', 7),
('San Bernardo', 7),
('Valparaíso', 8),
('Quilpué', 8),
('Villa Alemana', 8),
('Concepción', 9),
('Talcahuano', 9),
('Chillán', 9),
('Montevideo', 10),
('Pando', 10),
('Las Piedras', 10),
('Canelones', 11),
('Santa Lucía', 11),
('San José de Mayo', 11),
('Maldonado', 12),
('Punta del Este', 12),
('San Carlos', 12),
('Asunción', 13),
('San Lorenzo', 13),
('Fernando de la Mora', 13),
('Ciudad del Este', 14),
('Hernandarias', 14),
('Presidente Franco', 14),
('Encarnación', 15),
('San Juan del Paraná', 15),
('Cambyretá', 15),
('La Paz', 16),
('El Alto', 16),
('Viacha', 16),
('Santa Cruz de la Sierra', 17),
('Montero', 17),
('Warnes', 17),
('Cochabamba', 18),
('Sacaba', 18),
('Quillacollo', 18),
('Lima', 19),
('Callao', 19),
('Chorrillos', 19),
('Arequipa', 20),
('Camaná', 20),
('Mollendo', 20),
('Cusco', 21),
('Urubamba', 21),
('Calca', 21),
('Quito', 22),
('Sangolquí', 22),
('Cayambe', 22),
('Guayaquil', 23),
('Daule', 23),
('Samborondón', 23),
('Manta', 24),
('Portoviejo', 24),
('Jipijapa', 24),
('Bogotá', 25),
('Soacha', 25),
('Facatativá', 25),
('Medellín', 26),
('Bello', 26),
('Envigado', 26),
('Cali', 27),
('Palmira', 27),
('Jamundí', 27),
('Caracas', 28),
('Los Teques', 28),
('Guarenas', 28),
('Maracaibo', 29),
('Cabimas', 29),
('Ciudad Ojeda', 29),
('Georgetown', 30),
('Linden', 30),
('Bartica', 30),
('Paramaribo', 31),
('Nieuw Nickerie', 31),
('Moengo', 31),
('Cayenne', 32),
('Saint-Laurent-du-Maroni', 32),
('Kourou', 32),
('Mexico City', 33),
('Ecatepec', 33),
('Naucalpan', 33);


INSERT INTO `order_status` (`description`) VALUES
('Pending'),
('Processing'),
('Shipped'),
('Delivered'),
('Cancelled'),
('Returned'),
('On Hold'),
('Completed'),
('Awaiting Payment'),
('Payment Received'),
('Payment Failed'),
('Refunded'),
('Out for Delivery'),
('Ready for Pickup'),
('Partially Shipped'),
('Backordered'),
('Awaiting Fulfillment'),
('Awaiting Shipment'),
('Awaiting Pickup'),
('In Transit'),
('Order Received'),
('Order Confirmed'),
('Order Packed'),
('Failed Delivery Attempt'),
('Reattempting Delivery'),
('Dispatched'),
('Pick Up Scheduled'),
('Pick Up Completed'),
('Awaiting Confirmation'),
('Order Modified'),
('Partially Completed'),
('Order Expired'),
('Pre-Order'),
('Custom Order Processing'),
('Subscription Active'),
('Subscription Cancelled'),
('Subscription Expired'),
('Hold for Verification'),
('Awaiting Approval'),
('Approved'),
('Rejected'),
('Awaiting Restock'),
('Restocked'),
('Limited Availability'),
('Special Handling Required'),
('Awaiting Supplier'),
('Supplier Confirmed'),
('In Production'),
('Production Completed'),
('Recalled'),
('Lost in Transit');

INSERT INTO `buyers` (`id_card_number`, `first_name`, `last_name`) VALUES
('ID10001', 'Juan', 'Pérez'),
('ID10002', 'María', 'Gómez'),
('ID10003', 'Carlos', 'Rodríguez'),
('ID10004', 'Ana', 'Fernández'),
('ID10005', 'Luis', 'González'),
('ID10006', 'Laura', 'Martínez'),
('ID10007', 'Pedro', 'López'),
('ID10008', 'Sofía', 'Díaz'),
('ID10009', 'Diego', 'Torres'),
('ID10010', 'Camila', 'Ramírez'),
('ID10011', 'Javier', 'Sánchez'),
('ID10012', 'Valentina', 'Moreno'),
('ID10013', 'Ricardo', 'Álvarez'),
('ID10014', 'Lucía', 'Molina'),
('ID10015', 'Fernando', 'Castro'),
('ID10016', 'Martina', 'Ortiz'),
('ID10017', 'Nicolás', 'Silva'),
('ID10018', 'Paula', 'Medina'),
('ID10019', 'Gabriel', 'Vega'),
('ID10020', 'Isabela', 'Rojas');


INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES
('CID1001', 'Distribuidora La Central', 'Av. Siempre Viva 123', '1122334455', 1),
('CID1002', 'Frutas del Sur', 'Calle Primavera 456', '1144778899', 2),
('CID1003', 'Verdulería El Mercado', 'Av. Libertad 789', '1133557799', 3),
('CID1004', 'Hortalizas Frescas', 'Calle Esperanza 101', '1122113344', 4),
('CID1005', 'Mayorista Las Delicias', 'Av. San Martín 202', '1155667788', 5),
('CID1006', 'Frutas y Más', 'Calle 9 de Julio 303', '1177889900', 6),
('CID1007', 'El Campo Natural', 'Av. Independencia 404', '1166448822', 7),
('CID1008', 'La Huerta Express', 'Calle Mitre 505', '1199112233', 8),
('CID1009', 'Distribuciones Frutales', 'Av. Córdoba 606', '1122331100', 9),
('CID1010', 'Verduras al Día', 'Calle Sarmiento 707', '1133774499', 10),
('CID1011', 'Frescos del Valle', 'Av. Belgrano 808', '1144667788', 11),
('CID1012', 'Mayorista AgroVerde', 'Calle Moreno 909', '1177223344', 12),
('CID1013', 'La Frutería Natural', 'Av. Rivadavia 111', '1188556677', 13),
('CID1014', 'El Rincón de la Huerta', 'Calle Alberdi 222', '1199445566', 14),
('CID1015', 'Verdulería Santa Fe', 'Av. Santa Fe 333', '1155112233', 15),
('CID1016', 'Distribuidora La Quinta', 'Calle Córdoba 444', '1166998811', 16),
('CID1017', 'Eco Frutas', 'Av. San Juan 555', '1177552299', 17),
('CID1018', 'Hortalizas del Norte', 'Calle Roca 666', '1199223344', 18),
('CID1019', 'Frescuras del Campo', 'Av. Tucumán 777', '1188112233', 19),
('CID1020', 'El Rey de las Frutas', 'Calle Mitre 888', '1133445566', 20);

INSERT INTO `carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES
('CAR1001', 'Transporte Rápido Express', 'Av. Siempre Viva 123', '1122334455', 1),
('CAR1002', 'Logística del Sur', 'Calle Primavera 456', '1144778899', 2),
('CAR1003', 'Fletes y Mudanzas Martínez', 'Av. Libertad 789', '1133557799', 3),
('CAR1004', 'Transportes La Rápida', 'Calle Esperanza 101', '1122113344', 4),
('CAR1005', 'Carga Segura', 'Av. San Martín 202', '1155667788', 5),
('CAR1006', 'Camioneros del Oeste', 'Calle 9 de Julio 303', '1177889900', 6),
('CAR1007', 'MoviCargo Express', 'Av. Independencia 404', '1166448822', 7),
('CAR1008', 'Transporte Nacional', 'Calle Mitre 505', '1199112233', 8),
('CAR1009', 'Distribuciones en Ruta', 'Av. Córdoba 606', '1122331100', 9),
('CAR1010', 'Camiones del Norte', 'Calle Sarmiento 707', '1133774499', 10),
('CAR1011', 'Fletes Urbanos', 'Av. Belgrano 808', '1144667788', 11),
('CAR1012', 'Carga y Logística Global', 'Calle Moreno 909', '1177223344', 12),
('CAR1013', 'Transportes Patagónicos', 'Av. Rivadavia 111', '1188556677', 13),
('CAR1014', 'Mudanzas y Fletes Express', 'Calle Alberdi 222', '1199445566', 14),
('CAR1015', 'Logística Santa Fe', 'Av. Santa Fe 333', '1155112233', 15),
('CAR1016', 'Distribuidora de Cargas', 'Calle Córdoba 444', '1166998811', 16),
('CAR1017', 'Camiones al Instante', 'Av. San Juan 555', '1177552299', 17),
('CAR1018', 'Fletes del Norte', 'Calle Roca 666', '1199223344', 18),
('CAR1019', 'Logística Express Global', 'Av. Tucumán 777', '1188112233', 19),
('CAR1020', 'Transportes y Mudanzas Mitre', 'Calle Mitre 888', '1133445566', 20);

INSERT INTO `warehouses` (`warehouse_code`, `address`, `telephone`, `locality_id`) VALUES
('WH001', 'Av. Siempre Viva 123', '1122334455', 1),
('WH002', 'Calle Primavera 456', '1144778899', 2),
('WH003', 'Av. Libertad 789', '1133557799', 3),
('WH004', 'Calle Esperanza 101', '1122113344', 4),
('WH005', 'Av. San Martín 202', '1155667788', 5),
('WH006', 'Calle 9 de Julio 303', '1177889900', 6),
('WH007', 'Av. Independencia 404', '1166448822', 7),
('WH008', 'Calle Mitre 505', '1199112233', 8),
('WH009', 'Av. Córdoba 606', '1122331100', 9),
('WH010', 'Calle Sarmiento 707', '1133774499', 10),
('WH011', 'Av. Belgrano 808', '1144667788', 11),
('WH012', 'Calle Moreno 909', '1177223344', 12),
('WH013', 'Av. Rivadavia 111', '1188556677', 13),
('WH014', 'Calle Alberdi 222', '1199445566', 14),
('WH015', 'Av. Santa Fe 333', '1155112233', 15),
('WH016', 'Calle Córdoba 444', '1166998811', 16),
('WH017', 'Av. San Juan 555', '1177552299', 17),
('WH018', 'Calle Roca 666', '1199223344', 18),
('WH019', 'Av. Tucumán 777', '1188112233', 19),
('WH020', 'Calle Mitre 888', '1133445566', 20);

INSERT INTO `products_types` (`description`) VALUES
('Frutas'),
('Verduras'),
('Lácteos'),
('Carnes'),
('Pescados y Mariscos'),
('Panadería'),
('Bebidas'),
('Cereales y Legumbres'),
('Condimentos y Especias'),
('Dulces y Postres'),
('Aceites y Grasas'),
('Comida enlatada'),
('Productos congelados'),
('Snacks'),
('Productos de panadería industrial'),
('Huevos'),
('Productos sin gluten'),
('Alimentos orgánicos'),
('Alimentos veganos'),
('Alimentos para bebés');

INSERT INTO `products` (`description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id`) VALUES
('Manzanas Rojas', 10.50, 0.00, 12.50, 15.00, 1.20, 'P001', -1.00, 10.00, 1, 1),
('Bananas', 8.30, 0.00, 20.00, 25.00, 1.80, 'P002', -2.00, 15.00, 1, 2),
('Leche Entera', 7.00, 5.00, 25.00, 10.00, 2.00, 'P003', 4.00, 10.00, 3, 3),
('Queso Fresco', 15.00, 3.00, 10.00, 10.00, 0.50, 'P004', 2.00, 8.00, 3, 4),
('Carne de Res', 20.00, 10.00, 30.00, 40.00, 5.00, 'P005', -18.00, 25.00, 4, 5),
('Salmón Fresco', 12.00, 8.00, 20.00, 30.00, 3.50, 'P006', -20.00, 15.00, 5, 6),
('Pan Integral', 5.00, 0.00, 10.00, 20.00, 1.00, 'P007', 0.00, 10.00, 6, 7),
('Jugo de Naranja', 30.00, 0.00, 15.00, 10.00, 1.50, 'P008', 2.00, 10.00, 7, 8),
('Arroz Blanco', 360.00, 0.00, 20.00, 10.00, 2.00, 'P009', 0.00, 15.00, 8, 9),
('Pimienta Negra', 540.00, 0.00, 5.00, 5.00, 0.10, 'P010', 0.00, 5.00, 9, 10),
('Dulce de Leche', 60.00, 0.00, 12.00, 15.00, 1.00, 'P011', 2.00, 10.00, 10, 11),
('Aceite de Oliva', 180.00, 0.00, 25.00, 8.00, 0.75, 'P012', 0.00, 8.00, 11, 12),
('Atún enlatado', 720.00, 0.00, 10.00, 10.00, 0.25, 'P013', 0.00, 10.00, 12, 13),
('Papas Congeladas', 12.00, 15.00, 30.00, 20.00, 2.50, 'P014', -18.00, 15.00, 13, 14),
('Chips de Papa', 120.00, 0.00, 15.00, 10.00, 0.50, 'P015', 0.00, 8.00, 14, 15),
('Galletas Dulces', 180.00, 0.00, 10.00, 15.00, 0.75, 'P016', 0.00, 10.00, 15, 16),
('Huevos Orgánicos', 30.00, 0.00, 8.00, 12.00, 0.60, 'P017', 5.00, 10.00, 16, 17),
('Harina de Trigo', 360.00, 0.00, 20.00, 15.00, 2.00, 'P018', 0.00, 15.00, 17, 18),
('Tofu', 15.00, 5.00, 10.00, 10.00, 0.30, 'P019', 2.00, 10.00, 18, 19),
('Papilla para Bebés', 90.00, 0.00, 12.00, 8.00, 0.50, 'P020', 4.00, 8.00, 19, 20);

INSERT INTO `sections` (`section_number`, `current_capacity`, `current_temperature`, `maximum_capacity`, `minimum_capacity`, `minimum_temperature`, `product_type_id`, `warehouse_id`) VALUES
('S001', 100, 4.00, 150, 50, 0.00, 1, 1),
('S002', 200, -2.00, 300, 100, -5.00, 2, 2),
('S003', 150, 2.00, 200, 80, -1.00, 3, 3),
('S004', 180, -10.00, 250, 90, -15.00, 4, 4),
('S005', 120, -18.00, 200, 60, -20.00, 5, 5),
('S006', 250, 1.00, 350, 150, 0.00, 6, 6),
('S007', 300, -5.00, 400, 200, -10.00, 7, 7),
('S008', 400, 3.00, 500, 250, -2.00, 8, 8),
('S009', 350, -12.00, 450, 200, -18.00, 9, 9),
('S010', 100, 5.00, 150, 50, 0.00, 10, 10),
('S011', 220, -8.00, 300, 120, -15.00, 11, 11),
('S012', 280, 0.00, 350, 170, -2.00, 12, 12),
('S013', 90, 6.00, 130, 40, 1.00, 13, 13),
('S014', 180, -4.00, 230, 80, -10.00, 14, 14),
('S015', 260, -20.00, 320, 150, -25.00, 15, 15),
('S016', 140, -1.00, 200, 70, -5.00, 16, 16),
('S017', 320, -3.00, 400, 180, -8.00, 17, 17),
('S018', 370, 2.00, 450, 250, -1.00, 18, 18),
('S019', 90, -6.00, 140, 50, -12.00, 19, 19),
('S020', 210, 0.00, 280, 110, -3.00, 20, 20);


INSERT INTO `employees` (`id_card_number`, `first_name`, `last_name`, `warehouse_id`) VALUES
('EMP1001', 'Juan', 'Pérez', 1),
('EMP1002', 'María', 'Gómez', 2),
('EMP1003', 'Carlos', 'Rodríguez', 3),
('EMP1004', 'Ana', 'Fernández', 4),
('EMP1005', 'Luis', 'González', 5),
('EMP1006', 'Laura', 'Martínez', 6),
('EMP1007', 'Pedro', 'López', 7),
('EMP1008', 'Sofía', 'Díaz', 8),
('EMP1009', 'Diego', 'Torres', 9),
('EMP1010', 'Camila', 'Ramírez', 10),
('EMP1011', 'Javier', 'Sánchez', 11),
('EMP1012', 'Valentina', 'Moreno', 12),
('EMP1013', 'Ricardo', 'Álvarez', 13),
('EMP1014', 'Lucía', 'Molina', 14),
('EMP1015', 'Fernando', 'Castro', 15),
('EMP1016', 'Martina', 'Ortiz', 16),
('EMP1017', 'Nicolás', 'Silva', 17),
('EMP1018', 'Paula', 'Medina', 18),
('EMP1019', 'Gabriel', 'Vega', 19),
('EMP1020', 'Isabela', 'Rojas', 20);

INSERT INTO `purchase_orders` (`order_number`, `order_date`, `tracking_code`, `buyer_id`, `carrier_id`, `order_status_id`, `warehouse_id`) VALUES
('PO1001', '2024-02-01 10:15:30.000000', 'TRK001', 1, 1, 1, 1),
('PO1002', '2024-02-02 11:30:45.000000', 'TRK002', 2, 2, 2, 2),
('PO1003', '2024-02-03 12:45:20.000000', 'TRK003', 3, 3, 3, 3),
('PO1004', '2024-02-04 14:00:10.000000', 'TRK004', 4, 4, 4, 4),
('PO1005', '2024-02-05 15:20:55.000000', 'TRK005', 5, 5, 5, 5),
('PO1006', '2024-02-06 16:35:40.000000', 'TRK006', 6, 6, 6, 6),
('PO1007', '2024-02-07 17:50:25.000000', 'TRK007', 7, 7, 7, 7),
('PO1008', '2024-02-08 18:05:10.000000', 'TRK008', 8, 8, 8, 8),
('PO1009', '2024-02-09 19:15:50.000000', 'TRK009', 9, 9, 9, 9),
('PO1010', '2024-02-10 20:30:35.000000', 'TRK010', 10, 10, 10, 10),
('PO1011', '2024-02-11 21:45:20.000000', 'TRK011', 11, 11, 11, 11),
('PO1012', '2024-02-12 22:50:05.000000', 'TRK012', 12, 12, 12, 12),
('PO1013', '2024-02-13 23:05:50.000000', 'TRK013', 13, 13, 13, 13),
('PO1014', '2024-02-14 08:20:35.000000', 'TRK014', 14, 14, 14, 14),
('PO1015', '2024-02-15 09:30:25.000000', 'TRK015', 15, 15, 15, 15),
('PO1016', '2024-02-16 10:40:15.000000', 'TRK016', 16, 16, 16, 16),
('PO1017', '2024-02-17 11:50:05.000000', 'TRK017', 17, 17, 17, 17),
('PO1018', '2024-02-18 12:05:55.000000', 'TRK018', 18, 18, 18, 18),
('PO1019', '2024-02-19 13:15:45.000000', 'TRK019', 19, 19, 19, 19),
('PO1020', '2024-02-20 14:25:35.000000', 'TRK020', 20, 20, 20, 20);


INSERT INTO `product_batches` (`batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES
('BATCH001', 80, -2.00, '2025-03-10 23:59:59.000000', 100, '2024-02-01 08:00:00.000000', '2024-02-01 08:00:00.000000', -5.00, 1, 1),
('BATCH002', 120, -3.00, '2025-03-15 23:59:59.000000', 150, '2024-02-02 09:15:00.000000', '2024-02-02 09:15:00.000000', -6.00, 2, 2),
('BATCH003', 90, 1.00, '2025-04-01 23:59:59.000000', 110, '2024-02-03 10:30:00.000000', '2024-02-03 10:30:00.000000', -2.00, 3, 3),
('BATCH004', 75, -10.00, '2025-04-10 23:59:59.000000', 90, '2024-02-04 11:45:00.000000', '2024-02-04 11:45:00.000000', -12.00, 4, 4),
('BATCH005', 200, -18.00, '2025-05-01 23:59:59.000000', 250, '2024-02-05 13:00:00.000000', '2024-02-05 13:00:00.000000', -20.00, 5, 5),
('BATCH006', 50, 0.00, '2025-05-15 23:59:59.000000', 60, '2024-02-06 14:15:00.000000', '2024-02-06 14:15:00.000000', -1.00, 6, 6),
('BATCH007', 130, -5.00, '2025-06-01 23:59:59.000000', 150, '2024-02-07 15:30:00.000000', '2024-02-07 15:30:00.000000', -8.00, 7, 7),
('BATCH008', 180, 2.00, '2025-06-15 23:59:59.000000', 200, '2024-02-08 16:45:00.000000', '2024-02-08 16:45:00.000000', 0.00, 8, 8),
('BATCH009', 75, -12.00, '2025-07-01 23:59:59.000000', 90, '2024-02-09 18:00:00.000000', '2024-02-09 18:00:00.000000', -15.00, 9, 9),
('BATCH010', 200, 3.00, '2025-07-15 23:59:59.000000', 220, '2024-02-10 19:15:00.000000', '2024-02-10 19:15:00.000000', 0.00, 10, 10),
('BATCH011', 95, -8.00, '2025-08-01 23:59:59.000000', 110, '2024-02-11 20:30:00.000000', '2024-02-11 20:30:00.000000', -10.00, 11, 11),
('BATCH012', 120, 0.00, '2025-08-15 23:59:59.000000', 140, '2024-02-12 21:45:00.000000', '2024-02-12 21:45:00.000000', -1.00, 12, 12),
('BATCH013', 110, 6.00, '2025-09-01 23:59:59.000000', 130, '2024-02-13 08:00:00.000000', '2024-02-13 08:00:00.000000', 2.00, 13, 13),
('BATCH014', 150, -4.00, '2025-09-15 23:59:59.000000', 170, '2024-02-14 09:15:00.000000', '2024-02-14 09:15:00.000000', -6.00, 14, 14),
('BATCH015', 180, -20.00, '2025-10-01 23:59:59.000000', 200, '2024-02-15 10:30:00.000000', '2024-02-15 10:30:00.000000', -22.00, 15, 15),
('BATCH016', 140, -1.00, '2025-10-15 23:59:59.000000', 160, '2024-02-16 11:45:00.000000', '2024-02-16 11:45:00.000000', -3.00, 16, 16),
('BATCH017', 190, -3.00, '2025-11-01 23:59:59.000000', 220, '2024-02-17 13:00:00.000000', '2024-02-17 13:00:00.000000', -5.00, 17, 17),
('BATCH018', 250, 2.00, '2025-11-15 23:59:59.000000', 270, '2024-02-18 14:15:00.000000', '2024-02-18 14:15:00.000000', 0.00, 18, 18),
('BATCH019', 70, -6.00, '2025-12-01 23:59:59.000000', 85, '2024-02-19 15:30:00.000000', '2024-02-19 15:30:00.000000', -8.00, 19, 19),
('BATCH020', 200, 0.00, '2025-12-15 23:59:59.000000', 220, '2024-02-20 16:45:00.000000', '2024-02-20 16:45:00.000000', -1.00, 20, 20);

INSERT INTO `inbound_orders` (`order_date`, `order_number`, `employee_id`, `product_batch_id`, `warehouse_id`) VALUES
('2024-02-01 08:30:00.000000', 'INB001', 1, 1, 1),
('2024-02-02 09:45:00.000000', 'INB002', 2, 2, 2),
('2024-02-03 10:15:00.000000', 'INB003', 3, 3, 3),
('2024-02-04 11:30:00.000000', 'INB004', 4, 4, 4),
('2024-02-05 12:00:00.000000', 'INB005', 5, 5, 5),
('2024-02-06 13:20:00.000000', 'INB006', 6, 6, 6),
('2024-02-07 14:35:00.000000', 'INB007', 7, 7, 7),
('2024-02-08 15:50:00.000000', 'INB008', 8, 8, 8),
('2024-02-09 16:10:00.000000', 'INB009', 9, 9, 9),
('2024-02-10 17:25:00.000000', 'INB010', 10, 10, 10),
('2024-02-11 18:40:00.000000', 'INB011', 11, 11, 11),
('2024-02-12 19:55:00.000000', 'INB012', 12, 12, 12),
('2024-02-13 20:10:00.000000', 'INB013', 13, 13, 13),
('2024-02-14 21:25:00.000000', 'INB014', 14, 14, 14),
('2024-02-15 22:40:00.000000', 'INB015', 15, 15, 15),
('2024-02-16 23:55:00.000000', 'INB016', 16, 16, 16),
('2024-02-17 08:10:00.000000', 'INB017', 17, 17, 17),
('2024-02-18 09:25:00.000000', 'INB018', 18, 18, 18),
('2024-02-19 10:40:00.000000', 'INB019', 19, 19, 19),
('2024-02-20 11:55:00.000000', 'INB020', 20, 20, 20);


INSERT INTO `product_records` (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES
('2024-02-01 08:00:00.000000', 50.00, 75.00, 1),
('2024-02-02 09:15:00.000000', 30.00, 45.00, 2),
('2024-02-03 10:30:00.000000', 60.00, 90.00, 3),
('2024-02-04 11:45:00.000000', 20.00, 35.00, 4),
('2024-02-05 13:00:00.000000', 100.00, 150.00, 5),
('2024-02-06 14:15:00.000000', 40.00, 60.00, 6),
('2024-02-07 15:30:00.000000', 25.00, 38.00, 7),
('2024-02-08 16:45:00.000000', 80.00, 120.00, 8),
('2024-02-09 18:00:00.000000', 90.00, 135.00, 9),
('2024-02-10 19:15:00.000000', 15.00, 25.00, 10),
('2024-02-11 20:30:00.000000', 55.00, 82.00, 11),
('2024-02-12 21:45:00.000000', 70.00, 105.00, 12),
('2024-02-13 08:00:00.000000', 65.00, 98.00, 13),
('2024-02-14 09:15:00.000000', 110.00, 165.00, 14),
('2024-02-15 10:30:00.000000', 45.00, 68.00, 15),
('2024-02-16 11:45:00.000000', 95.00, 145.00, 16),
('2024-02-17 13:00:00.000000', 35.00, 55.00, 17),
('2024-02-18 14:15:00.000000', 120.00, 180.00, 18),
('2024-02-19 15:30:00.000000', 85.00, 130.00, 19),
('2024-02-20 16:45:00.000000', 50.00, 75.00, 20);



INSERT INTO `order_details` (`clean_liness_status`, `quantity`, `temperature`, `product_record_id`, `purchase_order_id`) VALUES
('Clean', 10, -2.00, 1, 1),
('Dirty', 5, -3.00, 2, 2),
('Clean', 20, 1.00, 3, 3),
('Moderate', 15, -10.00, 4, 4),
('Clean', 30, -18.00, 5, 5),
('Dirty', 8, 0.00, 6, 6),
('Moderate', 25, -5.00, 7, 7),
('Clean', 40, 2.00, 8, 8),
('Dirty', 12, -12.00, 9, 9),
('Clean', 50, 3.00, 10, 10),
('Moderate', 18, -8.00, 11, 11),
('Clean', 22, 0.00, 12, 12),
('Dirty', 35, 6.00, 13, 13),
('Clean', 28, -4.00, 14, 14),
('Moderate', 32, -20.00, 15, 15),
('Clean', 19, -1.00, 16, 16),
('Dirty', 27, -3.00, 17, 17),
('Clean', 55, 2.00, 18, 18),
('Moderate', 14, -6.00, 19, 19),
('Clean', 38, 0.00, 20, 20);
