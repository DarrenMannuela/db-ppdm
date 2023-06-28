package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-digitalimagewelllog/dto"
    "github.com/DarrenMannuela/svc-datatype-digitalimagewelllog/utils"

    "github.com/gofiber/fiber/v2"
)
func GetDigitalImageWellLog(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM digital_image_well_log_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_image_well_log    
    
        for rows.Next() {
    
            var diwlt dto.Digital_image_well_log
            if err := rows.Scan(&diwlt.Id, &diwlt.Ba_long_name, &diwlt.Ba_type, &diwlt.Area_id, &diwlt.Area_type, &diwlt.Field_name, &diwlt.Well_name, &diwlt.Uwi, &diwlt.Logging_company, &diwlt.Scale_ratio, &diwlt.Media_type, &diwlt.Digital_format, &diwlt.Well_log_class_id, &diwlt.Log_title, &diwlt.Report_log_run, &diwlt.Trip_date, &diwlt.Top_depth, &diwlt.Top_depth_ouom, &diwlt.Base_depth, &diwlt.Base_depth_ouom, &diwlt.Original_file_name, &diwlt.Password, &diwlt.Digital_size, &diwlt.Digital_size_uom, &diwlt.Ba_long_name_2, &diwlt.Ba_type_2, &diwlt.Data_store_name, &diwlt.Remark, &diwlt.Source, &diwlt.Qc_status, &diwlt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, diwlt)
        }
    
    return c.JSON(results)
}
func SetDigitalImageWellLog(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var diwlt dto.Digital_image_well_log
    setDefaults(&diwlt)

    if err := c.BodyParser(&diwlt); err != nil{
        return err
    }
    
    diwlt.Create_date = formatDateString(diwlt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO digital_image_well_log_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, uwi, logging_company, scale_ratio, media_type, digital_format, well_log_class_id, log_title, report_log_run, trip_date, top_depth, top_depth_ouom, base_depth, base_depth_ouom, original_file_name, password, digital_size, digital_size_uom, ba_long_name_2, ba_type_2, data_store_name, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30) RETURNING id INTO :31`, &diwlt.Ba_long_name, &diwlt.Ba_type, &diwlt.Area_id, &diwlt.Area_type, &diwlt.Field_name, &diwlt.Well_name, &diwlt.Uwi, &diwlt.Logging_company, &diwlt.Scale_ratio, &diwlt.Media_type, &diwlt.Digital_format, &diwlt.Well_log_class_id, &diwlt.Log_title, &diwlt.Report_log_run, &diwlt.Trip_date, &diwlt.Top_depth, &diwlt.Top_depth_ouom, &diwlt.Base_depth, &diwlt.Base_depth_ouom, &diwlt.Original_file_name, &diwlt.Password, &diwlt.Digital_size, &diwlt.Digital_size_uom, &diwlt.Ba_long_name_2, &diwlt.Ba_type_2, &diwlt.Data_store_name, &diwlt.Remark, &diwlt.Source, &diwlt.Qc_status, &diwlt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(DIGITAL_IMAGE_WELL_LOG_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetDigitalImageWellLogById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM digital_image_well_log_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_image_well_log    
    
        for rows.Next() {
    
            var diwlt dto.Digital_image_well_log
            if err := rows.Scan(&diwlt.Id, &diwlt.Ba_long_name, &diwlt.Ba_type, &diwlt.Area_id, &diwlt.Area_type, &diwlt.Field_name, &diwlt.Well_name, &diwlt.Uwi, &diwlt.Logging_company, &diwlt.Scale_ratio, &diwlt.Media_type, &diwlt.Digital_format, &diwlt.Well_log_class_id, &diwlt.Log_title, &diwlt.Report_log_run, &diwlt.Trip_date, &diwlt.Top_depth, &diwlt.Top_depth_ouom, &diwlt.Base_depth, &diwlt.Base_depth_ouom, &diwlt.Original_file_name, &diwlt.Password, &diwlt.Digital_size, &diwlt.Digital_size_uom, &diwlt.Ba_long_name_2, &diwlt.Ba_type_2, &diwlt.Data_store_name, &diwlt.Remark, &diwlt.Source, &diwlt.Qc_status, &diwlt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, diwlt)
        }
    
    return c.JSON(results)
}
func PutDigitalImageWellLog(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var diwlt dto.Digital_image_well_log
    setDefaults(&diwlt)

    if err := c.BodyParser(&diwlt); err != nil{
        return err
    }
    
    diwlt.Create_date = formatDateString(diwlt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM digital_image_well_log_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE digital_image_well_log_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, uwi = :7, logging_company = :8, scale_ratio = :9, media_type = :10, digital_format = :11, well_log_class_id = :12, log_title = :13, report_log_run = :14, trip_date = :15, top_depth = :16, top_depth_ouom = :17, base_depth = :18, base_depth_ouom = :19, original_file_name = :20, password = :21, digital_size = :22, digital_size_uom = :23, ba_long_name_2 = :24, ba_type_2 = :25, data_store_name = :26, remark = :27, source = :28, qc_status = :29, checked_by_ba_id = :30 WHERE id = :31`, &diwlt.Ba_long_name, &diwlt.Ba_type, &diwlt.Area_id, &diwlt.Area_type, &diwlt.Field_name, &diwlt.Well_name, &diwlt.Uwi, &diwlt.Logging_company, &diwlt.Scale_ratio, &diwlt.Media_type, &diwlt.Digital_format, &diwlt.Well_log_class_id, &diwlt.Log_title, &diwlt.Report_log_run, &diwlt.Trip_date, &diwlt.Top_depth, &diwlt.Top_depth_ouom, &diwlt.Base_depth, &diwlt.Base_depth_ouom, &diwlt.Original_file_name, &diwlt.Password, &diwlt.Digital_size, &diwlt.Digital_size_uom, &diwlt.Ba_long_name_2, &diwlt.Ba_type_2, &diwlt.Data_store_name, &diwlt.Remark, &diwlt.Source, &diwlt.Qc_status, &diwlt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(DIGITAL_IMAGE_WELL_LOG_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteDigitalImageWellLog(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT digital_image_well_log_id FROM digital_image_well_log_workspace WHERE digital_image_well_log_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM digital_image_well_log_workspace WHERE digital_image_well_log_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM digital_image_well_log_table WHERE id = :1`, id)
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
func PatchDigitalImageWellLog(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var diwlt dto.Digital_image_well_log
    setDefaults(&diwlt)

    if err := c.BodyParser(&diwlt); err != nil{
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
    err = db.QueryRow(`SELECT digital_image_well_log_id FROM digital_image_well_log_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if diwlt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET ba_long_name = :1 WHERE id = :2`, diwlt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET ba_type = :1 WHERE id = :2`, diwlt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Area_id != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET area_id = :1 WHERE id = :2`, diwlt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Area_type != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET area_type = :1 WHERE id = :2`, diwlt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Field_name != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET field_name = :1 WHERE id = :2`, diwlt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Well_name != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET well_name = :1 WHERE id = :2`, diwlt.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Uwi != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET uwi = :1 WHERE id = :2`, diwlt.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Logging_company != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET logging_company = :1 WHERE id = :2`, diwlt.Logging_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Scale_ratio != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET scale_ratio = :1 WHERE id = :2`, diwlt.Scale_ratio, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Media_type != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET media_type = :1 WHERE id = :2`, diwlt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET digital_format = :1 WHERE id = :2`, diwlt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Well_log_class_id != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET well_log_class_id = :1 WHERE id = :2`, diwlt.Well_log_class_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Log_title != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET log_title = :1 WHERE id = :2`, diwlt.Log_title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Report_log_run != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET report_log_run = :1 WHERE id = :2`, diwlt.Report_log_run, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Trip_date != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET trip_date = :1 WHERE id = :2`, diwlt.Trip_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Top_depth != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET top_depth = :1 WHERE id = :2`, diwlt.Top_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Top_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET top_depth_ouom = :1 WHERE id = :2`, diwlt.Top_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Base_depth != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET base_depth = :1 WHERE id = :2`, diwlt.Base_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Base_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET base_depth_ouom = :1 WHERE id = :2`, diwlt.Base_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET original_file_name = :1 WHERE id = :2`, diwlt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Password != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET password = :1 WHERE id = :2`, diwlt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET digital_size = :1 WHERE id = :2`, diwlt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET digital_size_uom = :1 WHERE id = :2`, diwlt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET ba_long_name_2 = :1 WHERE id = :2`, diwlt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET ba_type_2 = :1 WHERE id = :2`, diwlt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET data_store_name = :1 WHERE id = :2`, diwlt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Remark != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET remark = :1 WHERE id = :2`, diwlt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Source != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET source = :1 WHERE id = :2`, diwlt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET qc_status = :1 WHERE id = :2`, diwlt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if diwlt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE digital_image_well_log_table SET checked_by_ba_id = :1 WHERE id = :2`, diwlt.Checked_by_ba_id, id)
        
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

    