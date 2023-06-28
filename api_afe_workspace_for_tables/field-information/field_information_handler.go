package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-fieldinformation/dto"
    "github.com/DarrenMannuela/svc-datatype-fieldinformation/utils"

    "github.com/gofiber/fiber/v2"
)
func GetFieldInformation(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM field_information_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Field_information    
    
        for rows.Next() {
    
            var fit dto.Field_information
            if err := rows.Scan(&fit.Id, &fit.Area_id, &fit.Area_type, &fit.Field_name, &fit.Discovery_date, &fit.Field_type, &fit.Original_file_name, &fit.Password, &fit.Digital_size, &fit.Digital_size_uom, &fit.Media_type, &fit.Data_store_name, &fit.Remark, &fit.Qc_status, &fit.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, fit)
        }
    
    return c.JSON(results)
}
func SetFieldInformation(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var fit dto.Field_information
    setDefaults(&fit)

    if err := c.BodyParser(&fit); err != nil{
        return err
    }
    
    fit.Create_date = formatDateString(fit.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO field_information_table (area_id, area_type, field_name, discovery_date, field_type, original_file_name, password, digital_size, digital_size_uom, media_type, data_store_name, remark, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14) RETURNING id INTO :15`, &fit.Area_id, &fit.Area_type, &fit.Field_name, &fit.Discovery_date, &fit.Field_type, &fit.Original_file_name, &fit.Password, &fit.Digital_size, &fit.Digital_size_uom, &fit.Media_type, &fit.Data_store_name, &fit.Remark, &fit.Qc_status, &fit.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(FIELD_INFORMATION_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetFieldInformationById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM field_information_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Field_information    
    
        for rows.Next() {
    
            var fit dto.Field_information
            if err := rows.Scan(&fit.Id, &fit.Area_id, &fit.Area_type, &fit.Field_name, &fit.Discovery_date, &fit.Field_type, &fit.Original_file_name, &fit.Password, &fit.Digital_size, &fit.Digital_size_uom, &fit.Media_type, &fit.Data_store_name, &fit.Remark, &fit.Qc_status, &fit.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, fit)
        }
    
    return c.JSON(results)
}
func PutFieldInformation(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var fit dto.Field_information
    setDefaults(&fit)

    if err := c.BodyParser(&fit); err != nil{
        return err
    }
    
    fit.Create_date = formatDateString(fit.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM field_information_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE field_information_table SET
        area_id = :1, area_type = :2, field_name = :3, discovery_date = :4, field_type = :5, original_file_name = :6, password = :7, digital_size = :8, digital_size_uom = :9, media_type = :10, data_store_name = :11, remark = :12, qc_status = :13, checked_by_ba_id = :14 WHERE id = :15`, &fit.Area_id, &fit.Area_type, &fit.Field_name, &fit.Discovery_date, &fit.Field_type, &fit.Original_file_name, &fit.Password, &fit.Digital_size, &fit.Digital_size_uom, &fit.Media_type, &fit.Data_store_name, &fit.Remark, &fit.Qc_status, &fit.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(FIELD_INFORMATION_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteFieldInformation(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT field_information_id FROM field_information_workspace WHERE field_information_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM field_information_workspace WHERE field_information_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM field_information_table WHERE id = :1`, id)
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
func PatchFieldInformation(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var fit dto.Field_information
    setDefaults(&fit)

    if err := c.BodyParser(&fit); err != nil{
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
    err = db.QueryRow(`SELECT field_information_id FROM field_information_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if fit.Area_id != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET area_id = :1 WHERE id = :2`, fit.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Area_type != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET area_type = :1 WHERE id = :2`, fit.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Field_name != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET field_name = :1 WHERE id = :2`, fit.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Discovery_date != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET discovery_date = :1 WHERE id = :2`, fit.Discovery_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Field_type != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET field_type = :1 WHERE id = :2`, fit.Field_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET original_file_name = :1 WHERE id = :2`, fit.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Password != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET password = :1 WHERE id = :2`, fit.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Digital_size != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET digital_size = :1 WHERE id = :2`, fit.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET digital_size_uom = :1 WHERE id = :2`, fit.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Media_type != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET media_type = :1 WHERE id = :2`, fit.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET data_store_name = :1 WHERE id = :2`, fit.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Remark != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET remark = :1 WHERE id = :2`, fit.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Qc_status != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET qc_status = :1 WHERE id = :2`, fit.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if fit.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE field_information_table SET checked_by_ba_id = :1 WHERE id = :2`, fit.Checked_by_ba_id, id)
        
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

    