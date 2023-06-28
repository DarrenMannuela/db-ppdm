package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-t3dseismicsummary/dto"
    "github.com/DarrenMannuela/svc-datatype-t3dseismicsummary/utils"

    "github.com/gofiber/fiber/v2"
)
func GetT3DSeismicSummary(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_summary_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_summary    
    
        for rows.Next() {
    
            var tsst dto.T3D_seismic_summary
            if err := rows.Scan(&tsst.Id, &tsst.Ba_long_name, &tsst.Ba_type, &tsst.Area_id, &tsst.Area_type, &tsst.Acqtn_survey_name, &tsst.Area_size, &tsst.Area_size_ouom, &tsst.Acqtn_inline_bin_size, &tsst.Acqtn_inline_bin_size_ouom, &tsst.Acqtn_xline_bin_size, &tsst.Acqtn_xline_bin_size_ouom, &tsst.Seis_dimension, &tsst.Start_date, &tsst.Shot_by, &tsst.Crew_long_name, &tsst.Acqtn_shot_line_spacing, &tsst.Acqtn_shot_line_spacing_ouom, &tsst.Rcvr_spacing, &tsst.Rcvr_spacing_ouom, &tsst.Rcvr_line_spacing, &tsst.Rcvr_line_spacing_ouom, &tsst.Seis_3d_type, &tsst.Energy_type, &tsst.Fold_coverage, &tsst.Rcrd_channel_count, &tsst.Rcrd_rec_length, &tsst.Rcrd_rec_length_ouom, &tsst.Rcrd_sample_rate, &tsst.Rcrd_sample_rate_ouom, &tsst.Alias_long_name, &tsst.Environment, &tsst.Remark, &tsst.Source, &tsst.Qc_status, &tsst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsst)
        }
    
    return c.JSON(results)
}
func SetT3DSeismicSummary(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var tsst dto.T3D_seismic_summary
    setDefaults(&tsst)

    if err := c.BodyParser(&tsst); err != nil{
        return err
    }
    
    tsst.Create_date = formatDateString(tsst.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO t3d_seismic_summary_table (ba_long_name, ba_type, area_id, area_type, acqtn_survey_name, area_size, area_size_ouom, acqtn_inline_bin_size, acqtn_inline_bin_size_ouom, acqtn_xline_bin_size, acqtn_xline_bin_size_ouom, seis_dimension, start_date, shot_by, crew_long_name, acqtn_shot_line_spacing, acqtn_shot_line_spacing_ouom, rcvr_spacing, rcvr_spacing_ouom, rcvr_line_spacing, rcvr_line_spacing_ouom, seis_3d_type, energy_type, fold_coverage, rcrd_channel_count, rcrd_rec_length, rcrd_rec_length_ouom, rcrd_sample_rate, rcrd_sample_rate_ouom, alias_long_name, environment, remark, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35) RETURNING id INTO :36`, &tsst.Ba_long_name, &tsst.Ba_type, &tsst.Area_id, &tsst.Area_type, &tsst.Acqtn_survey_name, &tsst.Area_size, &tsst.Area_size_ouom, &tsst.Acqtn_inline_bin_size, &tsst.Acqtn_inline_bin_size_ouom, &tsst.Acqtn_xline_bin_size, &tsst.Acqtn_xline_bin_size_ouom, &tsst.Seis_dimension, &tsst.Start_date, &tsst.Shot_by, &tsst.Crew_long_name, &tsst.Acqtn_shot_line_spacing, &tsst.Acqtn_shot_line_spacing_ouom, &tsst.Rcvr_spacing, &tsst.Rcvr_spacing_ouom, &tsst.Rcvr_line_spacing, &tsst.Rcvr_line_spacing_ouom, &tsst.Seis_3d_type, &tsst.Energy_type, &tsst.Fold_coverage, &tsst.Rcrd_channel_count, &tsst.Rcrd_rec_length, &tsst.Rcrd_rec_length_ouom, &tsst.Rcrd_sample_rate, &tsst.Rcrd_sample_rate_ouom, &tsst.Alias_long_name, &tsst.Environment, &tsst.Remark, &tsst.Source, &tsst.Qc_status, &tsst.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(T3D_SEISMIC_SUMMARY_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetT3DSeismicSummaryById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM t3d_seismic_summary_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.T3D_seismic_summary    
    
        for rows.Next() {
    
            var tsst dto.T3D_seismic_summary
            if err := rows.Scan(&tsst.Id, &tsst.Ba_long_name, &tsst.Ba_type, &tsst.Area_id, &tsst.Area_type, &tsst.Acqtn_survey_name, &tsst.Area_size, &tsst.Area_size_ouom, &tsst.Acqtn_inline_bin_size, &tsst.Acqtn_inline_bin_size_ouom, &tsst.Acqtn_xline_bin_size, &tsst.Acqtn_xline_bin_size_ouom, &tsst.Seis_dimension, &tsst.Start_date, &tsst.Shot_by, &tsst.Crew_long_name, &tsst.Acqtn_shot_line_spacing, &tsst.Acqtn_shot_line_spacing_ouom, &tsst.Rcvr_spacing, &tsst.Rcvr_spacing_ouom, &tsst.Rcvr_line_spacing, &tsst.Rcvr_line_spacing_ouom, &tsst.Seis_3d_type, &tsst.Energy_type, &tsst.Fold_coverage, &tsst.Rcrd_channel_count, &tsst.Rcrd_rec_length, &tsst.Rcrd_rec_length_ouom, &tsst.Rcrd_sample_rate, &tsst.Rcrd_sample_rate_ouom, &tsst.Alias_long_name, &tsst.Environment, &tsst.Remark, &tsst.Source, &tsst.Qc_status, &tsst.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, tsst)
        }
    
    return c.JSON(results)
}
func PutT3DSeismicSummary(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsst dto.T3D_seismic_summary
    setDefaults(&tsst)

    if err := c.BodyParser(&tsst); err != nil{
        return err
    }
    
    tsst.Create_date = formatDateString(tsst.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM t3d_seismic_summary_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, acqtn_survey_name = :5, area_size = :6, area_size_ouom = :7, acqtn_inline_bin_size = :8, acqtn_inline_bin_size_ouom = :9, acqtn_xline_bin_size = :10, acqtn_xline_bin_size_ouom = :11, seis_dimension = :12, start_date = :13, shot_by = :14, crew_long_name = :15, acqtn_shot_line_spacing = :16, acqtn_shot_line_spacing_ouom = :17, rcvr_spacing = :18, rcvr_spacing_ouom = :19, rcvr_line_spacing = :20, rcvr_line_spacing_ouom = :21, seis_3d_type = :22, energy_type = :23, fold_coverage = :24, rcrd_channel_count = :25, rcrd_rec_length = :26, rcrd_rec_length_ouom = :27, rcrd_sample_rate = :28, rcrd_sample_rate_ouom = :29, alias_long_name = :30, environment = :31, remark = :32, source = :33, qc_status = :34, checked_by_ba_id = :35 WHERE id = :36`, &tsst.Ba_long_name, &tsst.Ba_type, &tsst.Area_id, &tsst.Area_type, &tsst.Acqtn_survey_name, &tsst.Area_size, &tsst.Area_size_ouom, &tsst.Acqtn_inline_bin_size, &tsst.Acqtn_inline_bin_size_ouom, &tsst.Acqtn_xline_bin_size, &tsst.Acqtn_xline_bin_size_ouom, &tsst.Seis_dimension, &tsst.Start_date, &tsst.Shot_by, &tsst.Crew_long_name, &tsst.Acqtn_shot_line_spacing, &tsst.Acqtn_shot_line_spacing_ouom, &tsst.Rcvr_spacing, &tsst.Rcvr_spacing_ouom, &tsst.Rcvr_line_spacing, &tsst.Rcvr_line_spacing_ouom, &tsst.Seis_3d_type, &tsst.Energy_type, &tsst.Fold_coverage, &tsst.Rcrd_channel_count, &tsst.Rcrd_rec_length, &tsst.Rcrd_rec_length_ouom, &tsst.Rcrd_sample_rate, &tsst.Rcrd_sample_rate_ouom, &tsst.Alias_long_name, &tsst.Environment, &tsst.Remark, &tsst.Source, &tsst.Qc_status, &tsst.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(T3D_SEISMIC_SUMMARY_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteT3DSeismicSummary(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT t3d_seismic_summary_id FROM t3d_seismic_summary_workspace WHERE t3d_seismic_summary_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM t3d_seismic_summary_workspace WHERE t3d_seismic_summary_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM t3d_seismic_summary_table WHERE id = :1`, id)
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
func PatchT3DSeismicSummary(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var tsst dto.T3D_seismic_summary
    setDefaults(&tsst)

    if err := c.BodyParser(&tsst); err != nil{
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
    err = db.QueryRow(`SELECT t3d_seismic_summary_id FROM t3d_seismic_summary_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if tsst.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET ba_long_name = :1 WHERE id = :2`, tsst.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Ba_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET ba_type = :1 WHERE id = :2`, tsst.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Area_id != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET area_id = :1 WHERE id = :2`, tsst.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Area_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET area_type = :1 WHERE id = :2`, tsst.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Acqtn_survey_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET acqtn_survey_name = :1 WHERE id = :2`, tsst.Acqtn_survey_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Area_size != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET area_size = :1 WHERE id = :2`, tsst.Area_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Area_size_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET area_size_ouom = :1 WHERE id = :2`, tsst.Area_size_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Acqtn_inline_bin_size != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET acqtn_inline_bin_size = :1 WHERE id = :2`, tsst.Acqtn_inline_bin_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Acqtn_inline_bin_size_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET acqtn_inline_bin_size_ouom = :1 WHERE id = :2`, tsst.Acqtn_inline_bin_size_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Acqtn_xline_bin_size != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET acqtn_xline_bin_size = :1 WHERE id = :2`, tsst.Acqtn_xline_bin_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Acqtn_xline_bin_size_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET acqtn_xline_bin_size_ouom = :1 WHERE id = :2`, tsst.Acqtn_xline_bin_size_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Seis_dimension != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET seis_dimension = :1 WHERE id = :2`, tsst.Seis_dimension, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Start_date != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET start_date = :1 WHERE id = :2`, tsst.Start_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Shot_by != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET shot_by = :1 WHERE id = :2`, tsst.Shot_by, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Crew_long_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET crew_long_name = :1 WHERE id = :2`, tsst.Crew_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Acqtn_shot_line_spacing != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET acqtn_shot_line_spacing = :1 WHERE id = :2`, tsst.Acqtn_shot_line_spacing, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Acqtn_shot_line_spacing_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET acqtn_shot_line_spacing_ouom = :1 WHERE id = :2`, tsst.Acqtn_shot_line_spacing_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Rcvr_spacing != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET rcvr_spacing = :1 WHERE id = :2`, tsst.Rcvr_spacing, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Rcvr_spacing_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET rcvr_spacing_ouom = :1 WHERE id = :2`, tsst.Rcvr_spacing_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Rcvr_line_spacing != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET rcvr_line_spacing = :1 WHERE id = :2`, tsst.Rcvr_line_spacing, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Rcvr_line_spacing_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET rcvr_line_spacing_ouom = :1 WHERE id = :2`, tsst.Rcvr_line_spacing_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Seis_3d_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET seis_3d_type = :1 WHERE id = :2`, tsst.Seis_3d_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Energy_type != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET energy_type = :1 WHERE id = :2`, tsst.Energy_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Fold_coverage != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET fold_coverage = :1 WHERE id = :2`, tsst.Fold_coverage, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Rcrd_channel_count != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET rcrd_channel_count = :1 WHERE id = :2`, tsst.Rcrd_channel_count, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Rcrd_rec_length != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET rcrd_rec_length = :1 WHERE id = :2`, tsst.Rcrd_rec_length, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Rcrd_rec_length_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET rcrd_rec_length_ouom = :1 WHERE id = :2`, tsst.Rcrd_rec_length_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Rcrd_sample_rate != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET rcrd_sample_rate = :1 WHERE id = :2`, tsst.Rcrd_sample_rate, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Rcrd_sample_rate_ouom != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET rcrd_sample_rate_ouom = :1 WHERE id = :2`, tsst.Rcrd_sample_rate_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Alias_long_name != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET alias_long_name = :1 WHERE id = :2`, tsst.Alias_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Environment != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET environment = :1 WHERE id = :2`, tsst.Environment, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Remark != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET remark = :1 WHERE id = :2`, tsst.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Source != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET source = :1 WHERE id = :2`, tsst.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Qc_status != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET qc_status = :1 WHERE id = :2`, tsst.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if tsst.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE t3d_seismic_summary_table SET checked_by_ba_id = :1 WHERE id = :2`, tsst.Checked_by_ba_id, id)
        
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

    