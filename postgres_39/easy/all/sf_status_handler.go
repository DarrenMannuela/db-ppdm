package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_status

	for rows.Next() {
		var sf_status dto.Sf_status
		if err := rows.Scan(&sf_status.Sf_subtype, &sf_status.Support_facility_id, &sf_status.Status_id, &sf_status.Active_ind, &sf_status.Effective_date, &sf_status.Expiry_date, &sf_status.Ppdm_guid, &sf_status.Remark, &sf_status.Source, &sf_status.Status, &sf_status.Status_type, &sf_status.Row_changed_by, &sf_status.Row_changed_date, &sf_status.Row_created_by, &sf_status.Row_created_date, &sf_status.Row_effective_date, &sf_status.Row_expiry_date, &sf_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_status to result
		result = append(result, sf_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfStatus(c *fiber.Ctx) error {
	var sf_status dto.Sf_status

	setDefaults(&sf_status)

	if err := json.Unmarshal(c.Body(), &sf_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	sf_status.Row_created_date = formatDate(sf_status.Row_created_date)
	sf_status.Row_changed_date = formatDate(sf_status.Row_changed_date)
	sf_status.Effective_date = formatDateString(sf_status.Effective_date)
	sf_status.Expiry_date = formatDateString(sf_status.Expiry_date)
	sf_status.Row_effective_date = formatDateString(sf_status.Row_effective_date)
	sf_status.Row_expiry_date = formatDateString(sf_status.Row_expiry_date)






	rows, err := stmt.Exec(sf_status.Sf_subtype, sf_status.Support_facility_id, sf_status.Status_id, sf_status.Active_ind, sf_status.Effective_date, sf_status.Expiry_date, sf_status.Ppdm_guid, sf_status.Remark, sf_status.Source, sf_status.Status, sf_status.Status_type, sf_status.Row_changed_by, sf_status.Row_changed_date, sf_status.Row_created_by, sf_status.Row_created_date, sf_status.Row_effective_date, sf_status.Row_expiry_date, sf_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfStatus(c *fiber.Ctx) error {
	var sf_status dto.Sf_status

	setDefaults(&sf_status)

	if err := json.Unmarshal(c.Body(), &sf_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_status.Sf_subtype = id

    if sf_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_status set support_facility_id = :1, status_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, status = :9, status_type = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where sf_subtype = :18")
	if err != nil {
		return err
	}

	sf_status.Row_changed_date = formatDate(sf_status.Row_changed_date)
	sf_status.Effective_date = formatDateString(sf_status.Effective_date)
	sf_status.Expiry_date = formatDateString(sf_status.Expiry_date)
	sf_status.Row_effective_date = formatDateString(sf_status.Row_effective_date)
	sf_status.Row_expiry_date = formatDateString(sf_status.Row_expiry_date)






	rows, err := stmt.Exec(sf_status.Support_facility_id, sf_status.Status_id, sf_status.Active_ind, sf_status.Effective_date, sf_status.Expiry_date, sf_status.Ppdm_guid, sf_status.Remark, sf_status.Source, sf_status.Status, sf_status.Status_type, sf_status.Row_changed_by, sf_status.Row_changed_date, sf_status.Row_created_by, sf_status.Row_effective_date, sf_status.Row_expiry_date, sf_status.Row_quality, sf_status.Sf_subtype)
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

func PatchSfStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_status set "
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

func DeleteSfStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_status dto.Sf_status
	sf_status.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_status where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_status.Sf_subtype)
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


