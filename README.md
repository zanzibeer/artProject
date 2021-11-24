#rest-api-tutorial

#user-service

#REST-API

GET /users -- list of users -- 200, 404, 500
GET /users/:id -- users by id -- 200, 404, 500
POST /users/:id -- create user -- 204, 4xx, Header Location: url
PUT /users/:id -- fully update user by id -- 200/204, 400, 404, 500
PATCH /users/:id -- partialy update user bu id -- 200/204, 400, 404, 500
DELETE /users/:id -- delete user bu id -- 204, 404, 500