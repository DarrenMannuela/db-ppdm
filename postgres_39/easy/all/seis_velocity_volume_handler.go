package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisVelocityVolume(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_velocity_volume")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_velocity_volume

	for rows.Next() {
		var seis_velocity_volume dto.Seis_velocity_volume
		if err := rows.Scan(&seis_velocity_volume.Velocity_volume_id, &seis_velocity_volume.Active_ind, &seis_velocity_volume.Bin_grid_id, &seis_velocity_volume.Bin_grid_seis_set_id, &seis_velocity_volume.Bin_grid_seis_set_subtype, &seis_velocity_volume.Bin_grid_source, &seis_velocity_volume.Created_date, &seis_velocity_volume.Effective_date, &seis_velocity_volume.Expiry_date, &seis_velocity_volume.Picked_by, &seis_velocity_volume.Ppdm_guid, &seis_velocity_volume.Remark, &seis_velocity_volume.Seis_well_set_id, &seis_velocity_volume.Seis_well_set_subtype, &seis_velocity_volume.Source, &seis_velocity_volume.Velocity_dimension, &seis_velocity_volume.Velocity_origin, &seis_velocity_volume.Row_changed_by, &seis_velocity_volume.Row_changed_date, &seis_velocity_volume.Row_created_by, &seis_velocity_volume.Row_created_date, &seis_velocity_volume.Row_effective_date, &seis_velocity_volume.Row_expiry_date, &seis_velocity_volume.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_velocity_volume to result
		result = append(result, seis_velocity_volume)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisVelocityVolume(c *fiber.Ctx) error {
	var seis_velocity_volume dto.Seis_velocity_volume

	setDefaults(&seis_velocity_volume)

	if err := json.Unmarshal(c.Body(), &seis_velocity_volume); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_velocity_volume values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	seis_velocity_volume.Row_created_date = formatDate(seis_velocity_volume.Row_created_date)
	seis_velocity_volume.Row_changed_date = formatDate(seis_velocity_volume.Row_changed_date)
	seis_velocity_volume.Created_date = formatDateString(seis_velocity_volume.Created_date)
	seis_velocity_volume.Effective_date = formatDateString(seis_velocity_volume.Effective_date)
	seis_velocity_volume.Expiry_date = formatDateString(seis_velocity_volume.Expiry_date)
	seis_velocity_volume.Row_effective_date = formatDateString(seis_velocity_volume.Row_effective_date)
	seis_velocity_volume.Row_expiry_date = formatDateString(seis_velocity_volume.Row_expiry_date)






	rows, err := stmt.Exec(seis_velocity_volume.Velocity_volume_id, seis_velocity_volume.Active_ind, seis_velocity_volume.Bin_grid_id, seis_velocity_volume.Bin_grid_seis_set_id, seis_velocity_volume.Bin_grid_seis_set_subtype, seis_velocity_volume.Bin_grid_source, seis_velocity_volume.Created_date, seis_velocity_volume.Effective_date, seis_velocity_volume.Expiry_date, seis_velocity_volume.Picked_by, seis_velocity_volume.Ppdm_guid, seis_velocity_volume.Remark, seis_velocity_volume.Seis_well_set_id, seis_velocity_volume.Seis_well_set_subtype, seis_velocity_volume.Source, seis_velocity_volume.Velocity_dimension, seis_velocity_volume.Velocity_origin, seis_velocity_volume.Row_changed_by, seis_velocity_volume.Row_changed_date, seis_velocity_volume.Row_created_by, seis_velocity_volume.Row_created_date, seis_velocity_volume.Row_effective_date, seis_velocity_volume.Row_expiry_date, seis_velocity_volume.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisVelocityVolume(c *fiber.Ctx) error {
	var seis_velocity_volume dto.Seis_velocity_volume

	setDefaults(&seis_velocity_volume)

	if err := json.Unmarshal(c.Body(), &seis_velocity_volume); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_velocity_volume.Velocity_volume_id = id

    if seis_velocity_volume.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_velocity_volume set active_ind = :1, bin_grid_id = :2, bin_grid_seis_set_id = :3, bin_grid_seis_set_subtype = :4, bin_grid_source = :5, created_date = :6, effective_date = :7, expiry_date = :8, picked_by = :9, ppdm_guid = :10, remark = :11, seis_well_set_id = :12, seis_well_set_subtype = :13, source = :14, velocity_dimension = :15, velocity_origin = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where velocity_volume_id = :24")
	if err != nil {
		return err
	}

	seis_velocity_volume.Row_changed_date = formatDate(seis_velocity_volume.Row_changed_date)
	seis_velocity_volume.Created_date = formatDateString(seis_velocity_volume.Created_date)
	seis_velocity_volume.Effective_date = formatDateString(seis_velocity_volume.Effective_date)
	seis_velocity_volume.Expiry_date = formatDateString(seis_velocity_volume.Expiry_date)
	seis_velocity_volume.Row_effective_date = formatDateString(seis_velocity_volume.Row_effective_date)
	seis_velocity_volume.Row_expiry_date = formatDateString(seis_velocity_volume.Row_expiry_date)






	rows, err := stmt.Exec(seis_velocity_volume.Active_ind, seis_velocity_volume.Bin_grid_id, seis_velocity_volume.Bin_grid_seis_set_id, seis_velocity_volume.Bin_grid_seis_set_subtype, seis_velocity_volume.Bin_grid_source, seis_velocity_volume.Created_date, seis_velocity_volume.Effective_date, seis_velocity_volume.Expiry_date, seis_velocity_volume.Picked_by, seis_velocity_volume.Ppdm_guid, seis_velocity_volume.Remark, seis_velocity_volume.Seis_well_set_id, seis_velocity_volume.Seis_well_set_subtype, seis_velocity_volume.Source, seis_velocity_volume.Velocity_dimension, seis_velocity_volume.Velocity_origin, seis_velocity_volume.Row_changed_by, seis_velocity_volume.Row_changed_date, seis_velocity_volume.Row_created_by, seis_velocity_volume.Row_effective_date, seis_velocity_volume.Row_expiry_date, seis_velocity_volume.Row_quality, seis_velocity_volume.Velocity_volume_id)
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

func PatchSeisVelocityVolume(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_velocity_volume set "
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
		} else if key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSeisVelocityVolume(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_velocity_volume dto.Seis_velocity_volume
	seis_velocity_volume.Velocity_volume_id = id

	stmt, err := db.Prepare("delete from seis_velocity_volume where velocity_volume_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_velocity_volume.Velocity_volume_id)
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


