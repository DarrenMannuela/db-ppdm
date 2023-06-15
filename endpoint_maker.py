import yaml
import re
import sys
import lib_platform as platform

if platform.system == "darwin":
    sys.path.insert(0, "database")
elif platform.system == "windows":
    sys.path.insert(0, "../database")

from conn import connection

id = '{id}'

def get_table_name():
    cursor = connection.cursor()
    cursor.execute("SELECT table_name FROM user_tables WHERE table_name NOT LIKE '%USER%' AND table_name NOT LIKE '%ROLE%'")
    tables = [row[0].lower() for row in cursor.fetchall()]
    return tables 

get_sql = get_table_name()

for table in  get_sql:

   table_name = table.replace("_", "-")
   
   with open(rf'schema/{table}.yaml', 'r') as f:
      data = yaml.load(f, Loader=yaml.FullLoader)


   endpoints= {rf'/{table_name}': 
                  {'get':{'security': [{'Authorization':[]}], 'tags': ['Query'], 'summary': rf'Get {table}', 
                           'responses':{'200': {'description': rf'get {table} data to be returned', 'content': {'application/json': {'example': [{}]}}}}}, 
                  'post': {'security': [{'Authorization': []}], 'tags': ['Create'], 'summary': rf'Post a new {table}', 
                           'description': rf'Create a new {table} data', 
                           'requestBody': {'required': True, 'description': rf'Request body to create {table} data', 
                                             'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{table}'}, 
                                                                              'example': {}}}}, 
                                                                              'responses': {'200': {'description': rf'{table} data is added', 'content': {'application/json': {'example': {'message': rf'The {table} data was successfully added'}}}}}}}, 
                  rf'/{table_name}/{id}':
                  {'parameters': [{'in': 'path', 'name': 'id', 'required': True, 'description': rf'ID of {table} data to fetch', 'schema':{'type': 'string'}}],
                  'put': {'security': [{'Authorization':[]}], 'tags': ['Put'], 'summary': rf'Update a new {table} data', 'description': rf'Update a new {table} data', 
                              'requestBody': {'required': True, 'description': rf'Request body to update {table} data', 
                                             'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{table}'}, 'example': {}}}},
                                                         'responses': {'200': {'description': rf'{table} data is completely updated', 'content': {'application/json': {'example': {'message': rf'The {table} data is completely updated'}}}}}},

                  'patch': {'security': [{'Authorization':[]}], 'tags': ['Patch'], 'summary': rf'Update a new {table} data', 'description': rf'Update a new {table} data', 
                                          'requestBody': {'required': True, 'description': rf'Request body to update {table} data', 
                                                         'content': {'application/json': {'schema': {"$ref": rf'#/components/schemas/{table}'}, 'example': {}}}},
                                                                     'responses': {'200': {'description': rf'one row in {table} data is updated', 'content': {'application/json': {'example': {'message': rf'The one row in {table} data is updated'}}}}}},
                                                                     
                  'delete': {'security': [{'Authorization':[]}], 'tags': ['Delete'], 'summary': rf'Delete {table} data', 'description': rf'Delete {table} data',
                              'responses':{'200': {'description': rf'{table} data is deleted', 'content': {'application/json': {'example': {'message': rf'The {table} data was successfully deleted'}}}}}}}}





   keys = list(data[table]['items']['properties'].keys())


   for key in keys:
      cur_key = data[table]['items']['properties'][key]
      if cur_key['type'] == 'string':
         if 'maxLength' not in cur_key and 'format' not in cur_key:
            endpoints[rf'/{table_name}']['get']['responses']['200']['content']['application/json']['example'][0].update({key:'12/12/2023'})
            endpoints[rf'/{table_name}']['post']['requestBody']['content']['application/json']['example'].update({key:'12/12/2023'})
            endpoints[rf'/{table_name}/{id}']['put']['requestBody']['content']['application/json']['example'].update({key:'12/12/2023'})
            endpoints[rf'/{table_name}/{id}']['patch']['requestBody']['content']['application/json']['example'].update({key:'12/12/2023'})

         if 'maxLength' in cur_key and cur_key['maxLength'] != 1:
            endpoints[rf'/{table_name}']['get']['responses']['200']['content']['application/json']['example'][0].update({key:'Lorem Ipsum'})
            endpoints[rf'/{table_name}']['post']['requestBody']['content']['application/json']['example'].update({key:'Lorem Ipsum'})
            endpoints[rf'/{table_name}/{id}']['put']['requestBody']['content']['application/json']['example'].update({key:'Lorem Ipsum'})
            endpoints[rf'/{table_name}/{id}']['patch']['requestBody']['content']['application/json']['example'].update({key:'Lorem Ipsum'})

         if 'maxLength' in cur_key and cur_key['maxLength'] == 1:
            endpoints[rf'/{table_name}']['get']['responses']['200']['content']['application/json']['example'][0].update({key:'y'})
            endpoints[rf'/{table_name}']['post']['requestBody']['content']['application/json']['example'].update({key:'y'})    
            endpoints[rf'/{table_name}/{id}']['put']['requestBody']['content']['application/json']['example'].update({key:'y'})
            endpoints[rf'/{table_name}/{id}']['patch']['requestBody']['content']['application/json']['example'].update({key:'y'})

         if 'format' in cur_key:
            endpoints[rf'/{table_name}']['get']['responses']['200']['content']['application/json']['example'][0].update({key:'010101010'})
            endpoints[rf'/{table_name}']['post']['requestBody']['content']['application/json']['example'].update({key:'010101010'})
            endpoints[rf'/{table_name}/{id}']['put']['requestBody']['content']['application/json']['example'].update({key:'010101010'})
            endpoints[rf'/{table_name}/{id}']['patch']['requestBody']['content']['application/json']['example'].update({key:'010101010'})

      elif cur_key['type'] == 'number':

         if 'format' not in cur_key:
            endpoints[rf'/{table_name}']['get']['responses']['200']['content']['application/json']['example'][0].update({key: 1}) 
            endpoints[rf'/{table_name}']['post']['requestBody']['content']['application/json']['example'].update({key: 1})
            endpoints[rf'/{table_name}/{id}']['put']['requestBody']['content']['application/json']['example'].update({key: 1})
            endpoints[rf'/{table_name}/{id}']['patch']['requestBody']['content']['application/json']['example'].update({key: 1})           

         elif 'format' in cur_key:
               if cur_key['format'] == 'decimal':
                  endpoints[rf'/{table_name}']['get']['responses']['200']['content']['application/json']['example'][0].update({key:0.1}) 
                  endpoints[rf'/{table_name}']['post']['requestBody']['content']['application/json']['example'].update({key:0.1}) 
                  endpoints[rf'/{table_name}/{id}']['put']['requestBody']['content']['application/json']['example'].update({key:0.1}) 
                  endpoints[rf'/{table_name}/{id}']['patch']['requestBody']['content']['application/json']['example'].update({key:0.1}) 

               elif cur_key['format'] == 'double':
                  endpoints[rf'/{table_name}']['get']['responses']['200']['content']['application/json']['example'][0].update({key:1.0})
                  endpoints[rf'/{table_name}']['post']['requestBody']['content']['application/json']['example'].update({key:1.0})
                  endpoints[rf'/{table_name}/{id}']['put']['requestBody']['content']['application/json']['example'].update({key:1.0})
                  endpoints[rf'/{table_name}/{id}']['patch']['requestBody']['content']['application/json']['example'].update({key:1.0})             

               elif cur_key['format'] == 'int16':
                  endpoints[rf'/{table_name}']['get']['responses']['200']['content']['application/json']['example'][0].update({key:32767})
                  endpoints[rf'/{table_name}']['post']['requestBody']['content']['application/json']['example'].update({key:32767})
                  endpoints[rf'/{table_name}/{id}']['put']['requestBody']['content']['application/json']['example'].update({key:32767})
                  endpoints[rf'/{table_name}/{id}']['patch']['requestBody']['content']['application/json']['example'].update({key:32767})



   with open(rf"endpoints/{table}.yaml", "w") as f:
      yaml.dump(endpoints, f, sort_keys=False)









