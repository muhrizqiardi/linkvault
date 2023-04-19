-- name: GetUsers :many
select * from public.users;

-- name: CreateUser :exec
insert into public.users (email, full_name, password) values ($1, $2, $3);

-- name: GetOneUserById :one
select * from public.users 
    where id = $1;
