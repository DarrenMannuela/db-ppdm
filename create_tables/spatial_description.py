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
        execute immediate 'drop table spatial_description';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table spatial_description(
    SPATIAL_DESCRIPTION_ID VARCHAR(40) PRIMARY KEY NOT NULL,
    SPATIAL_OBS_NO INT,
    ACTIVE_IND VARCHAR(1),
    CARTER_IND VARCHAR(1),
    CONGRESS_IND VARCHAR(1),
    COORD_ACQUISITION_ID VARCHAR(40),
    COORD_SYSTEM_ID VARCHAR(40),
    DLS_IND VARCHAR(1),
    EFFECTIVE_DATE TIMESTAMP(0),
    EXPIRY_DATE TIMESTAMP(0),
    FPS_IND VARCHAR(1),
    GEODETIC_IND VARCHAR(1),
    GROSS_SIZE DECIMAL(20,10),
    GROSS_SIZE_OUOM VARCHAR(40),
    INACTIVATION_DATE TIMESTAMP(0),
    LIBYA_IND VARCHAR(1),
    LINE_IND VARCHAR(1),
    LINE_VERSION_IND VARCHAR(1),
    LOCAL_COORD_SYSTEM_ID VARCHAR(40),
    MAX_LATITUDE DECIMAL(14,9),
    MAX_LONGITUDE DECIMAL(14,9),
    MINERAL_ZONE_IND VARCHAR(1),
    MIN_LATITUDE DECIMAL(14,9),
    MIN_LONGITUDE DECIMAL(14,9),
    NE_LOC_IND VARCHAR(1),
    NORTH_SEA_IND VARCHAR(1),
    NTS_IND VARCHAR(1),
    OFFSHORE_IND VARCHAR(1),
    OHIO_IND VARCHAR(1),
    PBL_IND VARCHAR(1),
    POINT_IND VARCHAR(1),
    POINT_VERSION_IND VARCHAR(1),
    POLYGON_IND VARCHAR(1),
    POLYGON_VERSION_IND VARCHAR(1),
    PPDM_GUID VARCHAR(38),
    REFERENCE_NUM VARCHAR(30),
    REMARK VARCHAR(2000),
    SOURCE VARCHAR(40),
    SPATIAL_DESC_TEXT_IND VARCHAR(1),
    SPATIAL_DESC_TYPE VARCHAR(40),
    TEXAS_IND VARCHAR(1),
    ROW_CHANGED_BY VARCHAR(30),
    ROW_CHANGED_DATE TIMESTAMP(0),
    ROW_CREATED_BY VARCHAR(30),
    ROW_CREATED_DATE TIMESTAMP(0),
    ROW_EFFECTIVE_DATE TIMESTAMP(0),
    ROW_EXPIRY_DATE TIMESTAMP(0),
    ROW_QUALITY VARCHAR(40))""")