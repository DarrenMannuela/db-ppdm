package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratHierarchyDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_hierarchy_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_hierarchy_desc

	for rows.Next() {
		var strat_hierarchy_desc dto.Strat_hierarchy_desc
		if err := rows.Scan(&strat_hierarchy_desc.Hierarchy_desc_id, &strat_hierarchy_desc.Hierarchy_seq_no, &strat_hierarchy_desc.Active_ind, &strat_hierarchy_desc.Effective_date, &strat_hierarchy_desc.Expiry_date, &strat_hierarchy_desc.Parent_strat_unit_type, &strat_hierarchy_desc.Ppdm_guid, &strat_hierarchy_desc.Remark, &strat_hierarchy_desc.Source, &strat_hierarchy_desc.Source_document_id, &strat_hierarchy_desc.Strat_hierarchy_type, &strat_hierarchy_desc.Strat_type, &strat_hierarchy_desc.Strat_unit_type, &strat_hierarchy_desc.Row_changed_by, &strat_hierarchy_desc.Row_changed_date, &strat_hierarchy_desc.Row_created_by, &strat_hierarchy_desc.Row_created_date, &strat_hierarchy_desc.Row_effective_date, &strat_hierarchy_desc.Row_expiry_date, &strat_hierarchy_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_hierarchy_desc to result
		result = append(result, strat_hierarchy_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratHierarchyDesc(c *fiber.Ctx) error {
	var strat_hierarchy_desc dto.Strat_hierarchy_desc

	setDefaults(&strat_hierarchy_desc)

	if err := json.Unmarshal(c.Body(), &strat_hierarchy_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_hierarchy_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	strat_hierarchy_desc.Row_created_date = formatDate(strat_hierarchy_desc.Row_created_date)
	strat_hierarchy_desc.Row_changed_date = formatDate(strat_hierarchy_desc.Row_changed_date)
	strat_hierarchy_desc.Effective_date = formatDateString(strat_hierarchy_desc.Effective_date)
	strat_hierarchy_desc.Expiry_date = formatDateString(strat_hierarchy_desc.Expiry_date)
	strat_hierarchy_desc.Row_effective_date = formatDateString(strat_hierarchy_desc.Row_effective_date)
	strat_hierarchy_desc.Row_expiry_date = formatDateString(strat_hierarchy_desc.Row_expiry_date)






	rows, err := stmt.Exec(strat_hierarchy_desc.Hierarchy_desc_id, strat_hierarchy_desc.Hierarchy_seq_no, strat_hierarchy_desc.Active_ind, strat_hierarchy_desc.Effective_date, strat_hierarchy_desc.Expiry_date, strat_hierarchy_desc.Parent_strat_unit_type, strat_hierarchy_desc.Ppdm_guid, strat_hierarchy_desc.Remark, strat_hierarchy_desc.Source, strat_hierarchy_desc.Source_document_id, strat_hierarchy_desc.Strat_hierarchy_type, strat_hierarchy_desc.Strat_type, strat_hierarchy_desc.Strat_unit_type, strat_hierarchy_desc.Row_changed_by, strat_hierarchy_desc.Row_changed_date, strat_hierarchy_desc.Row_created_by, strat_hierarchy_desc.Row_created_date, strat_hierarchy_desc.Row_effective_date, strat_hierarchy_desc.Row_expiry_date, strat_hierarchy_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratHierarchyDesc(c *fiber.Ctx) error {
	var strat_hierarchy_desc dto.Strat_hierarchy_desc

	setDefaults(&strat_hierarchy_desc)

	if err := json.Unmarshal(c.Body(), &strat_hierarchy_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_hierarchy_desc.Hierarchy_desc_id = id

    if strat_hierarchy_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_hierarchy_desc set hierarchy_seq_no = :1, active_ind = :2, effective_date = :3, expiry_date = :4, parent_strat_unit_type = :5, ppdm_guid = :6, remark = :7, source = :8, source_document_id = :9, strat_hierarchy_type = :10, strat_type = :11, strat_unit_type = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where hierarchy_desc_id = :20")
	if err != nil {
		return err
	}

	strat_hierarchy_desc.Row_changed_date = formatDate(strat_hierarchy_desc.Row_changed_date)
	strat_hierarchy_desc.Effective_date = formatDateString(strat_hierarchy_desc.Effective_date)
	strat_hierarchy_desc.Expiry_date = formatDateString(strat_hierarchy_desc.Expiry_date)
	strat_hierarchy_desc.Row_effective_date = formatDateString(strat_hierarchy_desc.Row_effective_date)
	strat_hierarchy_desc.Row_expiry_date = formatDateString(strat_hierarchy_desc.Row_expiry_date)






	rows, err := stmt.Exec(strat_hierarchy_desc.Hierarchy_seq_no, strat_hierarchy_desc.Active_ind, strat_hierarchy_desc.Effective_date, strat_hierarchy_desc.Expiry_date, strat_hierarchy_desc.Parent_strat_unit_type, strat_hierarchy_desc.Ppdm_guid, strat_hierarchy_desc.Remark, strat_hierarchy_desc.Source, strat_hierarchy_desc.Source_document_id, strat_hierarchy_desc.Strat_hierarchy_type, strat_hierarchy_desc.Strat_type, strat_hierarchy_desc.Strat_unit_type, strat_hierarchy_desc.Row_changed_by, strat_hierarchy_desc.Row_changed_date, strat_hierarchy_desc.Row_created_by, strat_hierarchy_desc.Row_effective_date, strat_hierarchy_desc.Row_expiry_date, strat_hierarchy_desc.Row_quality, strat_hierarchy_desc.Hierarchy_desc_id)
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

func PatchStratHierarchyDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_hierarchy_desc set "
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
	query += " where hierarchy_desc_id = :id"

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

func DeleteStratHierarchyDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_hierarchy_desc dto.Strat_hierarchy_desc
	strat_hierarchy_desc.Hierarchy_desc_id = id

	stmt, err := db.Prepare("delete from strat_hierarchy_desc where hierarchy_desc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_hierarchy_desc.Hierarchy_desc_id)
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


