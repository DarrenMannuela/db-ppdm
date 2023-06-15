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
        execute immediate 'drop view bibliography';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

# cursor.execute("""
#     create view bibliography as select source_document.publisher, source_document.document_title, source_document.issue, source_doc_author.author_id, source_document.publication_date, source_document.document_type, rm_data_store.data_store_name from source_document, source_doc_author, rm_data_store""")

cursor.execute("""
    CREATE OR REPLACE FORCE EDITIONABLE VIEW "ADMIN"."BIBLIOGRAPHY" ("PPDM_GUID", "PUBLISHER", "DOCUMENT_TITLE", "ISSUE", "AUTHOR_ID", "PUBLICATION_DATE", "DOCUMENT_TYPE", "DATA_STORE_NAME") AS 
    SELECT sd.ppdm_guid, sd.publisher, sd.document_title, sd.issue, sda.author_id, sd.publication_date, sd.document_type, rds.data_store_name 
    FROM source_document sd
    LEFT JOIN source_doc_author sda ON sda.ppdm_guid = sd.ppdm_guid 
    LEFT JOIN rm_data_store rds ON sd.ppdm_guid = rds.ppdm_guid""")