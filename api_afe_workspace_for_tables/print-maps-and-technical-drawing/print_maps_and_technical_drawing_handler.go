package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-printmapsandtechnicaldrawing/dto"
    "github.com/DarrenMannuela/svc-datatype-printmapsandtechnicaldrawing/utils"

    "github.com/gofiber/fiber/v2"
)
func GetPrintMapsAndTechnicalDrawing(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM print_maps_and_technical_drawing_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Print_maps_and_technical_drawing    
    
        for rows.Next() {
    
            var pmatdt dto.Print_maps_and_technical_drawing
            if err := rows.Scan(&pmatdt.Id, &pmatdt.Ba_long_name, &pmatdt.Ba_type, &pmatdt.Area_id, &pmatdt.Area_type, &pmatdt.Title, &pmatdt.Creator_name, &pmatdt.Create_date, &pmatdt.Map_scale, &pmatdt.Projection_type, &pmatdt.Geodetic_datum_name, &pmatdt.Media_type, &pmatdt.Document_type, &pmatdt.Item_category, &pmatdt.Ba_long_name_2, &pmatdt.Ba_type_2, &pmatdt.Data_store_name, &pmatdt.Data_store_type, &pmatdt.Location_id, &pmatdt.Remark, &pmatdt.Source, &pmatdt.Qc_status, &pmatdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, pmatdt)
        }
    
    return c.JSON(results)
}
func SetPrintMapsAndTechnicalDrawing(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var pmatdt dto.Print_maps_and_technical_drawing
    setDefaults(&pmatdt)

    if err := c.BodyParser(&pmatdt); err != nil{
        return err
    }
    
    pmatdt.Create_date = formatDateString(pmatdt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO print_maps_and_technical_drawing_table (ba_long_name, ba_type, area_id, area_type, title, creator_name, create_date, map_scale, projection_type, geodetic_datum_name, media_type, document_type, item_category, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22) RETURNING id INTO :23`, &pmatdt.Ba_long_name, &pmatdt.Ba_type, &pmatdt.Area_id, &pmatdt.Area_type, &pmatdt.Title, &pmatdt.Creator_name, &pmatdt.Create_date, &pmatdt.Map_scale, &pmatdt.Projection_type, &pmatdt.Geodetic_datum_name, &pmatdt.Media_type, &pmatdt.Document_type, &pmatdt.Item_category, &pmatdt.Ba_long_name_2, &pmatdt.Ba_type_2, &pmatdt.Data_store_name, &pmatdt.Data_store_type, &pmatdt.Location_id, &pmatdt.Remark, &pmatdt.Source, &pmatdt.Qc_status, &pmatdt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(PRINT_MAPS_AND_TECHNICAL_DRAWING_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetPrintMapsAndTechnicalDrawingById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM print_maps_and_technical_drawing_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Print_maps_and_technical_drawing    
    
        for rows.Next() {
    
            var pmatdt dto.Print_maps_and_technical_drawing
            if err := rows.Scan(&pmatdt.Id, &pmatdt.Ba_long_name, &pmatdt.Ba_type, &pmatdt.Area_id, &pmatdt.Area_type, &pmatdt.Title, &pmatdt.Creator_name, &pmatdt.Create_date, &pmatdt.Map_scale, &pmatdt.Projection_type, &pmatdt.Geodetic_datum_name, &pmatdt.Media_type, &pmatdt.Document_type, &pmatdt.Item_category, &pmatdt.Ba_long_name_2, &pmatdt.Ba_type_2, &pmatdt.Data_store_name, &pmatdt.Data_store_type, &pmatdt.Location_id, &pmatdt.Remark, &pmatdt.Source, &pmatdt.Qc_status, &pmatdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, pmatdt)
        }
    
    return c.JSON(results)
}
func PutPrintMapsAndTechnicalDrawing(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var pmatdt dto.Print_maps_and_technical_drawing
    setDefaults(&pmatdt)

    if err := c.BodyParser(&pmatdt); err != nil{
        return err
    }
    
    pmatdt.Create_date = formatDateString(pmatdt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM print_maps_and_technical_drawing_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, title = :5, creator_name = :6, create_date = :7, map_scale = :8, projection_type = :9, geodetic_datum_name = :10, media_type = :11, document_type = :12, item_category = :13, ba_long_name_2 = :14, ba_type_2 = :15, data_store_name = :16, data_store_type = :17, location_id = :18, remark = :19, source = :20, qc_status = :21, checked_by_ba_id = :22 WHERE id = :23`, &pmatdt.Ba_long_name, &pmatdt.Ba_type, &pmatdt.Area_id, &pmatdt.Area_type, &pmatdt.Title, &pmatdt.Creator_name, &pmatdt.Create_date, &pmatdt.Map_scale, &pmatdt.Projection_type, &pmatdt.Geodetic_datum_name, &pmatdt.Media_type, &pmatdt.Document_type, &pmatdt.Item_category, &pmatdt.Ba_long_name_2, &pmatdt.Ba_type_2, &pmatdt.Data_store_name, &pmatdt.Data_store_type, &pmatdt.Location_id, &pmatdt.Remark, &pmatdt.Source, &pmatdt.Qc_status, &pmatdt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(PRINT_MAPS_AND_TECHNICAL_DRAWING_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeletePrintMapsAndTechnicalDrawing(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT print_maps_and_technical_drawing_id FROM print_maps_and_technical_drawing_workspace WHERE print_maps_and_technical_drawing_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM print_maps_and_technical_drawing_workspace WHERE print_maps_and_technical_drawing_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM print_maps_and_technical_drawing_table WHERE id = :1`, id)
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
func PatchPrintMapsAndTechnicalDrawing(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var pmatdt dto.Print_maps_and_technical_drawing
    setDefaults(&pmatdt)

    if err := c.BodyParser(&pmatdt); err != nil{
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
    err = db.QueryRow(`SELECT print_maps_and_technical_drawing_id FROM print_maps_and_technical_drawing_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if pmatdt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET ba_long_name = :1 WHERE id = :2`, pmatdt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET ba_type = :1 WHERE id = :2`, pmatdt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Area_id != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET area_id = :1 WHERE id = :2`, pmatdt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Area_type != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET area_type = :1 WHERE id = :2`, pmatdt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Title != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET title = :1 WHERE id = :2`, pmatdt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET creator_name = :1 WHERE id = :2`, pmatdt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Create_date != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET create_date = :1 WHERE id = :2`, pmatdt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Map_scale != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET map_scale = :1 WHERE id = :2`, pmatdt.Map_scale, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Projection_type != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET projection_type = :1 WHERE id = :2`, pmatdt.Projection_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Geodetic_datum_name != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET geodetic_datum_name = :1 WHERE id = :2`, pmatdt.Geodetic_datum_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Media_type != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET media_type = :1 WHERE id = :2`, pmatdt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Document_type != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET document_type = :1 WHERE id = :2`, pmatdt.Document_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Item_category != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET item_category = :1 WHERE id = :2`, pmatdt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET ba_long_name_2 = :1 WHERE id = :2`, pmatdt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET ba_type_2 = :1 WHERE id = :2`, pmatdt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET data_store_name = :1 WHERE id = :2`, pmatdt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET data_store_type = :1 WHERE id = :2`, pmatdt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Location_id != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET location_id = :1 WHERE id = :2`, pmatdt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Remark != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET remark = :1 WHERE id = :2`, pmatdt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Source != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET source = :1 WHERE id = :2`, pmatdt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET qc_status = :1 WHERE id = :2`, pmatdt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pmatdt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE print_maps_and_technical_drawing_table SET checked_by_ba_id = :1 WHERE id = :2`, pmatdt.Checked_by_ba_id, id)
        
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

    