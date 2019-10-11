FROM mysql:8.0.16
COPY schema.sql /docker-entrypoint-initdb.d/schema.sql
CMD ["--character-set-server=utf8mb4", "--collation-server=utf8mb4_general_ci"]

