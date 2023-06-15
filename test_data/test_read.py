import csv

# opening the CSV file
with open('/Users/darrenmp/Documents/vscode/pt_gtn/test_data/anl_accuracy_data.csv', mode ='r')as file:
   
  # reading the CSV file
  csvFile = csv.reader(file)
 
  # displaying the contents of the CSV file
