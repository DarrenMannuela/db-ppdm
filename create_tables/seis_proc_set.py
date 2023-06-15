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
        execute immediate 'drop table seis_proc_set';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """CREATE TABLE seis_proc_set(
        SEIS_SET_SUBTYPE VARCHAR(30) NOT NULL,
        SEIS_PROC_SET_ID VARCHAR(40) NOT NULL,
        ACTIVE_IND VARCHAR(1),
        AREA_ID VARCHAR(40),
        AREA_TYPE VARCHAR(40),
        DESCRIPTION VARCHAR(2000),
        EFFECTIVE_DATE TIMESTAMP(0),
        END_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        MAX_LATITUDE DECIMAL(14,9),
        MAX_LONGITUDE DECIMAL(14,9),
        MIN_LATITUDE DECIMAL(14,9),
        MIN_LONGITUDE DECIMAL(14,9),
        OBJECTIVE VARCHAR(2000),
        ORIGINAL_PROC_IND VARCHAR(1),
        PPDM_GUID VARCHAR(38),
        PROCESSED_FOR VARCHAR(40),
        PROCESSING_COMPANY VARCHAR(40),
        PROCESSING_NAME VARCHAR(255),
        PROCESS_STATUS VARCHAR(40),
        PROC_SET_TYPE VARCHAR(40),
        REMARK VARCHAR(2000),
        REPROCESSED_IND VARCHAR(1),
        SOURCE VARCHAR(40),
        START_DATE TIMESTAMP(0),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")