package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisLicenseType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_license_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_license_type

	for rows.Next() {
		var seis_license_type dto.Seis_license_type
		if err := rows.Scan(&seis_license_type.Granted_by_ba_id, &seis_license_type.Seis_license_type_id, &seis_license_type.Abbreviation, &seis_license_type.Active_ind, &seis_license_type.Effective_date, &seis_license_type.Expiry_date, &seis_license_type.Long_name, &seis_license_type.Ppdm_guid, &seis_license_type.Rate_schedule_id, &seis_license_type.Remark, &seis_license_type.Short_name, &seis_license_type.Source, &seis_license_type.Row_changed_by, &seis_license_type.Row_changed_date, &seis_license_type.Row_created_by, &seis_license_type.Row_created_date, &seis_license_type.Row_effective_date, &seis_license_type.Row_expiry_date, &seis_license_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_license_type to result
		result = append(result, seis_license_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisLicenseType(c *fiber.Ctx) error {
	var seis_license_type dto.Seis_license_type

	setDefaults(&seis_license_type)

	if err := json.Unmarshal(c.Body(), &seis_license_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_license_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	seis_license_type.Row_created_date = formatDate(seis_license_type.Row_created_date)
	seis_license_type.Row_changed_date = formatDate(seis_license_type.Row_changed_date)
	seis_license_type.Effective_date = formatDateString(seis_license_type.Effective_date)
	seis_license_type.Expiry_date = formatDateString(seis_license_type.Expiry_date)
	seis_license_type.Row_effective_date = formatDateString(seis_license_type.Row_effective_date)
	seis_license_type.Row_expiry_date = formatDateString(seis_license_type.Row_expiry_date)






	rows, err := stmt.Exec(seis_license_type.Granted_by_ba_id, seis_license_type.Seis_license_type_id, seis_license_type.Abbreviation, seis_license_type.Active_ind, seis_license_type.Effective_date, seis_license_type.Expiry_date, seis_license_type.Long_name, seis_license_type.Ppdm_guid, seis_license_type.Rate_schedule_id, seis_license_type.Remark, seis_license_type.Short_name, seis_license_type.Source, seis_license_type.Row_changed_by, seis_license_type.Row_changed_date, seis_license_type.Row_created_by, seis_license_type.Row_created_date, seis_license_type.Row_effective_date, seis_license_type.Row_expiry_date, seis_license_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisLicenseType(c *fiber.Ctx) error {
	var seis_license_type dto.Seis_license_type

	setDefaults(&seis_license_type)

	if err := json.Unmarshal(c.Body(), &seis_license_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_license_type.Granted_by_ba_id = id

    if seis_license_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_license_type set seis_license_type_id = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, rate_schedule_id = :8, remark = :9, short_name = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where granted_by_ba_id = :19")
	if err != nil {
		return err
	}

	seis_license_type.Row_changed_date = formatDate(seis_license_type.Row_changed_date)
	seis_license_type.Effective_date = formatDateString(seis_license_type.Effective_date)
	seis_license_type.Expiry_date = formatDateString(seis_license_type.Expiry_date)
	seis_license_type.Row_effective_date = formatDateString(seis_license_type.Row_effective_date)
	seis_license_type.Row_expiry_date = formatDateString(seis_license_type.Row_expiry_date)






	rows, err := stmt.Exec(seis_license_type.Seis_license_type_id, seis_license_type.Abbreviation, seis_license_type.Active_ind, seis_license_type.Effective_date, seis_license_type.Expiry_date, seis_license_type.Long_name, seis_license_type.Ppdm_guid, seis_license_type.Rate_schedule_id, seis_license_type.Remark, seis_license_type.Short_name, seis_license_type.Source, seis_license_type.Row_changed_by, seis_license_type.Row_changed_date, seis_license_type.Row_created_by, seis_license_type.Row_effective_date, seis_license_type.Row_expiry_date, seis_license_type.Row_quality, seis_license_type.Granted_by_ba_id)
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

func PatchSeisLicenseType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_license_type set "
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

func DeleteSeisLicenseType(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_license_type dto.Seis_license_type
	seis_license_type.Granted_by_ba_id = id

	stmt, err := db.Prepare("delete from seis_license_type where granted_by_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_license_type.Granted_by_ba_id)
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


