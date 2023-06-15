import yaml
import re
import sys
import lib_platform as platform

if platform.system == "darwin":
    sys.path.insert(0, "database")
elif platform.system == "windows":
    sys.path.insert(0, "../database")

from conn import connection
with open('postgres_39/PPDM39_TAB.sql', 'r') as sql_file:
    sql_tables = sql_file.read()

combined_data = {'components': {'schemas': {}}}
list_types = []

def get_sql_str(table_name: str):

    get_table = []
    start_table = re.search(fr'\b {table_name.lower()}\b', sql_tables)
    start_index = sql_tables.index(start_table.group())
    cur_line = ""

    for line in sql_tables[start_index:]:
        if line == "\n":
            get_table.append(cur_line)
            cur_line = ""
        else:
            cur_line += line

        if ";" in line:
            break
    return get_table


def make_dict(table):
    make_yaml = {}
    make_yaml.update({table[0].strip(): {'type': 'array', 'items': {'type': 'object', 'properties': {}}}})

    for line in range (1, len(table)):
        if table[line] == '(':
            pass
        else:
            process_type = ""
            split_stmt = table[line].split(" ")
            if "(" in split_stmt[1]:
                process_type = re.findall(r'\w+(?=\(|$)', split_stmt[1])
                process_type = "".join(process_type)
            else:
                process_type = split_stmt[1]

            if process_type not in list_types:
                list_types.append(process_type)

            if "VARCHAR" in table[line]:
                get_len = re.findall(r"\((\d+)\)", table[line])
                make_yaml[table[0].strip()]['items']['properties'].update({split_stmt[0].lower(): {'type': 'string', 'maxLength': int(get_len[0])}})

            if "BYTEA" in table[line]:
                make_yaml[table[0].strip()]['items']['properties'].update({split_stmt[0].lower(): {'type': 'string', 'format': 'binary'}})    

            elif "TIMESTAMP" in table[line]:
                make_yaml[table[0].strip()]['items']['properties'].update({split_stmt[0].lower(): {'type': 'string'}})

            elif "DECIMAL" in table[line]:
                make_yaml[table[0].strip()]['items']['properties'].update({split_stmt[0].lower(): {'type': 'number', 'format': 'decimal'}})

            elif "DOUBLE" in table[line]:
                make_yaml[table[0].strip()]['items']['properties'].update({split_stmt[0].lower(): {'type': 'number', 'format': 'double'}}) 

            elif "INT" in table[line]:
                make_yaml[table[0].strip()]['items']['properties'].update({split_stmt[0].lower(): {'type': 'integer'}})

            elif "SMALLINT" in table[line]:
                make_yaml[table[0].strip()]['items']['properties'].update({split_stmt[0].lower(): {'type': 'integer', 'format': 'int16'}})


#   "properties": {
#             "source": {
#               "type": "string",
#               "maxLength": 40
#             },
            
# ['VARCHAR', 'INT,', 'BYTEA,', 'TIMESTAMP', 'DECIMAL', 'DOUBLE', 'SMALLINT,']

    return make_yaml

def get_table_name():
    cursor = connection.cursor()
    cursor.execute("SELECT table_name FROM user_tables WHERE table_name NOT LIKE '%USER%' AND table_name NOT LIKE '%ROLE%'")
    tables = [row[0].lower() for row in cursor.fetchall()]
    return tables 


# test = get_sql_str('anl_accuracy')
# test_dict = make_dict(test)

# with open('anl_accuracy.yaml', 'w') as file:
#     yaml.dump(test_dict, file, sort_keys=False)
# print("Done")

get_sql = get_table_name()

for table in range (0, len(get_sql)):
    cur_sql = get_sql_str(get_sql[table])
    cur_dict = make_dict(cur_sql)


    with open(rf'schema/{get_sql[table]}.yaml', 'w') as file:
        yaml.dump(cur_dict, file, sort_keys=False)


for file in range (0, len(get_sql)):
    with open (rf'schema/{get_sql[file]}.yaml', "r") as f:
        data = yaml.load(f, Loader=yaml.FullLoader)
        combined_data['components']['schemas'].update(data)


with open("schema/all-schemas.yaml", "w") as f:
    yaml.dump(combined_data, f, sort_keys=False)


# print(list_types)





# test = get_sql_str('cs_ellipsoid')
# test_dict = make_dict(test)


# print(get_sql)
# with open('test.yaml', 'w') as file:
#     yaml.dump(test_dict, file)

