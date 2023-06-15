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

# dotenv = Path("../config/conf/.env")
# load_dotenv(dotenv_path=dotenv)

# connection = oracledb.connect(
#     user = os.environ.get("DB_USER"),
#     password = os.environ.get("DB_PASS"),
#     dsn = os.environ.get("DB_URL")
# )


# print("Successfully connected to Oracle Database")

cursor = connection.cursor()


cursor.execute("""
    begin
        execute immediate 'drop view basin';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute("""
    create view basin as select project.project_name, strat_unit.strat_name_set_id, 
    r_strat_status.strat_status, resent_product.product_type, reserve_class.reserve_class_id, 
    resent_vol_summary.open_balance, resent_vol_summary.open_balance_ouom, ra_lr_size_type.size_type,
    land_size.gross_size, r_strat_type.strat_type, r_fault_type.fault_type, r_source.source, ppdm_quality_control.checked_by_ba_id,
    ppdm_quality_control.qc_status from project join strat_unit on project.source = strat_unit.source
    join r_strat_status on strat_unit.source = r_strat_status.source join resent_product on r_strat_status.source
    = resent_product.source join reserve_class on resent_product.source = reserve_class.source join 
    resent_vol_summary on reserve_class.source = resent_vol_summary.source join ra_lr_size_type
    on resent_vol_summary.source = ra_lr_size_type.source join land_size on ra_lr_size_type.source = land_size.source
    join r_strat_type on land_size.source = r_strat_type.source join r_fault_type on r_strat_type.source = r_fault_type.source
    join r_source on r_fault_type.source = r_source.source join ppdm_quality_control on r_source = ppdm_quality_control.source""")