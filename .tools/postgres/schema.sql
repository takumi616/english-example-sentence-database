create table userinfo(
    id  serial PRIMARY KEY,
    user_name varchar(30) not null,
    user_password varchar(30) not null,
    user_role varchar(80) not null,
    created timestamp not null,
    updated timestamp not null
);

create table vocabulary(
    id       serial     PRIMARY KEY,
    title    varchar(30)   not null,
    example varchar(200)  not null,
    created  timestamp  not null,
    updated  timestamp  not null
);
