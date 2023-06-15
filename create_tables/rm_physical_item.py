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
        execute immediate 'drop table rm_physical_item';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table rm_physical_item(
    PHYSICAL_ITEM_ID VARCHAR(40) PRIMARY KEY NOT NULL,
    ACTIVE_IND VARCHAR(1),
    BAR_CODE VARCHAR(240),
    CATALOGUE_BY_NAME VARCHAR(255),
    CATALOGUE_DATE TIMESTAMP(0),
    CERTIFIED_TRUE_COPY_IND VARCHAR(1),
    CIRCULATION_ALLOWED_IND VARCHAR(1),
    CIRCULATION_OUT_IND VARCHAR(1),
    COLOR_FORMAT VARCHAR(40),
    CREATE_DATE TIMESTAMP(0),
    DIGITAL_FORMAT VARCHAR(40),
    DIGITAL_SIZE DOUBLE PRECISION,
    DIGITAL_SIZE_UOM VARCHAR(40),
    DIM_HEIGHT DECIMAL(10,5),
    DIM_HEIGHT_UOM VARCHAR(40),
    DIM_LENGTH DECIMAL(10,5),
    DIM_LENGTH_UOM VARCHAR(40),
    DIM_WEIGHT DECIMAL(10,5),
    DIM_WEIGHT_UOM VARCHAR(40),
    DIM_WIDTH DECIMAL(10,5),
    DIM_WIDTH_UOM VARCHAR(40),
    EFFECTIVE_DATE TIMESTAMP(0),
    EXPIRY_DATE TIMESTAMP(0),
    FILE_COUNT INT,
    GROUP_IND VARCHAR(1),
    IMAGE_RESOLUTION_UOM VARCHAR(40),
    IMAGE_X_RESOLUTION DECIMAL(10,5),
    IMAGE_Y_RESOLUTION DECIMAL(10,5),
    ITEM_CATEGORY VARCHAR(40),
    ITEM_SUB_CATEGORY VARCHAR(40),
    LABEL VARCHAR(30),
    LANGUAGE VARCHAR(40),
    LAST_CONDITION VARCHAR(40),
    LOCATION_REFERENCE VARCHAR(240),
    MEDIA_TYPE VARCHAR(40),
    ORIGINAL_FILE_NAME VARCHAR(255),
    ORIGINAL_IND VARCHAR(1),
    OUTPUT_FILE_NAME VARCHAR(255),
    PAGE_COUNT DECIMAL(10,5),
    PHYSICAL_ITEM_STATUS VARCHAR(40),
    PPDM_GUID VARCHAR(38),
    PREFERRED_IND VARCHAR(1),
    QC_BY_NAME VARCHAR(255),
    REMARK VARCHAR(2000),
    RETENTION_PERIOD VARCHAR(40),
    SALE_ALLOWED_IND VARCHAR(1),
    SOURCE VARCHAR(40),
    STORE_ID VARCHAR(40),
    ROW_CHANGED_BY VARCHAR(30),
    ROW_CHANGED_DATE TIMESTAMP(0),
    ROW_CREATED_BY VARCHAR(30),
    ROW_CREATED_DATE TIMESTAMP(0),
    ROW_EFFECTIVE_DATE TIMESTAMP(0),
    ROW_EXPIRY_DATE TIMESTAMP(0),
    ROW_QUALITY VARCHAR(40))""")