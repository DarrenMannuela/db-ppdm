FROM doctorkirk/oracle-19c:latest
COPY ./sql_tables ./sql_tables
COPY ./startup/oracle/ ./
# COPY ./startup/oracle/start-db.sh ./
# COPY ./startup/oracle/start.sh ./
CMD ./start.sh