package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandUnit(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_unit")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_unit

	for rows.Next() {
		var land_unit dto.Land_unit
		if err := rows.Scan(&land_unit.Land_right_subtype, &land_unit.Land_right_id, &land_unit.Active_ind, &land_unit.Effective_date, &land_unit.Expiry_date, &land_unit.Land_unit_name, &land_unit.Land_unit_number, &land_unit.Land_unit_type, &land_unit.Ppdm_guid, &land_unit.Remark, &land_unit.Source, &land_unit.Row_changed_by, &land_unit.Row_changed_date, &land_unit.Row_created_by, &land_unit.Row_created_date, &land_unit.Row_effective_date, &land_unit.Row_expiry_date, &land_unit.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_unit to result
		result = append(result, land_unit)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandUnit(c *fiber.Ctx) error {
	var land_unit dto.Land_unit

	setDefaults(&land_unit)

	if err := json.Unmarshal(c.Body(), &land_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_unit values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	land_unit.Row_created_date = formatDate(land_unit.Row_created_date)
	land_unit.Row_changed_date = formatDate(land_unit.Row_changed_date)
	land_unit.Effective_date = formatDateString(land_unit.Effective_date)
	land_unit.Expiry_date = formatDateString(land_unit.Expiry_date)
	land_unit.Row_effective_date = formatDateString(land_unit.Row_effective_date)
	land_unit.Row_expiry_date = formatDateString(land_unit.Row_expiry_date)






	rows, err := stmt.Exec(land_unit.Land_right_subtype, land_unit.Land_right_id, land_unit.Active_ind, land_unit.Effective_date, land_unit.Expiry_date, land_unit.Land_unit_name, land_unit.Land_unit_number, land_unit.Land_unit_type, land_unit.Ppdm_guid, land_unit.Remark, land_unit.Source, land_unit.Row_changed_by, land_unit.Row_changed_date, land_unit.Row_created_by, land_unit.Row_created_date, land_unit.Row_effective_date, land_unit.Row_expiry_date, land_unit.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandUnit(c *fiber.Ctx) error {
	var land_unit dto.Land_unit

	setDefaults(&land_unit)

	if err := json.Unmarshal(c.Body(), &land_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_unit.Land_right_subtype = id

    if land_unit.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_unit set land_right_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, land_unit_name = :5, land_unit_number = :6, land_unit_type = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where land_right_subtype = :18")
	if err != nil {
		return err
	}

	land_unit.Row_changed_date = formatDate(land_unit.Row_changed_date)
	land_unit.Effective_date = formatDateString(land_unit.Effective_date)
	land_unit.Expiry_date = formatDateString(land_unit.Expiry_date)
	land_unit.Row_effective_date = formatDateString(land_unit.Row_effective_date)
	land_unit.Row_expiry_date = formatDateString(land_unit.Row_expiry_date)






	rows, err := stmt.Exec(land_unit.Land_right_id, land_unit.Active_ind, land_unit.Effective_date, land_unit.Expiry_date, land_unit.Land_unit_name, land_unit.Land_unit_number, land_unit.Land_unit_type, land_unit.Ppdm_guid, land_unit.Remark, land_unit.Source, land_unit.Row_changed_by, land_unit.Row_changed_date, land_unit.Row_created_by, land_unit.Row_effective_date, land_unit.Row_expiry_date, land_unit.Row_quality, land_unit.Land_right_subtype)
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

func PatchLandUnit(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_unit set "
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
	query += " where land_right_subtype = :id"

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

func DeleteLandUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_unit dto.Land_unit
	land_unit.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_unit where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_unit.Land_right_subtype)
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


