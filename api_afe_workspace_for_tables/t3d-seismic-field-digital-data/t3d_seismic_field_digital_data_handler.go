package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t3dseismicfielddigitaldata/dto"
    "github.com/DarrenMannuela/svc-datatype-t3dseismicfielddigitaldata/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT3DSeismicFieldDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_field_digital_data_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_field_digital_data    
    
        for rows.Next() {
    
            var tsfddt dto.T3D_seismic_field_digital_data
            if err := rows.Scan(&tsfddt.Id, &tsfddt.Ba_long_name, &tsfddt.Ba_type, &tsfddt.Area_id, &tsfddt.Area_type, &tsfddt.Acqtn_survey_name, &tsfddt.Start_date, &tsfddt.Shot_by, &tsfddt.Rcrd_rec_length, &tsfddt.Rcrd_rec_length_ouom, &tsfddt.Rcrd_sample_rate, &tsfddt.Rcrd_sample_rate_ouom, &tsfddt.First_seis_point_id, &tsfddt.Last_seis_point_id, &tsfddt.Create_date, &tsfddt.Proc_set_type, &tsfddt.Field_file_number, &tsfddt.Tape_number, &tsfddt.Rcrd_format_type, &tsfddt.Ba_long_name_2, &tsfddt.Ba_type_2, &tsfddt.Data_store_name, &tsfddt.Original_file_name, &tsfddt.Password, &tsfddt.Digital_size, &tsfddt.Digital_size_uom, &tsfddt.Remark, &tsfddt.Source, &tsfddt.Qc_status, &tsfddt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsfddt)
        }
    
    return c.JSON(results)
}
func SetT3DSeismicFieldDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tsfddt dto.T3D_seismic_field_digital_data
    setDefaults(&tsfddt)

    if err := c.BodyParser(&tsfddt); err != nil{
        return err
    }
    
    tsfddt.Create_date = formatDateString(tsfddt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t3d_seismic_field_digital_data_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, start_date, shot_by, rcrd_rec_length, rcrd_rec_length_ouom, rcrd_sample_rate, rcrd_sample_rate_ouom, first_seis_point_id, last_seis_point_id, create_date, proc_set_type, field_file_number, tape_number, rcrd_format_type, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29) RETURNING id INTO :30`, &tsfddt.Ba_long_name, &tsfddt.Ba_type, &tsfddt.Area_id, &tsfddt.Area_type, &tsfddt.Acqtn_survey_name, &tsfddt.Start_date, &tsfddt.Shot_by, &tsfddt.Rcrd_rec_length, &tsfddt.Rcrd_rec_length_ouom, &tsfddt.Rcrd_sample_rate, &tsfddt.Rcrd_sample_rate_ouom, &tsfddt.First_seis_point_id, &tsfddt.Last_seis_point_id, &tsfddt.Create_date, &tsfddt.Proc_set_type, &tsfddt.Field_file_number, &tsfddt.Tape_number, &tsfddt.Rcrd_format_type, &tsfddt.Ba_long_name_2, &tsfddt.Ba_type_2, &tsfddt.Data_store_name, &tsfddt.Original_file_name, &tsfddt.Password, &tsfddt.Digital_size, &tsfddt.Digital_size_uom, &tsfddt.Remark, &tsfddt.Source, &tsfddt.Qc_status, &tsfddt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T3D_SEISMIC_FIELD_DIGITAL_DATA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT3DSeismicFieldDigitalDataById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_field_digital_data_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_field_digital_data    
    
        for rows.Next() {
    
            var tsfddt dto.T3D_seismic_field_digital_data
            if err := rows.Scan(&tsfddt.Id, &tsfddt.Ba_long_name, &tsfddt.Ba_type, &tsfddt.Area_id, &tsfddt.Area_type, &tsfddt.Acqtn_survey_name, &tsfddt.Start_date, &tsfddt.Shot_by, &tsfddt.Rcrd_rec_length, &tsfddt.Rcrd_rec_length_ouom, &tsfddt.Rcrd_sample_rate, &tsfddt.Rcrd_sample_rate_ouom, &tsfddt.First_seis_point_id, &tsfddt.Last_seis_point_id, &tsfddt.Create_date, &tsfddt.Proc_set_type, &tsfddt.Field_file_number, &tsfddt.Tape_number, &tsfddt.Rcrd_format_type, &tsfddt.Ba_long_name_2, &tsfddt.Ba_type_2, &tsfddt.Data_store_name, &tsfddt.Original_file_name, &tsfddt.Password, &tsfddt.Digital_size, &tsfddt.Digital_size_uom, &tsfddt.Remark, &tsfddt.Source, &tsfddt.Qc_status, &tsfddt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsfddt)
        }
    
    return c.JSON(results)
}
func PutT3DSeismicFieldDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsfddt dto.T3D_seismic_field_digital_data
    setDefaults(&tsfddt)

    if err := c.BodyParser(&tsfddt); err != nil{
        return err
    }
    
    tsfddt.Create_date = formatDateString(tsfddt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t3d_seismic_field_digital_data_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t3d_seismic_field_digital_data_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, start_date = :6, shot_by = :7, rcrd_rec_length = :8, rcrd_rec_length_ouom = :9, rcrd_sample_rate = :10, rcrd_sample_rate_ouom = :11, first_seis_point_id = :12, last_seis_point_id = :13, create_date = :14, proc_set_type = :15, field_file_number = :16, tape_number = :17, rcrd_format_type = :18, ba_long_name_2 = :19, ba_type_2 = :20, data_store_name = :21, original_file_name = :22, password = :23, digital_size = :24, digital_size_uom = :25, remark = :26, source = :27, qc_status = :28, checked_by_ba_id = :29 WHERE id = :30`, &tsfddt.Ba_long_name, &tsfddt.Ba_type, &tsfddt.Area_id, &tsfddt.Area_type, &tsfddt.Acqtn_survey_name, &tsfddt.Start_date, &tsfddt.Shot_by, &tsfddt.Rcrd_rec_length, &tsfddt.Rcrd_rec_length_ouom, &tsfddt.Rcrd_sample_rate, &tsfddt.Rcrd_sample_rate_ouom, &tsfddt.First_seis_point_id, &tsfddt.Last_seis_point_id, &tsfddt.Create_date, &tsfddt.Proc_set_type, &tsfddt.Field_file_number, &tsfddt.Tape_number, &tsfddt.Rcrd_format_type, &tsfddt.Ba_long_name_2, &tsfddt.Ba_type_2, &tsfddt.Data_store_name, &tsfddt.Original_file_name, &tsfddt.Password, &tsfddt.Digital_size, &tsfddt.Digital_size_uom, &tsfddt.Remark, &tsfddt.Source, &tsfddt.Qc_status, &tsfddt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T3D_SEISMIC_FIELD_DIGITAL_DATA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT3DSeismicFieldDigitalData(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t3d_seismic_field_digital_data_id FROM t3d_seismic_field_digital_data_workspace WHERE t3d_seismic_field_digital_data_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t3d_seismic_field_digital_data_workspace WHERE t3d_seismic_field_digital_data_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t3d_seismic_field_digital_data_table WHERE id = :1`, id)
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
func PatchT3DSeismicFieldDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsfddt dto.T3D_seismic_field_digital_data
    setDefaults(&tsfddt)

    if err := c.BodyParser(&tsfddt); err != nil{
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
    err = db.QueryRow(`SELECT t3d_seismic_field_digital_data_id FROM t3d_seismic_field_digital_data_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tsfddt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, tsfddt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, tsfddt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, tsfddt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, tsfddt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET acqtn_survey_name = :1 WHERE id = :2`, tsfddt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Start_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET start_date = :1 WHERE id = :2`, tsfddt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Shot_by != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET shot_by = :1 WHERE id = :2`, tsfddt.Shot_by, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Rcrd_rec_length != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_rec_length = :1 WHERE id = :2`, tsfddt.Rcrd_rec_length, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Rcrd_rec_length_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_rec_length_ouom = :1 WHERE id = :2`, tsfddt.Rcrd_rec_length_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Rcrd_sample_rate != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_sample_rate = :1 WHERE id = :2`, tsfddt.Rcrd_sample_rate, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Rcrd_sample_rate_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_sample_rate_ouom = :1 WHERE id = :2`, tsfddt.Rcrd_sample_rate_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.First_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET first_seis_point_id = :1 WHERE id = :2`, tsfddt.First_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Last_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET last_seis_point_id = :1 WHERE id = :2`, tsfddt.Last_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, tsfddt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET proc_set_type = :1 WHERE id = :2`, tsfddt.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Field_file_number != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_file_number = :1 WHERE id = :2`, tsfddt.Field_file_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Tape_number != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET tape_number = :1 WHERE id = :2`, tsfddt.Tape_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Rcrd_format_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_format_type = :1 WHERE id = :2`, tsfddt.Rcrd_format_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, tsfddt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, tsfddt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, tsfddt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, tsfddt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, tsfddt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, tsfddt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, tsfddt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, tsfddt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, tsfddt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, tsfddt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsfddt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, tsfddt.Checked_by_ba_id, id)
        
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

    