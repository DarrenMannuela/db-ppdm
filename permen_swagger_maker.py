import os 
import re
import yaml


folder = '/Users/darrenmp/Documents/vscode/db-ppdm-copy/permen_dto'
views = []


def afe_filler():
    afe_example = {"example": {
                    "afe_number": 1,
                    "workspace_name": "LoremIpsum",
                    "kkks_name": "LoremIpsum",
                    "working_area": "LoremIpsum",
                    "submission_type": "Lorem",
                    "data_type": rf"{tag}"},
                    "user_id": 1
                    }
    return afe_example


def workspace_filler():
    workspace_example = {"example": {
                    "Id": 1,
                    "Afe_number": 1,
                    rf"{workspace_id}": 1}
                    }
    return workspace_example


def afe_struct_filler():

    afe_struct = {"afe" : {
        "type" : "array",
        "items" : {
          "type" : "object",
          "properties" : {
            "afe_number": {
              "type": "integer"
            },
            "workspace_name": {
              "type": "string"
            },
            "kkks_name" : {
              "type" : "string"
            },
            "working_area" : {
              "type" : "string"
            },
            "submission_type" : {
              "type" : "string"
            },
            "data_type" : {
              "type" : "string"
            }
          }
        }
      }}
    return afe_struct


# Loop through files in the folder and get file names
for filename in os.listdir(folder):
    # Create the full file path by joining the folder path and filename
    file_path = os.path.join(folder, filename)


    # Check if the path corresponds to a file (not a subfolder)
    if os.path.isfile(file_path):
        # Extract the file name from the full file path
        file_name = os.path.basename(file_path)

        # Process the file name
        views.append(file_name)


views.remove(".DS_Store")

print(views)


for file in views:
    afe = "afe"
    title = ""
    tag = ""


    endpoint_holder = {'paths': {}}
    cur_file = f"/Users/darrenmp/Documents/vscode/db-ppdm-copy/permen_dto/{file}"


    split_file = file.split(".")
    seperated = re.findall('[A-Z][^A-Z]*', split_file[0])


    workspace_file = f"/Users/darrenmp/Documents/vscode/db-ppdm-copy/permen_workspace_dto/{split_file[0]}/workspace.go"



    for word in range (len(seperated)):
        if word == len(seperated)-1:
            title += seperated[word]
            tag += seperated[word]
        else:
            title += seperated[word]+"_"
            tag += seperated[word]+" "
    

    cur_table = {'components': {'schemas': {}, 'securitySchemes': {'OAuth2PasswordBearer': {'type': 'oauth2', "flows": {"password":{"scopes": {}, "tokenUrl": "v1/token"}}}}}}

    cur_table['components']["schemas"].update(afe_struct_filler())

    base = {'openapi': '3.0.0', 'info': {'description': rf'This is the swagger API for {title}', 'version': '1.0.0', 'title': rf'{title}', 'termsOfService': 'http://swagger.io/terms/', 
        'contact':{'email': 'darren.mannuela@gmail.com'}, 'license': {'name': 'Apache 2.0', 'url': 'http://www.apache.org/licenses/LICENSE-2.0.html'}},
        'servers': [{'description': rf'{title}', 'url': 'http://localhost:8080/api/v1'},
                    {'description': 'SwaggerHub API Auto Mocking', 'url': rf'https://virtserver.swaggerhub.com/DarrenMannuela/{title.replace(" ", "")}/1.0.0'}], 
        
        
        
        'tags':[{'name': 'Afe', 'description': rf'All endpoints related to get {title} AFE'}, 
                {'name': 'Workspace', 'description': rf'All endpoints related to {title} Workspace'},
                {'name': rf'{tag}', 'description': rf'All endpoints related to {title}'},
                {'name': rf'{tag}'+" Dummy Data", 'description': rf'All endpoints related to {title}'}]}
    
    endpoint_name = ""
    for word in range (len(seperated)):
        if word == len(seperated)-1:
            endpoint_name += seperated[word]
        else:
            endpoint_name += seperated[word]+"-"


    afe_endpoint = {rf'/{endpoint_name}'+"-afe": 
            {'get':{'security': [{'Authorization':[]}], 'tags': ['Afe'], 'summary': rf'Get {title} AFE', 
                    'responses':{'200': {'description': rf'get {title} AFE data to be returned', 'content': {'application/json': {'example': [[]]}}}}}, 
            'post': {'security': [{'OAuth2PasswordBearer': []}], 'tags': ['Afe'], 'summary': rf'Post a new {title} AFE', 
                    'description': rf'Create a new {title} AFE data', 
                    'requestBody': {'required': True, 'description': rf'Request body to create {title} AFE data', 
                                        'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 
                                                                        'example': []}}}, 
                                                                        'responses': {'200': {'description': rf'{title} data is added', 'content': {'application/json': {'example': {'message': rf'The {title} AFE data was successfully added'}}}}}}},
                                                                                          
            rf'/{endpoint_name}'+"-afe/"+"{afe}":
                {'parameters': [{'in': 'path', 'name': 'afe', 'required': True, 'description': rf'afe of {title} data to fetch', 'schema':{'type': 'interger'}}],
                'get':{'security': [{'Authorization':[]}], 'tags': ['Afe'], 'summary': rf'Get {title} AFE', 
                    'responses':{'200': {'description': rf'get {title} AFE data to be returned', 'content': {'application/json': {'example': [[]]}}}}},
                'put': {'security': [{'Authorization':[]}], 'tags': ['Afe'], 'summary': rf'Update a new {title} AFE data', 'description': rf'Update a new {title} data', 
                            'requestBody': {'required': True, 'description': rf'Request body to update {title} AFE data', 
                                            'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 'example': []}}},
                                                        'responses': {'200': {'description': rf'{title} data is completely updated', 'content': {'application/json': {'example': {'message': rf'The {title} AFE data is completely updated'}}}}}},

                'patch': {'security': [{'Authorization':[]}], 'tags': ['Afe'], 'summary': rf'Update a new {title} AFE data', 'description': rf'Update a new {title} data', 
                                        'requestBody': {'required': True, 'description': rf'Request body to update {title} AFE data', 
                                                        'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 'example': []}}},
                                                                    'responses': {'200': {'description': rf'one row in {title} data is updated', 'content': {'application/json': {'example': {'message': rf'The one row in {title} AFE data is updated'}}}}}},
                                                                    
                'delete': {'security': [{'Authorization':[]}], 'tags': ['Afe'], 'summary': rf'Delete {title} AFE data', 'description': rf'Delete {title} AFE data',
                            'responses':{'200': {'description': rf'{title} AFE data is deleted', 'content': {'application/json': {'example': {'message': rf'The {title} AFE data was successfully deleted'}}}}}}}}
    
    
    afe_endpoint[rf'/{endpoint_name}'+"-afe"]['get']['responses']['200']['content']['application/json'].update(afe_filler())
    afe_endpoint[rf'/{endpoint_name}'+"-afe"]['post']['requestBody']['content']['application/json'].update(afe_filler())
    afe_endpoint[rf'/{endpoint_name}'+"-afe/"+"{afe}"]['get']['responses']['200']['content']['application/json'].update(afe_filler())
    afe_endpoint[rf'/{endpoint_name}'+"-afe/"+"{afe}"]['put']['requestBody']['content']['application/json'].update(afe_filler())
    afe_endpoint[rf'/{endpoint_name}'+"-afe/"+"{afe}"]['patch']['requestBody']['content']['application/json'].update(afe_filler())

    endpoint_holder['paths'].update(afe_endpoint)


    with open(workspace_file, "r+") as workspace:
        def workspace_struct_filler():
            workspace_struct = {
                rf"{title}"+"_workspace" : {
                    "type" : "array",
                    "items" : {
                        "type" : "object",
                    "properties" : {
                        "id": {
                        "type": "integer"
                        },
                        "afe_number": {
                        "type": "integer"
                        },
                        rf"{workspace_id}" : {
                        "type" : "integer"
                        }
                    }
                }
            }
            }
            return workspace_struct
            

        cur_workspace = workspace.read()
        get_field = r"\s+([^\s]+)\s+\*int"
        fields = re.findall(get_field, cur_workspace)
        workspace_id = fields[1]

        for field in fields:
            workspace_endpoint = {rf'/{endpoint_name}'+"-workspace": 
            {'get':{'security': [{'Authorization':[]}], 'tags': ['Workspace'], 'summary': rf'Get {title} Workspace', 
                    'responses':{'200': {'description': rf'get {title} Workspace data to be returned', 'content': {'application/json': {'example': [[]]}}}}}, 
            'post': {'security': [{'OAuth2PasswordBearer': []}], 'tags': ['Workspace'], 'summary': rf'Post a new {title} Workspace', 
                    'description': rf'Create a new {title} Workspace data', 
                    'requestBody': {'required': True, 'description': rf'Request body to create {title} Workspace data', 
                                        'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 
                                                                        'example': []}}}, 
                                                                            'responses': {'200': {'description': rf'{title} Workspace data is added', 'content': {'application/json': {'example': {'message': rf'The {title} Workspace data was successfully added'}}}}}}},
            rf'/{endpoint_name}'+"-workspace/"+"{afe}}": 
                {'parameters': [{'in': 'path', 'name': 'afe', 'required': True, 'description': rf'afe of {title} data to fetch', 'schema':{'type': 'interger'}}],
                    'get':{'security': [{'Authorization':[]}], 'tags': ['Workspace'], 'summary': rf'Get {title} Workspace', 
                        'responses':{'200': {'description': rf'get {title} Workspace data to be returned', 'content': {'application/json': {'example': [[]]}}}}}},
        
        
                                                                                          
            rf'/{endpoint_name}'+"-workspace/"+"{id}":
                {'parameters': [{'in': 'path', 'name': 'id', 'required': True, 'description': rf'id of {title} data to fetch', 'schema':{'type': 'interger'}}],
                'put': {'security': [{'Authorization':[]}], 'tags': ['Workspace'], 'summary': rf'Update a new {title} Workspace data', 'description': rf'Update a new {title} Workspace Workspace data', 
                            'requestBody': {'required': True, 'description': rf'Request body to update {title} data', 
                                            'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 'example': []}}},
                                                        'responses': {'200': {'description': rf'{title} Workspace data is completely updated', 'content': {'application/json': {'example': {'message': rf'The {title} Workspace data is completely updated'}}}}}},

                'patch': {'security': [{'Authorization':[]}], 'tags': ['Workspace'], 'summary': rf'Update a new {title} data', 'description': rf'Update a new {title} data', 
                                        'requestBody': {'required': True, 'description': rf'Request body to update {title} data', 
                                                        'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 'example': []}}},
                                                                    'responses': {'200': {'description': rf'one row in {title} data is updated', 'content': {'application/json': {'example': {'message': rf'The one row in {title} Workspace data is updated'}}}}}},
                                                                    
                'delete': {'security': [{'Authorization':[]}], 'tags': ['Workspace'], 'summary': rf'Delete {title} Workspace data', 'description': rf'Delete {title} data',
                            'responses':{'200': {'description': rf'{title} Workspace data is deleted', 'content': {'application/json': {'example': {'message': rf'The {title} Workspace data was successfully deleted'}}}}}}}}
            
            workspace_endpoint[rf'/{endpoint_name}'+"-workspace"]['get']['responses']['200']['content']['application/json'].update(workspace_filler())
            workspace_endpoint[rf'/{endpoint_name}'+"-workspace"]['post']['requestBody']['content']['application/json'].update(workspace_filler())
            workspace_endpoint[rf'/{endpoint_name}'+"-workspace/"+"{afe}}"]['get']['responses']['200']['content']['application/json'].update(workspace_filler())
            workspace_endpoint[rf'/{endpoint_name}'+"-workspace/"+"{id}"]['put']['requestBody']['content']['application/json'].update(workspace_filler())
            workspace_endpoint[rf'/{endpoint_name}'+"-workspace/"+"{id}"]['patch']['requestBody']['content']['application/json'].update(workspace_filler())

            

        cur_table['components']["schemas"].update(workspace_struct_filler())
        cur_table['components']["schemas"].update(afe_struct_filler())

        endpoint_holder['paths'].update(workspace_endpoint)

    with open(cur_file, "r+") as dto:
       filler = {'example': {}}
       schema =  {}
       def table_filler(filler: dict):
           table_filler = filler
           return table_filler
       

       cur_dto =  dto.read()
       get_field = r"\s+([^\s]+)\s+\*string"
       fields = re.findall(get_field, cur_dto)
       endpoint = {rf'/{endpoint_name}': 
        {'get':{'security': [{'Authorization':[]}], 'tags': [rf'{tag}'], 'summary': rf'Get {title}', 
                'responses':{'200': {'description': rf'get {title} data to be returned', 'content': {'application/json': {'example': [[]]}}}}}, 
        'post': {'security': [{'OAuth2PasswordBearer': []}], 'tags': [rf'{tag}'], 'summary': rf'Post a new {title}', 
                'description': rf'Create a new {title} data', 
                'requestBody': {'required': True, 'description': rf'Request body to create {title} data', 
                                    'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 
                                                                    'example': []}}}, 
                                                                    'responses': {'200': {'description': rf'{title} data is added', 'content': {'application/json': {'example': {'message': rf'The {title} data was successfully added'}}}}}}},
                                                                                        
        rf'/{endpoint_name}/'+"{id}":
            {'parameters': [{'in': 'path', 'name': 'id', 'required': True, 'description': rf'id of {title} data to fetch', 'schema':{'type': 'interger'}}],
            'get':{'security': [{'Authorization':[]}], 'tags': [rf'{tag}'], 'summary': rf'Get {title}', 
                'responses':{'200': {'description': rf'get {title} data to be returned', 'content': {'application/json': {'example': [[]]}}}}},
            'put': {'security': [{'Authorization':[]}], 'tags': [rf'{tag}'], 'summary': rf'Update a new {title} data', 'description': rf'Update a new {title} data', 
                        'requestBody': {'required': True, 'description': rf'Request body to update {title} data', 
                                        'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 'example': []}}},
                                                    'responses': {'200': {'description': rf'{title} data is completely updated', 'content': {'application/json': {'example': {'message': rf'The {title} data is completely updated'}}}}}},

            'patch': {'security': [{'Authorization':[]}], 'tags': [rf'{tag}'], 'summary': rf'Update a new {title} data', 'description': rf'Update a new {title} data', 
                                    'requestBody': {'required': True, 'description': rf'Request body to update {title} data', 
                                                    'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 'example': []}}},
                                                                'responses': {'200': {'description': rf'one row in {title} data is updated', 'content': {'application/json': {'example': {'message': rf'The one row in {title} data is updated'}}}}}},
                                                                
            'delete': {'security': [{'Authorization':[]}], 'tags': [rf'{tag}'], 'summary': rf'Delete {title} data', 'description': rf'Delete {title} data',
                        'responses':{'200': {'description': rf'{title} data is deleted', 'content': {'application/json': {'example': {'message': rf'The {title} data was successfully deleted'}}}}}}},

        rf'/{endpoint_name}/'+"{num}": 
        {'parameters': [{'in': 'path', 'name': 'num', 'required': True, 'description': rf'number of dummy data to add', 'schema':{'type': 'interger'}}],
            'get':{'security': [{'Authorization':[]}], 'tags': [rf'{tag}'+" Dummy Data"], 'summary': rf'Add dummy data to {title}', 
                'responses':{'200': {'description': rf'add data to {title}', 'content': {'application/json': {}}}}}}}
       
       schema[title] = {"type": "object", "properties":{}}

       cur_schema = {title:{
           "type": "object",
           "properties": {}
       }}
       
       for field in fields:
           filler['example'][field] = "LoremIpsum"
           cur_field = {field:{'type': "string"}}
           cur_schema[title]['properties'].update(cur_field)

    endpoint[rf'/{endpoint_name}']['get']['responses']['200']['content']['application/json'].update(table_filler(filler))

    cur_table['components']["schemas"].update(cur_schema)

    endpoint_holder["paths"].update(endpoint)
    base.update(endpoint_holder)
    base.update(cur_table)

    print(cur_table)



    with open(f"/Users/darrenmp/Documents/vscode/db-ppdm-copy/permen_swagger/{title}.yaml", "w") as json_file:
        yaml.dump(base, json_file, sort_keys=False, default_flow_style=False)


    
           
          
               
       




        
    #     #    endpoints = {rf'/{title_name}': 
    #               {'get':{'security': [{'Authorization':[]}], 'tags': ['Query'], 'summary': rf'Get {title}', 
    #                        'responses':{'200': {'description': rf'get {title} data to be returned', 'content': {'application/json': {'example': [{}]}}}}}, 
    #               'post': {'security': [{'OAuth2PasswordBearer': []}], 'tags': ['Create'], 'summary': rf'Post a new {title}', 
    #                        'description': rf'Create a new {title} data', 
    #                        'requestBody': {'required': True, 'description': rf'Request body to create {title} data', 
    #                                          'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 
    #                                                                           'example': {}}}}, 
    #                                                                           'responses': {'200': {'description': rf'{title} data is added', 'content': {'application/json': {'example': {'message': rf'The {title} data was successfully added'}}}}}}}, 
    #               rf'/{title_name}/{afe}':
    #               {'parameters': [{'in': 'path', 'name': 'afe', 'required': True, 'description': rf'afe of {title} data to fetch', 'schema':{'type': 'string'}}],
    #               'put': {'security': [{'Authorization':[]}], 'tags': ['Put'], 'summary': rf'Update a new {title} data', 'description': rf'Update a new {title} data', 
    #                           'requestBody': {'required': True, 'description': rf'Request body to update {title} data', 
    #                                          'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 'example': {}}}},
    #                                                      'responses': {'200': {'description': rf'{title} data is completely updated', 'content': {'application/json': {'example': {'message': rf'The {title} data is completely updated'}}}}}},

    #               'patch': {'security': [{'Authorization':[]}], 'tags': ['Patch'], 'summary': rf'Update a new {title} data', 'description': rf'Update a new {title} data', 
    #                                       'requestBody': {'required': True, 'description': rf'Request body to update {title} data', 
    #                                                      'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{title}'}, 'example': {}}}},
    #                                                                  'responses': {'200': {'description': rf'one row in {title} data is updated', 'content': {'application/json': {'example': {'message': rf'The one row in {title} data is updated'}}}}}},
                                                                     
    #               'delete': {'security': [{'Authorization':[]}], 'tags': ['Delete'], 'summary': rf'Delete {title} data', 'description': rf'Delete {title} data',
    #                           'responses':{'200': {'description': rf'{title} data is deleted', 'content': {'application/json': {'example': {'message': rf'The {title} data was successfully deleted'}}}}}}}}
           



    # #    print(fields)