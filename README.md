# Bible API

Bible API to get bible character data based on chapter and location. See graph/schema/schema.graphqls to see the available queries and mutations.

---

## Database Setup:

run `docker compose up -d db --build` to start mysql database. Need to have a .env file with the following (can change values, this is just for example):

```
MYSQL_ROOT_PASSWORD= root
DB_USER= user
DB_PASSWORD= dbpass
DB_NAME= bible_api
DB_HOST= localhost:3305
SERVER_PORT= 8080
```
