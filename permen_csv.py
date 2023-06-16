import pandas as pd
import os
import re
import shutil

# Specify the path to your CSV file
folder = '/Users/darrenmp/Documents/vscode/db-ppdm-copy/permen_csv'
data_types = []

# Loop through files in the folder and get file names
for filename in os.listdir(folder):
    # Create the full file path by joining the folder path and filename
    file_path = os.path.join(folder, filename)


    # Check if the path corresponds to a file (not a subfolder)
    if os.path.isfile(file_path):
        # Extract the file name from the full file path
        file_name = os.path.basename(file_path)

        # Process the file name
        data_types.append(file_name)


data_types.remove(".DS_Store")

for file in data_types:
    cur_file = f"/Users/darrenmp/Documents/vscode/db-ppdm-copy/permen_csv/{file}"
    # Read the CSV file
    df = pd.read_csv(cur_file)

    # Extract the first column
    first_column = df.iloc[:, 0]

    second_column = df.iloc[:, 1]


    cur_fields = []
    field_types = []

    count_col1 = 0
    count_col2 = 0
    data_type_check = 0

    for value in first_column:
        if pd.isna(value):
            pass
        else:
            field = value.split('_')

            go_field = ""

            for i in range(len(field)):
                if i == 0 and len(field) > 1:
                    go_field += field[i].title()+"_"
                elif i == 0 and len(field) == 1:
                    go_field+= field[i].title()
                elif i == len(field)-1:
                    go_field+= field[i].lower()
                else:
                    go_field+= field[i].lower()+"_"

            cur_fields.append(go_field)
            count_col1 += 1
    # print(cur_fields)

    for value in second_column:
        if isinstance(value, str):
            if re.match(r'^\s*Contoh\s*:\s*.*$', value):
                if "Tabel Isian Parameter Casing Sumur (Well Casing) Yang Digunakan " in cur_fields:
                    cur_fields.remove("Tabel Isian Parameter Casing Sumur (Well Casing) Yang Digunakan ")

                checker = False
                get_data_type = value.split(":")
                get_data_type[0] = get_data_type[0].replace(" ", "")

                # if file == "WELL DATA.csv":
                #     print(f'{get_data_type[0]} {cur_fields[count_col2]} :{get_data_type[-1]}')


                if re.search(r'\d',get_data_type[-1]) or re.search(r'[a-zA-Z]', get_data_type[-1]) or "-" and checker == False:
                    field_types.append("*string")
                    data_type_check += 1
                    # if file == "WELL DATA.csv":
                    #     print(f'{get_data_type[0]} {cur_fields[count_col2]} :{get_data_type[-1]}')
                    checker = True

                if get_data_type[-1].isdigit() and checker == False:
                    field_types.append("*int")
                    data_type_check += 1
                    # if file == "WELL DATA.csv":
                    #     print(f'{get_data_type[0]} {cur_fields[count_col2]} :{get_data_type[-1]}')
                    checker = True

                
                if re.match(r'^[\d+\-*/.]+$', get_data_type[-1]) and checker == False:
                    field_types.append("*int")
                    data_type_check += 1
                    # if file == "WELL DATA.csv":
                    #     print(f'{get_data_type[0]} {cur_fields[count_col2]} :{get_data_type[-1]}')
                    checker = True

                count_col2 += 1

    print(f"\n{file.split('.')[0]}\nThe 1st column contains: {len(cur_fields)}\nThe 2nd column contains: {len(field_types)}\nThe amount of data types found column contains: {data_type_check}\n")

    # table_name  = f"\n{file.split('.')[0]}"
    # with open("field.txt", 'a') as file:
    #     file.write(table_name)
    #     file.write(str(cur_fields))

    get_struct_name = file.split('.')[0]
    get_struct_name = get_struct_name.split("_")

    struct_name = ""

    for word in range(len(get_struct_name)):
        if len(get_struct_name) == 1:
            struct_name += get_struct_name[word].title()
        elif len(get_struct_name) > 1:
            if word == 0:
              struct_name += get_struct_name[word].title()+"_"
            elif word == len(get_struct_name)-1:
                struct_name += get_struct_name[word].lower()
            else:
                struct_name += get_struct_name[word].lower()+"_"



    print(struct_name)

    opener = "package dto\n\n"+"type "+struct_name+" struct{\n\n"


    content = ""

    for field in range(len(cur_fields)):
        if cur_fields[field] not in content:
            content += f'{cur_fields[field].replace(" ", "")}       {field_types[field]}   `json:"{cur_fields[field].replace(" ", "").lower()}" default:""`\n'

    closer = "}"

    abbreviation_list = struct_name.split("_")

    abbreviation = ""

    for word in range(len(abbreviation_list)):
        if len(abbreviation_list) == 1:
            abbreviation += abbreviation_list[0]
        else:
            get_letter = abbreviation_list[word][0]
            abbreviation += get_letter

    abbreviation = abbreviation.title()
    abbreviation += "_id"

    file_name_list = struct_name.split("_")
    file_name = ""

    for name in file_name_list:
        file_name += name.title()


            


    workspace_field = ["Id", "Afe_number"]

    workspace_field.append(abbreviation)
    print(workspace_field)


    content = ""

    for field in range(len(cur_fields)):
        if cur_fields[field] not in content:
            content += f'{cur_fields[field].replace(" ", "")}       {field_types[field]}   `json:"{cur_fields[field].replace(" ", "").lower()}" default:""`\n'

    closer = "}"


    opener_workspace = "package dto\n\n"+"type Workspace struct{\n\n"

    content_workspace = ""


    for field in range(len(workspace_field)):
        if field == 0:
            content_workspace += f'{workspace_field[field].replace(" ", "")}       int  `json:"{workspace_field[field].replace(" ", "").lower()}" default:""`\n'
        else:
             content_workspace += f'{workspace_field[field]}       *int  `json:"{workspace_field[field].replace(" ", "").lower()}" default:""`\n'
            


    # os.makedirs(f"permen_workspace_dto/{file_name}")


    with open(f"permen_workspace_dto/{file_name}/workspace.go", 'w') as file:
        file.write(opener_workspace+content_workspace+closer)


    with open(f"permen_dto/{file_name}.go", 'w') as file:
        file.write(opener+content+closer)


    afe = "/Users/darrenmp/Documents/vscode/db-ppdm-copy/temp_dto/Afe.go"
    submission = "/Users/darrenmp/Documents/vscode/db-ppdm-copy/temp_dto/submission.go"
    folder = f"permen_workspace_dto/{file_name}/"


    shutil.copy(afe, folder)
    shutil.copy(submission, folder)


# print(opener+content+closer)





