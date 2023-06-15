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
        execute immediate 'drop table resent_vol_summary';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create table resent_vol_summary (
        RESENT_ID VARCHAR(40) PRIMARY KEY NOT NULL,
        RESERVE_CLASS_ID VARCHAR(40) NOT NULL,
        PRODUCT_TYPE VARCHAR(40) NOT NULL,
        SUMMARY_ID VARCHAR(40) NOT NULL,
        SUMMARY_OBS_NO INT,
        ACTIVE_IND VARCHAR(1),
        ANALYST_BA_ID VARCHAR(40),
        APPROVED_BY_BA_ID VARCHAR(40),
        APPROVED_DATE TIMESTAMP(0),
        CREATED_DATE TIMESTAMP(0),
        CREATED_DATE_DESC VARCHAR(8),
        CURRENT_BALANCE DECIMAL(14,4),
        CURRENT_BALANCE_OUOM VARCHAR(40),
        CURRENT_BALANCE_UOM VARCHAR(40),
        DECLINE_CASE_ID VARCHAR(40),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        GROSS_SUMMARY_IND VARCHAR(1),
        INTEREST_SET_ID VARCHAR(40),
        INTEREST_SET_PARTNER VARCHAR(40),
        INTEREST_SET_SEQ_NO INT,
        MATERIAL_BALANCE_CASE_ID VARCHAR(40),
        NET_SUMMARY_IND VARCHAR(1),
        OPEN_BALANCE DECIMAL(14,4),
        OPEN_BALANCE_OUOM VARCHAR(40),
        OPEN_BALANCE_UOM VARCHAR(40),
        PARTNER_OBS_NO INT,
        PDEN_ID VARCHAR(40),
        PDEN_PRODUCT_TYPE VARCHAR(40),
        PDEN_SOURCE VARCHAR(40),
        PDEN_SUBTYPE VARCHAR(30),
        PPDM_GUID VARCHAR(38),
        REMARK VARCHAR(2000),
        REPORT_IND VARCHAR(1),
        SOURCE VARCHAR(40),
        VOLUME_METHOD VARCHAR(40),
        VOL_ANAL_CASE_ID VARCHAR(40),
        YIELD_PARENT_PRODUCT VARCHAR(40),
        YIELD_RATE DECIMAL(14,4),
        YIELD_RATE_OUOM VARCHAR(40),
        YIELD_RATE_UOM VARCHAR(40),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")
