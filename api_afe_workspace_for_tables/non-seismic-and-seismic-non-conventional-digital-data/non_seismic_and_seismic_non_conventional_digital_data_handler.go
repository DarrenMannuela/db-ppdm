package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionaldigitaldata/dto"
    "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionaldigitaldata/utils"

    "github.com/gofiber/fiber/v2"
)
func GetNonSeismicAndSeismicNonConventionalDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM non_seismic_and_seismic_non_conventional_digital_data_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Non_seismic_and_seismic_non_conventional_digital_data    
    
        for rows.Next() {
    
            var nsasncddt dto.Non_seismic_and_seismic_non_conventional_digital_data
            if err := rows.Scan(&nsasncddt.Id, &nsasncddt.Ba_long_name, &nsasncddt.Ba_type, &nsasncddt.Area_id, &nsasncddt.Area_type, &nsasncddt.Acqtn_survey_name, &nsasncddt.Processing_company, &nsasncddt.Start_date, &nsasncddt.Seis_station_type, &nsasncddt.Create_date, &nsasncddt.Proc_set_type, &nsasncddt.Media_type, &nsasncddt.Tape_number, &nsasncddt.Digital_format, &nsasncddt.Ba_long_name_2, &nsasncddt.Ba_type_2, &nsasncddt.Data_store_name, &nsasncddt.Original_file_name, &nsasncddt.Password, &nsasncddt.Digital_size, &nsasncddt.Digital_size_uom, &nsasncddt.Remark, &nsasncddt.Source, &nsasncddt.Qc_status, &nsasncddt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, nsasncddt)
        }
    
    return c.JSON(results)
}
func SetNonSeismicAndSeismicNonConventionalDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var nsasncddt dto.Non_seismic_and_seismic_non_conventional_digital_data
    setDefaults(&nsasncddt)

    if err := c.BodyParser(&nsasncddt); err != nil{
        return err
    }
    
    nsasncddt.Create_date = formatDateString(nsasncddt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO non_seismic_and_seismic_non_conventional_digital_data_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, processing_company, start_date, seis_station_type, create_date, proc_set_type, media_type, tape_number, digital_format, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24) RETURNING id INTO :25`, &nsasncddt.Ba_long_name, &nsasncddt.Ba_type, &nsasncddt.Area_id, &nsasncddt.Area_type, &nsasncddt.Acqtn_survey_name, &nsasncddt.Processing_company, &nsasncddt.Start_date, &nsasncddt.Seis_station_type, &nsasncddt.Create_date, &nsasncddt.Proc_set_type, &nsasncddt.Media_type, &nsasncddt.Tape_number, &nsasncddt.Digital_format, &nsasncddt.Ba_long_name_2, &nsasncddt.Ba_type_2, &nsasncddt.Data_store_name, &nsasncddt.Original_file_name, &nsasncddt.Password, &nsasncddt.Digital_size, &nsasncddt.Digital_size_uom, &nsasncddt.Remark, &nsasncddt.Source, &nsasncddt.Qc_status, &nsasncddt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DIGITAL_DATA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetNonSeismicAndSeismicNonConventionalDigitalDataById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM non_seismic_and_seismic_non_conventional_digital_data_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Non_seismic_and_seismic_non_conventional_digital_data    
    
        for rows.Next() {
    
            var nsasncddt dto.Non_seismic_and_seismic_non_conventional_digital_data
            if err := rows.Scan(&nsasncddt.Id, &nsasncddt.Ba_long_name, &nsasncddt.Ba_type, &nsasncddt.Area_id, &nsasncddt.Area_type, &nsasncddt.Acqtn_survey_name, &nsasncddt.Processing_company, &nsasncddt.Start_date, &nsasncddt.Seis_station_type, &nsasncddt.Create_date, &nsasncddt.Proc_set_type, &nsasncddt.Media_type, &nsasncddt.Tape_number, &nsasncddt.Digital_format, &nsasncddt.Ba_long_name_2, &nsasncddt.Ba_type_2, &nsasncddt.Data_store_name, &nsasncddt.Original_file_name, &nsasncddt.Password, &nsasncddt.Digital_size, &nsasncddt.Digital_size_uom, &nsasncddt.Remark, &nsasncddt.Source, &nsasncddt.Qc_status, &nsasncddt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, nsasncddt)
        }
    
    return c.JSON(results)
}
func PutNonSeismicAndSeismicNonConventionalDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var nsasncddt dto.Non_seismic_and_seismic_non_conventional_digital_data
    setDefaults(&nsasncddt)

    if err := c.BodyParser(&nsasncddt); err != nil{
        return err
    }
    
    nsasncddt.Create_date = formatDateString(nsasncddt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM non_seismic_and_seismic_non_conventional_digital_data_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, processing_company = :6, start_date = :7, seis_station_type = :8, create_date = :9, proc_set_type = :10, media_type = :11, tape_number = :12, digital_format = :13, ba_long_name_2 = :14, ba_type_2 = :15, data_store_name = :16, original_file_name = :17, password = :18, digital_size = :19, digital_size_uom = :20, remark = :21, source = :22, qc_status = :23, checked_by_ba_id = :24 WHERE id = :25`, &nsasncddt.Ba_long_name, &nsasncddt.Ba_type, &nsasncddt.Area_id, &nsasncddt.Area_type, &nsasncddt.Acqtn_survey_name, &nsasncddt.Processing_company, &nsasncddt.Start_date, &nsasncddt.Seis_station_type, &nsasncddt.Create_date, &nsasncddt.Proc_set_type, &nsasncddt.Media_type, &nsasncddt.Tape_number, &nsasncddt.Digital_format, &nsasncddt.Ba_long_name_2, &nsasncddt.Ba_type_2, &nsasncddt.Data_store_name, &nsasncddt.Original_file_name, &nsasncddt.Password, &nsasncddt.Digital_size, &nsasncddt.Digital_size_uom, &nsasncddt.Remark, &nsasncddt.Source, &nsasncddt.Qc_status, &nsasncddt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DIGITAL_DATA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteNonSeismicAndSeismicNonConventionalDigitalData(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT non_seismic_and_seismic_non_conventional_digital_data_id FROM non_seismic_and_seismic_non_conventional_digital_data_workspace WHERE non_seismic_and_seismic_non_conventional_digital_data_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_digital_data_workspace WHERE non_seismic_and_seismic_non_conventional_digital_data_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_digital_data_table WHERE id = :1`, id)
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
func PatchNonSeismicAndSeismicNonConventionalDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var nsasncddt dto.Non_seismic_and_seismic_non_conventional_digital_data
    setDefaults(&nsasncddt)

    if err := c.BodyParser(&nsasncddt); err != nil{
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
    err = db.QueryRow(`SELECT non_seismic_and_seismic_non_conventional_digital_data_id FROM non_seismic_and_seismic_non_conventional_digital_data_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if nsasncddt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET ba_long_name = :1 WHERE id = :2`, nsasncddt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET ba_type = :1 WHERE id = :2`, nsasncddt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET area_id = :1 WHERE id = :2`, nsasncddt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET area_type = :1 WHERE id = :2`, nsasncddt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET acqtn_survey_name = :1 WHERE id = :2`, nsasncddt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Processing_company != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET processing_company = :1 WHERE id = :2`, nsasncddt.Processing_company, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Start_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET start_date = :1 WHERE id = :2`, nsasncddt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Seis_station_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET seis_station_type = :1 WHERE id = :2`, nsasncddt.Seis_station_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET create_date = :1 WHERE id = :2`, nsasncddt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Proc_set_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET proc_set_type = :1 WHERE id = :2`, nsasncddt.Proc_set_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET media_type = :1 WHERE id = :2`, nsasncddt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Tape_number != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET tape_number = :1 WHERE id = :2`, nsasncddt.Tape_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET digital_format = :1 WHERE id = :2`, nsasncddt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET ba_long_name_2 = :1 WHERE id = :2`, nsasncddt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET ba_type_2 = :1 WHERE id = :2`, nsasncddt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET data_store_name = :1 WHERE id = :2`, nsasncddt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET original_file_name = :1 WHERE id = :2`, nsasncddt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET password = :1 WHERE id = :2`, nsasncddt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET digital_size = :1 WHERE id = :2`, nsasncddt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET digital_size_uom = :1 WHERE id = :2`, nsasncddt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET remark = :1 WHERE id = :2`, nsasncddt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET source = :1 WHERE id = :2`, nsasncddt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET qc_status = :1 WHERE id = :2`, nsasncddt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncddt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_digital_data_table SET checked_by_ba_id = :1 WHERE id = :2`, nsasncddt.Checked_by_ba_id, id)
        
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

    