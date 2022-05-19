FROM mysql/mysql-server:latest
COPY ./mysql/sql-scripts/ /docker-entrypoint-initdb.d/
COPY ./data/ /var/lib/mysql-files