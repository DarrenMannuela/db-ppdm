package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t2dseismicfielddata/dto"
    "github.com/DarrenMannuela/svc-datatype-t2dseismicfielddata/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT2DSeismicFieldData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_field_data_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_field_data    
    
        for rows.Next() {
    
            var tsfdt dto.T2D_seismic_field_data
            if err := rows.Scan(&tsfdt.Id, &tsfdt.Ba_long_name, &tsfdt.Ba_type, &tsfdt.Area_id, &tsfdt.Area_type, &tsfdt.Acqtn_survey_name, &tsfdt.Shot_by, &tsfdt.Start_date, &tsfdt.Rcrd_rec_length, &tsfdt.Rcrd_rec_length_ouom, &tsfdt.Rcrd_sample_rate, &tsfdt.Rcrd_sample_rate_ouom, &tsfdt.Line_name, &tsfdt.First_seis_point_id, &tsfdt.Last_seis_point_id, &tsfdt.Create_date, &tsfdt.Proc_set_type, &tsfdt.Field_file_number, &tsfdt.Media_type, &tsfdt.Tape_number, &tsfdt.Rcrd_format_type, &tsfdt.Ba_long_name_2, &tsfdt.Ba_type_2, &tsfdt.Data_store_name, &tsfdt.Data_store_type, &tsfdt.Location_id, &tsfdt.Remark, &tsfdt.Source, &tsfdt.Qc_status, &tsfdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsfdt)
        }
    
    return c.JSON(results)
}
func SetT2DSeismicFieldData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tsfdt dto.T2D_seismic_field_data
    setDefaults(&tsfdt)

    if err := c.BodyParser(&tsfdt); err != nil{
        return err
    }
    
    tsfdt.Create_date = formatDateString(tsfdt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t2d_seismic_field_data_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, shot_by, start_date, rcrd_rec_length, rcrd_rec_length_ouom, rcrd_sample_rate, rcrd_sample_rate_ouom, line_name, first_seis_point_id, last_seis_point_id, create_date, proc_set_type, field_file_number, media_type, tape_number, rcrd_format_type, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29) RETURNING id INTO :30`, &tsfdt.Ba_long_name, &tsfdt.Ba_type, &tsfdt.Area_id, &tsfdt.Area_type, &tsfdt.Acqtn_survey_name, &tsfdt.Shot_by, &tsfdt.Start_date, &tsfdt.Rcrd_rec_length, &tsfdt.Rcrd_rec_length_ouom, &tsfdt.Rcrd_sample_rate, &tsfdt.Rcrd_sample_rate_ouom, &tsfdt.Line_name, &tsfdt.First_seis_point_id, &tsfdt.Last_seis_point_id, &tsfdt.Create_date, &tsfdt.Proc_set_type, &tsfdt.Field_file_number, &tsfdt.Media_type, &tsfdt.Tape_number, &tsfdt.Rcrd_format_type, &tsfdt.Ba_long_name_2, &tsfdt.Ba_type_2, &tsfdt.Data_store_name, &tsfdt.Data_store_type, &tsfdt.Location_id, &tsfdt.Remark, &tsfdt.Source, &tsfdt.Qc_status, &tsfdt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T2D_SEISMIC_FIELD_DATA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT2DSeismicFieldDataById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_field_data_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_field_data    
    
        for rows.Next() {
    
            var tsfdt dto.T2D_seismic_field_data
            if err := rows.Scan(&tsfdt.Id, &tsfdt.Ba_long_name, &tsfdt.Ba_type, &tsfdt.Area_id, &tsfdt.Area_type, &tsfdt.Acqtn_survey_name, &tsfdt.Shot_by, &tsfdt.Start_date, &tsfdt.Rcrd_rec_length, &tsfdt.Rcrd_rec_length_ouom, &tsfdt.Rcrd_sample_rate, &tsfdt.Rcrd_sample_rate_ouom, &tsfdt.Line_name, &tsfdt.First_seis_point_id, &tsfdt.Last_seis_point_id, &tsfdt.Create_date, &tsfdt.Proc_set_type, &tsfdt.Field_file_number, &tsfdt.Media_type, &tsfdt.Tape_number, &tsfdt.Rcrd_format_type, &tsfdt.Ba_long_name_2, &tsfdt.Ba_type_2, &tsfdt.Data_store_name, &tsfdt.Data_store_type, &tsfdt.Location_id, &tsfdt.Remark, &tsfdt.Source, &tsfdt.Qc_status, &tsfdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsfdt)
        }
    
    return c.JSON(results)
}
func PutT2DSeismicFieldData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsfdt dto.T2D_seismic_field_data
    setDefaults(&tsfdt)

    if err := c.BodyParser(&tsfdt); err != nil{
        return err
    }
    
    tsfdt.Create_date = formatDateString(tsfdt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t2d_seismic_field_data_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, shot_by = :6, start_date = :7, rcrd_rec_length = :8, rcrd_rec_length_ouom = :9, rcrd_sample_rate = :10, rcrd_sample_rate_ouom = :11, line_name = :12, first_seis_point_id = :13, last_seis_point_id = :14, create_date = :15, proc_set_type = :16, field_file_number = :17, media_type = :18, tape_number = :19, rcrd_format_type = :20, ba_long_name_2 = :21, ba_type_2 = :22, data_store_name = :23, data_store_type = :24, location_id = :25, remark = :26, source = :27, qc_status = :28, checked_by_ba_id = :29 WHERE id = :30`, &tsfdt.Ba_long_name, &tsfdt.Ba_type, &tsfdt.Area_id, &tsfdt.Area_type, &tsfdt.Acqtn_survey_name, &tsfdt.Shot_by, &tsfdt.Start_date, &tsfdt.Rcrd_rec_length, &tsfdt.Rcrd_rec_length_ouom, &tsfdt.Rcrd_sample_rate, &tsfdt.Rcrd_sample_rate_ouom, &tsfdt.Line_name, &tsfdt.First_seis_point_id, &tsfdt.Last_seis_point_id, &tsfdt.Create_date, &tsfdt.Proc_set_type, &tsfdt.Field_file_number, &tsfdt.Media_type, &tsfdt.Tape_number, &tsfdt.Rcrd_format_type, &tsfdt.Ba_long_name_2, &tsfdt.Ba_type_2, &tsfdt.Data_store_name, &tsfdt.Data_store_type, &tsfdt.Location_id, &tsfdt.Remark, &tsfdt.Source, &tsfdt.Qc_status, &tsfdt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T2D_SEISMIC_FIELD_DATA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT2DSeismicFieldData(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t2d_seismic_field_data_id FROM t2d_seismic_field_data_workspace WHERE t2d_seismic_field_data_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t2d_seismic_field_data_workspace WHERE t2d_seismic_field_data_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t2d_seismic_field_data_table WHERE id = :1`, id)
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
func PatchT2DSeismicFieldData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsfdt dto.T2D_seismic_field_data
    setDefaults(&tsfdt)

    if err := c.BodyParser(&tsfdt); err != nil{
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
    err = db.QueryRow(`SELECT t2d_seismic_field_data_id FROM t2d_seismic_field_data_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tsfdt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET ba_long_name = :1 WHERE id = :2`, tsfdt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET ba_type = :1 WHERE id = :2`, tsfdt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Area_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET area_id = :1 WHERE id = :2`, tsfdt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Area_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET area_type = :1 WHERE id = :2`, tsfdt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET acqtn_survey_name = :1 WHERE id = :2`, tsfdt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Shot_by != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET shot_by = :1 WHERE id = :2`, tsfdt.Shot_by, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Start_date != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET start_date = :1 WHERE id = :2`, tsfdt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Rcrd_rec_length != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET rcrd_rec_length = :1 WHERE id = :2`, tsfdt.Rcrd_rec_length, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Rcrd_rec_length_ouom != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET rcrd_rec_length_ouom = :1 WHERE id = :2`, tsfdt.Rcrd_rec_length_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Rcrd_sample_rate != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET rcrd_sample_rate = :1 WHERE id = :2`, tsfdt.Rcrd_sample_rate, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Rcrd_sample_rate_ouom != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET rcrd_sample_rate_ouom = :1 WHERE id = :2`, tsfdt.Rcrd_sample_rate_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Line_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET line_name = :1 WHERE id = :2`, tsfdt.Line_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.First_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET first_seis_point_id = :1 WHERE id = :2`, tsfdt.First_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Last_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET last_seis_point_id = :1 WHERE id = :2`, tsfdt.Last_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Create_date != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET create_date = :1 WHERE id = :2`, tsfdt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET proc_set_type = :1 WHERE id = :2`, tsfdt.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Field_file_number != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET field_file_number = :1 WHERE id = :2`, tsfdt.Field_file_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Media_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET media_type = :1 WHERE id = :2`, tsfdt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Tape_number != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET tape_number = :1 WHERE id = :2`, tsfdt.Tape_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Rcrd_format_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET rcrd_format_type = :1 WHERE id = :2`, tsfdt.Rcrd_format_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET ba_long_name_2 = :1 WHERE id = :2`, tsfdt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET ba_type_2 = :1 WHERE id = :2`, tsfdt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET data_store_name = :1 WHERE id = :2`, tsfdt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET data_store_type = :1 WHERE id = :2`, tsfdt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Location_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET location_id = :1 WHERE id = :2`, tsfdt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Remark != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET remark = :1 WHERE id = :2`, tsfdt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Source != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET source = :1 WHERE id = :2`, tsfdt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET qc_status = :1 WHERE id = :2`, tsfdt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfdt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_field_data_table SET checked_by_ba_id = :1 WHERE id = :2`, tsfdt.Checked_by_ba_id, id)
        
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

    