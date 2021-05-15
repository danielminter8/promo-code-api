# Basic Go API


### Small golang project using the gin web framework along with gorm and postgres for the database.

### Instructions to run project using docker:


- Clone this repo first.
- Make sure you have docker installed.
- In terminal, navigate to the project directory and type the following:
``` docker compose build ```
- Once the build completes run the following:
- ``` docker compose start ```
- Navigate to http://localhost:8080/swagger/index.html#/ in your browser where you will see documentation of how to use the api.

### Instructions to run program without docker:

- Clone this repo first.
- Make sure you have golang and postgres installed.
- Before running app, create a .env file in the project root. Inside the ***.env*** you should have the following with your database details filled in:
```
e.g
DB_USER=postgres
DB_PASSWORD=pg123
DB_NAME=postgres
DB_HOST=localhost
DB_PORT=5432

```
- In terminal, navigate to project directory and type the following:
``` go run .```
- Navigate to http://localhost:5000/swagger/index.html#/ in your browser where you will see documentation of how to use the api.
