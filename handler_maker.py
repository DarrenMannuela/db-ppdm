import shutil
import os

old_dir = "postgres_39/easy/all/"


print_well_view = ['business_associate', 'area', 'field', 'well', 'rm_information_item', 'rm_creator', 'r_media_type', 'r_document_type', 'r_item_category', 'r_item_sub_category', 'rm_physical_item', 'anl_accuracy', 'rm_data_store', 'r_source','ppdm_quality_control']

def make_folder(view_name):
    view_name = view_name.lower()
    view_name = view_name.replace(" ", "_")
    curr_dir = os.getcwd()
    handler_dir = os.path.join(curr_dir, "handler_maker")
    if not os.path.exists(handler_dir):
        os.mkdir(handler_dir)
    new_folder = os.path.join( handler_dir, view_name)
    if not os.path.exists(new_folder):
        os.mkdir(new_folder)
    return view_name


def make_handler(view: list, file: str):
    for table in view:
        name =  table+"_handler.go"
        new_dir = f'handler_maker/{file}/{name}'
        shutil.copy(old_dir+name, new_dir)

def replace_link(view, name):
    name = name.strip()
    name = name + "_handler.go"
    curr_dir = os.getcwd()
    handler_dir = os.path.join(curr_dir, "handler_maker")
    view_dir = os.path.join(handler_dir, view)
    filename = os.path.join(view_dir, name)
    with open(filename, 'r') as f:
        data = f.read()
    data = data.replace("pt_gtn_bibliography", "pt_gtn_"+view)
    with open(filename, "w") as f:
        data = f.write(data)

if __name__ == "__main__":
    view = 'print well report'
    file_path = make_folder(view)
    make_handler(print_well_view, file_path)
    
    for i in print_well_view:
        replace_link("print_well_report", i)
