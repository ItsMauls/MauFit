#!/bin/bash
# The script will exit immediately if a command fails
set -e

# Function to check for and create a database if it does not exist.
# It takes one argument: the name of the database to create.
create_database() {
  local db_name=$1
  # Execute a psql command to check if the database already exists.
  # The query will return '1' if the database exists, and an empty string if it does not.
  local DB_EXISTS=$(psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" -tAc "SELECT 1 FROM pg_database WHERE datname = '$db_name'")

  # Check if the DB_EXISTS variable is empty
  if [ -z "$DB_EXISTS" ]; then
    echo ">>> Database '$db_name' not found. Creating a new database..."
    # Execute a psql command to create the database.
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
      CREATE DATABASE "$db_name";
EOSQL
    echo ">>> Database '$db_name' created successfully."
  else
    echo ">>> Database '$db_name' already exists. No action taken."
  fi
}

# Call the function for each database you want to create.
create_database "maufit-main"
create_database "maufit-user"


# You can add other SQL commands here for setting up tables, etc.
# Make sure to connect to the correct database (--dbname).
# Example for 'maufit-main':
# echo ">>> Running schema and table setup for maufit-main..."
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "maufit-main" <<-EOSQL
#   CREATE TABLE IF NOT EXISTS users (
#     id SERIAL PRIMARY KEY,
#     username VARCHAR(50) UNIQUE NOT NULL
#   );
# EOSQL
# echo ">>> Schema setup for maufit-main finished."

