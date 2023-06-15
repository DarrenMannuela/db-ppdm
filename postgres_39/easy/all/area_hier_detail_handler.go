package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAreaHierDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from area_hier_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Area_hier_detail

	for rows.Next() {
		var area_hier_detail dto.Area_hier_detail
		if err := rows.Scan(&area_hier_detail.Area_hierarchy_id, &area_hier_detail.Area_hier_level_seq_no, &area_hier_detail.Active_ind, &area_hier_detail.Area_type, &area_hier_detail.Effective_date, &area_hier_detail.Expiry_date, &area_hier_detail.Ppdm_guid, &area_hier_detail.Remark, &area_hier_detail.Source, &area_hier_detail.Row_changed_by, &area_hier_detail.Row_changed_date, &area_hier_detail.Row_created_by, &area_hier_detail.Row_created_date, &area_hier_detail.Row_effective_date, &area_hier_detail.Row_expiry_date, &area_hier_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append area_hier_detail to result
		result = append(result, area_hier_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAreaHierDetail(c *fiber.Ctx) error {
	var area_hier_detail dto.Area_hier_detail

	setDefaults(&area_hier_detail)

	if err := json.Unmarshal(c.Body(), &area_hier_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into area_hier_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	area_hier_detail.Row_created_date = formatDate(area_hier_detail.Row_created_date)
	area_hier_detail.Row_changed_date = formatDate(area_hier_detail.Row_changed_date)
	area_hier_detail.Effective_date = formatDateString(area_hier_detail.Effective_date)
	area_hier_detail.Expiry_date = formatDateString(area_hier_detail.Expiry_date)
	area_hier_detail.Row_effective_date = formatDateString(area_hier_detail.Row_effective_date)
	area_hier_detail.Row_expiry_date = formatDateString(area_hier_detail.Row_expiry_date)






	rows, err := stmt.Exec(area_hier_detail.Area_hierarchy_id, area_hier_detail.Area_hier_level_seq_no, area_hier_detail.Active_ind, area_hier_detail.Area_type, area_hier_detail.Effective_date, area_hier_detail.Expiry_date, area_hier_detail.Ppdm_guid, area_hier_detail.Remark, area_hier_detail.Source, area_hier_detail.Row_changed_by, area_hier_detail.Row_changed_date, area_hier_detail.Row_created_by, area_hier_detail.Row_created_date, area_hier_detail.Row_effective_date, area_hier_detail.Row_expiry_date, area_hier_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAreaHierDetail(c *fiber.Ctx) error {
	var area_hier_detail dto.Area_hier_detail

	setDefaults(&area_hier_detail)

	if err := json.Unmarshal(c.Body(), &area_hier_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	area_hier_detail.Area_hierarchy_id = id

    if area_hier_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update area_hier_detail set area_hier_level_seq_no = :1, active_ind = :2, area_type = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where area_hierarchy_id = :16")
	if err != nil {
		return err
	}

	area_hier_detail.Row_changed_date = formatDate(area_hier_detail.Row_changed_date)
	area_hier_detail.Effective_date = formatDateString(area_hier_detail.Effective_date)
	area_hier_detail.Expiry_date = formatDateString(area_hier_detail.Expiry_date)
	area_hier_detail.Row_effective_date = formatDateString(area_hier_detail.Row_effective_date)
	area_hier_detail.Row_expiry_date = formatDateString(area_hier_detail.Row_expiry_date)






	rows, err := stmt.Exec(area_hier_detail.Area_hier_level_seq_no, area_hier_detail.Active_ind, area_hier_detail.Area_type, area_hier_detail.Effective_date, area_hier_detail.Expiry_date, area_hier_detail.Ppdm_guid, area_hier_detail.Remark, area_hier_detail.Source, area_hier_detail.Row_changed_by, area_hier_detail.Row_changed_date, area_hier_detail.Row_created_by, area_hier_detail.Row_effective_date, area_hier_detail.Row_expiry_date, area_hier_detail.Row_quality, area_hier_detail.Area_hierarchy_id)
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

func PatchAreaHierDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update area_hier_detail set "
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

func DeleteAreaHierDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var area_hier_detail dto.Area_hier_detail
	area_hier_detail.Area_hierarchy_id = id

	stmt, err := db.Prepare("delete from area_hier_detail where area_hierarchy_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(area_hier_detail.Area_hierarchy_id)
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


