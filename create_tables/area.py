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
        execute immediate 'drop table area';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table area(
    AREA_ID VARCHAR(40) PRIMARY KEY NOT NULL,
    AREA_TYPE VARCHAR(40) NOT NULL,
    ACTIVE_IND VARCHAR(1),
    AREA_MAX_LATITUDE DECIMAL(14,9),
    AREA_MAX_LONGITUDE DECIMAL(14,9),
    AREA_MIN_LATITUDE DECIMAL(14,9),
    AREA_MIN_LONGITUDE DECIMAL(14,9),
    COORD_ACQUISITION_ID VARCHAR(40),
    COORD_SYSTEM_ID VARCHAR(40),
    EFFECTIVE_DATE TIMESTAMP(0),
    EXPIRY_DATE TIMESTAMP(0),
    LOCAL_COORD_SYSTEM_ID VARCHAR(40),
    PPDM_GUID VARCHAR(38),
    PREFERRED_NAME VARCHAR(255),
    REMARK VARCHAR(240),
    SOURCE VARCHAR(40),
    SOURCE_DOCUMENT_ID VARCHAR(40),
    ROW_CHANGED_BY VARCHAR(30),
    ROW_CHANGED_DATE TIMESTAMP(0),
    ROW_CREATED_BY VARCHAR(30),
    ROW_CREATED_DATE TIMESTAMP(0),
    ROW_EFFECTIVE_DATE TIMESTAMP(0),
    ROW_EXPIRY_DATE TIMESTAMP(0),
    ROW_QUALITY VARCHAR(40))""")