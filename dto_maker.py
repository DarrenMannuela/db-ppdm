import shutil
import os

old_dir = "postgres_39/all_struct_with_default/"


print_well_view = ['business_associate', 'area', 'field', 'well', 'rm_information_item', 'rm_creator', 'r_media_type', 'r_document_type', 'r_item_category', 'r_item_sub_category', 'rm_physical_item', 'anl_accuracy', 'rm_data_store', 'r_source','ppdm_quality_control']

def make_folder(view_name):
    view_name = view_name.lower()
    view_name = view_name.replace(" ", "_")
    curr_dir = os.getcwd()
    dto_dir = os.path.join(curr_dir, "dto_maker")
    if not os.path.exists(dto_dir):
        os.mkdir(dto_dir)
    new_folder = os.path.join(dto_dir, view_name)
    if not os.path.exists(new_folder):
        os.mkdir(new_folder)
    return view_name


def make_dtos(view: list, file: str):
    for table in view:
        name =  table.replace("_", " ")
        name = name.title()
        name = name.replace(" ", "")
        name = name+".go"
        new_dir = f'dto_maker/{file}/{name}'
        shutil.copy(old_dir+name, new_dir)

if __name__ == "__main__":
    view = 'print well report'
    file_path = make_folder(view)
    make_dtos(print_well_view, file_path)
