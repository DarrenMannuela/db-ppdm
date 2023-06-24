# Startup script to insert sql tables
# Written by Bryn Ghiffar

RACLE_PORT=1521
ORACLE_USER=system
ORACLE_PWD=$(printenv ORACLE_PWD)
ORACLE_SID=$(printenv ORACLE_SID)
ORACLE_CONN_STR="${ORACLE_USER}/${ORACLE_PWD}@localhost:${ORACLE_PORT}/${ORACLE_SID}"

# echo "the password is ${ORACLE_CONN_STR}"
until echo "exit" | sqlplus $ORACLE_CONN_STR | grep -i "Connected to:"
do 
	echo "Waiting until ready ($count) * 10s"
	count=$[$count + 1]
	sleep 10s
done
echo "@/home/oracle/sql_tables/submission.sql" | sqlplus $ORACLE_CONN_STR
echo "@/home/oracle/sql_tables/roles.sql" | sqlplus $ORACLE_CONN_STR
echo "@/home/oracle/sql_tables/users.sql" | sqlplus $ORACLE_CONN_STR
echo "@/home/oracle/sql_tables/permen.sql" | sqlplus $ORACLE_CONN_STR
echo "@/home/oracle/sql_tables/afe.sql" | sqlplus $ORACLE_CONN_STR
echo "@/home/oracle/sql_tables/workspace.sql" | sqlplus $ORACLE_CONN_STR
