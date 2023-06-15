package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRInstDetailRefValue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_inst_detail_ref_value")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_inst_detail_ref_value

	for rows.Next() {
		var r_inst_detail_ref_value dto.R_inst_detail_ref_value
		if err := rows.Scan(&r_inst_detail_ref_value.Reference_value_type, &r_inst_detail_ref_value.Abbreviation, &r_inst_detail_ref_value.Active_ind, &r_inst_detail_ref_value.Effective_date, &r_inst_detail_ref_value.Expiry_date, &r_inst_detail_ref_value.Long_name, &r_inst_detail_ref_value.Ppdm_guid, &r_inst_detail_ref_value.Remark, &r_inst_detail_ref_value.Short_name, &r_inst_detail_ref_value.Source, &r_inst_detail_ref_value.Source_document, &r_inst_detail_ref_value.Row_changed_by, &r_inst_detail_ref_value.Row_changed_date, &r_inst_detail_ref_value.Row_created_by, &r_inst_detail_ref_value.Row_created_date, &r_inst_detail_ref_value.Row_effective_date, &r_inst_detail_ref_value.Row_expiry_date, &r_inst_detail_ref_value.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_inst_detail_ref_value to result
		result = append(result, r_inst_detail_ref_value)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRInstDetailRefValue(c *fiber.Ctx) error {
	var r_inst_detail_ref_value dto.R_inst_detail_ref_value

	setDefaults(&r_inst_detail_ref_value)

	if err := json.Unmarshal(c.Body(), &r_inst_detail_ref_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_inst_detail_ref_value values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	r_inst_detail_ref_value.Row_created_date = formatDate(r_inst_detail_ref_value.Row_created_date)
	r_inst_detail_ref_value.Row_changed_date = formatDate(r_inst_detail_ref_value.Row_changed_date)
	r_inst_detail_ref_value.Effective_date = formatDateString(r_inst_detail_ref_value.Effective_date)
	r_inst_detail_ref_value.Expiry_date = formatDateString(r_inst_detail_ref_value.Expiry_date)
	r_inst_detail_ref_value.Row_effective_date = formatDateString(r_inst_detail_ref_value.Row_effective_date)
	r_inst_detail_ref_value.Row_expiry_date = formatDateString(r_inst_detail_ref_value.Row_expiry_date)






	rows, err := stmt.Exec(r_inst_detail_ref_value.Reference_value_type, r_inst_detail_ref_value.Abbreviation, r_inst_detail_ref_value.Active_ind, r_inst_detail_ref_value.Effective_date, r_inst_detail_ref_value.Expiry_date, r_inst_detail_ref_value.Long_name, r_inst_detail_ref_value.Ppdm_guid, r_inst_detail_ref_value.Remark, r_inst_detail_ref_value.Short_name, r_inst_detail_ref_value.Source, r_inst_detail_ref_value.Source_document, r_inst_detail_ref_value.Row_changed_by, r_inst_detail_ref_value.Row_changed_date, r_inst_detail_ref_value.Row_created_by, r_inst_detail_ref_value.Row_created_date, r_inst_detail_ref_value.Row_effective_date, r_inst_detail_ref_value.Row_expiry_date, r_inst_detail_ref_value.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRInstDetailRefValue(c *fiber.Ctx) error {
	var r_inst_detail_ref_value dto.R_inst_detail_ref_value

	setDefaults(&r_inst_detail_ref_value)

	if err := json.Unmarshal(c.Body(), &r_inst_detail_ref_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_inst_detail_ref_value.Reference_value_type = id

    if r_inst_detail_ref_value.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_inst_detail_ref_value set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, source_document = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where reference_value_type = :18")
	if err != nil {
		return err
	}

	r_inst_detail_ref_value.Row_changed_date = formatDate(r_inst_detail_ref_value.Row_changed_date)
	r_inst_detail_ref_value.Effective_date = formatDateString(r_inst_detail_ref_value.Effective_date)
	r_inst_detail_ref_value.Expiry_date = formatDateString(r_inst_detail_ref_value.Expiry_date)
	r_inst_detail_ref_value.Row_effective_date = formatDateString(r_inst_detail_ref_value.Row_effective_date)
	r_inst_detail_ref_value.Row_expiry_date = formatDateString(r_inst_detail_ref_value.Row_expiry_date)






	rows, err := stmt.Exec(r_inst_detail_ref_value.Abbreviation, r_inst_detail_ref_value.Active_ind, r_inst_detail_ref_value.Effective_date, r_inst_detail_ref_value.Expiry_date, r_inst_detail_ref_value.Long_name, r_inst_detail_ref_value.Ppdm_guid, r_inst_detail_ref_value.Remark, r_inst_detail_ref_value.Short_name, r_inst_detail_ref_value.Source, r_inst_detail_ref_value.Source_document, r_inst_detail_ref_value.Row_changed_by, r_inst_detail_ref_value.Row_changed_date, r_inst_detail_ref_value.Row_created_by, r_inst_detail_ref_value.Row_effective_date, r_inst_detail_ref_value.Row_expiry_date, r_inst_detail_ref_value.Row_quality, r_inst_detail_ref_value.Reference_value_type)
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

func PatchRInstDetailRefValue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_inst_detail_ref_value set "
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
	query += " where reference_value_type = :id"

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

func DeleteRInstDetailRefValue(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_inst_detail_ref_value dto.R_inst_detail_ref_value
	r_inst_detail_ref_value.Reference_value_type = id

	stmt, err := db.Prepare("delete from r_inst_detail_ref_value where reference_value_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_inst_detail_ref_value.Reference_value_type)
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


