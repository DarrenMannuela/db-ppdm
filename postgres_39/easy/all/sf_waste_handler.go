package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfWaste(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_waste")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_waste

	for rows.Next() {
		var sf_waste dto.Sf_waste
		if err := rows.Scan(&sf_waste.Sf_subtype, &sf_waste.Waste_facility_id, &sf_waste.Active_ind, &sf_waste.Capacity, &sf_waste.Capacity_ouom, &sf_waste.Current_owner, &sf_waste.Effective_date, &sf_waste.Expiry_date, &sf_waste.License_expiry_date, &sf_waste.License_num, &sf_waste.License_register_ba_id, &sf_waste.Ppdm_guid, &sf_waste.Remark, &sf_waste.Source, &sf_waste.Waste_facility_name, &sf_waste.Waste_facility_type, &sf_waste.Row_changed_by, &sf_waste.Row_changed_date, &sf_waste.Row_created_by, &sf_waste.Row_created_date, &sf_waste.Row_effective_date, &sf_waste.Row_expiry_date, &sf_waste.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_waste to result
		result = append(result, sf_waste)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfWaste(c *fiber.Ctx) error {
	var sf_waste dto.Sf_waste

	setDefaults(&sf_waste)

	if err := json.Unmarshal(c.Body(), &sf_waste); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_waste values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	sf_waste.Row_created_date = formatDate(sf_waste.Row_created_date)
	sf_waste.Row_changed_date = formatDate(sf_waste.Row_changed_date)
	sf_waste.Effective_date = formatDateString(sf_waste.Effective_date)
	sf_waste.Expiry_date = formatDateString(sf_waste.Expiry_date)
	sf_waste.License_expiry_date = formatDateString(sf_waste.License_expiry_date)
	sf_waste.Row_effective_date = formatDateString(sf_waste.Row_effective_date)
	sf_waste.Row_expiry_date = formatDateString(sf_waste.Row_expiry_date)






	rows, err := stmt.Exec(sf_waste.Sf_subtype, sf_waste.Waste_facility_id, sf_waste.Active_ind, sf_waste.Capacity, sf_waste.Capacity_ouom, sf_waste.Current_owner, sf_waste.Effective_date, sf_waste.Expiry_date, sf_waste.License_expiry_date, sf_waste.License_num, sf_waste.License_register_ba_id, sf_waste.Ppdm_guid, sf_waste.Remark, sf_waste.Source, sf_waste.Waste_facility_name, sf_waste.Waste_facility_type, sf_waste.Row_changed_by, sf_waste.Row_changed_date, sf_waste.Row_created_by, sf_waste.Row_created_date, sf_waste.Row_effective_date, sf_waste.Row_expiry_date, sf_waste.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfWaste(c *fiber.Ctx) error {
	var sf_waste dto.Sf_waste

	setDefaults(&sf_waste)

	if err := json.Unmarshal(c.Body(), &sf_waste); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_waste.Sf_subtype = id

    if sf_waste.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_waste set waste_facility_id = :1, active_ind = :2, capacity = :3, capacity_ouom = :4, current_owner = :5, effective_date = :6, expiry_date = :7, license_expiry_date = :8, license_num = :9, license_register_ba_id = :10, ppdm_guid = :11, remark = :12, source = :13, waste_facility_name = :14, waste_facility_type = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where sf_subtype = :23")
	if err != nil {
		return err
	}

	sf_waste.Row_changed_date = formatDate(sf_waste.Row_changed_date)
	sf_waste.Effective_date = formatDateString(sf_waste.Effective_date)
	sf_waste.Expiry_date = formatDateString(sf_waste.Expiry_date)
	sf_waste.License_expiry_date = formatDateString(sf_waste.License_expiry_date)
	sf_waste.Row_effective_date = formatDateString(sf_waste.Row_effective_date)
	sf_waste.Row_expiry_date = formatDateString(sf_waste.Row_expiry_date)






	rows, err := stmt.Exec(sf_waste.Waste_facility_id, sf_waste.Active_ind, sf_waste.Capacity, sf_waste.Capacity_ouom, sf_waste.Current_owner, sf_waste.Effective_date, sf_waste.Expiry_date, sf_waste.License_expiry_date, sf_waste.License_num, sf_waste.License_register_ba_id, sf_waste.Ppdm_guid, sf_waste.Remark, sf_waste.Source, sf_waste.Waste_facility_name, sf_waste.Waste_facility_type, sf_waste.Row_changed_by, sf_waste.Row_changed_date, sf_waste.Row_created_by, sf_waste.Row_effective_date, sf_waste.Row_expiry_date, sf_waste.Row_quality, sf_waste.Sf_subtype)
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

func PatchSfWaste(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_waste set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "license_expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSfWaste(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_waste dto.Sf_waste
	sf_waste.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_waste where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_waste.Sf_subtype)
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


