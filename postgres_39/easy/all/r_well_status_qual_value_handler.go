package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRWellStatusQualValue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_well_status_qual_value")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_well_status_qual_value

	for rows.Next() {
		var r_well_status_qual_value dto.R_well_status_qual_value
		if err := rows.Scan(&r_well_status_qual_value.Status_type, &r_well_status_qual_value.Status, &r_well_status_qual_value.Status_qualifier, &r_well_status_qual_value.Status_qualifier_value, &r_well_status_qual_value.Abbreviation, &r_well_status_qual_value.Active_ind, &r_well_status_qual_value.Effective_date, &r_well_status_qual_value.Expiry_date, &r_well_status_qual_value.Long_name, &r_well_status_qual_value.Ppdm_guid, &r_well_status_qual_value.Remark, &r_well_status_qual_value.Short_name, &r_well_status_qual_value.Source, &r_well_status_qual_value.Status_group, &r_well_status_qual_value.Row_changed_by, &r_well_status_qual_value.Row_changed_date, &r_well_status_qual_value.Row_created_by, &r_well_status_qual_value.Row_created_date, &r_well_status_qual_value.Row_effective_date, &r_well_status_qual_value.Row_expiry_date, &r_well_status_qual_value.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_well_status_qual_value to result
		result = append(result, r_well_status_qual_value)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRWellStatusQualValue(c *fiber.Ctx) error {
	var r_well_status_qual_value dto.R_well_status_qual_value

	setDefaults(&r_well_status_qual_value)

	if err := json.Unmarshal(c.Body(), &r_well_status_qual_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_well_status_qual_value values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	r_well_status_qual_value.Row_created_date = formatDate(r_well_status_qual_value.Row_created_date)
	r_well_status_qual_value.Row_changed_date = formatDate(r_well_status_qual_value.Row_changed_date)
	r_well_status_qual_value.Effective_date = formatDateString(r_well_status_qual_value.Effective_date)
	r_well_status_qual_value.Expiry_date = formatDateString(r_well_status_qual_value.Expiry_date)
	r_well_status_qual_value.Row_effective_date = formatDateString(r_well_status_qual_value.Row_effective_date)
	r_well_status_qual_value.Row_expiry_date = formatDateString(r_well_status_qual_value.Row_expiry_date)






	rows, err := stmt.Exec(r_well_status_qual_value.Status_type, r_well_status_qual_value.Status, r_well_status_qual_value.Status_qualifier, r_well_status_qual_value.Status_qualifier_value, r_well_status_qual_value.Abbreviation, r_well_status_qual_value.Active_ind, r_well_status_qual_value.Effective_date, r_well_status_qual_value.Expiry_date, r_well_status_qual_value.Long_name, r_well_status_qual_value.Ppdm_guid, r_well_status_qual_value.Remark, r_well_status_qual_value.Short_name, r_well_status_qual_value.Source, r_well_status_qual_value.Status_group, r_well_status_qual_value.Row_changed_by, r_well_status_qual_value.Row_changed_date, r_well_status_qual_value.Row_created_by, r_well_status_qual_value.Row_created_date, r_well_status_qual_value.Row_effective_date, r_well_status_qual_value.Row_expiry_date, r_well_status_qual_value.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRWellStatusQualValue(c *fiber.Ctx) error {
	var r_well_status_qual_value dto.R_well_status_qual_value

	setDefaults(&r_well_status_qual_value)

	if err := json.Unmarshal(c.Body(), &r_well_status_qual_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_well_status_qual_value.Status_type = id

    if r_well_status_qual_value.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_well_status_qual_value set status = :1, status_qualifier = :2, status_qualifier_value = :3, abbreviation = :4, active_ind = :5, effective_date = :6, expiry_date = :7, long_name = :8, ppdm_guid = :9, remark = :10, short_name = :11, source = :12, status_group = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where status_type = :21")
	if err != nil {
		return err
	}

	r_well_status_qual_value.Row_changed_date = formatDate(r_well_status_qual_value.Row_changed_date)
	r_well_status_qual_value.Effective_date = formatDateString(r_well_status_qual_value.Effective_date)
	r_well_status_qual_value.Expiry_date = formatDateString(r_well_status_qual_value.Expiry_date)
	r_well_status_qual_value.Row_effective_date = formatDateString(r_well_status_qual_value.Row_effective_date)
	r_well_status_qual_value.Row_expiry_date = formatDateString(r_well_status_qual_value.Row_expiry_date)






	rows, err := stmt.Exec(r_well_status_qual_value.Status, r_well_status_qual_value.Status_qualifier, r_well_status_qual_value.Status_qualifier_value, r_well_status_qual_value.Abbreviation, r_well_status_qual_value.Active_ind, r_well_status_qual_value.Effective_date, r_well_status_qual_value.Expiry_date, r_well_status_qual_value.Long_name, r_well_status_qual_value.Ppdm_guid, r_well_status_qual_value.Remark, r_well_status_qual_value.Short_name, r_well_status_qual_value.Source, r_well_status_qual_value.Status_group, r_well_status_qual_value.Row_changed_by, r_well_status_qual_value.Row_changed_date, r_well_status_qual_value.Row_created_by, r_well_status_qual_value.Row_effective_date, r_well_status_qual_value.Row_expiry_date, r_well_status_qual_value.Row_quality, r_well_status_qual_value.Status_type)
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

func PatchRWellStatusQualValue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_well_status_qual_value set "
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
	query += " where status_type = :id"

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

func DeleteRWellStatusQualValue(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_well_status_qual_value dto.R_well_status_qual_value
	r_well_status_qual_value.Status_type = id

	stmt, err := db.Prepare("delete from r_well_status_qual_value where status_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_well_status_qual_value.Status_type)
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


