#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=toor psql -h "postgres" -U "user" -d "taskDB" -c '\q'>/dev/null 2>&1; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing migrations"

for file in /app/migrations/*.sql; do
  PGPASSWORD=toor psql -h "postgres" -U "user" -d "taskDB" -f "$file"
done

exec $cmd