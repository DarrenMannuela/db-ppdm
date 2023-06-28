package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t3dseismicnavigationdatastoredinmedia/dto"
    "github.com/DarrenMannuela/svc-datatype-t3dseismicnavigationdatastoredinmedia/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT3DSeismicNavigationDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_navigation_data_stored_in_media_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_navigation_data_stored_in_media    
    
        for rows.Next() {
    
            var tsndsimt dto.T3D_seismic_navigation_data_stored_in_media
            if err := rows.Scan(&tsndsimt.Id, &tsndsimt.Ba_long_name, &tsndsimt.Ba_type, &tsndsimt.Area_id, &tsndsimt.Area_type, &tsndsimt.Acqtn_survey_name, &tsndsimt.Seis_dimension, &tsndsimt.Shot_by, &tsndsimt.Process_date, &tsndsimt.Digital_format, &tsndsimt.Media_type, &tsndsimt.Tape_number, &tsndsimt.Ba_long_name_2, &tsndsimt.Ba_type_2, &tsndsimt.Data_store_name, &tsndsimt.Data_store_type, &tsndsimt.Location_id, &tsndsimt.Remark, &tsndsimt.Source, &tsndsimt.Qc_status, &tsndsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsndsimt)
        }
    
    return c.JSON(results)
}
func SetT3DSeismicNavigationDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tsndsimt dto.T3D_seismic_navigation_data_stored_in_media
    setDefaults(&tsndsimt)

    if err := c.BodyParser(&tsndsimt); err != nil{
        return err
    }
    
    tsndsimt.Create_date = formatDateString(tsndsimt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t3d_seismic_navigation_data_stored_in_media_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, seis_dimension, shot_by, process_date, digital_format, media_type, tape_number, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20) RETURNING id INTO :21`, &tsndsimt.Ba_long_name, &tsndsimt.Ba_type, &tsndsimt.Area_id, &tsndsimt.Area_type, &tsndsimt.Acqtn_survey_name, &tsndsimt.Seis_dimension, &tsndsimt.Shot_by, &tsndsimt.Process_date, &tsndsimt.Digital_format, &tsndsimt.Media_type, &tsndsimt.Tape_number, &tsndsimt.Ba_long_name_2, &tsndsimt.Ba_type_2, &tsndsimt.Data_store_name, &tsndsimt.Data_store_type, &tsndsimt.Location_id, &tsndsimt.Remark, &tsndsimt.Source, &tsndsimt.Qc_status, &tsndsimt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T3D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT3DSeismicNavigationDataStoredInMediaById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_navigation_data_stored_in_media_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_navigation_data_stored_in_media    
    
        for rows.Next() {
    
            var tsndsimt dto.T3D_seismic_navigation_data_stored_in_media
            if err := rows.Scan(&tsndsimt.Id, &tsndsimt.Ba_long_name, &tsndsimt.Ba_type, &tsndsimt.Area_id, &tsndsimt.Area_type, &tsndsimt.Acqtn_survey_name, &tsndsimt.Seis_dimension, &tsndsimt.Shot_by, &tsndsimt.Process_date, &tsndsimt.Digital_format, &tsndsimt.Media_type, &tsndsimt.Tape_number, &tsndsimt.Ba_long_name_2, &tsndsimt.Ba_type_2, &tsndsimt.Data_store_name, &tsndsimt.Data_store_type, &tsndsimt.Location_id, &tsndsimt.Remark, &tsndsimt.Source, &tsndsimt.Qc_status, &tsndsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsndsimt)
        }
    
    return c.JSON(results)
}
func PutT3DSeismicNavigationDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsndsimt dto.T3D_seismic_navigation_data_stored_in_media
    setDefaults(&tsndsimt)

    if err := c.BodyParser(&tsndsimt); err != nil{
        return err
    }
    
    tsndsimt.Create_date = formatDateString(tsndsimt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t3d_seismic_navigation_data_stored_in_media_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, seis_dimension = :6, shot_by = :7, process_date = :8, digital_format = :9, media_type = :10, tape_number = :11, ba_long_name_2 = :12, ba_type_2 = :13, data_store_name = :14, data_store_type = :15, location_id = :16, remark = :17, source = :18, qc_status = :19, checked_by_ba_id = :20 WHERE id = :21`, &tsndsimt.Ba_long_name, &tsndsimt.Ba_type, &tsndsimt.Area_id, &tsndsimt.Area_type, &tsndsimt.Acqtn_survey_name, &tsndsimt.Seis_dimension, &tsndsimt.Shot_by, &tsndsimt.Process_date, &tsndsimt.Digital_format, &tsndsimt.Media_type, &tsndsimt.Tape_number, &tsndsimt.Ba_long_name_2, &tsndsimt.Ba_type_2, &tsndsimt.Data_store_name, &tsndsimt.Data_store_type, &tsndsimt.Location_id, &tsndsimt.Remark, &tsndsimt.Source, &tsndsimt.Qc_status, &tsndsimt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T3D_SEISMIC_NAVIGATION_DATA_STORED_IN_MEDIA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT3DSeismicNavigationDataStoredInMedia(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t3d_seismic_navigation_data_stored_in_media_id FROM t3d_seismic_navigation_data_stored_in_media_workspace WHERE t3d_seismic_navigation_data_stored_in_media_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t3d_seismic_navigation_data_stored_in_media_workspace WHERE t3d_seismic_navigation_data_stored_in_media_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t3d_seismic_navigation_data_stored_in_media_table WHERE id = :1`, id)
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
func PatchT3DSeismicNavigationDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsndsimt dto.T3D_seismic_navigation_data_stored_in_media
    setDefaults(&tsndsimt)

    if err := c.BodyParser(&tsndsimt); err != nil{
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
    err = db.QueryRow(`SELECT t3d_seismic_navigation_data_stored_in_media_id FROM t3d_seismic_navigation_data_stored_in_media_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tsndsimt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET ba_long_name = :1 WHERE id = :2`, tsndsimt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET ba_type = :1 WHERE id = :2`, tsndsimt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Area_id != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET area_id = :1 WHERE id = :2`, tsndsimt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Area_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET area_type = :1 WHERE id = :2`, tsndsimt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET acqtn_survey_name = :1 WHERE id = :2`, tsndsimt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Seis_dimension != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET seis_dimension = :1 WHERE id = :2`, tsndsimt.Seis_dimension, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Shot_by != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET shot_by = :1 WHERE id = :2`, tsndsimt.Shot_by, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Process_date != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET process_date = :1 WHERE id = :2`, tsndsimt.Process_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET digital_format = :1 WHERE id = :2`, tsndsimt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Media_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET media_type = :1 WHERE id = :2`, tsndsimt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Tape_number != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET tape_number = :1 WHERE id = :2`, tsndsimt.Tape_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET ba_long_name_2 = :1 WHERE id = :2`, tsndsimt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET ba_type_2 = :1 WHERE id = :2`, tsndsimt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET data_store_name = :1 WHERE id = :2`, tsndsimt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET data_store_type = :1 WHERE id = :2`, tsndsimt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Location_id != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET location_id = :1 WHERE id = :2`, tsndsimt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Remark != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET remark = :1 WHERE id = :2`, tsndsimt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Source != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET source = :1 WHERE id = :2`, tsndsimt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET qc_status = :1 WHERE id = :2`, tsndsimt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsndsimt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_navigation_data_stored_in_media_table SET checked_by_ba_id = :1 WHERE id = :2`, tsndsimt.Checked_by_ba_id, id)
        
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

    