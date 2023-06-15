package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillBitJet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_bit_jet")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_bit_jet

	for rows.Next() {
		var well_drill_bit_jet dto.Well_drill_bit_jet
		if err := rows.Scan(&well_drill_bit_jet.Uwi, &well_drill_bit_jet.Bit_source, &well_drill_bit_jet.Bit_interval_obs_no, &well_drill_bit_jet.Jet_obs_no, &well_drill_bit_jet.Active_ind, &well_drill_bit_jet.Effective_date, &well_drill_bit_jet.Expiry_date, &well_drill_bit_jet.Jet_diameter, &well_drill_bit_jet.Jet_diameter_ouom, &well_drill_bit_jet.Jet_type, &well_drill_bit_jet.Ppdm_guid, &well_drill_bit_jet.Remark, &well_drill_bit_jet.Source, &well_drill_bit_jet.Row_changed_by, &well_drill_bit_jet.Row_changed_date, &well_drill_bit_jet.Row_created_by, &well_drill_bit_jet.Row_created_date, &well_drill_bit_jet.Row_effective_date, &well_drill_bit_jet.Row_expiry_date, &well_drill_bit_jet.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_bit_jet to result
		result = append(result, well_drill_bit_jet)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillBitJet(c *fiber.Ctx) error {
	var well_drill_bit_jet dto.Well_drill_bit_jet

	setDefaults(&well_drill_bit_jet)

	if err := json.Unmarshal(c.Body(), &well_drill_bit_jet); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_bit_jet values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	well_drill_bit_jet.Row_created_date = formatDate(well_drill_bit_jet.Row_created_date)
	well_drill_bit_jet.Row_changed_date = formatDate(well_drill_bit_jet.Row_changed_date)
	well_drill_bit_jet.Effective_date = formatDateString(well_drill_bit_jet.Effective_date)
	well_drill_bit_jet.Expiry_date = formatDateString(well_drill_bit_jet.Expiry_date)
	well_drill_bit_jet.Row_effective_date = formatDateString(well_drill_bit_jet.Row_effective_date)
	well_drill_bit_jet.Row_expiry_date = formatDateString(well_drill_bit_jet.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_bit_jet.Uwi, well_drill_bit_jet.Bit_source, well_drill_bit_jet.Bit_interval_obs_no, well_drill_bit_jet.Jet_obs_no, well_drill_bit_jet.Active_ind, well_drill_bit_jet.Effective_date, well_drill_bit_jet.Expiry_date, well_drill_bit_jet.Jet_diameter, well_drill_bit_jet.Jet_diameter_ouom, well_drill_bit_jet.Jet_type, well_drill_bit_jet.Ppdm_guid, well_drill_bit_jet.Remark, well_drill_bit_jet.Source, well_drill_bit_jet.Row_changed_by, well_drill_bit_jet.Row_changed_date, well_drill_bit_jet.Row_created_by, well_drill_bit_jet.Row_created_date, well_drill_bit_jet.Row_effective_date, well_drill_bit_jet.Row_expiry_date, well_drill_bit_jet.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillBitJet(c *fiber.Ctx) error {
	var well_drill_bit_jet dto.Well_drill_bit_jet

	setDefaults(&well_drill_bit_jet)

	if err := json.Unmarshal(c.Body(), &well_drill_bit_jet); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_bit_jet.Uwi = id

    if well_drill_bit_jet.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_bit_jet set bit_source = :1, bit_interval_obs_no = :2, jet_obs_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, jet_diameter = :7, jet_diameter_ouom = :8, jet_type = :9, ppdm_guid = :10, remark = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where uwi = :20")
	if err != nil {
		return err
	}

	well_drill_bit_jet.Row_changed_date = formatDate(well_drill_bit_jet.Row_changed_date)
	well_drill_bit_jet.Effective_date = formatDateString(well_drill_bit_jet.Effective_date)
	well_drill_bit_jet.Expiry_date = formatDateString(well_drill_bit_jet.Expiry_date)
	well_drill_bit_jet.Row_effective_date = formatDateString(well_drill_bit_jet.Row_effective_date)
	well_drill_bit_jet.Row_expiry_date = formatDateString(well_drill_bit_jet.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_bit_jet.Bit_source, well_drill_bit_jet.Bit_interval_obs_no, well_drill_bit_jet.Jet_obs_no, well_drill_bit_jet.Active_ind, well_drill_bit_jet.Effective_date, well_drill_bit_jet.Expiry_date, well_drill_bit_jet.Jet_diameter, well_drill_bit_jet.Jet_diameter_ouom, well_drill_bit_jet.Jet_type, well_drill_bit_jet.Ppdm_guid, well_drill_bit_jet.Remark, well_drill_bit_jet.Source, well_drill_bit_jet.Row_changed_by, well_drill_bit_jet.Row_changed_date, well_drill_bit_jet.Row_created_by, well_drill_bit_jet.Row_effective_date, well_drill_bit_jet.Row_expiry_date, well_drill_bit_jet.Row_quality, well_drill_bit_jet.Uwi)
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

func PatchWellDrillBitJet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_bit_jet set "
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

func DeleteWellDrillBitJet(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_bit_jet dto.Well_drill_bit_jet
	well_drill_bit_jet.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_bit_jet where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_bit_jet.Uwi)
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


