
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


-- name: CreateProfile :one
insert into user_profiles (
    user_id ,
    bio ,
    location 
) values ($1,$2,$3) returning *;


-- name: GetUserProfile :one
select * from users 
join user_profiles 
on users.id =user_profiles.user_id;