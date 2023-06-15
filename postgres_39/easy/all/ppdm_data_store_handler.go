package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmDataStore(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_data_store")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_data_store

	for rows.Next() {
		var ppdm_data_store dto.Ppdm_data_store
		if err := rows.Scan(&ppdm_data_store.Uom_system, &ppdm_data_store.Name, &ppdm_data_store.Active_ind, &ppdm_data_store.Effective_date, &ppdm_data_store.Expiry_date, &ppdm_data_store.Ppdm_guid, &ppdm_data_store.Remark, &ppdm_data_store.Source, &ppdm_data_store.Row_changed_by, &ppdm_data_store.Row_changed_date, &ppdm_data_store.Row_created_by, &ppdm_data_store.Row_created_date, &ppdm_data_store.Row_effective_date, &ppdm_data_store.Row_expiry_date, &ppdm_data_store.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_data_store to result
		result = append(result, ppdm_data_store)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmDataStore(c *fiber.Ctx) error {
	var ppdm_data_store dto.Ppdm_data_store

	setDefaults(&ppdm_data_store)

	if err := json.Unmarshal(c.Body(), &ppdm_data_store); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_data_store values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	ppdm_data_store.Row_created_date = formatDate(ppdm_data_store.Row_created_date)
	ppdm_data_store.Row_changed_date = formatDate(ppdm_data_store.Row_changed_date)
	ppdm_data_store.Effective_date = formatDateString(ppdm_data_store.Effective_date)
	ppdm_data_store.Expiry_date = formatDateString(ppdm_data_store.Expiry_date)
	ppdm_data_store.Row_effective_date = formatDateString(ppdm_data_store.Row_effective_date)
	ppdm_data_store.Row_expiry_date = formatDateString(ppdm_data_store.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_data_store.Uom_system, ppdm_data_store.Name, ppdm_data_store.Active_ind, ppdm_data_store.Effective_date, ppdm_data_store.Expiry_date, ppdm_data_store.Ppdm_guid, ppdm_data_store.Remark, ppdm_data_store.Source, ppdm_data_store.Row_changed_by, ppdm_data_store.Row_changed_date, ppdm_data_store.Row_created_by, ppdm_data_store.Row_created_date, ppdm_data_store.Row_effective_date, ppdm_data_store.Row_expiry_date, ppdm_data_store.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmDataStore(c *fiber.Ctx) error {
	var ppdm_data_store dto.Ppdm_data_store

	setDefaults(&ppdm_data_store)

	if err := json.Unmarshal(c.Body(), &ppdm_data_store); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_data_store.Uom_system = id

    if ppdm_data_store.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_data_store set name = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where uom_system = :15")
	if err != nil {
		return err
	}

	ppdm_data_store.Row_changed_date = formatDate(ppdm_data_store.Row_changed_date)
	ppdm_data_store.Effective_date = formatDateString(ppdm_data_store.Effective_date)
	ppdm_data_store.Expiry_date = formatDateString(ppdm_data_store.Expiry_date)
	ppdm_data_store.Row_effective_date = formatDateString(ppdm_data_store.Row_effective_date)
	ppdm_data_store.Row_expiry_date = formatDateString(ppdm_data_store.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_data_store.Name, ppdm_data_store.Active_ind, ppdm_data_store.Effective_date, ppdm_data_store.Expiry_date, ppdm_data_store.Ppdm_guid, ppdm_data_store.Remark, ppdm_data_store.Source, ppdm_data_store.Row_changed_by, ppdm_data_store.Row_changed_date, ppdm_data_store.Row_created_by, ppdm_data_store.Row_effective_date, ppdm_data_store.Row_expiry_date, ppdm_data_store.Row_quality, ppdm_data_store.Uom_system)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchPpdmDataStore(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_data_store set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uom_system = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeletePpdmDataStore(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_data_store dto.Ppdm_data_store
	ppdm_data_store.Uom_system = id

	stmt, err := db.Prepare("delete from ppdm_data_store where uom_system = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_data_store.Uom_system)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


