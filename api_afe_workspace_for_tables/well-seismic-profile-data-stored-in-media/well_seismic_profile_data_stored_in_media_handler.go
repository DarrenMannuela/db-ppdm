package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-wellseismicprofiledatastoredinmedia/dto"
    "github.com/DarrenMannuela/svc-datatype-wellseismicprofiledatastoredinmedia/utils"

    "github.com/gofiber/fiber/v2"
)
func GetWellSeismicProfileDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM well_seismic_profile_data_stored_in_media_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_seismic_profile_data_stored_in_media    
    
        for rows.Next() {
    
            var wspdsimt dto.Well_seismic_profile_data_stored_in_media
            if err := rows.Scan(&wspdsimt.Id, &wspdsimt.Ba_long_name, &wspdsimt.Ba_type, &wspdsimt.Area_id, &wspdsimt.Area_type, &wspdsimt.Field_name, &wspdsimt.Well_name, &wspdsimt.Alias_long_name, &wspdsimt.Uwi, &wspdsimt.Trip_date, &wspdsimt.Survey_company_ba_id, &wspdsimt.Top_depth, &wspdsimt.Top_depth_ouom, &wspdsimt.Base_depth, &wspdsimt.Base_depth_ouom, &wspdsimt.Checkshot_survey_type, &wspdsimt.Vsp_type, &wspdsimt.Digital_format, &wspdsimt.Media_type, &wspdsimt.Ba_long_name_2, &wspdsimt.Ba_type_2, &wspdsimt.Data_store_name, &wspdsimt.Data_store_type, &wspdsimt.Location_id, &wspdsimt.Remark, &wspdsimt.Source, &wspdsimt.Qc_status, &wspdsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wspdsimt)
        }
    
    return c.JSON(results)
}
func SetWellSeismicProfileDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var wspdsimt dto.Well_seismic_profile_data_stored_in_media
    setDefaults(&wspdsimt)

    if err := c.BodyParser(&wspdsimt); err != nil{
        return err
    }
    
    wspdsimt.Create_date = formatDateString(wspdsimt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO well_seismic_profile_data_stored_in_media_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, alias_long_name, uwi, trip_date, survey_company_ba_id, top_depth, top_depth_ouom, base_depth, base_depth_ouom, checkshot_survey_type, vsp_type, digital_format, media_type, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27) RETURNING id INTO :28`, &wspdsimt.Ba_long_name, &wspdsimt.Ba_type, &wspdsimt.Area_id, &wspdsimt.Area_type, &wspdsimt.Field_name, &wspdsimt.Well_name, &wspdsimt.Alias_long_name, &wspdsimt.Uwi, &wspdsimt.Trip_date, &wspdsimt.Survey_company_ba_id, &wspdsimt.Top_depth, &wspdsimt.Top_depth_ouom, &wspdsimt.Base_depth, &wspdsimt.Base_depth_ouom, &wspdsimt.Checkshot_survey_type, &wspdsimt.Vsp_type, &wspdsimt.Digital_format, &wspdsimt.Media_type, &wspdsimt.Ba_long_name_2, &wspdsimt.Ba_type_2, &wspdsimt.Data_store_name, &wspdsimt.Data_store_type, &wspdsimt.Location_id, &wspdsimt.Remark, &wspdsimt.Source, &wspdsimt.Qc_status, &wspdsimt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(WELL_SEISMIC_PROFILE_DATA_STORED_IN_MEDIA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetWellSeismicProfileDataStoredInMediaById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM well_seismic_profile_data_stored_in_media_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_seismic_profile_data_stored_in_media    
    
        for rows.Next() {
    
            var wspdsimt dto.Well_seismic_profile_data_stored_in_media
            if err := rows.Scan(&wspdsimt.Id, &wspdsimt.Ba_long_name, &wspdsimt.Ba_type, &wspdsimt.Area_id, &wspdsimt.Area_type, &wspdsimt.Field_name, &wspdsimt.Well_name, &wspdsimt.Alias_long_name, &wspdsimt.Uwi, &wspdsimt.Trip_date, &wspdsimt.Survey_company_ba_id, &wspdsimt.Top_depth, &wspdsimt.Top_depth_ouom, &wspdsimt.Base_depth, &wspdsimt.Base_depth_ouom, &wspdsimt.Checkshot_survey_type, &wspdsimt.Vsp_type, &wspdsimt.Digital_format, &wspdsimt.Media_type, &wspdsimt.Ba_long_name_2, &wspdsimt.Ba_type_2, &wspdsimt.Data_store_name, &wspdsimt.Data_store_type, &wspdsimt.Location_id, &wspdsimt.Remark, &wspdsimt.Source, &wspdsimt.Qc_status, &wspdsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wspdsimt)
        }
    
    return c.JSON(results)
}
func PutWellSeismicProfileDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wspdsimt dto.Well_seismic_profile_data_stored_in_media
    setDefaults(&wspdsimt)

    if err := c.BodyParser(&wspdsimt); err != nil{
        return err
    }
    
    wspdsimt.Create_date = formatDateString(wspdsimt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM well_seismic_profile_data_stored_in_media_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE well_seismic_profile_data_stored_in_media_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, alias_long_name = :7, uwi = :8, trip_date = :9, survey_company_ba_id = :10, top_depth = :11, top_depth_ouom = :12, base_depth = :13, base_depth_ouom = :14, checkshot_survey_type = :15, vsp_type = :16, digital_format = :17, media_type = :18, ba_long_name_2 = :19, ba_type_2 = :20, data_store_name = :21, data_store_type = :22, location_id = :23, remark = :24, source = :25, qc_status = :26, checked_by_ba_id = :27 WHERE id = :28`, &wspdsimt.Ba_long_name, &wspdsimt.Ba_type, &wspdsimt.Area_id, &wspdsimt.Area_type, &wspdsimt.Field_name, &wspdsimt.Well_name, &wspdsimt.Alias_long_name, &wspdsimt.Uwi, &wspdsimt.Trip_date, &wspdsimt.Survey_company_ba_id, &wspdsimt.Top_depth, &wspdsimt.Top_depth_ouom, &wspdsimt.Base_depth, &wspdsimt.Base_depth_ouom, &wspdsimt.Checkshot_survey_type, &wspdsimt.Vsp_type, &wspdsimt.Digital_format, &wspdsimt.Media_type, &wspdsimt.Ba_long_name_2, &wspdsimt.Ba_type_2, &wspdsimt.Data_store_name, &wspdsimt.Data_store_type, &wspdsimt.Location_id, &wspdsimt.Remark, &wspdsimt.Source, &wspdsimt.Qc_status, &wspdsimt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(WELL_SEISMIC_PROFILE_DATA_STORED_IN_MEDIA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteWellSeismicProfileDataStoredInMedia(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT well_seismic_profile_data_stored_in_media_id FROM well_seismic_profile_data_stored_in_media_workspace WHERE well_seismic_profile_data_stored_in_media_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM well_seismic_profile_data_stored_in_media_workspace WHERE well_seismic_profile_data_stored_in_media_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM well_seismic_profile_data_stored_in_media_table WHERE id = :1`, id)
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
func PatchWellSeismicProfileDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wspdsimt dto.Well_seismic_profile_data_stored_in_media
    setDefaults(&wspdsimt)

    if err := c.BodyParser(&wspdsimt); err != nil{
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
    err = db.QueryRow(`SELECT well_seismic_profile_data_stored_in_media_id FROM well_seismic_profile_data_stored_in_media_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if wspdsimt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, wspdsimt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, wspdsimt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, wspdsimt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, wspdsimt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Field_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_name = :1 WHERE id = :2`, wspdsimt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Well_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET well_name = :1 WHERE id = :2`, wspdsimt.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Alias_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET alias_long_name = :1 WHERE id = :2`, wspdsimt.Alias_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Uwi != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET uwi = :1 WHERE id = :2`, wspdsimt.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Trip_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET trip_date = :1 WHERE id = :2`, wspdsimt.Trip_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Survey_company_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET survey_company_ba_id = :1 WHERE id = :2`, wspdsimt.Survey_company_ba_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Top_depth != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET top_depth = :1 WHERE id = :2`, wspdsimt.Top_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Top_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET top_depth_ouom = :1 WHERE id = :2`, wspdsimt.Top_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Base_depth != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET base_depth = :1 WHERE id = :2`, wspdsimt.Base_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Base_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET base_depth_ouom = :1 WHERE id = :2`, wspdsimt.Base_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Checkshot_survey_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checkshot_survey_type = :1 WHERE id = :2`, wspdsimt.Checkshot_survey_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Vsp_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET vsp_type = :1 WHERE id = :2`, wspdsimt.Vsp_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, wspdsimt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, wspdsimt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, wspdsimt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, wspdsimt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, wspdsimt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_type = :1 WHERE id = :2`, wspdsimt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Location_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET location_id = :1 WHERE id = :2`, wspdsimt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, wspdsimt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, wspdsimt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, wspdsimt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wspdsimt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, wspdsimt.Checked_by_ba_id, id)
        
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

    