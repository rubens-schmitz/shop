create table cart (
    id int generated always as identity primary key,
    price dec not null check(price >= 0),
    quantity int not null check(quantity >= 0),
    datestamp text not null
);

create table category (
    id int generated always as identity primary key,
    title text not null,
    deleted boolean default false
);

create table access (
    id int generated always as identity primary key,
    class text not null,
    code text not null
);

create table product (
    id int generated always as identity primary key,
    title text not null,
    price dec not null check(price >= 0),
    deleted boolean default false,
    categoryId int not null,
    foreign key(categoryId) references category(id) on delete cascade
);

create table item (
    id int generated always as identity primary key,
    quantity int not null default 1 check(quantity >= 1),
    productId int not null,
    cartId int not null,
    foreign key(productId) references product(id) on delete cascade,
    foreign key(cartId) references cart(id) on delete cascade
);

create table picture (
    id int generated always as identity primary key,
    bytes bytea not null,
    productId int not null,
    foreign key(productId) references product(id) on delete cascade
);

create table deal (
    id int generated always as identity primary key,
    accessId int not null,
    cartId int not null,
    foreign key(accessId) references access(id) on delete cascade,
    foreign key(cartId) references cart(id) on delete cascade
);
