package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmSchemaGroup(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_schema_group")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_schema_group

	for rows.Next() {
		var ppdm_schema_group dto.Ppdm_schema_group
		if err := rows.Scan(&ppdm_schema_group.Group_system_id, &ppdm_schema_group.Group_schema_entity_id, &ppdm_schema_group.Comp_schema_entity_id, &ppdm_schema_group.Active_ind, &ppdm_schema_group.Effective_date, &ppdm_schema_group.Entity_seq_no, &ppdm_schema_group.Expiry_date, &ppdm_schema_group.Extension_ind, &ppdm_schema_group.Ppdm_guid, &ppdm_schema_group.Remark, &ppdm_schema_group.Schema_group_type, &ppdm_schema_group.Source, &ppdm_schema_group.Surrogate_ind, &ppdm_schema_group.Row_changed_by, &ppdm_schema_group.Row_changed_date, &ppdm_schema_group.Row_created_by, &ppdm_schema_group.Row_created_date, &ppdm_schema_group.Row_effective_date, &ppdm_schema_group.Row_expiry_date, &ppdm_schema_group.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_schema_group to result
		result = append(result, ppdm_schema_group)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmSchemaGroup(c *fiber.Ctx) error {
	var ppdm_schema_group dto.Ppdm_schema_group

	setDefaults(&ppdm_schema_group)

	if err := json.Unmarshal(c.Body(), &ppdm_schema_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_schema_group values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	ppdm_schema_group.Row_created_date = formatDate(ppdm_schema_group.Row_created_date)
	ppdm_schema_group.Row_changed_date = formatDate(ppdm_schema_group.Row_changed_date)
	ppdm_schema_group.Effective_date = formatDateString(ppdm_schema_group.Effective_date)
	ppdm_schema_group.Expiry_date = formatDateString(ppdm_schema_group.Expiry_date)
	ppdm_schema_group.Row_effective_date = formatDateString(ppdm_schema_group.Row_effective_date)
	ppdm_schema_group.Row_expiry_date = formatDateString(ppdm_schema_group.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_schema_group.Group_system_id, ppdm_schema_group.Group_schema_entity_id, ppdm_schema_group.Comp_schema_entity_id, ppdm_schema_group.Active_ind, ppdm_schema_group.Effective_date, ppdm_schema_group.Entity_seq_no, ppdm_schema_group.Expiry_date, ppdm_schema_group.Extension_ind, ppdm_schema_group.Ppdm_guid, ppdm_schema_group.Remark, ppdm_schema_group.Schema_group_type, ppdm_schema_group.Source, ppdm_schema_group.Surrogate_ind, ppdm_schema_group.Row_changed_by, ppdm_schema_group.Row_changed_date, ppdm_schema_group.Row_created_by, ppdm_schema_group.Row_created_date, ppdm_schema_group.Row_effective_date, ppdm_schema_group.Row_expiry_date, ppdm_schema_group.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmSchemaGroup(c *fiber.Ctx) error {
	var ppdm_schema_group dto.Ppdm_schema_group

	setDefaults(&ppdm_schema_group)

	if err := json.Unmarshal(c.Body(), &ppdm_schema_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_schema_group.Group_system_id = id

    if ppdm_schema_group.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_schema_group set group_schema_entity_id = :1, comp_schema_entity_id = :2, active_ind = :3, effective_date = :4, entity_seq_no = :5, expiry_date = :6, extension_ind = :7, ppdm_guid = :8, remark = :9, schema_group_type = :10, source = :11, surrogate_ind = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where group_system_id = :20")
	if err != nil {
		return err
	}

	ppdm_schema_group.Row_changed_date = formatDate(ppdm_schema_group.Row_changed_date)
	ppdm_schema_group.Effective_date = formatDateString(ppdm_schema_group.Effective_date)
	ppdm_schema_group.Expiry_date = formatDateString(ppdm_schema_group.Expiry_date)
	ppdm_schema_group.Row_effective_date = formatDateString(ppdm_schema_group.Row_effective_date)
	ppdm_schema_group.Row_expiry_date = formatDateString(ppdm_schema_group.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_schema_group.Group_schema_entity_id, ppdm_schema_group.Comp_schema_entity_id, ppdm_schema_group.Active_ind, ppdm_schema_group.Effective_date, ppdm_schema_group.Entity_seq_no, ppdm_schema_group.Expiry_date, ppdm_schema_group.Extension_ind, ppdm_schema_group.Ppdm_guid, ppdm_schema_group.Remark, ppdm_schema_group.Schema_group_type, ppdm_schema_group.Source, ppdm_schema_group.Surrogate_ind, ppdm_schema_group.Row_changed_by, ppdm_schema_group.Row_changed_date, ppdm_schema_group.Row_created_by, ppdm_schema_group.Row_effective_date, ppdm_schema_group.Row_expiry_date, ppdm_schema_group.Row_quality, ppdm_schema_group.Group_system_id)
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

func PatchPpdmSchemaGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_schema_group set "
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
	query += " where group_system_id = :id"

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

func DeletePpdmSchemaGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_schema_group dto.Ppdm_schema_group
	ppdm_schema_group.Group_system_id = id

	stmt, err := db.Prepare("delete from ppdm_schema_group where group_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_schema_group.Group_system_id)
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


