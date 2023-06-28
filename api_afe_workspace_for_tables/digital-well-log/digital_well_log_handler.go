package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-digitalwelllog/dto"
    "github.com/DarrenMannuela/svc-datatype-digitalwelllog/utils"

    "github.com/gofiber/fiber/v2"
)
func GetDigitalWellLog(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM digital_well_log_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_well_log    
    
        for rows.Next() {
    
            var dwlt dto.Digital_well_log
            if err := rows.Scan(&dwlt.Id, &dwlt.Ba_long_name, &dwlt.Ba_type, &dwlt.Area_id, &dwlt.Area_type, &dwlt.Field_name, &dwlt.Well_name, &dwlt.Uwi, &dwlt.Logging_company, &dwlt.Media_type, &dwlt.Well_log_class_id, &dwlt.Digital_format, &dwlt.Report_log_run, &dwlt.Trip_date, &dwlt.Top_depth, &dwlt.Top_depth_ouom, &dwlt.Base_depth, &dwlt.Base_depth_ouom, &dwlt.Original_file_name, &dwlt.Password, &dwlt.Digital_size, &dwlt.Digital_size_uom, &dwlt.Ba_long_name_2, &dwlt.Ba_type_2, &dwlt.Data_store_name, &dwlt.Remark, &dwlt.Source, &dwlt.Qc_status, &dwlt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dwlt)
        }
    
    return c.JSON(results)
}
func SetDigitalWellLog(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var dwlt dto.Digital_well_log
    setDefaults(&dwlt)

    if err := c.BodyParser(&dwlt); err != nil{
        return err
    }
    
    dwlt.Create_date = formatDateString(dwlt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO digital_well_log_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, uwi, logging_company, media_type, well_log_class_id, digital_format, report_log_run, trip_date, top_depth, top_depth_ouom, base_depth, base_depth_ouom, original_file_name, password, digital_size, digital_size_uom, ba_long_name_2, ba_type_2, data_store_name, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28) RETURNING id INTO :29`, &dwlt.Ba_long_name, &dwlt.Ba_type, &dwlt.Area_id, &dwlt.Area_type, &dwlt.Field_name, &dwlt.Well_name, &dwlt.Uwi, &dwlt.Logging_company, &dwlt.Media_type, &dwlt.Well_log_class_id, &dwlt.Digital_format, &dwlt.Report_log_run, &dwlt.Trip_date, &dwlt.Top_depth, &dwlt.Top_depth_ouom, &dwlt.Base_depth, &dwlt.Base_depth_ouom, &dwlt.Original_file_name, &dwlt.Password, &dwlt.Digital_size, &dwlt.Digital_size_uom, &dwlt.Ba_long_name_2, &dwlt.Ba_type_2, &dwlt.Data_store_name, &dwlt.Remark, &dwlt.Source, &dwlt.Qc_status, &dwlt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(DIGITAL_WELL_LOG_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetDigitalWellLogById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM digital_well_log_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_well_log    
    
        for rows.Next() {
    
            var dwlt dto.Digital_well_log
            if err := rows.Scan(&dwlt.Id, &dwlt.Ba_long_name, &dwlt.Ba_type, &dwlt.Area_id, &dwlt.Area_type, &dwlt.Field_name, &dwlt.Well_name, &dwlt.Uwi, &dwlt.Logging_company, &dwlt.Media_type, &dwlt.Well_log_class_id, &dwlt.Digital_format, &dwlt.Report_log_run, &dwlt.Trip_date, &dwlt.Top_depth, &dwlt.Top_depth_ouom, &dwlt.Base_depth, &dwlt.Base_depth_ouom, &dwlt.Original_file_name, &dwlt.Password, &dwlt.Digital_size, &dwlt.Digital_size_uom, &dwlt.Ba_long_name_2, &dwlt.Ba_type_2, &dwlt.Data_store_name, &dwlt.Remark, &dwlt.Source, &dwlt.Qc_status, &dwlt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dwlt)
        }
    
    return c.JSON(results)
}
func PutDigitalWellLog(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dwlt dto.Digital_well_log
    setDefaults(&dwlt)

    if err := c.BodyParser(&dwlt); err != nil{
        return err
    }
    
    dwlt.Create_date = formatDateString(dwlt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM digital_well_log_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE digital_well_log_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, uwi = :7, logging_company = :8, media_type = :9, well_log_class_id = :10, digital_format = :11, report_log_run = :12, trip_date = :13, top_depth = :14, top_depth_ouom = :15, base_depth = :16, base_depth_ouom = :17, original_file_name = :18, password = :19, digital_size = :20, digital_size_uom = :21, ba_long_name_2 = :22, ba_type_2 = :23, data_store_name = :24, remark = :25, source = :26, qc_status = :27, checked_by_ba_id = :28 WHERE id = :29`, &dwlt.Ba_long_name, &dwlt.Ba_type, &dwlt.Area_id, &dwlt.Area_type, &dwlt.Field_name, &dwlt.Well_name, &dwlt.Uwi, &dwlt.Logging_company, &dwlt.Media_type, &dwlt.Well_log_class_id, &dwlt.Digital_format, &dwlt.Report_log_run, &dwlt.Trip_date, &dwlt.Top_depth, &dwlt.Top_depth_ouom, &dwlt.Base_depth, &dwlt.Base_depth_ouom, &dwlt.Original_file_name, &dwlt.Password, &dwlt.Digital_size, &dwlt.Digital_size_uom, &dwlt.Ba_long_name_2, &dwlt.Ba_type_2, &dwlt.Data_store_name, &dwlt.Remark, &dwlt.Source, &dwlt.Qc_status, &dwlt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(DIGITAL_WELL_LOG_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteDigitalWellLog(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT digital_well_log_id FROM digital_well_log_workspace WHERE digital_well_log_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM digital_well_log_workspace WHERE digital_well_log_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM digital_well_log_table WHERE id = :1`, id)
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
func PatchDigitalWellLog(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dwlt dto.Digital_well_log
    setDefaults(&dwlt)

    if err := c.BodyParser(&dwlt); err != nil{
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
    err = db.QueryRow(`SELECT digital_well_log_id FROM digital_well_log_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if dwlt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET ba_long_name = :1 WHERE id = :2`, dwlt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET ba_type = :1 WHERE id = :2`, dwlt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Area_id != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET area_id = :1 WHERE id = :2`, dwlt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Area_type != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET area_type = :1 WHERE id = :2`, dwlt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Field_name != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET field_name = :1 WHERE id = :2`, dwlt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Well_name != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET well_name = :1 WHERE id = :2`, dwlt.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Uwi != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET uwi = :1 WHERE id = :2`, dwlt.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Logging_company != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET logging_company = :1 WHERE id = :2`, dwlt.Logging_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Media_type != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET media_type = :1 WHERE id = :2`, dwlt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Well_log_class_id != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET well_log_class_id = :1 WHERE id = :2`, dwlt.Well_log_class_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET digital_format = :1 WHERE id = :2`, dwlt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Report_log_run != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET report_log_run = :1 WHERE id = :2`, dwlt.Report_log_run, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Trip_date != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET trip_date = :1 WHERE id = :2`, dwlt.Trip_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Top_depth != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET top_depth = :1 WHERE id = :2`, dwlt.Top_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Top_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET top_depth_ouom = :1 WHERE id = :2`, dwlt.Top_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Base_depth != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET base_depth = :1 WHERE id = :2`, dwlt.Base_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Base_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET base_depth_ouom = :1 WHERE id = :2`, dwlt.Base_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET original_file_name = :1 WHERE id = :2`, dwlt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Password != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET password = :1 WHERE id = :2`, dwlt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET digital_size = :1 WHERE id = :2`, dwlt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET digital_size_uom = :1 WHERE id = :2`, dwlt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET ba_long_name_2 = :1 WHERE id = :2`, dwlt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET ba_type_2 = :1 WHERE id = :2`, dwlt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET data_store_name = :1 WHERE id = :2`, dwlt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Remark != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET remark = :1 WHERE id = :2`, dwlt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Source != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET source = :1 WHERE id = :2`, dwlt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET qc_status = :1 WHERE id = :2`, dwlt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwlt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE digital_well_log_table SET checked_by_ba_id = :1 WHERE id = :2`, dwlt.Checked_by_ba_id, id)
        
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

    