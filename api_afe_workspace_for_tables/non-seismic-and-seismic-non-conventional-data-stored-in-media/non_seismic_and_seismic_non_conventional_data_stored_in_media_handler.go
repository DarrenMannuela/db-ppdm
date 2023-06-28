package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionaldatastoredinmedia/dto"
    "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionaldatastoredinmedia/utils"

    "github.com/gofiber/fiber/v2"
)
func GetNonSeismicAndSeismicNonConventionalDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM non_seismic_and_seismic_non_conventional_data_stored_in_media_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Non_seismic_and_seismic_non_conventional_data_stored_in_media    
    
        for rows.Next() {
    
            var nsasncdsimt dto.Non_seismic_and_seismic_non_conventional_data_stored_in_media
            if err := rows.Scan(&nsasncdsimt.Id, &nsasncdsimt.Ba_long_name, &nsasncdsimt.Ba_type, &nsasncdsimt.Area_id, &nsasncdsimt.Area_type, &nsasncdsimt.Acqtn_survey_name, &nsasncdsimt.Processing_company, &nsasncdsimt.Start_date, &nsasncdsimt.Seis_station_type, &nsasncdsimt.Create_date, &nsasncdsimt.Proc_set_type, &nsasncdsimt.Media_type, &nsasncdsimt.Tape_number, &nsasncdsimt.Digital_format, &nsasncdsimt.Ba_long_name_2, &nsasncdsimt.Ba_type_2, &nsasncdsimt.Data_store_name, &nsasncdsimt.Data_store_type, &nsasncdsimt.Location_id, &nsasncdsimt.Remark, &nsasncdsimt.Source, &nsasncdsimt.Qc_status, &nsasncdsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, nsasncdsimt)
        }
    
    return c.JSON(results)
}
func SetNonSeismicAndSeismicNonConventionalDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var nsasncdsimt dto.Non_seismic_and_seismic_non_conventional_data_stored_in_media
    setDefaults(&nsasncdsimt)

    if err := c.BodyParser(&nsasncdsimt); err != nil{
        return err
    }
    
    nsasncdsimt.Create_date = formatDateString(nsasncdsimt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO non_seismic_and_seismic_non_conventional_data_stored_in_media_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, processing_company, start_date, seis_station_type, create_date, proc_set_type, media_type, tape_number, digital_format, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22) RETURNING id INTO :23`, &nsasncdsimt.Ba_long_name, &nsasncdsimt.Ba_type, &nsasncdsimt.Area_id, &nsasncdsimt.Area_type, &nsasncdsimt.Acqtn_survey_name, &nsasncdsimt.Processing_company, &nsasncdsimt.Start_date, &nsasncdsimt.Seis_station_type, &nsasncdsimt.Create_date, &nsasncdsimt.Proc_set_type, &nsasncdsimt.Media_type, &nsasncdsimt.Tape_number, &nsasncdsimt.Digital_format, &nsasncdsimt.Ba_long_name_2, &nsasncdsimt.Ba_type_2, &nsasncdsimt.Data_store_name, &nsasncdsimt.Data_store_type, &nsasncdsimt.Location_id, &nsasncdsimt.Remark, &nsasncdsimt.Source, &nsasncdsimt.Qc_status, &nsasncdsimt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_STORED_IN_MEDIA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetNonSeismicAndSeismicNonConventionalDataStoredInMediaById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM non_seismic_and_seismic_non_conventional_data_stored_in_media_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Non_seismic_and_seismic_non_conventional_data_stored_in_media    
    
        for rows.Next() {
    
            var nsasncdsimt dto.Non_seismic_and_seismic_non_conventional_data_stored_in_media
            if err := rows.Scan(&nsasncdsimt.Id, &nsasncdsimt.Ba_long_name, &nsasncdsimt.Ba_type, &nsasncdsimt.Area_id, &nsasncdsimt.Area_type, &nsasncdsimt.Acqtn_survey_name, &nsasncdsimt.Processing_company, &nsasncdsimt.Start_date, &nsasncdsimt.Seis_station_type, &nsasncdsimt.Create_date, &nsasncdsimt.Proc_set_type, &nsasncdsimt.Media_type, &nsasncdsimt.Tape_number, &nsasncdsimt.Digital_format, &nsasncdsimt.Ba_long_name_2, &nsasncdsimt.Ba_type_2, &nsasncdsimt.Data_store_name, &nsasncdsimt.Data_store_type, &nsasncdsimt.Location_id, &nsasncdsimt.Remark, &nsasncdsimt.Source, &nsasncdsimt.Qc_status, &nsasncdsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, nsasncdsimt)
        }
    
    return c.JSON(results)
}
func PutNonSeismicAndSeismicNonConventionalDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var nsasncdsimt dto.Non_seismic_and_seismic_non_conventional_data_stored_in_media
    setDefaults(&nsasncdsimt)

    if err := c.BodyParser(&nsasncdsimt); err != nil{
        return err
    }
    
    nsasncdsimt.Create_date = formatDateString(nsasncdsimt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM non_seismic_and_seismic_non_conventional_data_stored_in_media_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_data_stored_in_media_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, processing_company = :6, start_date = :7, seis_station_type = :8, create_date = :9, proc_set_type = :10, media_type = :11, tape_number = :12, digital_format = :13, ba_long_name_2 = :14, ba_type_2 = :15, data_store_name = :16, data_store_type = :17, location_id = :18, remark = :19, source = :20, qc_status = :21, checked_by_ba_id = :22 WHERE id = :23`, &nsasncdsimt.Ba_long_name, &nsasncdsimt.Ba_type, &nsasncdsimt.Area_id, &nsasncdsimt.Area_type, &nsasncdsimt.Acqtn_survey_name, &nsasncdsimt.Processing_company, &nsasncdsimt.Start_date, &nsasncdsimt.Seis_station_type, &nsasncdsimt.Create_date, &nsasncdsimt.Proc_set_type, &nsasncdsimt.Media_type, &nsasncdsimt.Tape_number, &nsasncdsimt.Digital_format, &nsasncdsimt.Ba_long_name_2, &nsasncdsimt.Ba_type_2, &nsasncdsimt.Data_store_name, &nsasncdsimt.Data_store_type, &nsasncdsimt.Location_id, &nsasncdsimt.Remark, &nsasncdsimt.Source, &nsasncdsimt.Qc_status, &nsasncdsimt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_STORED_IN_MEDIA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteNonSeismicAndSeismicNonConventionalDataStoredInMedia(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT non_seismic_and_seismic_non_conventional_data_stored_in_media_id FROM non_seismic_and_seismic_non_conventional_data_stored_in_media_workspace WHERE non_seismic_and_seismic_non_conventional_data_stored_in_media_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_data_stored_in_media_workspace WHERE non_seismic_and_seismic_non_conventional_data_stored_in_media_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_data_stored_in_media_table WHERE id = :1`, id)
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
func PatchNonSeismicAndSeismicNonConventionalDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var nsasncdsimt dto.Non_seismic_and_seismic_non_conventional_data_stored_in_media
    setDefaults(&nsasncdsimt)

    if err := c.BodyParser(&nsasncdsimt); err != nil{
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
    err = db.QueryRow(`SELECT non_seismic_and_seismic_non_conventional_data_stored_in_media_id FROM non_seismic_and_seismic_non_conventional_data_stored_in_media_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if nsasncdsimt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, nsasncdsimt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, nsasncdsimt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, nsasncdsimt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, nsasncdsimt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET acqtn_survey_name = :1 WHERE id = :2`, nsasncdsimt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Processing_company != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET processing_company = :1 WHERE id = :2`, nsasncdsimt.Processing_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Start_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET start_date = :1 WHERE id = :2`, nsasncdsimt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Seis_station_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET seis_station_type = :1 WHERE id = :2`, nsasncdsimt.Seis_station_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, nsasncdsimt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET proc_set_type = :1 WHERE id = :2`, nsasncdsimt.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, nsasncdsimt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Tape_number != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET tape_number = :1 WHERE id = :2`, nsasncdsimt.Tape_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, nsasncdsimt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, nsasncdsimt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, nsasncdsimt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, nsasncdsimt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_type = :1 WHERE id = :2`, nsasncdsimt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Location_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET location_id = :1 WHERE id = :2`, nsasncdsimt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, nsasncdsimt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, nsasncdsimt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, nsasncdsimt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdsimt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, nsasncdsimt.Checked_by_ba_id, id)
        
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

    