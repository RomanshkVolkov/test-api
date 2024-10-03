#!/usr/bin/env bash
# Wait for database to startup

sleep 30
# sqlcmd init-database.sql
/opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P $SA_PASSWORD -C -i init-database.sql
tail -f /dev/null
