import oracledb
import os
from pathlib import Path
from dotenv import load_dotenv
import lib_platform as platform

# connection = oracledb.connect(
#     user="admin",
#     password="ppdm-3.9",
#     dsn="/kei-oracle-test.chkxr21l2l5s.ap-southeast-3.rds.amazonaws.com:1521/ORCL")

if platform.system == "darwin":
    dotenv = Path("config/conf/.env")
    load_dotenv(dotenv_path=dotenv)
elif platform.system == "windows":
    dotenv = Path("../config/conf/.env")
    load_dotenv(dotenv_path=dotenv)

user = os.environ.get("DB_USER")
password = os.environ.get("DB_PASS")
host = os.environ.get("DB_URL")
port = os.environ.get("DB_PORT")
service_name = os.environ.get("DB_DATABASE")
dsn = f'{user}/{password}@{host}:{port}/{service_name}'

connection = oracledb.connect(dsn)

print("Successfully connected to Oracle Database")
