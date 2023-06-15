package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillAddInv(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_add_inv")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_add_inv

	for rows.Next() {
		var well_drill_add_inv dto.Well_drill_add_inv
		if err := rows.Scan(&well_drill_add_inv.Uwi, &well_drill_add_inv.Drill_period_obs_no, &well_drill_add_inv.Additive_id, &well_drill_add_inv.Active_ind, &well_drill_add_inv.Effective_date, &well_drill_add_inv.Expiry_date, &well_drill_add_inv.Ppdm_guid, &well_drill_add_inv.Quantity_count_opening, &well_drill_add_inv.Quantity_count_ordered, &well_drill_add_inv.Quantity_count_remaining, &well_drill_add_inv.Quantity_count_used, &well_drill_add_inv.Quantity_value_opening, &well_drill_add_inv.Quantity_value_ordered, &well_drill_add_inv.Quantity_value_ouom, &well_drill_add_inv.Quantity_value_remaining, &well_drill_add_inv.Quantity_value_uom, &well_drill_add_inv.Quantity_value_used, &well_drill_add_inv.Remark, &well_drill_add_inv.Report_time_type, &well_drill_add_inv.Sf_subtype, &well_drill_add_inv.Source, &well_drill_add_inv.Support_facility_id, &well_drill_add_inv.Row_changed_by, &well_drill_add_inv.Row_changed_date, &well_drill_add_inv.Row_created_by, &well_drill_add_inv.Row_created_date, &well_drill_add_inv.Row_effective_date, &well_drill_add_inv.Row_expiry_date, &well_drill_add_inv.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_add_inv to result
		result = append(result, well_drill_add_inv)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillAddInv(c *fiber.Ctx) error {
	var well_drill_add_inv dto.Well_drill_add_inv

	setDefaults(&well_drill_add_inv)

	if err := json.Unmarshal(c.Body(), &well_drill_add_inv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_add_inv values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	well_drill_add_inv.Row_created_date = formatDate(well_drill_add_inv.Row_created_date)
	well_drill_add_inv.Row_changed_date = formatDate(well_drill_add_inv.Row_changed_date)
	well_drill_add_inv.Effective_date = formatDateString(well_drill_add_inv.Effective_date)
	well_drill_add_inv.Expiry_date = formatDateString(well_drill_add_inv.Expiry_date)
	well_drill_add_inv.Row_effective_date = formatDateString(well_drill_add_inv.Row_effective_date)
	well_drill_add_inv.Row_expiry_date = formatDateString(well_drill_add_inv.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_add_inv.Uwi, well_drill_add_inv.Drill_period_obs_no, well_drill_add_inv.Additive_id, well_drill_add_inv.Active_ind, well_drill_add_inv.Effective_date, well_drill_add_inv.Expiry_date, well_drill_add_inv.Ppdm_guid, well_drill_add_inv.Quantity_count_opening, well_drill_add_inv.Quantity_count_ordered, well_drill_add_inv.Quantity_count_remaining, well_drill_add_inv.Quantity_count_used, well_drill_add_inv.Quantity_value_opening, well_drill_add_inv.Quantity_value_ordered, well_drill_add_inv.Quantity_value_ouom, well_drill_add_inv.Quantity_value_remaining, well_drill_add_inv.Quantity_value_uom, well_drill_add_inv.Quantity_value_used, well_drill_add_inv.Remark, well_drill_add_inv.Report_time_type, well_drill_add_inv.Sf_subtype, well_drill_add_inv.Source, well_drill_add_inv.Support_facility_id, well_drill_add_inv.Row_changed_by, well_drill_add_inv.Row_changed_date, well_drill_add_inv.Row_created_by, well_drill_add_inv.Row_created_date, well_drill_add_inv.Row_effective_date, well_drill_add_inv.Row_expiry_date, well_drill_add_inv.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillAddInv(c *fiber.Ctx) error {
	var well_drill_add_inv dto.Well_drill_add_inv

	setDefaults(&well_drill_add_inv)

	if err := json.Unmarshal(c.Body(), &well_drill_add_inv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_add_inv.Uwi = id

    if well_drill_add_inv.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_add_inv set drill_period_obs_no = :1, additive_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, quantity_count_opening = :7, quantity_count_ordered = :8, quantity_count_remaining = :9, quantity_count_used = :10, quantity_value_opening = :11, quantity_value_ordered = :12, quantity_value_ouom = :13, quantity_value_remaining = :14, quantity_value_uom = :15, quantity_value_used = :16, remark = :17, report_time_type = :18, sf_subtype = :19, source = :20, support_facility_id = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where uwi = :29")
	if err != nil {
		return err
	}

	well_drill_add_inv.Row_changed_date = formatDate(well_drill_add_inv.Row_changed_date)
	well_drill_add_inv.Effective_date = formatDateString(well_drill_add_inv.Effective_date)
	well_drill_add_inv.Expiry_date = formatDateString(well_drill_add_inv.Expiry_date)
	well_drill_add_inv.Row_effective_date = formatDateString(well_drill_add_inv.Row_effective_date)
	well_drill_add_inv.Row_expiry_date = formatDateString(well_drill_add_inv.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_add_inv.Drill_period_obs_no, well_drill_add_inv.Additive_id, well_drill_add_inv.Active_ind, well_drill_add_inv.Effective_date, well_drill_add_inv.Expiry_date, well_drill_add_inv.Ppdm_guid, well_drill_add_inv.Quantity_count_opening, well_drill_add_inv.Quantity_count_ordered, well_drill_add_inv.Quantity_count_remaining, well_drill_add_inv.Quantity_count_used, well_drill_add_inv.Quantity_value_opening, well_drill_add_inv.Quantity_value_ordered, well_drill_add_inv.Quantity_value_ouom, well_drill_add_inv.Quantity_value_remaining, well_drill_add_inv.Quantity_value_uom, well_drill_add_inv.Quantity_value_used, well_drill_add_inv.Remark, well_drill_add_inv.Report_time_type, well_drill_add_inv.Sf_subtype, well_drill_add_inv.Source, well_drill_add_inv.Support_facility_id, well_drill_add_inv.Row_changed_by, well_drill_add_inv.Row_changed_date, well_drill_add_inv.Row_created_by, well_drill_add_inv.Row_effective_date, well_drill_add_inv.Row_expiry_date, well_drill_add_inv.Row_quality, well_drill_add_inv.Uwi)
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

func PatchWellDrillAddInv(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_add_inv set "
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

func DeleteWellDrillAddInv(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_add_inv dto.Well_drill_add_inv
	well_drill_add_inv.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_add_inv where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_add_inv.Uwi)
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


