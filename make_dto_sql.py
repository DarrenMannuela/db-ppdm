import re 
import shutil
import os


# Specify the path to your CSV file
file = '/Users/darrenmp/Documents/vscode/db-ppdm-copy/sql_tables/permen.sql'
data_types = []
table_names = []
cur_tables = {}

check_data_types = []


def split_name(split_name: str):
    struct_name = ""

    split_name_list = split_name.split("_")

    for word in range(len(split_name_list)):
        if word == 0 :
            struct_name += split_name_list[word].title()+"_"

        elif word != len(split_name_list)-1:
            struct_name += split_name_list[word].lower()+"_"
        else:
            struct_name += split_name_list[word].lower()
    return struct_name


with open(file, "r") as file:
    dto_file = file.read()
    pattern = r"CREATE\s+TABLE\s+(\w+)\s*\((.+?)\);"
    tables = re.findall(pattern, dto_file, re.IGNORECASE | re.DOTALL)
    
    for table in tables:
        workspace_field = ["Id", "Afe_number"]
        table_fields = {}
        table_name = table[0]
        struct_name = split_name(table_name)

        file_name_list = split_name(table_name)
        file_name = file_name_list.split("_table")[0]

        workspace_field.append(file_name+"_id")


        opener = "package dto\n\n"+"type "+file_name+" struct{\n\n"
        content = ""

        gorm_content = ""
        closer = "}"


        opener_workspace = "package dto\n\n"+"type Workspace struct{\n\n"
        content_workspace = ""
        gorm_content_workspace = ""


        for field in range(len(workspace_field)):
            if field == 0:
                content_workspace += f'{workspace_field[field].replace(" ", "")}       int  `json:"{workspace_field[field].replace(" ", "").lower()}" default:""`\n'
                gorm_content_workspace += f'{workspace_field[field].replace(" ", "")}       int  `json:"{workspace_field[field].replace(" ", "").lower()}" default:""`\n'
            else:
                content_workspace += f'{workspace_field[field]}       *int  `json:"{workspace_field[field].replace(" ", "").lower()}" default:""`\n'
                gorm_content_workspace += f'{workspace_field[field]}       *int  `json:"{workspace_field[field].replace(" ", "").lower()}" default:""`\n'

        gorm_content_workspace += f'{file_name}       {file_name}  `json:"{file_name.lower()}" default:"" gorm:"foreignKey:{file_name+"_id"}"`\n'

        fields = table[1].strip().split('\n')
        fields = [field.strip() for field in fields]
        print("Table Name:", struct_name)
        print("Fields:")
        for field in fields:
            cur_field_list = field.split(" ")
            print(cur_field_list[0])
            # print(split_name(cur_field_list[0]))
            if cur_field_list[1] not in check_data_types:
                check_data_types.append(cur_field_list[1])

            if cur_field_list[0] == "ID":
                table_fields['Id'] = "int"
                content += f'Id      int  `json:"id" default:""`\n'
                gorm_content += f'Id      int  `json:"id" default:"" gorm:"primaryKey"`\n'
            else:
                if "VARCHAR" in cur_field_list[1]:
                    print("I am varchar")
                    content += f'{split_name(cur_field_list[0])}   *string  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                    gorm_content += f'{split_name(cur_field_list[0])}   *string  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                if "FLOAT" in cur_field_list[1]:
                    print("I am float")
                    content += f'{split_name(cur_field_list[0])}   *int  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                    gorm_content += f'{split_name(cur_field_list[0])}   *int  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                if "TIMESTAMP" in cur_field_list[1]:
                    print("I am time")
                    content += f'{split_name(cur_field_list[0])}   *string  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                    gorm_content += f'{split_name(cur_field_list[0])}   *string  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                if "NUMBER" in cur_field_list[1]:
                    print("I am number")
                    content += f'{split_name(cur_field_list[0])}   *int  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                    gorm_content += f'{split_name(cur_field_list[0])}   *int  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                if "INT" in cur_field_list[1]:
                    print("I am int")
                    content += f'{split_name(cur_field_list[0])}   *int  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                    gorm_content += f'{split_name(cur_field_list[0])}   *int  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                if "DECIMAL" in cur_field_list[1]:
                    print("I am int")
                    content += f'{split_name(cur_field_list[0])}   *int  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                    gorm_content += f'{split_name(cur_field_list[0])}   *int  `json:"{split_name(cur_field_list[0]).lower()}" default:""`\n'
                

        # os.makedirs(f"permen_workspace_dto/gorm/{file_name}")
        # os.makedirs(f"permen_workspace_dto/no_gorm/{file_name}")

        with open(f"permen_dto/no_gorm/{file_name}.go", 'w') as file:
            file.write(opener+content+closer)

        with open(f"permen_dto/gorm/{file_name}.go", 'w') as file:
            file.write(opener+gorm_content+closer)
        
        with open(f"permen_workspace_dto/gorm/{file_name}/workspace.go", 'w') as file:
            file.write(opener_workspace+gorm_content_workspace+closer)

        with open(f"permen_workspace_dto/no_gorm/{file_name}/workspace.go", 'w') as file:
            file.write(opener_workspace+content_workspace+closer)


        afe = "/Users/darrenmp/Documents/vscode/db-ppdm-copy/temp_dto/Afe.go"
        submission = "/Users/darrenmp/Documents/vscode/db-ppdm-copy/temp_dto/submission.go"
        token = "/Users/darrenmp/Documents/vscode/db-ppdm-copy/temp_dto/Token.go"
        credential = "/Users/darrenmp/Documents/vscode/db-ppdm-copy/temp_dto/Credential.go"
        gorm_dto = f"/Users/darrenmp/Documents/vscode/db-ppdm-copy/permen_dto/gorm/{file_name}.go"
        no_gorm_dto = f"/Users/darrenmp/Documents/vscode/db-ppdm-copy/permen_dto/no_gorm/{file_name}.go"
        gorm_folder = f"permen_workspace_dto/gorm/{file_name}/"
        no_gorm_folder = f"permen_workspace_dto/no_gorm/{file_name}/"



        shutil.copy(afe, gorm_folder)
        shutil.copy(submission, gorm_folder)
        shutil.copy(gorm_dto, gorm_folder)
        shutil.copy(token, gorm_folder)
        shutil.copy(credential, gorm_folder)


        shutil.copy(afe, no_gorm_folder)
        shutil.copy(submission, no_gorm_folder)
        shutil.copy(no_gorm_dto, no_gorm_folder)
        shutil.copy(token, no_gorm_folder)
        shutil.copy(credential, no_gorm_folder)


            


print(opener+content+closer)
print(check_data_types)