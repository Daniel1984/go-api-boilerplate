#!/bin/bash

# this script makes sure we have test db created for us
# when docker-compose is run
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    DROP DATABASE IF EXISTS boilerplatetest;
    CREATE DATABASE boilerplatetest;
EOSQL
