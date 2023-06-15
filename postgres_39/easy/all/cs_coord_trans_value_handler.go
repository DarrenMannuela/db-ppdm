package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsCoordTransValue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_coord_trans_value")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_coord_trans_value

	for rows.Next() {
		var cs_coord_trans_value dto.Cs_coord_trans_value
		if err := rows.Scan(&cs_coord_trans_value.Transform_id, &cs_coord_trans_value.Transform_type, &cs_coord_trans_value.Parameter_id, &cs_coord_trans_value.Value_id, &cs_coord_trans_value.Active_ind, &cs_coord_trans_value.Effective_date, &cs_coord_trans_value.Expiry_date, &cs_coord_trans_value.Ppdm_guid, &cs_coord_trans_value.Remark, &cs_coord_trans_value.Source, &cs_coord_trans_value.Source_document_id, &cs_coord_trans_value.Text_value, &cs_coord_trans_value.Trans_value, &cs_coord_trans_value.Trans_value_uom, &cs_coord_trans_value.Row_changed_by, &cs_coord_trans_value.Row_changed_date, &cs_coord_trans_value.Row_created_by, &cs_coord_trans_value.Row_created_date, &cs_coord_trans_value.Row_effective_date, &cs_coord_trans_value.Row_expiry_date, &cs_coord_trans_value.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_coord_trans_value to result
		result = append(result, cs_coord_trans_value)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsCoordTransValue(c *fiber.Ctx) error {
	var cs_coord_trans_value dto.Cs_coord_trans_value

	setDefaults(&cs_coord_trans_value)

	if err := json.Unmarshal(c.Body(), &cs_coord_trans_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_coord_trans_value values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	cs_coord_trans_value.Row_created_date = formatDate(cs_coord_trans_value.Row_created_date)
	cs_coord_trans_value.Row_changed_date = formatDate(cs_coord_trans_value.Row_changed_date)
	cs_coord_trans_value.Effective_date = formatDateString(cs_coord_trans_value.Effective_date)
	cs_coord_trans_value.Expiry_date = formatDateString(cs_coord_trans_value.Expiry_date)
	cs_coord_trans_value.Row_effective_date = formatDateString(cs_coord_trans_value.Row_effective_date)
	cs_coord_trans_value.Row_expiry_date = formatDateString(cs_coord_trans_value.Row_expiry_date)






	rows, err := stmt.Exec(cs_coord_trans_value.Transform_id, cs_coord_trans_value.Transform_type, cs_coord_trans_value.Parameter_id, cs_coord_trans_value.Value_id, cs_coord_trans_value.Active_ind, cs_coord_trans_value.Effective_date, cs_coord_trans_value.Expiry_date, cs_coord_trans_value.Ppdm_guid, cs_coord_trans_value.Remark, cs_coord_trans_value.Source, cs_coord_trans_value.Source_document_id, cs_coord_trans_value.Text_value, cs_coord_trans_value.Trans_value, cs_coord_trans_value.Trans_value_uom, cs_coord_trans_value.Row_changed_by, cs_coord_trans_value.Row_changed_date, cs_coord_trans_value.Row_created_by, cs_coord_trans_value.Row_created_date, cs_coord_trans_value.Row_effective_date, cs_coord_trans_value.Row_expiry_date, cs_coord_trans_value.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsCoordTransValue(c *fiber.Ctx) error {
	var cs_coord_trans_value dto.Cs_coord_trans_value

	setDefaults(&cs_coord_trans_value)

	if err := json.Unmarshal(c.Body(), &cs_coord_trans_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_coord_trans_value.Transform_id = id

    if cs_coord_trans_value.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_coord_trans_value set transform_type = :1, parameter_id = :2, value_id = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, source_document_id = :10, text_value = :11, trans_value = :12, trans_value_uom = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where transform_id = :21")
	if err != nil {
		return err
	}

	cs_coord_trans_value.Row_changed_date = formatDate(cs_coord_trans_value.Row_changed_date)
	cs_coord_trans_value.Effective_date = formatDateString(cs_coord_trans_value.Effective_date)
	cs_coord_trans_value.Expiry_date = formatDateString(cs_coord_trans_value.Expiry_date)
	cs_coord_trans_value.Row_effective_date = formatDateString(cs_coord_trans_value.Row_effective_date)
	cs_coord_trans_value.Row_expiry_date = formatDateString(cs_coord_trans_value.Row_expiry_date)






	rows, err := stmt.Exec(cs_coord_trans_value.Transform_type, cs_coord_trans_value.Parameter_id, cs_coord_trans_value.Value_id, cs_coord_trans_value.Active_ind, cs_coord_trans_value.Effective_date, cs_coord_trans_value.Expiry_date, cs_coord_trans_value.Ppdm_guid, cs_coord_trans_value.Remark, cs_coord_trans_value.Source, cs_coord_trans_value.Source_document_id, cs_coord_trans_value.Text_value, cs_coord_trans_value.Trans_value, cs_coord_trans_value.Trans_value_uom, cs_coord_trans_value.Row_changed_by, cs_coord_trans_value.Row_changed_date, cs_coord_trans_value.Row_created_by, cs_coord_trans_value.Row_effective_date, cs_coord_trans_value.Row_expiry_date, cs_coord_trans_value.Row_quality, cs_coord_trans_value.Transform_id)
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

func PatchCsCoordTransValue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_coord_trans_value set "
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

func DeleteCsCoordTransValue(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_coord_trans_value dto.Cs_coord_trans_value
	cs_coord_trans_value.Transform_id = id

	stmt, err := db.Prepare("delete from cs_coord_trans_value where transform_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_coord_trans_value.Transform_id)
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


