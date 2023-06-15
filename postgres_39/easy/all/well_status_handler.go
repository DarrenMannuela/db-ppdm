package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_status

	for rows.Next() {
		var well_status dto.Well_status
		if err := rows.Scan(&well_status.Uwi, &well_status.Source, &well_status.Status_id, &well_status.Active_ind, &well_status.Effective_date, &well_status.End_time, &well_status.Expiry_date, &well_status.Percent_capability, &well_status.Ppdm_guid, &well_status.Remark, &well_status.Start_time, &well_status.Status, &well_status.Status_date, &well_status.Status_depth, &well_status.Status_depth_ouom, &well_status.Status_qualifier, &well_status.Status_qualifier_value, &well_status.Status_type, &well_status.Timezone, &well_status.Row_changed_by, &well_status.Row_changed_date, &well_status.Row_created_by, &well_status.Row_created_date, &well_status.Row_effective_date, &well_status.Row_expiry_date, &well_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_status to result
		result = append(result, well_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellStatus(c *fiber.Ctx) error {
	var well_status dto.Well_status

	setDefaults(&well_status)

	if err := json.Unmarshal(c.Body(), &well_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	well_status.Row_created_date = formatDate(well_status.Row_created_date)
	well_status.Row_changed_date = formatDate(well_status.Row_changed_date)
	well_status.Effective_date = formatDateString(well_status.Effective_date)
	well_status.End_time = formatDateString(well_status.End_time)
	well_status.Expiry_date = formatDateString(well_status.Expiry_date)
	well_status.Start_time = formatDateString(well_status.Start_time)
	well_status.Status_date = formatDateString(well_status.Status_date)
	well_status.Row_effective_date = formatDateString(well_status.Row_effective_date)
	well_status.Row_expiry_date = formatDateString(well_status.Row_expiry_date)






	rows, err := stmt.Exec(well_status.Uwi, well_status.Source, well_status.Status_id, well_status.Active_ind, well_status.Effective_date, well_status.End_time, well_status.Expiry_date, well_status.Percent_capability, well_status.Ppdm_guid, well_status.Remark, well_status.Start_time, well_status.Status, well_status.Status_date, well_status.Status_depth, well_status.Status_depth_ouom, well_status.Status_qualifier, well_status.Status_qualifier_value, well_status.Status_type, well_status.Timezone, well_status.Row_changed_by, well_status.Row_changed_date, well_status.Row_created_by, well_status.Row_created_date, well_status.Row_effective_date, well_status.Row_expiry_date, well_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellStatus(c *fiber.Ctx) error {
	var well_status dto.Well_status

	setDefaults(&well_status)

	if err := json.Unmarshal(c.Body(), &well_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_status.Uwi = id

    if well_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_status set source = :1, status_id = :2, active_ind = :3, effective_date = :4, end_time = :5, expiry_date = :6, percent_capability = :7, ppdm_guid = :8, remark = :9, start_time = :10, status = :11, status_date = :12, status_depth = :13, status_depth_ouom = :14, status_qualifier = :15, status_qualifier_value = :16, status_type = :17, timezone = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where uwi = :26")
	if err != nil {
		return err
	}

	well_status.Row_changed_date = formatDate(well_status.Row_changed_date)
	well_status.Effective_date = formatDateString(well_status.Effective_date)
	well_status.End_time = formatDateString(well_status.End_time)
	well_status.Expiry_date = formatDateString(well_status.Expiry_date)
	well_status.Start_time = formatDateString(well_status.Start_time)
	well_status.Status_date = formatDateString(well_status.Status_date)
	well_status.Row_effective_date = formatDateString(well_status.Row_effective_date)
	well_status.Row_expiry_date = formatDateString(well_status.Row_expiry_date)






	rows, err := stmt.Exec(well_status.Source, well_status.Status_id, well_status.Active_ind, well_status.Effective_date, well_status.End_time, well_status.Expiry_date, well_status.Percent_capability, well_status.Ppdm_guid, well_status.Remark, well_status.Start_time, well_status.Status, well_status.Status_date, well_status.Status_depth, well_status.Status_depth_ouom, well_status.Status_qualifier, well_status.Status_qualifier_value, well_status.Status_type, well_status.Timezone, well_status.Row_changed_by, well_status.Row_changed_date, well_status.Row_created_by, well_status.Row_effective_date, well_status.Row_expiry_date, well_status.Row_quality, well_status.Uwi)
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

func PatchWellStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_status set "
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
		} else if key == "effective_date" || key == "end_time" || key == "expiry_date" || key == "start_time" || key == "status_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_status dto.Well_status
	well_status.Uwi = id

	stmt, err := db.Prepare("delete from well_status where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_status.Uwi)
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


