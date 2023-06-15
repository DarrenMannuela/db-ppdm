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
        execute immediate 'drop table users';
        exception when others then if sqlcode <> -942 then raise; end if;
    end;""")

cursor.execute(
    """create table users(
        ID INT GENERATED ALWAYS AS IDENTITY(START WITH 1 INCREMENT BY 1) PRIMARY KEY NOT NULL,
        FIRST_NAME VARCHAR(255) NOT NULL,
        LAST_NAME VARCHAR(255) NOT NULL,
        EMAIL VARCHAR(255) UNIQUE NOT NULL,
        PASSWORD VARCHAR(255) NOT NULL,
        DATE_JOINED DATE DEFAULT SYSDATE NOT NULL,
        EXPIRED_DATE DATE NOT NULL,
        ROLE_NAME VARCHAR(255) NOT NULL,
        PROFILE_PICTURE CLOB, 
        CONSTRAINT FK_CUSTOMER_ROLE FOREIGN KEY (ROLE_NAME) REFERENCES ROLES(ROLE_NAME)
    )""")