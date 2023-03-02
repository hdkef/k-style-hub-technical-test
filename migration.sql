/*
CREATE DATABASE FIRST
*/
CREATE DATABASE IF NOT EXISTS db_kstyle;

/*
USE DATABASE THAT WAS CREATED
*/
USE db_kstyle;

/*
CREATE TABLE PRODUCT
*/
CREATE TABLE product(id_product int primary key auto_increment, name_product VARCHAR(50) NOT NULL, price DECIMAL NOT NULL);

/*
CREATE TABLE MEMBER
*/
CREATE TABLE member(id_member int primary key auto_increment, username VARCHAR(25) UNIQUE NOT NULL,
gender ENUM('Female','Male') NOT NULL, skintype ENUM('Normal','Dry','Oily') NOT NULL,
skincolor ENUM('Type 1','Type 2','Type 3','Type 4','Type 5','Type 6') NOT NULL);

/*
CREATE TABLE REVIEW
*/
CREATE TABLE review(id_review int primary key auto_increment, id_product INT NOT NULL, 
id_member INT NOT NULL, desc_review TEXT, 
CONSTRAINT fk_rev_mem FOREIGN KEY (id_member) REFERENCES member(id_member),
CONSTRAINT fk_rev_prod FOREIGN KEY (id_product) REFERENCES product(id_product)
);

/*
CREATE TABLE LIKE_REVIEW
*/
CREATE TABLE like_review(id_review int NOT NULL, id_member INT NOT NULL, 
CONSTRAINT fk_lr_mem FOREIGN KEY (id_member) REFERENCES member(id_member),
CONSTRAINT fk_lr_rev FOREIGN KEY (id_review) REFERENCES review(id_review)
);