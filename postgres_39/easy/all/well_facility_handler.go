package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellFacility(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_facility")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_facility

	for rows.Next() {
		var well_facility dto.Well_facility
		if err := rows.Scan(&well_facility.Uwi, &well_facility.Facility_id, &well_facility.Facility_type, &well_facility.Active_obs_no, &well_facility.Active_ind, &well_facility.Effective_date, &well_facility.Expiry_date, &well_facility.Facility_use_type, &well_facility.Ppdm_guid, &well_facility.Remark, &well_facility.Single_source_ind, &well_facility.Source, &well_facility.String_id, &well_facility.String_source, &well_facility.Row_changed_by, &well_facility.Row_changed_date, &well_facility.Row_created_by, &well_facility.Row_created_date, &well_facility.Row_effective_date, &well_facility.Row_expiry_date, &well_facility.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_facility to result
		result = append(result, well_facility)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellFacility(c *fiber.Ctx) error {
	var well_facility dto.Well_facility

	setDefaults(&well_facility)

	if err := json.Unmarshal(c.Body(), &well_facility); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_facility values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	well_facility.Row_created_date = formatDate(well_facility.Row_created_date)
	well_facility.Row_changed_date = formatDate(well_facility.Row_changed_date)
	well_facility.Effective_date = formatDateString(well_facility.Effective_date)
	well_facility.Expiry_date = formatDateString(well_facility.Expiry_date)
	well_facility.Row_effective_date = formatDateString(well_facility.Row_effective_date)
	well_facility.Row_expiry_date = formatDateString(well_facility.Row_expiry_date)






	rows, err := stmt.Exec(well_facility.Uwi, well_facility.Facility_id, well_facility.Facility_type, well_facility.Active_obs_no, well_facility.Active_ind, well_facility.Effective_date, well_facility.Expiry_date, well_facility.Facility_use_type, well_facility.Ppdm_guid, well_facility.Remark, well_facility.Single_source_ind, well_facility.Source, well_facility.String_id, well_facility.String_source, well_facility.Row_changed_by, well_facility.Row_changed_date, well_facility.Row_created_by, well_facility.Row_created_date, well_facility.Row_effective_date, well_facility.Row_expiry_date, well_facility.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellFacility(c *fiber.Ctx) error {
	var well_facility dto.Well_facility

	setDefaults(&well_facility)

	if err := json.Unmarshal(c.Body(), &well_facility); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_facility.Uwi = id

    if well_facility.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_facility set facility_id = :1, facility_type = :2, active_obs_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, facility_use_type = :7, ppdm_guid = :8, remark = :9, single_source_ind = :10, source = :11, string_id = :12, string_source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where uwi = :21")
	if err != nil {
		return err
	}

	well_facility.Row_changed_date = formatDate(well_facility.Row_changed_date)
	well_facility.Effective_date = formatDateString(well_facility.Effective_date)
	well_facility.Expiry_date = formatDateString(well_facility.Expiry_date)
	well_facility.Row_effective_date = formatDateString(well_facility.Row_effective_date)
	well_facility.Row_expiry_date = formatDateString(well_facility.Row_expiry_date)






	rows, err := stmt.Exec(well_facility.Facility_id, well_facility.Facility_type, well_facility.Active_obs_no, well_facility.Active_ind, well_facility.Effective_date, well_facility.Expiry_date, well_facility.Facility_use_type, well_facility.Ppdm_guid, well_facility.Remark, well_facility.Single_source_ind, well_facility.Source, well_facility.String_id, well_facility.String_source, well_facility.Row_changed_by, well_facility.Row_changed_date, well_facility.Row_created_by, well_facility.Row_effective_date, well_facility.Row_expiry_date, well_facility.Row_quality, well_facility.Uwi)
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

func PatchWellFacility(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_facility set "
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

func DeleteWellFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_facility dto.Well_facility
	well_facility.Uwi = id

	stmt, err := db.Prepare("delete from well_facility where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_facility.Uwi)
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


