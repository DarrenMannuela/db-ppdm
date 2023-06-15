import oracledb
import sys
import lib_platform as platform

if platform.system == "darwin":
    sys.path.insert(0, "database")
elif platform.system == "windows":
    sys.path.insert(0, "../database")

from conn import connection

# pw = getpass.getpass("Enter password: ")

# connection = oracledb.connect(
#     user="admin",
#     password="ppdm-3.9",
#     dsn="/kei-oracle-test.chkxr21l2l5s.ap-southeast-3.rds.amazonaws.com:1521/ORCL")

# print("Successfully connected to Oracle Database")

cursor = connection.cursor()


cursor.execute("""
    begin
        execute immediate 'drop table project';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create table project (
        PROJECT_ID VARCHAR(40) PRIMARY KEY NOT NULL,
        ACTIVE_IND VARCHAR(1),
        COMPLETE_DATE TIMESTAMP(0),
        CONFIDENTIAL_IND VARCHAR(1),
        CONFIDENTIAL_RELEASE_DATE TIMESTAMP(0),
        DESCRIPTION VARCHAR(240),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        FIELD_STATION_IND VARCHAR(1),
        LAND_RIGHT_IND VARCHAR(1),
        PDEN_IND VARCHAR(1),
        PPDM_GUID VARCHAR(38),
        PROJECT_NAME VARCHAR(255),
        PROJECT_NUM VARCHAR(30),
        PROJECT_TYPE VARCHAR(40),
        REMARK VARCHAR(2000),
        SEIS_SET_IND VARCHAR(1),
        SOURCE VARCHAR(40),
        START_DATE TIMESTAMP(0),
        STRAT_COLUMN_IND VARCHAR(1),
        STRAT_INTERPRETATION_IND VARCHAR(1),
        WELL_IND VARCHAR(1),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")
