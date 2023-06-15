package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfBridge(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_bridge")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_bridge

	for rows.Next() {
		var sf_bridge dto.Sf_bridge
		if err := rows.Scan(&sf_bridge.Sf_subtype, &sf_bridge.Bridge_id, &sf_bridge.Active_ind, &sf_bridge.Bridge_capacity, &sf_bridge.Bridge_capacity_ouom, &sf_bridge.Bridge_height, &sf_bridge.Bridge_height_ouom, &sf_bridge.Bridge_length, &sf_bridge.Bridge_length_ouom, &sf_bridge.Bridge_type, &sf_bridge.Bridge_width, &sf_bridge.Bridge_width_ouom, &sf_bridge.Effective_date, &sf_bridge.Expiry_date, &sf_bridge.Ppdm_guid, &sf_bridge.Remark, &sf_bridge.Source, &sf_bridge.Surface_type, &sf_bridge.Row_changed_by, &sf_bridge.Row_changed_date, &sf_bridge.Row_created_by, &sf_bridge.Row_created_date, &sf_bridge.Row_effective_date, &sf_bridge.Row_expiry_date, &sf_bridge.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_bridge to result
		result = append(result, sf_bridge)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfBridge(c *fiber.Ctx) error {
	var sf_bridge dto.Sf_bridge

	setDefaults(&sf_bridge)

	if err := json.Unmarshal(c.Body(), &sf_bridge); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_bridge values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	sf_bridge.Row_created_date = formatDate(sf_bridge.Row_created_date)
	sf_bridge.Row_changed_date = formatDate(sf_bridge.Row_changed_date)
	sf_bridge.Effective_date = formatDateString(sf_bridge.Effective_date)
	sf_bridge.Expiry_date = formatDateString(sf_bridge.Expiry_date)
	sf_bridge.Row_effective_date = formatDateString(sf_bridge.Row_effective_date)
	sf_bridge.Row_expiry_date = formatDateString(sf_bridge.Row_expiry_date)






	rows, err := stmt.Exec(sf_bridge.Sf_subtype, sf_bridge.Bridge_id, sf_bridge.Active_ind, sf_bridge.Bridge_capacity, sf_bridge.Bridge_capacity_ouom, sf_bridge.Bridge_height, sf_bridge.Bridge_height_ouom, sf_bridge.Bridge_length, sf_bridge.Bridge_length_ouom, sf_bridge.Bridge_type, sf_bridge.Bridge_width, sf_bridge.Bridge_width_ouom, sf_bridge.Effective_date, sf_bridge.Expiry_date, sf_bridge.Ppdm_guid, sf_bridge.Remark, sf_bridge.Source, sf_bridge.Surface_type, sf_bridge.Row_changed_by, sf_bridge.Row_changed_date, sf_bridge.Row_created_by, sf_bridge.Row_created_date, sf_bridge.Row_effective_date, sf_bridge.Row_expiry_date, sf_bridge.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfBridge(c *fiber.Ctx) error {
	var sf_bridge dto.Sf_bridge

	setDefaults(&sf_bridge)

	if err := json.Unmarshal(c.Body(), &sf_bridge); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_bridge.Sf_subtype = id

    if sf_bridge.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_bridge set bridge_id = :1, active_ind = :2, bridge_capacity = :3, bridge_capacity_ouom = :4, bridge_height = :5, bridge_height_ouom = :6, bridge_length = :7, bridge_length_ouom = :8, bridge_type = :9, bridge_width = :10, bridge_width_ouom = :11, effective_date = :12, expiry_date = :13, ppdm_guid = :14, remark = :15, source = :16, surface_type = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where sf_subtype = :25")
	if err != nil {
		return err
	}

	sf_bridge.Row_changed_date = formatDate(sf_bridge.Row_changed_date)
	sf_bridge.Effective_date = formatDateString(sf_bridge.Effective_date)
	sf_bridge.Expiry_date = formatDateString(sf_bridge.Expiry_date)
	sf_bridge.Row_effective_date = formatDateString(sf_bridge.Row_effective_date)
	sf_bridge.Row_expiry_date = formatDateString(sf_bridge.Row_expiry_date)






	rows, err := stmt.Exec(sf_bridge.Bridge_id, sf_bridge.Active_ind, sf_bridge.Bridge_capacity, sf_bridge.Bridge_capacity_ouom, sf_bridge.Bridge_height, sf_bridge.Bridge_height_ouom, sf_bridge.Bridge_length, sf_bridge.Bridge_length_ouom, sf_bridge.Bridge_type, sf_bridge.Bridge_width, sf_bridge.Bridge_width_ouom, sf_bridge.Effective_date, sf_bridge.Expiry_date, sf_bridge.Ppdm_guid, sf_bridge.Remark, sf_bridge.Source, sf_bridge.Surface_type, sf_bridge.Row_changed_by, sf_bridge.Row_changed_date, sf_bridge.Row_created_by, sf_bridge.Row_effective_date, sf_bridge.Row_expiry_date, sf_bridge.Row_quality, sf_bridge.Sf_subtype)
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

func PatchSfBridge(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_bridge set "
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

func DeleteSfBridge(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_bridge dto.Sf_bridge
	sf_bridge.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_bridge where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_bridge.Sf_subtype)
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


