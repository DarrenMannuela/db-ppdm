import sys
import lib_platform as platform

if platform.system == "darwin":
    sys.path.insert(0, "database")
elif platform.system == "windows":
    sys.path.insert(0, "../database")

from conn import connection

def fill_bind(n):
    holder = ''
    for i in range (0, n):
        holder += ":"+str(i+1)+", "
    return holder


def fill_param(n):
    holder = []
    holder.append("A1")
    for i in range (1, n):
        holder.append(None)
    return holder


def get_description(table_name: str):
    field_names = ''
    cursor = connection.cursor()
    cursor.execute('SELECT * FROM '+table_name)
    for i in cursor.description:
        field_names += i[0].lower()+", "

    return field_names


def get_attr(table_name: str):
    field_names = ''
    var_name = ''
    var =  table_name.split("_")
    if len(var) > 1:
        for word in range(len(var)):
            if word == 0:
                var_name += var[word]
            else:
                var_name += var[word].capitalize()
    else:
        var_name += var[0]

    cursor = connection.cursor()
    cursor.execute('SELECT * FROM '+table_name)
    for i in cursor.description:
        field_names += var_name+"."+i[0].capitalize()+", "
        


    return field_names



table_name = "print_well_report_table"
# print("FILL THE BIND:\n")
# print(fill_bind(30))
# print("\n\n") 

cursor = connection.cursor()
cursor.execute(f"SELECT COUNT(*) FROM all_tab_columns WHERE table_name = '{table_name.upper()}'")
num_columns = cursor.fetchone()[0]

bind = fill_bind(num_columns)

# print("FILL THE PARAM:\n")
# print(fill_param(30))
# print("\n\n") 



# print("FILL THE DESC:\n")
# print(get_description(table_name))
# print("\n\n") 

desc = get_description(table_name)


# print("MODEL:\n")
# print(get_attr(table_name))
# print("\n\n") 

attr = get_attr(table_name)

print(f"_, err = tx.Exec(`INSERT INTO {table_name} ({desc[:-2]}) VALUES ({bind[:-2]})`, {attr[:-2]})"+"\nif err != nil {\n\ttx.Rollback()"+f"\n\tfmt.Println(\"{table_name.upper()}\")\n\treturn err"+"\n}")

print(attr)