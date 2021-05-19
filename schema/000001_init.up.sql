CREATE TABLE contacts
(
    id serial not null unique,
    phone varchar(255) not null unique,
    name varchar(255) not null
);