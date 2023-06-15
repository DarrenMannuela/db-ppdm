package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEntGroup(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ent_group")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ent_group

	for rows.Next() {
		var ent_group dto.Ent_group
		if err := rows.Scan(&ent_group.Entitlement_id, &ent_group.Security_group_id, &ent_group.Group_obs_no, &ent_group.Access_condition, &ent_group.Access_type, &ent_group.Active_ind, &ent_group.Effective_date, &ent_group.Expiry_date, &ent_group.Ppdm_guid, &ent_group.Remark, &ent_group.Restriction_desc, &ent_group.Security_desc, &ent_group.Source, &ent_group.Use_desc, &ent_group.Row_changed_by, &ent_group.Row_changed_date, &ent_group.Row_created_by, &ent_group.Row_created_date, &ent_group.Row_effective_date, &ent_group.Row_expiry_date, &ent_group.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ent_group to result
		result = append(result, ent_group)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEntGroup(c *fiber.Ctx) error {
	var ent_group dto.Ent_group

	setDefaults(&ent_group)

	if err := json.Unmarshal(c.Body(), &ent_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ent_group values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	ent_group.Row_created_date = formatDate(ent_group.Row_created_date)
	ent_group.Row_changed_date = formatDate(ent_group.Row_changed_date)
	ent_group.Effective_date = formatDateString(ent_group.Effective_date)
	ent_group.Expiry_date = formatDateString(ent_group.Expiry_date)
	ent_group.Row_effective_date = formatDateString(ent_group.Row_effective_date)
	ent_group.Row_expiry_date = formatDateString(ent_group.Row_expiry_date)






	rows, err := stmt.Exec(ent_group.Entitlement_id, ent_group.Security_group_id, ent_group.Group_obs_no, ent_group.Access_condition, ent_group.Access_type, ent_group.Active_ind, ent_group.Effective_date, ent_group.Expiry_date, ent_group.Ppdm_guid, ent_group.Remark, ent_group.Restriction_desc, ent_group.Security_desc, ent_group.Source, ent_group.Use_desc, ent_group.Row_changed_by, ent_group.Row_changed_date, ent_group.Row_created_by, ent_group.Row_created_date, ent_group.Row_effective_date, ent_group.Row_expiry_date, ent_group.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEntGroup(c *fiber.Ctx) error {
	var ent_group dto.Ent_group

	setDefaults(&ent_group)

	if err := json.Unmarshal(c.Body(), &ent_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ent_group.Entitlement_id = id

    if ent_group.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ent_group set security_group_id = :1, group_obs_no = :2, access_condition = :3, access_type = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, restriction_desc = :10, security_desc = :11, source = :12, use_desc = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where entitlement_id = :21")
	if err != nil {
		return err
	}

	ent_group.Row_changed_date = formatDate(ent_group.Row_changed_date)
	ent_group.Effective_date = formatDateString(ent_group.Effective_date)
	ent_group.Expiry_date = formatDateString(ent_group.Expiry_date)
	ent_group.Row_effective_date = formatDateString(ent_group.Row_effective_date)
	ent_group.Row_expiry_date = formatDateString(ent_group.Row_expiry_date)






	rows, err := stmt.Exec(ent_group.Security_group_id, ent_group.Group_obs_no, ent_group.Access_condition, ent_group.Access_type, ent_group.Active_ind, ent_group.Effective_date, ent_group.Expiry_date, ent_group.Ppdm_guid, ent_group.Remark, ent_group.Restriction_desc, ent_group.Security_desc, ent_group.Source, ent_group.Use_desc, ent_group.Row_changed_by, ent_group.Row_changed_date, ent_group.Row_created_by, ent_group.Row_effective_date, ent_group.Row_expiry_date, ent_group.Row_quality, ent_group.Entitlement_id)
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

func PatchEntGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ent_group set "
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
	query += " where entitlement_id = :id"

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

func DeleteEntGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	var ent_group dto.Ent_group
	ent_group.Entitlement_id = id

	stmt, err := db.Prepare("delete from ent_group where entitlement_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ent_group.Entitlement_id)
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


