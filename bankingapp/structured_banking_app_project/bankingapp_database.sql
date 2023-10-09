create database bankingapp;

use bankingapp;

drop database bankingapp;
show tables;

SELECT * from customers;
DELETE from customers where id = 4;
SELECT * from banks;
SELECT * from accounts;
DELETE from accounts where id = 1 OR id = 2;
SELECT * from offers;
SELECT * from account_offers_join;
SELECT * from bank_passbooks;
SELECT * from bank_entries;
SELECT * from passbooks;
SELECT * from entries;



DROP table bank_offers;

DROP table account_offers;








SELECT * from users;
DELETE from users where id = 10 OR id = 11 OR id = 12;

SELECT * from contacts;
DELETE from contacts where id = 4;

SELECT * from contact_infos;
TRUNCATE TABLE contact_infos;