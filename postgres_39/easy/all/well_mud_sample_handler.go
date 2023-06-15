package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellMudSample(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_mud_sample")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_mud_sample

	for rows.Next() {
		var well_mud_sample dto.Well_mud_sample
		if err := rows.Scan(&well_mud_sample.Anl_id, &well_mud_sample.Anl_source, &well_mud_sample.Step_seq_no, &well_mud_sample.Mud_anl_obs_no, &well_mud_sample.Active_ind, &well_mud_sample.Circulation_stop_date, &well_mud_sample.Circulation_stop_time, &well_mud_sample.Circulation_stop_timezone, &well_mud_sample.Effective_date, &well_mud_sample.Expiry_date, &well_mud_sample.Fluid_flow, &well_mud_sample.Fluid_flow_ouom, &well_mud_sample.Fluid_loss, &well_mud_sample.Fluid_loss_ouom, &well_mud_sample.Mud_collect_reason, &well_mud_sample.Mud_density, &well_mud_sample.Mud_density_ouom, &well_mud_sample.Mud_density_uom, &well_mud_sample.Mud_ph, &well_mud_sample.Mud_ph_ouom, &well_mud_sample.Mud_sample_type, &well_mud_sample.Mud_type, &well_mud_sample.Mud_viscosity, &well_mud_sample.Mud_viscosity_ouom, &well_mud_sample.Period_obs_no, &well_mud_sample.Ppdm_guid, &well_mud_sample.Preferred_ind, &well_mud_sample.Problem_ind, &well_mud_sample.Remark, &well_mud_sample.Sample_date, &well_mud_sample.Sample_depth, &well_mud_sample.Sample_depth_ouom, &well_mud_sample.Sample_time, &well_mud_sample.Sample_timezone, &well_mud_sample.Source, &well_mud_sample.Source_location, &well_mud_sample.Time_since_circ, &well_mud_sample.Time_since_circ_ouom, &well_mud_sample.Uwi, &well_mud_sample.Row_changed_by, &well_mud_sample.Row_changed_date, &well_mud_sample.Row_created_by, &well_mud_sample.Row_created_date, &well_mud_sample.Row_effective_date, &well_mud_sample.Row_expiry_date, &well_mud_sample.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_mud_sample to result
		result = append(result, well_mud_sample)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellMudSample(c *fiber.Ctx) error {
	var well_mud_sample dto.Well_mud_sample

	setDefaults(&well_mud_sample)

	if err := json.Unmarshal(c.Body(), &well_mud_sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_mud_sample values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46)")
	if err != nil {
		return err
	}
	well_mud_sample.Row_created_date = formatDate(well_mud_sample.Row_created_date)
	well_mud_sample.Row_changed_date = formatDate(well_mud_sample.Row_changed_date)
	well_mud_sample.Circulation_stop_date = formatDateString(well_mud_sample.Circulation_stop_date)
	well_mud_sample.Circulation_stop_time = formatDateString(well_mud_sample.Circulation_stop_time)
	well_mud_sample.Effective_date = formatDateString(well_mud_sample.Effective_date)
	well_mud_sample.Expiry_date = formatDateString(well_mud_sample.Expiry_date)
	well_mud_sample.Sample_date = formatDateString(well_mud_sample.Sample_date)
	well_mud_sample.Sample_time = formatDateString(well_mud_sample.Sample_time)
	well_mud_sample.Row_effective_date = formatDateString(well_mud_sample.Row_effective_date)
	well_mud_sample.Row_expiry_date = formatDateString(well_mud_sample.Row_expiry_date)






	rows, err := stmt.Exec(well_mud_sample.Anl_id, well_mud_sample.Anl_source, well_mud_sample.Step_seq_no, well_mud_sample.Mud_anl_obs_no, well_mud_sample.Active_ind, well_mud_sample.Circulation_stop_date, well_mud_sample.Circulation_stop_time, well_mud_sample.Circulation_stop_timezone, well_mud_sample.Effective_date, well_mud_sample.Expiry_date, well_mud_sample.Fluid_flow, well_mud_sample.Fluid_flow_ouom, well_mud_sample.Fluid_loss, well_mud_sample.Fluid_loss_ouom, well_mud_sample.Mud_collect_reason, well_mud_sample.Mud_density, well_mud_sample.Mud_density_ouom, well_mud_sample.Mud_density_uom, well_mud_sample.Mud_ph, well_mud_sample.Mud_ph_ouom, well_mud_sample.Mud_sample_type, well_mud_sample.Mud_type, well_mud_sample.Mud_viscosity, well_mud_sample.Mud_viscosity_ouom, well_mud_sample.Period_obs_no, well_mud_sample.Ppdm_guid, well_mud_sample.Preferred_ind, well_mud_sample.Problem_ind, well_mud_sample.Remark, well_mud_sample.Sample_date, well_mud_sample.Sample_depth, well_mud_sample.Sample_depth_ouom, well_mud_sample.Sample_time, well_mud_sample.Sample_timezone, well_mud_sample.Source, well_mud_sample.Source_location, well_mud_sample.Time_since_circ, well_mud_sample.Time_since_circ_ouom, well_mud_sample.Uwi, well_mud_sample.Row_changed_by, well_mud_sample.Row_changed_date, well_mud_sample.Row_created_by, well_mud_sample.Row_created_date, well_mud_sample.Row_effective_date, well_mud_sample.Row_expiry_date, well_mud_sample.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellMudSample(c *fiber.Ctx) error {
	var well_mud_sample dto.Well_mud_sample

	setDefaults(&well_mud_sample)

	if err := json.Unmarshal(c.Body(), &well_mud_sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_mud_sample.Anl_id = id

    if well_mud_sample.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_mud_sample set anl_source = :1, step_seq_no = :2, mud_anl_obs_no = :3, active_ind = :4, circulation_stop_date = :5, circulation_stop_time = :6, circulation_stop_timezone = :7, effective_date = :8, expiry_date = :9, fluid_flow = :10, fluid_flow_ouom = :11, fluid_loss = :12, fluid_loss_ouom = :13, mud_collect_reason = :14, mud_density = :15, mud_density_ouom = :16, mud_density_uom = :17, mud_ph = :18, mud_ph_ouom = :19, mud_sample_type = :20, mud_type = :21, mud_viscosity = :22, mud_viscosity_ouom = :23, period_obs_no = :24, ppdm_guid = :25, preferred_ind = :26, problem_ind = :27, remark = :28, sample_date = :29, sample_depth = :30, sample_depth_ouom = :31, sample_time = :32, sample_timezone = :33, source = :34, source_location = :35, time_since_circ = :36, time_since_circ_ouom = :37, uwi = :38, row_changed_by = :39, row_changed_date = :40, row_created_by = :41, row_effective_date = :42, row_expiry_date = :43, row_quality = :44 where anl_id = :46")
	if err != nil {
		return err
	}

	well_mud_sample.Row_changed_date = formatDate(well_mud_sample.Row_changed_date)
	well_mud_sample.Circulation_stop_date = formatDateString(well_mud_sample.Circulation_stop_date)
	well_mud_sample.Circulation_stop_time = formatDateString(well_mud_sample.Circulation_stop_time)
	well_mud_sample.Effective_date = formatDateString(well_mud_sample.Effective_date)
	well_mud_sample.Expiry_date = formatDateString(well_mud_sample.Expiry_date)
	well_mud_sample.Sample_date = formatDateString(well_mud_sample.Sample_date)
	well_mud_sample.Sample_time = formatDateString(well_mud_sample.Sample_time)
	well_mud_sample.Row_effective_date = formatDateString(well_mud_sample.Row_effective_date)
	well_mud_sample.Row_expiry_date = formatDateString(well_mud_sample.Row_expiry_date)






	rows, err := stmt.Exec(well_mud_sample.Anl_source, well_mud_sample.Step_seq_no, well_mud_sample.Mud_anl_obs_no, well_mud_sample.Active_ind, well_mud_sample.Circulation_stop_date, well_mud_sample.Circulation_stop_time, well_mud_sample.Circulation_stop_timezone, well_mud_sample.Effective_date, well_mud_sample.Expiry_date, well_mud_sample.Fluid_flow, well_mud_sample.Fluid_flow_ouom, well_mud_sample.Fluid_loss, well_mud_sample.Fluid_loss_ouom, well_mud_sample.Mud_collect_reason, well_mud_sample.Mud_density, well_mud_sample.Mud_density_ouom, well_mud_sample.Mud_density_uom, well_mud_sample.Mud_ph, well_mud_sample.Mud_ph_ouom, well_mud_sample.Mud_sample_type, well_mud_sample.Mud_type, well_mud_sample.Mud_viscosity, well_mud_sample.Mud_viscosity_ouom, well_mud_sample.Period_obs_no, well_mud_sample.Ppdm_guid, well_mud_sample.Preferred_ind, well_mud_sample.Problem_ind, well_mud_sample.Remark, well_mud_sample.Sample_date, well_mud_sample.Sample_depth, well_mud_sample.Sample_depth_ouom, well_mud_sample.Sample_time, well_mud_sample.Sample_timezone, well_mud_sample.Source, well_mud_sample.Source_location, well_mud_sample.Time_since_circ, well_mud_sample.Time_since_circ_ouom, well_mud_sample.Uwi, well_mud_sample.Row_changed_by, well_mud_sample.Row_changed_date, well_mud_sample.Row_created_by, well_mud_sample.Row_effective_date, well_mud_sample.Row_expiry_date, well_mud_sample.Row_quality, well_mud_sample.Anl_id)
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

func PatchWellMudSample(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_mud_sample set "
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
		} else if key == "circulation_stop_date" || key == "circulation_stop_time" || key == "effective_date" || key == "expiry_date" || key == "sample_date" || key == "sample_time" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where anl_id = :id"

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

func DeleteWellMudSample(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_mud_sample dto.Well_mud_sample
	well_mud_sample.Anl_id = id

	stmt, err := db.Prepare("delete from well_mud_sample where anl_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_mud_sample.Anl_id)
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


