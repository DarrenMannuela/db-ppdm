package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-digitalmapsandtechnicaldrawing/dto"
    "github.com/DarrenMannuela/svc-datatype-digitalmapsandtechnicaldrawing/utils"

    "github.com/gofiber/fiber/v2"
)
func GetDigitalMapsAndTechnicalDrawing(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM digital_maps_and_technical_drawing_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_maps_and_technical_drawing    
    
        for rows.Next() {
    
            var dmatdt dto.Digital_maps_and_technical_drawing
            if err := rows.Scan(&dmatdt.Id, &dmatdt.Ba_long_name, &dmatdt.Ba_type, &dmatdt.Area_id, &dmatdt.Area_type, &dmatdt.Title, &dmatdt.Creator_name, &dmatdt.Create_date, &dmatdt.Digital_format, &dmatdt.Media_type, &dmatdt.Map_scale, &dmatdt.Projection_type, &dmatdt.Geodetic_datum_name, &dmatdt.Document_type, &dmatdt.Item_category, &dmatdt.Ba_long_name_2, &dmatdt.Ba_type_2, &dmatdt.Data_store_name, &dmatdt.Original_file_name, &dmatdt.Password, &dmatdt.Digital_size, &dmatdt.Digital_size_uom, &dmatdt.Remark, &dmatdt.Source, &dmatdt.Qc_status, &dmatdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dmatdt)
        }
    
    return c.JSON(results)
}
func SetDigitalMapsAndTechnicalDrawing(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var dmatdt dto.Digital_maps_and_technical_drawing
    setDefaults(&dmatdt)

    if err := c.BodyParser(&dmatdt); err != nil{
        return err
    }
    
    dmatdt.Create_date = formatDateString(dmatdt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO digital_maps_and_technical_drawing_table (ba_long_name, ba_type, area_id, area_type, title, creator_name, create_date, digital_format, media_type, map_scale, projection_type, geodetic_datum_name, document_type, item_category, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25) RETURNING id INTO :26`, &dmatdt.Ba_long_name, &dmatdt.Ba_type, &dmatdt.Area_id, &dmatdt.Area_type, &dmatdt.Title, &dmatdt.Creator_name, &dmatdt.Create_date, &dmatdt.Digital_format, &dmatdt.Media_type, &dmatdt.Map_scale, &dmatdt.Projection_type, &dmatdt.Geodetic_datum_name, &dmatdt.Document_type, &dmatdt.Item_category, &dmatdt.Ba_long_name_2, &dmatdt.Ba_type_2, &dmatdt.Data_store_name, &dmatdt.Original_file_name, &dmatdt.Password, &dmatdt.Digital_size, &dmatdt.Digital_size_uom, &dmatdt.Remark, &dmatdt.Source, &dmatdt.Qc_status, &dmatdt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(DIGITAL_MAPS_AND_TECHNICAL_DRAWING_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetDigitalMapsAndTechnicalDrawingById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM digital_maps_and_technical_drawing_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_maps_and_technical_drawing    
    
        for rows.Next() {
    
            var dmatdt dto.Digital_maps_and_technical_drawing
            if err := rows.Scan(&dmatdt.Id, &dmatdt.Ba_long_name, &dmatdt.Ba_type, &dmatdt.Area_id, &dmatdt.Area_type, &dmatdt.Title, &dmatdt.Creator_name, &dmatdt.Create_date, &dmatdt.Digital_format, &dmatdt.Media_type, &dmatdt.Map_scale, &dmatdt.Projection_type, &dmatdt.Geodetic_datum_name, &dmatdt.Document_type, &dmatdt.Item_category, &dmatdt.Ba_long_name_2, &dmatdt.Ba_type_2, &dmatdt.Data_store_name, &dmatdt.Original_file_name, &dmatdt.Password, &dmatdt.Digital_size, &dmatdt.Digital_size_uom, &dmatdt.Remark, &dmatdt.Source, &dmatdt.Qc_status, &dmatdt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dmatdt)
        }
    
    return c.JSON(results)
}
func PutDigitalMapsAndTechnicalDrawing(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dmatdt dto.Digital_maps_and_technical_drawing
    setDefaults(&dmatdt)

    if err := c.BodyParser(&dmatdt); err != nil{
        return err
    }
    
    dmatdt.Create_date = formatDateString(dmatdt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM digital_maps_and_technical_drawing_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE digital_maps_and_technical_drawing_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, title = :5, creator_name = :6, create_date = :7, digital_format = :8, media_type = :9, map_scale = :10, projection_type = :11, geodetic_datum_name = :12, document_type = :13, item_category = :14, ba_long_name_2 = :15, ba_type_2 = :16, data_store_name = :17, original_file_name = :18, password = :19, digital_size = :20, digital_size_uom = :21, remark = :22, source = :23, qc_status = :24, checked_by_ba_id = :25 WHERE id = :26`, &dmatdt.Ba_long_name, &dmatdt.Ba_type, &dmatdt.Area_id, &dmatdt.Area_type, &dmatdt.Title, &dmatdt.Creator_name, &dmatdt.Create_date, &dmatdt.Digital_format, &dmatdt.Media_type, &dmatdt.Map_scale, &dmatdt.Projection_type, &dmatdt.Geodetic_datum_name, &dmatdt.Document_type, &dmatdt.Item_category, &dmatdt.Ba_long_name_2, &dmatdt.Ba_type_2, &dmatdt.Data_store_name, &dmatdt.Original_file_name, &dmatdt.Password, &dmatdt.Digital_size, &dmatdt.Digital_size_uom, &dmatdt.Remark, &dmatdt.Source, &dmatdt.Qc_status, &dmatdt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(DIGITAL_MAPS_AND_TECHNICAL_DRAWING_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteDigitalMapsAndTechnicalDrawing(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT digital_maps_and_technical_drawing_id FROM digital_maps_and_technical_drawing_workspace WHERE digital_maps_and_technical_drawing_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM digital_maps_and_technical_drawing_workspace WHERE digital_maps_and_technical_drawing_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM digital_maps_and_technical_drawing_table WHERE id = :1`, id)
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
func PatchDigitalMapsAndTechnicalDrawing(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dmatdt dto.Digital_maps_and_technical_drawing
    setDefaults(&dmatdt)

    if err := c.BodyParser(&dmatdt); err != nil{
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
    err = db.QueryRow(`SELECT digital_maps_and_technical_drawing_id FROM digital_maps_and_technical_drawing_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if dmatdt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, dmatdt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, dmatdt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, dmatdt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, dmatdt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Title != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET title = :1 WHERE id = :2`, dmatdt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET creator_name = :1 WHERE id = :2`, dmatdt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, dmatdt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, dmatdt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, dmatdt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Map_scale != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET map_scale = :1 WHERE id = :2`, dmatdt.Map_scale, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Projection_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET projection_type = :1 WHERE id = :2`, dmatdt.Projection_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Geodetic_datum_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET geodetic_datum_name = :1 WHERE id = :2`, dmatdt.Geodetic_datum_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Document_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET document_type = :1 WHERE id = :2`, dmatdt.Document_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Item_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_category = :1 WHERE id = :2`, dmatdt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, dmatdt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, dmatdt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, dmatdt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, dmatdt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, dmatdt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, dmatdt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, dmatdt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, dmatdt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, dmatdt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, dmatdt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dmatdt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, dmatdt.Checked_by_ba_id, id)
        
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

    