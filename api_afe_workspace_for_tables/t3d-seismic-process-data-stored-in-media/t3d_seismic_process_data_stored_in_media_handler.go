package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t3dseismicprocessdatastoredinmedia/dto"
    "github.com/DarrenMannuela/svc-datatype-t3dseismicprocessdatastoredinmedia/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT3DSeismicProcessDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_process_data_stored_in_media_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_process_data_stored_in_media    
    
        for rows.Next() {
    
            var tspdsimt dto.T3D_seismic_process_data_stored_in_media
            if err := rows.Scan(&tspdsimt.Id, &tspdsimt.Ba_long_name, &tspdsimt.Ba_type, &tspdsimt.Area_id, &tspdsimt.Area_type, &tspdsimt.Acqtn_survey_name, &tspdsimt.Processing_company, &tspdsimt.Start_date, &tspdsimt.Rcrd_rec_length, &tspdsimt.Rcrd_rec_length_ouom, &tspdsimt.Rcrd_sample_rate, &tspdsimt.Rcrd_sample_rate_ouom, &tspdsimt.First_nline_no, &tspdsimt.Last_nline_no, &tspdsimt.First_xline_no, &tspdsimt.Last_xline_no, &tspdsimt.Create_date, &tspdsimt.Proc_set_type, &tspdsimt.Media_type, &tspdsimt.Tape_number, &tspdsimt.Digital_format, &tspdsimt.Polarity, &tspdsimt.Ba_long_name_2, &tspdsimt.Ba_type_2, &tspdsimt.Data_store_name, &tspdsimt.Data_store_type, &tspdsimt.Location_id, &tspdsimt.Remark, &tspdsimt.Source, &tspdsimt.Qc_status, &tspdsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tspdsimt)
        }
    
    return c.JSON(results)
}
func SetT3DSeismicProcessDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tspdsimt dto.T3D_seismic_process_data_stored_in_media
    setDefaults(&tspdsimt)

    if err := c.BodyParser(&tspdsimt); err != nil{
        return err
    }
    
    tspdsimt.Create_date = formatDateString(tspdsimt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t3d_seismic_process_data_stored_in_media_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, processing_company, start_date, rcrd_rec_length, rcrd_rec_length_ouom, rcrd_sample_rate, rcrd_sample_rate_ouom, first_nline_no, last_nline_no, first_xline_no, last_xline_no, create_date, proc_set_type, media_type, tape_number, digital_format, polarity, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30) RETURNING id INTO :31`, &tspdsimt.Ba_long_name, &tspdsimt.Ba_type, &tspdsimt.Area_id, &tspdsimt.Area_type, &tspdsimt.Acqtn_survey_name, &tspdsimt.Processing_company, &tspdsimt.Start_date, &tspdsimt.Rcrd_rec_length, &tspdsimt.Rcrd_rec_length_ouom, &tspdsimt.Rcrd_sample_rate, &tspdsimt.Rcrd_sample_rate_ouom, &tspdsimt.First_nline_no, &tspdsimt.Last_nline_no, &tspdsimt.First_xline_no, &tspdsimt.Last_xline_no, &tspdsimt.Create_date, &tspdsimt.Proc_set_type, &tspdsimt.Media_type, &tspdsimt.Tape_number, &tspdsimt.Digital_format, &tspdsimt.Polarity, &tspdsimt.Ba_long_name_2, &tspdsimt.Ba_type_2, &tspdsimt.Data_store_name, &tspdsimt.Data_store_type, &tspdsimt.Location_id, &tspdsimt.Remark, &tspdsimt.Source, &tspdsimt.Qc_status, &tspdsimt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T3D_SEISMIC_PROCESS_DATA_STORED_IN_MEDIA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT3DSeismicProcessDataStoredInMediaById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_process_data_stored_in_media_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_process_data_stored_in_media    
    
        for rows.Next() {
    
            var tspdsimt dto.T3D_seismic_process_data_stored_in_media
            if err := rows.Scan(&tspdsimt.Id, &tspdsimt.Ba_long_name, &tspdsimt.Ba_type, &tspdsimt.Area_id, &tspdsimt.Area_type, &tspdsimt.Acqtn_survey_name, &tspdsimt.Processing_company, &tspdsimt.Start_date, &tspdsimt.Rcrd_rec_length, &tspdsimt.Rcrd_rec_length_ouom, &tspdsimt.Rcrd_sample_rate, &tspdsimt.Rcrd_sample_rate_ouom, &tspdsimt.First_nline_no, &tspdsimt.Last_nline_no, &tspdsimt.First_xline_no, &tspdsimt.Last_xline_no, &tspdsimt.Create_date, &tspdsimt.Proc_set_type, &tspdsimt.Media_type, &tspdsimt.Tape_number, &tspdsimt.Digital_format, &tspdsimt.Polarity, &tspdsimt.Ba_long_name_2, &tspdsimt.Ba_type_2, &tspdsimt.Data_store_name, &tspdsimt.Data_store_type, &tspdsimt.Location_id, &tspdsimt.Remark, &tspdsimt.Source, &tspdsimt.Qc_status, &tspdsimt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tspdsimt)
        }
    
    return c.JSON(results)
}
func PutT3DSeismicProcessDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tspdsimt dto.T3D_seismic_process_data_stored_in_media
    setDefaults(&tspdsimt)

    if err := c.BodyParser(&tspdsimt); err != nil{
        return err
    }
    
    tspdsimt.Create_date = formatDateString(tspdsimt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t3d_seismic_process_data_stored_in_media_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, processing_company = :6, start_date = :7, rcrd_rec_length = :8, rcrd_rec_length_ouom = :9, rcrd_sample_rate = :10, rcrd_sample_rate_ouom = :11, first_nline_no = :12, last_nline_no = :13, first_xline_no = :14, last_xline_no = :15, create_date = :16, proc_set_type = :17, media_type = :18, tape_number = :19, digital_format = :20, polarity = :21, ba_long_name_2 = :22, ba_type_2 = :23, data_store_name = :24, data_store_type = :25, location_id = :26, remark = :27, source = :28, qc_status = :29, checked_by_ba_id = :30 WHERE id = :31`, &tspdsimt.Ba_long_name, &tspdsimt.Ba_type, &tspdsimt.Area_id, &tspdsimt.Area_type, &tspdsimt.Acqtn_survey_name, &tspdsimt.Processing_company, &tspdsimt.Start_date, &tspdsimt.Rcrd_rec_length, &tspdsimt.Rcrd_rec_length_ouom, &tspdsimt.Rcrd_sample_rate, &tspdsimt.Rcrd_sample_rate_ouom, &tspdsimt.First_nline_no, &tspdsimt.Last_nline_no, &tspdsimt.First_xline_no, &tspdsimt.Last_xline_no, &tspdsimt.Create_date, &tspdsimt.Proc_set_type, &tspdsimt.Media_type, &tspdsimt.Tape_number, &tspdsimt.Digital_format, &tspdsimt.Polarity, &tspdsimt.Ba_long_name_2, &tspdsimt.Ba_type_2, &tspdsimt.Data_store_name, &tspdsimt.Data_store_type, &tspdsimt.Location_id, &tspdsimt.Remark, &tspdsimt.Source, &tspdsimt.Qc_status, &tspdsimt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T3D_SEISMIC_PROCESS_DATA_STORED_IN_MEDIA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT3DSeismicProcessDataStoredInMedia(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t3d_seismic_process_data_stored_in_media_id FROM t3d_seismic_process_data_stored_in_media_workspace WHERE t3d_seismic_process_data_stored_in_media_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t3d_seismic_process_data_stored_in_media_workspace WHERE t3d_seismic_process_data_stored_in_media_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t3d_seismic_process_data_stored_in_media_table WHERE id = :1`, id)
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
func PatchT3DSeismicProcessDataStoredInMedia(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tspdsimt dto.T3D_seismic_process_data_stored_in_media
    setDefaults(&tspdsimt)

    if err := c.BodyParser(&tspdsimt); err != nil{
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
    err = db.QueryRow(`SELECT t3d_seismic_process_data_stored_in_media_id FROM t3d_seismic_process_data_stored_in_media_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tspdsimt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET ba_long_name = :1 WHERE id = :2`, tspdsimt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET ba_type = :1 WHERE id = :2`, tspdsimt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Area_id != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET area_id = :1 WHERE id = :2`, tspdsimt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Area_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET area_type = :1 WHERE id = :2`, tspdsimt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET acqtn_survey_name = :1 WHERE id = :2`, tspdsimt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Processing_company != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET processing_company = :1 WHERE id = :2`, tspdsimt.Processing_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Start_date != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET start_date = :1 WHERE id = :2`, tspdsimt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Rcrd_rec_length != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET rcrd_rec_length = :1 WHERE id = :2`, tspdsimt.Rcrd_rec_length, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Rcrd_rec_length_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET rcrd_rec_length_ouom = :1 WHERE id = :2`, tspdsimt.Rcrd_rec_length_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Rcrd_sample_rate != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET rcrd_sample_rate = :1 WHERE id = :2`, tspdsimt.Rcrd_sample_rate, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Rcrd_sample_rate_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET rcrd_sample_rate_ouom = :1 WHERE id = :2`, tspdsimt.Rcrd_sample_rate_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.First_nline_no != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET first_nline_no = :1 WHERE id = :2`, tspdsimt.First_nline_no, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Last_nline_no != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET last_nline_no = :1 WHERE id = :2`, tspdsimt.Last_nline_no, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.First_xline_no != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET first_xline_no = :1 WHERE id = :2`, tspdsimt.First_xline_no, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Last_xline_no != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET last_xline_no = :1 WHERE id = :2`, tspdsimt.Last_xline_no, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Create_date != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET create_date = :1 WHERE id = :2`, tspdsimt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET proc_set_type = :1 WHERE id = :2`, tspdsimt.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Media_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET media_type = :1 WHERE id = :2`, tspdsimt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Tape_number != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET tape_number = :1 WHERE id = :2`, tspdsimt.Tape_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET digital_format = :1 WHERE id = :2`, tspdsimt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Polarity != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET polarity = :1 WHERE id = :2`, tspdsimt.Polarity, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET ba_long_name_2 = :1 WHERE id = :2`, tspdsimt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET ba_type_2 = :1 WHERE id = :2`, tspdsimt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET data_store_name = :1 WHERE id = :2`, tspdsimt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET data_store_type = :1 WHERE id = :2`, tspdsimt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Location_id != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET location_id = :1 WHERE id = :2`, tspdsimt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Remark != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET remark = :1 WHERE id = :2`, tspdsimt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Source != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET source = :1 WHERE id = :2`, tspdsimt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET qc_status = :1 WHERE id = :2`, tspdsimt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdsimt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_process_data_stored_in_media_table SET checked_by_ba_id = :1 WHERE id = :2`, tspdsimt.Checked_by_ba_id, id)
        
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

    