package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-digitalprojectfile/dto"
    "github.com/DarrenMannuela/svc-datatype-digitalprojectfile/utils"

    "github.com/gofiber/fiber/v2"
)
func GetDigitalProjectFile(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM digital_project_file_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_project_file    
    
        for rows.Next() {
    
            var dpft dto.Digital_project_file
            if err := rows.Scan(&dpft.Id, &dpft.Ba_long_name, &dpft.Ba_type, &dpft.Area_id, &dpft.Area_type, &dpft.Field_name, &dpft.Project_name, &dpft.Sw_application_id, &dpft.Application_version, &dpft.Item_category, &dpft.Process_date, &dpft.Interpreter, &dpft.Digital_format, &dpft.Media_type, &dpft.Original_file_name, &dpft.Password, &dpft.Digital_size, &dpft.Digital_size_uom, &dpft.Ba_long_name_2, &dpft.Ba_type_2, &dpft.Data_store_name, &dpft.Remark, &dpft.Source, &dpft.Qc_status, &dpft.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dpft)
        }
    
    return c.JSON(results)
}
func SetDigitalProjectFile(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var dpft dto.Digital_project_file
    setDefaults(&dpft)

    if err := c.BodyParser(&dpft); err != nil{
        return err
    }
    
    dpft.Create_date = formatDateString(dpft.Create_date)
    

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}()

    var generatedID int64    
    _, err = tx.Exec(`INSERT INTO digital_project_file_table (ba_long_name, ba_type, area_id, area_type, field_name, project_name, sw_application_id, application_version, item_category, process_date, interpreter, digital_format, media_type, original_file_name, password, digital_size, digital_size_uom, ba_long_name_2, ba_type_2, data_store_name, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24) RETURNING id INTO :25`, &dpft.Ba_long_name, &dpft.Ba_type, &dpft.Area_id, &dpft.Area_type, &dpft.Field_name, &dpft.Project_name, &dpft.Sw_application_id, &dpft.Application_version, &dpft.Item_category, &dpft.Process_date, &dpft.Interpreter, &dpft.Digital_format, &dpft.Media_type, &dpft.Original_file_name, &dpft.Password, &dpft.Digital_size, &dpft.Digital_size_uom, &dpft.Ba_long_name_2, &dpft.Ba_type_2, &dpft.Data_store_name, &dpft.Remark, &dpft.Source, &dpft.Qc_status, &dpft.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(DIGITAL_PROJECT_FILE_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetDigitalProjectFileById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM digital_project_file_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_project_file    
    
        for rows.Next() {
    
            var dpft dto.Digital_project_file
            if err := rows.Scan(&dpft.Id, &dpft.Ba_long_name, &dpft.Ba_type, &dpft.Area_id, &dpft.Area_type, &dpft.Field_name, &dpft.Project_name, &dpft.Sw_application_id, &dpft.Application_version, &dpft.Item_category, &dpft.Process_date, &dpft.Interpreter, &dpft.Digital_format, &dpft.Media_type, &dpft.Original_file_name, &dpft.Password, &dpft.Digital_size, &dpft.Digital_size_uom, &dpft.Ba_long_name_2, &dpft.Ba_type_2, &dpft.Data_store_name, &dpft.Remark, &dpft.Source, &dpft.Qc_status, &dpft.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dpft)
        }
    
    return c.JSON(results)
}
func PutDigitalProjectFile(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dpft dto.Digital_project_file
    setDefaults(&dpft)

    if err := c.BodyParser(&dpft); err != nil{
        return err
    }
    
    dpft.Create_date = formatDateString(dpft.Create_date)
    

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}()
    
    
    var idExist string
    err = db.QueryRow(`SELECT * FROM digital_project_file_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE digital_project_file_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, project_name = :6, sw_application_id = :7, application_version = :8, item_category = :9, process_date = :10, interpreter = :11, digital_format = :12, media_type = :13, original_file_name = :14, password = :15, digital_size = :16, digital_size_uom = :17, ba_long_name_2 = :18, ba_type_2 = :19, data_store_name = :20, remark = :21, source = :22, qc_status = :23, checked_by_ba_id = :24 WHERE id = :25`, &dpft.Ba_long_name, &dpft.Ba_type, &dpft.Area_id, &dpft.Area_type, &dpft.Field_name, &dpft.Project_name, &dpft.Sw_application_id, &dpft.Application_version, &dpft.Item_category, &dpft.Process_date, &dpft.Interpreter, &dpft.Digital_format, &dpft.Media_type, &dpft.Original_file_name, &dpft.Password, &dpft.Digital_size, &dpft.Digital_size_uom, &dpft.Ba_long_name_2, &dpft.Ba_type_2, &dpft.Data_store_name, &dpft.Remark, &dpft.Source, &dpft.Qc_status, &dpft.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(DIGITAL_PROJECT_FILE_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteDigitalProjectFile(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}()
    
    
    var idExist string
	_ = tx.QueryRow(`SELECT digital_project_file_id FROM digital_project_file_workspace WHERE digital_project_file_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM digital_project_file_workspace WHERE digital_project_file_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM digital_project_file_table WHERE id = :1`, id)
	if err != nil{
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
    
    return c.SendStatus(fiber.StatusOK)
}
func PatchDigitalProjectFile(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dpft dto.Digital_project_file
    setDefaults(&dpft)

    if err := c.BodyParser(&dpft); err != nil{
        return err
    }
    
    tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}()
    
    
    var idExist string
    err = db.QueryRow(`SELECT digital_project_file_id FROM digital_project_file_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if dpft.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, dpft.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, dpft.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, dpft.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, dpft.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Field_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_name = :1 WHERE id = :2`, dpft.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Project_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET project_name = :1 WHERE id = :2`, dpft.Project_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Sw_application_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sw_application_id = :1 WHERE id = :2`, dpft.Sw_application_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Application_version != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET application_version = :1 WHERE id = :2`, dpft.Application_version, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Item_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_category = :1 WHERE id = :2`, dpft.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Process_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET process_date = :1 WHERE id = :2`, dpft.Process_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Interpreter != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET interpreter = :1 WHERE id = :2`, dpft.Interpreter, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, dpft.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, dpft.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, dpft.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, dpft.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, dpft.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, dpft.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, dpft.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, dpft.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, dpft.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, dpft.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, dpft.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, dpft.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dpft.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, dpft.Checked_by_ba_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
    }
    err = tx.Commit()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(map[string]string{"message": "Record updated successfully"})
}

    