CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE invoices
(
    id serial not null unique,
    amount int not null,
    account varchar(255) not null,
    message varchar(255)
);

CREATE TABLE users_invoices
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    invoice_id int references invoices (id) on delete cascade not null
);