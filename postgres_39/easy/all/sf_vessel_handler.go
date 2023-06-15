package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfVessel(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_vessel")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_vessel

	for rows.Next() {
		var sf_vessel dto.Sf_vessel
		if err := rows.Scan(&sf_vessel.Sf_subtype, &sf_vessel.Vessel_id, &sf_vessel.Active_ind, &sf_vessel.Area_id, &sf_vessel.Area_type, &sf_vessel.Backup_antenna_location, &sf_vessel.Backup_antenna_type, &sf_vessel.Current_operator, &sf_vessel.Drill_hole_position, &sf_vessel.Effective_date, &sf_vessel.Expiry_date, &sf_vessel.Positioning_method, &sf_vessel.Ppdm_guid, &sf_vessel.Primary_antenna_location, &sf_vessel.Primary_antenna_type, &sf_vessel.Remark, &sf_vessel.Secondary_antennal_location, &sf_vessel.Secondary_antenna_type, &sf_vessel.Source, &sf_vessel.Vessel_beam, &sf_vessel.Vessel_displacement, &sf_vessel.Vessel_draft, &sf_vessel.Vessel_length, &sf_vessel.Vessel_name, &sf_vessel.Vessel_size, &sf_vessel.Vessel_type, &sf_vessel.Row_changed_by, &sf_vessel.Row_changed_date, &sf_vessel.Row_created_by, &sf_vessel.Row_created_date, &sf_vessel.Row_effective_date, &sf_vessel.Row_expiry_date, &sf_vessel.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_vessel to result
		result = append(result, sf_vessel)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfVessel(c *fiber.Ctx) error {
	var sf_vessel dto.Sf_vessel

	setDefaults(&sf_vessel)

	if err := json.Unmarshal(c.Body(), &sf_vessel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_vessel values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	sf_vessel.Row_created_date = formatDate(sf_vessel.Row_created_date)
	sf_vessel.Row_changed_date = formatDate(sf_vessel.Row_changed_date)
	sf_vessel.Effective_date = formatDateString(sf_vessel.Effective_date)
	sf_vessel.Expiry_date = formatDateString(sf_vessel.Expiry_date)
	sf_vessel.Row_effective_date = formatDateString(sf_vessel.Row_effective_date)
	sf_vessel.Row_expiry_date = formatDateString(sf_vessel.Row_expiry_date)






	rows, err := stmt.Exec(sf_vessel.Sf_subtype, sf_vessel.Vessel_id, sf_vessel.Active_ind, sf_vessel.Area_id, sf_vessel.Area_type, sf_vessel.Backup_antenna_location, sf_vessel.Backup_antenna_type, sf_vessel.Current_operator, sf_vessel.Drill_hole_position, sf_vessel.Effective_date, sf_vessel.Expiry_date, sf_vessel.Positioning_method, sf_vessel.Ppdm_guid, sf_vessel.Primary_antenna_location, sf_vessel.Primary_antenna_type, sf_vessel.Remark, sf_vessel.Secondary_antennal_location, sf_vessel.Secondary_antenna_type, sf_vessel.Source, sf_vessel.Vessel_beam, sf_vessel.Vessel_displacement, sf_vessel.Vessel_draft, sf_vessel.Vessel_length, sf_vessel.Vessel_name, sf_vessel.Vessel_size, sf_vessel.Vessel_type, sf_vessel.Row_changed_by, sf_vessel.Row_changed_date, sf_vessel.Row_created_by, sf_vessel.Row_created_date, sf_vessel.Row_effective_date, sf_vessel.Row_expiry_date, sf_vessel.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfVessel(c *fiber.Ctx) error {
	var sf_vessel dto.Sf_vessel

	setDefaults(&sf_vessel)

	if err := json.Unmarshal(c.Body(), &sf_vessel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_vessel.Sf_subtype = id

    if sf_vessel.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_vessel set vessel_id = :1, active_ind = :2, area_id = :3, area_type = :4, backup_antenna_location = :5, backup_antenna_type = :6, current_operator = :7, drill_hole_position = :8, effective_date = :9, expiry_date = :10, positioning_method = :11, ppdm_guid = :12, primary_antenna_location = :13, primary_antenna_type = :14, remark = :15, secondary_antennal_location = :16, secondary_antenna_type = :17, source = :18, vessel_beam = :19, vessel_displacement = :20, vessel_draft = :21, vessel_length = :22, vessel_name = :23, vessel_size = :24, vessel_type = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where sf_subtype = :33")
	if err != nil {
		return err
	}

	sf_vessel.Row_changed_date = formatDate(sf_vessel.Row_changed_date)
	sf_vessel.Effective_date = formatDateString(sf_vessel.Effective_date)
	sf_vessel.Expiry_date = formatDateString(sf_vessel.Expiry_date)
	sf_vessel.Row_effective_date = formatDateString(sf_vessel.Row_effective_date)
	sf_vessel.Row_expiry_date = formatDateString(sf_vessel.Row_expiry_date)






	rows, err := stmt.Exec(sf_vessel.Vessel_id, sf_vessel.Active_ind, sf_vessel.Area_id, sf_vessel.Area_type, sf_vessel.Backup_antenna_location, sf_vessel.Backup_antenna_type, sf_vessel.Current_operator, sf_vessel.Drill_hole_position, sf_vessel.Effective_date, sf_vessel.Expiry_date, sf_vessel.Positioning_method, sf_vessel.Ppdm_guid, sf_vessel.Primary_antenna_location, sf_vessel.Primary_antenna_type, sf_vessel.Remark, sf_vessel.Secondary_antennal_location, sf_vessel.Secondary_antenna_type, sf_vessel.Source, sf_vessel.Vessel_beam, sf_vessel.Vessel_displacement, sf_vessel.Vessel_draft, sf_vessel.Vessel_length, sf_vessel.Vessel_name, sf_vessel.Vessel_size, sf_vessel.Vessel_type, sf_vessel.Row_changed_by, sf_vessel.Row_changed_date, sf_vessel.Row_created_by, sf_vessel.Row_effective_date, sf_vessel.Row_expiry_date, sf_vessel.Row_quality, sf_vessel.Sf_subtype)
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

func PatchSfVessel(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_vessel set "
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
	query += " where sf_subtype = :id"

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

func DeleteSfVessel(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_vessel dto.Sf_vessel
	sf_vessel.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_vessel where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_vessel.Sf_subtype)
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


