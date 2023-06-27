import os
import sys
import lib_platform as platform

if platform.system == "darwin":
    sys.path.insert(0, "database")
elif platform.system == "windows":
    sys.path.insert(0, "../database")

from conn import connection

# Specify the directory path
directory = '/Users/darrenmp/Documents/vscode/db-ppdm-copy/permen_workspace_dto/gorm'

# Get a list of all items (files and folders) in the directory
items = os.listdir(directory)

# Filter out the files and keep only the folder names
folders = [item for item in items if os.path.isdir(os.path.join(directory, item))]

folders_name_change = []
table_names = []
func_names = []

# Print the folder names
for folder in folders:
    cur_name = ""
    func_name = ""
    table_name = ""
    cur_folder_list = folder.split("_")

    for word in cur_folder_list:
        cur_name += word.lower()
        func_name += word.title()
    folders_name_change.append(cur_name)
    func_names.append(func_name)


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
    abbreviation = "&"
    field_names = ''
    table_name_list =  table_name.split("_")
    for name in range(len(table_name_list)):
        cur_word = table_name_list[name]
        abbreviation += cur_word[0].lower()

    cursor = connection.cursor()
    cursor.execute('SELECT * FROM '+table_name)
    for i in cursor.description:
        field_names += abbreviation+"."+i[0].capitalize()+", "
        
    return field_names

def insert_attr(table_name: str):
    abbreviation = "&"
    field_names = ''
    table_name_list =  table_name.split("_")
    for name in range(len(table_name_list)):
        cur_word = table_name_list[name]
        abbreviation += cur_word[0].lower()

    cursor = connection.cursor()
    cursor.execute('SELECT * FROM '+table_name)
    for i in cursor.description:
        if i[0] != "ID":
            field_names += abbreviation+"."+i[0].capitalize()+", "
        
    return field_names


def get_num_bind (table_name: str):
    cursor = connection.cursor()
    cursor.execute(f"SELECT COUNT(*) FROM all_tab_columns WHERE table_name = '{table_name.upper()}'")
    num_columns = cursor.fetchone()[0]

    bind = fill_bind(num_columns)

    return bind




def make_get_sql(table_name: str):
    get = f'rows, err := db.Query("SELECT * FROM {table_name}")'+"""
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }"""
    return get


def make_get_sql_by_id(table_name: str):
    get = f'rows, err := db.Query("SELECT * FROM {table_name} WHERE id = :1", id)'+"""
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }"""
    return get




def make_insert_sql(table_name: str, desc: str, bind: str, attr: str):
    insert = f"_, err = tx.Exec(`INSERT INTO {table_name} ({desc[:-2]}) VALUES ({bind[:-2]})`, {attr[:-2]})"+"""
    if err != nil {"""+f"""
        tx.Rollback()
        fmt.Println({table_name.upper()})
        return err"""+"""
        }"""
    return insert



def make_type_handler(table_name: str, import_name: str, func_name: str, struct_name: str):
    abbreviation = ""
    table_name_list =  table_name.split("_")
    for name in range(len(table_name_list)):
        cur_word = table_name_list[name]
        abbreviation += cur_word[0].lower()


    

    attr = get_attr(table_name)
    insert = insert_attr(table_name)

    check_validity = """validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    """

    header = f"""package apiv1\nimport (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-{import_name}/dto"
    "github.com/DarrenMannuela/svc-datatype-{import_name}/utils"

    "github.com/gofiber/fiber/v2"
)
"""
    get_all = f"""func Get{func_name}(c *fiber.Ctx) error"""+"{"+f"""
    {check_validity}"""+f"""
    {make_get_sql(table_name)}\n
    """+f"""
    defer rows.Close()
    var results []dto.{struct_name}    
    """+"""
        for rows.Next() {
    """+f"""
            var {abbreviation} dto.{struct_name}
            if err := rows.Scan({attr[:-2]}); err != nil"""+"{"+"""
                return err
            }"""+f"""
            results = append(results, {abbreviation})"""+"""
        }
    """+"""
    return c.JSON(results)
}
"""



    # get_by_id = f"func Get{func_name}ById(c *fiber.Ctx) error"+"{"+"""\n
    #   {check_validity}
    # """
    set_all = f"""func Set{func_name}(c *fiber.Ctx) error"""+"{"+f"""
    {check_validity}"""+f"""
    var {abbreviation} dto.{struct_name}
    setDefaults(&{abbreviation})

    if err := c.BodyParser(&{abbreviation}); err != nill"""+"{"+"""
        return err
    }
    """+f"""
    {abbreviation}.Create_date = formatDateString(printWellReport.Create_date)
    """+"""

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

    var generatedID int64    
    """+f"""{make_insert_sql(table_name, get_description(table_name), get_num_bind(table_name), insert)}"""+"""
    return c.SendStatus(fiber.StatusOK)
}
"""

    get_by_id = f"""func Get{func_name}ById(c *fiber.Ctx) error"""+"{"+f"""
    {check_validity}
    id := c.Params("id")
    """+f"""
    {make_get_sql_by_id(table_name)}\n
    """+f"""
    defer rows.Close()
    var results []dto.{struct_name}    
    """+"""
        for rows.Next() {
    """+f"""
            var {abbreviation} dto.{struct_name}
            if err := rows.Scan({attr[:-2]}); err != nil"""+"{"+"""
                return err
            }"""+f"""
            results = append(results, {abbreviation})"""+"""
        }
    """+"""
    return c.JSON(results)
}
"""
    
    update_all = ""

    return header+get_all+set_all+get_by_id


test = "print_well_report_table"
desc = get_description(test)
bind = get_num_bind(test)
attr = get_attr(test)


testing = make_type_handler(test, 'printwellreporttable', "PrintWellReport", "Print_well_report")


with open(f"test.go", 'w') as file:
    file.write(testing)
