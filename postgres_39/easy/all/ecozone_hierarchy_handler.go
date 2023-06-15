package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEcozoneHierarchy(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ecozone_hierarchy")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ecozone_hierarchy

	for rows.Next() {
		var ecozone_hierarchy dto.Ecozone_hierarchy
		if err := rows.Scan(&ecozone_hierarchy.Ecozone_set_id, &ecozone_hierarchy.Parent_ecozone_id, &ecozone_hierarchy.Child_ecozone_id, &ecozone_hierarchy.Active_ind, &ecozone_hierarchy.Effective_date, &ecozone_hierarchy.Expiry_date, &ecozone_hierarchy.Hierarchy_level, &ecozone_hierarchy.Ppdm_guid, &ecozone_hierarchy.Remark, &ecozone_hierarchy.Source, &ecozone_hierarchy.Row_changed_by, &ecozone_hierarchy.Row_changed_date, &ecozone_hierarchy.Row_created_by, &ecozone_hierarchy.Row_created_date, &ecozone_hierarchy.Row_effective_date, &ecozone_hierarchy.Row_expiry_date, &ecozone_hierarchy.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ecozone_hierarchy to result
		result = append(result, ecozone_hierarchy)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEcozoneHierarchy(c *fiber.Ctx) error {
	var ecozone_hierarchy dto.Ecozone_hierarchy

	setDefaults(&ecozone_hierarchy)

	if err := json.Unmarshal(c.Body(), &ecozone_hierarchy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ecozone_hierarchy values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	ecozone_hierarchy.Row_created_date = formatDate(ecozone_hierarchy.Row_created_date)
	ecozone_hierarchy.Row_changed_date = formatDate(ecozone_hierarchy.Row_changed_date)
	ecozone_hierarchy.Effective_date = formatDateString(ecozone_hierarchy.Effective_date)
	ecozone_hierarchy.Expiry_date = formatDateString(ecozone_hierarchy.Expiry_date)
	ecozone_hierarchy.Row_effective_date = formatDateString(ecozone_hierarchy.Row_effective_date)
	ecozone_hierarchy.Row_expiry_date = formatDateString(ecozone_hierarchy.Row_expiry_date)






	rows, err := stmt.Exec(ecozone_hierarchy.Ecozone_set_id, ecozone_hierarchy.Parent_ecozone_id, ecozone_hierarchy.Child_ecozone_id, ecozone_hierarchy.Active_ind, ecozone_hierarchy.Effective_date, ecozone_hierarchy.Expiry_date, ecozone_hierarchy.Hierarchy_level, ecozone_hierarchy.Ppdm_guid, ecozone_hierarchy.Remark, ecozone_hierarchy.Source, ecozone_hierarchy.Row_changed_by, ecozone_hierarchy.Row_changed_date, ecozone_hierarchy.Row_created_by, ecozone_hierarchy.Row_created_date, ecozone_hierarchy.Row_effective_date, ecozone_hierarchy.Row_expiry_date, ecozone_hierarchy.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEcozoneHierarchy(c *fiber.Ctx) error {
	var ecozone_hierarchy dto.Ecozone_hierarchy

	setDefaults(&ecozone_hierarchy)

	if err := json.Unmarshal(c.Body(), &ecozone_hierarchy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ecozone_hierarchy.Ecozone_set_id = id

    if ecozone_hierarchy.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ecozone_hierarchy set parent_ecozone_id = :1, child_ecozone_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, hierarchy_level = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where ecozone_set_id = :17")
	if err != nil {
		return err
	}

	ecozone_hierarchy.Row_changed_date = formatDate(ecozone_hierarchy.Row_changed_date)
	ecozone_hierarchy.Effective_date = formatDateString(ecozone_hierarchy.Effective_date)
	ecozone_hierarchy.Expiry_date = formatDateString(ecozone_hierarchy.Expiry_date)
	ecozone_hierarchy.Row_effective_date = formatDateString(ecozone_hierarchy.Row_effective_date)
	ecozone_hierarchy.Row_expiry_date = formatDateString(ecozone_hierarchy.Row_expiry_date)






	rows, err := stmt.Exec(ecozone_hierarchy.Parent_ecozone_id, ecozone_hierarchy.Child_ecozone_id, ecozone_hierarchy.Active_ind, ecozone_hierarchy.Effective_date, ecozone_hierarchy.Expiry_date, ecozone_hierarchy.Hierarchy_level, ecozone_hierarchy.Ppdm_guid, ecozone_hierarchy.Remark, ecozone_hierarchy.Source, ecozone_hierarchy.Row_changed_by, ecozone_hierarchy.Row_changed_date, ecozone_hierarchy.Row_created_by, ecozone_hierarchy.Row_effective_date, ecozone_hierarchy.Row_expiry_date, ecozone_hierarchy.Row_quality, ecozone_hierarchy.Ecozone_set_id)
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

func PatchEcozoneHierarchy(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ecozone_hierarchy set "
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
	query += " where ecozone_set_id = :id"

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

func DeleteEcozoneHierarchy(c *fiber.Ctx) error {
	id := c.Params("id")
	var ecozone_hierarchy dto.Ecozone_hierarchy
	ecozone_hierarchy.Ecozone_set_id = id

	stmt, err := db.Prepare("delete from ecozone_hierarchy where ecozone_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ecozone_hierarchy.Ecozone_set_id)
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


