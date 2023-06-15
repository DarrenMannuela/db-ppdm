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
        execute immediate 'drop table anl_accuracy';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table anl_accuracy(
    ANALYSIS_ID VARCHAR(40) PRIMARY KEY NOT NULL,
    ANL_SOURCE VARCHAR(40) NOT NULL,
    ACCURACY_OBS_NO INT,
    ACCURACY_DESC VARCHAR(240),
    ACCURACY_TYPE VARCHAR(40),
    ACTIVE_IND VARCHAR(1),
    CALCULATED_IND VARCHAR(1),
    CALCULATE_METHOD_ID VARCHAR(40),
    CONFIDENCE_TYPE VARCHAR(40),
    EFFECTIVE_DATE TIMESTAMP(0),
    EXPIRY_DATE TIMESTAMP(0),
    MEASURED_VALUE_IND VARCHAR(1),
    PPDM_GUID VARCHAR(38),
    RAW_VALUE_IND VARCHAR(1),
    REFERENCED_COLUMN_NAME VARCHAR(30),
    REFERENCED_PPDM_GUID VARCHAR(38),
    REFERENCED_SYSTEM_ID VARCHAR(40),
    REFERENCED_TABLE_NAME VARCHAR(30),
    REMARK VARCHAR(2000),
    REPEATABILITY_TYPE VARCHAR(40),
    REPORTED_IND VARCHAR(1),
    SOURCE VARCHAR(40),
    STEP_SEQ_NO INT,
    ROW_CHANGED_BY VARCHAR(30),
    ROW_CHANGED_DATE TIMESTAMP(0),
    ROW_CREATED_BY VARCHAR(30),
    ROW_CREATED_DATE TIMESTAMP(0),
    ROW_EFFECTIVE_DATE TIMESTAMP(0),
    ROW_EXPIRY_DATE TIMESTAMP(0),
    ROW_QUALITY VARCHAR(40))""")