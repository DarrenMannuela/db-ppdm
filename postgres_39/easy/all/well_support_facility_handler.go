package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellSupportFacility(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_support_facility")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_support_facility

	for rows.Next() {
		var well_support_facility dto.Well_support_facility
		if err := rows.Scan(&well_support_facility.Uwi, &well_support_facility.Sf_subtype, &well_support_facility.Support_facility_id, &well_support_facility.Sf_obs_no, &well_support_facility.Active_ind, &well_support_facility.Depth_datum, &well_support_facility.Depth_datum_elev, &well_support_facility.Depth_datum_elev_ouom, &well_support_facility.Derrick_floor_elev, &well_support_facility.Derrick_floor_elev_ouom, &well_support_facility.Effective_date, &well_support_facility.Elev_ref_datum, &well_support_facility.Environment_type, &well_support_facility.Expiry_date, &well_support_facility.Kb_elev, &well_support_facility.Kb_elev_ouom, &well_support_facility.Ppdm_guid, &well_support_facility.Remark, &well_support_facility.Rig_on_site_date, &well_support_facility.Rig_release_date, &well_support_facility.Rotary_table_elev, &well_support_facility.Rotary_table_elev_ouom, &well_support_facility.Sf_use_type, &well_support_facility.Source, &well_support_facility.Subsea_elev_ref_type, &well_support_facility.Row_changed_by, &well_support_facility.Row_changed_date, &well_support_facility.Row_created_by, &well_support_facility.Row_created_date, &well_support_facility.Row_effective_date, &well_support_facility.Row_expiry_date, &well_support_facility.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_support_facility to result
		result = append(result, well_support_facility)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellSupportFacility(c *fiber.Ctx) error {
	var well_support_facility dto.Well_support_facility

	setDefaults(&well_support_facility)

	if err := json.Unmarshal(c.Body(), &well_support_facility); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_support_facility values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	well_support_facility.Row_created_date = formatDate(well_support_facility.Row_created_date)
	well_support_facility.Row_changed_date = formatDate(well_support_facility.Row_changed_date)
	well_support_facility.Effective_date = formatDateString(well_support_facility.Effective_date)
	well_support_facility.Expiry_date = formatDateString(well_support_facility.Expiry_date)
	well_support_facility.Rig_on_site_date = formatDateString(well_support_facility.Rig_on_site_date)
	well_support_facility.Rig_release_date = formatDateString(well_support_facility.Rig_release_date)
	well_support_facility.Row_effective_date = formatDateString(well_support_facility.Row_effective_date)
	well_support_facility.Row_expiry_date = formatDateString(well_support_facility.Row_expiry_date)






	rows, err := stmt.Exec(well_support_facility.Uwi, well_support_facility.Sf_subtype, well_support_facility.Support_facility_id, well_support_facility.Sf_obs_no, well_support_facility.Active_ind, well_support_facility.Depth_datum, well_support_facility.Depth_datum_elev, well_support_facility.Depth_datum_elev_ouom, well_support_facility.Derrick_floor_elev, well_support_facility.Derrick_floor_elev_ouom, well_support_facility.Effective_date, well_support_facility.Elev_ref_datum, well_support_facility.Environment_type, well_support_facility.Expiry_date, well_support_facility.Kb_elev, well_support_facility.Kb_elev_ouom, well_support_facility.Ppdm_guid, well_support_facility.Remark, well_support_facility.Rig_on_site_date, well_support_facility.Rig_release_date, well_support_facility.Rotary_table_elev, well_support_facility.Rotary_table_elev_ouom, well_support_facility.Sf_use_type, well_support_facility.Source, well_support_facility.Subsea_elev_ref_type, well_support_facility.Row_changed_by, well_support_facility.Row_changed_date, well_support_facility.Row_created_by, well_support_facility.Row_created_date, well_support_facility.Row_effective_date, well_support_facility.Row_expiry_date, well_support_facility.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellSupportFacility(c *fiber.Ctx) error {
	var well_support_facility dto.Well_support_facility

	setDefaults(&well_support_facility)

	if err := json.Unmarshal(c.Body(), &well_support_facility); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_support_facility.Uwi = id

    if well_support_facility.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_support_facility set sf_subtype = :1, support_facility_id = :2, sf_obs_no = :3, active_ind = :4, depth_datum = :5, depth_datum_elev = :6, depth_datum_elev_ouom = :7, derrick_floor_elev = :8, derrick_floor_elev_ouom = :9, effective_date = :10, elev_ref_datum = :11, environment_type = :12, expiry_date = :13, kb_elev = :14, kb_elev_ouom = :15, ppdm_guid = :16, remark = :17, rig_on_site_date = :18, rig_release_date = :19, rotary_table_elev = :20, rotary_table_elev_ouom = :21, sf_use_type = :22, source = :23, subsea_elev_ref_type = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where uwi = :32")
	if err != nil {
		return err
	}

	well_support_facility.Row_changed_date = formatDate(well_support_facility.Row_changed_date)
	well_support_facility.Effective_date = formatDateString(well_support_facility.Effective_date)
	well_support_facility.Expiry_date = formatDateString(well_support_facility.Expiry_date)
	well_support_facility.Rig_on_site_date = formatDateString(well_support_facility.Rig_on_site_date)
	well_support_facility.Rig_release_date = formatDateString(well_support_facility.Rig_release_date)
	well_support_facility.Row_effective_date = formatDateString(well_support_facility.Row_effective_date)
	well_support_facility.Row_expiry_date = formatDateString(well_support_facility.Row_expiry_date)






	rows, err := stmt.Exec(well_support_facility.Sf_subtype, well_support_facility.Support_facility_id, well_support_facility.Sf_obs_no, well_support_facility.Active_ind, well_support_facility.Depth_datum, well_support_facility.Depth_datum_elev, well_support_facility.Depth_datum_elev_ouom, well_support_facility.Derrick_floor_elev, well_support_facility.Derrick_floor_elev_ouom, well_support_facility.Effective_date, well_support_facility.Elev_ref_datum, well_support_facility.Environment_type, well_support_facility.Expiry_date, well_support_facility.Kb_elev, well_support_facility.Kb_elev_ouom, well_support_facility.Ppdm_guid, well_support_facility.Remark, well_support_facility.Rig_on_site_date, well_support_facility.Rig_release_date, well_support_facility.Rotary_table_elev, well_support_facility.Rotary_table_elev_ouom, well_support_facility.Sf_use_type, well_support_facility.Source, well_support_facility.Subsea_elev_ref_type, well_support_facility.Row_changed_by, well_support_facility.Row_changed_date, well_support_facility.Row_created_by, well_support_facility.Row_effective_date, well_support_facility.Row_expiry_date, well_support_facility.Row_quality, well_support_facility.Uwi)
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

func PatchWellSupportFacility(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_support_facility set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "rig_on_site_date" || key == "rig_release_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellSupportFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_support_facility dto.Well_support_facility
	well_support_facility.Uwi = id

	stmt, err := db.Prepare("delete from well_support_facility where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_support_facility.Uwi)
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


