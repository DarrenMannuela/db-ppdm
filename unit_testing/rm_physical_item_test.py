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

class TestRmPhysicalItem(unittest.TestCase):
    def test_insert_rm_physical_item(self):
        try:
            cursor = connection.cursor()
            query = "INSERT INTO RM_PHYSICAL_ITEM VALUES(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55)"
            val = ('A1', None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None)
            start = time.time()
            cursor.execute(query, val)
            connection.commit()
            end = time.time()
            print(end-start)
        
        except oracledb.IntegrityError as e:
            self.assertIn('ORA-00001', str(e), "Primary key Restriction Failed")

    def test_select_rm_physical_item(self):
        primary_key = 'A1'
        cursor = connection.cursor()
        query = "SELECT * FROM RM_PHYSICAL_ITEM WHERE PHYSICAL_ITEM_ID = 'A1'"
        cursor.execute(query)
        start = time.time()
        cursor.execute(query)
        end = time.time()
        get_key = cursor.fetchone()[0]
        print(end-start)
        self.assertEqual(primary_key, get_key, "Data doesnt exist")

    def test_alter_rm_physical_item(self):
        try:
            cursor = connection.cursor()
            query = "UPDATE RM_PHYSICAL_ITEM SET PHYSICAL_ITEM_ID = 'B1' WHERE PHYSICAL_ITEM_ID = 'A1'"
            cursor.execute(query)
            connection.commit()
        except oracledb.Error as e:
            print(e)

    def test_delete_rm_physical_item(self):
        try:
            cursor = connection.cursor()
            query = "DELETE FROM RM_PHYSICAL_ITEM WHERE PHYSICAL_ITEM_ID = 'A1'"
            cursor.execute(query)
            connection.commit()
            #self.assertEqual(primary_key, get_key, "Data doesnt exist")
        except oracledb.Error as e:
            print(e)



unittest.main()


