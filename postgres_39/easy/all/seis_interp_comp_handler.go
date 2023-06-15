package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisInterpComp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_interp_comp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_interp_comp

	for rows.Next() {
		var seis_interp_comp dto.Seis_interp_comp
		if err := rows.Scan(&seis_interp_comp.Seis_set_subtype, &seis_interp_comp.Interp_set_id, &seis_interp_comp.Component_id, &seis_interp_comp.Active_ind, &seis_interp_comp.Area_id, &seis_interp_comp.Area_type, &seis_interp_comp.Coord_acquisition_id, &seis_interp_comp.Coord_system_id, &seis_interp_comp.Corner1_lat, &seis_interp_comp.Corner1_long, &seis_interp_comp.Corner2_lat, &seis_interp_comp.Corner2_long, &seis_interp_comp.Corner3_lat, &seis_interp_comp.Corner3_long, &seis_interp_comp.Data_sample_size, &seis_interp_comp.Effective_date, &seis_interp_comp.Expiry_date, &seis_interp_comp.Information_item_id, &seis_interp_comp.Info_item_subtype, &seis_interp_comp.Input_ind, &seis_interp_comp.Local_coord_system_id, &seis_interp_comp.Origin_type, &seis_interp_comp.Output_ind, &seis_interp_comp.Ppdm_guid, &seis_interp_comp.Process_step_id, &seis_interp_comp.Proc_component_id, &seis_interp_comp.Proc_set_id, &seis_interp_comp.Proc_set_subtype, &seis_interp_comp.Remark, &seis_interp_comp.Source, &seis_interp_comp.Uwi, &seis_interp_comp.Velocity_volume_id, &seis_interp_comp.Row_changed_by, &seis_interp_comp.Row_changed_date, &seis_interp_comp.Row_created_by, &seis_interp_comp.Row_created_date, &seis_interp_comp.Row_effective_date, &seis_interp_comp.Row_expiry_date, &seis_interp_comp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_interp_comp to result
		result = append(result, seis_interp_comp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisInterpComp(c *fiber.Ctx) error {
	var seis_interp_comp dto.Seis_interp_comp

	setDefaults(&seis_interp_comp)

	if err := json.Unmarshal(c.Body(), &seis_interp_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_interp_comp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	seis_interp_comp.Row_created_date = formatDate(seis_interp_comp.Row_created_date)
	seis_interp_comp.Row_changed_date = formatDate(seis_interp_comp.Row_changed_date)
	seis_interp_comp.Effective_date = formatDateString(seis_interp_comp.Effective_date)
	seis_interp_comp.Expiry_date = formatDateString(seis_interp_comp.Expiry_date)
	seis_interp_comp.Row_effective_date = formatDateString(seis_interp_comp.Row_effective_date)
	seis_interp_comp.Row_expiry_date = formatDateString(seis_interp_comp.Row_expiry_date)






	rows, err := stmt.Exec(seis_interp_comp.Seis_set_subtype, seis_interp_comp.Interp_set_id, seis_interp_comp.Component_id, seis_interp_comp.Active_ind, seis_interp_comp.Area_id, seis_interp_comp.Area_type, seis_interp_comp.Coord_acquisition_id, seis_interp_comp.Coord_system_id, seis_interp_comp.Corner1_lat, seis_interp_comp.Corner1_long, seis_interp_comp.Corner2_lat, seis_interp_comp.Corner2_long, seis_interp_comp.Corner3_lat, seis_interp_comp.Corner3_long, seis_interp_comp.Data_sample_size, seis_interp_comp.Effective_date, seis_interp_comp.Expiry_date, seis_interp_comp.Information_item_id, seis_interp_comp.Info_item_subtype, seis_interp_comp.Input_ind, seis_interp_comp.Local_coord_system_id, seis_interp_comp.Origin_type, seis_interp_comp.Output_ind, seis_interp_comp.Ppdm_guid, seis_interp_comp.Process_step_id, seis_interp_comp.Proc_component_id, seis_interp_comp.Proc_set_id, seis_interp_comp.Proc_set_subtype, seis_interp_comp.Remark, seis_interp_comp.Source, seis_interp_comp.Uwi, seis_interp_comp.Velocity_volume_id, seis_interp_comp.Row_changed_by, seis_interp_comp.Row_changed_date, seis_interp_comp.Row_created_by, seis_interp_comp.Row_created_date, seis_interp_comp.Row_effective_date, seis_interp_comp.Row_expiry_date, seis_interp_comp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisInterpComp(c *fiber.Ctx) error {
	var seis_interp_comp dto.Seis_interp_comp

	setDefaults(&seis_interp_comp)

	if err := json.Unmarshal(c.Body(), &seis_interp_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_interp_comp.Seis_set_subtype = id

    if seis_interp_comp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_interp_comp set interp_set_id = :1, component_id = :2, active_ind = :3, area_id = :4, area_type = :5, coord_acquisition_id = :6, coord_system_id = :7, corner1_lat = :8, corner1_long = :9, corner2_lat = :10, corner2_long = :11, corner3_lat = :12, corner3_long = :13, data_sample_size = :14, effective_date = :15, expiry_date = :16, information_item_id = :17, info_item_subtype = :18, input_ind = :19, local_coord_system_id = :20, origin_type = :21, output_ind = :22, ppdm_guid = :23, process_step_id = :24, proc_component_id = :25, proc_set_id = :26, proc_set_subtype = :27, remark = :28, source = :29, uwi = :30, velocity_volume_id = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where seis_set_subtype = :39")
	if err != nil {
		return err
	}

	seis_interp_comp.Row_changed_date = formatDate(seis_interp_comp.Row_changed_date)
	seis_interp_comp.Effective_date = formatDateString(seis_interp_comp.Effective_date)
	seis_interp_comp.Expiry_date = formatDateString(seis_interp_comp.Expiry_date)
	seis_interp_comp.Row_effective_date = formatDateString(seis_interp_comp.Row_effective_date)
	seis_interp_comp.Row_expiry_date = formatDateString(seis_interp_comp.Row_expiry_date)






	rows, err := stmt.Exec(seis_interp_comp.Interp_set_id, seis_interp_comp.Component_id, seis_interp_comp.Active_ind, seis_interp_comp.Area_id, seis_interp_comp.Area_type, seis_interp_comp.Coord_acquisition_id, seis_interp_comp.Coord_system_id, seis_interp_comp.Corner1_lat, seis_interp_comp.Corner1_long, seis_interp_comp.Corner2_lat, seis_interp_comp.Corner2_long, seis_interp_comp.Corner3_lat, seis_interp_comp.Corner3_long, seis_interp_comp.Data_sample_size, seis_interp_comp.Effective_date, seis_interp_comp.Expiry_date, seis_interp_comp.Information_item_id, seis_interp_comp.Info_item_subtype, seis_interp_comp.Input_ind, seis_interp_comp.Local_coord_system_id, seis_interp_comp.Origin_type, seis_interp_comp.Output_ind, seis_interp_comp.Ppdm_guid, seis_interp_comp.Process_step_id, seis_interp_comp.Proc_component_id, seis_interp_comp.Proc_set_id, seis_interp_comp.Proc_set_subtype, seis_interp_comp.Remark, seis_interp_comp.Source, seis_interp_comp.Uwi, seis_interp_comp.Velocity_volume_id, seis_interp_comp.Row_changed_by, seis_interp_comp.Row_changed_date, seis_interp_comp.Row_created_by, seis_interp_comp.Row_effective_date, seis_interp_comp.Row_expiry_date, seis_interp_comp.Row_quality, seis_interp_comp.Seis_set_subtype)
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

func PatchSeisInterpComp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_interp_comp set "
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

func DeleteSeisInterpComp(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_interp_comp dto.Seis_interp_comp
	seis_interp_comp.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_interp_comp where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_interp_comp.Seis_set_subtype)
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


