package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisProcComponent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_proc_component")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_proc_component

	for rows.Next() {
		var seis_proc_component dto.Seis_proc_component
		if err := rows.Scan(&seis_proc_component.Seis_set_subtype, &seis_proc_component.Seis_proc_set_id, &seis_proc_component.Component_id, &seis_proc_component.Active_ind, &seis_proc_component.Bin_grid_id, &seis_proc_component.Bin_grid_source, &seis_proc_component.Comp_seis_set_id, &seis_proc_component.Comp_seis_set_subtype, &seis_proc_component.Effective_date, &seis_proc_component.Expiry_date, &seis_proc_component.Information_item_id, &seis_proc_component.Info_item_subtype, &seis_proc_component.Input_ind, &seis_proc_component.Output_ind, &seis_proc_component.Ppdm_guid, &seis_proc_component.Processing_component_type, &seis_proc_component.Remark, &seis_proc_component.Seis_proc_plan_id, &seis_proc_component.Source, &seis_proc_component.Uwi, &seis_proc_component.Velocity_volume_id, &seis_proc_component.Well_log_curve_id, &seis_proc_component.Row_changed_by, &seis_proc_component.Row_changed_date, &seis_proc_component.Row_created_by, &seis_proc_component.Row_created_date, &seis_proc_component.Row_effective_date, &seis_proc_component.Row_expiry_date, &seis_proc_component.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_proc_component to result
		result = append(result, seis_proc_component)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisProcComponent(c *fiber.Ctx) error {
	var seis_proc_component dto.Seis_proc_component

	setDefaults(&seis_proc_component)

	if err := json.Unmarshal(c.Body(), &seis_proc_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_proc_component values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	seis_proc_component.Row_created_date = formatDate(seis_proc_component.Row_created_date)
	seis_proc_component.Row_changed_date = formatDate(seis_proc_component.Row_changed_date)
	seis_proc_component.Effective_date = formatDateString(seis_proc_component.Effective_date)
	seis_proc_component.Expiry_date = formatDateString(seis_proc_component.Expiry_date)
	seis_proc_component.Row_effective_date = formatDateString(seis_proc_component.Row_effective_date)
	seis_proc_component.Row_expiry_date = formatDateString(seis_proc_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_component.Seis_set_subtype, seis_proc_component.Seis_proc_set_id, seis_proc_component.Component_id, seis_proc_component.Active_ind, seis_proc_component.Bin_grid_id, seis_proc_component.Bin_grid_source, seis_proc_component.Comp_seis_set_id, seis_proc_component.Comp_seis_set_subtype, seis_proc_component.Effective_date, seis_proc_component.Expiry_date, seis_proc_component.Information_item_id, seis_proc_component.Info_item_subtype, seis_proc_component.Input_ind, seis_proc_component.Output_ind, seis_proc_component.Ppdm_guid, seis_proc_component.Processing_component_type, seis_proc_component.Remark, seis_proc_component.Seis_proc_plan_id, seis_proc_component.Source, seis_proc_component.Uwi, seis_proc_component.Velocity_volume_id, seis_proc_component.Well_log_curve_id, seis_proc_component.Row_changed_by, seis_proc_component.Row_changed_date, seis_proc_component.Row_created_by, seis_proc_component.Row_created_date, seis_proc_component.Row_effective_date, seis_proc_component.Row_expiry_date, seis_proc_component.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisProcComponent(c *fiber.Ctx) error {
	var seis_proc_component dto.Seis_proc_component

	setDefaults(&seis_proc_component)

	if err := json.Unmarshal(c.Body(), &seis_proc_component); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_proc_component.Seis_set_subtype = id

    if seis_proc_component.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_proc_component set seis_proc_set_id = :1, component_id = :2, active_ind = :3, bin_grid_id = :4, bin_grid_source = :5, comp_seis_set_id = :6, comp_seis_set_subtype = :7, effective_date = :8, expiry_date = :9, information_item_id = :10, info_item_subtype = :11, input_ind = :12, output_ind = :13, ppdm_guid = :14, processing_component_type = :15, remark = :16, seis_proc_plan_id = :17, source = :18, uwi = :19, velocity_volume_id = :20, well_log_curve_id = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where seis_set_subtype = :29")
	if err != nil {
		return err
	}

	seis_proc_component.Row_changed_date = formatDate(seis_proc_component.Row_changed_date)
	seis_proc_component.Effective_date = formatDateString(seis_proc_component.Effective_date)
	seis_proc_component.Expiry_date = formatDateString(seis_proc_component.Expiry_date)
	seis_proc_component.Row_effective_date = formatDateString(seis_proc_component.Row_effective_date)
	seis_proc_component.Row_expiry_date = formatDateString(seis_proc_component.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_component.Seis_proc_set_id, seis_proc_component.Component_id, seis_proc_component.Active_ind, seis_proc_component.Bin_grid_id, seis_proc_component.Bin_grid_source, seis_proc_component.Comp_seis_set_id, seis_proc_component.Comp_seis_set_subtype, seis_proc_component.Effective_date, seis_proc_component.Expiry_date, seis_proc_component.Information_item_id, seis_proc_component.Info_item_subtype, seis_proc_component.Input_ind, seis_proc_component.Output_ind, seis_proc_component.Ppdm_guid, seis_proc_component.Processing_component_type, seis_proc_component.Remark, seis_proc_component.Seis_proc_plan_id, seis_proc_component.Source, seis_proc_component.Uwi, seis_proc_component.Velocity_volume_id, seis_proc_component.Well_log_curve_id, seis_proc_component.Row_changed_by, seis_proc_component.Row_changed_date, seis_proc_component.Row_created_by, seis_proc_component.Row_effective_date, seis_proc_component.Row_expiry_date, seis_proc_component.Row_quality, seis_proc_component.Seis_set_subtype)
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

func PatchSeisProcComponent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_proc_component set "
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

func DeleteSeisProcComponent(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_proc_component dto.Seis_proc_component
	seis_proc_component.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_proc_component where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_proc_component.Seis_set_subtype)
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


