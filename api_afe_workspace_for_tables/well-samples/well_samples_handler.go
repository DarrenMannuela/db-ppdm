package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-wellsamples/dto"
    "github.com/DarrenMannuela/svc-datatype-wellsamples/utils"

    "github.com/gofiber/fiber/v2"
)
func GetWellSamples(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM well_samples_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_samples    
    
        for rows.Next() {
    
            var wst dto.Well_samples
            if err := rows.Scan(&wst.Id, &wst.Ba_long_name, &wst.Ba_type, &wst.Area_id, &wst.Area_type, &wst.Field_name, &wst.Well_name, &wst.Uwi, &wst.Sample_type, &wst.Sample_num, &wst.Sample_count, &wst.Top_md, &wst.Top_md_ouom, &wst.Base_md, &wst.Base_md_ouom, &wst.Study_type, &wst.Ba_long_name_2, &wst.Ba_type_2, &wst.Data_store_name, &wst.Data_store_type, &wst.Location_id, &wst.Remark, &wst.Source, &wst.Qc_status, &wst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wst)
        }
    
    return c.JSON(results)
}
func SetWellSamples(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var wst dto.Well_samples
    setDefaults(&wst)

    if err := c.BodyParser(&wst); err != nil{
        return err
    }
    
    wst.Create_date = formatDateString(wst.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO well_samples_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, uwi, sample_type, sample_num, sample_count, top_md, top_md_ouom, base_md, base_md_ouom, study_type, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24) RETURNING id INTO :25`, &wst.Ba_long_name, &wst.Ba_type, &wst.Area_id, &wst.Area_type, &wst.Field_name, &wst.Well_name, &wst.Uwi, &wst.Sample_type, &wst.Sample_num, &wst.Sample_count, &wst.Top_md, &wst.Top_md_ouom, &wst.Base_md, &wst.Base_md_ouom, &wst.Study_type, &wst.Ba_long_name_2, &wst.Ba_type_2, &wst.Data_store_name, &wst.Data_store_type, &wst.Location_id, &wst.Remark, &wst.Source, &wst.Qc_status, &wst.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(WELL_SAMPLES_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetWellSamplesById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM well_samples_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_samples    
    
        for rows.Next() {
    
            var wst dto.Well_samples
            if err := rows.Scan(&wst.Id, &wst.Ba_long_name, &wst.Ba_type, &wst.Area_id, &wst.Area_type, &wst.Field_name, &wst.Well_name, &wst.Uwi, &wst.Sample_type, &wst.Sample_num, &wst.Sample_count, &wst.Top_md, &wst.Top_md_ouom, &wst.Base_md, &wst.Base_md_ouom, &wst.Study_type, &wst.Ba_long_name_2, &wst.Ba_type_2, &wst.Data_store_name, &wst.Data_store_type, &wst.Location_id, &wst.Remark, &wst.Source, &wst.Qc_status, &wst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wst)
        }
    
    return c.JSON(results)
}
func PutWellSamples(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wst dto.Well_samples
    setDefaults(&wst)

    if err := c.BodyParser(&wst); err != nil{
        return err
    }
    
    wst.Create_date = formatDateString(wst.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM well_samples_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE well_samples_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, uwi = :7, sample_type = :8, sample_num = :9, sample_count = :10, top_md = :11, top_md_ouom = :12, base_md = :13, base_md_ouom = :14, study_type = :15, ba_long_name_2 = :16, ba_type_2 = :17, data_store_name = :18, data_store_type = :19, location_id = :20, remark = :21, source = :22, qc_status = :23, checked_by_ba_id = :24 WHERE id = :25`, &wst.Ba_long_name, &wst.Ba_type, &wst.Area_id, &wst.Area_type, &wst.Field_name, &wst.Well_name, &wst.Uwi, &wst.Sample_type, &wst.Sample_num, &wst.Sample_count, &wst.Top_md, &wst.Top_md_ouom, &wst.Base_md, &wst.Base_md_ouom, &wst.Study_type, &wst.Ba_long_name_2, &wst.Ba_type_2, &wst.Data_store_name, &wst.Data_store_type, &wst.Location_id, &wst.Remark, &wst.Source, &wst.Qc_status, &wst.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(WELL_SAMPLES_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteWellSamples(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT well_samples_id FROM well_samples_workspace WHERE well_samples_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM well_samples_workspace WHERE well_samples_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM well_samples_table WHERE id = :1`, id)
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
func PatchWellSamples(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wst dto.Well_samples
    setDefaults(&wst)

    if err := c.BodyParser(&wst); err != nil{
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
    err = db.QueryRow(`SELECT well_samples_id FROM well_samples_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if wst.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, wst.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, wst.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, wst.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, wst.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Field_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_name = :1 WHERE id = :2`, wst.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Well_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET well_name = :1 WHERE id = :2`, wst.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Uwi != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET uwi = :1 WHERE id = :2`, wst.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Sample_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sample_type = :1 WHERE id = :2`, wst.Sample_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Sample_num != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sample_num = :1 WHERE id = :2`, wst.Sample_num, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Sample_count != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sample_count = :1 WHERE id = :2`, wst.Sample_count, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Top_md != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET top_md = :1 WHERE id = :2`, wst.Top_md, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Top_md_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET top_md_ouom = :1 WHERE id = :2`, wst.Top_md_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Base_md != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET base_md = :1 WHERE id = :2`, wst.Base_md, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Base_md_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET base_md_ouom = :1 WHERE id = :2`, wst.Base_md_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Study_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET study_type = :1 WHERE id = :2`, wst.Study_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, wst.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, wst.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, wst.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_type = :1 WHERE id = :2`, wst.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Location_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET location_id = :1 WHERE id = :2`, wst.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, wst.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, wst.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, wst.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wst.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, wst.Checked_by_ba_id, id)
        
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

    