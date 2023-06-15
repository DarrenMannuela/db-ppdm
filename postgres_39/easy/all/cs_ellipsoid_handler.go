package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsEllipsoid(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_ellipsoid")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_ellipsoid

	for rows.Next() {
		var cs_ellipsoid dto.Cs_ellipsoid
		if err := rows.Scan(&cs_ellipsoid.Ellipsoid_id, &cs_ellipsoid.Active_ind, &cs_ellipsoid.Axis_ouom, &cs_ellipsoid.Eccentricity, &cs_ellipsoid.Effective_date, &cs_ellipsoid.Ellipsoid_name, &cs_ellipsoid.Expiry_date, &cs_ellipsoid.Inverse_flattening, &cs_ellipsoid.Ppdm_guid, &cs_ellipsoid.Remark, &cs_ellipsoid.Semi_major_axis, &cs_ellipsoid.Semi_major_axis_accuracy, &cs_ellipsoid.Semi_minor_axis, &cs_ellipsoid.Semi_minor_axis_accuracy, &cs_ellipsoid.Source, &cs_ellipsoid.Source_document_id, &cs_ellipsoid.Row_changed_by, &cs_ellipsoid.Row_changed_date, &cs_ellipsoid.Row_created_by, &cs_ellipsoid.Row_created_date, &cs_ellipsoid.Row_effective_date, &cs_ellipsoid.Row_expiry_date, &cs_ellipsoid.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_ellipsoid to result
		result = append(result, cs_ellipsoid)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsEllipsoid(c *fiber.Ctx) error {
	var cs_ellipsoid dto.Cs_ellipsoid

	setDefaults(&cs_ellipsoid)

	if err := json.Unmarshal(c.Body(), &cs_ellipsoid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_ellipsoid values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	cs_ellipsoid.Row_created_date = formatDate(cs_ellipsoid.Row_created_date)
	cs_ellipsoid.Row_changed_date = formatDate(cs_ellipsoid.Row_changed_date)
	cs_ellipsoid.Effective_date = formatDateString(cs_ellipsoid.Effective_date)
	cs_ellipsoid.Expiry_date = formatDateString(cs_ellipsoid.Expiry_date)
	cs_ellipsoid.Row_effective_date = formatDateString(cs_ellipsoid.Row_effective_date)
	cs_ellipsoid.Row_expiry_date = formatDateString(cs_ellipsoid.Row_expiry_date)






	rows, err := stmt.Exec(cs_ellipsoid.Ellipsoid_id, cs_ellipsoid.Active_ind, cs_ellipsoid.Axis_ouom, cs_ellipsoid.Eccentricity, cs_ellipsoid.Effective_date, cs_ellipsoid.Ellipsoid_name, cs_ellipsoid.Expiry_date, cs_ellipsoid.Inverse_flattening, cs_ellipsoid.Ppdm_guid, cs_ellipsoid.Remark, cs_ellipsoid.Semi_major_axis, cs_ellipsoid.Semi_major_axis_accuracy, cs_ellipsoid.Semi_minor_axis, cs_ellipsoid.Semi_minor_axis_accuracy, cs_ellipsoid.Source, cs_ellipsoid.Source_document_id, cs_ellipsoid.Row_changed_by, cs_ellipsoid.Row_changed_date, cs_ellipsoid.Row_created_by, cs_ellipsoid.Row_created_date, cs_ellipsoid.Row_effective_date, cs_ellipsoid.Row_expiry_date, cs_ellipsoid.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsEllipsoid(c *fiber.Ctx) error {
	var cs_ellipsoid dto.Cs_ellipsoid

	setDefaults(&cs_ellipsoid)

	if err := json.Unmarshal(c.Body(), &cs_ellipsoid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_ellipsoid.Ellipsoid_id = id

    if cs_ellipsoid.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_ellipsoid set active_ind = :1, axis_ouom = :2, eccentricity = :3, effective_date = :4, ellipsoid_name = :5, expiry_date = :6, inverse_flattening = :7, ppdm_guid = :8, remark = :9, semi_major_axis = :10, semi_major_axis_accuracy = :11, semi_minor_axis = :12, semi_minor_axis_accuracy = :13, source = :14, source_document_id = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where ellipsoid_id = :23")
	if err != nil {
		return err
	}

	cs_ellipsoid.Row_changed_date = formatDate(cs_ellipsoid.Row_changed_date)
	cs_ellipsoid.Effective_date = formatDateString(cs_ellipsoid.Effective_date)
	cs_ellipsoid.Expiry_date = formatDateString(cs_ellipsoid.Expiry_date)
	cs_ellipsoid.Row_effective_date = formatDateString(cs_ellipsoid.Row_effective_date)
	cs_ellipsoid.Row_expiry_date = formatDateString(cs_ellipsoid.Row_expiry_date)






	rows, err := stmt.Exec(cs_ellipsoid.Active_ind, cs_ellipsoid.Axis_ouom, cs_ellipsoid.Eccentricity, cs_ellipsoid.Effective_date, cs_ellipsoid.Ellipsoid_name, cs_ellipsoid.Expiry_date, cs_ellipsoid.Inverse_flattening, cs_ellipsoid.Ppdm_guid, cs_ellipsoid.Remark, cs_ellipsoid.Semi_major_axis, cs_ellipsoid.Semi_major_axis_accuracy, cs_ellipsoid.Semi_minor_axis, cs_ellipsoid.Semi_minor_axis_accuracy, cs_ellipsoid.Source, cs_ellipsoid.Source_document_id, cs_ellipsoid.Row_changed_by, cs_ellipsoid.Row_changed_date, cs_ellipsoid.Row_created_by, cs_ellipsoid.Row_effective_date, cs_ellipsoid.Row_expiry_date, cs_ellipsoid.Row_quality, cs_ellipsoid.Ellipsoid_id)
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

func PatchCsEllipsoid(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_ellipsoid set "
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
	query += " where ellipsoid_id = :id"

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

func DeleteCsEllipsoid(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_ellipsoid dto.Cs_ellipsoid
	cs_ellipsoid.Ellipsoid_id = id

	stmt, err := db.Prepare("delete from cs_ellipsoid where ellipsoid_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_ellipsoid.Ellipsoid_id)
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


