package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-bibliography/dto"
    "github.com/DarrenMannuela/svc-datatype-bibliography/utils"

    "github.com/gofiber/fiber/v2"
)
func GetBibliography(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query(`SELECT * FROM bibliography_table`)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Bibliography    
    
        for rows.Next() {
    
            var bt dto.Bibliography
            if err := rows.Scan(&bt.Id, &bt.Publisher, &bt.Document_title, &bt.Issue, &bt.Author_id, &bt.Publication_date, &bt.Document_type, &bt.Data_store_name); err != nil{
                return err
            }
            results = append(results, bt)
        }
    
    return c.JSON(results)
}
func SetBibliography(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var bt dto.Bibliography
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
    _, err = tx.Exec(`INSERT INTO bibliography_table (publisher, document_title, issue, author_id, publication_date, document_type, data_store_name) VALUES (:1, :2, :3, :4, :5, :6, :7) RETURNING id INTO :8`, &bt.Publisher, &bt.Document_title, &bt.Issue, &bt.Author_id, &bt.Publication_date, &bt.Document_type, &bt.Data_store_name, sql.Out{Dest: &generatedID})
    if err != nil {
        tx.Rollback()
        fmt.Println(BIBLIOGRAPHY_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetBibliographyById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query(`SELECT * FROM bibliography_table WHERE id = :1`, id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Bibliography    
    
        for rows.Next() {
    
            var bt dto.Bibliography
            if err := rows.Scan(&bt.Id, &bt.Publisher, &bt.Document_title, &bt.Issue, &bt.Author_id, &bt.Publication_date, &bt.Document_type, &bt.Data_store_name); err != nil{
                return err
            }
            results = append(results, bt)
        }
    
    return c.JSON(results)
}
func PutBibliography(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var bt dto.Bibliography
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
    err = db.QueryRow(`SELECT * FROM bibliography_table WHERE id = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != "" {

    
        _, err = tx.Exec(`UPDATE bibliography_table SET
        publisher = :1, document_title = :2, issue = :3, author_id = :4, publication_date = :5, document_type = :6, data_store_name = :7 WHERE id = :8`, &bt.Publisher, &bt.Document_title, &bt.Issue, &bt.Author_id, &bt.Publication_date, &bt.Document_type, &bt.Data_store_name, id)
        if err != nil {
            tx.Rollback()
            fmt.Println(BIBLIOGRAPHY_TABLE)
            return err
        }
    
    }

    if err := tx.Commit(); err != nil {
		tx.Rollback()
		return c.Status(500).SendString(err.Error())
	}
    
    
    return c.SendStatus(fiber.StatusOK)
}
func DeleteBibliography(c *fiber.Ctx) error{
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
	_ = tx.QueryRow(`SELECT bibliography_id FROM bibliography_workspace WHERE bibliography_id = :1`, id).Scan(&idExist)
	if idExist != "" {
		_, err = tx.Exec(`DELETE FROM bibliography_workspace WHERE bibliography_id = :1`, id)
		if err != nil{
			tx.Rollback()
			return c.Status(500).SendString(err.Error())
		}

	}
    
    _, err = tx.Exec(`DELETE FROM bibliography_table WHERE id = :1`, id)
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
func PatchBibliography(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    var bt dto.Bibliography
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
    err = db.QueryRow(`SELECT bibliography_id FROM bibliography_table WHERE afe_number = :1`, id).Scan(&idExist)
    if err != nil{
        tx.Rollback()
		return c.Status(400).SendString("ID does not exist")
    }
    
    if idExist != ""{
        if bt.Publisher != nil{
             _, err = tx.Exec(`UPDATE bibliography_table SET publisher = :1 WHERE id = :2`, bt.Publisher, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Document_title != nil{
             _, err = tx.Exec(`UPDATE bibliography_table SET document_title = :1 WHERE id = :2`, bt.Document_title, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Issue != nil{
             _, err = tx.Exec(`UPDATE bibliography_table SET issue = :1 WHERE id = :2`, bt.Issue, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Author_id != nil{
             _, err = tx.Exec(`UPDATE bibliography_table SET author_id = :1 WHERE id = :2`, bt.Author_id, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Publication_date != nil{
             _, err = tx.Exec(`UPDATE bibliography_table SET publication_date = :1 WHERE id = :2`, bt.Publication_date, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Document_type != nil{
             _, err = tx.Exec(`UPDATE bibliography_table SET document_type = :1 WHERE id = :2`, bt.Document_type, id)
        
            if err != nil {
                tx.Rollback()
                return c.Status(500).SendString(err.Error())
            }
        }
            
        if bt.Data_store_name != nil{
             _, err = tx.Exec(`UPDATE bibliography_table SET data_store_name = :1 WHERE id = :2`, bt.Data_store_name, id)
        
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

    