import sys
import lib_platform as platform

if platform.system == "darwin":
    sys.path.insert(0, "database")
elif platform.system == "windows":
    sys.path.insert(0, "../database")

from conn import connection

cursor = connection.cursor()

cursor.execute("""
    begin
        execute immediate 'drop table roles';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table roles(
        ID INT GENERATED ALWAYS AS IDENTITY(START WITH 1 INCREMENT BY 1) PRIMARY KEY NOT NULL,
        ROLE_NAME VARCHAR(255) UNIQUE NOT NULL
    )""")


