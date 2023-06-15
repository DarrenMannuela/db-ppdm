package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmSystemMap(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_system_map")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_system_map

	for rows.Next() {
		var ppdm_system_map dto.Ppdm_system_map
		if err := rows.Scan(&ppdm_system_map.Map_id, &ppdm_system_map.Active_ind, &ppdm_system_map.Effective_date, &ppdm_system_map.Expiry_date, &ppdm_system_map.Map_owner_ba_id, &ppdm_system_map.Map_version_num, &ppdm_system_map.Ppdm_guid, &ppdm_system_map.Preferred_ind, &ppdm_system_map.Remark, &ppdm_system_map.Source, &ppdm_system_map.Sw_application_id, &ppdm_system_map.System_id1, &ppdm_system_map.System_id2, &ppdm_system_map.Row_changed_by, &ppdm_system_map.Row_changed_date, &ppdm_system_map.Row_created_by, &ppdm_system_map.Row_created_date, &ppdm_system_map.Row_effective_date, &ppdm_system_map.Row_expiry_date, &ppdm_system_map.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_system_map to result
		result = append(result, ppdm_system_map)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmSystemMap(c *fiber.Ctx) error {
	var ppdm_system_map dto.Ppdm_system_map

	setDefaults(&ppdm_system_map)

	if err := json.Unmarshal(c.Body(), &ppdm_system_map); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_system_map values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	ppdm_system_map.Row_created_date = formatDate(ppdm_system_map.Row_created_date)
	ppdm_system_map.Row_changed_date = formatDate(ppdm_system_map.Row_changed_date)
	ppdm_system_map.Effective_date = formatDateString(ppdm_system_map.Effective_date)
	ppdm_system_map.Expiry_date = formatDateString(ppdm_system_map.Expiry_date)
	ppdm_system_map.Row_effective_date = formatDateString(ppdm_system_map.Row_effective_date)
	ppdm_system_map.Row_expiry_date = formatDateString(ppdm_system_map.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_system_map.Map_id, ppdm_system_map.Active_ind, ppdm_system_map.Effective_date, ppdm_system_map.Expiry_date, ppdm_system_map.Map_owner_ba_id, ppdm_system_map.Map_version_num, ppdm_system_map.Ppdm_guid, ppdm_system_map.Preferred_ind, ppdm_system_map.Remark, ppdm_system_map.Source, ppdm_system_map.Sw_application_id, ppdm_system_map.System_id1, ppdm_system_map.System_id2, ppdm_system_map.Row_changed_by, ppdm_system_map.Row_changed_date, ppdm_system_map.Row_created_by, ppdm_system_map.Row_created_date, ppdm_system_map.Row_effective_date, ppdm_system_map.Row_expiry_date, ppdm_system_map.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmSystemMap(c *fiber.Ctx) error {
	var ppdm_system_map dto.Ppdm_system_map

	setDefaults(&ppdm_system_map)

	if err := json.Unmarshal(c.Body(), &ppdm_system_map); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_system_map.Map_id = id

    if ppdm_system_map.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_system_map set active_ind = :1, effective_date = :2, expiry_date = :3, map_owner_ba_id = :4, map_version_num = :5, ppdm_guid = :6, preferred_ind = :7, remark = :8, source = :9, sw_application_id = :10, system_id1 = :11, system_id2 = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where map_id = :20")
	if err != nil {
		return err
	}

	ppdm_system_map.Row_changed_date = formatDate(ppdm_system_map.Row_changed_date)
	ppdm_system_map.Effective_date = formatDateString(ppdm_system_map.Effective_date)
	ppdm_system_map.Expiry_date = formatDateString(ppdm_system_map.Expiry_date)
	ppdm_system_map.Row_effective_date = formatDateString(ppdm_system_map.Row_effective_date)
	ppdm_system_map.Row_expiry_date = formatDateString(ppdm_system_map.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_system_map.Active_ind, ppdm_system_map.Effective_date, ppdm_system_map.Expiry_date, ppdm_system_map.Map_owner_ba_id, ppdm_system_map.Map_version_num, ppdm_system_map.Ppdm_guid, ppdm_system_map.Preferred_ind, ppdm_system_map.Remark, ppdm_system_map.Source, ppdm_system_map.Sw_application_id, ppdm_system_map.System_id1, ppdm_system_map.System_id2, ppdm_system_map.Row_changed_by, ppdm_system_map.Row_changed_date, ppdm_system_map.Row_created_by, ppdm_system_map.Row_effective_date, ppdm_system_map.Row_expiry_date, ppdm_system_map.Row_quality, ppdm_system_map.Map_id)
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

func PatchPpdmSystemMap(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_system_map set "
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
	query += " where map_id = :id"

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

func DeletePpdmSystemMap(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_system_map dto.Ppdm_system_map
	ppdm_system_map.Map_id = id

	stmt, err := db.Prepare("delete from ppdm_system_map where map_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_system_map.Map_id)
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


