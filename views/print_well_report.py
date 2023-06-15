import oracledb
import sys
import lib_platform as platform

if platform.system == "darwin":
    sys.path.insert(0, "database")
elif platform.system == "windows":
    sys.path.insert(0, "../database")

from conn import connection


# import os
# from pathlib import Path
# from dotenv import load_dotenv

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
        execute immediate 'drop view print_well_report';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

# cursor.execute("""
#     create view print_well_report as select source_document.publisher, source_document.document_title, source_document.issue, source_doc_author.author_id, source_document.publication_date, source_document.document_type, rm_data_store.data_store_name from source_document, source_doc_author, rm_data_store""")

cursor.execute("""
    CREATE OR REPLACE FORCE EDITIONABLE VIEW "ADMIN"."PRINT_WELL_REPORT" ("PPDM_GUID", "BA_LONG_NAME1", "BA_TYPE1", "AREA_ID", "AREA_TYPE", "FIELD_NAME", "WELL_NAME", "UWI", "TITLE", "CREATOR_NAME", "CREATE_DATE", "MEDIA_TYPE", "DOCUMENT_TYPE", "ITEM_CATEGORY", "ITEM_SUB_CATEGORY", "PAGE_COUNT", "REMARK", "BA_LONG_NAME2", "BA_TYPE2", "DATA_STORE_NAME", "DATA_STORE_TYPE", "SOURCE", "QC_STATUS", "CHECKED_BY_BA_ID") AS 
  SELECT ba.ppdm_guid, ba.ba_long_name AS ba_long_name1, ba.ba_type AS ba_type1, a.area_id, a.area_type, 
f.field_name, w.well_name, w.uwi, rii.title, rc.creator_name, 
rpi.create_date, rmt.media_type, rdt.document_type, ric.item_category, 
risc.item_sub_category, rpi.page_count, aa.remark, ba.ba_long_name AS ba_long_name2,
ba.ba_type AS ba_type2, rds.data_store_name, rds.data_store_type, rs.source, 
pqc.qc_status, pqc.checked_by_ba_id 
FROM business_associate ba
JOIN area a ON ba.ppdm_guid = a.ppdm_guid
JOIN field f ON ba.ppdm_guid = f.ppdm_guid 
JOIN well w ON ba.ppdm_guid = w.ppdm_guid
JOIN rm_information_item rii ON ba.ppdm_guid = rii.ppdm_guid
JOIN rm_creator rc ON ba.ppdm_guid = rc.ppdm_guid
JOIN r_media_type rmt ON ba.ppdm_guid = rmt.ppdm_guid
JOIN r_document_type rdt ON ba.ppdm_guid = rdt.ppdm_guid
JOIN r_item_category ric ON ba.ppdm_guid = ric.ppdm_guid
JOIN r_item_sub_category risc ON ba.ppdm_guid = risc.ppdm_guid
JOIN rm_physical_item rpi ON ba.ppdm_guid = rpi.ppdm_guid
JOIN anl_accuracy aa ON ba.ppdm_guid = aa.ppdm_guid
JOIN rm_data_store rds ON ba.ppdm_guid = rds.ppdm_guid
JOIN r_source rs ON ba.ppdm_guid = rs.ppdm_guid
JOIN ppdm_quality_control pqc ON ba.ppdm_guid = pqc.ppdm_guid;""")