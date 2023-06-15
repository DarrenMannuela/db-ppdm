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
        execute immediate 'drop table ppdm_quality_control';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create table ppdm_quality_control (
        SYSTEM_ID VARCHAR(40) PRIMARY KEY NOT NULL,
        TABLE_NAME VARCHAR(30) NOT NULL,
        QC_SEQ_NO INT,
        ACTIVE_IND VARCHAR(1),
        CHECKED_BY_BA_ID VARCHAR(40),
        COLUMN_NAME VARCHAR(30),
        CURRENT_DATE_VALUE TIMESTAMP(0),
        CURRENT_NUMERIC_VALUE DOUBLE PRECISION,
        CURRENT_NUMERIC_VALUE_OUOM VARCHAR(40),
        CURRENT_NUMERIC_VALUE_UOM VARCHAR(40),
        CURRENT_ROW_GUID VARCHAR(38),
        CURRENT_TEXT_VALUE VARCHAR(2000),
        DATA_TYPE VARCHAR(40),
        DONE_BY_BA_ID VARCHAR(40),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        NULL_DESCRIPTION VARCHAR(1024),
        PPDM_GUID VARCHAR(38),
        QC_DATETIME TIMESTAMP(0),
        QC_DESC VARCHAR(1024),
        QC_STATUS VARCHAR(40),
        QC_TYPE VARCHAR(40),
        QUALITY_TYPE VARCHAR(40),
        REMARK VARCHAR(2000),
        RETENTION_PERIOD VARCHAR(40),
        SOURCE VARCHAR(40),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")
