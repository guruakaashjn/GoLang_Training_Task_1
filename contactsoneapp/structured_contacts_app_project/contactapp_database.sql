create database contactapp;

use contactapp;


show tables;

SELECT * from users;
DELETE from users where id = 5 OR id = 6 OR id = 7;

SELECT * from contacts;
DELETE from contacts where id = 4;

SELECT * from contact_infos;
TRUNCATE TABLE contact_infos;