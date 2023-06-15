package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillPeriodServ(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_period_serv")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_period_serv

	for rows.Next() {
		var well_drill_period_serv dto.Well_drill_period_serv
		if err := rows.Scan(&well_drill_period_serv.Uwi, &well_drill_period_serv.Period_obs_no, &well_drill_period_serv.Provided_by, &well_drill_period_serv.Service_seq_no, &well_drill_period_serv.Segment_obs_no, &well_drill_period_serv.Active_ind, &well_drill_period_serv.Description, &well_drill_period_serv.Effective_date, &well_drill_period_serv.End_time, &well_drill_period_serv.End_timezone, &well_drill_period_serv.Expiry_date, &well_drill_period_serv.Metric_code, &well_drill_period_serv.Metric_max_value, &well_drill_period_serv.Metric_min_value, &well_drill_period_serv.Metric_value, &well_drill_period_serv.Metric_value_ouom, &well_drill_period_serv.Metric_value_uom, &well_drill_period_serv.Ppdm_guid, &well_drill_period_serv.Remark, &well_drill_period_serv.Service_desc, &well_drill_period_serv.Service_metric, &well_drill_period_serv.Source, &well_drill_period_serv.Start_time, &well_drill_period_serv.Start_timezone, &well_drill_period_serv.Row_changed_by, &well_drill_period_serv.Row_changed_date, &well_drill_period_serv.Row_created_by, &well_drill_period_serv.Row_created_date, &well_drill_period_serv.Row_effective_date, &well_drill_period_serv.Row_expiry_date, &well_drill_period_serv.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_period_serv to result
		result = append(result, well_drill_period_serv)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillPeriodServ(c *fiber.Ctx) error {
	var well_drill_period_serv dto.Well_drill_period_serv

	setDefaults(&well_drill_period_serv)

	if err := json.Unmarshal(c.Body(), &well_drill_period_serv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_period_serv values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	well_drill_period_serv.Row_created_date = formatDate(well_drill_period_serv.Row_created_date)
	well_drill_period_serv.Row_changed_date = formatDate(well_drill_period_serv.Row_changed_date)
	well_drill_period_serv.Effective_date = formatDateString(well_drill_period_serv.Effective_date)
	well_drill_period_serv.End_time = formatDateString(well_drill_period_serv.End_time)
	well_drill_period_serv.Expiry_date = formatDateString(well_drill_period_serv.Expiry_date)
	well_drill_period_serv.Start_time = formatDateString(well_drill_period_serv.Start_time)
	well_drill_period_serv.Row_effective_date = formatDateString(well_drill_period_serv.Row_effective_date)
	well_drill_period_serv.Row_expiry_date = formatDateString(well_drill_period_serv.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period_serv.Uwi, well_drill_period_serv.Period_obs_no, well_drill_period_serv.Provided_by, well_drill_period_serv.Service_seq_no, well_drill_period_serv.Segment_obs_no, well_drill_period_serv.Active_ind, well_drill_period_serv.Description, well_drill_period_serv.Effective_date, well_drill_period_serv.End_time, well_drill_period_serv.End_timezone, well_drill_period_serv.Expiry_date, well_drill_period_serv.Metric_code, well_drill_period_serv.Metric_max_value, well_drill_period_serv.Metric_min_value, well_drill_period_serv.Metric_value, well_drill_period_serv.Metric_value_ouom, well_drill_period_serv.Metric_value_uom, well_drill_period_serv.Ppdm_guid, well_drill_period_serv.Remark, well_drill_period_serv.Service_desc, well_drill_period_serv.Service_metric, well_drill_period_serv.Source, well_drill_period_serv.Start_time, well_drill_period_serv.Start_timezone, well_drill_period_serv.Row_changed_by, well_drill_period_serv.Row_changed_date, well_drill_period_serv.Row_created_by, well_drill_period_serv.Row_created_date, well_drill_period_serv.Row_effective_date, well_drill_period_serv.Row_expiry_date, well_drill_period_serv.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillPeriodServ(c *fiber.Ctx) error {
	var well_drill_period_serv dto.Well_drill_period_serv

	setDefaults(&well_drill_period_serv)

	if err := json.Unmarshal(c.Body(), &well_drill_period_serv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_period_serv.Uwi = id

    if well_drill_period_serv.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_period_serv set period_obs_no = :1, provided_by = :2, service_seq_no = :3, segment_obs_no = :4, active_ind = :5, description = :6, effective_date = :7, end_time = :8, end_timezone = :9, expiry_date = :10, metric_code = :11, metric_max_value = :12, metric_min_value = :13, metric_value = :14, metric_value_ouom = :15, metric_value_uom = :16, ppdm_guid = :17, remark = :18, service_desc = :19, service_metric = :20, source = :21, start_time = :22, start_timezone = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where uwi = :31")
	if err != nil {
		return err
	}

	well_drill_period_serv.Row_changed_date = formatDate(well_drill_period_serv.Row_changed_date)
	well_drill_period_serv.Effective_date = formatDateString(well_drill_period_serv.Effective_date)
	well_drill_period_serv.End_time = formatDateString(well_drill_period_serv.End_time)
	well_drill_period_serv.Expiry_date = formatDateString(well_drill_period_serv.Expiry_date)
	well_drill_period_serv.Start_time = formatDateString(well_drill_period_serv.Start_time)
	well_drill_period_serv.Row_effective_date = formatDateString(well_drill_period_serv.Row_effective_date)
	well_drill_period_serv.Row_expiry_date = formatDateString(well_drill_period_serv.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period_serv.Period_obs_no, well_drill_period_serv.Provided_by, well_drill_period_serv.Service_seq_no, well_drill_period_serv.Segment_obs_no, well_drill_period_serv.Active_ind, well_drill_period_serv.Description, well_drill_period_serv.Effective_date, well_drill_period_serv.End_time, well_drill_period_serv.End_timezone, well_drill_period_serv.Expiry_date, well_drill_period_serv.Metric_code, well_drill_period_serv.Metric_max_value, well_drill_period_serv.Metric_min_value, well_drill_period_serv.Metric_value, well_drill_period_serv.Metric_value_ouom, well_drill_period_serv.Metric_value_uom, well_drill_period_serv.Ppdm_guid, well_drill_period_serv.Remark, well_drill_period_serv.Service_desc, well_drill_period_serv.Service_metric, well_drill_period_serv.Source, well_drill_period_serv.Start_time, well_drill_period_serv.Start_timezone, well_drill_period_serv.Row_changed_by, well_drill_period_serv.Row_changed_date, well_drill_period_serv.Row_created_by, well_drill_period_serv.Row_effective_date, well_drill_period_serv.Row_expiry_date, well_drill_period_serv.Row_quality, well_drill_period_serv.Uwi)
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

func PatchWellDrillPeriodServ(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_period_serv set "
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
		} else if key == "effective_date" || key == "end_time" || key == "expiry_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteWellDrillPeriodServ(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_period_serv dto.Well_drill_period_serv
	well_drill_period_serv.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_period_serv where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_period_serv.Uwi)
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


