CREATE TABLE property_type (
    id serial primary key,
    name varchar not null
);

CREATE TABLE confectionary_type (
    id serial primary key,
    name varchar not null
);

CREATE TABLE units (
    id serial primary key,
    name varchar not null
);

CREATE TABLE packaging (
    id serial primary key,
    name varchar not null
);

CREATE TABLE city (
    id serial primary key,
    name varchar not null
);

CREATE TABLE district (
    id serial primary key,
    name varchar not null,
    city_id integer references city(id) not null
);

CREATE TABLE factory (
    id serial primary key,
    name varchar not null,
    city_id integer references city(id),
    phone_number char(11) not null,
    year int not null,
    property_type integer references property_type(id) not null
);

CREATE TABLE product (
    id serial primary key,
    name varchar not null,
    type_id integer references confectionary_type(id) not null,
    factory_id integer references factory(id) not null
);

CREATE TABLE product_catalogue (
    id serial primary key,
    product_id integer references product(id),
    package_id integer references packaging(id),
    unit_id integer references units(id)
);

CREATE TABLE shop (
    id serial primary key,
    name varchar not null,
    district_id integer references district(id),
    opened_at integer not null,
    phone_number char(11) not null,
    employees_count integer not null
);

CREATE TABLE shipments (
    id serial primary key,
    catalogue_id integer references product_catalogue(id) not null,
    shop_id integer references shop(id) not null,
    quantity integer not null,
    cost decimal not null
);
