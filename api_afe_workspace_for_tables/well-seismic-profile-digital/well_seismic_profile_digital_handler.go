package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-wellseismicprofiledigital/dto"
    "github.com/DarrenMannuela/svc-datatype-wellseismicprofiledigital/utils"

    "github.com/gofiber/fiber/v2"
)
func GetWellSeismicProfileDigital(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM well_seismic_profile_digital_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_seismic_profile_digital    
    
        for rows.Next() {
    
            var wspdt dto.Well_seismic_profile_digital
            if err := rows.Scan(&wspdt.Id, &wspdt.Ba_long_name, &wspdt.Ba_type, &wspdt.Area_id, &wspdt.Area_type, &wspdt.Field_name, &wspdt.Well_name, &wspdt.Alias_long_name, &wspdt.Uwi, &wspdt.Trip_date, &wspdt.Survey_company_ba_id, &wspdt.Top_depth, &wspdt.Top_depth_ouom, &wspdt.Base_depth, &wspdt.Base_depth_ouom, &wspdt.Checkshot_survey_type, &wspdt.Vsp_type, &wspdt.Digital_format, &wspdt.Media_type, &wspdt.Original_file_name, &wspdt.Password, &wspdt.Digital_size, &wspdt.Digital_size_uom, &wspdt.Ba_long_name_2, &wspdt.Ba_type_2, &wspdt.Data_store_name, &wspdt.Remark, &wspdt.Source, &wspdt.Qc_status, &wspdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wspdt)
        }
    
    return c.JSON(results)
}
func SetWellSeismicProfileDigital(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var wspdt dto.Well_seismic_profile_digital
    setDefaults(&wspdt)

    if err := c.BodyParser(&wspdt); err != nil{
        return err
    }
    
    wspdt.Create_date = formatDateString(wspdt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO well_seismic_profile_digital_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, alias_long_name, uwi, trip_date, survey_company_ba_id, top_depth, top_depth_ouom, base_depth, base_depth_ouom, checkshot_survey_type, vsp_type, digital_format, media_type, original_file_name, password, digital_size, digital_size_uom, ba_long_name_2, ba_type_2, data_store_name, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29) RETURNING id INTO :30`, &wspdt.Ba_long_name, &wspdt.Ba_type, &wspdt.Area_id, &wspdt.Area_type, &wspdt.Field_name, &wspdt.Well_name, &wspdt.Alias_long_name, &wspdt.Uwi, &wspdt.Trip_date, &wspdt.Survey_company_ba_id, &wspdt.Top_depth, &wspdt.Top_depth_ouom, &wspdt.Base_depth, &wspdt.Base_depth_ouom, &wspdt.Checkshot_survey_type, &wspdt.Vsp_type, &wspdt.Digital_format, &wspdt.Media_type, &wspdt.Original_file_name, &wspdt.Password, &wspdt.Digital_size, &wspdt.Digital_size_uom, &wspdt.Ba_long_name_2, &wspdt.Ba_type_2, &wspdt.Data_store_name, &wspdt.Remark, &wspdt.Source, &wspdt.Qc_status, &wspdt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(WELL_SEISMIC_PROFILE_DIGITAL_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetWellSeismicProfileDigitalById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM well_seismic_profile_digital_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_seismic_profile_digital    
    
        for rows.Next() {
    
            var wspdt dto.Well_seismic_profile_digital
            if err := rows.Scan(&wspdt.Id, &wspdt.Ba_long_name, &wspdt.Ba_type, &wspdt.Area_id, &wspdt.Area_type, &wspdt.Field_name, &wspdt.Well_name, &wspdt.Alias_long_name, &wspdt.Uwi, &wspdt.Trip_date, &wspdt.Survey_company_ba_id, &wspdt.Top_depth, &wspdt.Top_depth_ouom, &wspdt.Base_depth, &wspdt.Base_depth_ouom, &wspdt.Checkshot_survey_type, &wspdt.Vsp_type, &wspdt.Digital_format, &wspdt.Media_type, &wspdt.Original_file_name, &wspdt.Password, &wspdt.Digital_size, &wspdt.Digital_size_uom, &wspdt.Ba_long_name_2, &wspdt.Ba_type_2, &wspdt.Data_store_name, &wspdt.Remark, &wspdt.Source, &wspdt.Qc_status, &wspdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wspdt)
        }
    
    return c.JSON(results)
}
func PutWellSeismicProfileDigital(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wspdt dto.Well_seismic_profile_digital
    setDefaults(&wspdt)

    if err := c.BodyParser(&wspdt); err != nil{
        return err
    }
    
    wspdt.Create_date = formatDateString(wspdt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM well_seismic_profile_digital_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE well_seismic_profile_digital_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, alias_long_name = :7, uwi = :8, trip_date = :9, survey_company_ba_id = :10, top_depth = :11, top_depth_ouom = :12, base_depth = :13, base_depth_ouom = :14, checkshot_survey_type = :15, vsp_type = :16, digital_format = :17, media_type = :18, original_file_name = :19, password = :20, digital_size = :21, digital_size_uom = :22, ba_long_name_2 = :23, ba_type_2 = :24, data_store_name = :25, remark = :26, source = :27, qc_status = :28, checked_by_ba_id = :29 WHERE id = :30`, &wspdt.Ba_long_name, &wspdt.Ba_type, &wspdt.Area_id, &wspdt.Area_type, &wspdt.Field_name, &wspdt.Well_name, &wspdt.Alias_long_name, &wspdt.Uwi, &wspdt.Trip_date, &wspdt.Survey_company_ba_id, &wspdt.Top_depth, &wspdt.Top_depth_ouom, &wspdt.Base_depth, &wspdt.Base_depth_ouom, &wspdt.Checkshot_survey_type, &wspdt.Vsp_type, &wspdt.Digital_format, &wspdt.Media_type, &wspdt.Original_file_name, &wspdt.Password, &wspdt.Digital_size, &wspdt.Digital_size_uom, &wspdt.Ba_long_name_2, &wspdt.Ba_type_2, &wspdt.Data_store_name, &wspdt.Remark, &wspdt.Source, &wspdt.Qc_status, &wspdt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(WELL_SEISMIC_PROFILE_DIGITAL_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteWellSeismicProfileDigital(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT well_seismic_profile_digital_id FROM well_seismic_profile_digital_workspace WHERE well_seismic_profile_digital_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM well_seismic_profile_digital_workspace WHERE well_seismic_profile_digital_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM well_seismic_profile_digital_table WHERE id = :1`, id)
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
func PatchWellSeismicProfileDigital(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wspdt dto.Well_seismic_profile_digital
    setDefaults(&wspdt)

    if err := c.BodyParser(&wspdt); err != nil{
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
    err = db.QueryRow(`SELECT well_seismic_profile_digital_id FROM well_seismic_profile_digital_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if wspdt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, wspdt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, wspdt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, wspdt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, wspdt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Field_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_name = :1 WHERE id = :2`, wspdt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Well_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET well_name = :1 WHERE id = :2`, wspdt.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Alias_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET alias_long_name = :1 WHERE id = :2`, wspdt.Alias_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Uwi != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET uwi = :1 WHERE id = :2`, wspdt.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Trip_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET trip_date = :1 WHERE id = :2`, wspdt.Trip_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Survey_company_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET survey_company_ba_id = :1 WHERE id = :2`, wspdt.Survey_company_ba_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Top_depth != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET top_depth = :1 WHERE id = :2`, wspdt.Top_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Top_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET top_depth_ouom = :1 WHERE id = :2`, wspdt.Top_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Base_depth != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET base_depth = :1 WHERE id = :2`, wspdt.Base_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Base_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET base_depth_ouom = :1 WHERE id = :2`, wspdt.Base_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Checkshot_survey_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checkshot_survey_type = :1 WHERE id = :2`, wspdt.Checkshot_survey_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Vsp_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET vsp_type = :1 WHERE id = :2`, wspdt.Vsp_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, wspdt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, wspdt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, wspdt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, wspdt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, wspdt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, wspdt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, wspdt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, wspdt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, wspdt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, wspdt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, wspdt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, wspdt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, wspdt.Checked_by_ba_id, id)
        
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

    