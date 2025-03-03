

-- name: CreateUser :one
insert into users(
    id ,
    user_name,
    email,
    password,
    role
) values ($1,$2,$3,$4,$5) returning * ;


-- name: GetUserByEmail :one 
select * from users where email=$1 ;