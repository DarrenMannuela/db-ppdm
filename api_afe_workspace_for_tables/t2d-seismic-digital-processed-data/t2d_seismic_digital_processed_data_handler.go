package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t2dseismicdigitalprocesseddata/dto"
    "github.com/DarrenMannuela/svc-datatype-t2dseismicdigitalprocesseddata/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT2DSeismicDigitalProcessedData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_digital_processed_data_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_digital_processed_data    
    
        for rows.Next() {
    
            var tsdpdt dto.T2D_seismic_digital_processed_data
            if err := rows.Scan(&tsdpdt.Id, &tsdpdt.Ba_long_name, &tsdpdt.Ba_type, &tsdpdt.Area_id, &tsdpdt.Area_type, &tsdpdt.Acqtn_survey_name, &tsdpdt.Processing_company, &tsdpdt.Start_date, &tsdpdt.Rcrd_rec_length, &tsdpdt.Rcrd_rec_length_ouom, &tsdpdt.Rcrd_sample_rate, &tsdpdt.Rcrd_sample_rate_ouom, &tsdpdt.Line_name, &tsdpdt.Digital_format, &tsdpdt.Media_type, &tsdpdt.First_seis_point_id, &tsdpdt.Last_seis_point_id, &tsdpdt.Proc_set_type, &tsdpdt.Polarity, &tsdpdt.Ba_long_name_2, &tsdpdt.Ba_type_2, &tsdpdt.Data_store_name, &tsdpdt.Original_file_name, &tsdpdt.Password, &tsdpdt.Sw_application_id, &tsdpdt.Application_version, &tsdpdt.Digital_size, &tsdpdt.Digital_size_uom, &tsdpdt.Remark, &tsdpdt.Source, &tsdpdt.Qc_status, &tsdpdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsdpdt)
        }
    
    return c.JSON(results)
}
func SetT2DSeismicDigitalProcessedData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tsdpdt dto.T2D_seismic_digital_processed_data
    setDefaults(&tsdpdt)

    if err := c.BodyParser(&tsdpdt); err != nil{
        return err
    }
    
    tsdpdt.Create_date = formatDateString(tsdpdt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t2d_seismic_digital_processed_data_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, processing_company, start_date, rcrd_rec_length, rcrd_rec_length_ouom, rcrd_sample_rate, rcrd_sample_rate_ouom, line_name, digital_format, media_type, first_seis_point_id, last_seis_point_id, proc_set_type, polarity, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, sw_application_id, application_version, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31) RETURNING id INTO :32`, &tsdpdt.Ba_long_name, &tsdpdt.Ba_type, &tsdpdt.Area_id, &tsdpdt.Area_type, &tsdpdt.Acqtn_survey_name, &tsdpdt.Processing_company, &tsdpdt.Start_date, &tsdpdt.Rcrd_rec_length, &tsdpdt.Rcrd_rec_length_ouom, &tsdpdt.Rcrd_sample_rate, &tsdpdt.Rcrd_sample_rate_ouom, &tsdpdt.Line_name, &tsdpdt.Digital_format, &tsdpdt.Media_type, &tsdpdt.First_seis_point_id, &tsdpdt.Last_seis_point_id, &tsdpdt.Proc_set_type, &tsdpdt.Polarity, &tsdpdt.Ba_long_name_2, &tsdpdt.Ba_type_2, &tsdpdt.Data_store_name, &tsdpdt.Original_file_name, &tsdpdt.Password, &tsdpdt.Sw_application_id, &tsdpdt.Application_version, &tsdpdt.Digital_size, &tsdpdt.Digital_size_uom, &tsdpdt.Remark, &tsdpdt.Source, &tsdpdt.Qc_status, &tsdpdt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T2D_SEISMIC_DIGITAL_PROCESSED_DATA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT2DSeismicDigitalProcessedDataById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_digital_processed_data_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_digital_processed_data    
    
        for rows.Next() {
    
            var tsdpdt dto.T2D_seismic_digital_processed_data
            if err := rows.Scan(&tsdpdt.Id, &tsdpdt.Ba_long_name, &tsdpdt.Ba_type, &tsdpdt.Area_id, &tsdpdt.Area_type, &tsdpdt.Acqtn_survey_name, &tsdpdt.Processing_company, &tsdpdt.Start_date, &tsdpdt.Rcrd_rec_length, &tsdpdt.Rcrd_rec_length_ouom, &tsdpdt.Rcrd_sample_rate, &tsdpdt.Rcrd_sample_rate_ouom, &tsdpdt.Line_name, &tsdpdt.Digital_format, &tsdpdt.Media_type, &tsdpdt.First_seis_point_id, &tsdpdt.Last_seis_point_id, &tsdpdt.Proc_set_type, &tsdpdt.Polarity, &tsdpdt.Ba_long_name_2, &tsdpdt.Ba_type_2, &tsdpdt.Data_store_name, &tsdpdt.Original_file_name, &tsdpdt.Password, &tsdpdt.Sw_application_id, &tsdpdt.Application_version, &tsdpdt.Digital_size, &tsdpdt.Digital_size_uom, &tsdpdt.Remark, &tsdpdt.Source, &tsdpdt.Qc_status, &tsdpdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsdpdt)
        }
    
    return c.JSON(results)
}
func PutT2DSeismicDigitalProcessedData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsdpdt dto.T2D_seismic_digital_processed_data
    setDefaults(&tsdpdt)

    if err := c.BodyParser(&tsdpdt); err != nil{
        return err
    }
    
    tsdpdt.Create_date = formatDateString(tsdpdt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t2d_seismic_digital_processed_data_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t2d_seismic_digital_processed_data_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, processing_company = :6, start_date = :7, rcrd_rec_length = :8, rcrd_rec_length_ouom = :9, rcrd_sample_rate = :10, rcrd_sample_rate_ouom = :11, line_name = :12, digital_format = :13, media_type = :14, first_seis_point_id = :15, last_seis_point_id = :16, proc_set_type = :17, polarity = :18, ba_long_name_2 = :19, ba_type_2 = :20, data_store_name = :21, original_file_name = :22, password = :23, sw_application_id = :24, application_version = :25, digital_size = :26, digital_size_uom = :27, remark = :28, source = :29, qc_status = :30, checked_by_ba_id = :31 WHERE id = :32`, &tsdpdt.Ba_long_name, &tsdpdt.Ba_type, &tsdpdt.Area_id, &tsdpdt.Area_type, &tsdpdt.Acqtn_survey_name, &tsdpdt.Processing_company, &tsdpdt.Start_date, &tsdpdt.Rcrd_rec_length, &tsdpdt.Rcrd_rec_length_ouom, &tsdpdt.Rcrd_sample_rate, &tsdpdt.Rcrd_sample_rate_ouom, &tsdpdt.Line_name, &tsdpdt.Digital_format, &tsdpdt.Media_type, &tsdpdt.First_seis_point_id, &tsdpdt.Last_seis_point_id, &tsdpdt.Proc_set_type, &tsdpdt.Polarity, &tsdpdt.Ba_long_name_2, &tsdpdt.Ba_type_2, &tsdpdt.Data_store_name, &tsdpdt.Original_file_name, &tsdpdt.Password, &tsdpdt.Sw_application_id, &tsdpdt.Application_version, &tsdpdt.Digital_size, &tsdpdt.Digital_size_uom, &tsdpdt.Remark, &tsdpdt.Source, &tsdpdt.Qc_status, &tsdpdt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T2D_SEISMIC_DIGITAL_PROCESSED_DATA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT2DSeismicDigitalProcessedData(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t2d_seismic_digital_processed_data_id FROM t2d_seismic_digital_processed_data_workspace WHERE t2d_seismic_digital_processed_data_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t2d_seismic_digital_processed_data_workspace WHERE t2d_seismic_digital_processed_data_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t2d_seismic_digital_processed_data_table WHERE id = :1`, id)
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
func PatchT2DSeismicDigitalProcessedData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsdpdt dto.T2D_seismic_digital_processed_data
    setDefaults(&tsdpdt)

    if err := c.BodyParser(&tsdpdt); err != nil{
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
    err = db.QueryRow(`SELECT t2d_seismic_digital_processed_data_id FROM t2d_seismic_digital_processed_data_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tsdpdt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, tsdpdt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, tsdpdt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, tsdpdt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, tsdpdt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET acqtn_survey_name = :1 WHERE id = :2`, tsdpdt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Processing_company != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET processing_company = :1 WHERE id = :2`, tsdpdt.Processing_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Start_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET start_date = :1 WHERE id = :2`, tsdpdt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Rcrd_rec_length != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_rec_length = :1 WHERE id = :2`, tsdpdt.Rcrd_rec_length, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Rcrd_rec_length_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_rec_length_ouom = :1 WHERE id = :2`, tsdpdt.Rcrd_rec_length_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Rcrd_sample_rate != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_sample_rate = :1 WHERE id = :2`, tsdpdt.Rcrd_sample_rate, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Rcrd_sample_rate_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET rcrd_sample_rate_ouom = :1 WHERE id = :2`, tsdpdt.Rcrd_sample_rate_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Line_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET line_name = :1 WHERE id = :2`, tsdpdt.Line_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, tsdpdt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, tsdpdt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.First_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET first_seis_point_id = :1 WHERE id = :2`, tsdpdt.First_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Last_seis_point_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET last_seis_point_id = :1 WHERE id = :2`, tsdpdt.Last_seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET proc_set_type = :1 WHERE id = :2`, tsdpdt.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Polarity != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET polarity = :1 WHERE id = :2`, tsdpdt.Polarity, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, tsdpdt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, tsdpdt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, tsdpdt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, tsdpdt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, tsdpdt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Sw_application_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sw_application_id = :1 WHERE id = :2`, tsdpdt.Sw_application_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Application_version != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET application_version = :1 WHERE id = :2`, tsdpdt.Application_version, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, tsdpdt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, tsdpdt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, tsdpdt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, tsdpdt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, tsdpdt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsdpdt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, tsdpdt.Checked_by_ba_id, id)
        
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

    