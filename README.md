# 📝 Sharing Vision - Backend Golang

Test Koding untuk Fullstack Developer Sharing Vision Golang

## Table of contents
1. [Installation](#Installation)
2. [API Docs](#API_Docs)
3. [API Spec](#API_Spec)


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


## API_Docs
- [Postman Collection](https://google.com)
- [Swagger (Open API)](https://google.com)


## API_Spec

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
    "status" : "string, enum: Publish|Draft|Thrash"
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

### Retrieve a list of articles

Request :
- Method : GET
- Endpoint : `/article?limit=10&offset=1`
- Query Param :
  - **limit** : number,
  - **offset** : number
- Header :
    - Accept: application/json

Response :

```json 
{
    "code": "int",
    "status": "int",
    "data" : [
      {
        "id" : "int, unique",
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

### Retrieve a single of article

Request :
- Method : PUT
- Endpoint : `/article/{articleId}`
- Header :
    - Content-Type: application/json
    - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id" : "int, unique",
        "title" : "string",
        "slug" : "long",
        "content" : "string",
        "category" : "string",
        "status" : "string",
        "createdDate" : "date",
        "updatedDate" : "date"
     }
}
```

## Update a article

Request :
- Method : PUT
- Endpoint : `/article/{articleID}`
- Header :
    - Accept: application/json
- Query Param :
    - size : number,
    - page : number

Request Body :

```json
{
  "title": "string",
  "content": "string",
  "category": "string",
  "status": "string, enum:Publish,Draft,Thrash"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
        "id" : "int, unique",
        "title" : "string",
        "slug" : "long",
        "content" : "string",
        "category" : "string",
        "status" : "string",
        "createdDate" : "date",
        "updatedDate" : "date"
    }
}
```

## Delete a article

Request :
- Method : DELETE
- Endpoint : `/article/{articleID}`
- Header :
    - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "message": "string",
}
```