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
        execute immediate 'drop table seis_activity';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table seis_activity(
        SEIS_SET_SUBTYPE VARCHAR(30) NOT NULL,
        SEIS_SET_ID VARCHAR(40) NOT NULL,
        ACTIVITY_OBS_NO INT,
        ACTIVE_IND VARCHAR(1),
        ACTIVITY_DURATION DECIMAL(10,5),
        ACTIVITY_DURATION_OUOM VARCHAR(40),
        AREA_SIZE DECIMAL(20,10),
        AREA_SIZE_OUOM VARCHAR(40),
        CREW_COMPANY_BA_ID VARCHAR(40),
        CREW_ID VARCHAR(40),
        EFFECTIVE_DATE TIMESTAMP(0),
        END_DATE TIMESTAMP(0),
        END_TIME TIMESTAMP(0),
        END_TIMEZONE VARCHAR(40),
        EXPIRY_DATE TIMESTAMP(0),
        FIRST_NLINE_NO INT,
        FIRST_SEIS_POINT_ID VARCHAR(16),
        FIRST_XLINE_NO INT,
        LAST_NLINE_NO INT,
        LAST_SEIS_POINT_ID VARCHAR(16),
        LAST_XLINE_NO INT,
        LINE_LENGTH DECIMAL(10,3),
        LINE_LENGTH_OUOM VARCHAR(40),
        OWNER_BA_ID VARCHAR(40),
        PPDM_GUID VARCHAR(38),
        REMARK VARCHAR(2000),
        SEIS_ACTIVITY_CLASS VARCHAR(40),
        SEIS_ACTIVITY_TYPE VARCHAR(40),
        SOURCE VARCHAR(40),
        START_DATE TIMESTAMP(0),
        START_TIME TIMESTAMP(0),
        START_TIMEZONE VARCHAR(40),
        TOTAL_TIME_ELAPSED DECIMAL(10,5),
        TOTAL_TIME_ELAPSED_OUOM VARCHAR(40),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")