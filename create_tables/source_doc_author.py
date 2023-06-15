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
        execute immediate 'drop table source_doc_author';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create table source_doc_author (
        SOURCE_DOCUMENT_ID VARCHAR(40) PRIMARY KEY NOT NULL,
        AUTHOR_ID VARCHAR(40) NOT NULL,
        ACTIVE_IND VARCHAR(1),
        AUTHOR_BA_ID VARCHAR(40),
        AUTHOR_FIRST_NAME VARCHAR(30),
        AUTHOR_INITIAL VARCHAR(30),
        AUTHOR_LAST_NAME VARCHAR(30),
        AUTHOR_SEQ_NO INT,
        AUTHOR_TYPE VARCHAR(40),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
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
