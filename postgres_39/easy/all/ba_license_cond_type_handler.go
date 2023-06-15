package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaLicenseCondType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_license_cond_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_license_cond_type

	for rows.Next() {
		var ba_license_cond_type dto.Ba_license_cond_type
		if err := rows.Scan(&ba_license_cond_type.Granted_by_ba_id, &ba_license_cond_type.License_type, &ba_license_cond_type.Condition_type, &ba_license_cond_type.Abbreviation, &ba_license_cond_type.Active_ind, &ba_license_cond_type.Effective_date, &ba_license_cond_type.Expiry_date, &ba_license_cond_type.Long_name, &ba_license_cond_type.Ppdm_guid, &ba_license_cond_type.Property_set_id, &ba_license_cond_type.Remark, &ba_license_cond_type.Short_name, &ba_license_cond_type.Source, &ba_license_cond_type.Row_changed_by, &ba_license_cond_type.Row_changed_date, &ba_license_cond_type.Row_created_by, &ba_license_cond_type.Row_created_date, &ba_license_cond_type.Row_effective_date, &ba_license_cond_type.Row_expiry_date, &ba_license_cond_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_license_cond_type to result
		result = append(result, ba_license_cond_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaLicenseCondType(c *fiber.Ctx) error {
	var ba_license_cond_type dto.Ba_license_cond_type

	setDefaults(&ba_license_cond_type)

	if err := json.Unmarshal(c.Body(), &ba_license_cond_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_license_cond_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	ba_license_cond_type.Row_created_date = formatDate(ba_license_cond_type.Row_created_date)
	ba_license_cond_type.Row_changed_date = formatDate(ba_license_cond_type.Row_changed_date)
	ba_license_cond_type.Effective_date = formatDateString(ba_license_cond_type.Effective_date)
	ba_license_cond_type.Expiry_date = formatDateString(ba_license_cond_type.Expiry_date)
	ba_license_cond_type.Row_effective_date = formatDateString(ba_license_cond_type.Row_effective_date)
	ba_license_cond_type.Row_expiry_date = formatDateString(ba_license_cond_type.Row_expiry_date)






	rows, err := stmt.Exec(ba_license_cond_type.Granted_by_ba_id, ba_license_cond_type.License_type, ba_license_cond_type.Condition_type, ba_license_cond_type.Abbreviation, ba_license_cond_type.Active_ind, ba_license_cond_type.Effective_date, ba_license_cond_type.Expiry_date, ba_license_cond_type.Long_name, ba_license_cond_type.Ppdm_guid, ba_license_cond_type.Property_set_id, ba_license_cond_type.Remark, ba_license_cond_type.Short_name, ba_license_cond_type.Source, ba_license_cond_type.Row_changed_by, ba_license_cond_type.Row_changed_date, ba_license_cond_type.Row_created_by, ba_license_cond_type.Row_created_date, ba_license_cond_type.Row_effective_date, ba_license_cond_type.Row_expiry_date, ba_license_cond_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaLicenseCondType(c *fiber.Ctx) error {
	var ba_license_cond_type dto.Ba_license_cond_type

	setDefaults(&ba_license_cond_type)

	if err := json.Unmarshal(c.Body(), &ba_license_cond_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_license_cond_type.Granted_by_ba_id = id

    if ba_license_cond_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_license_cond_type set license_type = :1, condition_type = :2, abbreviation = :3, active_ind = :4, effective_date = :5, expiry_date = :6, long_name = :7, ppdm_guid = :8, property_set_id = :9, remark = :10, short_name = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where granted_by_ba_id = :20")
	if err != nil {
		return err
	}

	ba_license_cond_type.Row_changed_date = formatDate(ba_license_cond_type.Row_changed_date)
	ba_license_cond_type.Effective_date = formatDateString(ba_license_cond_type.Effective_date)
	ba_license_cond_type.Expiry_date = formatDateString(ba_license_cond_type.Expiry_date)
	ba_license_cond_type.Row_effective_date = formatDateString(ba_license_cond_type.Row_effective_date)
	ba_license_cond_type.Row_expiry_date = formatDateString(ba_license_cond_type.Row_expiry_date)






	rows, err := stmt.Exec(ba_license_cond_type.License_type, ba_license_cond_type.Condition_type, ba_license_cond_type.Abbreviation, ba_license_cond_type.Active_ind, ba_license_cond_type.Effective_date, ba_license_cond_type.Expiry_date, ba_license_cond_type.Long_name, ba_license_cond_type.Ppdm_guid, ba_license_cond_type.Property_set_id, ba_license_cond_type.Remark, ba_license_cond_type.Short_name, ba_license_cond_type.Source, ba_license_cond_type.Row_changed_by, ba_license_cond_type.Row_changed_date, ba_license_cond_type.Row_created_by, ba_license_cond_type.Row_effective_date, ba_license_cond_type.Row_expiry_date, ba_license_cond_type.Row_quality, ba_license_cond_type.Granted_by_ba_id)
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

func PatchBaLicenseCondType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_license_cond_type set "
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
	query += " where granted_by_ba_id = :id"

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

func DeleteBaLicenseCondType(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_license_cond_type dto.Ba_license_cond_type
	ba_license_cond_type.Granted_by_ba_id = id

	stmt, err := db.Prepare("delete from ba_license_cond_type where granted_by_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_license_cond_type.Granted_by_ba_id)
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


