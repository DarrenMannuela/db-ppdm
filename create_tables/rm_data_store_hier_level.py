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
        execute immediate 'drop table rm_data_store_hier_level';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
   create table rm_data_store_hier_level(
        DATA_STORE_HIER_ID VARCHAR(40) NOT NULL,
        HIER_LEVEL_ID VARCHAR(40) PRIMARY KEY NOT NULL,
        ACTIVE_IND VARCHAR(1),
        DATA_STORE_TYPE VARCHAR(40),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        LEVEL_NAME VARCHAR(255),
        LEVEL_SEQ_NO INT,
        PARENT_HIER_LEVEL_ID VARCHAR(40),
        PPDM_GUID VARCHAR(38),
        REMARK VARCHAR(2000),
        SOURCE VARCHAR(40),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")
