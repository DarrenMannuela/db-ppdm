package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRFacLicCond(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_fac_lic_cond")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_fac_lic_cond

	for rows.Next() {
		var r_fac_lic_cond dto.R_fac_lic_cond
		if err := rows.Scan(&r_fac_lic_cond.Condition_type, &r_fac_lic_cond.Abbreviation, &r_fac_lic_cond.Active_ind, &r_fac_lic_cond.Effective_date, &r_fac_lic_cond.Expiry_date, &r_fac_lic_cond.Long_name, &r_fac_lic_cond.Ppdm_guid, &r_fac_lic_cond.Property_set_id, &r_fac_lic_cond.Remark, &r_fac_lic_cond.Short_name, &r_fac_lic_cond.Source, &r_fac_lic_cond.Row_changed_by, &r_fac_lic_cond.Row_changed_date, &r_fac_lic_cond.Row_created_by, &r_fac_lic_cond.Row_created_date, &r_fac_lic_cond.Row_effective_date, &r_fac_lic_cond.Row_expiry_date, &r_fac_lic_cond.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_fac_lic_cond to result
		result = append(result, r_fac_lic_cond)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRFacLicCond(c *fiber.Ctx) error {
	var r_fac_lic_cond dto.R_fac_lic_cond

	setDefaults(&r_fac_lic_cond)

	if err := json.Unmarshal(c.Body(), &r_fac_lic_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_fac_lic_cond values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	r_fac_lic_cond.Row_created_date = formatDate(r_fac_lic_cond.Row_created_date)
	r_fac_lic_cond.Row_changed_date = formatDate(r_fac_lic_cond.Row_changed_date)
	r_fac_lic_cond.Effective_date = formatDateString(r_fac_lic_cond.Effective_date)
	r_fac_lic_cond.Expiry_date = formatDateString(r_fac_lic_cond.Expiry_date)
	r_fac_lic_cond.Row_effective_date = formatDateString(r_fac_lic_cond.Row_effective_date)
	r_fac_lic_cond.Row_expiry_date = formatDateString(r_fac_lic_cond.Row_expiry_date)






	rows, err := stmt.Exec(r_fac_lic_cond.Condition_type, r_fac_lic_cond.Abbreviation, r_fac_lic_cond.Active_ind, r_fac_lic_cond.Effective_date, r_fac_lic_cond.Expiry_date, r_fac_lic_cond.Long_name, r_fac_lic_cond.Ppdm_guid, r_fac_lic_cond.Property_set_id, r_fac_lic_cond.Remark, r_fac_lic_cond.Short_name, r_fac_lic_cond.Source, r_fac_lic_cond.Row_changed_by, r_fac_lic_cond.Row_changed_date, r_fac_lic_cond.Row_created_by, r_fac_lic_cond.Row_created_date, r_fac_lic_cond.Row_effective_date, r_fac_lic_cond.Row_expiry_date, r_fac_lic_cond.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRFacLicCond(c *fiber.Ctx) error {
	var r_fac_lic_cond dto.R_fac_lic_cond

	setDefaults(&r_fac_lic_cond)

	if err := json.Unmarshal(c.Body(), &r_fac_lic_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_fac_lic_cond.Condition_type = id

    if r_fac_lic_cond.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_fac_lic_cond set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, property_set_id = :7, remark = :8, short_name = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where condition_type = :18")
	if err != nil {
		return err
	}

	r_fac_lic_cond.Row_changed_date = formatDate(r_fac_lic_cond.Row_changed_date)
	r_fac_lic_cond.Effective_date = formatDateString(r_fac_lic_cond.Effective_date)
	r_fac_lic_cond.Expiry_date = formatDateString(r_fac_lic_cond.Expiry_date)
	r_fac_lic_cond.Row_effective_date = formatDateString(r_fac_lic_cond.Row_effective_date)
	r_fac_lic_cond.Row_expiry_date = formatDateString(r_fac_lic_cond.Row_expiry_date)






	rows, err := stmt.Exec(r_fac_lic_cond.Abbreviation, r_fac_lic_cond.Active_ind, r_fac_lic_cond.Effective_date, r_fac_lic_cond.Expiry_date, r_fac_lic_cond.Long_name, r_fac_lic_cond.Ppdm_guid, r_fac_lic_cond.Property_set_id, r_fac_lic_cond.Remark, r_fac_lic_cond.Short_name, r_fac_lic_cond.Source, r_fac_lic_cond.Row_changed_by, r_fac_lic_cond.Row_changed_date, r_fac_lic_cond.Row_created_by, r_fac_lic_cond.Row_effective_date, r_fac_lic_cond.Row_expiry_date, r_fac_lic_cond.Row_quality, r_fac_lic_cond.Condition_type)
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

func PatchRFacLicCond(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_fac_lic_cond set "
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
	query += " where condition_type = :id"

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

func DeleteRFacLicCond(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_fac_lic_cond dto.R_fac_lic_cond
	r_fac_lic_cond.Condition_type = id

	stmt, err := db.Prepare("delete from r_fac_lic_cond where condition_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_fac_lic_cond.Condition_type)
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


