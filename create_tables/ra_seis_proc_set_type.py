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
        execute immediate 'drop table ra_seis_proc_set_type';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table ra_seis_proc_set_type(
        PROC_SET_TYPE VARCHAR(40) NOT NULL,
        ALIAS_ID VARCHAR(40) NOT NULL,
        ABBREVIATION VARCHAR(30),
        ACTIVE_IND VARCHAR(1),
        ALIAS_LONG_NAME VARCHAR(255),
        ALIAS_REASON_TYPE VARCHAR(40),
        ALIAS_SHORT_NAME VARCHAR(30),
        ALIAS_TYPE VARCHAR(40),
        AMENDED_DATE TIMESTAMP(0),
        CREATED_DATE TIMESTAMP(0),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        LANGUAGE_ID VARCHAR(40),
        ORIGINAL_IND VARCHAR(1),
        OWNER_BA_ID VARCHAR(40),
        PPDM_GUID VARCHAR(38),
        PREFERRED_IND VARCHAR(1),
        REASON_DESC VARCHAR(2000),
        REMARK VARCHAR(2000),
        SOURCE VARCHAR(40),
        SOURCE_DOCUMENT VARCHAR(40),
        STRUCKOFF_DATE TIMESTAMP(0),
        SW_APPLICATION_ID VARCHAR(40),
        USE_RULE VARCHAR(2000),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")