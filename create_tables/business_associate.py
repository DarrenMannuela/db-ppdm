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
        execute immediate 'drop table business_associate';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table business_associate(
    BUSINESS_ASSOCIATE_ID VARCHAR(40) PRIMARY KEY NOT NULL,
    ACTIVE_IND VARCHAR(1),
    BA_ABBREVIATION VARCHAR(12),
    BA_CATEGORY VARCHAR(40),
    BA_CODE VARCHAR(40),
    BA_LONG_NAME VARCHAR(255),
    BA_SHORT_NAME VARCHAR(30),
    BA_TYPE VARCHAR(40),
    CREDIT_CHECK_DATE TIMESTAMP(0),
    CREDIT_CHECK_IND VARCHAR(1),
    CREDIT_CHECK_SOURCE VARCHAR(40),
    CREDIT_RATING VARCHAR(40),
    CREDIT_RATING_SOURCE VARCHAR(40),
    CURRENT_STATUS VARCHAR(40),
    EFFECTIVE_DATE TIMESTAMP(0),
    EXPIRY_DATE TIMESTAMP(0),
    FIRST_NAME VARCHAR(255),
    LAST_NAME VARCHAR(255),
    MIDDLE_INITIAL VARCHAR(30),
    PPDM_GUID VARCHAR(38),
    REMARK VARCHAR(2000),
    SERVICE_PERIOD VARCHAR(240),
    SOURCE VARCHAR(40),
    ROW_CHANGED_BY VARCHAR(30),
    ROW_CHANGED_DATE TIMESTAMP(0),
    ROW_CREATED_BY VARCHAR(30),
    ROW_CREATED_DATE TIMESTAMP(0),
    ROW_EFFECTIVE_DATE TIMESTAMP(0),
    ROW_EXPIRY_DATE TIMESTAMP(0),
    ROW_QUALITY VARCHAR(40))""")