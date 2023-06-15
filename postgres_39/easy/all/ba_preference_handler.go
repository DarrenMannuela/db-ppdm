package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaPreference(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_preference")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_preference

	for rows.Next() {
		var ba_preference dto.Ba_preference
		if err := rows.Scan(&ba_preference.Business_associate_id, &ba_preference.Preference_type, &ba_preference.Preference_obs_no, &ba_preference.Active_ind, &ba_preference.Description, &ba_preference.Effective_date, &ba_preference.Expiry_date, &ba_preference.Language, &ba_preference.Ppdm_guid, &ba_preference.Remark, &ba_preference.Source, &ba_preference.Sw_application_id, &ba_preference.Wl_curve_class_id, &ba_preference.Wl_parameter_class_id, &ba_preference.Row_changed_by, &ba_preference.Row_changed_date, &ba_preference.Row_created_by, &ba_preference.Row_created_date, &ba_preference.Row_effective_date, &ba_preference.Row_expiry_date, &ba_preference.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_preference to result
		result = append(result, ba_preference)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaPreference(c *fiber.Ctx) error {
	var ba_preference dto.Ba_preference

	setDefaults(&ba_preference)

	if err := json.Unmarshal(c.Body(), &ba_preference); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_preference values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	ba_preference.Row_created_date = formatDate(ba_preference.Row_created_date)
	ba_preference.Row_changed_date = formatDate(ba_preference.Row_changed_date)
	ba_preference.Effective_date = formatDateString(ba_preference.Effective_date)
	ba_preference.Expiry_date = formatDateString(ba_preference.Expiry_date)
	ba_preference.Row_effective_date = formatDateString(ba_preference.Row_effective_date)
	ba_preference.Row_expiry_date = formatDateString(ba_preference.Row_expiry_date)






	rows, err := stmt.Exec(ba_preference.Business_associate_id, ba_preference.Preference_type, ba_preference.Preference_obs_no, ba_preference.Active_ind, ba_preference.Description, ba_preference.Effective_date, ba_preference.Expiry_date, ba_preference.Language, ba_preference.Ppdm_guid, ba_preference.Remark, ba_preference.Source, ba_preference.Sw_application_id, ba_preference.Wl_curve_class_id, ba_preference.Wl_parameter_class_id, ba_preference.Row_changed_by, ba_preference.Row_changed_date, ba_preference.Row_created_by, ba_preference.Row_created_date, ba_preference.Row_effective_date, ba_preference.Row_expiry_date, ba_preference.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaPreference(c *fiber.Ctx) error {
	var ba_preference dto.Ba_preference

	setDefaults(&ba_preference)

	if err := json.Unmarshal(c.Body(), &ba_preference); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_preference.Business_associate_id = id

    if ba_preference.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_preference set preference_type = :1, preference_obs_no = :2, active_ind = :3, description = :4, effective_date = :5, expiry_date = :6, language = :7, ppdm_guid = :8, remark = :9, source = :10, sw_application_id = :11, wl_curve_class_id = :12, wl_parameter_class_id = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where business_associate_id = :21")
	if err != nil {
		return err
	}

	ba_preference.Row_changed_date = formatDate(ba_preference.Row_changed_date)
	ba_preference.Effective_date = formatDateString(ba_preference.Effective_date)
	ba_preference.Expiry_date = formatDateString(ba_preference.Expiry_date)
	ba_preference.Row_effective_date = formatDateString(ba_preference.Row_effective_date)
	ba_preference.Row_expiry_date = formatDateString(ba_preference.Row_expiry_date)






	rows, err := stmt.Exec(ba_preference.Preference_type, ba_preference.Preference_obs_no, ba_preference.Active_ind, ba_preference.Description, ba_preference.Effective_date, ba_preference.Expiry_date, ba_preference.Language, ba_preference.Ppdm_guid, ba_preference.Remark, ba_preference.Source, ba_preference.Sw_application_id, ba_preference.Wl_curve_class_id, ba_preference.Wl_parameter_class_id, ba_preference.Row_changed_by, ba_preference.Row_changed_date, ba_preference.Row_created_by, ba_preference.Row_effective_date, ba_preference.Row_expiry_date, ba_preference.Row_quality, ba_preference.Business_associate_id)
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

func PatchBaPreference(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_preference set "
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
	query += " where business_associate_id = :id"

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

func DeleteBaPreference(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_preference dto.Ba_preference
	ba_preference.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_preference where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_preference.Business_associate_id)
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


