package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t3dseismicfielddatastoredinmedia/dto"
    "github.com/DarrenMannuela/svc-datatype-t3dseismicfielddatastoredinmedia/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT3DSeismicFieldDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_field_data_stored_in_media_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_field_data_stored_in_media    
    
        for rows.Next() {
    
            var tsfdsimt dto.T3D_seismic_field_data_stored_in_media
            if err := rows.Scan(&tsfdsimt.Id, &tsfdsimt.Ba_long_name, &tsfdsimt.Ba_type, &tsfdsimt.Area_id, &tsfdsimt.Area_type, &tsfdsimt.Acqtn_survey_name, &tsfdsimt.Start_date, &tsfdsimt.Shot_by, &tsfdsimt.Rcrd_rec_length, &tsfdsimt.Rcrd_rec_length_ouom, &tsfdsimt.Rcrd_sample_rate, &tsfdsimt.Rcrd_sample_rate_ouom, &tsfdsimt.First_seis_point_id, &tsfdsimt.Last_seis_point_id, &tsfdsimt.Create_date, &tsfdsimt.Proc_set_type, &tsfdsimt.Media_type, &tsfdsimt.Field_file_number, &tsfdsimt.Tape_number, &tsfdsimt.Rcrd_format_type, &tsfdsimt.Ba_long_name_2, &tsfdsimt.Ba_type_2, &tsfdsimt.Data_store_name, &tsfdsimt.Data_store_type, &tsfdsimt.Location_id, &tsfdsimt.Remark, &tsfdsimt.Source, &tsfdsimt.Qc_status, &tsfdsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsfdsimt)
        }
    
    return c.JSON(results)
}
func SetT3DSeismicFieldDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tsfdsimt dto.T3D_seismic_field_data_stored_in_media
    setDefaults(&tsfdsimt)

    if err := c.BodyParser(&tsfdsimt); err != nil{
        return err
    }
    
    tsfdsimt.Create_date = formatDateString(tsfdsimt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t3d_seismic_field_data_stored_in_media_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, start_date, shot_by, rcrd_rec_length, rcrd_rec_length_ouom, rcrd_sample_rate, rcrd_sample_rate_ouom, first_seis_point_id, last_seis_point_id, create_date, proc_set_type, media_type, field_file_number, tape_number, rcrd_format_type, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28) RETURNING id INTO :29`, &tsfdsimt.Ba_long_name, &tsfdsimt.Ba_type, &tsfdsimt.Area_id, &tsfdsimt.Area_type, &tsfdsimt.Acqtn_survey_name, &tsfdsimt.Start_date, &tsfdsimt.Shot_by, &tsfdsimt.Rcrd_rec_length, &tsfdsimt.Rcrd_rec_length_ouom, &tsfdsimt.Rcrd_sample_rate, &tsfdsimt.Rcrd_sample_rate_ouom, &tsfdsimt.First_seis_point_id, &tsfdsimt.Last_seis_point_id, &tsfdsimt.Create_date, &tsfdsimt.Proc_set_type, &tsfdsimt.Media_type, &tsfdsimt.Field_file_number, &tsfdsimt.Tape_number, &tsfdsimt.Rcrd_format_type, &tsfdsimt.Ba_long_name_2, &tsfdsimt.Ba_type_2, &tsfdsimt.Data_store_name, &tsfdsimt.Data_store_type, &tsfdsimt.Location_id, &tsfdsimt.Remark, &tsfdsimt.Source, &tsfdsimt.Qc_status, &tsfdsimt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T3D_SEISMIC_FIELD_DATA_STORED_IN_MEDIA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT3DSeismicFieldDataStoredInMediaById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_field_data_stored_in_media_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_field_data_stored_in_media    
    
        for rows.Next() {
    
            var tsfdsimt dto.T3D_seismic_field_data_stored_in_media
            if err := rows.Scan(&tsfdsimt.Id, &tsfdsimt.Ba_long_name, &tsfdsimt.Ba_type, &tsfdsimt.Area_id, &tsfdsimt.Area_type, &tsfdsimt.Acqtn_survey_name, &tsfdsimt.Start_date, &tsfdsimt.Shot_by, &tsfdsimt.Rcrd_rec_length, &tsfdsimt.Rcrd_rec_length_ouom, &tsfdsimt.Rcrd_sample_rate, &tsfdsimt.Rcrd_sample_rate_ouom, &tsfdsimt.First_seis_point_id, &tsfdsimt.Last_seis_point_id, &tsfdsimt.Create_date, &tsfdsimt.Proc_set_type, &tsfdsimt.Media_type, &tsfdsimt.Field_file_number, &tsfdsimt.Tape_number, &tsfdsimt.Rcrd_format_type, &tsfdsimt.Ba_long_name_2, &tsfdsimt.Ba_type_2, &tsfdsimt.Data_store_name, &tsfdsimt.Data_store_type, &tsfdsimt.Location_id, &tsfdsimt.Remark, &tsfdsimt.Source, &tsfdsimt.Qc_status, &tsfdsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsfdsimt)
        }
    
    return c.JSON(results)
}
func PutT3DSeismicFieldDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsfdsimt dto.T3D_seismic_field_data_stored_in_media
    setDefaults(&tsfdsimt)

    if err := c.BodyParser(&tsfdsimt); err != nil{
        return err
    }
    
    tsfdsimt.Create_date = formatDateString(tsfdsimt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t3d_seismic_field_data_stored_in_media_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t3d_seismic_field_data_stored_in_media_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, start_date = :6, shot_by = :7, rcrd_rec_length = :8, rcrd_rec_length_ouom = :9, rcrd_sample_rate = :10, rcrd_sample_rate_ouom = :11, first_seis_point_id = :12, last_seis_point_id = :13, create_date = :14, proc_set_type = :15, media_type = :16, field_file_number = :17, tape_number = :18, rcrd_format_type = :19, ba_long_name_2 = :20, ba_type_2 = :21, data_store_name = :22, data_store_type = :23, location_id = :24, remark = :25, source = :26, qc_status = :27, checked_by_ba_id = :28 WHERE id = :29`, &tsfdsimt.Ba_long_name, &tsfdsimt.Ba_type, &tsfdsimt.Area_id, &tsfdsimt.Area_type, &tsfdsimt.Acqtn_survey_name, &tsfdsimt.Start_date, &tsfdsimt.Shot_by, &tsfdsimt.Rcrd_rec_length, &tsfdsimt.Rcrd_rec_length_ouom, &tsfdsimt.Rcrd_sample_rate, &tsfdsimt.Rcrd_sample_rate_ouom, &tsfdsimt.First_seis_point_id, &tsfdsimt.Last_seis_point_id, &tsfdsimt.Create_date, &tsfdsimt.Proc_set_type, &tsfdsimt.Media_type, &tsfdsimt.Field_file_number, &tsfdsimt.Tape_number, &tsfdsimt.Rcrd_format_type, &tsfdsimt.Ba_long_name_2, &tsfdsimt.Ba_type_2, &tsfdsimt.Data_store_name, &tsfdsimt.Data_store_type, &tsfdsimt.Location_id, &tsfdsimt.Remark, &tsfdsimt.Source, &tsfdsimt.Qc_status, &tsfdsimt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T3D_SEISMIC_FIELD_DATA_STORED_IN_MEDIA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT3DSeismicFieldDataStoredInMedia(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t3d_seismic_field_data_stored_in_media_id FROM t3d_seismic_field_data_stored_in_media_workspace WHERE t3d_seismic_field_data_stored_in_media_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t3d_seismic_field_data_stored_in_media_workspace WHERE t3d_seismic_field_data_stored_in_media_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t3d_seismic_field_data_stored_in_media_table WHERE id = :1`, id)
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
func PatchT3DSeismicFieldDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsfdsimt dto.T3D_seismic_field_data_stored_in_media
    setDefaults(&tsfdsimt)

    if err := c.BodyParser(&tsfdsimt); err != nil{
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
    err = db.QueryRow(`SELECT t3d_seismic_field_data_stored_in_media_id FROM t3d_seismic_field_data_stored_in_media_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tsfdsimt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, tsfdsimt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, tsfdsimt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, tsfdsimt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, tsfdsimt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET acqtn_survey_name = :1 WHERE id = :2`, tsfdsimt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Start_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET start_date = :1 WHERE id = :2`, tsfdsimt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Shot_by != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET shot_by = :1 WHERE id = :2`, tsfdsimt.Shot_by, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Rcrd_rec_length != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_rec_length = :1 WHERE id = :2`, tsfdsimt.Rcrd_rec_length, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Rcrd_rec_length_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_rec_length_ouom = :1 WHERE id = :2`, tsfdsimt.Rcrd_rec_length_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Rcrd_sample_rate != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_sample_rate = :1 WHERE id = :2`, tsfdsimt.Rcrd_sample_rate, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Rcrd_sample_rate_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_sample_rate_ouom = :1 WHERE id = :2`, tsfdsimt.Rcrd_sample_rate_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.First_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET first_seis_point_id = :1 WHERE id = :2`, tsfdsimt.First_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Last_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET last_seis_point_id = :1 WHERE id = :2`, tsfdsimt.Last_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, tsfdsimt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET proc_set_type = :1 WHERE id = :2`, tsfdsimt.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, tsfdsimt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Field_file_number != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_file_number = :1 WHERE id = :2`, tsfdsimt.Field_file_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Tape_number != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET tape_number = :1 WHERE id = :2`, tsfdsimt.Tape_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Rcrd_format_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_format_type = :1 WHERE id = :2`, tsfdsimt.Rcrd_format_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, tsfdsimt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, tsfdsimt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, tsfdsimt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_type = :1 WHERE id = :2`, tsfdsimt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Location_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET location_id = :1 WHERE id = :2`, tsfdsimt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, tsfdsimt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, tsfdsimt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, tsfdsimt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdsimt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, tsfdsimt.Checked_by_ba_id, id)
        
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

    