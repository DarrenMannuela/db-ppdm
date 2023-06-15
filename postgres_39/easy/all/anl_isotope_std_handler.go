package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlIsotopeStd(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_isotope_std")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_isotope_std

	for rows.Next() {
		var anl_isotope_std dto.Anl_isotope_std
		if err := rows.Scan(&anl_isotope_std.Standard_id, &anl_isotope_std.Substance_id, &anl_isotope_std.Active_ind, &anl_isotope_std.Effective_date, &anl_isotope_std.Expiry_date, &anl_isotope_std.Ppdm_guid, &anl_isotope_std.Remark, &anl_isotope_std.Source, &anl_isotope_std.Source_document_id, &anl_isotope_std.Standard_name, &anl_isotope_std.Standard_value, &anl_isotope_std.Standard_value_error, &anl_isotope_std.Standard_value_ouom, &anl_isotope_std.Standard_value_uom, &anl_isotope_std.Row_changed_by, &anl_isotope_std.Row_changed_date, &anl_isotope_std.Row_created_by, &anl_isotope_std.Row_created_date, &anl_isotope_std.Row_effective_date, &anl_isotope_std.Row_expiry_date, &anl_isotope_std.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_isotope_std to result
		result = append(result, anl_isotope_std)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlIsotopeStd(c *fiber.Ctx) error {
	var anl_isotope_std dto.Anl_isotope_std

	setDefaults(&anl_isotope_std)

	if err := json.Unmarshal(c.Body(), &anl_isotope_std); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_isotope_std values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	anl_isotope_std.Row_created_date = formatDate(anl_isotope_std.Row_created_date)
	anl_isotope_std.Row_changed_date = formatDate(anl_isotope_std.Row_changed_date)
	anl_isotope_std.Effective_date = formatDateString(anl_isotope_std.Effective_date)
	anl_isotope_std.Expiry_date = formatDateString(anl_isotope_std.Expiry_date)
	anl_isotope_std.Row_effective_date = formatDateString(anl_isotope_std.Row_effective_date)
	anl_isotope_std.Row_expiry_date = formatDateString(anl_isotope_std.Row_expiry_date)






	rows, err := stmt.Exec(anl_isotope_std.Standard_id, anl_isotope_std.Substance_id, anl_isotope_std.Active_ind, anl_isotope_std.Effective_date, anl_isotope_std.Expiry_date, anl_isotope_std.Ppdm_guid, anl_isotope_std.Remark, anl_isotope_std.Source, anl_isotope_std.Source_document_id, anl_isotope_std.Standard_name, anl_isotope_std.Standard_value, anl_isotope_std.Standard_value_error, anl_isotope_std.Standard_value_ouom, anl_isotope_std.Standard_value_uom, anl_isotope_std.Row_changed_by, anl_isotope_std.Row_changed_date, anl_isotope_std.Row_created_by, anl_isotope_std.Row_created_date, anl_isotope_std.Row_effective_date, anl_isotope_std.Row_expiry_date, anl_isotope_std.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlIsotopeStd(c *fiber.Ctx) error {
	var anl_isotope_std dto.Anl_isotope_std

	setDefaults(&anl_isotope_std)

	if err := json.Unmarshal(c.Body(), &anl_isotope_std); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_isotope_std.Standard_id = id

    if anl_isotope_std.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_isotope_std set substance_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, source_document_id = :8, standard_name = :9, standard_value = :10, standard_value_error = :11, standard_value_ouom = :12, standard_value_uom = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where standard_id = :21")
	if err != nil {
		return err
	}

	anl_isotope_std.Row_changed_date = formatDate(anl_isotope_std.Row_changed_date)
	anl_isotope_std.Effective_date = formatDateString(anl_isotope_std.Effective_date)
	anl_isotope_std.Expiry_date = formatDateString(anl_isotope_std.Expiry_date)
	anl_isotope_std.Row_effective_date = formatDateString(anl_isotope_std.Row_effective_date)
	anl_isotope_std.Row_expiry_date = formatDateString(anl_isotope_std.Row_expiry_date)






	rows, err := stmt.Exec(anl_isotope_std.Substance_id, anl_isotope_std.Active_ind, anl_isotope_std.Effective_date, anl_isotope_std.Expiry_date, anl_isotope_std.Ppdm_guid, anl_isotope_std.Remark, anl_isotope_std.Source, anl_isotope_std.Source_document_id, anl_isotope_std.Standard_name, anl_isotope_std.Standard_value, anl_isotope_std.Standard_value_error, anl_isotope_std.Standard_value_ouom, anl_isotope_std.Standard_value_uom, anl_isotope_std.Row_changed_by, anl_isotope_std.Row_changed_date, anl_isotope_std.Row_created_by, anl_isotope_std.Row_effective_date, anl_isotope_std.Row_expiry_date, anl_isotope_std.Row_quality, anl_isotope_std.Standard_id)
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

func PatchAnlIsotopeStd(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_isotope_std set "
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
	query += " where standard_id = :id"

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

func DeleteAnlIsotopeStd(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_isotope_std dto.Anl_isotope_std
	anl_isotope_std.Standard_id = id

	stmt, err := db.Prepare("delete from anl_isotope_std where standard_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_isotope_std.Standard_id)
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


