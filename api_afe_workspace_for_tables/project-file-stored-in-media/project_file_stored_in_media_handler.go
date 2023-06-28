package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-projectfilestoredinmedia/dto"
    "github.com/DarrenMannuela/svc-datatype-projectfilestoredinmedia/utils"

    "github.com/gofiber/fiber/v2"
)
func GetProjectFileStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM project_file_stored_in_media_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Project_file_stored_in_media    
    
        for rows.Next() {
    
            var pfsimt dto.Project_file_stored_in_media
            if err := rows.Scan(&pfsimt.Id, &pfsimt.Ba_long_name, &pfsimt.Ba_type, &pfsimt.Area_id, &pfsimt.Area_type, &pfsimt.Field_name, &pfsimt.Project_name, &pfsimt.Sw_application_id, &pfsimt.Application_version, &pfsimt.Item_category, &pfsimt.Process_date, &pfsimt.Interpreter, &pfsimt.Digital_format, &pfsimt.Media_type, &pfsimt.Ba_long_name_2, &pfsimt.Ba_type_2, &pfsimt.Data_store_name, &pfsimt.Data_store_type, &pfsimt.Location_id, &pfsimt.Remark, &pfsimt.Source, &pfsimt.Qc_status, &pfsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, pfsimt)
        }
    
    return c.JSON(results)
}
func SetProjectFileStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var pfsimt dto.Project_file_stored_in_media
    setDefaults(&pfsimt)

    if err := c.BodyParser(&pfsimt); err != nil{
        return err
    }
    
    pfsimt.Create_date = formatDateString(pfsimt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO project_file_stored_in_media_table (ba_long_name, ba_type, area_id, area_type, field_name, project_name, sw_application_id, application_version, item_category, process_date, interpreter, digital_format, media_type, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22) RETURNING id INTO :23`, &pfsimt.Ba_long_name, &pfsimt.Ba_type, &pfsimt.Area_id, &pfsimt.Area_type, &pfsimt.Field_name, &pfsimt.Project_name, &pfsimt.Sw_application_id, &pfsimt.Application_version, &pfsimt.Item_category, &pfsimt.Process_date, &pfsimt.Interpreter, &pfsimt.Digital_format, &pfsimt.Media_type, &pfsimt.Ba_long_name_2, &pfsimt.Ba_type_2, &pfsimt.Data_store_name, &pfsimt.Data_store_type, &pfsimt.Location_id, &pfsimt.Remark, &pfsimt.Source, &pfsimt.Qc_status, &pfsimt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(PROJECT_FILE_STORED_IN_MEDIA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetProjectFileStoredInMediaById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM project_file_stored_in_media_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Project_file_stored_in_media    
    
        for rows.Next() {
    
            var pfsimt dto.Project_file_stored_in_media
            if err := rows.Scan(&pfsimt.Id, &pfsimt.Ba_long_name, &pfsimt.Ba_type, &pfsimt.Area_id, &pfsimt.Area_type, &pfsimt.Field_name, &pfsimt.Project_name, &pfsimt.Sw_application_id, &pfsimt.Application_version, &pfsimt.Item_category, &pfsimt.Process_date, &pfsimt.Interpreter, &pfsimt.Digital_format, &pfsimt.Media_type, &pfsimt.Ba_long_name_2, &pfsimt.Ba_type_2, &pfsimt.Data_store_name, &pfsimt.Data_store_type, &pfsimt.Location_id, &pfsimt.Remark, &pfsimt.Source, &pfsimt.Qc_status, &pfsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, pfsimt)
        }
    
    return c.JSON(results)
}
func PutProjectFileStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var pfsimt dto.Project_file_stored_in_media
    setDefaults(&pfsimt)

    if err := c.BodyParser(&pfsimt); err != nil{
        return err
    }
    
    pfsimt.Create_date = formatDateString(pfsimt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM project_file_stored_in_media_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, project_name = :6, sw_application_id = :7, application_version = :8, item_category = :9, process_date = :10, interpreter = :11, digital_format = :12, media_type = :13, ba_long_name_2 = :14, ba_type_2 = :15, data_store_name = :16, data_store_type = :17, location_id = :18, remark = :19, source = :20, qc_status = :21, checked_by_ba_id = :22 WHERE id = :23`, &pfsimt.Ba_long_name, &pfsimt.Ba_type, &pfsimt.Area_id, &pfsimt.Area_type, &pfsimt.Field_name, &pfsimt.Project_name, &pfsimt.Sw_application_id, &pfsimt.Application_version, &pfsimt.Item_category, &pfsimt.Process_date, &pfsimt.Interpreter, &pfsimt.Digital_format, &pfsimt.Media_type, &pfsimt.Ba_long_name_2, &pfsimt.Ba_type_2, &pfsimt.Data_store_name, &pfsimt.Data_store_type, &pfsimt.Location_id, &pfsimt.Remark, &pfsimt.Source, &pfsimt.Qc_status, &pfsimt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(PROJECT_FILE_STORED_IN_MEDIA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteProjectFileStoredInMedia(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT project_file_stored_in_media_id FROM project_file_stored_in_media_workspace WHERE project_file_stored_in_media_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM project_file_stored_in_media_workspace WHERE project_file_stored_in_media_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM project_file_stored_in_media_table WHERE id = :1`, id)
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
func PatchProjectFileStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var pfsimt dto.Project_file_stored_in_media
    setDefaults(&pfsimt)

    if err := c.BodyParser(&pfsimt); err != nil{
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
    err = db.QueryRow(`SELECT project_file_stored_in_media_id FROM project_file_stored_in_media_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if pfsimt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET ba_long_name = :1 WHERE id = :2`, pfsimt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET ba_type = :1 WHERE id = :2`, pfsimt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Area_id != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET area_id = :1 WHERE id = :2`, pfsimt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Area_type != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET area_type = :1 WHERE id = :2`, pfsimt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Field_name != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET field_name = :1 WHERE id = :2`, pfsimt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Project_name != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET project_name = :1 WHERE id = :2`, pfsimt.Project_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Sw_application_id != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET sw_application_id = :1 WHERE id = :2`, pfsimt.Sw_application_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Application_version != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET application_version = :1 WHERE id = :2`, pfsimt.Application_version, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Item_category != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET item_category = :1 WHERE id = :2`, pfsimt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Process_date != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET process_date = :1 WHERE id = :2`, pfsimt.Process_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Interpreter != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET interpreter = :1 WHERE id = :2`, pfsimt.Interpreter, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET digital_format = :1 WHERE id = :2`, pfsimt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Media_type != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET media_type = :1 WHERE id = :2`, pfsimt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET ba_long_name_2 = :1 WHERE id = :2`, pfsimt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET ba_type_2 = :1 WHERE id = :2`, pfsimt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET data_store_name = :1 WHERE id = :2`, pfsimt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET data_store_type = :1 WHERE id = :2`, pfsimt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Location_id != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET location_id = :1 WHERE id = :2`, pfsimt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Remark != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET remark = :1 WHERE id = :2`, pfsimt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Source != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET source = :1 WHERE id = :2`, pfsimt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET qc_status = :1 WHERE id = :2`, pfsimt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pfsimt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE project_file_stored_in_media_table SET checked_by_ba_id = :1 WHERE id = :2`, pfsimt.Checked_by_ba_id, id)
        
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

    