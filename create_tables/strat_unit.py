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
        execute immediate 'drop table strat_unit';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create table strat_unit (
        STRAT_NAME_SET_ID VARCHAR(40) PRIMARY KEY NOT NULL,
        STRAT_UNIT_ID VARCHAR(40) NOT NULL,
        ABBREVIATION VARCHAR(12),
        ACTIVE_IND VARCHAR(1),
        AREA_ID VARCHAR(40),
        AREA_TYPE VARCHAR(40),
        BUSINESS_ASSOCIATE_ID VARCHAR(40),
        CONFIDENCE_ID VARCHAR(40),
        CURRENT_STATUS_DATE TIMESTAMP(0),
        DESCRIPTION VARCHAR(240),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        FAULT_TYPE VARCHAR(40),
        FORM_CODE VARCHAR(20),
        GROUP_CODE VARCHAR(20),
        LONG_NAME VARCHAR(255),
        ORDINAL_AGE_CODE DECIMAL(15,5),
        PPDM_GUID VARCHAR(38),
        PREFERRED_IND VARCHAR(1),
        REMARK VARCHAR(2000),
        SHORT_NAME VARCHAR(30),
        SOURCE VARCHAR(40),
        STRAT_INTERPRET_METHOD VARCHAR(40),
        STRAT_STATUS VARCHAR(40),
        STRAT_TYPE VARCHAR(40),
        STRAT_UNIT_TYPE VARCHAR(40),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")
