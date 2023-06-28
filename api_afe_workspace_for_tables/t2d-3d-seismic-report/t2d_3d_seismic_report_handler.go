package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t2d3dseismicreport/dto"
    "github.com/DarrenMannuela/svc-datatype-t2d3dseismicreport/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT2D3DSeismicReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t2d_3d_seismic_report_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_3d_seismic_report    
    
        for rows.Next() {
    
            var t3srt dto.T2D_3d_seismic_report
            if err := rows.Scan(&t3srt.Id, &t3srt.Ba_long_name, &t3srt.Ba_type, &t3srt.Area_id, &t3srt.Area_type, &t3srt.Acqtn_survey_name, &t3srt.Start_date, &t3srt.Seis_dimension, &t3srt.Line_name, &t3srt.Title, &t3srt.Creator_name, &t3srt.Create_date, &t3srt.Item_category, &t3srt.Media_type, &t3srt.Ba_long_name_2, &t3srt.Ba_type_2, &t3srt.Data_store_name, &t3srt.Data_store_type, &t3srt.Location_id, &t3srt.Remark, &t3srt.Source, &t3srt.Qc_status, &t3srt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, t3srt)
        }
    
    return c.JSON(results)
}
func SetT2D3DSeismicReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var t3srt dto.T2D_3d_seismic_report
    setDefaults(&t3srt)

    if err := c.BodyParser(&t3srt); err != nil{
        return err
    }
    
    t3srt.Create_date = formatDateString(t3srt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t2d_3d_seismic_report_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, start_date, seis_dimension, line_name, title, creator_name, create_date, item_category, media_type, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22) RETURNING id INTO :23`, &t3srt.Ba_long_name, &t3srt.Ba_type, &t3srt.Area_id, &t3srt.Area_type, &t3srt.Acqtn_survey_name, &t3srt.Start_date, &t3srt.Seis_dimension, &t3srt.Line_name, &t3srt.Title, &t3srt.Creator_name, &t3srt.Create_date, &t3srt.Item_category, &t3srt.Media_type, &t3srt.Ba_long_name_2, &t3srt.Ba_type_2, &t3srt.Data_store_name, &t3srt.Data_store_type, &t3srt.Location_id, &t3srt.Remark, &t3srt.Source, &t3srt.Qc_status, &t3srt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T2D_3D_SEISMIC_REPORT_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT2D3DSeismicReportById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t2d_3d_seismic_report_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_3d_seismic_report    
    
        for rows.Next() {
    
            var t3srt dto.T2D_3d_seismic_report
            if err := rows.Scan(&t3srt.Id, &t3srt.Ba_long_name, &t3srt.Ba_type, &t3srt.Area_id, &t3srt.Area_type, &t3srt.Acqtn_survey_name, &t3srt.Start_date, &t3srt.Seis_dimension, &t3srt.Line_name, &t3srt.Title, &t3srt.Creator_name, &t3srt.Create_date, &t3srt.Item_category, &t3srt.Media_type, &t3srt.Ba_long_name_2, &t3srt.Ba_type_2, &t3srt.Data_store_name, &t3srt.Data_store_type, &t3srt.Location_id, &t3srt.Remark, &t3srt.Source, &t3srt.Qc_status, &t3srt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, t3srt)
        }
    
    return c.JSON(results)
}
func PutT2D3DSeismicReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var t3srt dto.T2D_3d_seismic_report
    setDefaults(&t3srt)

    if err := c.BodyParser(&t3srt); err != nil{
        return err
    }
    
    t3srt.Create_date = formatDateString(t3srt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t2d_3d_seismic_report_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, start_date = :6, seis_dimension = :7, line_name = :8, title = :9, creator_name = :10, create_date = :11, item_category = :12, media_type = :13, ba_long_name_2 = :14, ba_type_2 = :15, data_store_name = :16, data_store_type = :17, location_id = :18, remark = :19, source = :20, qc_status = :21, checked_by_ba_id = :22 WHERE id = :23`, &t3srt.Ba_long_name, &t3srt.Ba_type, &t3srt.Area_id, &t3srt.Area_type, &t3srt.Acqtn_survey_name, &t3srt.Start_date, &t3srt.Seis_dimension, &t3srt.Line_name, &t3srt.Title, &t3srt.Creator_name, &t3srt.Create_date, &t3srt.Item_category, &t3srt.Media_type, &t3srt.Ba_long_name_2, &t3srt.Ba_type_2, &t3srt.Data_store_name, &t3srt.Data_store_type, &t3srt.Location_id, &t3srt.Remark, &t3srt.Source, &t3srt.Qc_status, &t3srt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T2D_3D_SEISMIC_REPORT_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT2D3DSeismicReport(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t2d_3d_seismic_report_id FROM t2d_3d_seismic_report_workspace WHERE t2d_3d_seismic_report_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t2d_3d_seismic_report_workspace WHERE t2d_3d_seismic_report_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t2d_3d_seismic_report_table WHERE id = :1`, id)
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
func PatchT2D3DSeismicReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var t3srt dto.T2D_3d_seismic_report
    setDefaults(&t3srt)

    if err := c.BodyParser(&t3srt); err != nil{
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
    err = db.QueryRow(`SELECT t2d_3d_seismic_report_id FROM t2d_3d_seismic_report_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if t3srt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET ba_long_name = :1 WHERE id = :2`, t3srt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET ba_type = :1 WHERE id = :2`, t3srt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Area_id != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET area_id = :1 WHERE id = :2`, t3srt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Area_type != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET area_type = :1 WHERE id = :2`, t3srt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET acqtn_survey_name = :1 WHERE id = :2`, t3srt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Start_date != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET start_date = :1 WHERE id = :2`, t3srt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Seis_dimension != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET seis_dimension = :1 WHERE id = :2`, t3srt.Seis_dimension, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Line_name != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET line_name = :1 WHERE id = :2`, t3srt.Line_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Title != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET title = :1 WHERE id = :2`, t3srt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET creator_name = :1 WHERE id = :2`, t3srt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Create_date != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET create_date = :1 WHERE id = :2`, t3srt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Item_category != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET item_category = :1 WHERE id = :2`, t3srt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Media_type != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET media_type = :1 WHERE id = :2`, t3srt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET ba_long_name_2 = :1 WHERE id = :2`, t3srt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET ba_type_2 = :1 WHERE id = :2`, t3srt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET data_store_name = :1 WHERE id = :2`, t3srt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET data_store_type = :1 WHERE id = :2`, t3srt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Location_id != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET location_id = :1 WHERE id = :2`, t3srt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Remark != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET remark = :1 WHERE id = :2`, t3srt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Source != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET source = :1 WHERE id = :2`, t3srt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET qc_status = :1 WHERE id = :2`, t3srt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if t3srt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE t2d_3d_seismic_report_table SET checked_by_ba_id = :1 WHERE id = :2`, t3srt.Checked_by_ba_id, id)
        
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

    