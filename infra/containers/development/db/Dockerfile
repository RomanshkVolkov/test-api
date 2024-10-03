FROM mcr.microsoft.com/mssql/server:2022-CU14-ubuntu-22.04

ENV SA_PASSWORD=${SA_PASSWORD}
ENV ACCEPT_EULA=Y 

COPY init-database.sql init-database.sql
COPY db_setup.sh db_setup.sh 
COPY entrypoint.sh entrypoint.sh

CMD /bin/bash ./entrypoint.sh