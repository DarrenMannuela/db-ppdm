import unittest
import oracledb
import sys
import lib_platform as platform
import time

if platform.system == "darwin":
    sys.path.insert(0, "database")
elif platform.system == "windows":
    sys.path.insert(0, "../database")

from conn import connection

class TestBaAlias(unittest.TestCase):
    # def test_insert_field(self):
    #     cursor = connection.cursor()
    #     query = "INSERT INTO FIELD (field_id, active_ind, area_id, area_type, discovery_date, effective_date, expiry_date, field_abbreviation, field_area_obs_no, field_name, field_type, group_field_id, ppdm_guid, remark, source, row_changed_by, row_changed_date, row_created_by, row_created_date, row_effective_date, row_expiry_date, row_quality) VALUES(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)"
    #     val = ("A2", None, None, None, None, None, None, None, None, None, None,  None, None, None, None, None, None, None, None, None, None, None)
    #     cursor.execute(query, val)
    #     connection.commit()
    #     print(cursor.rowcount, "record(s) inserted")

    def test_insert_ba_alias(self):
        try:
            cursor = connection.cursor()
            query = "INSERT INTO BA_ALIAS VALUES(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)"
            val = ('A2', 'B1', 'C1', None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None)
            cursor.execute(query, val)
            start = time.time()
            connection.commit()
            end = time.time()
            print(end-start)
        except oracledb.IntegrityError as e:
            self.assertIn('ORA-00001', str(e), "Primary key Restriction Failed")

    def test_select_ba_alias(self):
        primary_key  = 'A1'
        cursor = connection.cursor()
        query = "SELECT * FROM BA_ALIAS WHERE BUSINESS_ASSOCIATE_ID = 'A1'"
        cursor.execute(query)
        start = time.time()
        cursor.execute(query)
        end = time.time()
        get_key = cursor.fetchone()[0]
        print(end-start)
        self.assertEqual(primary_key, get_key, "Data doesnt exist")

    def test_alter_ba_alias(self):
        try:
            cursor = connection.cursor()
            query = "UPDATE BA_ALIAS SET BUSINESS_ASSOCIATE_ID = 'B1' WHERE BUSINESS_ASSOCIATE_ID = 'A1'"
            cursor.execute(query)
            connection.commit()
        except oracledb.Error as e:
            print(e)

    def test_delete_ba_alias(self):
        try:
            cursor = connection.cursor()
            query = "DELETE FROM BA_ALIAS WHERE BUSINESS_ASSOCIATE_ID = 'A2'"
            cursor.execute(query)
            connection.commit()
            #self.assertEqual(primary_key, get_key, "Data doesnt exist")
        except oracledb.Error as e:
            print(e)
            
        cursor.close()
        connection.close()


    




unittest.main()