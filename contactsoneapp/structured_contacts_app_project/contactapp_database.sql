create database contactapp;

use contactapp;

drop database contactapp;
show tables;

SELECT * from users;
DELETE from users where id = 10 OR id = 11 OR id = 12;

SELECT * from contacts;
DELETE from contacts where id = 4;

SELECT * from contact_infos;
TRUNCATE TABLE contact_infos;