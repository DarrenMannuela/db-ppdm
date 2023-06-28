package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-welldata/dto"
    "github.com/DarrenMannuela/svc-datatype-welldata/utils"

    "github.com/gofiber/fiber/v2"
)
func GetWellData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM well_data_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_data    
    
        for rows.Next() {
    
            var wdt dto.Well_data
            if err := rows.Scan(&wdt.Id, &wdt.Ba_long_name, &wdt.Ba_type, &wdt.Area_id, &wdt.Area_type, &wdt.Field_name, &wdt.Well_name, &wdt.Alias_long_name, &wdt.Uwi, &wdt.Status_type, &wdt.Current_class, &wdt.Well_level_type, &wdt.Profile_type, &wdt.Projected_strat_unit_id, &wdt.Surface_longitude, &wdt.Surface_latitude, &wdt.Easting, &wdt.Easting_ouom, &wdt.Northing, &wdt.Northing_ouom, &wdt.Utm_quadrant, &wdt.Projection_type, &wdt.Geodetic_datum_name, &wdt.Environment_type, &wdt.Line_name, &wdt.Seis_point_id, &wdt.Spud_date, &wdt.Final_drill_date, &wdt.Completion_date, &wdt.Rotary_table_elev, &wdt.Rotary_table_elev_ouom, &wdt.Kb_elev, &wdt.Kb_elev_ouom, &wdt.Derrick_floor_elev, &wdt.Derrick_floor_elev_ouom, &wdt.Water_depth, &wdt.Water_depth_ouom, &wdt.Ground_elev, &wdt.Ground_elev_ouom, &wdt.Depth_datum_elev, &wdt.Drill_td, &wdt.Drill_td_ouom, &wdt.Log_td, &wdt.Log_td_ouom, &wdt.Max_tvd, &wdt.Max_tvd_ouom, &wdt.Projected_depth, &wdt.Projected_depth_ouom, &wdt.Final_td, &wdt.Final_td_ouom, &wdt.Operator_ba_id, &wdt.Rig_name, &wdt.Rig_type, &wdt.Test_result_code, &wdt.Remark, &wdt.Source, &wdt.Qc_status, &wdt.Checked_by_ba_id, &wdt.Tubing_obs_no, &wdt.Outside_diameter, &wdt.Outside_diameter_ouom, &wdt.Base_depth, &wdt.Base_depth_ouom); err != nil{
                return err
            }
            results = append(results, wdt)
        }
    
    return c.JSON(results)
}
func SetWellData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var wdt dto.Well_data
    setDefaults(&wdt)

    if err := c.BodyParser(&wdt); err != nil{
        return err
    }
    
    wdt.Create_date = formatDateString(wdt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO well_data_table (ba_long_name, ba_type, area_id, area_type, field_name, well_name, alias_long_name, uwi, status_type, current_class, well_level_type, profile_type, projected_strat_unit_id, surface_longitude, surface_latitude, easting, easting_ouom, northing, northing_ouom, utm_quadrant, projection_type, geodetic_datum_name, environment_type, line_name, seis_point_id, spud_date, final_drill_date, completion_date, rotary_table_elev, rotary_table_elev_ouom, kb_elev, kb_elev_ouom, derrick_floor_elev, derrick_floor_elev_ouom, water_depth, water_depth_ouom, ground_elev, ground_elev_ouom, depth_datum_elev, drill_td, drill_td_ouom, log_td, log_td_ouom, max_tvd, max_tvd_ouom, projected_depth, projected_depth_ouom, final_td, final_td_ouom, operator_ba_id, rig_name, rig_type, test_result_code, remark, source, qc_status, checked_by_ba_id, tubing_obs_no, outside_diameter, outside_diameter_ouom, base_depth, base_depth_ouom) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62) RETURNING id INTO :63`, &wdt.Ba_long_name, &wdt.Ba_type, &wdt.Area_id, &wdt.Area_type, &wdt.Field_name, &wdt.Well_name, &wdt.Alias_long_name, &wdt.Uwi, &wdt.Status_type, &wdt.Current_class, &wdt.Well_level_type, &wdt.Profile_type, &wdt.Projected_strat_unit_id, &wdt.Surface_longitude, &wdt.Surface_latitude, &wdt.Easting, &wdt.Easting_ouom, &wdt.Northing, &wdt.Northing_ouom, &wdt.Utm_quadrant, &wdt.Projection_type, &wdt.Geodetic_datum_name, &wdt.Environment_type, &wdt.Line_name, &wdt.Seis_point_id, &wdt.Spud_date, &wdt.Final_drill_date, &wdt.Completion_date, &wdt.Rotary_table_elev, &wdt.Rotary_table_elev_ouom, &wdt.Kb_elev, &wdt.Kb_elev_ouom, &wdt.Derrick_floor_elev, &wdt.Derrick_floor_elev_ouom, &wdt.Water_depth, &wdt.Water_depth_ouom, &wdt.Ground_elev, &wdt.Ground_elev_ouom, &wdt.Depth_datum_elev, &wdt.Drill_td, &wdt.Drill_td_ouom, &wdt.Log_td, &wdt.Log_td_ouom, &wdt.Max_tvd, &wdt.Max_tvd_ouom, &wdt.Projected_depth, &wdt.Projected_depth_ouom, &wdt.Final_td, &wdt.Final_td_ouom, &wdt.Operator_ba_id, &wdt.Rig_name, &wdt.Rig_type, &wdt.Test_result_code, &wdt.Remark, &wdt.Source, &wdt.Qc_status, &wdt.Checked_by_ba_id, &wdt.Tubing_obs_no, &wdt.Outside_diameter, &wdt.Outside_diameter_ouom, &wdt.Base_depth, &wdt.Base_depth_ouom, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(WELL_DATA_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetWellDataById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM well_data_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Well_data    
    
        for rows.Next() {
    
            var wdt dto.Well_data
            if err := rows.Scan(&wdt.Id, &wdt.Ba_long_name, &wdt.Ba_type, &wdt.Area_id, &wdt.Area_type, &wdt.Field_name, &wdt.Well_name, &wdt.Alias_long_name, &wdt.Uwi, &wdt.Status_type, &wdt.Current_class, &wdt.Well_level_type, &wdt.Profile_type, &wdt.Projected_strat_unit_id, &wdt.Surface_longitude, &wdt.Surface_latitude, &wdt.Easting, &wdt.Easting_ouom, &wdt.Northing, &wdt.Northing_ouom, &wdt.Utm_quadrant, &wdt.Projection_type, &wdt.Geodetic_datum_name, &wdt.Environment_type, &wdt.Line_name, &wdt.Seis_point_id, &wdt.Spud_date, &wdt.Final_drill_date, &wdt.Completion_date, &wdt.Rotary_table_elev, &wdt.Rotary_table_elev_ouom, &wdt.Kb_elev, &wdt.Kb_elev_ouom, &wdt.Derrick_floor_elev, &wdt.Derrick_floor_elev_ouom, &wdt.Water_depth, &wdt.Water_depth_ouom, &wdt.Ground_elev, &wdt.Ground_elev_ouom, &wdt.Depth_datum_elev, &wdt.Drill_td, &wdt.Drill_td_ouom, &wdt.Log_td, &wdt.Log_td_ouom, &wdt.Max_tvd, &wdt.Max_tvd_ouom, &wdt.Projected_depth, &wdt.Projected_depth_ouom, &wdt.Final_td, &wdt.Final_td_ouom, &wdt.Operator_ba_id, &wdt.Rig_name, &wdt.Rig_type, &wdt.Test_result_code, &wdt.Remark, &wdt.Source, &wdt.Qc_status, &wdt.Checked_by_ba_id, &wdt.Tubing_obs_no, &wdt.Outside_diameter, &wdt.Outside_diameter_ouom, &wdt.Base_depth, &wdt.Base_depth_ouom); err != nil{
                return err
            }
            results = append(results, wdt)
        }
    
    return c.JSON(results)
}
func PutWellData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wdt dto.Well_data
    setDefaults(&wdt)

    if err := c.BodyParser(&wdt); err != nil{
        return err
    }
    
    wdt.Create_date = formatDateString(wdt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM well_data_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE well_data_table SET
        ba_long_name = :1, ba_type = :2, area_id = :3, area_type = :4, field_name = :5, well_name = :6, alias_long_name = :7, uwi = :8, status_type = :9, current_class = :10, well_level_type = :11, profile_type = :12, projected_strat_unit_id = :13, surface_longitude = :14, surface_latitude = :15, easting = :16, easting_ouom = :17, northing = :18, northing_ouom = :19, utm_quadrant = :20, projection_type = :21, geodetic_datum_name = :22, environment_type = :23, line_name = :24, seis_point_id = :25, spud_date = :26, final_drill_date = :27, completion_date = :28, rotary_table_elev = :29, rotary_table_elev_ouom = :30, kb_elev = :31, kb_elev_ouom = :32, derrick_floor_elev = :33, derrick_floor_elev_ouom = :34, water_depth = :35, water_depth_ouom = :36, ground_elev = :37, ground_elev_ouom = :38, depth_datum_elev = :39, drill_td = :40, drill_td_ouom = :41, log_td = :42, log_td_ouom = :43, max_tvd = :44, max_tvd_ouom = :45, projected_depth = :46, projected_depth_ouom = :47, final_td = :48, final_td_ouom = :49, operator_ba_id = :50, rig_name = :51, rig_type = :52, test_result_code = :53, remark = :54, source = :55, qc_status = :56, checked_by_ba_id = :57, tubing_obs_no = :58, outside_diameter = :59, outside_diameter_ouom = :60, base_depth = :61, base_depth_ouom = :62 WHERE id = :63`, &wdt.Ba_long_name, &wdt.Ba_type, &wdt.Area_id, &wdt.Area_type, &wdt.Field_name, &wdt.Well_name, &wdt.Alias_long_name, &wdt.Uwi, &wdt.Status_type, &wdt.Current_class, &wdt.Well_level_type, &wdt.Profile_type, &wdt.Projected_strat_unit_id, &wdt.Surface_longitude, &wdt.Surface_latitude, &wdt.Easting, &wdt.Easting_ouom, &wdt.Northing, &wdt.Northing_ouom, &wdt.Utm_quadrant, &wdt.Projection_type, &wdt.Geodetic_datum_name, &wdt.Environment_type, &wdt.Line_name, &wdt.Seis_point_id, &wdt.Spud_date, &wdt.Final_drill_date, &wdt.Completion_date, &wdt.Rotary_table_elev, &wdt.Rotary_table_elev_ouom, &wdt.Kb_elev, &wdt.Kb_elev_ouom, &wdt.Derrick_floor_elev, &wdt.Derrick_floor_elev_ouom, &wdt.Water_depth, &wdt.Water_depth_ouom, &wdt.Ground_elev, &wdt.Ground_elev_ouom, &wdt.Depth_datum_elev, &wdt.Drill_td, &wdt.Drill_td_ouom, &wdt.Log_td, &wdt.Log_td_ouom, &wdt.Max_tvd, &wdt.Max_tvd_ouom, &wdt.Projected_depth, &wdt.Projected_depth_ouom, &wdt.Final_td, &wdt.Final_td_ouom, &wdt.Operator_ba_id, &wdt.Rig_name, &wdt.Rig_type, &wdt.Test_result_code, &wdt.Remark, &wdt.Source, &wdt.Qc_status, &wdt.Checked_by_ba_id, &wdt.Tubing_obs_no, &wdt.Outside_diameter, &wdt.Outside_diameter_ouom, &wdt.Base_depth, &wdt.Base_depth_ouom, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(WELL_DATA_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteWellData(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT well_data_id FROM well_data_workspace WHERE well_data_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM well_data_workspace WHERE well_data_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM well_data_table WHERE id = :1`, id)
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
func PatchWellData(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var wdt dto.Well_data
    setDefaults(&wdt)

    if err := c.BodyParser(&wdt); err != nil{
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
    err = db.QueryRow(`SELECT well_data_id FROM well_data_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if wdt.Ba_long_name != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET ba_long_name = :1 WHERE id = :2`, wdt.Ba_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Ba_type != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET ba_type = :1 WHERE id = :2`, wdt.Ba_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Area_id != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET area_id = :1 WHERE id = :2`, wdt.Area_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Area_type != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET area_type = :1 WHERE id = :2`, wdt.Area_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Field_name != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET field_name = :1 WHERE id = :2`, wdt.Field_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Well_name != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET well_name = :1 WHERE id = :2`, wdt.Well_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Alias_long_name != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET alias_long_name = :1 WHERE id = :2`, wdt.Alias_long_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Uwi != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET uwi = :1 WHERE id = :2`, wdt.Uwi, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Status_type != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET status_type = :1 WHERE id = :2`, wdt.Status_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Current_class != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET current_class = :1 WHERE id = :2`, wdt.Current_class, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Well_level_type != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET well_level_type = :1 WHERE id = :2`, wdt.Well_level_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Profile_type != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET profile_type = :1 WHERE id = :2`, wdt.Profile_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Projected_strat_unit_id != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET projected_strat_unit_id = :1 WHERE id = :2`, wdt.Projected_strat_unit_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Surface_longitude != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET surface_longitude = :1 WHERE id = :2`, wdt.Surface_longitude, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Surface_latitude != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET surface_latitude = :1 WHERE id = :2`, wdt.Surface_latitude, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Easting != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET easting = :1 WHERE id = :2`, wdt.Easting, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Easting_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET easting_ouom = :1 WHERE id = :2`, wdt.Easting_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Northing != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET northing = :1 WHERE id = :2`, wdt.Northing, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Northing_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET northing_ouom = :1 WHERE id = :2`, wdt.Northing_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Utm_quadrant != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET utm_quadrant = :1 WHERE id = :2`, wdt.Utm_quadrant, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Projection_type != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET projection_type = :1 WHERE id = :2`, wdt.Projection_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Geodetic_datum_name != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET geodetic_datum_name = :1 WHERE id = :2`, wdt.Geodetic_datum_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Environment_type != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET environment_type = :1 WHERE id = :2`, wdt.Environment_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Line_name != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET line_name = :1 WHERE id = :2`, wdt.Line_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Seis_point_id != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET seis_point_id = :1 WHERE id = :2`, wdt.Seis_point_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Spud_date != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET spud_date = :1 WHERE id = :2`, wdt.Spud_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Final_drill_date != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET final_drill_date = :1 WHERE id = :2`, wdt.Final_drill_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Completion_date != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET completion_date = :1 WHERE id = :2`, wdt.Completion_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Rotary_table_elev != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET rotary_table_elev = :1 WHERE id = :2`, wdt.Rotary_table_elev, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Rotary_table_elev_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET rotary_table_elev_ouom = :1 WHERE id = :2`, wdt.Rotary_table_elev_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Kb_elev != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET kb_elev = :1 WHERE id = :2`, wdt.Kb_elev, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Kb_elev_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET kb_elev_ouom = :1 WHERE id = :2`, wdt.Kb_elev_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Derrick_floor_elev != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET derrick_floor_elev = :1 WHERE id = :2`, wdt.Derrick_floor_elev, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Derrick_floor_elev_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET derrick_floor_elev_ouom = :1 WHERE id = :2`, wdt.Derrick_floor_elev_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Water_depth != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET water_depth = :1 WHERE id = :2`, wdt.Water_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Water_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET water_depth_ouom = :1 WHERE id = :2`, wdt.Water_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Ground_elev != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET ground_elev = :1 WHERE id = :2`, wdt.Ground_elev, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Ground_elev_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET ground_elev_ouom = :1 WHERE id = :2`, wdt.Ground_elev_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Depth_datum_elev != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET depth_datum_elev = :1 WHERE id = :2`, wdt.Depth_datum_elev, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Drill_td != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET drill_td = :1 WHERE id = :2`, wdt.Drill_td, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Drill_td_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET drill_td_ouom = :1 WHERE id = :2`, wdt.Drill_td_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Log_td != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET log_td = :1 WHERE id = :2`, wdt.Log_td, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Log_td_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET log_td_ouom = :1 WHERE id = :2`, wdt.Log_td_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Max_tvd != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET max_tvd = :1 WHERE id = :2`, wdt.Max_tvd, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Max_tvd_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET max_tvd_ouom = :1 WHERE id = :2`, wdt.Max_tvd_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Projected_depth != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET projected_depth = :1 WHERE id = :2`, wdt.Projected_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Projected_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET projected_depth_ouom = :1 WHERE id = :2`, wdt.Projected_depth_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Final_td != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET final_td = :1 WHERE id = :2`, wdt.Final_td, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Final_td_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET final_td_ouom = :1 WHERE id = :2`, wdt.Final_td_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Operator_ba_id != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET operator_ba_id = :1 WHERE id = :2`, wdt.Operator_ba_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Rig_name != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET rig_name = :1 WHERE id = :2`, wdt.Rig_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Rig_type != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET rig_type = :1 WHERE id = :2`, wdt.Rig_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Test_result_code != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET test_result_code = :1 WHERE id = :2`, wdt.Test_result_code, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Remark != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET remark = :1 WHERE id = :2`, wdt.Remark, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Source != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET source = :1 WHERE id = :2`, wdt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET qc_status = :1 WHERE id = :2`, wdt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET checked_by_ba_id = :1 WHERE id = :2`, wdt.Checked_by_ba_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Tubing_obs_no != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET tubing_obs_no = :1 WHERE id = :2`, wdt.Tubing_obs_no, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Outside_diameter != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET outside_diameter = :1 WHERE id = :2`, wdt.Outside_diameter, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Outside_diameter_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET outside_diameter_ouom = :1 WHERE id = :2`, wdt.Outside_diameter_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Base_depth != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET base_depth = :1 WHERE id = :2`, wdt.Base_depth, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if wdt.Base_depth_ouom != nil{
             _, err = tx.Exec(`UPDATE well_data_table SET base_depth_ouom = :1 WHERE id = :2`, wdt.Base_depth_ouom, id)
        
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

    