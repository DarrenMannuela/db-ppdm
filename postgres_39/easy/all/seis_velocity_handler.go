package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisVelocity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_velocity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_velocity

	for rows.Next() {
		var seis_velocity dto.Seis_velocity
		if err := rows.Scan(&seis_velocity.Velocity_volume_id, &seis_velocity.Volume_point, &seis_velocity.Active_ind, &seis_velocity.Bin_grid_id, &seis_velocity.Bin_point_id, &seis_velocity.Bin_source, &seis_velocity.Compute_method, &seis_velocity.Effective_date, &seis_velocity.Expiry_date, &seis_velocity.Latitude, &seis_velocity.Longitude, &seis_velocity.Ppdm_guid, &seis_velocity.Receiver_md, &seis_velocity.Receiver_md_ouom, &seis_velocity.Receiver_vert_depth, &seis_velocity.Receiver_vert_depth_ouom, &seis_velocity.Remark, &seis_velocity.Seis_point_id, &seis_velocity.Seis_set_id, &seis_velocity.Seis_set_subtype, &seis_velocity.Seis_time, &seis_velocity.Seis_time_ouom, &seis_velocity.Seis_well_set_id, &seis_velocity.Seis_well_set_subtype, &seis_velocity.Source, &seis_velocity.Source_md, &seis_velocity.Source_md_ouom, &seis_velocity.Source_vert_depth, &seis_velocity.Source_vert_depth_ouom, &seis_velocity.Velocity, &seis_velocity.Velocity_azimuth, &seis_velocity.Velocity_depth, &seis_velocity.Velocity_inclination, &seis_velocity.Velocity_inclination_ouom, &seis_velocity.Velocity_ouom, &seis_velocity.Velocity_type, &seis_velocity.Velocity_x, &seis_velocity.Velocity_y, &seis_velocity.Velocity_z, &seis_velocity.Row_changed_by, &seis_velocity.Row_changed_date, &seis_velocity.Row_created_by, &seis_velocity.Row_created_date, &seis_velocity.Row_effective_date, &seis_velocity.Row_expiry_date, &seis_velocity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_velocity to result
		result = append(result, seis_velocity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisVelocity(c *fiber.Ctx) error {
	var seis_velocity dto.Seis_velocity

	setDefaults(&seis_velocity)

	if err := json.Unmarshal(c.Body(), &seis_velocity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_velocity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46)")
	if err != nil {
		return err
	}
	seis_velocity.Row_created_date = formatDate(seis_velocity.Row_created_date)
	seis_velocity.Row_changed_date = formatDate(seis_velocity.Row_changed_date)
	seis_velocity.Effective_date = formatDateString(seis_velocity.Effective_date)
	seis_velocity.Expiry_date = formatDateString(seis_velocity.Expiry_date)
	seis_velocity.Row_effective_date = formatDateString(seis_velocity.Row_effective_date)
	seis_velocity.Row_expiry_date = formatDateString(seis_velocity.Row_expiry_date)






	rows, err := stmt.Exec(seis_velocity.Velocity_volume_id, seis_velocity.Volume_point, seis_velocity.Active_ind, seis_velocity.Bin_grid_id, seis_velocity.Bin_point_id, seis_velocity.Bin_source, seis_velocity.Compute_method, seis_velocity.Effective_date, seis_velocity.Expiry_date, seis_velocity.Latitude, seis_velocity.Longitude, seis_velocity.Ppdm_guid, seis_velocity.Receiver_md, seis_velocity.Receiver_md_ouom, seis_velocity.Receiver_vert_depth, seis_velocity.Receiver_vert_depth_ouom, seis_velocity.Remark, seis_velocity.Seis_point_id, seis_velocity.Seis_set_id, seis_velocity.Seis_set_subtype, seis_velocity.Seis_time, seis_velocity.Seis_time_ouom, seis_velocity.Seis_well_set_id, seis_velocity.Seis_well_set_subtype, seis_velocity.Source, seis_velocity.Source_md, seis_velocity.Source_md_ouom, seis_velocity.Source_vert_depth, seis_velocity.Source_vert_depth_ouom, seis_velocity.Velocity, seis_velocity.Velocity_azimuth, seis_velocity.Velocity_depth, seis_velocity.Velocity_inclination, seis_velocity.Velocity_inclination_ouom, seis_velocity.Velocity_ouom, seis_velocity.Velocity_type, seis_velocity.Velocity_x, seis_velocity.Velocity_y, seis_velocity.Velocity_z, seis_velocity.Row_changed_by, seis_velocity.Row_changed_date, seis_velocity.Row_created_by, seis_velocity.Row_created_date, seis_velocity.Row_effective_date, seis_velocity.Row_expiry_date, seis_velocity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisVelocity(c *fiber.Ctx) error {
	var seis_velocity dto.Seis_velocity

	setDefaults(&seis_velocity)

	if err := json.Unmarshal(c.Body(), &seis_velocity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_velocity.Velocity_volume_id = id

    if seis_velocity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_velocity set volume_point = :1, active_ind = :2, bin_grid_id = :3, bin_point_id = :4, bin_source = :5, compute_method = :6, effective_date = :7, expiry_date = :8, latitude = :9, longitude = :10, ppdm_guid = :11, receiver_md = :12, receiver_md_ouom = :13, receiver_vert_depth = :14, receiver_vert_depth_ouom = :15, remark = :16, seis_point_id = :17, seis_set_id = :18, seis_set_subtype = :19, seis_time = :20, seis_time_ouom = :21, seis_well_set_id = :22, seis_well_set_subtype = :23, source = :24, source_md = :25, source_md_ouom = :26, source_vert_depth = :27, source_vert_depth_ouom = :28, velocity = :29, velocity_azimuth = :30, velocity_depth = :31, velocity_inclination = :32, velocity_inclination_ouom = :33, velocity_ouom = :34, velocity_type = :35, velocity_x = :36, velocity_y = :37, velocity_z = :38, row_changed_by = :39, row_changed_date = :40, row_created_by = :41, row_effective_date = :42, row_expiry_date = :43, row_quality = :44 where velocity_volume_id = :46")
	if err != nil {
		return err
	}

	seis_velocity.Row_changed_date = formatDate(seis_velocity.Row_changed_date)
	seis_velocity.Effective_date = formatDateString(seis_velocity.Effective_date)
	seis_velocity.Expiry_date = formatDateString(seis_velocity.Expiry_date)
	seis_velocity.Row_effective_date = formatDateString(seis_velocity.Row_effective_date)
	seis_velocity.Row_expiry_date = formatDateString(seis_velocity.Row_expiry_date)






	rows, err := stmt.Exec(seis_velocity.Volume_point, seis_velocity.Active_ind, seis_velocity.Bin_grid_id, seis_velocity.Bin_point_id, seis_velocity.Bin_source, seis_velocity.Compute_method, seis_velocity.Effective_date, seis_velocity.Expiry_date, seis_velocity.Latitude, seis_velocity.Longitude, seis_velocity.Ppdm_guid, seis_velocity.Receiver_md, seis_velocity.Receiver_md_ouom, seis_velocity.Receiver_vert_depth, seis_velocity.Receiver_vert_depth_ouom, seis_velocity.Remark, seis_velocity.Seis_point_id, seis_velocity.Seis_set_id, seis_velocity.Seis_set_subtype, seis_velocity.Seis_time, seis_velocity.Seis_time_ouom, seis_velocity.Seis_well_set_id, seis_velocity.Seis_well_set_subtype, seis_velocity.Source, seis_velocity.Source_md, seis_velocity.Source_md_ouom, seis_velocity.Source_vert_depth, seis_velocity.Source_vert_depth_ouom, seis_velocity.Velocity, seis_velocity.Velocity_azimuth, seis_velocity.Velocity_depth, seis_velocity.Velocity_inclination, seis_velocity.Velocity_inclination_ouom, seis_velocity.Velocity_ouom, seis_velocity.Velocity_type, seis_velocity.Velocity_x, seis_velocity.Velocity_y, seis_velocity.Velocity_z, seis_velocity.Row_changed_by, seis_velocity.Row_changed_date, seis_velocity.Row_created_by, seis_velocity.Row_effective_date, seis_velocity.Row_expiry_date, seis_velocity.Row_quality, seis_velocity.Velocity_volume_id)
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

func PatchSeisVelocity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_velocity set "
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
	query += " where velocity_volume_id = :id"

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

func DeleteSeisVelocity(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_velocity dto.Seis_velocity
	seis_velocity.Velocity_volume_id = id

	stmt, err := db.Prepare("delete from seis_velocity where velocity_volume_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_velocity.Velocity_volume_id)
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


