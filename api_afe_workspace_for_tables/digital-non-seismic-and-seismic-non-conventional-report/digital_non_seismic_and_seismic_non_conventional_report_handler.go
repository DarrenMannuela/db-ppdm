package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-digitalnonseismicandseismicnonconventionalreport/dto"
    "github.com/DarrenMannuela/svc-datatype-digitalnonseismicandseismicnonconventionalreport/utils"

    "github.com/gofiber/fiber/v2"
)
func GetDigitalNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM digital_non_seismic_and_seismic_non_conventional_report_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_non_seismic_and_seismic_non_conventional_report    
    
        for rows.Next() {
    
            var dnsasncrt dto.Digital_non_seismic_and_seismic_non_conventional_report
            if err := rows.Scan(&dnsasncrt.Id, &dnsasncrt.Ba_long_name, &dnsasncrt.Ba_type, &dnsasncrt.Area_id, &dnsasncrt.Area_type, &dnsasncrt.Acqtn_survey_name, &dnsasncrt.Start_date, &dnsasncrt.Title, &dnsasncrt.Creator_name, &dnsasncrt.Create_date, &dnsasncrt.Item_category, &dnsasncrt.Digital_format, &dnsasncrt.Media_type, &dnsasncrt.Ba_long_name_2, &dnsasncrt.Ba_type_2, &dnsasncrt.Data_store_name, &dnsasncrt.Original_file_name, &dnsasncrt.Password, &dnsasncrt.Digital_size, &dnsasncrt.Digital_size_uom, &dnsasncrt.Remark, &dnsasncrt.Source, &dnsasncrt.Qc_status, &dnsasncrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dnsasncrt)
        }
    
    return c.JSON(results)
}
func SetDigitalNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var dnsasncrt dto.Digital_non_seismic_and_seismic_non_conventional_report
    setDefaults(&dnsasncrt)

    if err := c.BodyParser(&dnsasncrt); err != nil{
        return err
    }
    
    dnsasncrt.Create_date = formatDateString(dnsasncrt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO digital_non_seismic_and_seismic_non_conventional_report_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, start_date, title, creator_name, create_date, item_category, digital_format, media_type, ba_long_name_2, ba_type_2, data_store_name, original_file_name, password, digital_size, digital_size_uom, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23) RETURNING id INTO :24`, &dnsasncrt.Ba_long_name, &dnsasncrt.Ba_type, &dnsasncrt.Area_id, &dnsasncrt.Area_type, &dnsasncrt.Acqtn_survey_name, &dnsasncrt.Start_date, &dnsasncrt.Title, &dnsasncrt.Creator_name, &dnsasncrt.Create_date, &dnsasncrt.Item_category, &dnsasncrt.Digital_format, &dnsasncrt.Media_type, &dnsasncrt.Ba_long_name_2, &dnsasncrt.Ba_type_2, &dnsasncrt.Data_store_name, &dnsasncrt.Original_file_name, &dnsasncrt.Password, &dnsasncrt.Digital_size, &dnsasncrt.Digital_size_uom, &dnsasncrt.Remark, &dnsasncrt.Source, &dnsasncrt.Qc_status, &dnsasncrt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(DIGITAL_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetDigitalNonSeismicAndSeismicNonConventionalReportById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM digital_non_seismic_and_seismic_non_conventional_report_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Digital_non_seismic_and_seismic_non_conventional_report    
    
        for rows.Next() {
    
            var dnsasncrt dto.Digital_non_seismic_and_seismic_non_conventional_report
            if err := rows.Scan(&dnsasncrt.Id, &dnsasncrt.Ba_long_name, &dnsasncrt.Ba_type, &dnsasncrt.Area_id, &dnsasncrt.Area_type, &dnsasncrt.Acqtn_survey_name, &dnsasncrt.Start_date, &dnsasncrt.Title, &dnsasncrt.Creator_name, &dnsasncrt.Create_date, &dnsasncrt.Item_category, &dnsasncrt.Digital_format, &dnsasncrt.Media_type, &dnsasncrt.Ba_long_name_2, &dnsasncrt.Ba_type_2, &dnsasncrt.Data_store_name, &dnsasncrt.Original_file_name, &dnsasncrt.Password, &dnsasncrt.Digital_size, &dnsasncrt.Digital_size_uom, &dnsasncrt.Remark, &dnsasncrt.Source, &dnsasncrt.Qc_status, &dnsasncrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, dnsasncrt)
        }
    
    return c.JSON(results)
}
func PutDigitalNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dnsasncrt dto.Digital_non_seismic_and_seismic_non_conventional_report
    setDefaults(&dnsasncrt)

    if err := c.BodyParser(&dnsasncrt); err != nil{
        return err
    }
    
    dnsasncrt.Create_date = formatDateString(dnsasncrt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM digital_non_seismic_and_seismic_non_conventional_report_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, start_date = :6, title = :7, creator_name = :8, create_date = :9, item_category = :10, digital_format = :11, media_type = :12, ba_long_name_2 = :13, ba_type_2 = :14, data_store_name = :15, original_file_name = :16, password = :17, digital_size = :18, digital_size_uom = :19, remark = :20, source = :21, qc_status = :22, checked_by_ba_id = :23 WHERE id = :24`, &dnsasncrt.Ba_long_name, &dnsasncrt.Ba_type, &dnsasncrt.Area_id, &dnsasncrt.Area_type, &dnsasncrt.Acqtn_survey_name, &dnsasncrt.Start_date, &dnsasncrt.Title, &dnsasncrt.Creator_name, &dnsasncrt.Create_date, &dnsasncrt.Item_category, &dnsasncrt.Digital_format, &dnsasncrt.Media_type, &dnsasncrt.Ba_long_name_2, &dnsasncrt.Ba_type_2, &dnsasncrt.Data_store_name, &dnsasncrt.Original_file_name, &dnsasncrt.Password, &dnsasncrt.Digital_size, &dnsasncrt.Digital_size_uom, &dnsasncrt.Remark, &dnsasncrt.Source, &dnsasncrt.Qc_status, &dnsasncrt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(DIGITAL_NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_REPORT_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteDigitalNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT digital_non_seismic_and_seismic_non_conventional_report_id FROM digital_non_seismic_and_seismic_non_conventional_report_workspace WHERE digital_non_seismic_and_seismic_non_conventional_report_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM digital_non_seismic_and_seismic_non_conventional_report_workspace WHERE digital_non_seismic_and_seismic_non_conventional_report_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM digital_non_seismic_and_seismic_non_conventional_report_table WHERE id = :1`, id)
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
func PatchDigitalNonSeismicAndSeismicNonConventionalReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var dnsasncrt dto.Digital_non_seismic_and_seismic_non_conventional_report
    setDefaults(&dnsasncrt)

    if err := c.BodyParser(&dnsasncrt); err != nil{
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
    err = db.QueryRow(`SELECT digital_non_seismic_and_seismic_non_conventional_report_id FROM digital_non_seismic_and_seismic_non_conventional_report_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if dnsasncrt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, dnsasncrt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, dnsasncrt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Area_id != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, dnsasncrt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Area_type != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, dnsasncrt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET acqtn_survey_name = :1 WHERE id = :2`, dnsasncrt.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Start_date != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET start_date = :1 WHERE id = :2`, dnsasncrt.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Title != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET title = :1 WHERE id = :2`, dnsasncrt.Title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Creator_name != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET creator_name = :1 WHERE id = :2`, dnsasncrt.Creator_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Create_date != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET create_date = :1 WHERE id = :2`, dnsasncrt.Create_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Item_category != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET item_category = :1 WHERE id = :2`, dnsasncrt.Item_category, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Digital_format != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET digital_format = :1 WHERE id = :2`, dnsasncrt.Digital_format, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Media_type != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET media_type = :1 WHERE id = :2`, dnsasncrt.Media_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, dnsasncrt.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, dnsasncrt.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, dnsasncrt.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Original_file_name != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET original_file_name = :1 WHERE id = :2`, dnsasncrt.Original_file_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Password != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET password = :1 WHERE id = :2`, dnsasncrt.Password, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Digital_size != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET digital_size = :1 WHERE id = :2`, dnsasncrt.Digital_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Digital_size_uom != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET digital_size_uom = :1 WHERE id = :2`, dnsasncrt.Digital_size_uom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Remark != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, dnsasncrt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Source != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, dnsasncrt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, dnsasncrt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if dnsasncrt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE digital_non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, dnsasncrt.Checked_by_ba_id, id)
        
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

    