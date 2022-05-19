FROM mysql/mysql-server:latest
COPY ./mysql/sql-scripts/ /docker-entrypoint-initdb.d/
COPY ./mysql/ /var/lib/mysql-files