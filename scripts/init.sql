-- USUARIOS

select * from users;

CREATE TABLE public.users
(
    id serial NOT NULL,
    username varchar NULL,
    password varchar NULL,
    CONSTRAINT user_pk PRIMARY KEY (id)
)

-- ROLES

select * from roles

CREATE TABLE public.roles
(
    id serial NOT NULL,
    nombre varchar NULL,
    CONSTRAINT roles_pk PRIMARY KEY (id)
)

insert into roles (nombre) values ('ROLE_ADMIN');
insert into roles (nombre) values ('ROLE_USER');


select * from usuarios_roles

CREATE TABLE public.usuarios_roles
(
    usuario_id int NOT NULL,
    role_id int NOT NULL,
    CONSTRAINT usuarios_roles_pk PRIMARY KEY (usuario_id, role_id)
)

insert into usuarios_roles (usuario_id, role_id) values (1, 1);
insert into usuarios_roles (usuario_id, role_id) values (2, 2);


select * from cars

CREATE TABLE public.cars
(
    id serial NOT NULL,
    marca varchar NULL,
    model varchar NULL,
    price int NULL,
    CONSTRAINT cars_pk PRIMARY KEY (id)
)
