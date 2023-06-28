package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-welllogs/dto"
    "github.com/DarrenMannuela/svc-datatype-welllogs/utils"

    "github.com/gofiber/fiber/v2"
)
func GetWellLogs(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM well_logs_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_logs    
    
        for rows.Next() {
    
            var wlt dto.Well_logs
            if err := rows.Scan(&wlt.Id, &wlt.Ba_long_name, &wlt.Ba_type, &wlt.Area_id, &wlt.Area_type, &wlt.Field_name, &wlt.Well_name, &wlt.Uwi, &wlt.Logging_company, &wlt.Scale_ratio, &wlt.Media_type, &wlt.Well_log_class_id, &wlt.Log_title, &wlt.Report_log_run, &wlt.Trip_date, &wlt.Top_depth, &wlt.Top_depth_ouom, &wlt.Base_depth, &wlt.Base_depth_ouom, &wlt.Ba_long_name_2, &wlt.Ba_type_2, &wlt.Data_store_name, &wlt.Data_store_type, &wlt.Location_id, &wlt.Remark, &wlt.Source, &wlt.Qc_status, &wlt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wlt)
        }
    
    return c.JSON(results)
}
func SetWellLogs(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var wlt dto.Well_logs
    setDefaults(&wlt)

    if err := c.BodyParser(&wlt); err != nil{
        return err
    }
    
    wlt.Create_date = formatDateString(wlt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO well_logs_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, uwi, logging_company, scale_ratio, media_type, well_log_class_id, log_title, report_log_run, trip_date, top_depth, top_depth_ouom, base_depth, base_depth_ouom, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27) RETURNING id INTO :28`, &wlt.Ba_long_name, &wlt.Ba_type, &wlt.Area_id, &wlt.Area_type, &wlt.Field_name, &wlt.Well_name, &wlt.Uwi, &wlt.Logging_company, &wlt.Scale_ratio, &wlt.Media_type, &wlt.Well_log_class_id, &wlt.Log_title, &wlt.Report_log_run, &wlt.Trip_date, &wlt.Top_depth, &wlt.Top_depth_ouom, &wlt.Base_depth, &wlt.Base_depth_ouom, &wlt.Ba_long_name_2, &wlt.Ba_type_2, &wlt.Data_store_name, &wlt.Data_store_type, &wlt.Location_id, &wlt.Remark, &wlt.Source, &wlt.Qc_status, &wlt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(WELL_LOGS_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetWellLogsById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM well_logs_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_logs    
    
        for rows.Next() {
    
            var wlt dto.Well_logs
            if err := rows.Scan(&wlt.Id, &wlt.Ba_long_name, &wlt.Ba_type, &wlt.Area_id, &wlt.Area_type, &wlt.Field_name, &wlt.Well_name, &wlt.Uwi, &wlt.Logging_company, &wlt.Scale_ratio, &wlt.Media_type, &wlt.Well_log_class_id, &wlt.Log_title, &wlt.Report_log_run, &wlt.Trip_date, &wlt.Top_depth, &wlt.Top_depth_ouom, &wlt.Base_depth, &wlt.Base_depth_ouom, &wlt.Ba_long_name_2, &wlt.Ba_type_2, &wlt.Data_store_name, &wlt.Data_store_type, &wlt.Location_id, &wlt.Remark, &wlt.Source, &wlt.Qc_status, &wlt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wlt)
        }
    
    return c.JSON(results)
}
func PutWellLogs(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wlt dto.Well_logs
    setDefaults(&wlt)

    if err := c.BodyParser(&wlt); err != nil{
        return err
    }
    
    wlt.Create_date = formatDateString(wlt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM well_logs_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE well_logs_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, uwi = :7, logging_company = :8, scale_ratio = :9, media_type = :10, well_log_class_id = :11, log_title = :12, report_log_run = :13, trip_date = :14, top_depth = :15, top_depth_ouom = :16, base_depth = :17, base_depth_ouom = :18, ba_long_name_2 = :19, ba_type_2 = :20, data_store_name = :21, data_store_type = :22, location_id = :23, remark = :24, source = :25, qc_status = :26, checked_by_ba_id = :27 WHERE id = :28`, &wlt.Ba_long_name, &wlt.Ba_type, &wlt.Area_id, &wlt.Area_type, &wlt.Field_name, &wlt.Well_name, &wlt.Uwi, &wlt.Logging_company, &wlt.Scale_ratio, &wlt.Media_type, &wlt.Well_log_class_id, &wlt.Log_title, &wlt.Report_log_run, &wlt.Trip_date, &wlt.Top_depth, &wlt.Top_depth_ouom, &wlt.Base_depth, &wlt.Base_depth_ouom, &wlt.Ba_long_name_2, &wlt.Ba_type_2, &wlt.Data_store_name, &wlt.Data_store_type, &wlt.Location_id, &wlt.Remark, &wlt.Source, &wlt.Qc_status, &wlt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(WELL_LOGS_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteWellLogs(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT well_logs_id FROM well_logs_workspace WHERE well_logs_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM well_logs_workspace WHERE well_logs_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM well_logs_table WHERE id = :1`, id)
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
func PatchWellLogs(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wlt dto.Well_logs
    setDefaults(&wlt)

    if err := c.BodyParser(&wlt); err != nil{
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
    err = db.QueryRow(`SELECT well_logs_id FROM well_logs_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if wlt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET ba_long_name = :1 WHERE id = :2`, wlt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET ba_type = :1 WHERE id = :2`, wlt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Area_id != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET area_id = :1 WHERE id = :2`, wlt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Area_type != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET area_type = :1 WHERE id = :2`, wlt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Field_name != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET field_name = :1 WHERE id = :2`, wlt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Well_name != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET well_name = :1 WHERE id = :2`, wlt.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Uwi != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET uwi = :1 WHERE id = :2`, wlt.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Logging_company != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET logging_company = :1 WHERE id = :2`, wlt.Logging_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Scale_ratio != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET scale_ratio = :1 WHERE id = :2`, wlt.Scale_ratio, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Media_type != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET media_type = :1 WHERE id = :2`, wlt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Well_log_class_id != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET well_log_class_id = :1 WHERE id = :2`, wlt.Well_log_class_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Log_title != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET log_title = :1 WHERE id = :2`, wlt.Log_title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Report_log_run != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET report_log_run = :1 WHERE id = :2`, wlt.Report_log_run, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Trip_date != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET trip_date = :1 WHERE id = :2`, wlt.Trip_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Top_depth != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET top_depth = :1 WHERE id = :2`, wlt.Top_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Top_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET top_depth_ouom = :1 WHERE id = :2`, wlt.Top_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Base_depth != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET base_depth = :1 WHERE id = :2`, wlt.Base_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Base_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET base_depth_ouom = :1 WHERE id = :2`, wlt.Base_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET ba_long_name_2 = :1 WHERE id = :2`, wlt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET ba_type_2 = :1 WHERE id = :2`, wlt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET data_store_name = :1 WHERE id = :2`, wlt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET data_store_type = :1 WHERE id = :2`, wlt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Location_id != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET location_id = :1 WHERE id = :2`, wlt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Remark != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET remark = :1 WHERE id = :2`, wlt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Source != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET source = :1 WHERE id = :2`, wlt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET qc_status = :1 WHERE id = :2`, wlt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wlt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE well_logs_table SET checked_by_ba_id = :1 WHERE id = :2`, wlt.Checked_by_ba_id, id)
        
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

    