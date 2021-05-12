# Basic Go API


### Small golang project using the gin web framework along with gorm and postgres for the database.

### Instructions to run project using docker:


- Clone this repo first.
- Make sure you have docker installed.
- In terminal, navigate to project directory and type the following:
``` docker compose up --build ```
- It will run once the build completes.
- Navigate to http://localhost:5000/swagger/index.html#/ in your browser where you will see documentation of our to use the api.

### Instructions to run program without docker:

- Make sure you have golang and postgres installed.
- In terminal, navigate to project directory and type the following:
``` go run .```
- Navigate to http://localhost:5000/swagger/index.html#/ in your browser where you will see documentation of our to use the api.
