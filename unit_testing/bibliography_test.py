import unittest
import sys
import lib_platform as platform

if platform.system == "darwin":
    sys.path.insert(0, "database")
elif platform.system == "windows":
    sys.path.insert(0, "../database")

from conn import connection

class TestField(unittest.TestCase):
    # def test_insert_field(self):
    #     cursor = connection.cursor()
    #     query = "INSERT INTO FIELD (field_id, active_ind, area_id, area_type, discovery_date, effective_date, expiry_date, field_abbreviation, field_area_obs_no, field_name, field_type, group_field_id, ppdm_guid, remark, source, row_changed_by, row_changed_date, row_created_by, row_created_date, row_effective_date, row_expiry_date, row_quality) VALUES(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)"
    #     val = ("A2", None, None, None, None, None, None, None, None, None, None,  None, None, None, None, None, None, None, None, None, None, None)
    #     cursor.execute(query, val)
    #     connection.commit()
    #     print(cursor.rowcount, "record(s) inserted")

    # def test_insert_source_document(self):
    #         cursor = connection.cursor()
    #         query = "INSERT INTO SOURCE_DOCUMENT VALUES(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)"
    #         val = ('A1', None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None)
    #         cursor.execute(query, val)
    #         connection.commit()
    #         print(cursor.rowcount, "record(s) inserted")

    def test_select_bibliography(self):
        cursor = connection.cursor()
        query = "SELECT * FROM BIBLIOGRAPHY"
        cursor.execute(query)
        for row in cursor:
            print(row)
        cursor.close()
        connection.close()



unittest.main()
