package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-digital2d3dseismicreport/dto"
    "github.com/DarrenMannuela/svc-datatype-digital2d3dseismicreport/utils"

    "github.com/gofiber/fiber/v2"
)
func GetDigital2D3DSeismicReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM digital_2d_3d_seismic_report_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_2d_3d_seismic_report    
    
        for rows.Next() {
    
            var d23srt dto.Digital_2d_3d_seismic_report
            if err := rows.Scan(&d23srt.Id, &d23srt.Ba_long_name, &d23srt.Ba_type, &d23srt.Area_id, &d23srt.Area_type, &d23srt.Acqtn_survey_name, &d23srt.Start_date, &d23srt.Seis_dimension, &d23srt.Line_name, &d23srt.Title, &d23srt.Creator_name, &d23srt.Create_date, &d23srt.Item_category, &d23srt.Digital_format, &d23srt.Media_type, &d23srt.Ba_long_name_2, &d23srt.Ba_type_2, &d23srt.Data_store_name, &d23srt.Original_file_name, &d23srt.Password, &d23srt.Digital_size, &d23srt.Digital_size_uom, &d23srt.Remark, &d23srt.Source, &d23srt.Qc_status, &d23srt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, d23srt)
        }
    
    return c.JSON(results)
}
func SetDigital2D3DSeismicReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var d23srt dto.Digital_2d_3d_seismic_report
    setDefaults(&d23srt)

    if err := c.BodyParser(&d23srt); err != nil{
        return err
    }
    
    d23srt.Create_date = formatDateString(d23srt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO digital_2d_3d_seismic_report_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, start_date, seis_dimension, line_name, title, creator_name, create_date, item_category, digital_format, media_type, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25) RETURNING id INTO :26`, &d23srt.Ba_long_name, &d23srt.Ba_type, &d23srt.Area_id, &d23srt.Area_type, &d23srt.Acqtn_survey_name, &d23srt.Start_date, &d23srt.Seis_dimension, &d23srt.Line_name, &d23srt.Title, &d23srt.Creator_name, &d23srt.Create_date, &d23srt.Item_category, &d23srt.Digital_format, &d23srt.Media_type, &d23srt.Ba_long_name_2, &d23srt.Ba_type_2, &d23srt.Data_store_name, &d23srt.Original_file_name, &d23srt.Password, &d23srt.Digital_size, &d23srt.Digital_size_uom, &d23srt.Remark, &d23srt.Source, &d23srt.Qc_status, &d23srt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(DIGITAL_2D_3D_SEISMIC_REPORT_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetDigital2D3DSeismicReportById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM digital_2d_3d_seismic_report_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_2d_3d_seismic_report    
    
        for rows.Next() {
    
            var d23srt dto.Digital_2d_3d_seismic_report
            if err := rows.Scan(&d23srt.Id, &d23srt.Ba_long_name, &d23srt.Ba_type, &d23srt.Area_id, &d23srt.Area_type, &d23srt.Acqtn_survey_name, &d23srt.Start_date, &d23srt.Seis_dimension, &d23srt.Line_name, &d23srt.Title, &d23srt.Creator_name, &d23srt.Create_date, &d23srt.Item_category, &d23srt.Digital_format, &d23srt.Media_type, &d23srt.Ba_long_name_2, &d23srt.Ba_type_2, &d23srt.Data_store_name, &d23srt.Original_file_name, &d23srt.Password, &d23srt.Digital_size, &d23srt.Digital_size_uom, &d23srt.Remark, &d23srt.Source, &d23srt.Qc_status, &d23srt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, d23srt)
        }
    
    return c.JSON(results)
}
func PutDigital2D3DSeismicReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var d23srt dto.Digital_2d_3d_seismic_report
    setDefaults(&d23srt)

    if err := c.BodyParser(&d23srt); err != nil{
        return err
    }
    
    d23srt.Create_date = formatDateString(d23srt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM digital_2d_3d_seismic_report_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE digital_2d_3d_seismic_report_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, start_date = :6, seis_dimension = :7, line_name = :8, title = :9, creator_name = :10, create_date = :11, item_category = :12, digital_format = :13, media_type = :14, ba_long_name_2 = :15, ba_type_2 = :16, data_store_name = :17, original_file_name = :18, password = :19, digital_size = :20, digital_size_uom = :21, remark = :22, source = :23, qc_status = :24, checked_by_ba_id = :25 WHERE id = :26`, &d23srt.Ba_long_name, &d23srt.Ba_type, &d23srt.Area_id, &d23srt.Area_type, &d23srt.Acqtn_survey_name, &d23srt.Start_date, &d23srt.Seis_dimension, &d23srt.Line_name, &d23srt.Title, &d23srt.Creator_name, &d23srt.Create_date, &d23srt.Item_category, &d23srt.Digital_format, &d23srt.Media_type, &d23srt.Ba_long_name_2, &d23srt.Ba_type_2, &d23srt.Data_store_name, &d23srt.Original_file_name, &d23srt.Password, &d23srt.Digital_size, &d23srt.Digital_size_uom, &d23srt.Remark, &d23srt.Source, &d23srt.Qc_status, &d23srt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(DIGITAL_2D_3D_SEISMIC_REPORT_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteDigital2D3DSeismicReport(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT digital_2d_3d_seismic_report_id FROM digital_2d_3d_seismic_report_workspace WHERE digital_2d_3d_seismic_report_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM digital_2d_3d_seismic_report_workspace WHERE digital_2d_3d_seismic_report_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM digital_2d_3d_seismic_report_table WHERE id = :1`, id)
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
func PatchDigital2D3DSeismicReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var d23srt dto.Digital_2d_3d_seismic_report
    setDefaults(&d23srt)

    if err := c.BodyParser(&d23srt); err != nil{
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
    err = db.QueryRow(`SELECT digital_2d_3d_seismic_report_id FROM digital_2d_3d_seismic_report_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if d23srt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, d23srt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, d23srt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, d23srt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, d23srt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET acqtn_survey_name = :1 WHERE id = :2`, d23srt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Start_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET start_date = :1 WHERE id = :2`, d23srt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Seis_dimension != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET seis_dimension = :1 WHERE id = :2`, d23srt.Seis_dimension, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Line_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET line_name = :1 WHERE id = :2`, d23srt.Line_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Title != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET title = :1 WHERE id = :2`, d23srt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET creator_name = :1 WHERE id = :2`, d23srt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, d23srt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Item_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_category = :1 WHERE id = :2`, d23srt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, d23srt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, d23srt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, d23srt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, d23srt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, d23srt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, d23srt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, d23srt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, d23srt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, d23srt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, d23srt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, d23srt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, d23srt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d23srt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, d23srt.Checked_by_ba_id, id)
        
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

    