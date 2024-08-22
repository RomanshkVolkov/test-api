#!/usr/bin/env bash
find / -name 'sqlcmd' 2>/dev/null
/opt/mssql/bin/sqlservr &
./db_setup.sh
