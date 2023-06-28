package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-digital2dseismicsection/dto"
    "github.com/DarrenMannuela/svc-datatype-digital2dseismicsection/utils"

    "github.com/gofiber/fiber/v2"
)
func GetDigital2DSeismicSection(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM digital_2d_seismic_section_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_2d_seismic_section    
    
        for rows.Next() {
    
            var d2sst dto.Digital_2d_seismic_section
            if err := rows.Scan(&d2sst.Id, &d2sst.Ba_long_name, &d2sst.Ba_type, &d2sst.Area_id, &d2sst.Area_type, &d2sst.Acqtn_survey_name, &d2sst.Processing_company, &d2sst.Start_date, &d2sst.Line_name, &d2sst.First_seis_point_id, &d2sst.Last_seis_point_id, &d2sst.Create_date, &d2sst.Proc_set_type, &d2sst.Digital_format, &d2sst.Media_type, &d2sst.Vertical_scale, &d2sst.Vertical_scale_ouom, &d2sst.Horizontal_scale, &d2sst.Horizontal_scale_ouom, &d2sst.Polarity, &d2sst.Ba_long_name_2, &d2sst.Ba_type_2, &d2sst.Data_store_name, &d2sst.Original_file_name, &d2sst.Password, &d2sst.Sw_application_id, &d2sst.Application_version, &d2sst.Digital_size, &d2sst.Digital_size_uom, &d2sst.Remark, &d2sst.Source, &d2sst.Qc_status, &d2sst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, d2sst)
        }
    
    return c.JSON(results)
}
func SetDigital2DSeismicSection(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var d2sst dto.Digital_2d_seismic_section
    setDefaults(&d2sst)

    if err := c.BodyParser(&d2sst); err != nil{
        return err
    }
    
    d2sst.Create_date = formatDateString(d2sst.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO digital_2d_seismic_section_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, processing_company, start_date, line_name, first_seis_point_id, last_seis_point_id, create_date, proc_set_type, digital_format, media_type, vertical_scale, vertical_scale_ouom, horizontal_scale, horizontal_scale_ouom, polarity, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, sw_application_id, application_version, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32) RETURNING id INTO :33`, &d2sst.Ba_long_name, &d2sst.Ba_type, &d2sst.Area_id, &d2sst.Area_type, &d2sst.Acqtn_survey_name, &d2sst.Processing_company, &d2sst.Start_date, &d2sst.Line_name, &d2sst.First_seis_point_id, &d2sst.Last_seis_point_id, &d2sst.Create_date, &d2sst.Proc_set_type, &d2sst.Digital_format, &d2sst.Media_type, &d2sst.Vertical_scale, &d2sst.Vertical_scale_ouom, &d2sst.Horizontal_scale, &d2sst.Horizontal_scale_ouom, &d2sst.Polarity, &d2sst.Ba_long_name_2, &d2sst.Ba_type_2, &d2sst.Data_store_name, &d2sst.Original_file_name, &d2sst.Password, &d2sst.Sw_application_id, &d2sst.Application_version, &d2sst.Digital_size, &d2sst.Digital_size_uom, &d2sst.Remark, &d2sst.Source, &d2sst.Qc_status, &d2sst.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(DIGITAL_2D_SEISMIC_SECTION_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetDigital2DSeismicSectionById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM digital_2d_seismic_section_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_2d_seismic_section    
    
        for rows.Next() {
    
            var d2sst dto.Digital_2d_seismic_section
            if err := rows.Scan(&d2sst.Id, &d2sst.Ba_long_name, &d2sst.Ba_type, &d2sst.Area_id, &d2sst.Area_type, &d2sst.Acqtn_survey_name, &d2sst.Processing_company, &d2sst.Start_date, &d2sst.Line_name, &d2sst.First_seis_point_id, &d2sst.Last_seis_point_id, &d2sst.Create_date, &d2sst.Proc_set_type, &d2sst.Digital_format, &d2sst.Media_type, &d2sst.Vertical_scale, &d2sst.Vertical_scale_ouom, &d2sst.Horizontal_scale, &d2sst.Horizontal_scale_ouom, &d2sst.Polarity, &d2sst.Ba_long_name_2, &d2sst.Ba_type_2, &d2sst.Data_store_name, &d2sst.Original_file_name, &d2sst.Password, &d2sst.Sw_application_id, &d2sst.Application_version, &d2sst.Digital_size, &d2sst.Digital_size_uom, &d2sst.Remark, &d2sst.Source, &d2sst.Qc_status, &d2sst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, d2sst)
        }
    
    return c.JSON(results)
}
func PutDigital2DSeismicSection(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var d2sst dto.Digital_2d_seismic_section
    setDefaults(&d2sst)

    if err := c.BodyParser(&d2sst); err != nil{
        return err
    }
    
    d2sst.Create_date = formatDateString(d2sst.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM digital_2d_seismic_section_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, processing_company = :6, start_date = :7, line_name = :8, first_seis_point_id = :9, last_seis_point_id = :10, create_date = :11, proc_set_type = :12, digital_format = :13, media_type = :14, vertical_scale = :15, vertical_scale_ouom = :16, horizontal_scale = :17, horizontal_scale_ouom = :18, polarity = :19, ba_long_name_2 = :20, ba_type_2 = :21, data_store_name = :22, original_file_name = :23, password = :24, sw_application_id = :25, application_version = :26, digital_size = :27, digital_size_uom = :28, remark = :29, source = :30, qc_status = :31, checked_by_ba_id = :32 WHERE id = :33`, &d2sst.Ba_long_name, &d2sst.Ba_type, &d2sst.Area_id, &d2sst.Area_type, &d2sst.Acqtn_survey_name, &d2sst.Processing_company, &d2sst.Start_date, &d2sst.Line_name, &d2sst.First_seis_point_id, &d2sst.Last_seis_point_id, &d2sst.Create_date, &d2sst.Proc_set_type, &d2sst.Digital_format, &d2sst.Media_type, &d2sst.Vertical_scale, &d2sst.Vertical_scale_ouom, &d2sst.Horizontal_scale, &d2sst.Horizontal_scale_ouom, &d2sst.Polarity, &d2sst.Ba_long_name_2, &d2sst.Ba_type_2, &d2sst.Data_store_name, &d2sst.Original_file_name, &d2sst.Password, &d2sst.Sw_application_id, &d2sst.Application_version, &d2sst.Digital_size, &d2sst.Digital_size_uom, &d2sst.Remark, &d2sst.Source, &d2sst.Qc_status, &d2sst.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(DIGITAL_2D_SEISMIC_SECTION_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteDigital2DSeismicSection(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT digital_2d_seismic_section_id FROM digital_2d_seismic_section_workspace WHERE digital_2d_seismic_section_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM digital_2d_seismic_section_workspace WHERE digital_2d_seismic_section_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM digital_2d_seismic_section_table WHERE id = :1`, id)
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
func PatchDigital2DSeismicSection(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var d2sst dto.Digital_2d_seismic_section
    setDefaults(&d2sst)

    if err := c.BodyParser(&d2sst); err != nil{
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
    err = db.QueryRow(`SELECT digital_2d_seismic_section_id FROM digital_2d_seismic_section_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if d2sst.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET ba_long_name = :1 WHERE id = :2`, d2sst.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Ba_type != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET ba_type = :1 WHERE id = :2`, d2sst.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Area_id != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET area_id = :1 WHERE id = :2`, d2sst.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Area_type != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET area_type = :1 WHERE id = :2`, d2sst.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET acqtn_survey_name = :1 WHERE id = :2`, d2sst.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Processing_company != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET processing_company = :1 WHERE id = :2`, d2sst.Processing_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Start_date != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET start_date = :1 WHERE id = :2`, d2sst.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Line_name != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET line_name = :1 WHERE id = :2`, d2sst.Line_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.First_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET first_seis_point_id = :1 WHERE id = :2`, d2sst.First_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Last_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET last_seis_point_id = :1 WHERE id = :2`, d2sst.Last_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Create_date != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET create_date = :1 WHERE id = :2`, d2sst.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET proc_set_type = :1 WHERE id = :2`, d2sst.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Digital_format != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET digital_format = :1 WHERE id = :2`, d2sst.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Media_type != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET media_type = :1 WHERE id = :2`, d2sst.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Vertical_scale != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET vertical_scale = :1 WHERE id = :2`, d2sst.Vertical_scale, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Vertical_scale_ouom != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET vertical_scale_ouom = :1 WHERE id = :2`, d2sst.Vertical_scale_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Horizontal_scale != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET horizontal_scale = :1 WHERE id = :2`, d2sst.Horizontal_scale, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Horizontal_scale_ouom != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET horizontal_scale_ouom = :1 WHERE id = :2`, d2sst.Horizontal_scale_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Polarity != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET polarity = :1 WHERE id = :2`, d2sst.Polarity, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET ba_long_name_2 = :1 WHERE id = :2`, d2sst.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET ba_type_2 = :1 WHERE id = :2`, d2sst.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET data_store_name = :1 WHERE id = :2`, d2sst.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET original_file_name = :1 WHERE id = :2`, d2sst.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Password != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET password = :1 WHERE id = :2`, d2sst.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Sw_application_id != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET sw_application_id = :1 WHERE id = :2`, d2sst.Sw_application_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Application_version != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET application_version = :1 WHERE id = :2`, d2sst.Application_version, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Digital_size != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET digital_size = :1 WHERE id = :2`, d2sst.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET digital_size_uom = :1 WHERE id = :2`, d2sst.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Remark != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET remark = :1 WHERE id = :2`, d2sst.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Source != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET source = :1 WHERE id = :2`, d2sst.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Qc_status != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET qc_status = :1 WHERE id = :2`, d2sst.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if d2sst.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE digital_2d_seismic_section_table SET checked_by_ba_id = :1 WHERE id = :2`, d2sst.Checked_by_ba_id, id)
        
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

    