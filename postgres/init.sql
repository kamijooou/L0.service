create database l0_test owner l0_user;
\connect l0_test

create domain posint as int check (value >= 0);
create domain bigposint as bigint check (value >= 0);

create table orders (
    id serial primary key,
    order_uid varchar(100) unique not null,
    track_number varchar(100) unique not null,
    entry varchar(30),
    locale varchar(15),
    internal_signature varchar(100),
    customer_id varchar(100) not null,
    delivery_service varchar(100),
    shardkey varchar(20),
    sm_id posint not null,
    date_created timestamp with time zone not null,
    oof_shard varchar(50)
);

create index if not exists o_uid on orders(order_uid);

create table deliveries(
    order_id int primary key,
    name varchar(50),
    phone varchar(11) not null,
    zip varchar(7),
    city varchar(50) not null,
    address varchar(100) not null,
    region varchar(50) not null,
    email varchar(50) not null,
    foreign key(order_id) references orders(id)
);

create table payments (
    transaction varchar(100) primary key,
    request_id varchar(100),
    currency varchar(30) not null,
    provider varchar(50) not null,
    amount posint not null,
    payment_dt bigposint not null,
    bank varchar(30) not null,
    delivery_cost posint not null,
    goods_total posint not null,
    custom_fee posint,
    foreign key(transaction) references orders(order_uid)
);

create table items (
    id serial primary key,
    chrt_id posint not null,
    track_number varchar(100) not null,
    price posint not null,
    rid varchar(100) not null,
    name varchar(100) not null,
    sale posint,
    size varchar(100) not null,
    total_price posint not null,
    nm_id bigposint not null,
    brand varchar(100),
    status posint not null,
    foreign key (track_number) references orders(track_number) on delete cascade
);