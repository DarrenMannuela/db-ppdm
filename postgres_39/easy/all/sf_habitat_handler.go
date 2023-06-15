package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfHabitat(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_habitat")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_habitat

	for rows.Next() {
		var sf_habitat dto.Sf_habitat
		if err := rows.Scan(&sf_habitat.Sf_subtype, &sf_habitat.Support_facility_id, &sf_habitat.Active_ind, &sf_habitat.Area_size, &sf_habitat.Area_size_ouom, &sf_habitat.Bed_count, &sf_habitat.Capacity, &sf_habitat.Capacity_ouom, &sf_habitat.Capacity_uom, &sf_habitat.Effective_date, &sf_habitat.Expiry_date, &sf_habitat.Ppdm_guid, &sf_habitat.Remark, &sf_habitat.Source, &sf_habitat.Row_changed_by, &sf_habitat.Row_changed_date, &sf_habitat.Row_created_by, &sf_habitat.Row_created_date, &sf_habitat.Row_effective_date, &sf_habitat.Row_expiry_date, &sf_habitat.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_habitat to result
		result = append(result, sf_habitat)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfHabitat(c *fiber.Ctx) error {
	var sf_habitat dto.Sf_habitat

	setDefaults(&sf_habitat)

	if err := json.Unmarshal(c.Body(), &sf_habitat); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_habitat values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	sf_habitat.Row_created_date = formatDate(sf_habitat.Row_created_date)
	sf_habitat.Row_changed_date = formatDate(sf_habitat.Row_changed_date)
	sf_habitat.Effective_date = formatDateString(sf_habitat.Effective_date)
	sf_habitat.Expiry_date = formatDateString(sf_habitat.Expiry_date)
	sf_habitat.Row_effective_date = formatDateString(sf_habitat.Row_effective_date)
	sf_habitat.Row_expiry_date = formatDateString(sf_habitat.Row_expiry_date)






	rows, err := stmt.Exec(sf_habitat.Sf_subtype, sf_habitat.Support_facility_id, sf_habitat.Active_ind, sf_habitat.Area_size, sf_habitat.Area_size_ouom, sf_habitat.Bed_count, sf_habitat.Capacity, sf_habitat.Capacity_ouom, sf_habitat.Capacity_uom, sf_habitat.Effective_date, sf_habitat.Expiry_date, sf_habitat.Ppdm_guid, sf_habitat.Remark, sf_habitat.Source, sf_habitat.Row_changed_by, sf_habitat.Row_changed_date, sf_habitat.Row_created_by, sf_habitat.Row_created_date, sf_habitat.Row_effective_date, sf_habitat.Row_expiry_date, sf_habitat.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfHabitat(c *fiber.Ctx) error {
	var sf_habitat dto.Sf_habitat

	setDefaults(&sf_habitat)

	if err := json.Unmarshal(c.Body(), &sf_habitat); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_habitat.Sf_subtype = id

    if sf_habitat.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_habitat set support_facility_id = :1, active_ind = :2, area_size = :3, area_size_ouom = :4, bed_count = :5, capacity = :6, capacity_ouom = :7, capacity_uom = :8, effective_date = :9, expiry_date = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where sf_subtype = :21")
	if err != nil {
		return err
	}

	sf_habitat.Row_changed_date = formatDate(sf_habitat.Row_changed_date)
	sf_habitat.Effective_date = formatDateString(sf_habitat.Effective_date)
	sf_habitat.Expiry_date = formatDateString(sf_habitat.Expiry_date)
	sf_habitat.Row_effective_date = formatDateString(sf_habitat.Row_effective_date)
	sf_habitat.Row_expiry_date = formatDateString(sf_habitat.Row_expiry_date)






	rows, err := stmt.Exec(sf_habitat.Support_facility_id, sf_habitat.Active_ind, sf_habitat.Area_size, sf_habitat.Area_size_ouom, sf_habitat.Bed_count, sf_habitat.Capacity, sf_habitat.Capacity_ouom, sf_habitat.Capacity_uom, sf_habitat.Effective_date, sf_habitat.Expiry_date, sf_habitat.Ppdm_guid, sf_habitat.Remark, sf_habitat.Source, sf_habitat.Row_changed_by, sf_habitat.Row_changed_date, sf_habitat.Row_created_by, sf_habitat.Row_effective_date, sf_habitat.Row_expiry_date, sf_habitat.Row_quality, sf_habitat.Sf_subtype)
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

func PatchSfHabitat(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_habitat set "
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

func DeleteSfHabitat(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_habitat dto.Sf_habitat
	sf_habitat.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_habitat where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_habitat.Sf_subtype)
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


