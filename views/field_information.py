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

# connection = oracledb.connect(
#     user = os.environ.get("DB_USER"),
#     password = os.environ.get("DB_PASS"),
#     dsn = os.environ.get("DB_URL")
# )

# print("Successfully connected to Oracle Database")

cursor = connection.cursor()


cursor.execute("""
    begin
        execute immediate 'drop view field_information';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")
    

cursor.execute("""
    create view field_information as select field.area_id, field.area_type, 
    field.field_name, field.discovery_date, field.field_type, ba_alias.password,
    rm_physical_item.digital_size, rm_physical_item.digital_size_uom, 
    ra_media_type.media_type, rm_data_store.data_store_name, anl_accuracy.remark,
    ppdm_quality_control.qc_status, ppdm_quality_control.checked_by_ba_id 
    from field join ba_alias on field.source = ba_alias.source join rm_physical_item on 
    ba_alias.source = rm_physical_item.source join ra_media_type on rm_physical_item.source
    = ra_media_type.source join rm_data_store on ra_media_type.source = rm_data_store.source
    join anl_accuracy on rm_data_store.source = anl_accuracy.source
    join ppdm_quality_control on anl_accuracy.source = ppdm_quality_control.source""")
