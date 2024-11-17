# Limiter API Request Rate

> Made with Golang with GORM and Gin
>
> Databases: SQLite3 and Redis

## .Env Example:

> REDIS_ADDR=<SEU_ENDERECO_REDIS>
>
> REDIS_PASSWORD=<SUA_SENHA_REDIS>

## Endpoints

### GET /user

### POST /user

> {<br>
    "username": "username",<br>
    "password": "password",<br>
    "email": "email"<br>
    "name": "name"<br>
> }

### PUT /user
> Pode aproveitar o payload da criação, e passar apenas o que deseja mudar

### DELETE /user

## Login

### POST /login
> {<br>
    "username": "username",<br>
    "password": "password"<br>
> }
> 
>MULTIPART FORM

### GET /login
> Retorna o UserID e o Username com base no Token Recebido


## Keys
### GET /keys
### POST /keys
### DELETE /keys

## Packages
### POST /pack/{amount: int}
### GET /pack

## Requests
### POST /request/
> Header: "X-API-KEY: <SUA_API_KEY>"
### GET /request/
> Header: "X-API-KEY: <SUA_API_KEY>"