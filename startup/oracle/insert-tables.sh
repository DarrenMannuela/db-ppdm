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
echo "Insert into ROLES (ROLE_NAME) values ('Admin');" | sqlplus $ORACLE_CONN_STR
echo "Insert into USERS (FIRST_NAME,LAST_NAME,EMAIL,PASSWORD,DATE_JOINED,EXPIRED_DATE,ROLE_NAME) values ('John','Richardson','john.richardson@gtn.id','$2a$14$wDOPvFpU5cWyNnRu2QpZYu/ArkVhTQkLVsYyY9tJkktLo8ZrFGIMG',to_date('23-03-2023','DD-MM-RRRR'),to_date('23-03-2025','DD-MM-RRRR'),'Admin');" | sqlplus $ORACLE_CONN_STR
