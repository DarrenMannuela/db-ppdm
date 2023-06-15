import unittest
import oracledb
import sys
import lib_platform as platform
import time
import csv


if platform.system == "darwin":
    sys.path.insert(0, "database")
    with open('test_data/r_source_data.csv', mode ='r')as file:
        test_file = csv.reader(file)
        next(test_file)
        test_data = [tuple(row) for row in test_file]
elif platform.system == "windows":
    sys.path.insert(0, "../database")
    with open('test_data/r_source_data.csv', mode ='r')as file:
        test_file = csv.reader(file)
        next(test_file)
        test_data = [tuple(row) for row in test_file]

from conn import connection

class TestRSource(unittest.TestCase):

    def test_insert_r_source(self):
        for data in range (0, len(test_data)):

            try:
                cursor = connection.cursor()
                query = "INSERT INTO R_SOURCE VALUES(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)"
                val = test_data[data]
                start = time.time()
                cursor.execute(query, val)
                connection.commit()
                end = time.time()
            
            except oracledb.IntegrityError as e:
                print(e)

    def test_select_r_source(self):
        primary_key  = 'A1'
        cursor = connection.cursor()
        query = "SELECT * FROM R_SOURCE WHERE SOURCE = 'A1'"
        cursor.execute(query)
        start = time.time()
        cursor.execute(query)
        end = time.time()
        get_key = cursor.fetchone()[0]
        self.assertEqual(primary_key, get_key, "Data doesnt exist")

    def test_alter_r_source(self):
        try:
            cursor = connection.cursor()
            query = "UPDATE R_SOURCE SET SOURCE = 'B1' WHERE SOURCE = 'A5'"
            cursor.execute(query)
            connection.commit()
        except oracledb.Error as e:
            print(e)

    def test_delete_r_source(self):
        for data in range (0, len(test_data)):
            try:
                cursor = connection.cursor()
                query = "DELETE FROM R_SOURCE WHERE SOURCE = :val"
                val = {'val':test_data[data][0]}
                cursor.execute(query, val)
                connection.commit()
                #self.assertEqual(primary_key, get_key, "Data doesnt exist")
            except oracledb.Error as e:
                print(e)





unittest.main()