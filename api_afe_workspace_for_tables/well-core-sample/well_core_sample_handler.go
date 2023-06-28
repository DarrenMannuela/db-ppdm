package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-wellcoresample/dto"
    "github.com/DarrenMannuela/svc-datatype-wellcoresample/utils"

    "github.com/gofiber/fiber/v2"
)
func GetWellCoreSample(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM well_core_sample_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_core_sample    
    
        for rows.Next() {
    
            var wcst dto.Well_core_sample
            if err := rows.Scan(&wcst.Id, &wcst.Ba_long_name, &wcst.Ba_type, &wcst.Area_id, &wcst.Area_type, &wcst.Field_name, &wcst.Well_name, &wcst.Uwi, &wcst.Core_type, &wcst.Sample_num, &wcst.Sample_count, &wcst.Top_depth, &wcst.Top_depth_ouom, &wcst.Base_depth, &wcst.Base_depth_ouom, &wcst.Portion_volume, &wcst.Portion_volume_ouom, &wcst.Study_type, &wcst.Ba_long_name_2, &wcst.Ba_type_2, &wcst.Data_store_name, &wcst.Data_store_type, &wcst.Location_id, &wcst.Remark, &wcst.Source, &wcst.Qc_status, &wcst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wcst)
        }
    
    return c.JSON(results)
}
func SetWellCoreSample(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var wcst dto.Well_core_sample
    setDefaults(&wcst)

    if err := c.BodyParser(&wcst); err != nil{
        return err
    }
    
    wcst.Create_date = formatDateString(wcst.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO well_core_sample_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, uwi, core_type, sample_num, sample_count, top_depth, top_depth_ouom, base_depth, base_depth_ouom, portion_volume, portion_volume_ouom, study_type, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26) RETURNING id INTO :27`, &wcst.Ba_long_name, &wcst.Ba_type, &wcst.Area_id, &wcst.Area_type, &wcst.Field_name, &wcst.Well_name, &wcst.Uwi, &wcst.Core_type, &wcst.Sample_num, &wcst.Sample_count, &wcst.Top_depth, &wcst.Top_depth_ouom, &wcst.Base_depth, &wcst.Base_depth_ouom, &wcst.Portion_volume, &wcst.Portion_volume_ouom, &wcst.Study_type, &wcst.Ba_long_name_2, &wcst.Ba_type_2, &wcst.Data_store_name, &wcst.Data_store_type, &wcst.Location_id, &wcst.Remark, &wcst.Source, &wcst.Qc_status, &wcst.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(WELL_CORE_SAMPLE_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetWellCoreSampleById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM well_core_sample_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_core_sample    
    
        for rows.Next() {
    
            var wcst dto.Well_core_sample
            if err := rows.Scan(&wcst.Id, &wcst.Ba_long_name, &wcst.Ba_type, &wcst.Area_id, &wcst.Area_type, &wcst.Field_name, &wcst.Well_name, &wcst.Uwi, &wcst.Core_type, &wcst.Sample_num, &wcst.Sample_count, &wcst.Top_depth, &wcst.Top_depth_ouom, &wcst.Base_depth, &wcst.Base_depth_ouom, &wcst.Portion_volume, &wcst.Portion_volume_ouom, &wcst.Study_type, &wcst.Ba_long_name_2, &wcst.Ba_type_2, &wcst.Data_store_name, &wcst.Data_store_type, &wcst.Location_id, &wcst.Remark, &wcst.Source, &wcst.Qc_status, &wcst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, wcst)
        }
    
    return c.JSON(results)
}
func PutWellCoreSample(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wcst dto.Well_core_sample
    setDefaults(&wcst)

    if err := c.BodyParser(&wcst); err != nil{
        return err
    }
    
    wcst.Create_date = formatDateString(wcst.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM well_core_sample_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE well_core_sample_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, uwi = :7, core_type = :8, sample_num = :9, sample_count = :10, top_depth = :11, top_depth_ouom = :12, base_depth = :13, base_depth_ouom = :14, portion_volume = :15, portion_volume_ouom = :16, study_type = :17, ba_long_name_2 = :18, ba_type_2 = :19, data_store_name = :20, data_store_type = :21, location_id = :22, remark = :23, source = :24, qc_status = :25, checked_by_ba_id = :26 WHERE id = :27`, &wcst.Ba_long_name, &wcst.Ba_type, &wcst.Area_id, &wcst.Area_type, &wcst.Field_name, &wcst.Well_name, &wcst.Uwi, &wcst.Core_type, &wcst.Sample_num, &wcst.Sample_count, &wcst.Top_depth, &wcst.Top_depth_ouom, &wcst.Base_depth, &wcst.Base_depth_ouom, &wcst.Portion_volume, &wcst.Portion_volume_ouom, &wcst.Study_type, &wcst.Ba_long_name_2, &wcst.Ba_type_2, &wcst.Data_store_name, &wcst.Data_store_type, &wcst.Location_id, &wcst.Remark, &wcst.Source, &wcst.Qc_status, &wcst.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(WELL_CORE_SAMPLE_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteWellCoreSample(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT well_core_sample_id FROM well_core_sample_workspace WHERE well_core_sample_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM well_core_sample_workspace WHERE well_core_sample_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM well_core_sample_table WHERE id = :1`, id)
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
func PatchWellCoreSample(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wcst dto.Well_core_sample
    setDefaults(&wcst)

    if err := c.BodyParser(&wcst); err != nil{
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
    err = db.QueryRow(`SELECT well_core_sample_id FROM well_core_sample_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if wcst.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, wcst.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, wcst.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, wcst.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, wcst.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Field_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_name = :1 WHERE id = :2`, wcst.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Well_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET well_name = :1 WHERE id = :2`, wcst.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Uwi != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET uwi = :1 WHERE id = :2`, wcst.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Core_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET core_type = :1 WHERE id = :2`, wcst.Core_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Sample_num != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sample_num = :1 WHERE id = :2`, wcst.Sample_num, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Sample_count != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sample_count = :1 WHERE id = :2`, wcst.Sample_count, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Top_depth != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET top_depth = :1 WHERE id = :2`, wcst.Top_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Top_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET top_depth_ouom = :1 WHERE id = :2`, wcst.Top_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Base_depth != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET base_depth = :1 WHERE id = :2`, wcst.Base_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Base_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET base_depth_ouom = :1 WHERE id = :2`, wcst.Base_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Portion_volume != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET portion_volume = :1 WHERE id = :2`, wcst.Portion_volume, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Portion_volume_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET portion_volume_ouom = :1 WHERE id = :2`, wcst.Portion_volume_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Study_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET study_type = :1 WHERE id = :2`, wcst.Study_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, wcst.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, wcst.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, wcst.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_type = :1 WHERE id = :2`, wcst.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Location_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET location_id = :1 WHERE id = :2`, wcst.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, wcst.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, wcst.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, wcst.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wcst.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, wcst.Checked_by_ba_id, id)
        
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

    