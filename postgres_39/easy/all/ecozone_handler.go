package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEcozone(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ecozone")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ecozone

	for rows.Next() {
		var ecozone dto.Ecozone
		if err := rows.Scan(&ecozone.Ecozone_id, &ecozone.Active_ind, &ecozone.Base_depth, &ecozone.Depth_ouom, &ecozone.Ecozone_type, &ecozone.Effective_date, &ecozone.Expiry_date, &ecozone.Ppdm_guid, &ecozone.Preferred_name, &ecozone.Remark, &ecozone.Source, &ecozone.Top_depth, &ecozone.Row_changed_by, &ecozone.Row_changed_date, &ecozone.Row_created_by, &ecozone.Row_created_date, &ecozone.Row_effective_date, &ecozone.Row_expiry_date, &ecozone.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ecozone to result
		result = append(result, ecozone)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEcozone(c *fiber.Ctx) error {
	var ecozone dto.Ecozone

	setDefaults(&ecozone)

	if err := json.Unmarshal(c.Body(), &ecozone); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ecozone values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	ecozone.Row_created_date = formatDate(ecozone.Row_created_date)
	ecozone.Row_changed_date = formatDate(ecozone.Row_changed_date)
	ecozone.Effective_date = formatDateString(ecozone.Effective_date)
	ecozone.Expiry_date = formatDateString(ecozone.Expiry_date)
	ecozone.Row_effective_date = formatDateString(ecozone.Row_effective_date)
	ecozone.Row_expiry_date = formatDateString(ecozone.Row_expiry_date)






	rows, err := stmt.Exec(ecozone.Ecozone_id, ecozone.Active_ind, ecozone.Base_depth, ecozone.Depth_ouom, ecozone.Ecozone_type, ecozone.Effective_date, ecozone.Expiry_date, ecozone.Ppdm_guid, ecozone.Preferred_name, ecozone.Remark, ecozone.Source, ecozone.Top_depth, ecozone.Row_changed_by, ecozone.Row_changed_date, ecozone.Row_created_by, ecozone.Row_created_date, ecozone.Row_effective_date, ecozone.Row_expiry_date, ecozone.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEcozone(c *fiber.Ctx) error {
	var ecozone dto.Ecozone

	setDefaults(&ecozone)

	if err := json.Unmarshal(c.Body(), &ecozone); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ecozone.Ecozone_id = id

    if ecozone.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ecozone set active_ind = :1, base_depth = :2, depth_ouom = :3, ecozone_type = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, preferred_name = :8, remark = :9, source = :10, top_depth = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where ecozone_id = :19")
	if err != nil {
		return err
	}

	ecozone.Row_changed_date = formatDate(ecozone.Row_changed_date)
	ecozone.Effective_date = formatDateString(ecozone.Effective_date)
	ecozone.Expiry_date = formatDateString(ecozone.Expiry_date)
	ecozone.Row_effective_date = formatDateString(ecozone.Row_effective_date)
	ecozone.Row_expiry_date = formatDateString(ecozone.Row_expiry_date)






	rows, err := stmt.Exec(ecozone.Active_ind, ecozone.Base_depth, ecozone.Depth_ouom, ecozone.Ecozone_type, ecozone.Effective_date, ecozone.Expiry_date, ecozone.Ppdm_guid, ecozone.Preferred_name, ecozone.Remark, ecozone.Source, ecozone.Top_depth, ecozone.Row_changed_by, ecozone.Row_changed_date, ecozone.Row_created_by, ecozone.Row_effective_date, ecozone.Row_expiry_date, ecozone.Row_quality, ecozone.Ecozone_id)
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

func PatchEcozone(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ecozone set "
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
	query += " where ecozone_id = :id"

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

func DeleteEcozone(c *fiber.Ctx) error {
	id := c.Params("id")
	var ecozone dto.Ecozone
	ecozone.Ecozone_id = id

	stmt, err := db.Prepare("delete from ecozone where ecozone_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ecozone.Ecozone_id)
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


