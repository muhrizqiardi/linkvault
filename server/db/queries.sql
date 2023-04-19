-- name: GetUsers :many
select * from public.users;

-- name: CreateUser :exec
insert into public.users (email, full_name, password) values ($1, $2, $3);

-- name: GetOneUserById :one
select * from public.users 
    where id = $1;

-- name: UpdateOneUserById :exec
update public.users
    set 
        email = $2,
        password = $3,
        full_name = $4
    where 
        id = $1;
        
-- name: DeleteOneUserById :exec
delete from public.users where id = $1;
