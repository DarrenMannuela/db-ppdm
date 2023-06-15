package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmMapLoadError(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_map_load_error")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_map_load_error

	for rows.Next() {
		var ppdm_map_load_error dto.Ppdm_map_load_error
		if err := rows.Scan(&ppdm_map_load_error.Map_id, &ppdm_map_load_error.Load_process_id, &ppdm_map_load_error.Error_id, &ppdm_map_load_error.Active_ind, &ppdm_map_load_error.Delete_error_ind, &ppdm_map_load_error.Effective_date, &ppdm_map_load_error.Error_cause_desc, &ppdm_map_load_error.Error_cause_type, &ppdm_map_load_error.Error_code, &ppdm_map_load_error.Error_date, &ppdm_map_load_error.Error_message, &ppdm_map_load_error.Expiry_date, &ppdm_map_load_error.Insert_error_ind, &ppdm_map_load_error.Ppdm_guid, &ppdm_map_load_error.Ref_map_detail_obs_no, &ppdm_map_load_error.Remark, &ppdm_map_load_error.Resolution_desc, &ppdm_map_load_error.Resolution_type, &ppdm_map_load_error.Source, &ppdm_map_load_error.Update_error_ind, &ppdm_map_load_error.Row_changed_by, &ppdm_map_load_error.Row_changed_date, &ppdm_map_load_error.Row_created_by, &ppdm_map_load_error.Row_created_date, &ppdm_map_load_error.Row_effective_date, &ppdm_map_load_error.Row_expiry_date, &ppdm_map_load_error.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_map_load_error to result
		result = append(result, ppdm_map_load_error)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmMapLoadError(c *fiber.Ctx) error {
	var ppdm_map_load_error dto.Ppdm_map_load_error

	setDefaults(&ppdm_map_load_error)

	if err := json.Unmarshal(c.Body(), &ppdm_map_load_error); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_map_load_error values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	ppdm_map_load_error.Row_created_date = formatDate(ppdm_map_load_error.Row_created_date)
	ppdm_map_load_error.Row_changed_date = formatDate(ppdm_map_load_error.Row_changed_date)
	ppdm_map_load_error.Effective_date = formatDateString(ppdm_map_load_error.Effective_date)
	ppdm_map_load_error.Error_date = formatDateString(ppdm_map_load_error.Error_date)
	ppdm_map_load_error.Expiry_date = formatDateString(ppdm_map_load_error.Expiry_date)
	ppdm_map_load_error.Row_effective_date = formatDateString(ppdm_map_load_error.Row_effective_date)
	ppdm_map_load_error.Row_expiry_date = formatDateString(ppdm_map_load_error.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_map_load_error.Map_id, ppdm_map_load_error.Load_process_id, ppdm_map_load_error.Error_id, ppdm_map_load_error.Active_ind, ppdm_map_load_error.Delete_error_ind, ppdm_map_load_error.Effective_date, ppdm_map_load_error.Error_cause_desc, ppdm_map_load_error.Error_cause_type, ppdm_map_load_error.Error_code, ppdm_map_load_error.Error_date, ppdm_map_load_error.Error_message, ppdm_map_load_error.Expiry_date, ppdm_map_load_error.Insert_error_ind, ppdm_map_load_error.Ppdm_guid, ppdm_map_load_error.Ref_map_detail_obs_no, ppdm_map_load_error.Remark, ppdm_map_load_error.Resolution_desc, ppdm_map_load_error.Resolution_type, ppdm_map_load_error.Source, ppdm_map_load_error.Update_error_ind, ppdm_map_load_error.Row_changed_by, ppdm_map_load_error.Row_changed_date, ppdm_map_load_error.Row_created_by, ppdm_map_load_error.Row_created_date, ppdm_map_load_error.Row_effective_date, ppdm_map_load_error.Row_expiry_date, ppdm_map_load_error.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmMapLoadError(c *fiber.Ctx) error {
	var ppdm_map_load_error dto.Ppdm_map_load_error

	setDefaults(&ppdm_map_load_error)

	if err := json.Unmarshal(c.Body(), &ppdm_map_load_error); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_map_load_error.Map_id = id

    if ppdm_map_load_error.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_map_load_error set load_process_id = :1, error_id = :2, active_ind = :3, delete_error_ind = :4, effective_date = :5, error_cause_desc = :6, error_cause_type = :7, error_code = :8, error_date = :9, error_message = :10, expiry_date = :11, insert_error_ind = :12, ppdm_guid = :13, ref_map_detail_obs_no = :14, remark = :15, resolution_desc = :16, resolution_type = :17, source = :18, update_error_ind = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where map_id = :27")
	if err != nil {
		return err
	}

	ppdm_map_load_error.Row_changed_date = formatDate(ppdm_map_load_error.Row_changed_date)
	ppdm_map_load_error.Effective_date = formatDateString(ppdm_map_load_error.Effective_date)
	ppdm_map_load_error.Error_date = formatDateString(ppdm_map_load_error.Error_date)
	ppdm_map_load_error.Expiry_date = formatDateString(ppdm_map_load_error.Expiry_date)
	ppdm_map_load_error.Row_effective_date = formatDateString(ppdm_map_load_error.Row_effective_date)
	ppdm_map_load_error.Row_expiry_date = formatDateString(ppdm_map_load_error.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_map_load_error.Load_process_id, ppdm_map_load_error.Error_id, ppdm_map_load_error.Active_ind, ppdm_map_load_error.Delete_error_ind, ppdm_map_load_error.Effective_date, ppdm_map_load_error.Error_cause_desc, ppdm_map_load_error.Error_cause_type, ppdm_map_load_error.Error_code, ppdm_map_load_error.Error_date, ppdm_map_load_error.Error_message, ppdm_map_load_error.Expiry_date, ppdm_map_load_error.Insert_error_ind, ppdm_map_load_error.Ppdm_guid, ppdm_map_load_error.Ref_map_detail_obs_no, ppdm_map_load_error.Remark, ppdm_map_load_error.Resolution_desc, ppdm_map_load_error.Resolution_type, ppdm_map_load_error.Source, ppdm_map_load_error.Update_error_ind, ppdm_map_load_error.Row_changed_by, ppdm_map_load_error.Row_changed_date, ppdm_map_load_error.Row_created_by, ppdm_map_load_error.Row_effective_date, ppdm_map_load_error.Row_expiry_date, ppdm_map_load_error.Row_quality, ppdm_map_load_error.Map_id)
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

func PatchPpdmMapLoadError(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_map_load_error set "
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
		} else if key == "effective_date" || key == "error_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where map_id = :id"

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

func DeletePpdmMapLoadError(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_map_load_error dto.Ppdm_map_load_error
	ppdm_map_load_error.Map_id = id

	stmt, err := db.Prepare("delete from ppdm_map_load_error where map_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_map_load_error.Map_id)
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


