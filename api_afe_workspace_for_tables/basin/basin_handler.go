package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-basin/dto"
    "github.com/DarrenMannuela/svc-datatype-basin/utils"

    "github.com/gofiber/fiber/v2"
)
func GetBasin(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM basin_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Basin    
    
        for rows.Next() {
    
            var bt dto.Basin
            if err := rows.Scan(&bt.Id, &bt.Project_name, &bt.Strat_name_set_id, &bt.Strat_status, &bt.Product_type, &bt.Reserve_class_id, &bt.Open_balance, &bt.Open_balance_ouom, &bt.Size_type, &bt.Gross_size, &bt.Size_ouom, &bt.Strat_type, &bt.Fault_type, &bt.Source, &bt.Qc_status, &bt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, bt)
        }
    
    return c.JSON(results)
}
func SetBasin(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var bt dto.Basin
    setDefaults(&bt)

    if err := c.BodyParser(&bt); err != nil{
        return err
    }
    
    bt.Create_date = formatDateString(bt.Create_date)
    

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
    _, err = tx.Exec(`INSERT INTO basin_table (project_name, strat_name_set_id, strat_status, product_type, reserve_class_id, open_balance, open_balance_ouom, size_type, gross_size, size_ouom, strat_type, fault_type, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15) RETURNING id INTO :16`, &bt.Project_name, &bt.Strat_name_set_id, &bt.Strat_status, &bt.Product_type, &bt.Reserve_class_id, &bt.Open_balance, &bt.Open_balance_ouom, &bt.Size_type, &bt.Gross_size, &bt.Size_ouom, &bt.Strat_type, &bt.Fault_type, &bt.Source, &bt.Qc_status, &bt.Checked_by_ba_id, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(BASIN_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetBasinById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM basin_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Basin    
    
        for rows.Next() {
    
            var bt dto.Basin
            if err := rows.Scan(&bt.Id, &bt.Project_name, &bt.Strat_name_set_id, &bt.Strat_status, &bt.Product_type, &bt.Reserve_class_id, &bt.Open_balance, &bt.Open_balance_ouom, &bt.Size_type, &bt.Gross_size, &bt.Size_ouom, &bt.Strat_type, &bt.Fault_type, &bt.Source, &bt.Qc_status, &bt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, bt)
        }
    
    return c.JSON(results)
}
func PutBasin(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var bt dto.Basin
    setDefaults(&bt)

    if err := c.BodyParser(&bt); err != nil{
        return err
    }
    
    bt.Create_date = formatDateString(bt.Create_date)
    

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
    err = db.QueryRow(`SELECT * FROM basin_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE basin_table SET
        project_name = :1, strat_name_set_id = :2, strat_status = :3, product_type = :4, reserve_class_id = :5, open_balance = :6, open_balance_ouom = :7, size_type = :8, gross_size = :9, size_ouom = :10, strat_type = :11, fault_type = :12, source = :13, qc_status = :14, checked_by_ba_id = :15 WHERE id = :16`, &bt.Project_name, &bt.Strat_name_set_id, &bt.Strat_status, &bt.Product_type, &bt.Reserve_class_id, &bt.Open_balance, &bt.Open_balance_ouom, &bt.Size_type, &bt.Gross_size, &bt.Size_ouom, &bt.Strat_type, &bt.Fault_type, &bt.Source, &bt.Qc_status, &bt.Checked_by_ba_id, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(BASIN_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteBasin(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT basin_id FROM basin_workspace WHERE basin_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM basin_workspace WHERE basin_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM basin_table WHERE id = :1`, id)
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
func PatchBasin(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var bt dto.Basin
    setDefaults(&bt)

    if err := c.BodyParser(&bt); err != nil{
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
    err = db.QueryRow(`SELECT basin_id FROM basin_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if bt.Project_name != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET project_name = :1 WHERE id = :2`, bt.Project_name, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Strat_name_set_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET strat_name_set_id = :1 WHERE id = :2`, bt.Strat_name_set_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Strat_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET strat_status = :1 WHERE id = :2`, bt.Strat_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Product_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET product_type = :1 WHERE id = :2`, bt.Product_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Reserve_class_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET reserve_class_id = :1 WHERE id = :2`, bt.Reserve_class_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Open_balance != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET open_balance = :1 WHERE id = :2`, bt.Open_balance, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Open_balance_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET open_balance_ouom = :1 WHERE id = :2`, bt.Open_balance_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Size_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET size_type = :1 WHERE id = :2`, bt.Size_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Gross_size != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET gross_size = :1 WHERE id = :2`, bt.Gross_size, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Size_ouom != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET size_ouom = :1 WHERE id = :2`, bt.Size_ouom, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Strat_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET strat_type = :1 WHERE id = :2`, bt.Strat_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Fault_type != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET fault_type = :1 WHERE id = :2`, bt.Fault_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Source != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET source = :1 WHERE id = :2`, bt.Source, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Qc_status != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET qc_status = :1 WHERE id = :2`, bt.Qc_status, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Checked_by_ba_id != nil{
             _, err = tx.Exec(`UPDATE non_seismic_and_seismic_non_conventional_report_table SET checked_by_ba_id = :1 WHERE id = :2`, bt.Checked_by_ba_id, id)
        
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

    