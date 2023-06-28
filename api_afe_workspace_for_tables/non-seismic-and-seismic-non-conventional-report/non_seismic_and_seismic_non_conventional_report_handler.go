package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionalreport/dto"
    "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionalreport/utils"

    "github.com/gofiber/fiber/v2"
)
func GetNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM non_seismic_and_seismic_non_conventional_report_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Non_seismic_and_seismic_non_conventional_report    
    
        for rows.Next() {
    
            var nsasncrt dto.Non_seismic_and_seismic_non_conventional_report
            if err := rows.Scan(&nsasncrt.Id, &nsasncrt.Ba_long_name, &nsasncrt.Ba_type, &nsasncrt.Area_id, &nsasncrt.Area_type, &nsasncrt.Acqtn_survey_name, &nsasncrt.Start_date, &nsasncrt.Title, &nsasncrt.Creator_name, &nsasncrt.Create_date, &nsasncrt.Item_category, &nsasncrt.Media_type, &nsasncrt.Ba_long_name_2, &nsasncrt.Ba_type_2, &nsasncrt.Data_store_name, &nsasncrt.Data_store_type, &nsasncrt.Location_id, &nsasncrt.Remark, &nsasncrt.Source, &nsasncrt.Qc_status, &nsasncrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, nsasncrt)
        }
    
    return c.JSON(results)
}
func SetNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var nsasncrt dto.Non_seismic_and_seismic_non_conventional_report
    setDefaults(&nsasncrt)

    if err := c.BodyParser(&nsasncrt); err != nil{
        return err
    }
    
    nsasncrt.Create_date = formatDateString(nsasncrt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO non_seismic_and_seismic_non_conventional_report_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, start_date, title, creator_name, create_date, item_category, media_type, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20) RETURNING id INTO :21`, &nsasncrt.Ba_long_name, &nsasncrt.Ba_type, &nsasncrt.Area_id, &nsasncrt.Area_type, &nsasncrt.Acqtn_survey_name, &nsasncrt.Start_date, &nsasncrt.Title, &nsasncrt.Creator_name, &nsasncrt.Create_date, &nsasncrt.Item_category, &nsasncrt.Media_type, &nsasncrt.Ba_long_name_2, &nsasncrt.Ba_type_2, &nsasncrt.Data_store_name, &nsasncrt.Data_store_type, &nsasncrt.Location_id, &nsasncrt.Remark, &nsasncrt.Source, &nsasncrt.Qc_status, &nsasncrt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetNonSeismicAndSeismicNonConventionalReportById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM non_seismic_and_seismic_non_conventional_report_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Non_seismic_and_seismic_non_conventional_report    
    
        for rows.Next() {
    
            var nsasncrt dto.Non_seismic_and_seismic_non_conventional_report
            if err := rows.Scan(&nsasncrt.Id, &nsasncrt.Ba_long_name, &nsasncrt.Ba_type, &nsasncrt.Area_id, &nsasncrt.Area_type, &nsasncrt.Acqtn_survey_name, &nsasncrt.Start_date, &nsasncrt.Title, &nsasncrt.Creator_name, &nsasncrt.Create_date, &nsasncrt.Item_category, &nsasncrt.Media_type, &nsasncrt.Ba_long_name_2, &nsasncrt.Ba_type_2, &nsasncrt.Data_store_name, &nsasncrt.Data_store_type, &nsasncrt.Location_id, &nsasncrt.Remark, &nsasncrt.Source, &nsasncrt.Qc_status, &nsasncrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, nsasncrt)
        }
    
    return c.JSON(results)
}
func PutNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var nsasncrt dto.Non_seismic_and_seismic_non_conventional_report
    setDefaults(&nsasncrt)

    if err := c.BodyParser(&nsasncrt); err != nil{
        return err
    }
    
    nsasncrt.Create_date = formatDateString(nsasncrt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM non_seismic_and_seismic_non_conventional_report_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, start_date = :6, title = :7, creator_name = :8, create_date = :9, item_category = :10, media_type = :11, ba_long_name_2 = :12, ba_type_2 = :13, data_store_name = :14, data_store_type = :15, location_id = :16, remark = :17, source = :18, qc_status = :19, checked_by_ba_id = :20 WHERE id = :21`, &nsasncrt.Ba_long_name, &nsasncrt.Ba_type, &nsasncrt.Area_id, &nsasncrt.Area_type, &nsasncrt.Acqtn_survey_name, &nsasncrt.Start_date, &nsasncrt.Title, &nsasncrt.Creator_name, &nsasncrt.Create_date, &nsasncrt.Item_category, &nsasncrt.Media_type, &nsasncrt.Ba_long_name_2, &nsasncrt.Ba_type_2, &nsasncrt.Data_store_name, &nsasncrt.Data_store_type, &nsasncrt.Location_id, &nsasncrt.Remark, &nsasncrt.Source, &nsasncrt.Qc_status, &nsasncrt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT non_seismic_and_seismic_non_conventional_report_id FROM non_seismic_and_seismic_non_conventional_report_workspace WHERE non_seismic_and_seismic_non_conventional_report_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_report_workspace WHERE non_seismic_and_seismic_non_conventional_report_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_report_table WHERE id = :1`, id)
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
func PatchNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var nsasncrt dto.Non_seismic_and_seismic_non_conventional_report
    setDefaults(&nsasncrt)

    if err := c.BodyParser(&nsasncrt); err != nil{
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
    err = db.QueryRow(`SELECT non_seismic_and_seismic_non_conventional_report_id FROM non_seismic_and_seismic_non_conventional_report_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if nsasncrt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, nsasncrt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, nsasncrt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, nsasncrt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, nsasncrt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET acqtn_survey_name = :1 WHERE id = :2`, nsasncrt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Start_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET start_date = :1 WHERE id = :2`, nsasncrt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Title != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET title = :1 WHERE id = :2`, nsasncrt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET creator_name = :1 WHERE id = :2`, nsasncrt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Create_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, nsasncrt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Item_category != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET item_category = :1 WHERE id = :2`, nsasncrt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Media_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, nsasncrt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, nsasncrt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, nsasncrt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, nsasncrt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_type = :1 WHERE id = :2`, nsasncrt.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Location_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET location_id = :1 WHERE id = :2`, nsasncrt.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, nsasncrt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, nsasncrt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, nsasncrt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncrt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, nsasncrt.Checked_by_ba_id, id)
        
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

    