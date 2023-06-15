package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellPermitType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_permit_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_permit_type

	for rows.Next() {
		var well_permit_type dto.Well_permit_type
		if err := rows.Scan(&well_permit_type.Granted_by_ba_id, &well_permit_type.Permit_type, &well_permit_type.Abbreviation, &well_permit_type.Active_ind, &well_permit_type.Effective_date, &well_permit_type.Expiry_date, &well_permit_type.Long_name, &well_permit_type.Ppdm_guid, &well_permit_type.Rate_schedule_id, &well_permit_type.Remark, &well_permit_type.Short_name, &well_permit_type.Source, &well_permit_type.Row_changed_by, &well_permit_type.Row_changed_date, &well_permit_type.Row_created_by, &well_permit_type.Row_created_date, &well_permit_type.Row_effective_date, &well_permit_type.Row_expiry_date, &well_permit_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_permit_type to result
		result = append(result, well_permit_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellPermitType(c *fiber.Ctx) error {
	var well_permit_type dto.Well_permit_type

	setDefaults(&well_permit_type)

	if err := json.Unmarshal(c.Body(), &well_permit_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_permit_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	well_permit_type.Row_created_date = formatDate(well_permit_type.Row_created_date)
	well_permit_type.Row_changed_date = formatDate(well_permit_type.Row_changed_date)
	well_permit_type.Effective_date = formatDateString(well_permit_type.Effective_date)
	well_permit_type.Expiry_date = formatDateString(well_permit_type.Expiry_date)
	well_permit_type.Row_effective_date = formatDateString(well_permit_type.Row_effective_date)
	well_permit_type.Row_expiry_date = formatDateString(well_permit_type.Row_expiry_date)






	rows, err := stmt.Exec(well_permit_type.Granted_by_ba_id, well_permit_type.Permit_type, well_permit_type.Abbreviation, well_permit_type.Active_ind, well_permit_type.Effective_date, well_permit_type.Expiry_date, well_permit_type.Long_name, well_permit_type.Ppdm_guid, well_permit_type.Rate_schedule_id, well_permit_type.Remark, well_permit_type.Short_name, well_permit_type.Source, well_permit_type.Row_changed_by, well_permit_type.Row_changed_date, well_permit_type.Row_created_by, well_permit_type.Row_created_date, well_permit_type.Row_effective_date, well_permit_type.Row_expiry_date, well_permit_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellPermitType(c *fiber.Ctx) error {
	var well_permit_type dto.Well_permit_type

	setDefaults(&well_permit_type)

	if err := json.Unmarshal(c.Body(), &well_permit_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_permit_type.Granted_by_ba_id = id

    if well_permit_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_permit_type set permit_type = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, rate_schedule_id = :8, remark = :9, short_name = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where granted_by_ba_id = :19")
	if err != nil {
		return err
	}

	well_permit_type.Row_changed_date = formatDate(well_permit_type.Row_changed_date)
	well_permit_type.Effective_date = formatDateString(well_permit_type.Effective_date)
	well_permit_type.Expiry_date = formatDateString(well_permit_type.Expiry_date)
	well_permit_type.Row_effective_date = formatDateString(well_permit_type.Row_effective_date)
	well_permit_type.Row_expiry_date = formatDateString(well_permit_type.Row_expiry_date)






	rows, err := stmt.Exec(well_permit_type.Permit_type, well_permit_type.Abbreviation, well_permit_type.Active_ind, well_permit_type.Effective_date, well_permit_type.Expiry_date, well_permit_type.Long_name, well_permit_type.Ppdm_guid, well_permit_type.Rate_schedule_id, well_permit_type.Remark, well_permit_type.Short_name, well_permit_type.Source, well_permit_type.Row_changed_by, well_permit_type.Row_changed_date, well_permit_type.Row_created_by, well_permit_type.Row_effective_date, well_permit_type.Row_expiry_date, well_permit_type.Row_quality, well_permit_type.Granted_by_ba_id)
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

func PatchWellPermitType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_permit_type set "
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
	query += " where granted_by_ba_id = :id"

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

func DeleteWellPermitType(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_permit_type dto.Well_permit_type
	well_permit_type.Granted_by_ba_id = id

	stmt, err := db.Prepare("delete from well_permit_type where granted_by_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_permit_type.Granted_by_ba_id)
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


