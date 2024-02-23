# Instagram backend

A monolith API for INSTAGRAM (not official) writen in go.

stack:

    - Go
    - Postgres
    - Gorilla Mux

Having `GO` and `POSGRES` installed on your system is all you need

# Authentication

This API uses JWT (JSON Web Tokens) for authentication. When a user logs in, 
the API generates a JWT containing the user's id and username.

### Routes:

Remember to prefix /v1/ on every url

&nbsp;

| Method | Url                 | Description               |
| ------ | ------------------- | ------------------------- |
| POST   | /auth/signup        | Creates a new user.       |
| POST   | /auth/login         | Login a user              |


# Post

To post photos, you need to be authenticated. The `auth_middleware` 
will first verify if you are authenticated.

Creating a post accepts form with field names `image` and `descriprion`

When a post if created, the Image is stored in a folder called `uploads`
and for every user, a folder is created to store photos.
The images are served and accessed on `/v1/uploads/{username}}/{name-of-image}",`

### Routes:

Remember to prefix /v1/ on every url

&nbsp;

| Method | Url                 | Description               |
| ------ | ------------------- | ------------------------- |
| POST   | /posts              | Creates a new posts.      |
| GET    | /posts              | Get all posts             |
| GET    | /posts/{postID}     | Get post by id            |
| UPDATE | /posts/{postID}     | Update post by ID         |
| DELETE | /posts/{postID}     | DELETE post by id         |


# Like

To like photos, you need to be authenticated. The `auth_middleware` 
will first verify if you are authenticated.

### Routes:

Remember to prefix /v1/ on every url

&nbsp;

| Method | Url                 | Description               |
| ------ | ------------------- | ------------------------- |
| POST   | /like/{postID}      | Like post by ID           |
| DELETE | /unlike/{postID}    | Unlike post by ID         |
| GET    | /likes/{postID}     | Get all likes post by id  |


# Follow

To follow users, you need to be authenticated. The `auth_middleware` 
will first verify if you are authenticated.

### Routes:

Remember to prefix /v1/ on every url

&nbsp;

| Method | Url                 | Description                  |
| ------ | ------------------- | ---------------------------- |
| POST   | /follow/{userID}    | FOllow user by ID            |
| DELETE | /unfollow/{userID}  | Unfollow user by ID          |
| GET    | /followers          | Get all followers for a user |



.