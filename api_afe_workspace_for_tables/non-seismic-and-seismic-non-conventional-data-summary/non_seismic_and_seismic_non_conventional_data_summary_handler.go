package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionaldatasummary/dto"
    "github.com/DarrenMannuela/svc-datatype-nonseismicandseismicnonconventionaldatasummary/utils"

    "github.com/gofiber/fiber/v2"
)
func GetNonSeismicAndSeismicNonConventionalDataSummary(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM non_seismic_and_seismic_non_conventional_data_summary_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Non_seismic_and_seismic_non_conventional_data_summary    
    
        for rows.Next() {
    
            var nsasncdst dto.Non_seismic_and_seismic_non_conventional_data_summary
            if err := rows.Scan(&nsasncdst.Id, &nsasncdst.Ba_long_name, &nsasncdst.Ba_type, &nsasncdst.Area_id, &nsasncdst.Area_type, &nsasncdst.Acqtn_survey_name, &nsasncdst.Seis_spectrum_type, &nsasncdst.Seis_dimension, &nsasncdst.Start_date, &nsasncdst.Shot_by, &nsasncdst.Crew_long_name, &nsasncdst.Acqtn_direction, &nsasncdst.Environment, &nsasncdst.Remark, &nsasncdst.Source, &nsasncdst.Qc_status, &nsasncdst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, nsasncdst)
        }
    
    return c.JSON(results)
}
func SetNonSeismicAndSeismicNonConventionalDataSummary(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var nsasncdst dto.Non_seismic_and_seismic_non_conventional_data_summary
    setDefaults(&nsasncdst)

    if err := c.BodyParser(&nsasncdst); err != nil{
        return err
    }
    
    nsasncdst.Create_date = formatDateString(nsasncdst.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO non_seismic_and_seismic_non_conventional_data_summary_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, seis_spectrum_type, seis_dimension, start_date, shot_by, crew_long_name, acqtn_direction, environment, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16) RETURNING id INTO :17`, &nsasncdst.Ba_long_name, &nsasncdst.Ba_type, &nsasncdst.Area_id, &nsasncdst.Area_type, &nsasncdst.Acqtn_survey_name, &nsasncdst.Seis_spectrum_type, &nsasncdst.Seis_dimension, &nsasncdst.Start_date, &nsasncdst.Shot_by, &nsasncdst.Crew_long_name, &nsasncdst.Acqtn_direction, &nsasncdst.Environment, &nsasncdst.Remark, &nsasncdst.Source, &nsasncdst.Qc_status, &nsasncdst.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetNonSeismicAndSeismicNonConventionalDataSummaryById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM non_seismic_and_seismic_non_conventional_data_summary_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Non_seismic_and_seismic_non_conventional_data_summary    
    
        for rows.Next() {
    
            var nsasncdst dto.Non_seismic_and_seismic_non_conventional_data_summary
            if err := rows.Scan(&nsasncdst.Id, &nsasncdst.Ba_long_name, &nsasncdst.Ba_type, &nsasncdst.Area_id, &nsasncdst.Area_type, &nsasncdst.Acqtn_survey_name, &nsasncdst.Seis_spectrum_type, &nsasncdst.Seis_dimension, &nsasncdst.Start_date, &nsasncdst.Shot_by, &nsasncdst.Crew_long_name, &nsasncdst.Acqtn_direction, &nsasncdst.Environment, &nsasncdst.Remark, &nsasncdst.Source, &nsasncdst.Qc_status, &nsasncdst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, nsasncdst)
        }
    
    return c.JSON(results)
}
func PutNonSeismicAndSeismicNonConventionalDataSummary(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var nsasncdst dto.Non_seismic_and_seismic_non_conventional_data_summary
    setDefaults(&nsasncdst)

    if err := c.BodyParser(&nsasncdst); err != nil{
        return err
    }
    
    nsasncdst.Create_date = formatDateString(nsasncdst.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM non_seismic_and_seismic_non_conventional_data_summary_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_data_summary_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, seis_spectrum_type = :6, seis_dimension = :7, start_date = :8, shot_by = :9, crew_long_name = :10, acqtn_direction = :11, environment = :12, remark = :13, source = :14, qc_status = :15, checked_by_ba_id = :16 WHERE id = :17`, &nsasncdst.Ba_long_name, &nsasncdst.Ba_type, &nsasncdst.Area_id, &nsasncdst.Area_type, &nsasncdst.Acqtn_survey_name, &nsasncdst.Seis_spectrum_type, &nsasncdst.Seis_dimension, &nsasncdst.Start_date, &nsasncdst.Shot_by, &nsasncdst.Crew_long_name, &nsasncdst.Acqtn_direction, &nsasncdst.Environment, &nsasncdst.Remark, &nsasncdst.Source, &nsasncdst.Qc_status, &nsasncdst.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(NON_SEISMIC_AND_SEISMIC_NON_CONVENTIONAL_DATA_SUMMARY_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteNonSeismicAndSeismicNonConventionalDataSummary(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT non_seismic_and_seismic_non_conventional_data_summary_id FROM non_seismic_and_seismic_non_conventional_data_summary_workspace WHERE non_seismic_and_seismic_non_conventional_data_summary_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_data_summary_workspace WHERE non_seismic_and_seismic_non_conventional_data_summary_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM non_seismic_and_seismic_non_conventional_data_summary_table WHERE id = :1`, id)
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
func PatchNonSeismicAndSeismicNonConventionalDataSummary(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var nsasncdst dto.Non_seismic_and_seismic_non_conventional_data_summary
    setDefaults(&nsasncdst)

    if err := c.BodyParser(&nsasncdst); err != nil{
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
    err = db.QueryRow(`SELECT non_seismic_and_seismic_non_conventional_data_summary_id FROM non_seismic_and_seismic_non_conventional_data_summary_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if nsasncdst.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_long_name = :1 WHERE id = :2`, nsasncdst.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Ba_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET ba_type = :1 WHERE id = :2`, nsasncdst.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Area_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_id = :1 WHERE id = :2`, nsasncdst.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Area_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET area_type = :1 WHERE id = :2`, nsasncdst.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET acqtn_survey_name = :1 WHERE id = :2`, nsasncdst.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Seis_spectrum_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET seis_spectrum_type = :1 WHERE id = :2`, nsasncdst.Seis_spectrum_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Seis_dimension != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET seis_dimension = :1 WHERE id = :2`, nsasncdst.Seis_dimension, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Start_date != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET start_date = :1 WHERE id = :2`, nsasncdst.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Shot_by != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET shot_by = :1 WHERE id = :2`, nsasncdst.Shot_by, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Crew_long_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET crew_long_name = :1 WHERE id = :2`, nsasncdst.Crew_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Acqtn_direction != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET acqtn_direction = :1 WHERE id = :2`, nsasncdst.Acqtn_direction, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Environment != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET environment = :1 WHERE id = :2`, nsasncdst.Environment, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Remark != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET remark = :1 WHERE id = :2`, nsasncdst.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, nsasncdst.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, nsasncdst.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if nsasncdst.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, nsasncdst.Checked_by_ba_id, id)
        
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

    