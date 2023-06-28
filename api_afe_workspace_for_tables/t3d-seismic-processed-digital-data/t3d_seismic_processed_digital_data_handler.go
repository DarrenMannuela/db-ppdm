package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t3dseismicprocesseddigitaldata/dto"
    "github.com/DarrenMannuela/svc-datatype-t3dseismicprocesseddigitaldata/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT3DSeismicProcessedDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_processed_digital_data_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_processed_digital_data    
    
        for rows.Next() {
    
            var tspddt dto.T3D_seismic_processed_digital_data
            if err := rows.Scan(&tspddt.Id, &tspddt.Ba_long_name, &tspddt.Ba_type, &tspddt.Area_id, &tspddt.Area_type, &tspddt.Acqtn_survey_name, &tspddt.Processing_company, &tspddt.Start_date, &tspddt.Rcrd_rec_length, &tspddt.Rcrd_rec_length_ouom, &tspddt.Rcrd_sample_rate, &tspddt.Rcrd_sample_rate_ouom, &tspddt.First_nline_no, &tspddt.Last_nline_no, &tspddt.First_xline_no, &tspddt.Last_xline_no, &tspddt.Digital_format, &tspddt.Media_type, &tspddt.Proc_set_type, &tspddt.Polarity, &tspddt.Ba_long_name_2, &tspddt.Ba_type_2, &tspddt.Data_store_name, &tspddt.Original_file_name, &tspddt.Password, &tspddt.Sw_application_id, &tspddt.Application_version, &tspddt.Digital_size, &tspddt.Digital_size_uom, &tspddt.Remark, &tspddt.Source, &tspddt.Qc_status, &tspddt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tspddt)
        }
    
    return c.JSON(results)
}
func SetT3DSeismicProcessedDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tspddt dto.T3D_seismic_processed_digital_data
    setDefaults(&tspddt)

    if err := c.BodyParser(&tspddt); err != nil{
        return err
    }
    
    tspddt.Create_date = formatDateString(tspddt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t3d_seismic_processed_digital_data_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, processing_company, start_date, rcrd_rec_length, rcrd_rec_length_ouom, rcrd_sample_rate, rcrd_sample_rate_ouom, first_nline_no, last_nline_no, first_xline_no, last_xline_no, digital_format, media_type, proc_set_type, polarity, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, sw_application_id, application_version, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32) RETURNING id INTO :33`, &tspddt.Ba_long_name, &tspddt.Ba_type, &tspddt.Area_id, &tspddt.Area_type, &tspddt.Acqtn_survey_name, &tspddt.Processing_company, &tspddt.Start_date, &tspddt.Rcrd_rec_length, &tspddt.Rcrd_rec_length_ouom, &tspddt.Rcrd_sample_rate, &tspddt.Rcrd_sample_rate_ouom, &tspddt.First_nline_no, &tspddt.Last_nline_no, &tspddt.First_xline_no, &tspddt.Last_xline_no, &tspddt.Digital_format, &tspddt.Media_type, &tspddt.Proc_set_type, &tspddt.Polarity, &tspddt.Ba_long_name_2, &tspddt.Ba_type_2, &tspddt.Data_store_name, &tspddt.Original_file_name, &tspddt.Password, &tspddt.Sw_application_id, &tspddt.Application_version, &tspddt.Digital_size, &tspddt.Digital_size_uom, &tspddt.Remark, &tspddt.Source, &tspddt.Qc_status, &tspddt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T3D_SEISMIC_PROCESSED_DIGITAL_DATA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT3DSeismicProcessedDigitalDataById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_processed_digital_data_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_processed_digital_data    
    
        for rows.Next() {
    
            var tspddt dto.T3D_seismic_processed_digital_data
            if err := rows.Scan(&tspddt.Id, &tspddt.Ba_long_name, &tspddt.Ba_type, &tspddt.Area_id, &tspddt.Area_type, &tspddt.Acqtn_survey_name, &tspddt.Processing_company, &tspddt.Start_date, &tspddt.Rcrd_rec_length, &tspddt.Rcrd_rec_length_ouom, &tspddt.Rcrd_sample_rate, &tspddt.Rcrd_sample_rate_ouom, &tspddt.First_nline_no, &tspddt.Last_nline_no, &tspddt.First_xline_no, &tspddt.Last_xline_no, &tspddt.Digital_format, &tspddt.Media_type, &tspddt.Proc_set_type, &tspddt.Polarity, &tspddt.Ba_long_name_2, &tspddt.Ba_type_2, &tspddt.Data_store_name, &tspddt.Original_file_name, &tspddt.Password, &tspddt.Sw_application_id, &tspddt.Application_version, &tspddt.Digital_size, &tspddt.Digital_size_uom, &tspddt.Remark, &tspddt.Source, &tspddt.Qc_status, &tspddt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tspddt)
        }
    
    return c.JSON(results)
}
func PutT3DSeismicProcessedDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tspddt dto.T3D_seismic_processed_digital_data
    setDefaults(&tspddt)

    if err := c.BodyParser(&tspddt); err != nil{
        return err
    }
    
    tspddt.Create_date = formatDateString(tspddt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t3d_seismic_processed_digital_data_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t3d_seismic_processed_digital_data_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, processing_company = :6, start_date = :7, rcrd_rec_length = :8, rcrd_rec_length_ouom = :9, rcrd_sample_rate = :10, rcrd_sample_rate_ouom = :11, first_nline_no = :12, last_nline_no = :13, first_xline_no = :14, last_xline_no = :15, digital_format = :16, media_type = :17, proc_set_type = :18, polarity = :19, ba_long_name_2 = :20, ba_type_2 = :21, data_store_name = :22, original_file_name = :23, password = :24, sw_application_id = :25, application_version = :26, digital_size = :27, digital_size_uom = :28, remark = :29, source = :30, qc_status = :31, checked_by_ba_id = :32 WHERE id = :33`, &tspddt.Ba_long_name, &tspddt.Ba_type, &tspddt.Area_id, &tspddt.Area_type, &tspddt.Acqtn_survey_name, &tspddt.Processing_company, &tspddt.Start_date, &tspddt.Rcrd_rec_length, &tspddt.Rcrd_rec_length_ouom, &tspddt.Rcrd_sample_rate, &tspddt.Rcrd_sample_rate_ouom, &tspddt.First_nline_no, &tspddt.Last_nline_no, &tspddt.First_xline_no, &tspddt.Last_xline_no, &tspddt.Digital_format, &tspddt.Media_type, &tspddt.Proc_set_type, &tspddt.Polarity, &tspddt.Ba_long_name_2, &tspddt.Ba_type_2, &tspddt.Data_store_name, &tspddt.Original_file_name, &tspddt.Password, &tspddt.Sw_application_id, &tspddt.Application_version, &tspddt.Digital_size, &tspddt.Digital_size_uom, &tspddt.Remark, &tspddt.Source, &tspddt.Qc_status, &tspddt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T3D_SEISMIC_PROCESSED_DIGITAL_DATA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT3DSeismicProcessedDigitalData(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t3d_seismic_processed_digital_data_id FROM t3d_seismic_processed_digital_data_workspace WHERE t3d_seismic_processed_digital_data_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t3d_seismic_processed_digital_data_workspace WHERE t3d_seismic_processed_digital_data_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t3d_seismic_processed_digital_data_table WHERE id = :1`, id)
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
func PatchT3DSeismicProcessedDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tspddt dto.T3D_seismic_processed_digital_data
    setDefaults(&tspddt)

    if err := c.BodyParser(&tspddt); err != nil{
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
    err = db.QueryRow(`SELECT t3d_seismic_processed_digital_data_id FROM t3d_seismic_processed_digital_data_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tspddt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, tspddt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, tspddt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, tspddt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, tspddt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET acqtn_survey_name = :1 WHERE id = :2`, tspddt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Processing_company != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET processing_company = :1 WHERE id = :2`, tspddt.Processing_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Start_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET start_date = :1 WHERE id = :2`, tspddt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Rcrd_rec_length != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_rec_length = :1 WHERE id = :2`, tspddt.Rcrd_rec_length, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Rcrd_rec_length_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_rec_length_ouom = :1 WHERE id = :2`, tspddt.Rcrd_rec_length_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Rcrd_sample_rate != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_sample_rate = :1 WHERE id = :2`, tspddt.Rcrd_sample_rate, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Rcrd_sample_rate_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_sample_rate_ouom = :1 WHERE id = :2`, tspddt.Rcrd_sample_rate_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.First_nline_no != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET first_nline_no = :1 WHERE id = :2`, tspddt.First_nline_no, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Last_nline_no != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET last_nline_no = :1 WHERE id = :2`, tspddt.Last_nline_no, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.First_xline_no != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET first_xline_no = :1 WHERE id = :2`, tspddt.First_xline_no, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Last_xline_no != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET last_xline_no = :1 WHERE id = :2`, tspddt.Last_xline_no, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, tspddt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, tspddt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET proc_set_type = :1 WHERE id = :2`, tspddt.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Polarity != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET polarity = :1 WHERE id = :2`, tspddt.Polarity, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, tspddt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, tspddt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, tspddt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, tspddt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, tspddt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Sw_application_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sw_application_id = :1 WHERE id = :2`, tspddt.Sw_application_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Application_version != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET application_version = :1 WHERE id = :2`, tspddt.Application_version, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, tspddt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, tspddt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, tspddt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, tspddt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, tspddt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tspddt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, tspddt.Checked_by_ba_id, id)
        
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

    