package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-printtechnicalreport/dto"
    "github.com/DarrenMannuela/svc-datatype-printtechnicalreport/utils"

    "github.com/gofiber/fiber/v2"
)
func GetPrintTechnicalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM print_technical_report_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Print_technical_report    
    
        for rows.Next() {
    
            var ptrt dto.Print_technical_report
            if err := rows.Scan(&ptrt.Id, &ptrt.Ba_long_name, &ptrt.Ba_type, &ptrt.Area_id, &ptrt.Area_type, &ptrt.Title, &ptrt.Creator_name, &ptrt.Create_date, &ptrt.Page_count, &ptrt.Media_type, &ptrt.Document_type, &ptrt.Item_category, &ptrt.Item_sub_category, &ptrt.Ba_long_name_2, &ptrt.Ba_type_2, &ptrt.Data_store_name, &ptrt.Data_store_type, &ptrt.Location_id, &ptrt.Remark, &ptrt.Source, &ptrt.Qc_status, &ptrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, ptrt)
        }
    
    return c.JSON(results)
}
func SetPrintTechnicalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var ptrt dto.Print_technical_report
    setDefaults(&ptrt)

    if err := c.BodyParser(&ptrt); err != nil{
        return err
    }
    
    ptrt.Create_date = formatDateString(ptrt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO print_technical_report_table (ba_long_name, ba_type, area_id, area_type, title, creator_name, create_date, page_count, media_type, document_type, item_category, item_sub_category, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21) RETURNING id INTO :22`, &ptrt.Ba_long_name, &ptrt.Ba_type, &ptrt.Area_id, &ptrt.Area_type, &ptrt.Title, &ptrt.Creator_name, &ptrt.Create_date, &ptrt.Page_count, &ptrt.Media_type, &ptrt.Document_type, &ptrt.Item_category, &ptrt.Item_sub_category, &ptrt.Ba_long_name_2, &ptrt.Ba_type_2, &ptrt.Data_store_name, &ptrt.Data_store_type, &ptrt.Location_id, &ptrt.Remark, &ptrt.Source, &ptrt.Qc_status, &ptrt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(PRINT_TECHNICAL_REPORT_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetPrintTechnicalReportById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM print_technical_report_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Print_technical_report    
    
        for rows.Next() {
    
            var ptrt dto.Print_technical_report
            if err := rows.Scan(&ptrt.Id, &ptrt.Ba_long_name, &ptrt.Ba_type, &ptrt.Area_id, &ptrt.Area_type, &ptrt.Title, &ptrt.Creator_name, &ptrt.Create_date, &ptrt.Page_count, &ptrt.Media_type, &ptrt.Document_type, &ptrt.Item_category, &ptrt.Item_sub_category, &ptrt.Ba_long_name_2, &ptrt.Ba_type_2, &ptrt.Data_store_name, &ptrt.Data_store_type, &ptrt.Location_id, &ptrt.Remark, &ptrt.Source, &ptrt.Qc_status, &ptrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, ptrt)
        }
    
    return c.JSON(results)
}
func PutPrintTechnicalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var ptrt dto.Print_technical_report
    setDefaults(&ptrt)

    if err := c.BodyParser(&ptrt); err != nil{
        return err
    }
    
    ptrt.Create_date = formatDateString(ptrt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM print_technical_report_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE print_technical_report_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, title = :5, creator_name = :6, create_date = :7, page_count = :8, media_type = :9, document_type = :10, item_category = :11, item_sub_category = :12, ba_long_name_2 = :13, ba_type_2 = :14, data_store_name = :15, data_store_type = :16, location_id = :17, remark = :18, source = :19, qc_status = :20, checked_by_ba_id = :21 WHERE id = :22`, &ptrt.Ba_long_name, &ptrt.Ba_type, &ptrt.Area_id, &ptrt.Area_type, &ptrt.Title, &ptrt.Creator_name, &ptrt.Create_date, &ptrt.Page_count, &ptrt.Media_type, &ptrt.Document_type, &ptrt.Item_category, &ptrt.Item_sub_category, &ptrt.Ba_long_name_2, &ptrt.Ba_type_2, &ptrt.Data_store_name, &ptrt.Data_store_type, &ptrt.Location_id, &ptrt.Remark, &ptrt.Source, &ptrt.Qc_status, &ptrt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(PRINT_TECHNICAL_REPORT_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeletePrintTechnicalReport(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT print_technical_report_id FROM print_technical_report_workspace WHERE print_technical_report_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM print_technical_report_workspace WHERE print_technical_report_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM print_technical_report_table WHERE id = :1`, id)
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
func PatchPrintTechnicalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var ptrt dto.Print_technical_report
    setDefaults(&ptrt)

    if err := c.BodyParser(&ptrt); err != nil{
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
    err = db.QueryRow(`SELECT print_technical_report_id FROM print_technical_report_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if ptrt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, ptrt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, ptrt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, ptrt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, ptrt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Title != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET title = :1 WHERE id = :2`, ptrt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET creator_name = :1 WHERE id = :2`, ptrt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, ptrt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Page_count != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET page_count = :1 WHERE id = :2`, ptrt.Page_count, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, ptrt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Document_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET document_type = :1 WHERE id = :2`, ptrt.Document_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Item_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_category = :1 WHERE id = :2`, ptrt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Item_sub_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_sub_category = :1 WHERE id = :2`, ptrt.Item_sub_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, ptrt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, ptrt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, ptrt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_type = :1 WHERE id = :2`, ptrt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Location_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET location_id = :1 WHERE id = :2`, ptrt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, ptrt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, ptrt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, ptrt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ptrt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, ptrt.Checked_by_ba_id, id)
        
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

    