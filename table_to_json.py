import tabula
import pandas as pd
import PyPDF2
import re
import json

# Set the input and output file paths
input_file = 'PPDM.pdf'


page_amount = 0


with open('PPDM.pdf', 'rb') as pdf_file:
    pdf_reader = PyPDF2.PdfReader(pdf_file)
    page_amount = len(pdf_reader.pages)




with open('PPDM.pdf', 'rb') as pdf_file:
    cur_dict = {}
    cur_table = ''
    cur_len = 0


    # Read the pdf file
    pdf_reader = PyPDF2.PdfReader(pdf_file)
    # Get current page and extract the text 
    for page in range(4, page_amount):
        # List indexing page must minus 1 due to list staring from index 0
        cur_page = pdf_reader.pages[page-1]
        text = cur_page.extract_text()

        # Use regex to search the table name
        pattern = re.compile(r'Table Name:\s*(.*)')
        find_name = pattern.search(text)


        # Get the contents of the table by using tabula for the current page
        cur_df_list = tabula.read_pdf(input_file, pages=page, lattice=True)
        cur_df = cur_df_list[0]
        df = pd.DataFrame(cur_df)
        df_dict = df.to_dict()

        # Indicating that its from the same table
        if find_name is None:
            # Check if the following page contains a table name
            fol_page = pdf_reader.pages[page]
            fol_table = fol_page.extract_text()

            # Use regex to search the table name
            pattern = re.compile(r'Table Name:\s*(.*)')
            fol_name = pattern.search(fol_table)


            # Assigning the keys stored with each column
            holder = {}
            columns = list(df_dict.keys())

            # Adjust the key, to be align with the cur_table variable
            for column in columns:
                holder[column] = {}
                for key, val in df_dict[column].items():  
                    new_key = cur_len + int(key)
                    holder[column][new_key] = val


            if fol_name is None:
                for column in columns:
                    cur_dict[cur_table][column].update(holder[column])
                   


                # Get the new update length of the table rows
                cur_len = len(cur_dict[cur_table]['Column Name'])

            else:
                for column in columns:
                    cur_dict[cur_table][column].update(holder[column])
                with open(f'table_from_pdf/{cur_table}.json', 'w') as f:
                    json.dump(cur_dict, f)
                    cur_dict = {}
                    cur_table = ''


        else:
            # Get the table name
            table_name = find_name.group(1)
            cur_table = table_name
            cur_dict[cur_table]= df_dict
            cur_len = len(cur_dict[cur_table]['Column Name'])

            print(table_name)

            # Check if the following page contains a table name
            fol_page = pdf_reader.pages[page]
            fol_table = fol_page.extract_text()

            # Use regex to search the table name
            pattern = re.compile(r'Table Name:\s*(.*)')
            fol_name = pattern.search(fol_table)

            if fol_name is None:
                # If there is no table name on the following table we pass, cause that is handled on the if statement
                pass
            else:
                # If the following page contains table name then we can just dump the current dict into a json
                with open(f'table_from_pdf/{cur_table}.json', 'w') as f:
                    json.dump(cur_dict, f)
                    cur_dict = {}
                    cur_table = ''







# Extract data from the PDF table and save it as a CSV file
# tabula.convert_into(input_file, output_file, output_format='csv', pages='4-4389', lattice=True)


# # Make a for loop that will be the range of the pages in the pdf other then the cover page (4-4389), 
# we use this to get the current page and check if the table name exist on the page (to handle multi page tables in the pdf), 
# if table name doesnt exist on the page then we assume that it belongs in the same tabel as the previous page, we then store the table_name as a dict key, with the value being a dict containing the table content from the pdf using tabula .convert_into

# with open('PPDM.pdf', 'rb') as pdf_file:
#      # Read the PDF file
#     pdf_reader = PyPDF2.PdfReader(pdf_file)
#     # Get the first page of the PDF
#     first_page = pdf_reader.pages[9]
#     # Get the text content of the first page
#     text_content = first_page.extract_text()
#     # Use regular expression to find the table name
#     table_name_pattern = re.compile(r'Table Name:\s*(.*)')
#     table_name_match = table_name_pattern.search(text_content)
#     # Extract the table name from the regex match
#     table_name = table_name_match.group(1)

# # Print the table name
# print(table_name)


# # Read CSV file into DataFrame
# df = pd.read_csv('PPDM.csv')

# # Print first 5 rows of DataFrame
# print(df.columns)


