import re
import json
import numpy as np
import os

folder = "table_from_pdf"

all_tablename = []
for names in os.listdir(folder):
    remove_json = names.strip('.json').strip()
    all_tablename.append(remove_json)

def get_referenced_table(table_name, visited=None):
    if visited is None:
        visited = set()
    if table_name in visited:
        return None
    visited.add(table_name)

    table = []
    with open(f'table_from_pdf/{table_name}.json', 'r') as f:
        cur_table = json.load(f)

    get_ref = list(cur_table[table_name]['Ref Table(s)'].values())

    del_ref = [ref for ref in get_ref if isinstance(ref, str)]

    for ref in del_ref:
        split = ref.split('\r')
        for item in split:
            if item in all_tablename:
                table.append(item)
            elif item == 'PPDM_MEASUREMENT_SYST':
                table.append('PPDM_MEASUREMENT_SYSTEM')

    referenced_tables = []
    for referenced_table in table:
        if referenced_table in visited:
            referenced_tables.append(referenced_table)
        else:
            sub_table = get_referenced_table(referenced_table, visited)
            referenced_tables.append(sub_table)

    return {table_name: referenced_tables}

tablenames = os.listdir('table_from_pdf')
# for names in tablenames:
#     remove_json = names.strip('.json')
#     print(get_referenced_table(remove_json))
# # print(get_referenced_table("ANL_ANALYSIS_REPORT"))

print(get_referenced_table('ANL_ACCURACY'))










# import json
# import numpy as np
# import os


# def get_referenced_table(table_name, visited=None):
#     if visited is None:
#         visited = set()
#     if table_name in visited:
#         return None
#     visited.add(table_name)

#     table = []
#     with open(f'table_from_pdf/{table_name}.json', 'r') as f:
#         cur_table = json.load(f)

#     get_ref = list(cur_table[table_name]['Ref Table(s)'].values())

#     del_ref = [ref for ref in get_ref if isinstance(ref, str)]

#     for ref in del_ref:
#         if '\r' in ref:
#             for split in ref.split('\r'):
#                 table.append(split)
#         else:
#             table.append(ref)

#     referenced_tables = []
#     for referenced_table in table:
#         if referenced_table in visited:
#             referenced_tables.append(referenced_table)
#         else:
#             sub_table = get_referenced_table(referenced_table, visited)
#             referenced_tables.append(sub_table)

#     return {table_name: referenced_tables}

# # tablenames = os.listdir('table_from_pdf')
# # for names in tablenames:
# #     remove_json = names.strip('.json')
#     # print(get_referenced_table(remove_json))
# print(get_referenced_table("SOURCE_DOCUMENT"))



