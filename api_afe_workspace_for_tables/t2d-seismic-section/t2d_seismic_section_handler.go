package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t2dseismicsection/dto"
    "github.com/DarrenMannuela/svc-datatype-t2dseismicsection/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT2DSeismicSection(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_section_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_section    
    
        for rows.Next() {
    
            var tsst dto.T2D_seismic_section
            if err := rows.Scan(&tsst.Id, &tsst.Ba_long_name, &tsst.Ba_type, &tsst.Area_id, &tsst.Area_type, &tsst.Acqtn_survey_name, &tsst.Processing_company, &tsst.Start_date, &tsst.Line_name, &tsst.First_seis_point_id, &tsst.Last_seis_point_id, &tsst.Create_date, &tsst.Proc_set_type, &tsst.Media_type, &tsst.Vertical_scale, &tsst.Vertical_scale_ouom, &tsst.Horizontal_scale, &tsst.Horizontal_scale_ouom, &tsst.Polarity, &tsst.Ba_long_name_2, &tsst.Ba_type_2, &tsst.Data_store_name, &tsst.Data_store_type, &tsst.Location_id, &tsst.Sw_application_id, &tsst.Application_version, &tsst.Remark, &tsst.Source, &tsst.Qc_status, &tsst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsst)
        }
    
    return c.JSON(results)
}
func SetT2DSeismicSection(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tsst dto.T2D_seismic_section
    setDefaults(&tsst)

    if err := c.BodyParser(&tsst); err != nil{
        return err
    }
    
    tsst.Create_date = formatDateString(tsst.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t2d_seismic_section_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, processing_company, start_date, line_name, first_seis_point_id, last_seis_point_id, create_date, proc_set_type, media_type, vertical_scale, vertical_scale_ouom, horizontal_scale, horizontal_scale_ouom, polarity, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, sw_application_id, application_version, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29) RETURNING id INTO :30`, &tsst.Ba_long_name, &tsst.Ba_type, &tsst.Area_id, &tsst.Area_type, &tsst.Acqtn_survey_name, &tsst.Processing_company, &tsst.Start_date, &tsst.Line_name, &tsst.First_seis_point_id, &tsst.Last_seis_point_id, &tsst.Create_date, &tsst.Proc_set_type, &tsst.Media_type, &tsst.Vertical_scale, &tsst.Vertical_scale_ouom, &tsst.Horizontal_scale, &tsst.Horizontal_scale_ouom, &tsst.Polarity, &tsst.Ba_long_name_2, &tsst.Ba_type_2, &tsst.Data_store_name, &tsst.Data_store_type, &tsst.Location_id, &tsst.Sw_application_id, &tsst.Application_version, &tsst.Remark, &tsst.Source, &tsst.Qc_status, &tsst.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T2D_SEISMIC_SECTION_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT2DSeismicSectionById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_section_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_section    
    
        for rows.Next() {
    
            var tsst dto.T2D_seismic_section
            if err := rows.Scan(&tsst.Id, &tsst.Ba_long_name, &tsst.Ba_type, &tsst.Area_id, &tsst.Area_type, &tsst.Acqtn_survey_name, &tsst.Processing_company, &tsst.Start_date, &tsst.Line_name, &tsst.First_seis_point_id, &tsst.Last_seis_point_id, &tsst.Create_date, &tsst.Proc_set_type, &tsst.Media_type, &tsst.Vertical_scale, &tsst.Vertical_scale_ouom, &tsst.Horizontal_scale, &tsst.Horizontal_scale_ouom, &tsst.Polarity, &tsst.Ba_long_name_2, &tsst.Ba_type_2, &tsst.Data_store_name, &tsst.Data_store_type, &tsst.Location_id, &tsst.Sw_application_id, &tsst.Application_version, &tsst.Remark, &tsst.Source, &tsst.Qc_status, &tsst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsst)
        }
    
    return c.JSON(results)
}
func PutT2DSeismicSection(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsst dto.T2D_seismic_section
    setDefaults(&tsst)

    if err := c.BodyParser(&tsst); err != nil{
        return err
    }
    
    tsst.Create_date = formatDateString(tsst.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t2d_seismic_section_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, processing_company = :6, start_date = :7, line_name = :8, first_seis_point_id = :9, last_seis_point_id = :10, create_date = :11, proc_set_type = :12, media_type = :13, vertical_scale = :14, vertical_scale_ouom = :15, horizontal_scale = :16, horizontal_scale_ouom = :17, polarity = :18, ba_long_name_2 = :19, ba_type_2 = :20, data_store_name = :21, data_store_type = :22, location_id = :23, sw_application_id = :24, application_version = :25, remark = :26, source = :27, qc_status = :28, checked_by_ba_id = :29 WHERE id = :30`, &tsst.Ba_long_name, &tsst.Ba_type, &tsst.Area_id, &tsst.Area_type, &tsst.Acqtn_survey_name, &tsst.Processing_company, &tsst.Start_date, &tsst.Line_name, &tsst.First_seis_point_id, &tsst.Last_seis_point_id, &tsst.Create_date, &tsst.Proc_set_type, &tsst.Media_type, &tsst.Vertical_scale, &tsst.Vertical_scale_ouom, &tsst.Horizontal_scale, &tsst.Horizontal_scale_ouom, &tsst.Polarity, &tsst.Ba_long_name_2, &tsst.Ba_type_2, &tsst.Data_store_name, &tsst.Data_store_type, &tsst.Location_id, &tsst.Sw_application_id, &tsst.Application_version, &tsst.Remark, &tsst.Source, &tsst.Qc_status, &tsst.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T2D_SEISMIC_SECTION_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT2DSeismicSection(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t2d_seismic_section_id FROM t2d_seismic_section_workspace WHERE t2d_seismic_section_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t2d_seismic_section_workspace WHERE t2d_seismic_section_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t2d_seismic_section_table WHERE id = :1`, id)
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
func PatchT2DSeismicSection(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsst dto.T2D_seismic_section
    setDefaults(&tsst)

    if err := c.BodyParser(&tsst); err != nil{
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
    err = db.QueryRow(`SELECT t2d_seismic_section_id FROM t2d_seismic_section_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tsst.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET ba_long_name = :1 WHERE id = :2`, tsst.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Ba_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET ba_type = :1 WHERE id = :2`, tsst.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Area_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET area_id = :1 WHERE id = :2`, tsst.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Area_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET area_type = :1 WHERE id = :2`, tsst.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET acqtn_survey_name = :1 WHERE id = :2`, tsst.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Processing_company != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET processing_company = :1 WHERE id = :2`, tsst.Processing_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Start_date != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET start_date = :1 WHERE id = :2`, tsst.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Line_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET line_name = :1 WHERE id = :2`, tsst.Line_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.First_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET first_seis_point_id = :1 WHERE id = :2`, tsst.First_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Last_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET last_seis_point_id = :1 WHERE id = :2`, tsst.Last_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Create_date != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET create_date = :1 WHERE id = :2`, tsst.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET proc_set_type = :1 WHERE id = :2`, tsst.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Media_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET media_type = :1 WHERE id = :2`, tsst.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Vertical_scale != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET vertical_scale = :1 WHERE id = :2`, tsst.Vertical_scale, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Vertical_scale_ouom != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET vertical_scale_ouom = :1 WHERE id = :2`, tsst.Vertical_scale_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Horizontal_scale != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET horizontal_scale = :1 WHERE id = :2`, tsst.Horizontal_scale, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Horizontal_scale_ouom != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET horizontal_scale_ouom = :1 WHERE id = :2`, tsst.Horizontal_scale_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Polarity != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET polarity = :1 WHERE id = :2`, tsst.Polarity, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET ba_long_name_2 = :1 WHERE id = :2`, tsst.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET ba_type_2 = :1 WHERE id = :2`, tsst.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET data_store_name = :1 WHERE id = :2`, tsst.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET data_store_type = :1 WHERE id = :2`, tsst.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Location_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET location_id = :1 WHERE id = :2`, tsst.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Sw_application_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET sw_application_id = :1 WHERE id = :2`, tsst.Sw_application_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Application_version != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET application_version = :1 WHERE id = :2`, tsst.Application_version, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Remark != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET remark = :1 WHERE id = :2`, tsst.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Source != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET source = :1 WHERE id = :2`, tsst.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Qc_status != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET qc_status = :1 WHERE id = :2`, tsst.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_section_table SET checked_by_ba_id = :1 WHERE id = :2`, tsst.Checked_by_ba_id, id)
        
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

    