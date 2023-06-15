import subprocess
from pathlib import Path
from dotenv import load_dotenv
import lib_platform as platform
import os


if platform.system == "darwin":
    dotenv = Path("config/conf/.env")
    load_dotenv(dotenv_path=dotenv)
elif platform.system == "windows":
    dotenv = Path("config/conf/.env")
    load_dotenv(dotenv_path=dotenv)


user = os.environ.get("DB_USER")
password = os.environ.get("DB_PASS")
host = os.environ.get("DB_URL")
port = os.environ.get("DB_PORT")
service_name = os.environ.get("DB_DATABASE")
dsn = f'{user}/{password}@/{host}:{port}/{service_name}'

cmd = "sql "+ dsn
sql_cmd = cmd +" @postgres_39/PPDM39_TAB.sql postgres_39/PPDM39_PK.sql postgres_39/PPDM39_CK.sql Vpostgres_39/PPDM39_FK.sql postgres_39/PPDM39_OUOM.sql postgres_39/PPDM39_UOM.sql postgres_39/PPDM39_RQUAL.sql postgres_39/PPDM39_RSRC.sql postgres_39/PPDM39_TCM.sql postgres_39/PPDM39_CCM.sql postgres_39/PPDM39_GUID.sql"

returned_value = subprocess.call(sql_cmd, shell=True)
