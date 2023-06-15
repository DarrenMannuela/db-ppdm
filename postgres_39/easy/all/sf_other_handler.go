package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfOther(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_other")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_other

	for rows.Next() {
		var sf_other dto.Sf_other
		if err := rows.Scan(&sf_other.Sf_subtype, &sf_other.Support_facility_id, &sf_other.Active_ind, &sf_other.Area_size, &sf_other.Area_size_ouom, &sf_other.Capacity, &sf_other.Capacity_ouom, &sf_other.Capacity_uom, &sf_other.Effective_date, &sf_other.Expiry_date, &sf_other.Height, &sf_other.Height_ouom, &sf_other.Length, &sf_other.Length_ouom, &sf_other.Ppdm_guid, &sf_other.Remark, &sf_other.Source, &sf_other.Width, &sf_other.Width_ouom, &sf_other.Row_changed_by, &sf_other.Row_changed_date, &sf_other.Row_created_by, &sf_other.Row_created_date, &sf_other.Row_effective_date, &sf_other.Row_expiry_date, &sf_other.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_other to result
		result = append(result, sf_other)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfOther(c *fiber.Ctx) error {
	var sf_other dto.Sf_other

	setDefaults(&sf_other)

	if err := json.Unmarshal(c.Body(), &sf_other); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_other values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	sf_other.Row_created_date = formatDate(sf_other.Row_created_date)
	sf_other.Row_changed_date = formatDate(sf_other.Row_changed_date)
	sf_other.Effective_date = formatDateString(sf_other.Effective_date)
	sf_other.Expiry_date = formatDateString(sf_other.Expiry_date)
	sf_other.Row_effective_date = formatDateString(sf_other.Row_effective_date)
	sf_other.Row_expiry_date = formatDateString(sf_other.Row_expiry_date)






	rows, err := stmt.Exec(sf_other.Sf_subtype, sf_other.Support_facility_id, sf_other.Active_ind, sf_other.Area_size, sf_other.Area_size_ouom, sf_other.Capacity, sf_other.Capacity_ouom, sf_other.Capacity_uom, sf_other.Effective_date, sf_other.Expiry_date, sf_other.Height, sf_other.Height_ouom, sf_other.Length, sf_other.Length_ouom, sf_other.Ppdm_guid, sf_other.Remark, sf_other.Source, sf_other.Width, sf_other.Width_ouom, sf_other.Row_changed_by, sf_other.Row_changed_date, sf_other.Row_created_by, sf_other.Row_created_date, sf_other.Row_effective_date, sf_other.Row_expiry_date, sf_other.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfOther(c *fiber.Ctx) error {
	var sf_other dto.Sf_other

	setDefaults(&sf_other)

	if err := json.Unmarshal(c.Body(), &sf_other); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_other.Sf_subtype = id

    if sf_other.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_other set support_facility_id = :1, active_ind = :2, area_size = :3, area_size_ouom = :4, capacity = :5, capacity_ouom = :6, capacity_uom = :7, effective_date = :8, expiry_date = :9, height = :10, height_ouom = :11, length = :12, length_ouom = :13, ppdm_guid = :14, remark = :15, source = :16, width = :17, width_ouom = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where sf_subtype = :26")
	if err != nil {
		return err
	}

	sf_other.Row_changed_date = formatDate(sf_other.Row_changed_date)
	sf_other.Effective_date = formatDateString(sf_other.Effective_date)
	sf_other.Expiry_date = formatDateString(sf_other.Expiry_date)
	sf_other.Row_effective_date = formatDateString(sf_other.Row_effective_date)
	sf_other.Row_expiry_date = formatDateString(sf_other.Row_expiry_date)






	rows, err := stmt.Exec(sf_other.Support_facility_id, sf_other.Active_ind, sf_other.Area_size, sf_other.Area_size_ouom, sf_other.Capacity, sf_other.Capacity_ouom, sf_other.Capacity_uom, sf_other.Effective_date, sf_other.Expiry_date, sf_other.Height, sf_other.Height_ouom, sf_other.Length, sf_other.Length_ouom, sf_other.Ppdm_guid, sf_other.Remark, sf_other.Source, sf_other.Width, sf_other.Width_ouom, sf_other.Row_changed_by, sf_other.Row_changed_date, sf_other.Row_created_by, sf_other.Row_effective_date, sf_other.Row_expiry_date, sf_other.Row_quality, sf_other.Sf_subtype)
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

func PatchSfOther(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_other set "
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

func DeleteSfOther(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_other dto.Sf_other
	sf_other.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_other where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_other.Sf_subtype)
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


