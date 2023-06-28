package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-printwellreport/dto"
    "github.com/DarrenMannuela/svc-datatype-printwellreport/utils"

    "github.com/gofiber/fiber/v2"
)
func GetPrintWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM print_well_report_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Print_well_report    
    
        for rows.Next() {
    
            var pwrt dto.Print_well_report
            if err := rows.Scan(&pwrt.Id, &pwrt.Ba_long_name, &pwrt.Ba_type, &pwrt.Area_id, &pwrt.Area_type, &pwrt.Field_name, &pwrt.Well_name, &pwrt.Uwi, &pwrt.Title, &pwrt.Creator_name, &pwrt.Create_date, &pwrt.Media_type, &pwrt.Document_type, &pwrt.Item_category, &pwrt.Item_sub_category, &pwrt.Page_count, &pwrt.Remark, &pwrt.Data_store_name, &pwrt.Data_store_type, &pwrt.Source, &pwrt.Qc_status, &pwrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, pwrt)
        }
    
    return c.JSON(results)
}
func SetPrintWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var pwrt dto.Print_well_report
    setDefaults(&pwrt)

    if err := c.BodyParser(&pwrt); err != nill{
        return err
    }
    
    pwrt.Create_date = formatDateString(pwrt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO print_well_report_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, uwi, title, creator_name, create_date, media_type, document_type, item_category, item_sub_category, page_count, remark, data_store_name, data_store_type, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21) RETURNING id INTO :22`, &pwrt.Ba_long_name, &pwrt.Ba_type, &pwrt.Area_id, &pwrt.Area_type, &pwrt.Field_name, &pwrt.Well_name, &pwrt.Uwi, &pwrt.Title, &pwrt.Creator_name, &pwrt.Create_date, &pwrt.Media_type, &pwrt.Document_type, &pwrt.Item_category, &pwrt.Item_sub_category, &pwrt.Page_count, &pwrt.Remark, &pwrt.Data_store_name, &pwrt.Data_store_type, &pwrt.Source, &pwrt.Qc_status, &pwrt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(PRINT_WELL_REPORT_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetPrintWellReportById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM print_well_report_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Print_well_report    
    
        for rows.Next() {
    
            var pwrt dto.Print_well_report
            if err := rows.Scan(&pwrt.Id, &pwrt.Ba_long_name, &pwrt.Ba_type, &pwrt.Area_id, &pwrt.Area_type, &pwrt.Field_name, &pwrt.Well_name, &pwrt.Uwi, &pwrt.Title, &pwrt.Creator_name, &pwrt.Create_date, &pwrt.Media_type, &pwrt.Document_type, &pwrt.Item_category, &pwrt.Item_sub_category, &pwrt.Page_count, &pwrt.Remark, &pwrt.Data_store_name, &pwrt.Data_store_type, &pwrt.Source, &pwrt.Qc_status, &pwrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, pwrt)
        }
    
    return c.JSON(results)
}
func PutPrintWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var pwrt dto.Print_well_report
    setDefaults(&pwrt)

    if err := c.BodyParser(&pwrt); err != nill{
        return err
    }
    
    pwrt.Create_date = formatDateString(pwrt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM print_well_report_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE print_well_report_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, uwi = :7, title = :8, creator_name = :9, create_date = :10, media_type = :11, document_type = :12, item_category = :13, item_sub_category = :14, page_count = :15, remark = :16, data_store_name = :17, data_store_type = :18, source = :19, qc_status = :20, checked_by_ba_id = :21 WHERE id = :22`, &pwrt.Ba_long_name, &pwrt.Ba_type, &pwrt.Area_id, &pwrt.Area_type, &pwrt.Field_name, &pwrt.Well_name, &pwrt.Uwi, &pwrt.Title, &pwrt.Creator_name, &pwrt.Create_date, &pwrt.Media_type, &pwrt.Document_type, &pwrt.Item_category, &pwrt.Item_sub_category, &pwrt.Page_count, &pwrt.Remark, &pwrt.Data_store_name, &pwrt.Data_store_type, &pwrt.Source, &pwrt.Qc_status, &pwrt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(PRINT_WELL_REPORT_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeletePrintWellReport(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT print_well_report_id FROM print_well_report_workspace WHERE print_well_report_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM print_well_report_workspace WHERE print_well_report_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM print_well_report_table WHERE id = :1`, id)
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
func PatchPrintWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var pwrt dto.Print_well_report
    setDefaults(&pwrt)

    if err := c.BodyParser(&pwrt); err != nill{
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
    err = db.QueryRow(`SELECT print_well_report_id FROM print_well_report_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if pwrt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, pwrt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, pwrt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, pwrt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, pwrt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Field_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_name = :1 WHERE id = :2`, pwrt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Well_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET well_name = :1 WHERE id = :2`, pwrt.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Uwi != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET uwi = :1 WHERE id = :2`, pwrt.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Title != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET title = :1 WHERE id = :2`, pwrt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET creator_name = :1 WHERE id = :2`, pwrt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, pwrt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, pwrt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Document_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET document_type = :1 WHERE id = :2`, pwrt.Document_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Item_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_category = :1 WHERE id = :2`, pwrt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Item_sub_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_sub_category = :1 WHERE id = :2`, pwrt.Item_sub_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Page_count != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET page_count = :1 WHERE id = :2`, pwrt.Page_count, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, pwrt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, pwrt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_type = :1 WHERE id = :2`, pwrt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, pwrt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, pwrt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if pwrt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, pwrt.Checked_by_ba_id, id)
        
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

    