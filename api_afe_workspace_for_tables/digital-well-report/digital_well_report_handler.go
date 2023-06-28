package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-digitalwellreport/dto"
    "github.com/DarrenMannuela/svc-datatype-digitalwellreport/utils"

    "github.com/gofiber/fiber/v2"
)
func GetDigitalWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM digital_well_report_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_well_report    
    
        for rows.Next() {
    
            var dwrt dto.Digital_well_report
            if err := rows.Scan(&dwrt.Id, &dwrt.Ba_long_name, &dwrt.Ba_type, &dwrt.Area_id, &dwrt.Area_type, &dwrt.Field_name, &dwrt.Well_name, &dwrt.Uwi, &dwrt.Title, &dwrt.Creator_name, &dwrt.Create_date, &dwrt.Media_type, &dwrt.Document_type, &dwrt.Item_category, &dwrt.Item_sub_category, &dwrt.Digital_format, &dwrt.Ba_long_name_2, &dwrt.Ba_type_2, &dwrt.Data_store_name, &dwrt.Original_file_name, &dwrt.Password, &dwrt.Digital_size, &dwrt.Digital_size_uom, &dwrt.Remark, &dwrt.Source, &dwrt.Qc_status, &dwrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dwrt)
        }
    
    return c.JSON(results)
}
func SetDigitalWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var dwrt dto.Digital_well_report
    setDefaults(&dwrt)

    if err := c.BodyParser(&dwrt); err != nil{
        return err
    }
    
    dwrt.Create_date = formatDateString(dwrt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO digital_well_report_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, uwi, title, creator_name, create_date, media_type, document_type, item_category, item_sub_category, digital_format, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26) RETURNING id INTO :27`, &dwrt.Ba_long_name, &dwrt.Ba_type, &dwrt.Area_id, &dwrt.Area_type, &dwrt.Field_name, &dwrt.Well_name, &dwrt.Uwi, &dwrt.Title, &dwrt.Creator_name, &dwrt.Create_date, &dwrt.Media_type, &dwrt.Document_type, &dwrt.Item_category, &dwrt.Item_sub_category, &dwrt.Digital_format, &dwrt.Ba_long_name_2, &dwrt.Ba_type_2, &dwrt.Data_store_name, &dwrt.Original_file_name, &dwrt.Password, &dwrt.Digital_size, &dwrt.Digital_size_uom, &dwrt.Remark, &dwrt.Source, &dwrt.Qc_status, &dwrt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(DIGITAL_WELL_REPORT_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetDigitalWellReportById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM digital_well_report_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_well_report    
    
        for rows.Next() {
    
            var dwrt dto.Digital_well_report
            if err := rows.Scan(&dwrt.Id, &dwrt.Ba_long_name, &dwrt.Ba_type, &dwrt.Area_id, &dwrt.Area_type, &dwrt.Field_name, &dwrt.Well_name, &dwrt.Uwi, &dwrt.Title, &dwrt.Creator_name, &dwrt.Create_date, &dwrt.Media_type, &dwrt.Document_type, &dwrt.Item_category, &dwrt.Item_sub_category, &dwrt.Digital_format, &dwrt.Ba_long_name_2, &dwrt.Ba_type_2, &dwrt.Data_store_name, &dwrt.Original_file_name, &dwrt.Password, &dwrt.Digital_size, &dwrt.Digital_size_uom, &dwrt.Remark, &dwrt.Source, &dwrt.Qc_status, &dwrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dwrt)
        }
    
    return c.JSON(results)
}
func PutDigitalWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dwrt dto.Digital_well_report
    setDefaults(&dwrt)

    if err := c.BodyParser(&dwrt); err != nil{
        return err
    }
    
    dwrt.Create_date = formatDateString(dwrt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM digital_well_report_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE digital_well_report_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, uwi = :7, title = :8, creator_name = :9, create_date = :10, media_type = :11, document_type = :12, item_category = :13, item_sub_category = :14, digital_format = :15, ba_long_name_2 = :16, ba_type_2 = :17, data_store_name = :18, original_file_name = :19, password = :20, digital_size = :21, digital_size_uom = :22, remark = :23, source = :24, qc_status = :25, checked_by_ba_id = :26 WHERE id = :27`, &dwrt.Ba_long_name, &dwrt.Ba_type, &dwrt.Area_id, &dwrt.Area_type, &dwrt.Field_name, &dwrt.Well_name, &dwrt.Uwi, &dwrt.Title, &dwrt.Creator_name, &dwrt.Create_date, &dwrt.Media_type, &dwrt.Document_type, &dwrt.Item_category, &dwrt.Item_sub_category, &dwrt.Digital_format, &dwrt.Ba_long_name_2, &dwrt.Ba_type_2, &dwrt.Data_store_name, &dwrt.Original_file_name, &dwrt.Password, &dwrt.Digital_size, &dwrt.Digital_size_uom, &dwrt.Remark, &dwrt.Source, &dwrt.Qc_status, &dwrt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(DIGITAL_WELL_REPORT_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteDigitalWellReport(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT digital_well_report_id FROM digital_well_report_workspace WHERE digital_well_report_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM digital_well_report_workspace WHERE digital_well_report_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM digital_well_report_table WHERE id = :1`, id)
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
func PatchDigitalWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dwrt dto.Digital_well_report
    setDefaults(&dwrt)

    if err := c.BodyParser(&dwrt); err != nil{
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
    err = db.QueryRow(`SELECT digital_well_report_id FROM digital_well_report_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if dwrt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, dwrt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, dwrt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, dwrt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, dwrt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Field_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_name = :1 WHERE id = :2`, dwrt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Well_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET well_name = :1 WHERE id = :2`, dwrt.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Uwi != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET uwi = :1 WHERE id = :2`, dwrt.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Title != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET title = :1 WHERE id = :2`, dwrt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET creator_name = :1 WHERE id = :2`, dwrt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, dwrt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, dwrt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Document_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET document_type = :1 WHERE id = :2`, dwrt.Document_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Item_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_category = :1 WHERE id = :2`, dwrt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Item_sub_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_sub_category = :1 WHERE id = :2`, dwrt.Item_sub_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, dwrt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, dwrt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, dwrt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, dwrt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, dwrt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Password != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, dwrt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, dwrt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, dwrt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, dwrt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, dwrt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, dwrt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dwrt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, dwrt.Checked_by_ba_id, id)
        
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

    