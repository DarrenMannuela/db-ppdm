package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsCoordTransform(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_coord_transform")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_coord_transform

	for rows.Next() {
		var cs_coord_transform dto.Cs_coord_transform
		if err := rows.Scan(&cs_coord_transform.Transform_id, &cs_coord_transform.Active_ind, &cs_coord_transform.Effective_date, &cs_coord_transform.Expiry_date, &cs_coord_transform.Ppdm_guid, &cs_coord_transform.Preferred_ind, &cs_coord_transform.Remark, &cs_coord_transform.Source, &cs_coord_transform.Source_coord_system_id, &cs_coord_transform.Source_document_id, &cs_coord_transform.Target_coord_system_id, &cs_coord_transform.Transform_name, &cs_coord_transform.Transform_type, &cs_coord_transform.Row_changed_by, &cs_coord_transform.Row_changed_date, &cs_coord_transform.Row_created_by, &cs_coord_transform.Row_created_date, &cs_coord_transform.Row_effective_date, &cs_coord_transform.Row_expiry_date, &cs_coord_transform.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_coord_transform to result
		result = append(result, cs_coord_transform)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsCoordTransform(c *fiber.Ctx) error {
	var cs_coord_transform dto.Cs_coord_transform

	setDefaults(&cs_coord_transform)

	if err := json.Unmarshal(c.Body(), &cs_coord_transform); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_coord_transform values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	cs_coord_transform.Row_created_date = formatDate(cs_coord_transform.Row_created_date)
	cs_coord_transform.Row_changed_date = formatDate(cs_coord_transform.Row_changed_date)
	cs_coord_transform.Effective_date = formatDateString(cs_coord_transform.Effective_date)
	cs_coord_transform.Expiry_date = formatDateString(cs_coord_transform.Expiry_date)
	cs_coord_transform.Row_effective_date = formatDateString(cs_coord_transform.Row_effective_date)
	cs_coord_transform.Row_expiry_date = formatDateString(cs_coord_transform.Row_expiry_date)






	rows, err := stmt.Exec(cs_coord_transform.Transform_id, cs_coord_transform.Active_ind, cs_coord_transform.Effective_date, cs_coord_transform.Expiry_date, cs_coord_transform.Ppdm_guid, cs_coord_transform.Preferred_ind, cs_coord_transform.Remark, cs_coord_transform.Source, cs_coord_transform.Source_coord_system_id, cs_coord_transform.Source_document_id, cs_coord_transform.Target_coord_system_id, cs_coord_transform.Transform_name, cs_coord_transform.Transform_type, cs_coord_transform.Row_changed_by, cs_coord_transform.Row_changed_date, cs_coord_transform.Row_created_by, cs_coord_transform.Row_created_date, cs_coord_transform.Row_effective_date, cs_coord_transform.Row_expiry_date, cs_coord_transform.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsCoordTransform(c *fiber.Ctx) error {
	var cs_coord_transform dto.Cs_coord_transform

	setDefaults(&cs_coord_transform)

	if err := json.Unmarshal(c.Body(), &cs_coord_transform); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_coord_transform.Transform_id = id

    if cs_coord_transform.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_coord_transform set active_ind = :1, effective_date = :2, expiry_date = :3, ppdm_guid = :4, preferred_ind = :5, remark = :6, source = :7, source_coord_system_id = :8, source_document_id = :9, target_coord_system_id = :10, transform_name = :11, transform_type = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where transform_id = :20")
	if err != nil {
		return err
	}

	cs_coord_transform.Row_changed_date = formatDate(cs_coord_transform.Row_changed_date)
	cs_coord_transform.Effective_date = formatDateString(cs_coord_transform.Effective_date)
	cs_coord_transform.Expiry_date = formatDateString(cs_coord_transform.Expiry_date)
	cs_coord_transform.Row_effective_date = formatDateString(cs_coord_transform.Row_effective_date)
	cs_coord_transform.Row_expiry_date = formatDateString(cs_coord_transform.Row_expiry_date)






	rows, err := stmt.Exec(cs_coord_transform.Active_ind, cs_coord_transform.Effective_date, cs_coord_transform.Expiry_date, cs_coord_transform.Ppdm_guid, cs_coord_transform.Preferred_ind, cs_coord_transform.Remark, cs_coord_transform.Source, cs_coord_transform.Source_coord_system_id, cs_coord_transform.Source_document_id, cs_coord_transform.Target_coord_system_id, cs_coord_transform.Transform_name, cs_coord_transform.Transform_type, cs_coord_transform.Row_changed_by, cs_coord_transform.Row_changed_date, cs_coord_transform.Row_created_by, cs_coord_transform.Row_effective_date, cs_coord_transform.Row_expiry_date, cs_coord_transform.Row_quality, cs_coord_transform.Transform_id)
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

func PatchCsCoordTransform(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_coord_transform set "
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
	query += " where transform_id = :id"

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

func DeleteCsCoordTransform(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_coord_transform dto.Cs_coord_transform
	cs_coord_transform.Transform_id = id

	stmt, err := db.Prepare("delete from cs_coord_transform where transform_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_coord_transform.Transform_id)
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


