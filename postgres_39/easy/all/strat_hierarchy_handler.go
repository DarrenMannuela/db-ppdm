package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratHierarchy(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_hierarchy")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_hierarchy

	for rows.Next() {
		var strat_hierarchy dto.Strat_hierarchy
		if err := rows.Scan(&strat_hierarchy.Parent_strat_name_set_id, &strat_hierarchy.Parent_strat_unit_id, &strat_hierarchy.Child_strat_name_set_id, &strat_hierarchy.Child_strat_unit_id, &strat_hierarchy.Hierarchy_id, &strat_hierarchy.Source, &strat_hierarchy.Active_ind, &strat_hierarchy.Area_id, &strat_hierarchy.Area_type, &strat_hierarchy.Child_interp_id2, &strat_hierarchy.Child_strat_column_id, &strat_hierarchy.Child_strat_column_source, &strat_hierarchy.Effective_date, &strat_hierarchy.Expiry_date, &strat_hierarchy.Parent_interp_id, &strat_hierarchy.Parent_strat_column_id, &strat_hierarchy.Parent_strat_column_source, &strat_hierarchy.Ppdm_guid, &strat_hierarchy.Preferred_hierarchy_ind, &strat_hierarchy.Remark, &strat_hierarchy.Source_document_id, &strat_hierarchy.Strat_hierarchy_type, &strat_hierarchy.Row_changed_by, &strat_hierarchy.Row_changed_date, &strat_hierarchy.Row_created_by, &strat_hierarchy.Row_created_date, &strat_hierarchy.Row_effective_date, &strat_hierarchy.Row_expiry_date, &strat_hierarchy.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_hierarchy to result
		result = append(result, strat_hierarchy)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratHierarchy(c *fiber.Ctx) error {
	var strat_hierarchy dto.Strat_hierarchy

	setDefaults(&strat_hierarchy)

	if err := json.Unmarshal(c.Body(), &strat_hierarchy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_hierarchy values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	strat_hierarchy.Row_created_date = formatDate(strat_hierarchy.Row_created_date)
	strat_hierarchy.Row_changed_date = formatDate(strat_hierarchy.Row_changed_date)
	strat_hierarchy.Effective_date = formatDateString(strat_hierarchy.Effective_date)
	strat_hierarchy.Expiry_date = formatDateString(strat_hierarchy.Expiry_date)
	strat_hierarchy.Row_effective_date = formatDateString(strat_hierarchy.Row_effective_date)
	strat_hierarchy.Row_expiry_date = formatDateString(strat_hierarchy.Row_expiry_date)






	rows, err := stmt.Exec(strat_hierarchy.Parent_strat_name_set_id, strat_hierarchy.Parent_strat_unit_id, strat_hierarchy.Child_strat_name_set_id, strat_hierarchy.Child_strat_unit_id, strat_hierarchy.Hierarchy_id, strat_hierarchy.Source, strat_hierarchy.Active_ind, strat_hierarchy.Area_id, strat_hierarchy.Area_type, strat_hierarchy.Child_interp_id2, strat_hierarchy.Child_strat_column_id, strat_hierarchy.Child_strat_column_source, strat_hierarchy.Effective_date, strat_hierarchy.Expiry_date, strat_hierarchy.Parent_interp_id, strat_hierarchy.Parent_strat_column_id, strat_hierarchy.Parent_strat_column_source, strat_hierarchy.Ppdm_guid, strat_hierarchy.Preferred_hierarchy_ind, strat_hierarchy.Remark, strat_hierarchy.Source_document_id, strat_hierarchy.Strat_hierarchy_type, strat_hierarchy.Row_changed_by, strat_hierarchy.Row_changed_date, strat_hierarchy.Row_created_by, strat_hierarchy.Row_created_date, strat_hierarchy.Row_effective_date, strat_hierarchy.Row_expiry_date, strat_hierarchy.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratHierarchy(c *fiber.Ctx) error {
	var strat_hierarchy dto.Strat_hierarchy

	setDefaults(&strat_hierarchy)

	if err := json.Unmarshal(c.Body(), &strat_hierarchy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_hierarchy.Parent_strat_name_set_id = id

    if strat_hierarchy.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_hierarchy set parent_strat_unit_id = :1, child_strat_name_set_id = :2, child_strat_unit_id = :3, hierarchy_id = :4, source = :5, active_ind = :6, area_id = :7, area_type = :8, child_interp_id2 = :9, child_strat_column_id = :10, child_strat_column_source = :11, effective_date = :12, expiry_date = :13, parent_interp_id = :14, parent_strat_column_id = :15, parent_strat_column_source = :16, ppdm_guid = :17, preferred_hierarchy_ind = :18, remark = :19, source_document_id = :20, strat_hierarchy_type = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where parent_strat_name_set_id = :29")
	if err != nil {
		return err
	}

	strat_hierarchy.Row_changed_date = formatDate(strat_hierarchy.Row_changed_date)
	strat_hierarchy.Effective_date = formatDateString(strat_hierarchy.Effective_date)
	strat_hierarchy.Expiry_date = formatDateString(strat_hierarchy.Expiry_date)
	strat_hierarchy.Row_effective_date = formatDateString(strat_hierarchy.Row_effective_date)
	strat_hierarchy.Row_expiry_date = formatDateString(strat_hierarchy.Row_expiry_date)






	rows, err := stmt.Exec(strat_hierarchy.Parent_strat_unit_id, strat_hierarchy.Child_strat_name_set_id, strat_hierarchy.Child_strat_unit_id, strat_hierarchy.Hierarchy_id, strat_hierarchy.Source, strat_hierarchy.Active_ind, strat_hierarchy.Area_id, strat_hierarchy.Area_type, strat_hierarchy.Child_interp_id2, strat_hierarchy.Child_strat_column_id, strat_hierarchy.Child_strat_column_source, strat_hierarchy.Effective_date, strat_hierarchy.Expiry_date, strat_hierarchy.Parent_interp_id, strat_hierarchy.Parent_strat_column_id, strat_hierarchy.Parent_strat_column_source, strat_hierarchy.Ppdm_guid, strat_hierarchy.Preferred_hierarchy_ind, strat_hierarchy.Remark, strat_hierarchy.Source_document_id, strat_hierarchy.Strat_hierarchy_type, strat_hierarchy.Row_changed_by, strat_hierarchy.Row_changed_date, strat_hierarchy.Row_created_by, strat_hierarchy.Row_effective_date, strat_hierarchy.Row_expiry_date, strat_hierarchy.Row_quality, strat_hierarchy.Parent_strat_name_set_id)
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

func PatchStratHierarchy(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_hierarchy set "
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
	query += " where parent_strat_name_set_id = :id"

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

func DeleteStratHierarchy(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_hierarchy dto.Strat_hierarchy
	strat_hierarchy.Parent_strat_name_set_id = id

	stmt, err := db.Prepare("delete from strat_hierarchy where parent_strat_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_hierarchy.Parent_strat_name_set_id)
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


