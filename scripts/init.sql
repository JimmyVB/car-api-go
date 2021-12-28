CREATE TABLE public.pokemon
(
    id varchar NULL,
    name numeric NULL,
    order numeric NULL,
    height numeric NULL,
    weight numeric NULL,
    category varchar NULL
);

CREATE TABLE public.users
(
    id serial NULL,
    username varchar NULL,
    password varchar NULL,
    CONSTRAINT user_pk PRIMARY KEY (id)
);
