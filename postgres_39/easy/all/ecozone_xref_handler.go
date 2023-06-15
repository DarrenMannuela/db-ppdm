package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEcozoneXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ecozone_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ecozone_xref

	for rows.Next() {
		var ecozone_xref dto.Ecozone_xref
		if err := rows.Scan(&ecozone_xref.Ecozone_id1, &ecozone_xref.Ecozone_id2, &ecozone_xref.Ecozone_xref_id, &ecozone_xref.Active_ind, &ecozone_xref.Ecozone_xref_type, &ecozone_xref.Effective_date, &ecozone_xref.Expiry_date, &ecozone_xref.Ppdm_guid, &ecozone_xref.Remark, &ecozone_xref.Source, &ecozone_xref.Row_changed_by, &ecozone_xref.Row_changed_date, &ecozone_xref.Row_created_by, &ecozone_xref.Row_created_date, &ecozone_xref.Row_effective_date, &ecozone_xref.Row_expiry_date, &ecozone_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ecozone_xref to result
		result = append(result, ecozone_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEcozoneXref(c *fiber.Ctx) error {
	var ecozone_xref dto.Ecozone_xref

	setDefaults(&ecozone_xref)

	if err := json.Unmarshal(c.Body(), &ecozone_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ecozone_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	ecozone_xref.Row_created_date = formatDate(ecozone_xref.Row_created_date)
	ecozone_xref.Row_changed_date = formatDate(ecozone_xref.Row_changed_date)
	ecozone_xref.Effective_date = formatDateString(ecozone_xref.Effective_date)
	ecozone_xref.Expiry_date = formatDateString(ecozone_xref.Expiry_date)
	ecozone_xref.Row_effective_date = formatDateString(ecozone_xref.Row_effective_date)
	ecozone_xref.Row_expiry_date = formatDateString(ecozone_xref.Row_expiry_date)






	rows, err := stmt.Exec(ecozone_xref.Ecozone_id1, ecozone_xref.Ecozone_id2, ecozone_xref.Ecozone_xref_id, ecozone_xref.Active_ind, ecozone_xref.Ecozone_xref_type, ecozone_xref.Effective_date, ecozone_xref.Expiry_date, ecozone_xref.Ppdm_guid, ecozone_xref.Remark, ecozone_xref.Source, ecozone_xref.Row_changed_by, ecozone_xref.Row_changed_date, ecozone_xref.Row_created_by, ecozone_xref.Row_created_date, ecozone_xref.Row_effective_date, ecozone_xref.Row_expiry_date, ecozone_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEcozoneXref(c *fiber.Ctx) error {
	var ecozone_xref dto.Ecozone_xref

	setDefaults(&ecozone_xref)

	if err := json.Unmarshal(c.Body(), &ecozone_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ecozone_xref.Ecozone_id1 = id

    if ecozone_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ecozone_xref set ecozone_id2 = :1, ecozone_xref_id = :2, active_ind = :3, ecozone_xref_type = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where ecozone_id1 = :17")
	if err != nil {
		return err
	}

	ecozone_xref.Row_changed_date = formatDate(ecozone_xref.Row_changed_date)
	ecozone_xref.Effective_date = formatDateString(ecozone_xref.Effective_date)
	ecozone_xref.Expiry_date = formatDateString(ecozone_xref.Expiry_date)
	ecozone_xref.Row_effective_date = formatDateString(ecozone_xref.Row_effective_date)
	ecozone_xref.Row_expiry_date = formatDateString(ecozone_xref.Row_expiry_date)






	rows, err := stmt.Exec(ecozone_xref.Ecozone_id2, ecozone_xref.Ecozone_xref_id, ecozone_xref.Active_ind, ecozone_xref.Ecozone_xref_type, ecozone_xref.Effective_date, ecozone_xref.Expiry_date, ecozone_xref.Ppdm_guid, ecozone_xref.Remark, ecozone_xref.Source, ecozone_xref.Row_changed_by, ecozone_xref.Row_changed_date, ecozone_xref.Row_created_by, ecozone_xref.Row_effective_date, ecozone_xref.Row_expiry_date, ecozone_xref.Row_quality, ecozone_xref.Ecozone_id1)
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

func PatchEcozoneXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ecozone_xref set "
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
	query += " where ecozone_id1 = :id"

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

func DeleteEcozoneXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var ecozone_xref dto.Ecozone_xref
	ecozone_xref.Ecozone_id1 = id

	stmt, err := db.Prepare("delete from ecozone_xref where ecozone_id1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ecozone_xref.Ecozone_id1)
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


