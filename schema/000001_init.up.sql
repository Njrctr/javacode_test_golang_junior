CREATE TABLE users
(
    id            serial       not null unique,
    email          varchar(255) not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE wallets
(
    id            serial       not null unique,
    uuid          uuid      DEFAULT gen_random_uuid(),
    balance  int   CHECK (balance >= 0) DEFAULT 0
);

CREATE TABLE users_wallets
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    wallet_id int references wallets (id) on delete cascade not null
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";