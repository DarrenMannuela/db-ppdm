package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogDictionary(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_dictionary")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_dictionary

	for rows.Next() {
		var well_log_dictionary dto.Well_log_dictionary
		if err := rows.Scan(&well_log_dictionary.Dictionary_id, &well_log_dictionary.Active_ind, &well_log_dictionary.Description, &well_log_dictionary.Effective_date, &well_log_dictionary.Expiry_date, &well_log_dictionary.Owned_by_ba_id, &well_log_dictionary.Ppdm_guid, &well_log_dictionary.Preferred_name, &well_log_dictionary.Remark, &well_log_dictionary.Source, &well_log_dictionary.Version_obs_no, &well_log_dictionary.Row_changed_by, &well_log_dictionary.Row_changed_date, &well_log_dictionary.Row_created_by, &well_log_dictionary.Row_created_date, &well_log_dictionary.Row_effective_date, &well_log_dictionary.Row_expiry_date, &well_log_dictionary.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_dictionary to result
		result = append(result, well_log_dictionary)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogDictionary(c *fiber.Ctx) error {
	var well_log_dictionary dto.Well_log_dictionary

	setDefaults(&well_log_dictionary)

	if err := json.Unmarshal(c.Body(), &well_log_dictionary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_dictionary values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	well_log_dictionary.Row_created_date = formatDate(well_log_dictionary.Row_created_date)
	well_log_dictionary.Row_changed_date = formatDate(well_log_dictionary.Row_changed_date)
	well_log_dictionary.Effective_date = formatDateString(well_log_dictionary.Effective_date)
	well_log_dictionary.Expiry_date = formatDateString(well_log_dictionary.Expiry_date)
	well_log_dictionary.Row_effective_date = formatDateString(well_log_dictionary.Row_effective_date)
	well_log_dictionary.Row_expiry_date = formatDateString(well_log_dictionary.Row_expiry_date)






	rows, err := stmt.Exec(well_log_dictionary.Dictionary_id, well_log_dictionary.Active_ind, well_log_dictionary.Description, well_log_dictionary.Effective_date, well_log_dictionary.Expiry_date, well_log_dictionary.Owned_by_ba_id, well_log_dictionary.Ppdm_guid, well_log_dictionary.Preferred_name, well_log_dictionary.Remark, well_log_dictionary.Source, well_log_dictionary.Version_obs_no, well_log_dictionary.Row_changed_by, well_log_dictionary.Row_changed_date, well_log_dictionary.Row_created_by, well_log_dictionary.Row_created_date, well_log_dictionary.Row_effective_date, well_log_dictionary.Row_expiry_date, well_log_dictionary.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogDictionary(c *fiber.Ctx) error {
	var well_log_dictionary dto.Well_log_dictionary

	setDefaults(&well_log_dictionary)

	if err := json.Unmarshal(c.Body(), &well_log_dictionary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_dictionary.Dictionary_id = id

    if well_log_dictionary.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_dictionary set active_ind = :1, description = :2, effective_date = :3, expiry_date = :4, owned_by_ba_id = :5, ppdm_guid = :6, preferred_name = :7, remark = :8, source = :9, version_obs_no = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where dictionary_id = :18")
	if err != nil {
		return err
	}

	well_log_dictionary.Row_changed_date = formatDate(well_log_dictionary.Row_changed_date)
	well_log_dictionary.Effective_date = formatDateString(well_log_dictionary.Effective_date)
	well_log_dictionary.Expiry_date = formatDateString(well_log_dictionary.Expiry_date)
	well_log_dictionary.Row_effective_date = formatDateString(well_log_dictionary.Row_effective_date)
	well_log_dictionary.Row_expiry_date = formatDateString(well_log_dictionary.Row_expiry_date)






	rows, err := stmt.Exec(well_log_dictionary.Active_ind, well_log_dictionary.Description, well_log_dictionary.Effective_date, well_log_dictionary.Expiry_date, well_log_dictionary.Owned_by_ba_id, well_log_dictionary.Ppdm_guid, well_log_dictionary.Preferred_name, well_log_dictionary.Remark, well_log_dictionary.Source, well_log_dictionary.Version_obs_no, well_log_dictionary.Row_changed_by, well_log_dictionary.Row_changed_date, well_log_dictionary.Row_created_by, well_log_dictionary.Row_effective_date, well_log_dictionary.Row_expiry_date, well_log_dictionary.Row_quality, well_log_dictionary.Dictionary_id)
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

func PatchWellLogDictionary(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_dictionary set "
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
	query += " where dictionary_id = :id"

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

func DeleteWellLogDictionary(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_dictionary dto.Well_log_dictionary
	well_log_dictionary.Dictionary_id = id

	stmt, err := db.Prepare("delete from well_log_dictionary where dictionary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_dictionary.Dictionary_id)
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


