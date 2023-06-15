package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmMapLoad(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_map_load")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_map_load

	for rows.Next() {
		var ppdm_map_load dto.Ppdm_map_load
		if err := rows.Scan(&ppdm_map_load.Map_id, &ppdm_map_load.Load_process_id, &ppdm_map_load.Active_ind, &ppdm_map_load.Delete_allowed_ind, &ppdm_map_load.Effective_date, &ppdm_map_load.End_date, &ppdm_map_load.Expiry_date, &ppdm_map_load.Insert_allowed_ind, &ppdm_map_load.Ppdm_group_id, &ppdm_map_load.Ppdm_guid, &ppdm_map_load.Procedure_id, &ppdm_map_load.Procedure_name, &ppdm_map_load.Remark, &ppdm_map_load.Source, &ppdm_map_load.Source_system_id, &ppdm_map_load.Start_date, &ppdm_map_load.Sw_application_id, &ppdm_map_load.Target_system_id, &ppdm_map_load.Update_allowed_ind, &ppdm_map_load.Row_changed_by, &ppdm_map_load.Row_changed_date, &ppdm_map_load.Row_created_by, &ppdm_map_load.Row_created_date, &ppdm_map_load.Row_effective_date, &ppdm_map_load.Row_expiry_date, &ppdm_map_load.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_map_load to result
		result = append(result, ppdm_map_load)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmMapLoad(c *fiber.Ctx) error {
	var ppdm_map_load dto.Ppdm_map_load

	setDefaults(&ppdm_map_load)

	if err := json.Unmarshal(c.Body(), &ppdm_map_load); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_map_load values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	ppdm_map_load.Row_created_date = formatDate(ppdm_map_load.Row_created_date)
	ppdm_map_load.Row_changed_date = formatDate(ppdm_map_load.Row_changed_date)
	ppdm_map_load.Effective_date = formatDateString(ppdm_map_load.Effective_date)
	ppdm_map_load.End_date = formatDateString(ppdm_map_load.End_date)
	ppdm_map_load.Expiry_date = formatDateString(ppdm_map_load.Expiry_date)
	ppdm_map_load.Start_date = formatDateString(ppdm_map_load.Start_date)
	ppdm_map_load.Row_effective_date = formatDateString(ppdm_map_load.Row_effective_date)
	ppdm_map_load.Row_expiry_date = formatDateString(ppdm_map_load.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_map_load.Map_id, ppdm_map_load.Load_process_id, ppdm_map_load.Active_ind, ppdm_map_load.Delete_allowed_ind, ppdm_map_load.Effective_date, ppdm_map_load.End_date, ppdm_map_load.Expiry_date, ppdm_map_load.Insert_allowed_ind, ppdm_map_load.Ppdm_group_id, ppdm_map_load.Ppdm_guid, ppdm_map_load.Procedure_id, ppdm_map_load.Procedure_name, ppdm_map_load.Remark, ppdm_map_load.Source, ppdm_map_load.Source_system_id, ppdm_map_load.Start_date, ppdm_map_load.Sw_application_id, ppdm_map_load.Target_system_id, ppdm_map_load.Update_allowed_ind, ppdm_map_load.Row_changed_by, ppdm_map_load.Row_changed_date, ppdm_map_load.Row_created_by, ppdm_map_load.Row_created_date, ppdm_map_load.Row_effective_date, ppdm_map_load.Row_expiry_date, ppdm_map_load.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmMapLoad(c *fiber.Ctx) error {
	var ppdm_map_load dto.Ppdm_map_load

	setDefaults(&ppdm_map_load)

	if err := json.Unmarshal(c.Body(), &ppdm_map_load); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_map_load.Map_id = id

    if ppdm_map_load.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_map_load set load_process_id = :1, active_ind = :2, delete_allowed_ind = :3, effective_date = :4, end_date = :5, expiry_date = :6, insert_allowed_ind = :7, ppdm_group_id = :8, ppdm_guid = :9, procedure_id = :10, procedure_name = :11, remark = :12, source = :13, source_system_id = :14, start_date = :15, sw_application_id = :16, target_system_id = :17, update_allowed_ind = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where map_id = :26")
	if err != nil {
		return err
	}

	ppdm_map_load.Row_changed_date = formatDate(ppdm_map_load.Row_changed_date)
	ppdm_map_load.Effective_date = formatDateString(ppdm_map_load.Effective_date)
	ppdm_map_load.End_date = formatDateString(ppdm_map_load.End_date)
	ppdm_map_load.Expiry_date = formatDateString(ppdm_map_load.Expiry_date)
	ppdm_map_load.Start_date = formatDateString(ppdm_map_load.Start_date)
	ppdm_map_load.Row_effective_date = formatDateString(ppdm_map_load.Row_effective_date)
	ppdm_map_load.Row_expiry_date = formatDateString(ppdm_map_load.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_map_load.Load_process_id, ppdm_map_load.Active_ind, ppdm_map_load.Delete_allowed_ind, ppdm_map_load.Effective_date, ppdm_map_load.End_date, ppdm_map_load.Expiry_date, ppdm_map_load.Insert_allowed_ind, ppdm_map_load.Ppdm_group_id, ppdm_map_load.Ppdm_guid, ppdm_map_load.Procedure_id, ppdm_map_load.Procedure_name, ppdm_map_load.Remark, ppdm_map_load.Source, ppdm_map_load.Source_system_id, ppdm_map_load.Start_date, ppdm_map_load.Sw_application_id, ppdm_map_load.Target_system_id, ppdm_map_load.Update_allowed_ind, ppdm_map_load.Row_changed_by, ppdm_map_load.Row_changed_date, ppdm_map_load.Row_created_by, ppdm_map_load.Row_effective_date, ppdm_map_load.Row_expiry_date, ppdm_map_load.Row_quality, ppdm_map_load.Map_id)
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

func PatchPpdmMapLoad(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_map_load set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePpdmMapLoad(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_map_load dto.Ppdm_map_load
	ppdm_map_load.Map_id = id

	stmt, err := db.Prepare("delete from ppdm_map_load where map_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_map_load.Map_id)
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


