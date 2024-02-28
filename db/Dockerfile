# db/Dockerfile

# Use an official MySQL image as a base image
FROM mysql:latest

# Set the root password and create the database
ENV MYSQL_ROOT_PASSWORD $DB_PASSWORD
ENV MYSQL_DATABASE $DB_NAME

# Copy the init.sql script to the docker-entrypoint-initdb.d/ directory
COPY ./init.sql /docker-entrypoint-initdb.d/
