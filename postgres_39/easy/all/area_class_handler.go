package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAreaClass(c *fiber.Ctx) error {
	rows, err := db.Query("select * from area_class")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Area_class

	for rows.Next() {
		var area_class dto.Area_class
		if err := rows.Scan(&area_class.Area_id_parent, &area_class.Area_type_parent, &area_class.Area_id, &area_class.Area_type, &area_class.Area_hierarchy_id, &area_class.Active_ind, &area_class.Effective_date, &area_class.Expiry_date, &area_class.Ppdm_guid, &area_class.Remark, &area_class.Source, &area_class.Row_changed_by, &area_class.Row_changed_date, &area_class.Row_created_by, &area_class.Row_created_date, &area_class.Row_effective_date, &area_class.Row_expiry_date, &area_class.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append area_class to result
		result = append(result, area_class)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAreaClass(c *fiber.Ctx) error {
	var area_class dto.Area_class

	setDefaults(&area_class)

	if err := json.Unmarshal(c.Body(), &area_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into area_class values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	area_class.Row_created_date = formatDate(area_class.Row_created_date)
	area_class.Row_changed_date = formatDate(area_class.Row_changed_date)
	area_class.Effective_date = formatDateString(area_class.Effective_date)
	area_class.Expiry_date = formatDateString(area_class.Expiry_date)
	area_class.Row_effective_date = formatDateString(area_class.Row_effective_date)
	area_class.Row_expiry_date = formatDateString(area_class.Row_expiry_date)






	rows, err := stmt.Exec(area_class.Area_id_parent, area_class.Area_type_parent, area_class.Area_id, area_class.Area_type, area_class.Area_hierarchy_id, area_class.Active_ind, area_class.Effective_date, area_class.Expiry_date, area_class.Ppdm_guid, area_class.Remark, area_class.Source, area_class.Row_changed_by, area_class.Row_changed_date, area_class.Row_created_by, area_class.Row_created_date, area_class.Row_effective_date, area_class.Row_expiry_date, area_class.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAreaClass(c *fiber.Ctx) error {
	var area_class dto.Area_class

	setDefaults(&area_class)

	if err := json.Unmarshal(c.Body(), &area_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	area_class.Area_id_parent = id

    if area_class.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update area_class set area_type_parent = :1, area_id = :2, area_type = :3, area_hierarchy_id = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where area_id_parent = :18")
	if err != nil {
		return err
	}

	area_class.Row_changed_date = formatDate(area_class.Row_changed_date)
	area_class.Effective_date = formatDateString(area_class.Effective_date)
	area_class.Expiry_date = formatDateString(area_class.Expiry_date)
	area_class.Row_effective_date = formatDateString(area_class.Row_effective_date)
	area_class.Row_expiry_date = formatDateString(area_class.Row_expiry_date)






	rows, err := stmt.Exec(area_class.Area_type_parent, area_class.Area_id, area_class.Area_type, area_class.Area_hierarchy_id, area_class.Active_ind, area_class.Effective_date, area_class.Expiry_date, area_class.Ppdm_guid, area_class.Remark, area_class.Source, area_class.Row_changed_by, area_class.Row_changed_date, area_class.Row_created_by, area_class.Row_effective_date, area_class.Row_expiry_date, area_class.Row_quality, area_class.Area_id_parent)
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

func PatchAreaClass(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update area_class set "
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
	query += " where area_id_parent = :id"

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

func DeleteAreaClass(c *fiber.Ctx) error {
	id := c.Params("id")
	var area_class dto.Area_class
	area_class.Area_id_parent = id

	stmt, err := db.Prepare("delete from area_class where area_id_parent = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(area_class.Area_id_parent)
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


