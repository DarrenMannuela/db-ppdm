package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-digitaltechnicalreport/dto"
    "github.com/DarrenMannuela/svc-datatype-digitaltechnicalreport/utils"

    "github.com/gofiber/fiber/v2"
)
func GetDigitalTechnicalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM digital_technical_report_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_technical_report    
    
        for rows.Next() {
    
            var dtrt dto.Digital_technical_report
            if err := rows.Scan(&dtrt.Id, &dtrt.Ba_long_name, &dtrt.Ba_type, &dtrt.Area_id, &dtrt.Area_type, &dtrt.Title, &dtrt.Creator_name, &dtrt.Create_date, &dtrt.Page_count, &dtrt.Media_type, &dtrt.Digital_format, &dtrt.Document_type, &dtrt.Item_category, &dtrt.Item_sub_category, &dtrt.Ba_long_name_2, &dtrt.Ba_type_2, &dtrt.Data_store_name, &dtrt.Original_file_name, &dtrt.Password, &dtrt.Digital_size, &dtrt.Digital_size_uom, &dtrt.Remark, &dtrt.Source, &dtrt.Qc_status, &dtrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dtrt)
        }
    
    return c.JSON(results)
}
func SetDigitalTechnicalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var dtrt dto.Digital_technical_report
    setDefaults(&dtrt)

    if err := c.BodyParser(&dtrt); err != nil{
        return err
    }
    
    dtrt.Create_date = formatDateString(dtrt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO digital_technical_report_table (ba_long_name, ba_type, area_id, area_type, title, creator_name, create_date, page_count, media_type, digital_format, document_type, item_category, item_sub_category, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24) RETURNING id INTO :25`, &dtrt.Ba_long_name, &dtrt.Ba_type, &dtrt.Area_id, &dtrt.Area_type, &dtrt.Title, &dtrt.Creator_name, &dtrt.Create_date, &dtrt.Page_count, &dtrt.Media_type, &dtrt.Digital_format, &dtrt.Document_type, &dtrt.Item_category, &dtrt.Item_sub_category, &dtrt.Ba_long_name_2, &dtrt.Ba_type_2, &dtrt.Data_store_name, &dtrt.Original_file_name, &dtrt.Password, &dtrt.Digital_size, &dtrt.Digital_size_uom, &dtrt.Remark, &dtrt.Source, &dtrt.Qc_status, &dtrt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(DIGITAL_TECHNICAL_REPORT_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetDigitalTechnicalReportById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM digital_technical_report_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_technical_report    
    
        for rows.Next() {
    
            var dtrt dto.Digital_technical_report
            if err := rows.Scan(&dtrt.Id, &dtrt.Ba_long_name, &dtrt.Ba_type, &dtrt.Area_id, &dtrt.Area_type, &dtrt.Title, &dtrt.Creator_name, &dtrt.Create_date, &dtrt.Page_count, &dtrt.Media_type, &dtrt.Digital_format, &dtrt.Document_type, &dtrt.Item_category, &dtrt.Item_sub_category, &dtrt.Ba_long_name_2, &dtrt.Ba_type_2, &dtrt.Data_store_name, &dtrt.Original_file_name, &dtrt.Password, &dtrt.Digital_size, &dtrt.Digital_size_uom, &dtrt.Remark, &dtrt.Source, &dtrt.Qc_status, &dtrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dtrt)
        }
    
    return c.JSON(results)
}
func PutDigitalTechnicalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dtrt dto.Digital_technical_report
    setDefaults(&dtrt)

    if err := c.BodyParser(&dtrt); err != nil{
        return err
    }
    
    dtrt.Create_date = formatDateString(dtrt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM digital_technical_report_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE digital_technical_report_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, title = :5, creator_name = :6, create_date = :7, page_count = :8, media_type = :9, digital_format = :10, document_type = :11, item_category = :12, item_sub_category = :13, ba_long_name_2 = :14, ba_type_2 = :15, data_store_name = :16, original_file_name = :17, password = :18, digital_size = :19, digital_size_uom = :20, remark = :21, source = :22, qc_status = :23, checked_by_ba_id = :24 WHERE id = :25`, &dtrt.Ba_long_name, &dtrt.Ba_type, &dtrt.Area_id, &dtrt.Area_type, &dtrt.Title, &dtrt.Creator_name, &dtrt.Create_date, &dtrt.Page_count, &dtrt.Media_type, &dtrt.Digital_format, &dtrt.Document_type, &dtrt.Item_category, &dtrt.Item_sub_category, &dtrt.Ba_long_name_2, &dtrt.Ba_type_2, &dtrt.Data_store_name, &dtrt.Original_file_name, &dtrt.Password, &dtrt.Digital_size, &dtrt.Digital_size_uom, &dtrt.Remark, &dtrt.Source, &dtrt.Qc_status, &dtrt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(DIGITAL_TECHNICAL_REPORT_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteDigitalTechnicalReport(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT digital_technical_report_id FROM digital_technical_report_workspace WHERE digital_technical_report_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM digital_technical_report_workspace WHERE digital_technical_report_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM digital_technical_report_table WHERE id = :1`, id)
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
func PatchDigitalTechnicalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dtrt dto.Digital_technical_report
    setDefaults(&dtrt)

    if err := c.BodyParser(&dtrt); err != nil{
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
    err = db.QueryRow(`SELECT digital_technical_report_id FROM digital_technical_report_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if dtrt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, dtrt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, dtrt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, dtrt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, dtrt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Title != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET title = :1 WHERE id = :2`, dtrt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET creator_name = :1 WHERE id = :2`, dtrt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, dtrt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Page_count != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET page_count = :1 WHERE id = :2`, dtrt.Page_count, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, dtrt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, dtrt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Document_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET document_type = :1 WHERE id = :2`, dtrt.Document_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Item_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_category = :1 WHERE id = :2`, dtrt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Item_sub_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_sub_category = :1 WHERE id = :2`, dtrt.Item_sub_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, dtrt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, dtrt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, dtrt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, dtrt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, dtrt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, dtrt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, dtrt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, dtrt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, dtrt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, dtrt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dtrt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, dtrt.Checked_by_ba_id, id)
        
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

    