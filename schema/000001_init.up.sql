CREATE TABLE users
(
    id serial not null unique,
    created_at timestamp not null,
    company_name varchar(255),
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    role int not null
);

CREATE TABLE invoices
(
    id serial not null unique,
    uuid int,
    created_at timestamp not null,
    account varchar(255) not null,
    amount int not null,
    client_name varchar(255),
    message varchar(255),
    status int not null
);

CREATE TABLE users_invoices
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    invoice_id int references invoices (id) on delete cascade not null
);