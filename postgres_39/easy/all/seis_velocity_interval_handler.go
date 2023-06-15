package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisVelocityInterval(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_velocity_interval")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_velocity_interval

	for rows.Next() {
		var seis_velocity_interval dto.Seis_velocity_interval
		if err := rows.Scan(&seis_velocity_interval.Interval_id, &seis_velocity_interval.Active_ind, &seis_velocity_interval.Base_depth, &seis_velocity_interval.Base_depth_ouom, &seis_velocity_interval.Base_strat_unit_id, &seis_velocity_interval.Compute_method, &seis_velocity_interval.Effective_date, &seis_velocity_interval.Expiry_date, &seis_velocity_interval.Last_update, &seis_velocity_interval.Ppdm_guid, &seis_velocity_interval.Remark, &seis_velocity_interval.Reported_tvd, &seis_velocity_interval.Reported_tvd_ouom, &seis_velocity_interval.Seis_set_id, &seis_velocity_interval.Seis_set_subtype, &seis_velocity_interval.Seis_time_datum, &seis_velocity_interval.Seis_time_datum_ouom, &seis_velocity_interval.Source, &seis_velocity_interval.Strat_name_set_id, &seis_velocity_interval.Top_depth, &seis_velocity_interval.Top_depth_ouom, &seis_velocity_interval.Top_strat_unit_id, &seis_velocity_interval.Uwi, &seis_velocity_interval.Velocity_quality, &seis_velocity_interval.Velocity_type, &seis_velocity_interval.Velocity_value, &seis_velocity_interval.Velocity_value_ouom, &seis_velocity_interval.Row_changed_by, &seis_velocity_interval.Row_changed_date, &seis_velocity_interval.Row_created_by, &seis_velocity_interval.Row_created_date, &seis_velocity_interval.Row_effective_date, &seis_velocity_interval.Row_expiry_date, &seis_velocity_interval.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_velocity_interval to result
		result = append(result, seis_velocity_interval)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisVelocityInterval(c *fiber.Ctx) error {
	var seis_velocity_interval dto.Seis_velocity_interval

	setDefaults(&seis_velocity_interval)

	if err := json.Unmarshal(c.Body(), &seis_velocity_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_velocity_interval values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	seis_velocity_interval.Row_created_date = formatDate(seis_velocity_interval.Row_created_date)
	seis_velocity_interval.Row_changed_date = formatDate(seis_velocity_interval.Row_changed_date)
	seis_velocity_interval.Effective_date = formatDateString(seis_velocity_interval.Effective_date)
	seis_velocity_interval.Expiry_date = formatDateString(seis_velocity_interval.Expiry_date)
	seis_velocity_interval.Last_update = formatDateString(seis_velocity_interval.Last_update)
	seis_velocity_interval.Row_effective_date = formatDateString(seis_velocity_interval.Row_effective_date)
	seis_velocity_interval.Row_expiry_date = formatDateString(seis_velocity_interval.Row_expiry_date)






	rows, err := stmt.Exec(seis_velocity_interval.Interval_id, seis_velocity_interval.Active_ind, seis_velocity_interval.Base_depth, seis_velocity_interval.Base_depth_ouom, seis_velocity_interval.Base_strat_unit_id, seis_velocity_interval.Compute_method, seis_velocity_interval.Effective_date, seis_velocity_interval.Expiry_date, seis_velocity_interval.Last_update, seis_velocity_interval.Ppdm_guid, seis_velocity_interval.Remark, seis_velocity_interval.Reported_tvd, seis_velocity_interval.Reported_tvd_ouom, seis_velocity_interval.Seis_set_id, seis_velocity_interval.Seis_set_subtype, seis_velocity_interval.Seis_time_datum, seis_velocity_interval.Seis_time_datum_ouom, seis_velocity_interval.Source, seis_velocity_interval.Strat_name_set_id, seis_velocity_interval.Top_depth, seis_velocity_interval.Top_depth_ouom, seis_velocity_interval.Top_strat_unit_id, seis_velocity_interval.Uwi, seis_velocity_interval.Velocity_quality, seis_velocity_interval.Velocity_type, seis_velocity_interval.Velocity_value, seis_velocity_interval.Velocity_value_ouom, seis_velocity_interval.Row_changed_by, seis_velocity_interval.Row_changed_date, seis_velocity_interval.Row_created_by, seis_velocity_interval.Row_created_date, seis_velocity_interval.Row_effective_date, seis_velocity_interval.Row_expiry_date, seis_velocity_interval.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisVelocityInterval(c *fiber.Ctx) error {
	var seis_velocity_interval dto.Seis_velocity_interval

	setDefaults(&seis_velocity_interval)

	if err := json.Unmarshal(c.Body(), &seis_velocity_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_velocity_interval.Interval_id = id

    if seis_velocity_interval.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_velocity_interval set active_ind = :1, base_depth = :2, base_depth_ouom = :3, base_strat_unit_id = :4, compute_method = :5, effective_date = :6, expiry_date = :7, last_update = :8, ppdm_guid = :9, remark = :10, reported_tvd = :11, reported_tvd_ouom = :12, seis_set_id = :13, seis_set_subtype = :14, seis_time_datum = :15, seis_time_datum_ouom = :16, source = :17, strat_name_set_id = :18, top_depth = :19, top_depth_ouom = :20, top_strat_unit_id = :21, uwi = :22, velocity_quality = :23, velocity_type = :24, velocity_value = :25, velocity_value_ouom = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where interval_id = :34")
	if err != nil {
		return err
	}

	seis_velocity_interval.Row_changed_date = formatDate(seis_velocity_interval.Row_changed_date)
	seis_velocity_interval.Effective_date = formatDateString(seis_velocity_interval.Effective_date)
	seis_velocity_interval.Expiry_date = formatDateString(seis_velocity_interval.Expiry_date)
	seis_velocity_interval.Last_update = formatDateString(seis_velocity_interval.Last_update)
	seis_velocity_interval.Row_effective_date = formatDateString(seis_velocity_interval.Row_effective_date)
	seis_velocity_interval.Row_expiry_date = formatDateString(seis_velocity_interval.Row_expiry_date)






	rows, err := stmt.Exec(seis_velocity_interval.Active_ind, seis_velocity_interval.Base_depth, seis_velocity_interval.Base_depth_ouom, seis_velocity_interval.Base_strat_unit_id, seis_velocity_interval.Compute_method, seis_velocity_interval.Effective_date, seis_velocity_interval.Expiry_date, seis_velocity_interval.Last_update, seis_velocity_interval.Ppdm_guid, seis_velocity_interval.Remark, seis_velocity_interval.Reported_tvd, seis_velocity_interval.Reported_tvd_ouom, seis_velocity_interval.Seis_set_id, seis_velocity_interval.Seis_set_subtype, seis_velocity_interval.Seis_time_datum, seis_velocity_interval.Seis_time_datum_ouom, seis_velocity_interval.Source, seis_velocity_interval.Strat_name_set_id, seis_velocity_interval.Top_depth, seis_velocity_interval.Top_depth_ouom, seis_velocity_interval.Top_strat_unit_id, seis_velocity_interval.Uwi, seis_velocity_interval.Velocity_quality, seis_velocity_interval.Velocity_type, seis_velocity_interval.Velocity_value, seis_velocity_interval.Velocity_value_ouom, seis_velocity_interval.Row_changed_by, seis_velocity_interval.Row_changed_date, seis_velocity_interval.Row_created_by, seis_velocity_interval.Row_effective_date, seis_velocity_interval.Row_expiry_date, seis_velocity_interval.Row_quality, seis_velocity_interval.Interval_id)
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

func PatchSeisVelocityInterval(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_velocity_interval set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "last_update" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where interval_id = :id"

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

func DeleteSeisVelocityInterval(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_velocity_interval dto.Seis_velocity_interval
	seis_velocity_interval.Interval_id = id

	stmt, err := db.Prepare("delete from seis_velocity_interval where interval_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_velocity_interval.Interval_id)
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


