package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-workingarea/dto"
    "github.com/DarrenMannuela/svc-datatype-workingarea/utils"

    "github.com/gofiber/fiber/v2"
)
func GetWorkingArea(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM working_area_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Working_area    
    
        for rows.Next() {
    
            var wat dto.Working_area
            if err := rows.Scan(&wat.Id, &wat.Area_id, &wat.Area_type, &wat.Ba_long_name, &wat.Ba_type, &wat.Effective_date, &wat.Term_exiry_date, &wat.Contract_type, &wat.R_granted_right_type, &wat.Gross_size, &wat.Gross_size_ouom, &wat.Land_right_category, &wat.Producing_ind, &wat.Original_file_name, &wat.Password, &wat.Digital_size, &wat.Digital_size_uom, &wat.Media_type, &wat.Data_store_name, &wat.Source, &wat.Qc_status, &wat.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wat)
        }
    
    return c.JSON(results)
}
func SetWorkingArea(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var wat dto.Working_area
    setDefaults(&wat)

    if err := c.BodyParser(&wat); err != nil{
        return err
    }
    
    wat.Create_date = formatDateString(wat.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO working_area_table (area_id, area_type, ba_long_name, ba_type, effective_date, term_exiry_date, contract_type, r_granted_right_type, gross_size, gross_size_ouom, land_right_category, producing_ind, original_file_name, password, digital_size, digital_size_uom, media_type, data_store_name, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21) RETURNING id INTO :22`, &wat.Area_id, &wat.Area_type, &wat.Ba_long_name, &wat.Ba_type, &wat.Effective_date, &wat.Term_exiry_date, &wat.Contract_type, &wat.R_granted_right_type, &wat.Gross_size, &wat.Gross_size_ouom, &wat.Land_right_category, &wat.Producing_ind, &wat.Original_file_name, &wat.Password, &wat.Digital_size, &wat.Digital_size_uom, &wat.Media_type, &wat.Data_store_name, &wat.Source, &wat.Qc_status, &wat.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(WORKING_AREA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetWorkingAreaById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM working_area_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Working_area    
    
        for rows.Next() {
    
            var wat dto.Working_area
            if err := rows.Scan(&wat.Id, &wat.Area_id, &wat.Area_type, &wat.Ba_long_name, &wat.Ba_type, &wat.Effective_date, &wat.Term_exiry_date, &wat.Contract_type, &wat.R_granted_right_type, &wat.Gross_size, &wat.Gross_size_ouom, &wat.Land_right_category, &wat.Producing_ind, &wat.Original_file_name, &wat.Password, &wat.Digital_size, &wat.Digital_size_uom, &wat.Media_type, &wat.Data_store_name, &wat.Source, &wat.Qc_status, &wat.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wat)
        }
    
    return c.JSON(results)
}
func PutWorkingArea(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wat dto.Working_area
    setDefaults(&wat)

    if err := c.BodyParser(&wat); err != nil{
        return err
    }
    
    wat.Create_date = formatDateString(wat.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM working_area_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE working_area_table SET
        area_id = :1, area_type = :2, ba_long_name = :3, ba_type = :4, effective_date = :5, term_exiry_date = :6, contract_type = :7, r_granted_right_type = :8, gross_size = :9, gross_size_ouom = :10, land_right_category = :11, producing_ind = :12, original_file_name = :13, password = :14, digital_size = :15, digital_size_uom = :16, media_type = :17, data_store_name = :18, source = :19, qc_status = :20, checked_by_ba_id = :21 WHERE id = :22`, &wat.Area_id, &wat.Area_type, &wat.Ba_long_name, &wat.Ba_type, &wat.Effective_date, &wat.Term_exiry_date, &wat.Contract_type, &wat.R_granted_right_type, &wat.Gross_size, &wat.Gross_size_ouom, &wat.Land_right_category, &wat.Producing_ind, &wat.Original_file_name, &wat.Password, &wat.Digital_size, &wat.Digital_size_uom, &wat.Media_type, &wat.Data_store_name, &wat.Source, &wat.Qc_status, &wat.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(WORKING_AREA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteWorkingArea(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT working_area_id FROM working_area_workspace WHERE working_area_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM working_area_workspace WHERE working_area_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM working_area_table WHERE id = :1`, id)
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
func PatchWorkingArea(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wat dto.Working_area
    setDefaults(&wat)

    if err := c.BodyParser(&wat); err != nil{
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
    err = db.QueryRow(`SELECT working_area_id FROM working_area_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if wat.Area_id != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET area_id = :1 WHERE id = :2`, wat.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Area_type != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET area_type = :1 WHERE id = :2`, wat.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET ba_long_name = :1 WHERE id = :2`, wat.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Ba_type != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET ba_type = :1 WHERE id = :2`, wat.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Effective_date != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET effective_date = :1 WHERE id = :2`, wat.Effective_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Term_exiry_date != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET term_exiry_date = :1 WHERE id = :2`, wat.Term_exiry_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Contract_type != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET contract_type = :1 WHERE id = :2`, wat.Contract_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.R_granted_right_type != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET r_granted_right_type = :1 WHERE id = :2`, wat.R_granted_right_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Gross_size != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET gross_size = :1 WHERE id = :2`, wat.Gross_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Gross_size_ouom != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET gross_size_ouom = :1 WHERE id = :2`, wat.Gross_size_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Land_right_category != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET land_right_category = :1 WHERE id = :2`, wat.Land_right_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Producing_ind != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET producing_ind = :1 WHERE id = :2`, wat.Producing_ind, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET original_file_name = :1 WHERE id = :2`, wat.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Password != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET password = :1 WHERE id = :2`, wat.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Digital_size != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET digital_size = :1 WHERE id = :2`, wat.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET digital_size_uom = :1 WHERE id = :2`, wat.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Media_type != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET media_type = :1 WHERE id = :2`, wat.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET data_store_name = :1 WHERE id = :2`, wat.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Source != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET source = :1 WHERE id = :2`, wat.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Qc_status != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET qc_status = :1 WHERE id = :2`, wat.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wat.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE working_area_table SET checked_by_ba_id = :1 WHERE id = :2`, wat.Checked_by_ba_id, id)
        
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

    