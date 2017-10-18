FROM postgres:9.6-alpine

RUN mkdir -p /docker-entrypoint-initdb.d
COPY init-db.sh /docker-entrypoint-initdb.d/

