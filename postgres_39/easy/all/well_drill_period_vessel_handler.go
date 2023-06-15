package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillPeriodVessel(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_period_vessel")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_period_vessel

	for rows.Next() {
		var well_drill_period_vessel dto.Well_drill_period_vessel
		if err := rows.Scan(&well_drill_period_vessel.Uwi, &well_drill_period_vessel.Period_obs_no, &well_drill_period_vessel.Sf_subtype, &well_drill_period_vessel.Vessel_id, &well_drill_period_vessel.Active_ind, &well_drill_period_vessel.Effective_date, &well_drill_period_vessel.Expiry_date, &well_drill_period_vessel.Heading, &well_drill_period_vessel.Heading_north_type, &well_drill_period_vessel.Passengers_off, &well_drill_period_vessel.Passengers_on, &well_drill_period_vessel.Ppdm_guid, &well_drill_period_vessel.Remark, &well_drill_period_vessel.Riser_angle, &well_drill_period_vessel.Riser_tension, &well_drill_period_vessel.Riser_tension_ouom, &well_drill_period_vessel.Riser_tension_uom, &well_drill_period_vessel.Source, &well_drill_period_vessel.Vessel_role, &well_drill_period_vessel.Row_changed_by, &well_drill_period_vessel.Row_changed_date, &well_drill_period_vessel.Row_created_by, &well_drill_period_vessel.Row_created_date, &well_drill_period_vessel.Row_effective_date, &well_drill_period_vessel.Row_expiry_date, &well_drill_period_vessel.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_period_vessel to result
		result = append(result, well_drill_period_vessel)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillPeriodVessel(c *fiber.Ctx) error {
	var well_drill_period_vessel dto.Well_drill_period_vessel

	setDefaults(&well_drill_period_vessel)

	if err := json.Unmarshal(c.Body(), &well_drill_period_vessel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_period_vessel values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	well_drill_period_vessel.Row_created_date = formatDate(well_drill_period_vessel.Row_created_date)
	well_drill_period_vessel.Row_changed_date = formatDate(well_drill_period_vessel.Row_changed_date)
	well_drill_period_vessel.Effective_date = formatDateString(well_drill_period_vessel.Effective_date)
	well_drill_period_vessel.Expiry_date = formatDateString(well_drill_period_vessel.Expiry_date)
	well_drill_period_vessel.Row_effective_date = formatDateString(well_drill_period_vessel.Row_effective_date)
	well_drill_period_vessel.Row_expiry_date = formatDateString(well_drill_period_vessel.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period_vessel.Uwi, well_drill_period_vessel.Period_obs_no, well_drill_period_vessel.Sf_subtype, well_drill_period_vessel.Vessel_id, well_drill_period_vessel.Active_ind, well_drill_period_vessel.Effective_date, well_drill_period_vessel.Expiry_date, well_drill_period_vessel.Heading, well_drill_period_vessel.Heading_north_type, well_drill_period_vessel.Passengers_off, well_drill_period_vessel.Passengers_on, well_drill_period_vessel.Ppdm_guid, well_drill_period_vessel.Remark, well_drill_period_vessel.Riser_angle, well_drill_period_vessel.Riser_tension, well_drill_period_vessel.Riser_tension_ouom, well_drill_period_vessel.Riser_tension_uom, well_drill_period_vessel.Source, well_drill_period_vessel.Vessel_role, well_drill_period_vessel.Row_changed_by, well_drill_period_vessel.Row_changed_date, well_drill_period_vessel.Row_created_by, well_drill_period_vessel.Row_created_date, well_drill_period_vessel.Row_effective_date, well_drill_period_vessel.Row_expiry_date, well_drill_period_vessel.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillPeriodVessel(c *fiber.Ctx) error {
	var well_drill_period_vessel dto.Well_drill_period_vessel

	setDefaults(&well_drill_period_vessel)

	if err := json.Unmarshal(c.Body(), &well_drill_period_vessel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_period_vessel.Uwi = id

    if well_drill_period_vessel.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_period_vessel set period_obs_no = :1, sf_subtype = :2, vessel_id = :3, active_ind = :4, effective_date = :5, expiry_date = :6, heading = :7, heading_north_type = :8, passengers_off = :9, passengers_on = :10, ppdm_guid = :11, remark = :12, riser_angle = :13, riser_tension = :14, riser_tension_ouom = :15, riser_tension_uom = :16, source = :17, vessel_role = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where uwi = :26")
	if err != nil {
		return err
	}

	well_drill_period_vessel.Row_changed_date = formatDate(well_drill_period_vessel.Row_changed_date)
	well_drill_period_vessel.Effective_date = formatDateString(well_drill_period_vessel.Effective_date)
	well_drill_period_vessel.Expiry_date = formatDateString(well_drill_period_vessel.Expiry_date)
	well_drill_period_vessel.Row_effective_date = formatDateString(well_drill_period_vessel.Row_effective_date)
	well_drill_period_vessel.Row_expiry_date = formatDateString(well_drill_period_vessel.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period_vessel.Period_obs_no, well_drill_period_vessel.Sf_subtype, well_drill_period_vessel.Vessel_id, well_drill_period_vessel.Active_ind, well_drill_period_vessel.Effective_date, well_drill_period_vessel.Expiry_date, well_drill_period_vessel.Heading, well_drill_period_vessel.Heading_north_type, well_drill_period_vessel.Passengers_off, well_drill_period_vessel.Passengers_on, well_drill_period_vessel.Ppdm_guid, well_drill_period_vessel.Remark, well_drill_period_vessel.Riser_angle, well_drill_period_vessel.Riser_tension, well_drill_period_vessel.Riser_tension_ouom, well_drill_period_vessel.Riser_tension_uom, well_drill_period_vessel.Source, well_drill_period_vessel.Vessel_role, well_drill_period_vessel.Row_changed_by, well_drill_period_vessel.Row_changed_date, well_drill_period_vessel.Row_created_by, well_drill_period_vessel.Row_effective_date, well_drill_period_vessel.Row_expiry_date, well_drill_period_vessel.Row_quality, well_drill_period_vessel.Uwi)
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

func PatchWellDrillPeriodVessel(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_period_vessel set "
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

func DeleteWellDrillPeriodVessel(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_period_vessel dto.Well_drill_period_vessel
	well_drill_period_vessel.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_period_vessel where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_period_vessel.Uwi)
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


