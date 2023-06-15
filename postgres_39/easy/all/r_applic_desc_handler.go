package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRApplicDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_applic_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_applic_desc

	for rows.Next() {
		var r_applic_desc dto.R_applic_desc
		if err := rows.Scan(&r_applic_desc.Application_type, &r_applic_desc.Applic_desc_type, &r_applic_desc.Abbreviation, &r_applic_desc.Active_ind, &r_applic_desc.Effective_date, &r_applic_desc.Expiry_date, &r_applic_desc.Long_name, &r_applic_desc.Ppdm_guid, &r_applic_desc.Property_set_id, &r_applic_desc.Remark, &r_applic_desc.Short_name, &r_applic_desc.Source, &r_applic_desc.Row_changed_by, &r_applic_desc.Row_changed_date, &r_applic_desc.Row_created_by, &r_applic_desc.Row_created_date, &r_applic_desc.Row_effective_date, &r_applic_desc.Row_expiry_date, &r_applic_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_applic_desc to result
		result = append(result, r_applic_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRApplicDesc(c *fiber.Ctx) error {
	var r_applic_desc dto.R_applic_desc

	setDefaults(&r_applic_desc)

	if err := json.Unmarshal(c.Body(), &r_applic_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_applic_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	r_applic_desc.Row_created_date = formatDate(r_applic_desc.Row_created_date)
	r_applic_desc.Row_changed_date = formatDate(r_applic_desc.Row_changed_date)
	r_applic_desc.Effective_date = formatDateString(r_applic_desc.Effective_date)
	r_applic_desc.Expiry_date = formatDateString(r_applic_desc.Expiry_date)
	r_applic_desc.Row_effective_date = formatDateString(r_applic_desc.Row_effective_date)
	r_applic_desc.Row_expiry_date = formatDateString(r_applic_desc.Row_expiry_date)






	rows, err := stmt.Exec(r_applic_desc.Application_type, r_applic_desc.Applic_desc_type, r_applic_desc.Abbreviation, r_applic_desc.Active_ind, r_applic_desc.Effective_date, r_applic_desc.Expiry_date, r_applic_desc.Long_name, r_applic_desc.Ppdm_guid, r_applic_desc.Property_set_id, r_applic_desc.Remark, r_applic_desc.Short_name, r_applic_desc.Source, r_applic_desc.Row_changed_by, r_applic_desc.Row_changed_date, r_applic_desc.Row_created_by, r_applic_desc.Row_created_date, r_applic_desc.Row_effective_date, r_applic_desc.Row_expiry_date, r_applic_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRApplicDesc(c *fiber.Ctx) error {
	var r_applic_desc dto.R_applic_desc

	setDefaults(&r_applic_desc)

	if err := json.Unmarshal(c.Body(), &r_applic_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_applic_desc.Application_type = id

    if r_applic_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_applic_desc set applic_desc_type = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, property_set_id = :8, remark = :9, short_name = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where application_type = :19")
	if err != nil {
		return err
	}

	r_applic_desc.Row_changed_date = formatDate(r_applic_desc.Row_changed_date)
	r_applic_desc.Effective_date = formatDateString(r_applic_desc.Effective_date)
	r_applic_desc.Expiry_date = formatDateString(r_applic_desc.Expiry_date)
	r_applic_desc.Row_effective_date = formatDateString(r_applic_desc.Row_effective_date)
	r_applic_desc.Row_expiry_date = formatDateString(r_applic_desc.Row_expiry_date)






	rows, err := stmt.Exec(r_applic_desc.Applic_desc_type, r_applic_desc.Abbreviation, r_applic_desc.Active_ind, r_applic_desc.Effective_date, r_applic_desc.Expiry_date, r_applic_desc.Long_name, r_applic_desc.Ppdm_guid, r_applic_desc.Property_set_id, r_applic_desc.Remark, r_applic_desc.Short_name, r_applic_desc.Source, r_applic_desc.Row_changed_by, r_applic_desc.Row_changed_date, r_applic_desc.Row_created_by, r_applic_desc.Row_effective_date, r_applic_desc.Row_expiry_date, r_applic_desc.Row_quality, r_applic_desc.Application_type)
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

func PatchRApplicDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_applic_desc set "
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
	query += " where application_type = :id"

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

func DeleteRApplicDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_applic_desc dto.R_applic_desc
	r_applic_desc.Application_type = id

	stmt, err := db.Prepare("delete from r_applic_desc where application_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_applic_desc.Application_type)
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


