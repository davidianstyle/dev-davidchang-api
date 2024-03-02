#!/bin/bash
set -e

# Set environment variables
source ./.env

# Create the dynamic SQL script
echo "CREATE DATABASE IF NOT EXISTS $DB_NAME ;" > ./init.sql

# Append other SQL commands for database initialization if needed
echo "USE $DB_NAME ;" >> ./init.sql
cat ./schema.sql >> ./init.sql

# Done
echo "init.sql file created"
