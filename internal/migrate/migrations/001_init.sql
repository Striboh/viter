-- see for more examples: https://github.com/rubenv/sql-migrate?tab=readme-ov-file#writing-migrations

-- +migrate Up

-- +migrate StatementBegin
create table roles (
name varchar(100) NOT NULL PRIMARY KEY
);
create table profiles (
id SERIAL NOT NULL PRIMARY KEY,
first_name varchar(100) NOT NULL,
last_name varchar(100) NOT NULL,
email varchar(100),
phone varchar(100) NOT NULL
);

create table role_records (
profile_id SERIAL NOT NULL references profiles(id),
role varchar(100) NOT NULL references roles(name),
PRIMARY KEY (profile_id, role)
);

create table warehouses (
id SERIAL NOT NULL PRIMARY KEY,
owner_id int NOT NULL references profiles(id),
adress varchar(300) NOT NULL
);

create table clients (
id int NOT NULL PRIMARY KEY references profiles(id),
capacity int NOT NULL --managed by admins
);

create table volunteers (
id int NOT NULL PRIMARY KEY references profiles(id),
score int NOT NULL --managed by system rule "donated - get score"
);

create table orders (
id SERIAL NOT NULL PRIMARY KEY,
consumer_id int NOT NULL references profiles(id),
date_created date NOT NULL
);


create table devices (
id SERIAL NOT NULL PRIMARY KEY,
name varchar(100) NOT NULL,
score int NOT NULL
);

create table items_ordered (
id SERIAL NOT NULL PRIMARY KEY,
order_id int NOT NULL references orders(id),
device_id int NOT NULL references devices(id),
quantity int NOT NULL,
UNIQUE (order_id, device_id)
);

create table donations (
id SERIAL NOT NULL PRIMARY KEY,
donator_id int NOT NULL references volunteers(id),
device_id int NOT NULL references devices(id),
warehouse_id int references warehouses(id),
quantity int NOT NULL
);

create table inventory_records (
id SERIAL NOT NULL PRIMARY KEY,
warehouse_id int NOT NULL references warehouses(id),
device_id int NOT NULL references devices(id),
quantity int NOT NULL,
UNIQUE (warehouse_id, device_id)
);

create table device_parts_info (
id SERIAL NOT NULL PRIMARY KEY,
composite_id int NOT NULL references devices(id),
part_id int NOT NULL references devices(id),
quantity int NOT NULL,
UNIQUE (composite_id, part_id)
);

create table api_clients (
id SERIAL NOT NULL PRIMARY KEY,
name varchar(100),
permissions_lvl varchar(100) NOT NULL,
hash varchar(100) NOT NULL
);
-- +migrate StatementEnd

-- +migrate Down

-- +migrate StatementBegin
drop table api_clients;
drop table device_parts_info;
drop table items_ordered;
drop table donations;
drop table orders;
drop table inventory_records;
drop table devices;
drop table warehouses;
drop table clients;
drop table volunteers;
drop table role_records;
drop table roles;
drop table profiles;
-- +migrate StatementEnd