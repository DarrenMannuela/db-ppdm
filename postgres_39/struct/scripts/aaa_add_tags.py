import subprocess, sys, os

def get_all_go_paths_in_dir():
    head = os.getcwd() + "/go"
    paths = list(os.listdir(head))
    go_paths = map(lambda f: os.path.join(head, f), paths)
    go_files = filter(lambda f: os.path.isfile(f) and f.endswith(".go"), go_paths)
    res = list(go_files)
    return res
    

def add_tags(file_path):
    # on macos
    command = f"~/go/bin/gomodifytags -file {file_path} -add-tags json -all -w -transform snakecase -quiet"
    os.system(command)

def remove_tags(file_path):
    # on macos
    command = f"~/go/bin/gomodifytags -file {file_path} -remove-tags json -all -w"
    os.system(command)    

def main():
    paths = get_all_go_paths_in_dir()
    n_paths = len(paths)
    for i, path in enumerate(paths):
        add_tags(file_path=path)
        print(f"Finished parsing file: {i + 1}/{n_paths}")
    pass


if __name__ == '__main__':
    main()
