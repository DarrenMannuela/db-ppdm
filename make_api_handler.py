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
        table_name += word.lower()+"_"
    table_name += "table"
    table_names.append(table_name)
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

def insert_description(table_name: str):
    field_names = ''
    cursor = connection.cursor()
    cursor.execute('SELECT * FROM '+table_name)
    for i in cursor.description:
        if i[0] != "ID":
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

    bind = fill_bind(num_columns-1)

    return bind

def get_num_col (table_name: str):
    cursor = connection.cursor()
    cursor.execute(f"SELECT COUNT(*) FROM all_tab_columns WHERE table_name = '{table_name.upper()}'")
    num_columns = cursor.fetchone()[0]


    return num_columns-1



def make_get_sql(table_name: str):
    get = f'rows, err := db.Query(`SELECT * FROM {table_name}`)'+"""
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }"""
    return get


def make_get_sql_by_id(table_name: str):
    get = f'rows, err := db.Query(`SELECT * FROM {table_name} WHERE id = :1`, id)'+"""
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }"""
    return get

def make_id_exist(table_name: str):
    get = f"""
    var idExist string
    err = db.QueryRow(`SELECT * FROM {table_name} WHERE id = :1`, id).Scan(&idExist)"""+"""
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }"""
    return get


def make_id_exist_del(workpsace: str):
    get = f"""
    var idExist string
	_ = tx.QueryRow(`SELECT {workpsace}id FROM {workpsace}workspace WHERE {workpsace}id = :1`, id).Scan(&idExist)
	if idExist != "" """+"{"+f"""
		_, err = tx.Exec(`DELETE FROM {workpsace}workspace WHERE {workpsace}id = :1`, id)
		if err != nil""" + "{"+"""
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    """
    return get


def make_insert_sql(table_name: str, desc: str, bind: str, attr: str):

    insert = f"_, err = tx.Exec(`INSERT INTO {table_name} ({desc[:-2]}) VALUES ({bind[:-2]}) RETURNING id INTO :{get_num_col(table_name)+1}`, {attr[:-2]}"+", sql.Out{Dest: &generatedID})"+"""
    if err != nil {"""+f"""
        tx.Rollback()
        fmt.Println({table_name.upper()})
        return err"""+"""
        }"""
    return insert


def make_update_sql(table_name: str, desc: str, bind: str, attr: str):
    combine = ""

    bind_list = bind.split(',')
    desc_list = desc.split(',')

    bind_list = bind_list[:-1]
    desc_list = desc_list[:-1]

    for bind in range(len(bind_list)):
        if bind < len(bind_list)-1:
           combine += desc_list[bind]+" = "+bind_list[bind].replace(" ", "")+","
        else:
            combine += desc_list[bind]+" = "+bind_list[bind].replace(" ", "")
    

    update = f"""_, err = tx.Exec(`UPDATE {table_name} SET
        {combine} WHERE id = :{get_num_col(table_name)+1}`, {attr}id)"""+"""
        if err != nil {"""+f"""
            tx.Rollback()
            fmt.Println({table_name.upper()})
            return err"""+"""
        }"""
    return update


def make_get_by_afe(table_name: str):
    get = f"""
    var idExist string
    err = db.QueryRow(`SELECT {table_name}id FROM {table_name}table WHERE afe_number = :1`, id).Scan(&idExist)"""+"""
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }"""
    return get

def make_patch_sql(insert_attr: str):
    insert_attr_list = insert_attr.split(",")

    patch = ""
    
    for attr in insert_attr_list:
        get_column = attr.split(".")
        get_field = attr.replace(" ", "")
        get_field = get_field[1:]

        if len(get_column) > 1:
            field_name = get_column[1].lower()
            cur_if = f"""
        if {get_field} != nil"""+"{"+f"""
             _, err = tx.Exec(`UPDATE {table_name} SET {field_name} = :1 WHERE id = :2`, {get_field}, id)
        """+"""
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            """
            patch += cur_if
    return patch

def make_type_handler(table_name: str, import_name: str, func_name: str, struct_name: str):
    abbreviation = ""
    workspace = ""
    table_name_list =  table_name.split("_")
    for name in range(len(table_name_list)):
        cur_word = table_name_list[name]
        abbreviation += cur_word[0].lower()

        if "table" not in cur_word:
            workspace += cur_word.lower()+"_"

    attr = get_attr(table_name)
    insert = insert_attr(table_name)

    insert_desc = insert_description(table_name)

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
    get= f"""func Get{func_name}(c *fiber.Ctx) error"""+"{"+f"""
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
    set = f"""func Set{func_name}(c *fiber.Ctx) error"""+"{"+f"""
    {check_validity}"""+f"""
    var {abbreviation} dto.{struct_name}
    setDefaults(&{abbreviation})

    if err := c.BodyParser(&{abbreviation}); err != nill"""+"{"+"""
        return err
    }
    """+f"""
    {abbreviation}.Create_date = formatDateString({abbreviation}.Create_date)
    """+"""

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}()

    var generatedID int64    
    """+f"""{make_insert_sql(table_name, insert_desc, get_num_bind(table_name), insert)}"""+"""
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
    
    update = f"""func Put{func_name}(c *fiber.Ctx) error"""+"{"+f"""
    {check_validity}"""+f"""
    id := c.Params("id")
    var {abbreviation} dto.{struct_name}
    setDefaults(&{abbreviation})

    if err := c.BodyParser(&{abbreviation}); err != nill"""+"{"+"""
        return err
    }
    """+f"""
    {abbreviation}.Create_date = formatDateString({abbreviation}.Create_date)
    """+"""

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}()
    """+f"""
    {make_id_exist(table_name)}
    """+"""
    if idExist != "" {

    """+f"""
        {make_update_sql(table_name, insert_desc, bind, insert)}
    """+"""
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    """+"""
    return c.SendStatus(fiber.StatusOK)
}
"""

    delete = f"""func Delete{func_name}(c *fiber.Ctx) error"""+"{"+f"""
    {check_validity}"""+f"""
    id := c.Params("id")"""+"""
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}()
    """+f"""
    {make_id_exist_del(workspace)}
    _, err = tx.Exec(`DELETE FROM {table_name} WHERE id = :1`, id)
	if err != nil"""+"{"+"""
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
    """+"""
    return c.SendStatus(fiber.StatusOK)
}
"""
    patch = f"""func Patch{func_name}(c *fiber.Ctx) error"""+"{"+f"""
    {check_validity}"""+f"""
    id := c.Params("id")
    var {abbreviation} dto.{struct_name}
    setDefaults(&{abbreviation})

    if err := c.BodyParser(&{abbreviation}); err != nill"""+"{"+"""
        return err
    }
    """ + """
    tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}()
    """+f"""
    {make_get_by_afe(workspace)}
    """+"""
    if idExist != ""{"""+f"""{make_patch_sql(insert)}"""+"""
    }
    err = tx.Commit()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(map[string]string{"message": "Record updated successfully"})
}

    """

    return header+get+set+get_by_id+update+delete+patch


test = "print_well_report_table"
desc = insert_description(test)
bind = get_num_bind(test)
attr =insert_attr(test)


testing = make_type_handler(test, 'printwellreport', "PrintWellReport", "Print_well_report")

make_update_sql(test, desc, bind, attr)


print(folders)

print(table_names)

print(folders_name_change)

print(func_names)

# for file in range(len(folders)):



with open(f"test.go", 'w') as file:
    file.write(testing)


