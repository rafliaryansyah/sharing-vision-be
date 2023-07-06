# 📝 Sharing Vision - Backend Golang

Test Koding untuk Fullstack Developer Sharing Vision Golang

## Table of contents
1. [Installation](#Installation)
2. [API Docs](#API Docs)
3. [API Spec](#Api Spec)


## Installation

Clone project from repository

```
git clone https://github.com/rafliaryansyah/sharing-vision-be.git
```

Switch to the repo folder

```
cd sharing-vision-be
```

#### Configuration .env file

Copy the example env file and make the required configuration changes in the .env file

```
cp .env.example .env
```


```
APP_PORT=

DATABASE_USERNAME=root
DATABASE_PASSWORD=
DATABASE_NETWORK=tcp
DATABASE_ADDRESS="localhost:3306"
DATABASE_NAME=sharing_visions
```

Run Database Migration, Make sure already install database [migration](https://github.com/golang-migrate/migrate) & Database already exists
```
migrate -database "mysql://USERNAME:PASSWORD@tcp(localhost:3306)/NAMA_DATABASE" -path database/migrations up
```

Install all the dependencies using go mod tidy

```
go mod tidy
```

Start the local development server

```
go run .
```

You can now access the server at http://localhost8080


## API Docs
- [Postman Collection](https://google.com)
- [Swagger (Open API)](https://google.com)


## Api Spec

### Create a article

Request :
- Method : POST
- Endpoint : `/article`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "title" : "string, unique",
    "content" : "string",
    "category" : "string",
    "status" : "enum: Publish|Draft|Thrash"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
         "id" : "string, unique",
         "title" : "string",
         "slug" : "string",
         "content" : "string",
         "category" : "string",
         "createdDate" : "date",
         "updatedDate" : "date"
     }
}
```

## Retrieve a list of articles

Request :
- Method : GET
- Endpoint : `/article?limit=10&offset=1`
- Query
  - limit: int
  - offset: int
- Header :
    - Accept: application/json

Response :

```json 
{
    "code": "int",
    "status": "int",
    "data" : [
      {
        "id" : "string, unique",
        "title" : "string",
        "slug" : "long",
        "content" : "string",
        "category" : "string",
        "status" : "string",
        "createdDate" : "date",
        "updatedDate" : "date"
      }
    ],
    "meta": {
      "total": "int",
      "totalPage": "int",
      "perPage": "int",
      "page": "int"
    }
}
```

## Update Product

Request :
- Method : PUT
- Endpoint : `/api/products/{id_product}`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
    "name" : "string",
    "price" : "long",
    "quantity" : "integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
         "id" : "string, unique",
         "name" : "string",
         "price" : "long",
         "quantity" : "integer",
         "createdAt" : "date",
         "updatedAt" : "date"
     }
}
```

## List Product

Request :
- Method : GET
- Endpoint : `/api/products`
- Header :
    - Accept: application/json
- Query Param :
    - size : number,
    - page : number

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
             "id" : "string, unique",
             "name" : "string",
             "price" : "long",
             "quantity" : "integer",
             "createdAt" : "date",
             "updatedAt" : "date"
        },
        {
             "id" : "string, unique",
             "name" : "string",
             "price" : "long",
             "quantity" : "integer",
             "createdAt" : "date",
             "updatedAt" : "date"
         }
    ]
}
```

## Delete Product

Request :
- Method : DELETE
- Endpoint : `/api/products/{id_product}`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string"
}
```