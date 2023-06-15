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
        execute immediate 'drop table land_size';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create table land_size(
        LAND_RIGHT_SUBTYPE VARCHAR(30) NOT NULL,
        LAND_RIGHT_ID VARCHAR(40) PRIMARY KEY NOT NULL,
        SIZE_TYPE VARCHAR(40) NOT NULL,
        SIZE_SEQ_NO INT,
        ACTIVE_IND VARCHAR(1),
        BUSINESS_ASSOCIATE_ID VARCHAR(40),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        GROSS_SIZE DECIMAL(20,10),
        INTEREST_SET_ID VARCHAR(40),
        INTEREST_SET_SEQ_NO INT,
        NET_SIZE DECIMAL(20,10),
        PARTNER_OBS_NO INT,
        PPDM_GUID VARCHAR(38),
        REMARK VARCHAR(2000),
        SIZE_OUOM VARCHAR(40),
        SOURCE VARCHAR(40),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")
