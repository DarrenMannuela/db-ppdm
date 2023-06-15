package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRAnalysisProperty(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_analysis_property")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_analysis_property

	for rows.Next() {
		var r_analysis_property dto.R_analysis_property
		if err := rows.Scan(&r_analysis_property.Analysis_property, &r_analysis_property.Abbreviation, &r_analysis_property.Active_ind, &r_analysis_property.Effective_date, &r_analysis_property.Expiry_date, &r_analysis_property.Long_name, &r_analysis_property.Ppdm_guid, &r_analysis_property.Property_set_id, &r_analysis_property.Remark, &r_analysis_property.Short_name, &r_analysis_property.Source, &r_analysis_property.Row_changed_by, &r_analysis_property.Row_changed_date, &r_analysis_property.Row_created_by, &r_analysis_property.Row_created_date, &r_analysis_property.Row_effective_date, &r_analysis_property.Row_expiry_date, &r_analysis_property.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_analysis_property to result
		result = append(result, r_analysis_property)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRAnalysisProperty(c *fiber.Ctx) error {
	var r_analysis_property dto.R_analysis_property

	setDefaults(&r_analysis_property)

	if err := json.Unmarshal(c.Body(), &r_analysis_property); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_analysis_property values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	r_analysis_property.Row_created_date = formatDate(r_analysis_property.Row_created_date)
	r_analysis_property.Row_changed_date = formatDate(r_analysis_property.Row_changed_date)
	r_analysis_property.Effective_date = formatDateString(r_analysis_property.Effective_date)
	r_analysis_property.Expiry_date = formatDateString(r_analysis_property.Expiry_date)
	r_analysis_property.Row_effective_date = formatDateString(r_analysis_property.Row_effective_date)
	r_analysis_property.Row_expiry_date = formatDateString(r_analysis_property.Row_expiry_date)






	rows, err := stmt.Exec(r_analysis_property.Analysis_property, r_analysis_property.Abbreviation, r_analysis_property.Active_ind, r_analysis_property.Effective_date, r_analysis_property.Expiry_date, r_analysis_property.Long_name, r_analysis_property.Ppdm_guid, r_analysis_property.Property_set_id, r_analysis_property.Remark, r_analysis_property.Short_name, r_analysis_property.Source, r_analysis_property.Row_changed_by, r_analysis_property.Row_changed_date, r_analysis_property.Row_created_by, r_analysis_property.Row_created_date, r_analysis_property.Row_effective_date, r_analysis_property.Row_expiry_date, r_analysis_property.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRAnalysisProperty(c *fiber.Ctx) error {
	var r_analysis_property dto.R_analysis_property

	setDefaults(&r_analysis_property)

	if err := json.Unmarshal(c.Body(), &r_analysis_property); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_analysis_property.Analysis_property = id

    if r_analysis_property.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_analysis_property set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, property_set_id = :7, remark = :8, short_name = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where analysis_property = :18")
	if err != nil {
		return err
	}

	r_analysis_property.Row_changed_date = formatDate(r_analysis_property.Row_changed_date)
	r_analysis_property.Effective_date = formatDateString(r_analysis_property.Effective_date)
	r_analysis_property.Expiry_date = formatDateString(r_analysis_property.Expiry_date)
	r_analysis_property.Row_effective_date = formatDateString(r_analysis_property.Row_effective_date)
	r_analysis_property.Row_expiry_date = formatDateString(r_analysis_property.Row_expiry_date)






	rows, err := stmt.Exec(r_analysis_property.Abbreviation, r_analysis_property.Active_ind, r_analysis_property.Effective_date, r_analysis_property.Expiry_date, r_analysis_property.Long_name, r_analysis_property.Ppdm_guid, r_analysis_property.Property_set_id, r_analysis_property.Remark, r_analysis_property.Short_name, r_analysis_property.Source, r_analysis_property.Row_changed_by, r_analysis_property.Row_changed_date, r_analysis_property.Row_created_by, r_analysis_property.Row_effective_date, r_analysis_property.Row_expiry_date, r_analysis_property.Row_quality, r_analysis_property.Analysis_property)
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

func PatchRAnalysisProperty(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_analysis_property set "
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
	query += " where analysis_property = :id"

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

func DeleteRAnalysisProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_analysis_property dto.R_analysis_property
	r_analysis_property.Analysis_property = id

	stmt, err := db.Prepare("delete from r_analysis_property where analysis_property = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_analysis_property.Analysis_property)
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


