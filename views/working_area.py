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
        execute immediate 'drop view working_area';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create view working_area as select area.area_id, area.area_type, 
    business_associate.ba_long_name, business_associate.ba_type, 
    business_associate.effective_date, business_associate.expiry_date, 
    cont_type.contract_type, land_right.granted_right_type, spatial_description.gross_size, 
    spatial_description.gross_size_ouom, land_right.land_right_category, land_right.producing_ind, 
    rm_physical_item.original_file_name, ba_alias.password, rm_physical_item.digital_size, 
    rm_physical_item.digital_size_uom, ra_media_type.media_type, rm_data_store.data_store_name, 
    r_source.source, ppdm_quality_control.qc_status, ppdm_quality_control.checked_by_ba_id 
    from area join business_associate on area.source = business_associate.source join cont_type
    on business_associate.source = cont_type.source join land_right on cont_type.source = land_right.source
    join spatial_description on land_right.source = spatial_description.source join rm_physical_item
    on spatial_description.source = rm_physical_item.source join ba_alias on rm_physical_item.source = ba_alias.source
    join ra_media_type on ba_alias.source = ra_media_type.source join rm_data_store on ra_media_type.source = rm_data_store.source
    join r_source on rm_data_store.source = r_source.source join ppdm_quality_control on r_source.source = ppdm_quality_control.source""")