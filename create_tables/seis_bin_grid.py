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
        execute immediate 'drop table seis_bin_grid';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table seis_bin_grid(
        SEIS_SET_SUBTYPE VARCHAR(30) NOT NULL,
        SEIS_SET_ID VARCHAR(40) NOT NULL,
        BIN_GRID_ID VARCHAR(40) NOT NULL,
        SOURCE VARCHAR(40) NOT NULL,
        ACTIVE_IND VARCHAR(1),
        ANGLE_ROTATION DECIMAL(10,5),
        BIN_GRID_TYPE VARCHAR(40),
        BIN_GRID_VERSION DECIMAL(5,2),
        BIN_METHOD VARCHAR(40),
        BIN_POINT_OUOM VARCHAR(40),
        COORD_ACQUISITION_ID VARCHAR(40),
        COORD_SYSTEM_ID VARCHAR(40),
        CORNER1_LAT DECIMAL(14,9),
        CORNER1_LONG DECIMAL(14,9),
        CORNER2_LAT DECIMAL(14,9),
        CORNER2_LONG DECIMAL(14,9),
        CORNER3_LAT DECIMAL(14,9),
        CORNER3_LONG DECIMAL(14,9),
        EFFECTIVE_DATE TIMESTAMP(0),
        EXPIRY_DATE TIMESTAMP(0),
        FOLD_COVERAGE DECIMAL(10,5),
        LOCAL_COORD_SYSTEM_ID VARCHAR(40),
        NLINE_AZIMUTH DECIMAL(10,5),
        NLINE_COUNT INT,
        NLINE_MAX_NO INT,
        NLINE_MIN_NO INT,
        NLINE_SPACING DECIMAL(10,3),
        NORTH_TYPE VARCHAR(40),
        POINT_ORIGIN_EASTING DECIMAL(15,7),
        POINT_ORIGIN_LATITUDE DECIMAL(14,9),
        POINT_ORIGIN_LONGITUDE DECIMAL(14,9),
        POINT_ORIGIN_NORTHING DECIMAL(15,7),
        PPDM_GUID VARCHAR(38),
        REMARK VARCHAR(2000),
        XLINE_AZIMUTH DECIMAL(10,5),
        XLINE_COUNT INT,
        XLINE_MAX_NO INT,
        XLINE_MIN_NO INT,
        XLINE_SPACING DECIMAL(10,3),
        ROW_CHANGED_BY VARCHAR(30),
        ROW_CHANGED_DATE TIMESTAMP(0),
        ROW_CREATED_BY VARCHAR(30),
        ROW_CREATED_DATE TIMESTAMP(0),
        ROW_EFFECTIVE_DATE TIMESTAMP(0),
        ROW_EXPIRY_DATE TIMESTAMP(0),
        ROW_QUALITY VARCHAR(40))""")