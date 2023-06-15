package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaPreferenceLevel(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_preference_level")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_preference_level

	for rows.Next() {
		var ba_preference_level dto.Ba_preference_level
		if err := rows.Scan(&ba_preference_level.Business_associate_id, &ba_preference_level.Preference_type, &ba_preference_level.Preference_obs_no, &ba_preference_level.Level_id, &ba_preference_level.Active_ind, &ba_preference_level.Currency_uom, &ba_preference_level.Description, &ba_preference_level.Effective_date, &ba_preference_level.Expiry_date, &ba_preference_level.Ppdm_guid, &ba_preference_level.Preference_level, &ba_preference_level.Preference_rule, &ba_preference_level.Preferred_ba_id, &ba_preference_level.Remark, &ba_preference_level.Source, &ba_preference_level.Wl_dictionary_id, &ba_preference_level.Wl_dict_curve_id, &ba_preference_level.Wl_parameter_id, &ba_preference_level.Row_changed_by, &ba_preference_level.Row_changed_date, &ba_preference_level.Row_created_by, &ba_preference_level.Row_created_date, &ba_preference_level.Row_effective_date, &ba_preference_level.Row_expiry_date, &ba_preference_level.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_preference_level to result
		result = append(result, ba_preference_level)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaPreferenceLevel(c *fiber.Ctx) error {
	var ba_preference_level dto.Ba_preference_level

	setDefaults(&ba_preference_level)

	if err := json.Unmarshal(c.Body(), &ba_preference_level); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_preference_level values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	ba_preference_level.Row_created_date = formatDate(ba_preference_level.Row_created_date)
	ba_preference_level.Row_changed_date = formatDate(ba_preference_level.Row_changed_date)
	ba_preference_level.Effective_date = formatDateString(ba_preference_level.Effective_date)
	ba_preference_level.Expiry_date = formatDateString(ba_preference_level.Expiry_date)
	ba_preference_level.Row_effective_date = formatDateString(ba_preference_level.Row_effective_date)
	ba_preference_level.Row_expiry_date = formatDateString(ba_preference_level.Row_expiry_date)






	rows, err := stmt.Exec(ba_preference_level.Business_associate_id, ba_preference_level.Preference_type, ba_preference_level.Preference_obs_no, ba_preference_level.Level_id, ba_preference_level.Active_ind, ba_preference_level.Currency_uom, ba_preference_level.Description, ba_preference_level.Effective_date, ba_preference_level.Expiry_date, ba_preference_level.Ppdm_guid, ba_preference_level.Preference_level, ba_preference_level.Preference_rule, ba_preference_level.Preferred_ba_id, ba_preference_level.Remark, ba_preference_level.Source, ba_preference_level.Wl_dictionary_id, ba_preference_level.Wl_dict_curve_id, ba_preference_level.Wl_parameter_id, ba_preference_level.Row_changed_by, ba_preference_level.Row_changed_date, ba_preference_level.Row_created_by, ba_preference_level.Row_created_date, ba_preference_level.Row_effective_date, ba_preference_level.Row_expiry_date, ba_preference_level.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaPreferenceLevel(c *fiber.Ctx) error {
	var ba_preference_level dto.Ba_preference_level

	setDefaults(&ba_preference_level)

	if err := json.Unmarshal(c.Body(), &ba_preference_level); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_preference_level.Business_associate_id = id

    if ba_preference_level.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_preference_level set preference_type = :1, preference_obs_no = :2, level_id = :3, active_ind = :4, currency_uom = :5, description = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, preference_level = :10, preference_rule = :11, preferred_ba_id = :12, remark = :13, source = :14, wl_dictionary_id = :15, wl_dict_curve_id = :16, wl_parameter_id = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where business_associate_id = :25")
	if err != nil {
		return err
	}

	ba_preference_level.Row_changed_date = formatDate(ba_preference_level.Row_changed_date)
	ba_preference_level.Effective_date = formatDateString(ba_preference_level.Effective_date)
	ba_preference_level.Expiry_date = formatDateString(ba_preference_level.Expiry_date)
	ba_preference_level.Row_effective_date = formatDateString(ba_preference_level.Row_effective_date)
	ba_preference_level.Row_expiry_date = formatDateString(ba_preference_level.Row_expiry_date)






	rows, err := stmt.Exec(ba_preference_level.Preference_type, ba_preference_level.Preference_obs_no, ba_preference_level.Level_id, ba_preference_level.Active_ind, ba_preference_level.Currency_uom, ba_preference_level.Description, ba_preference_level.Effective_date, ba_preference_level.Expiry_date, ba_preference_level.Ppdm_guid, ba_preference_level.Preference_level, ba_preference_level.Preference_rule, ba_preference_level.Preferred_ba_id, ba_preference_level.Remark, ba_preference_level.Source, ba_preference_level.Wl_dictionary_id, ba_preference_level.Wl_dict_curve_id, ba_preference_level.Wl_parameter_id, ba_preference_level.Row_changed_by, ba_preference_level.Row_changed_date, ba_preference_level.Row_created_by, ba_preference_level.Row_effective_date, ba_preference_level.Row_expiry_date, ba_preference_level.Row_quality, ba_preference_level.Business_associate_id)
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

func PatchBaPreferenceLevel(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_preference_level set "
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

func DeleteBaPreferenceLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_preference_level dto.Ba_preference_level
	ba_preference_level.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_preference_level where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_preference_level.Business_associate_id)
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


