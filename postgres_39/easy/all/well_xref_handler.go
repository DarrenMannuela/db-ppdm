package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_xref

	for rows.Next() {
		var well_xref dto.Well_xref
		if err := rows.Scan(&well_xref.Uwi, &well_xref.Xref_id, &well_xref.Active_ind, &well_xref.Effective_date, &well_xref.Expiry_date, &well_xref.Ppdm_guid, &well_xref.Remark, &well_xref.Source, &well_xref.Uwi2, &well_xref.Xref_owner_ba_id, &well_xref.Xref_type, &well_xref.Row_changed_by, &well_xref.Row_changed_date, &well_xref.Row_created_by, &well_xref.Row_created_date, &well_xref.Row_effective_date, &well_xref.Row_expiry_date, &well_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_xref to result
		result = append(result, well_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellXref(c *fiber.Ctx) error {
	var well_xref dto.Well_xref

	setDefaults(&well_xref)

	if err := json.Unmarshal(c.Body(), &well_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	well_xref.Row_created_date = formatDate(well_xref.Row_created_date)
	well_xref.Row_changed_date = formatDate(well_xref.Row_changed_date)
	well_xref.Effective_date = formatDateString(well_xref.Effective_date)
	well_xref.Expiry_date = formatDateString(well_xref.Expiry_date)
	well_xref.Row_effective_date = formatDateString(well_xref.Row_effective_date)
	well_xref.Row_expiry_date = formatDateString(well_xref.Row_expiry_date)






	rows, err := stmt.Exec(well_xref.Uwi, well_xref.Xref_id, well_xref.Active_ind, well_xref.Effective_date, well_xref.Expiry_date, well_xref.Ppdm_guid, well_xref.Remark, well_xref.Source, well_xref.Uwi2, well_xref.Xref_owner_ba_id, well_xref.Xref_type, well_xref.Row_changed_by, well_xref.Row_changed_date, well_xref.Row_created_by, well_xref.Row_created_date, well_xref.Row_effective_date, well_xref.Row_expiry_date, well_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellXref(c *fiber.Ctx) error {
	var well_xref dto.Well_xref

	setDefaults(&well_xref)

	if err := json.Unmarshal(c.Body(), &well_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_xref.Uwi = id

    if well_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_xref set xref_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, uwi2 = :8, xref_owner_ba_id = :9, xref_type = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where uwi = :18")
	if err != nil {
		return err
	}

	well_xref.Row_changed_date = formatDate(well_xref.Row_changed_date)
	well_xref.Effective_date = formatDateString(well_xref.Effective_date)
	well_xref.Expiry_date = formatDateString(well_xref.Expiry_date)
	well_xref.Row_effective_date = formatDateString(well_xref.Row_effective_date)
	well_xref.Row_expiry_date = formatDateString(well_xref.Row_expiry_date)






	rows, err := stmt.Exec(well_xref.Xref_id, well_xref.Active_ind, well_xref.Effective_date, well_xref.Expiry_date, well_xref.Ppdm_guid, well_xref.Remark, well_xref.Source, well_xref.Uwi2, well_xref.Xref_owner_ba_id, well_xref.Xref_type, well_xref.Row_changed_by, well_xref.Row_changed_date, well_xref.Row_created_by, well_xref.Row_effective_date, well_xref.Row_expiry_date, well_xref.Row_quality, well_xref.Uwi)
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

func PatchWellXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_xref set "
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
	query += " where uwi = :id"

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

func DeleteWellXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_xref dto.Well_xref
	well_xref.Uwi = id

	stmt, err := db.Prepare("delete from well_xref where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_xref.Uwi)
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


