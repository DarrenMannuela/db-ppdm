package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-seismicinterpretationdata/dto"
    "github.com/DarrenMannuela/svc-datatype-seismicinterpretationdata/utils"

    "github.com/gofiber/fiber/v2"
)
func GetSeismicInterpretationData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM seismic_interpretation_data_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Seismic_interpretation_data    
    
        for rows.Next() {
    
            var sidt dto.Seismic_interpretation_data
            if err := rows.Scan(&sidt.Id, &sidt.Ba_long_name, &sidt.Ba_type, &sidt.Area_id, &sidt.Area_type, &sidt.Field_name, &sidt.Project_name, &sidt.Interpreter, &sidt.Interp_date, &sidt.Interp_objective, &sidt.Interp_type, &sidt.Pick_method, &sidt.Sw_application_id, &sidt.Application_version, &sidt.Trace_position, &sidt.Media_type, &sidt.Tape_number, &sidt.Digital_format, &sidt.Remark, &sidt.Ba_long_name_2, &sidt.Ba_type_2, &sidt.Data_store_name, &sidt.Original_file_name, &sidt.Password, &sidt.Digital_size, &sidt.Digital_size_uom, &sidt.Source, &sidt.Qc_status, &sidt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, sidt)
        }
    
    return c.JSON(results)
}
func SetSeismicInterpretationData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var sidt dto.Seismic_interpretation_data
    setDefaults(&sidt)

    if err := c.BodyParser(&sidt); err != nil{
        return err
    }
    
    sidt.Create_date = formatDateString(sidt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO seismic_interpretation_data_table (ba_long_name, ba_type, area_id, area_type, field_name, project_name, interpreter, interp_date, interp_objective, interp_type, pick_method, sw_application_id, application_version, trace_position, media_type, tape_number, digital_format, remark, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, digital_size, digital_size_uom, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28) RETURNING id INTO :29`, &sidt.Ba_long_name, &sidt.Ba_type, &sidt.Area_id, &sidt.Area_type, &sidt.Field_name, &sidt.Project_name, &sidt.Interpreter, &sidt.Interp_date, &sidt.Interp_objective, &sidt.Interp_type, &sidt.Pick_method, &sidt.Sw_application_id, &sidt.Application_version, &sidt.Trace_position, &sidt.Media_type, &sidt.Tape_number, &sidt.Digital_format, &sidt.Remark, &sidt.Ba_long_name_2, &sidt.Ba_type_2, &sidt.Data_store_name, &sidt.Original_file_name, &sidt.Password, &sidt.Digital_size, &sidt.Digital_size_uom, &sidt.Source, &sidt.Qc_status, &sidt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(SEISMIC_INTERPRETATION_DATA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetSeismicInterpretationDataById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM seismic_interpretation_data_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Seismic_interpretation_data    
    
        for rows.Next() {
    
            var sidt dto.Seismic_interpretation_data
            if err := rows.Scan(&sidt.Id, &sidt.Ba_long_name, &sidt.Ba_type, &sidt.Area_id, &sidt.Area_type, &sidt.Field_name, &sidt.Project_name, &sidt.Interpreter, &sidt.Interp_date, &sidt.Interp_objective, &sidt.Interp_type, &sidt.Pick_method, &sidt.Sw_application_id, &sidt.Application_version, &sidt.Trace_position, &sidt.Media_type, &sidt.Tape_number, &sidt.Digital_format, &sidt.Remark, &sidt.Ba_long_name_2, &sidt.Ba_type_2, &sidt.Data_store_name, &sidt.Original_file_name, &sidt.Password, &sidt.Digital_size, &sidt.Digital_size_uom, &sidt.Source, &sidt.Qc_status, &sidt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, sidt)
        }
    
    return c.JSON(results)
}
func PutSeismicInterpretationData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var sidt dto.Seismic_interpretation_data
    setDefaults(&sidt)

    if err := c.BodyParser(&sidt); err != nil{
        return err
    }
    
    sidt.Create_date = formatDateString(sidt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM seismic_interpretation_data_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, project_name = :6, interpreter = :7, interp_date = :8, interp_objective = :9, interp_type = :10, pick_method = :11, sw_application_id = :12, application_version = :13, trace_position = :14, media_type = :15, tape_number = :16, digital_format = :17, remark = :18, ba_long_name_2 = :19, ba_type_2 = :20, data_store_name = :21, original_file_name = :22, password = :23, digital_size = :24, digital_size_uom = :25, source = :26, qc_status = :27, checked_by_ba_id = :28 WHERE id = :29`, &sidt.Ba_long_name, &sidt.Ba_type, &sidt.Area_id, &sidt.Area_type, &sidt.Field_name, &sidt.Project_name, &sidt.Interpreter, &sidt.Interp_date, &sidt.Interp_objective, &sidt.Interp_type, &sidt.Pick_method, &sidt.Sw_application_id, &sidt.Application_version, &sidt.Trace_position, &sidt.Media_type, &sidt.Tape_number, &sidt.Digital_format, &sidt.Remark, &sidt.Ba_long_name_2, &sidt.Ba_type_2, &sidt.Data_store_name, &sidt.Original_file_name, &sidt.Password, &sidt.Digital_size, &sidt.Digital_size_uom, &sidt.Source, &sidt.Qc_status, &sidt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(SEISMIC_INTERPRETATION_DATA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteSeismicInterpretationData(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT seismic_interpretation_data_id FROM seismic_interpretation_data_workspace WHERE seismic_interpretation_data_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM seismic_interpretation_data_workspace WHERE seismic_interpretation_data_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM seismic_interpretation_data_table WHERE id = :1`, id)
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
func PatchSeismicInterpretationData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var sidt dto.Seismic_interpretation_data
    setDefaults(&sidt)

    if err := c.BodyParser(&sidt); err != nil{
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
    err = db.QueryRow(`SELECT seismic_interpretation_data_id FROM seismic_interpretation_data_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if sidt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET ba_long_name = :1 WHERE id = :2`, sidt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET ba_type = :1 WHERE id = :2`, sidt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Area_id != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET area_id = :1 WHERE id = :2`, sidt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Area_type != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET area_type = :1 WHERE id = :2`, sidt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Field_name != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET field_name = :1 WHERE id = :2`, sidt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Project_name != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET project_name = :1 WHERE id = :2`, sidt.Project_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Interpreter != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET interpreter = :1 WHERE id = :2`, sidt.Interpreter, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Interp_date != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET interp_date = :1 WHERE id = :2`, sidt.Interp_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Interp_objective != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET interp_objective = :1 WHERE id = :2`, sidt.Interp_objective, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Interp_type != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET interp_type = :1 WHERE id = :2`, sidt.Interp_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Pick_method != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET pick_method = :1 WHERE id = :2`, sidt.Pick_method, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Sw_application_id != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET sw_application_id = :1 WHERE id = :2`, sidt.Sw_application_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Application_version != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET application_version = :1 WHERE id = :2`, sidt.Application_version, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Trace_position != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET trace_position = :1 WHERE id = :2`, sidt.Trace_position, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Media_type != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET media_type = :1 WHERE id = :2`, sidt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Tape_number != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET tape_number = :1 WHERE id = :2`, sidt.Tape_number, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET digital_format = :1 WHERE id = :2`, sidt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Remark != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET remark = :1 WHERE id = :2`, sidt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET ba_long_name_2 = :1 WHERE id = :2`, sidt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET ba_type_2 = :1 WHERE id = :2`, sidt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET data_store_name = :1 WHERE id = :2`, sidt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET original_file_name = :1 WHERE id = :2`, sidt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Password != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET password = :1 WHERE id = :2`, sidt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET digital_size = :1 WHERE id = :2`, sidt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET digital_size_uom = :1 WHERE id = :2`, sidt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Source != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET source = :1 WHERE id = :2`, sidt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET qc_status = :1 WHERE id = :2`, sidt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if sidt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE seismic_interpretation_data_table SET checked_by_ba_id = :1 WHERE id = :2`, sidt.Checked_by_ba_id, id)
        
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

    