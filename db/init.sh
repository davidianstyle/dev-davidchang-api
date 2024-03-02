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

# Create the dynamic SQL script in a writable directory
# echo "CREATE DATABASE IF NOT EXISTS \`$DB_NAME\` ;" > /tmp/init.sql
# echo "USE \`$DB_NAME\` ;" >> /tmp/init.sql
# echo "source /docker-entrypoint-initdb.d/schema.sql;" >> /tmp/init.sql

# # Move the script to the /docker-entrypoint-initdb.d/ directory
# sudo mv /tmp/init.sql /docker-entrypoint-initdb.d/

# Execute SQL commands directly using the mysql command
# echo "Creating database: \`$DB_NAME\`"
# mysql -h localhost -P $MYSQL_PORT -u $DB_USER -p$MYSQL_ROOT_PASSWORD -e "CREATE DATABASE IF NOT EXISTS \`$DB_NAME\` ;"
# echo "Using database: \`$DB_NAME\`"
# mysql -h localhost -P $MYSQL_PORT -u $DB_USER -p$MYSQL_ROOT_PASSWORD -e "USE \`$DB_NAME\` ;"
# echo "Running initialization"
# mysql -h localhost -P $MYSQL_PORT -u $DB_USER -p$MYSQL_ROOT_PASSWORD $DB_NAME < /docker-entrypoint-initdb.d/schema.sql

# Call the original entrypoint script to start MySQL
#exec /entrypoint.sh "$@"
