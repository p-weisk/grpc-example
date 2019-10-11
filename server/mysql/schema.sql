SET NAMES utf8mb4;
SET CHARSET utf8mb4;
CREATE DATABASE grpc CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
CREATE TABLE grpc.product (
    Id varchar(255) PRIMARY KEY
);
CREATE TABLE grpc.invoice (
    Number integer PRIMARY KEY,
    ClientId varchar(255) NOT NULL,
    P varchar(255) NOT NULL,
    FOREIGN KEY Fk_P(P) REFERENCES grpc.product(Id) ON DELETE CASCADE
);
CREATE USER dev@'%' IDENTIFIED BY 'dev';
GRANT ALL PRIVILEGES ON grpc.* TO dev@'%';

INSERT INTO grpc.product(Id) VALUES ('Product-0000');
INSERT INTO grpc.product(Id) VALUES ('Product-0001');
INSERT INTO grpc.product(Id) VALUES ('Product-0002');
INSERT INTO grpc.product(Id) VALUES ('Product-0003');

INSERT INTO grpc.invoice(Number, ClientId, P) VALUES ('1', 'Client-0001', 'Product-0000');
INSERT INTO grpc.invoice(Number, ClientId, P) VALUES ('2', 'Client-0002', 'Product-0000');
