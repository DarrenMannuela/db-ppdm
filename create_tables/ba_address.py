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
        execute immediate 'drop table ba_address';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create table ba_address(
        BUSINESS_ASSOCIATE_ID VARCHAR(40) NOT NULL,
        SOURCE VARCHAR(40) NOT NULL,
        ADDRESS_OBS_NO INT NOT NULL,
        ACTIVE_IND VARCHAR(1),
        ADDRESSEE_TEXT VARCHAR(240),
        ADDRESS_TYPE VARCHAR(40),
        CITY_AREA_ID VARCHAR(40),
        CITY_AREA_TYPE VARCHAR(40),
        COUNTRY_AREA_ID VARCHAR(40),
        COUNTRY_AREA_TYPE VARCHAR(40),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        FIRST_ADDRESS_LINE VARCHAR(1024),
        OFFICE_TYPE VARCHAR(40),
        POSTAL_ZIP_CODE VARCHAR(40),
        PPDM_GUID VARCHAR(38),
        PREFERRED_IND VARCHAR(1),
        PROVINCE_STATE_AREA_ID VARCHAR(40),
        PROVINCE_STATE_AREA_TYPE VARCHAR(40),
        REMARK VARCHAR(2000),
        SECOND_ADDRESS_LINE VARCHAR(1024),
        SERVICE_PERIOD VARCHAR(240),
        SERVICE_QUALITY VARCHAR(40),
        THIRD_ADDRESS_LINE VARCHAR(1024),
        WITHHOLDING_TAX_IND VARCHAR(1),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")