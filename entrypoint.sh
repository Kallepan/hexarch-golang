#!/bin/sh

set -e
COMMAND=$@

echo 'Waiting for database to be available...'
maxTries=10
while [ "$maxTries" -gt 0 ] && ! PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c '\q'; do
    maxTries=$(($maxTries - 1))
    sleep 3
done
echo
if [ "$maxTries" -le 0 ]; then
    echo >&2 'error: unable to contact postgres after 10 tries'
    exit 1
fi

exec $COMMAND