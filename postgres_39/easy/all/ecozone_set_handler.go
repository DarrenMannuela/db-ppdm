package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEcozoneSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ecozone_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ecozone_set

	for rows.Next() {
		var ecozone_set dto.Ecozone_set
		if err := rows.Scan(&ecozone_set.Ecozone_set_id, &ecozone_set.Active_ind, &ecozone_set.Effective_date, &ecozone_set.Expiry_date, &ecozone_set.Owner_ba_id, &ecozone_set.Ppdm_guid, &ecozone_set.Remark, &ecozone_set.Source, &ecozone_set.Source_document_id, &ecozone_set.Row_changed_by, &ecozone_set.Row_changed_date, &ecozone_set.Row_created_by, &ecozone_set.Row_created_date, &ecozone_set.Row_effective_date, &ecozone_set.Row_expiry_date, &ecozone_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ecozone_set to result
		result = append(result, ecozone_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEcozoneSet(c *fiber.Ctx) error {
	var ecozone_set dto.Ecozone_set

	setDefaults(&ecozone_set)

	if err := json.Unmarshal(c.Body(), &ecozone_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ecozone_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	ecozone_set.Row_created_date = formatDate(ecozone_set.Row_created_date)
	ecozone_set.Row_changed_date = formatDate(ecozone_set.Row_changed_date)
	ecozone_set.Effective_date = formatDateString(ecozone_set.Effective_date)
	ecozone_set.Expiry_date = formatDateString(ecozone_set.Expiry_date)
	ecozone_set.Row_effective_date = formatDateString(ecozone_set.Row_effective_date)
	ecozone_set.Row_expiry_date = formatDateString(ecozone_set.Row_expiry_date)






	rows, err := stmt.Exec(ecozone_set.Ecozone_set_id, ecozone_set.Active_ind, ecozone_set.Effective_date, ecozone_set.Expiry_date, ecozone_set.Owner_ba_id, ecozone_set.Ppdm_guid, ecozone_set.Remark, ecozone_set.Source, ecozone_set.Source_document_id, ecozone_set.Row_changed_by, ecozone_set.Row_changed_date, ecozone_set.Row_created_by, ecozone_set.Row_created_date, ecozone_set.Row_effective_date, ecozone_set.Row_expiry_date, ecozone_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEcozoneSet(c *fiber.Ctx) error {
	var ecozone_set dto.Ecozone_set

	setDefaults(&ecozone_set)

	if err := json.Unmarshal(c.Body(), &ecozone_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ecozone_set.Ecozone_set_id = id

    if ecozone_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ecozone_set set active_ind = :1, effective_date = :2, expiry_date = :3, owner_ba_id = :4, ppdm_guid = :5, remark = :6, source = :7, source_document_id = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where ecozone_set_id = :16")
	if err != nil {
		return err
	}

	ecozone_set.Row_changed_date = formatDate(ecozone_set.Row_changed_date)
	ecozone_set.Effective_date = formatDateString(ecozone_set.Effective_date)
	ecozone_set.Expiry_date = formatDateString(ecozone_set.Expiry_date)
	ecozone_set.Row_effective_date = formatDateString(ecozone_set.Row_effective_date)
	ecozone_set.Row_expiry_date = formatDateString(ecozone_set.Row_expiry_date)






	rows, err := stmt.Exec(ecozone_set.Active_ind, ecozone_set.Effective_date, ecozone_set.Expiry_date, ecozone_set.Owner_ba_id, ecozone_set.Ppdm_guid, ecozone_set.Remark, ecozone_set.Source, ecozone_set.Source_document_id, ecozone_set.Row_changed_by, ecozone_set.Row_changed_date, ecozone_set.Row_created_by, ecozone_set.Row_effective_date, ecozone_set.Row_expiry_date, ecozone_set.Row_quality, ecozone_set.Ecozone_set_id)
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

func PatchEcozoneSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ecozone_set set "
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
	query += " where ecozone_set_id = :id"

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

func DeleteEcozoneSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var ecozone_set dto.Ecozone_set
	ecozone_set.Ecozone_set_id = id

	stmt, err := db.Prepare("delete from ecozone_set where ecozone_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ecozone_set.Ecozone_set_id)
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


