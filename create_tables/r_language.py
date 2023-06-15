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
        execute immediate 'drop table r_language';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create table r_language(
        LANGUAGE_ID VARCHAR(40) PRIMARY KEY NOT NULL,
        ABBREVIATION VARCHAR(12),
        ACTIVE_IND VARCHAR(1),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        LONG_NAME VARCHAR(255),
        PPDM_GUID VARCHAR(38),
        REMARK VARCHAR(2000),
        SHORT_NAME VARCHAR(30),
        SOURCE VARCHAR(40),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40),
        CONSTRAINT FK_Language_Source FOREIGN KEY(SOURCE) 
        REFERENCES R_SOURCE(SOURCE),
        CONSTRAINT FK_Language_RowQuality FOREIGN KEY(ROW_QUALITY) 
        REFERENCES R_PPDM_ROW_QUALITY(ROW_QUALITY_ID))""")