package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t2dseismicprocessdata/dto"
    "github.com/DarrenMannuela/svc-datatype-t2dseismicprocessdata/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT2DSeismicProcessData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_process_data_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_process_data    
    
        for rows.Next() {
    
            var tspdt dto.T2D_seismic_process_data
            if err := rows.Scan(&tspdt.Id, &tspdt.Ba_long_name, &tspdt.Ba_type, &tspdt.Area_id, &tspdt.Area_type, &tspdt.Acqtn_survey_name, &tspdt.Processing_company, &tspdt.Start_date, &tspdt.Rcrd_rec_length, &tspdt.Rcrd_rec_length_ouom, &tspdt.Rcrd_sample_rate, &tspdt.Rcrd_sample_rate_ouom, &tspdt.Line_name, &tspdt.First_seis_point_id, &tspdt.Last_seis_point_id, &tspdt.Create_date, &tspdt.Proc_set_type, &tspdt.Media_type, &tspdt.Tape_number, &tspdt.Rcrd_format_type, &tspdt.Polarity, &tspdt.Ba_long_name_2, &tspdt.Ba_type_2, &tspdt.Data_store_name, &tspdt.Data_store_type, &tspdt.Location_id, &tspdt.Remark, &tspdt.Source, &tspdt.Qc_status, &tspdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tspdt)
        }
    
    return c.JSON(results)
}
func SetT2DSeismicProcessData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tspdt dto.T2D_seismic_process_data
    setDefaults(&tspdt)

    if err := c.BodyParser(&tspdt); err != nil{
        return err
    }
    
    tspdt.Create_date = formatDateString(tspdt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t2d_seismic_process_data_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, processing_company, start_date, rcrd_rec_length, rcrd_rec_length_ouom, rcrd_sample_rate, rcrd_sample_rate_ouom, line_name, first_seis_point_id, last_seis_point_id, create_date, proc_set_type, media_type, tape_number, rcrd_format_type, polarity, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29) RETURNING id INTO :30`, &tspdt.Ba_long_name, &tspdt.Ba_type, &tspdt.Area_id, &tspdt.Area_type, &tspdt.Acqtn_survey_name, &tspdt.Processing_company, &tspdt.Start_date, &tspdt.Rcrd_rec_length, &tspdt.Rcrd_rec_length_ouom, &tspdt.Rcrd_sample_rate, &tspdt.Rcrd_sample_rate_ouom, &tspdt.Line_name, &tspdt.First_seis_point_id, &tspdt.Last_seis_point_id, &tspdt.Create_date, &tspdt.Proc_set_type, &tspdt.Media_type, &tspdt.Tape_number, &tspdt.Rcrd_format_type, &tspdt.Polarity, &tspdt.Ba_long_name_2, &tspdt.Ba_type_2, &tspdt.Data_store_name, &tspdt.Data_store_type, &tspdt.Location_id, &tspdt.Remark, &tspdt.Source, &tspdt.Qc_status, &tspdt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T2D_SEISMIC_PROCESS_DATA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT2DSeismicProcessDataById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_process_data_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_process_data    
    
        for rows.Next() {
    
            var tspdt dto.T2D_seismic_process_data
            if err := rows.Scan(&tspdt.Id, &tspdt.Ba_long_name, &tspdt.Ba_type, &tspdt.Area_id, &tspdt.Area_type, &tspdt.Acqtn_survey_name, &tspdt.Processing_company, &tspdt.Start_date, &tspdt.Rcrd_rec_length, &tspdt.Rcrd_rec_length_ouom, &tspdt.Rcrd_sample_rate, &tspdt.Rcrd_sample_rate_ouom, &tspdt.Line_name, &tspdt.First_seis_point_id, &tspdt.Last_seis_point_id, &tspdt.Create_date, &tspdt.Proc_set_type, &tspdt.Media_type, &tspdt.Tape_number, &tspdt.Rcrd_format_type, &tspdt.Polarity, &tspdt.Ba_long_name_2, &tspdt.Ba_type_2, &tspdt.Data_store_name, &tspdt.Data_store_type, &tspdt.Location_id, &tspdt.Remark, &tspdt.Source, &tspdt.Qc_status, &tspdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tspdt)
        }
    
    return c.JSON(results)
}
func PutT2DSeismicProcessData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tspdt dto.T2D_seismic_process_data
    setDefaults(&tspdt)

    if err := c.BodyParser(&tspdt); err != nil{
        return err
    }
    
    tspdt.Create_date = formatDateString(tspdt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t2d_seismic_process_data_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, processing_company = :6, start_date = :7, rcrd_rec_length = :8, rcrd_rec_length_ouom = :9, rcrd_sample_rate = :10, rcrd_sample_rate_ouom = :11, line_name = :12, first_seis_point_id = :13, last_seis_point_id = :14, create_date = :15, proc_set_type = :16, media_type = :17, tape_number = :18, rcrd_format_type = :19, polarity = :20, ba_long_name_2 = :21, ba_type_2 = :22, data_store_name = :23, data_store_type = :24, location_id = :25, remark = :26, source = :27, qc_status = :28, checked_by_ba_id = :29 WHERE id = :30`, &tspdt.Ba_long_name, &tspdt.Ba_type, &tspdt.Area_id, &tspdt.Area_type, &tspdt.Acqtn_survey_name, &tspdt.Processing_company, &tspdt.Start_date, &tspdt.Rcrd_rec_length, &tspdt.Rcrd_rec_length_ouom, &tspdt.Rcrd_sample_rate, &tspdt.Rcrd_sample_rate_ouom, &tspdt.Line_name, &tspdt.First_seis_point_id, &tspdt.Last_seis_point_id, &tspdt.Create_date, &tspdt.Proc_set_type, &tspdt.Media_type, &tspdt.Tape_number, &tspdt.Rcrd_format_type, &tspdt.Polarity, &tspdt.Ba_long_name_2, &tspdt.Ba_type_2, &tspdt.Data_store_name, &tspdt.Data_store_type, &tspdt.Location_id, &tspdt.Remark, &tspdt.Source, &tspdt.Qc_status, &tspdt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T2D_SEISMIC_PROCESS_DATA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT2DSeismicProcessData(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t2d_seismic_process_data_id FROM t2d_seismic_process_data_workspace WHERE t2d_seismic_process_data_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t2d_seismic_process_data_workspace WHERE t2d_seismic_process_data_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t2d_seismic_process_data_table WHERE id = :1`, id)
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
func PatchT2DSeismicProcessData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tspdt dto.T2D_seismic_process_data
    setDefaults(&tspdt)

    if err := c.BodyParser(&tspdt); err != nil{
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
    err = db.QueryRow(`SELECT t2d_seismic_process_data_id FROM t2d_seismic_process_data_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tspdt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET ba_long_name = :1 WHERE id = :2`, tspdt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET ba_type = :1 WHERE id = :2`, tspdt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Area_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET area_id = :1 WHERE id = :2`, tspdt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Area_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET area_type = :1 WHERE id = :2`, tspdt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET acqtn_survey_name = :1 WHERE id = :2`, tspdt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Processing_company != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET processing_company = :1 WHERE id = :2`, tspdt.Processing_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Start_date != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET start_date = :1 WHERE id = :2`, tspdt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Rcrd_rec_length != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET rcrd_rec_length = :1 WHERE id = :2`, tspdt.Rcrd_rec_length, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Rcrd_rec_length_ouom != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET rcrd_rec_length_ouom = :1 WHERE id = :2`, tspdt.Rcrd_rec_length_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Rcrd_sample_rate != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET rcrd_sample_rate = :1 WHERE id = :2`, tspdt.Rcrd_sample_rate, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Rcrd_sample_rate_ouom != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET rcrd_sample_rate_ouom = :1 WHERE id = :2`, tspdt.Rcrd_sample_rate_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Line_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET line_name = :1 WHERE id = :2`, tspdt.Line_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.First_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET first_seis_point_id = :1 WHERE id = :2`, tspdt.First_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Last_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET last_seis_point_id = :1 WHERE id = :2`, tspdt.Last_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Create_date != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET create_date = :1 WHERE id = :2`, tspdt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET proc_set_type = :1 WHERE id = :2`, tspdt.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Media_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET media_type = :1 WHERE id = :2`, tspdt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Tape_number != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET tape_number = :1 WHERE id = :2`, tspdt.Tape_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Rcrd_format_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET rcrd_format_type = :1 WHERE id = :2`, tspdt.Rcrd_format_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Polarity != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET polarity = :1 WHERE id = :2`, tspdt.Polarity, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET ba_long_name_2 = :1 WHERE id = :2`, tspdt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET ba_type_2 = :1 WHERE id = :2`, tspdt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET data_store_name = :1 WHERE id = :2`, tspdt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET data_store_type = :1 WHERE id = :2`, tspdt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Location_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET location_id = :1 WHERE id = :2`, tspdt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Remark != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET remark = :1 WHERE id = :2`, tspdt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Source != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET source = :1 WHERE id = :2`, tspdt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET qc_status = :1 WHERE id = :2`, tspdt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspdt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_process_data_table SET checked_by_ba_id = :1 WHERE id = :2`, tspdt.Checked_by_ba_id, id)
        
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

    