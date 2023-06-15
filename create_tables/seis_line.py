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
        execute immediate 'drop table seis_line';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table seis_line(
        SEIS_SET_SUBTYPE VARCHAR(30) NOT NULL,
        SEIS_LINE_ID VARCHAR(40) NOT NULL,
        ACTIVE_IND VARCHAR(1),
        DIMENSION VARCHAR(40),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        LINE_NAME VARCHAR(255),
        PPDM_GUID VARCHAR(38),
        REFRACTION_REFLECTION VARCHAR(40),
        REMARK VARCHAR(2000),
        RESHOOT_OF_SET_ID VARCHAR(40),
        RESHOOT_SEIS_SET_SUBTYPE VARCHAR(30),
        SEIS_ACQTN_SET_SUBTYPE VARCHAR(30),
        SEIS_ACQTN_SURVEY_ID VARCHAR(40),
        SEIS_SPECTRUM_TYPE VARCHAR(40),
        SOURCE VARCHAR(40),
        TEST_EXPERIMENTAL VARCHAR(40),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")