FROM postgres:18.0 AS database

COPY ./create_database.sql /docker-entrypoint-initdb.d

CMD ["postgres"]