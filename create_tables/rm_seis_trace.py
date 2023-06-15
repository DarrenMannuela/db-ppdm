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
        execute immediate 'drop table rm_seis_trace';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table rm_seis_trace(
        INFO_ITEM_SUBTYPE VARCHAR(30) NOT NULL,
        INFORMATION_ITEM_ID VARCHAR(40) NOT NULL,
        ACTIVE_IND VARCHAR(1),
        DATUM_VEL_CORRECTION VARCHAR(240),
        DATUM_VEL_CORRECTION_OUOM VARCHAR(40),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        FIRST_SAMPLE_TIME DECIMAL(10,5),
        FIRST_SAMPLE_TIMEZONE VARCHAR(40),
        HORIZONTAL_SCALE VARCHAR(40),
        HORIZONTAL_SCALE_OUOM VARCHAR(40),
        PHASE VARCHAR(40),
        POLARITY VARCHAR(40),
        PPDM_GUID VARCHAR(38),
        RECORD_LENGTH DECIMAL(10,5),
        RECORD_LENGTH_OUOM VARCHAR(40),
        REFERENCE_DATUM DECIMAL(10,5),
        REFERENCE_DATUM_OUOM VARCHAR(40),
        REMARK VARCHAR(2000),
        REPLACEMENT_VEL DECIMAL(10,5),
        REPLACEMENT_VEL_OUOM VARCHAR(40),
        REPORTED_FIRST_REF_NUM VARCHAR(30),
        REPORTED_LAST_REF_NUM VARCHAR(30),
        REPORTED_REF_NUM_TYPE VARCHAR(40),
        SAMPLE_RATE DECIMAL(8,4),
        SAMPLE_RATE_OUOM VARCHAR(40),
        SAMPLE_TYPE VARCHAR(40),
        SOURCE VARCHAR(40),
        STATIC_CORRECTION DECIMAL(10,5),
        STATIC_CORRECTION_OUOM VARCHAR(40),
        VARIABLE_AREA_IND VARCHAR(1),
        VERTICAL_SCALE VARCHAR(40),
        VERTICAL_SCALE_OUOM VARCHAR(40),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")