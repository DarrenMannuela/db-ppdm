# importing the modules
import os
import shutil

# Providing the folder path
origin = 'C:/Users/johan/VSCode/pt_gtn/postgres_39/go/'
target = 'C:/Users/johan/VSCode/pt_gtn/bibliography/'

# Fetching the list of all the files
files =  ["Area.go", "SourceDocument.go", "RDocumentType.go", "RLanguage.go", "RPublicationName.go", 
"RSource.go", "RPpdmRowQuality.go", "SourceDocAuthor.go", "BusinessAssociate.go", "RAuthorType.go", "RBaCategory.go", 
"RBaType.go", "RBaStatus.go", "RmDataStore.go", "BaAddress.go", "RmDataStoreHier.go", "RmDataStoreHierLevel.go", 
"RStoreStatus.go", "RAddressType.go", "RServiceQuality.go", "RDataStoreType.go", "CsGeodeticDatum.go", "CsEllipsoid.go",
"CsPrincipalMeridian.go", "CsPrimeMeridian.go", "CsCoordinateSystem.go", "PpdmUnitOfMeasure.go", "RCoordSystemType.go",
"RProjectionType.go", "RVerticalDatumType.go", "RCoordCapture.go", "RCoordCompute.go", "RCoordQuality.go", "CsCoordTransform.go",
"RPpdmUomUsage.go", "PpdmMeasurementSystem.go", "PpdmQuantity.go", "RCsTransformType.go", "CsCoordAcquisition.go", "RAreaType.go",
"RDatumOrigin.go"]

# Fetching all the files to directory
for file_name in files:
   shutil.copy(origin+file_name, target+file_name)
print("Files are copied successfully")
