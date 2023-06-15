package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAreaHierarchy(c *fiber.Ctx) error {
	rows, err := db.Query("select * from area_hierarchy")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Area_hierarchy

	for rows.Next() {
		var area_hierarchy dto.Area_hierarchy
		if err := rows.Scan(&area_hierarchy.Area_hierarchy_id, &area_hierarchy.Active_ind, &area_hierarchy.Area_hierarchy_name, &area_hierarchy.Effective_date, &area_hierarchy.Expiry_date, &area_hierarchy.Ppdm_guid, &area_hierarchy.Remark, &area_hierarchy.Source, &area_hierarchy.Row_changed_by, &area_hierarchy.Row_changed_date, &area_hierarchy.Row_created_by, &area_hierarchy.Row_created_date, &area_hierarchy.Row_effective_date, &area_hierarchy.Row_expiry_date, &area_hierarchy.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append area_hierarchy to result
		result = append(result, area_hierarchy)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAreaHierarchy(c *fiber.Ctx) error {
	var area_hierarchy dto.Area_hierarchy

	setDefaults(&area_hierarchy)

	if err := json.Unmarshal(c.Body(), &area_hierarchy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into area_hierarchy values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	area_hierarchy.Row_created_date = formatDate(area_hierarchy.Row_created_date)
	area_hierarchy.Row_changed_date = formatDate(area_hierarchy.Row_changed_date)
	area_hierarchy.Effective_date = formatDateString(area_hierarchy.Effective_date)
	area_hierarchy.Expiry_date = formatDateString(area_hierarchy.Expiry_date)
	area_hierarchy.Row_effective_date = formatDateString(area_hierarchy.Row_effective_date)
	area_hierarchy.Row_expiry_date = formatDateString(area_hierarchy.Row_expiry_date)






	rows, err := stmt.Exec(area_hierarchy.Area_hierarchy_id, area_hierarchy.Active_ind, area_hierarchy.Area_hierarchy_name, area_hierarchy.Effective_date, area_hierarchy.Expiry_date, area_hierarchy.Ppdm_guid, area_hierarchy.Remark, area_hierarchy.Source, area_hierarchy.Row_changed_by, area_hierarchy.Row_changed_date, area_hierarchy.Row_created_by, area_hierarchy.Row_created_date, area_hierarchy.Row_effective_date, area_hierarchy.Row_expiry_date, area_hierarchy.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAreaHierarchy(c *fiber.Ctx) error {
	var area_hierarchy dto.Area_hierarchy

	setDefaults(&area_hierarchy)

	if err := json.Unmarshal(c.Body(), &area_hierarchy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	area_hierarchy.Area_hierarchy_id = id

    if area_hierarchy.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update area_hierarchy set active_ind = :1, area_hierarchy_name = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where area_hierarchy_id = :15")
	if err != nil {
		return err
	}

	area_hierarchy.Row_changed_date = formatDate(area_hierarchy.Row_changed_date)
	area_hierarchy.Effective_date = formatDateString(area_hierarchy.Effective_date)
	area_hierarchy.Expiry_date = formatDateString(area_hierarchy.Expiry_date)
	area_hierarchy.Row_effective_date = formatDateString(area_hierarchy.Row_effective_date)
	area_hierarchy.Row_expiry_date = formatDateString(area_hierarchy.Row_expiry_date)






	rows, err := stmt.Exec(area_hierarchy.Active_ind, area_hierarchy.Area_hierarchy_name, area_hierarchy.Effective_date, area_hierarchy.Expiry_date, area_hierarchy.Ppdm_guid, area_hierarchy.Remark, area_hierarchy.Source, area_hierarchy.Row_changed_by, area_hierarchy.Row_changed_date, area_hierarchy.Row_created_by, area_hierarchy.Row_effective_date, area_hierarchy.Row_expiry_date, area_hierarchy.Row_quality, area_hierarchy.Area_hierarchy_id)
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

func PatchAreaHierarchy(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update area_hierarchy set "
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
	query += " where area_hierarchy_id = :id"

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

func DeleteAreaHierarchy(c *fiber.Ctx) error {
	id := c.Params("id")
	var area_hierarchy dto.Area_hierarchy
	area_hierarchy.Area_hierarchy_id = id

	stmt, err := db.Prepare("delete from area_hierarchy where area_hierarchy_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(area_hierarchy.Area_hierarchy_id)
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


