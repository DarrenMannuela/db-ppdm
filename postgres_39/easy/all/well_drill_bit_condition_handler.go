package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillBitCondition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_bit_condition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_bit_condition

	for rows.Next() {
		var well_drill_bit_condition dto.Well_drill_bit_condition
		if err := rows.Scan(&well_drill_bit_condition.Uwi, &well_drill_bit_condition.Bit_source, &well_drill_bit_condition.Bit_interval_obs_no, &well_drill_bit_condition.Condition_obs_no, &well_drill_bit_condition.Active_ind, &well_drill_bit_condition.Bearing_condition, &well_drill_bit_condition.Drill_bit_condition, &well_drill_bit_condition.Effective_date, &well_drill_bit_condition.Expiry_date, &well_drill_bit_condition.Ppdm_guid, &well_drill_bit_condition.Remark, &well_drill_bit_condition.Source, &well_drill_bit_condition.Row_changed_by, &well_drill_bit_condition.Row_changed_date, &well_drill_bit_condition.Row_created_by, &well_drill_bit_condition.Row_created_date, &well_drill_bit_condition.Row_effective_date, &well_drill_bit_condition.Row_expiry_date, &well_drill_bit_condition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_bit_condition to result
		result = append(result, well_drill_bit_condition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillBitCondition(c *fiber.Ctx) error {
	var well_drill_bit_condition dto.Well_drill_bit_condition

	setDefaults(&well_drill_bit_condition)

	if err := json.Unmarshal(c.Body(), &well_drill_bit_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_bit_condition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	well_drill_bit_condition.Row_created_date = formatDate(well_drill_bit_condition.Row_created_date)
	well_drill_bit_condition.Row_changed_date = formatDate(well_drill_bit_condition.Row_changed_date)
	well_drill_bit_condition.Effective_date = formatDateString(well_drill_bit_condition.Effective_date)
	well_drill_bit_condition.Expiry_date = formatDateString(well_drill_bit_condition.Expiry_date)
	well_drill_bit_condition.Row_effective_date = formatDateString(well_drill_bit_condition.Row_effective_date)
	well_drill_bit_condition.Row_expiry_date = formatDateString(well_drill_bit_condition.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_bit_condition.Uwi, well_drill_bit_condition.Bit_source, well_drill_bit_condition.Bit_interval_obs_no, well_drill_bit_condition.Condition_obs_no, well_drill_bit_condition.Active_ind, well_drill_bit_condition.Bearing_condition, well_drill_bit_condition.Drill_bit_condition, well_drill_bit_condition.Effective_date, well_drill_bit_condition.Expiry_date, well_drill_bit_condition.Ppdm_guid, well_drill_bit_condition.Remark, well_drill_bit_condition.Source, well_drill_bit_condition.Row_changed_by, well_drill_bit_condition.Row_changed_date, well_drill_bit_condition.Row_created_by, well_drill_bit_condition.Row_created_date, well_drill_bit_condition.Row_effective_date, well_drill_bit_condition.Row_expiry_date, well_drill_bit_condition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillBitCondition(c *fiber.Ctx) error {
	var well_drill_bit_condition dto.Well_drill_bit_condition

	setDefaults(&well_drill_bit_condition)

	if err := json.Unmarshal(c.Body(), &well_drill_bit_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_bit_condition.Uwi = id

    if well_drill_bit_condition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_bit_condition set bit_source = :1, bit_interval_obs_no = :2, condition_obs_no = :3, active_ind = :4, bearing_condition = :5, drill_bit_condition = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where uwi = :19")
	if err != nil {
		return err
	}

	well_drill_bit_condition.Row_changed_date = formatDate(well_drill_bit_condition.Row_changed_date)
	well_drill_bit_condition.Effective_date = formatDateString(well_drill_bit_condition.Effective_date)
	well_drill_bit_condition.Expiry_date = formatDateString(well_drill_bit_condition.Expiry_date)
	well_drill_bit_condition.Row_effective_date = formatDateString(well_drill_bit_condition.Row_effective_date)
	well_drill_bit_condition.Row_expiry_date = formatDateString(well_drill_bit_condition.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_bit_condition.Bit_source, well_drill_bit_condition.Bit_interval_obs_no, well_drill_bit_condition.Condition_obs_no, well_drill_bit_condition.Active_ind, well_drill_bit_condition.Bearing_condition, well_drill_bit_condition.Drill_bit_condition, well_drill_bit_condition.Effective_date, well_drill_bit_condition.Expiry_date, well_drill_bit_condition.Ppdm_guid, well_drill_bit_condition.Remark, well_drill_bit_condition.Source, well_drill_bit_condition.Row_changed_by, well_drill_bit_condition.Row_changed_date, well_drill_bit_condition.Row_created_by, well_drill_bit_condition.Row_effective_date, well_drill_bit_condition.Row_expiry_date, well_drill_bit_condition.Row_quality, well_drill_bit_condition.Uwi)
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

func PatchWellDrillBitCondition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_bit_condition set "
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

func DeleteWellDrillBitCondition(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_bit_condition dto.Well_drill_bit_condition
	well_drill_bit_condition.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_bit_condition where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_bit_condition.Uwi)
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


