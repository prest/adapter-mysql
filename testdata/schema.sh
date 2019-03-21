DB_HOST=${PREST_PG_HOST:-localhost}
DB_USER=${PREST_PG_USER:-root}
DB_PORT=${PREST_PG_PORT:-3306}
DB_NAME=${PREST_PG_DATABASE:-prest}

mysql -h $DB_HOST -P $DB_PORT -u $DB_USER -e "DROP DATABASE IF EXISTS $DB_NAME;"
mysql -h $DB_HOST -P $DB_PORT -u $DB_USER -e "CREATE DATABASE $DB_NAME;"

cat testdata/schema.sql | mysql -h $DB_HOST -P $DB_PORT -u $DB_USER $DB_NAME
# psql -d $DB_NAME -h $DB_HOST -p $DB_PORT -U $DB_USER -f testdata/schema.sql
