# migrating data types
import os

def get_all_go_paths_in_dir():
    head = os.getcwd() + "/go"
    paths = list(os.listdir(head))
    go_paths = map(lambda f: os.path.join(head, f), paths)
    go_files = filter(lambda f: os.path.isfile(f) and f.endswith(".go"), go_paths)
    res = list(go_files)
    return res
    

def file_to_string(file_path) -> str:

    ss = ""
    with open(file_path, "r") as f:
        ss = "".join(f.readlines())\
            .replace("___", "_")

    with open(file_path, "w") as f:
        f.write(ss)

def main():
    paths = get_all_go_paths_in_dir()
    n_paths = len(paths)
    for i, path in enumerate(paths):
        file_to_string(path)
        print(f"file: {i+1}/{n_paths}")

if __name__ == '__main__':
    main()