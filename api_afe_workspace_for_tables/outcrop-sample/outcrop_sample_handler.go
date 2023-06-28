package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-outcropsample/dto"
    "github.com/DarrenMannuela/svc-datatype-outcropsample/utils"

    "github.com/gofiber/fiber/v2"
)
func GetOutcropSample(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM outcrop_sample_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Outcrop_sample    
    
        for rows.Next() {
    
            var ost dto.Outcrop_sample
            if err := rows.Scan(&ost.Id, &ost.Ba_long_name, &ost.Ba_type, &ost.Area_id, &ost.Area_type, &ost.Project_name, &ost.Field_station_id, &ost.Longitude, &ost.Latitude, &ost.Easting, &ost.Easting_ouom, &ost.Northing, &ost.Northing_ouom, &ost.Utm_quadrant, &ost.Geodetic_datum_name, &ost.Sample_num, &ost.Sample_count, &ost.Study_type, &ost.Collected_date, &ost.Pick_location, &ost.Ba_long_name_2, &ost.Ba_type_2, &ost.Data_store_name, &ost.Data_store_type, &ost.Location_id, &ost.Remark, &ost.Source, &ost.Qc_status, &ost.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, ost)
        }
    
    return c.JSON(results)
}
func SetOutcropSample(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var ost dto.Outcrop_sample
    setDefaults(&ost)

    if err := c.BodyParser(&ost); err != nil{
        return err
    }
    
    ost.Create_date = formatDateString(ost.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO outcrop_sample_table (ba_long_name, ba_type, area_id, area_type, project_name, field_station_id, longitude, latitude, easting, easting_ouom, northing, northing_ouom, utm_quadrant, geodetic_datum_name, sample_num, sample_count, study_type, collected_date, pick_location, ba_long_name_2, ba_type_2, data_store_name, data_store_type, location_id, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28) RETURNING id INTO :29`, &ost.Ba_long_name, &ost.Ba_type, &ost.Area_id, &ost.Area_type, &ost.Project_name, &ost.Field_station_id, &ost.Longitude, &ost.Latitude, &ost.Easting, &ost.Easting_ouom, &ost.Northing, &ost.Northing_ouom, &ost.Utm_quadrant, &ost.Geodetic_datum_name, &ost.Sample_num, &ost.Sample_count, &ost.Study_type, &ost.Collected_date, &ost.Pick_location, &ost.Ba_long_name_2, &ost.Ba_type_2, &ost.Data_store_name, &ost.Data_store_type, &ost.Location_id, &ost.Remark, &ost.Source, &ost.Qc_status, &ost.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(OUTCROP_SAMPLE_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetOutcropSampleById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM outcrop_sample_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Outcrop_sample    
    
        for rows.Next() {
    
            var ost dto.Outcrop_sample
            if err := rows.Scan(&ost.Id, &ost.Ba_long_name, &ost.Ba_type, &ost.Area_id, &ost.Area_type, &ost.Project_name, &ost.Field_station_id, &ost.Longitude, &ost.Latitude, &ost.Easting, &ost.Easting_ouom, &ost.Northing, &ost.Northing_ouom, &ost.Utm_quadrant, &ost.Geodetic_datum_name, &ost.Sample_num, &ost.Sample_count, &ost.Study_type, &ost.Collected_date, &ost.Pick_location, &ost.Ba_long_name_2, &ost.Ba_type_2, &ost.Data_store_name, &ost.Data_store_type, &ost.Location_id, &ost.Remark, &ost.Source, &ost.Qc_status, &ost.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, ost)
        }
    
    return c.JSON(results)
}
func PutOutcropSample(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var ost dto.Outcrop_sample
    setDefaults(&ost)

    if err := c.BodyParser(&ost); err != nil{
        return err
    }
    
    ost.Create_date = formatDateString(ost.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM outcrop_sample_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE outcrop_sample_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, project_name = :5, field_station_id = :6, longitude = :7, latitude = :8, easting = :9, easting_ouom = :10, northing = :11, northing_ouom = :12, utm_quadrant = :13, geodetic_datum_name = :14, sample_num = :15, sample_count = :16, study_type = :17, collected_date = :18, pick_location = :19, ba_long_name_2 = :20, ba_type_2 = :21, data_store_name = :22, data_store_type = :23, location_id = :24, remark = :25, source = :26, qc_status = :27, checked_by_ba_id = :28 WHERE id = :29`, &ost.Ba_long_name, &ost.Ba_type, &ost.Area_id, &ost.Area_type, &ost.Project_name, &ost.Field_station_id, &ost.Longitude, &ost.Latitude, &ost.Easting, &ost.Easting_ouom, &ost.Northing, &ost.Northing_ouom, &ost.Utm_quadrant, &ost.Geodetic_datum_name, &ost.Sample_num, &ost.Sample_count, &ost.Study_type, &ost.Collected_date, &ost.Pick_location, &ost.Ba_long_name_2, &ost.Ba_type_2, &ost.Data_store_name, &ost.Data_store_type, &ost.Location_id, &ost.Remark, &ost.Source, &ost.Qc_status, &ost.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(OUTCROP_SAMPLE_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteOutcropSample(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT outcrop_sample_id FROM outcrop_sample_workspace WHERE outcrop_sample_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM outcrop_sample_workspace WHERE outcrop_sample_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM outcrop_sample_table WHERE id = :1`, id)
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
func PatchOutcropSample(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var ost dto.Outcrop_sample
    setDefaults(&ost)

    if err := c.BodyParser(&ost); err != nil{
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
    err = db.QueryRow(`SELECT outcrop_sample_id FROM outcrop_sample_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if ost.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, ost.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, ost.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, ost.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, ost.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Project_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET project_name = :1 WHERE id = :2`, ost.Project_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Field_station_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET field_station_id = :1 WHERE id = :2`, ost.Field_station_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Longitude != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET longitude = :1 WHERE id = :2`, ost.Longitude, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Latitude != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET latitude = :1 WHERE id = :2`, ost.Latitude, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Easting != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET easting = :1 WHERE id = :2`, ost.Easting, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Easting_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET easting_ouom = :1 WHERE id = :2`, ost.Easting_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Northing != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET northing = :1 WHERE id = :2`, ost.Northing, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Northing_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET northing_ouom = :1 WHERE id = :2`, ost.Northing_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Utm_quadrant != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET utm_quadrant = :1 WHERE id = :2`, ost.Utm_quadrant, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Geodetic_datum_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET geodetic_datum_name = :1 WHERE id = :2`, ost.Geodetic_datum_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Sample_num != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sample_num = :1 WHERE id = :2`, ost.Sample_num, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Sample_count != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET sample_count = :1 WHERE id = :2`, ost.Sample_count, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Study_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET study_type = :1 WHERE id = :2`, ost.Study_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Collected_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET collected_date = :1 WHERE id = :2`, ost.Collected_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Pick_location != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET pick_location = :1 WHERE id = :2`, ost.Pick_location, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Ba_long_name_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name_2 = :1 WHERE id = :2`, ost.Ba_long_name_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Ba_type_2 != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type_2 = :1 WHERE id = :2`, ost.Ba_type_2, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_name = :1 WHERE id = :2`, ost.Data_store_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Data_store_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET data_store_type = :1 WHERE id = :2`, ost.Data_store_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Location_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET location_id = :1 WHERE id = :2`, ost.Location_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, ost.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, ost.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, ost.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if ost.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, ost.Checked_by_ba_id, id)
        
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

    