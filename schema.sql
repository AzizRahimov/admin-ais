CREATE table "user"
(
    id   bigserial primary key,
    name text NOT NULL
);


CREATE table users_roles
(
    user_id integer references "users",
    role_id integer references "roles"
);

CREATE TABLE users_rights
(
    user_id integer references "users",
    right_id integer references "rights"
);