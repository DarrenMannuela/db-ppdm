package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisProcStepComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_proc_step_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_proc_step_component

	for rows.Next() {
		var seis_proc_step_component dto.Seis_proc_step_component
		if err := rows.Scan(&seis_proc_step_component.Seis_set_subtype, &seis_proc_step_component.Seis_proc_set_id, &seis_proc_step_component.Component_id, &seis_proc_step_component.Process_step_id, &seis_proc_step_component.Active_ind, &seis_proc_step_component.Effective_date, &seis_proc_step_component.Expiry_date, &seis_proc_step_component.Input_ind, &seis_proc_step_component.Output_ind, &seis_proc_step_component.Ppdm_guid, &seis_proc_step_component.Remark, &seis_proc_step_component.Source, &seis_proc_step_component.Row_changed_by, &seis_proc_step_component.Row_changed_date, &seis_proc_step_component.Row_created_by, &seis_proc_step_component.Row_created_date, &seis_proc_step_component.Row_effective_date, &seis_proc_step_component.Row_expiry_date, &seis_proc_step_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_proc_step_component to result
		result = append(result, seis_proc_step_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisProcStepComponent(c *fiber.Ctx) error {
	var seis_proc_step_component dto.Seis_proc_step_component

	setDefaults(&seis_proc_step_component)

	if err := json.Unmarshal(c.Body(), &seis_proc_step_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_proc_step_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	seis_proc_step_component.Row_created_date = formatDate(seis_proc_step_component.Row_created_date)
	seis_proc_step_component.Row_changed_date = formatDate(seis_proc_step_component.Row_changed_date)
	seis_proc_step_component.Effective_date = formatDateString(seis_proc_step_component.Effective_date)
	seis_proc_step_component.Expiry_date = formatDateString(seis_proc_step_component.Expiry_date)
	seis_proc_step_component.Row_effective_date = formatDateString(seis_proc_step_component.Row_effective_date)
	seis_proc_step_component.Row_expiry_date = formatDateString(seis_proc_step_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_step_component.Seis_set_subtype, seis_proc_step_component.Seis_proc_set_id, seis_proc_step_component.Component_id, seis_proc_step_component.Process_step_id, seis_proc_step_component.Active_ind, seis_proc_step_component.Effective_date, seis_proc_step_component.Expiry_date, seis_proc_step_component.Input_ind, seis_proc_step_component.Output_ind, seis_proc_step_component.Ppdm_guid, seis_proc_step_component.Remark, seis_proc_step_component.Source, seis_proc_step_component.Row_changed_by, seis_proc_step_component.Row_changed_date, seis_proc_step_component.Row_created_by, seis_proc_step_component.Row_created_date, seis_proc_step_component.Row_effective_date, seis_proc_step_component.Row_expiry_date, seis_proc_step_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisProcStepComponent(c *fiber.Ctx) error {
	var seis_proc_step_component dto.Seis_proc_step_component

	setDefaults(&seis_proc_step_component)

	if err := json.Unmarshal(c.Body(), &seis_proc_step_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_proc_step_component.Seis_set_subtype = id

    if seis_proc_step_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_proc_step_component set seis_proc_set_id = :1, component_id = :2, process_step_id = :3, active_ind = :4, effective_date = :5, expiry_date = :6, input_ind = :7, output_ind = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where seis_set_subtype = :19")
	if err != nil {
		return err
	}

	seis_proc_step_component.Row_changed_date = formatDate(seis_proc_step_component.Row_changed_date)
	seis_proc_step_component.Effective_date = formatDateString(seis_proc_step_component.Effective_date)
	seis_proc_step_component.Expiry_date = formatDateString(seis_proc_step_component.Expiry_date)
	seis_proc_step_component.Row_effective_date = formatDateString(seis_proc_step_component.Row_effective_date)
	seis_proc_step_component.Row_expiry_date = formatDateString(seis_proc_step_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_step_component.Seis_proc_set_id, seis_proc_step_component.Component_id, seis_proc_step_component.Process_step_id, seis_proc_step_component.Active_ind, seis_proc_step_component.Effective_date, seis_proc_step_component.Expiry_date, seis_proc_step_component.Input_ind, seis_proc_step_component.Output_ind, seis_proc_step_component.Ppdm_guid, seis_proc_step_component.Remark, seis_proc_step_component.Source, seis_proc_step_component.Row_changed_by, seis_proc_step_component.Row_changed_date, seis_proc_step_component.Row_created_by, seis_proc_step_component.Row_effective_date, seis_proc_step_component.Row_expiry_date, seis_proc_step_component.Row_quality, seis_proc_step_component.Seis_set_subtype)
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

func PatchSeisProcStepComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_proc_step_component set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeisProcStepComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_proc_step_component dto.Seis_proc_step_component
	seis_proc_step_component.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_proc_step_component where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_proc_step_component.Seis_set_subtype)
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


