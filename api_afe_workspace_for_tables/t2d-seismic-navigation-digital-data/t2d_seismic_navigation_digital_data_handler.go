package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t2dseismicnavigationdigitaldata/dto"
    "github.com/DarrenMannuela/svc-datatype-t2dseismicnavigationdigitaldata/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT2DSeismicNavigationDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_navigation_digital_data_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_navigation_digital_data    
    
        for rows.Next() {
    
            var tsnddt dto.T2D_seismic_navigation_digital_data
            if err := rows.Scan(&tsnddt.Id, &tsnddt.Ba_long_name, &tsnddt.Ba_type, &tsnddt.Area_id, &tsnddt.Area_type, &tsnddt.Acqtn_survey_name, &tsnddt.Seis_dimension, &tsnddt.Process_date, &tsnddt.Shot_by, &tsnddt.Line_name, &tsnddt.Digital_format, &tsnddt.Media_type, &tsnddt.Original_file_name, &tsnddt.Password, &tsnddt.Digital_size, &tsnddt.Digital_size_uom, &tsnddt.Ba_long_name_2, &tsnddt.Ba_type_2, &tsnddt.Data_store_name, &tsnddt.Remark, &tsnddt.Source, &tsnddt.Qc_status, &tsnddt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsnddt)
        }
    
    return c.JSON(results)
}
func SetT2DSeismicNavigationDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tsnddt dto.T2D_seismic_navigation_digital_data
    setDefaults(&tsnddt)

    if err := c.BodyParser(&tsnddt); err != nil{
        return err
    }
    
    tsnddt.Create_date = formatDateString(tsnddt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t2d_seismic_navigation_digital_data_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, seis_dimension, process_date, shot_by, line_name, digital_format, media_type, original_file_name, password, digital_size, digital_size_uom, ba_long_name_2, ba_type_2, data_store_name, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22) RETURNING id INTO :23`, &tsnddt.Ba_long_name, &tsnddt.Ba_type, &tsnddt.Area_id, &tsnddt.Area_type, &tsnddt.Acqtn_survey_name, &tsnddt.Seis_dimension, &tsnddt.Process_date, &tsnddt.Shot_by, &tsnddt.Line_name, &tsnddt.Digital_format, &tsnddt.Media_type, &tsnddt.Original_file_name, &tsnddt.Password, &tsnddt.Digital_size, &tsnddt.Digital_size_uom, &tsnddt.Ba_long_name_2, &tsnddt.Ba_type_2, &tsnddt.Data_store_name, &tsnddt.Remark, &tsnddt.Source, &tsnddt.Qc_status, &tsnddt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T2D_SEISMIC_NAVIGATION_DIGITAL_DATA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT2DSeismicNavigationDigitalDataById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t2d_seismic_navigation_digital_data_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T2D_seismic_navigation_digital_data    
    
        for rows.Next() {
    
            var tsnddt dto.T2D_seismic_navigation_digital_data
            if err := rows.Scan(&tsnddt.Id, &tsnddt.Ba_long_name, &tsnddt.Ba_type, &tsnddt.Area_id, &tsnddt.Area_type, &tsnddt.Acqtn_survey_name, &tsnddt.Seis_dimension, &tsnddt.Process_date, &tsnddt.Shot_by, &tsnddt.Line_name, &tsnddt.Digital_format, &tsnddt.Media_type, &tsnddt.Original_file_name, &tsnddt.Password, &tsnddt.Digital_size, &tsnddt.Digital_size_uom, &tsnddt.Ba_long_name_2, &tsnddt.Ba_type_2, &tsnddt.Data_store_name, &tsnddt.Remark, &tsnddt.Source, &tsnddt.Qc_status, &tsnddt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsnddt)
        }
    
    return c.JSON(results)
}
func PutT2DSeismicNavigationDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsnddt dto.T2D_seismic_navigation_digital_data
    setDefaults(&tsnddt)

    if err := c.BodyParser(&tsnddt); err != nil{
        return err
    }
    
    tsnddt.Create_date = formatDateString(tsnddt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t2d_seismic_navigation_digital_data_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, seis_dimension = :6, process_date = :7, shot_by = :8, line_name = :9, digital_format = :10, media_type = :11, original_file_name = :12, password = :13, digital_size = :14, digital_size_uom = :15, ba_long_name_2 = :16, ba_type_2 = :17, data_store_name = :18, remark = :19, source = :20, qc_status = :21, checked_by_ba_id = :22 WHERE id = :23`, &tsnddt.Ba_long_name, &tsnddt.Ba_type, &tsnddt.Area_id, &tsnddt.Area_type, &tsnddt.Acqtn_survey_name, &tsnddt.Seis_dimension, &tsnddt.Process_date, &tsnddt.Shot_by, &tsnddt.Line_name, &tsnddt.Digital_format, &tsnddt.Media_type, &tsnddt.Original_file_name, &tsnddt.Password, &tsnddt.Digital_size, &tsnddt.Digital_size_uom, &tsnddt.Ba_long_name_2, &tsnddt.Ba_type_2, &tsnddt.Data_store_name, &tsnddt.Remark, &tsnddt.Source, &tsnddt.Qc_status, &tsnddt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T2D_SEISMIC_NAVIGATION_DIGITAL_DATA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT2DSeismicNavigationDigitalData(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t2d_seismic_navigation_digital_data_id FROM t2d_seismic_navigation_digital_data_workspace WHERE t2d_seismic_navigation_digital_data_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t2d_seismic_navigation_digital_data_workspace WHERE t2d_seismic_navigation_digital_data_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t2d_seismic_navigation_digital_data_table WHERE id = :1`, id)
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
func PatchT2DSeismicNavigationDigitalData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsnddt dto.T2D_seismic_navigation_digital_data
    setDefaults(&tsnddt)

    if err := c.BodyParser(&tsnddt); err != nil{
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
    err = db.QueryRow(`SELECT t2d_seismic_navigation_digital_data_id FROM t2d_seismic_navigation_digital_data_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tsnddt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET ba_long_name = :1 WHERE id = :2`, tsnddt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET ba_type = :1 WHERE id = :2`, tsnddt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Area_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET area_id = :1 WHERE id = :2`, tsnddt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Area_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET area_type = :1 WHERE id = :2`, tsnddt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET acqtn_survey_name = :1 WHERE id = :2`, tsnddt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Seis_dimension != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET seis_dimension = :1 WHERE id = :2`, tsnddt.Seis_dimension, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Process_date != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET process_date = :1 WHERE id = :2`, tsnddt.Process_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Shot_by != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET shot_by = :1 WHERE id = :2`, tsnddt.Shot_by, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Line_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET line_name = :1 WHERE id = :2`, tsnddt.Line_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET digital_format = :1 WHERE id = :2`, tsnddt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Media_type != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET media_type = :1 WHERE id = :2`, tsnddt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET original_file_name = :1 WHERE id = :2`, tsnddt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Password != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET password = :1 WHERE id = :2`, tsnddt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET digital_size = :1 WHERE id = :2`, tsnddt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET digital_size_uom = :1 WHERE id = :2`, tsnddt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET ba_long_name_2 = :1 WHERE id = :2`, tsnddt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET ba_type_2 = :1 WHERE id = :2`, tsnddt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET data_store_name = :1 WHERE id = :2`, tsnddt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Remark != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET remark = :1 WHERE id = :2`, tsnddt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Source != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET source = :1 WHERE id = :2`, tsnddt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET qc_status = :1 WHERE id = :2`, tsnddt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsnddt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE t2d_seismic_navigation_digital_data_table SET checked_by_ba_id = :1 WHERE id = :2`, tsnddt.Checked_by_ba_id, id)
        
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

    