package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEntSecurityGroupXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ent_security_group_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ent_security_group_xref

	for rows.Next() {
		var ent_security_group_xref dto.Ent_security_group_xref
		if err := rows.Scan(&ent_security_group_xref.Security_group_id1, &ent_security_group_xref.Security_group_id2, &ent_security_group_xref.Xref_obs_no, &ent_security_group_xref.Active_ind, &ent_security_group_xref.Effective_date, &ent_security_group_xref.Expiry_date, &ent_security_group_xref.Ppdm_guid, &ent_security_group_xref.Remark, &ent_security_group_xref.Source, &ent_security_group_xref.Xref_type, &ent_security_group_xref.Row_changed_by, &ent_security_group_xref.Row_changed_date, &ent_security_group_xref.Row_created_by, &ent_security_group_xref.Row_created_date, &ent_security_group_xref.Row_effective_date, &ent_security_group_xref.Row_expiry_date, &ent_security_group_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ent_security_group_xref to result
		result = append(result, ent_security_group_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEntSecurityGroupXref(c *fiber.Ctx) error {
	var ent_security_group_xref dto.Ent_security_group_xref

	setDefaults(&ent_security_group_xref)

	if err := json.Unmarshal(c.Body(), &ent_security_group_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ent_security_group_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	ent_security_group_xref.Row_created_date = formatDate(ent_security_group_xref.Row_created_date)
	ent_security_group_xref.Row_changed_date = formatDate(ent_security_group_xref.Row_changed_date)
	ent_security_group_xref.Effective_date = formatDateString(ent_security_group_xref.Effective_date)
	ent_security_group_xref.Expiry_date = formatDateString(ent_security_group_xref.Expiry_date)
	ent_security_group_xref.Row_effective_date = formatDateString(ent_security_group_xref.Row_effective_date)
	ent_security_group_xref.Row_expiry_date = formatDateString(ent_security_group_xref.Row_expiry_date)






	rows, err := stmt.Exec(ent_security_group_xref.Security_group_id1, ent_security_group_xref.Security_group_id2, ent_security_group_xref.Xref_obs_no, ent_security_group_xref.Active_ind, ent_security_group_xref.Effective_date, ent_security_group_xref.Expiry_date, ent_security_group_xref.Ppdm_guid, ent_security_group_xref.Remark, ent_security_group_xref.Source, ent_security_group_xref.Xref_type, ent_security_group_xref.Row_changed_by, ent_security_group_xref.Row_changed_date, ent_security_group_xref.Row_created_by, ent_security_group_xref.Row_created_date, ent_security_group_xref.Row_effective_date, ent_security_group_xref.Row_expiry_date, ent_security_group_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEntSecurityGroupXref(c *fiber.Ctx) error {
	var ent_security_group_xref dto.Ent_security_group_xref

	setDefaults(&ent_security_group_xref)

	if err := json.Unmarshal(c.Body(), &ent_security_group_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ent_security_group_xref.Security_group_id1 = id

    if ent_security_group_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ent_security_group_xref set security_group_id2 = :1, xref_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, xref_type = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where security_group_id1 = :17")
	if err != nil {
		return err
	}

	ent_security_group_xref.Row_changed_date = formatDate(ent_security_group_xref.Row_changed_date)
	ent_security_group_xref.Effective_date = formatDateString(ent_security_group_xref.Effective_date)
	ent_security_group_xref.Expiry_date = formatDateString(ent_security_group_xref.Expiry_date)
	ent_security_group_xref.Row_effective_date = formatDateString(ent_security_group_xref.Row_effective_date)
	ent_security_group_xref.Row_expiry_date = formatDateString(ent_security_group_xref.Row_expiry_date)






	rows, err := stmt.Exec(ent_security_group_xref.Security_group_id2, ent_security_group_xref.Xref_obs_no, ent_security_group_xref.Active_ind, ent_security_group_xref.Effective_date, ent_security_group_xref.Expiry_date, ent_security_group_xref.Ppdm_guid, ent_security_group_xref.Remark, ent_security_group_xref.Source, ent_security_group_xref.Xref_type, ent_security_group_xref.Row_changed_by, ent_security_group_xref.Row_changed_date, ent_security_group_xref.Row_created_by, ent_security_group_xref.Row_effective_date, ent_security_group_xref.Row_expiry_date, ent_security_group_xref.Row_quality, ent_security_group_xref.Security_group_id1)
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

func PatchEntSecurityGroupXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ent_security_group_xref set "
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
	query += " where security_group_id1 = :id"

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

func DeleteEntSecurityGroupXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var ent_security_group_xref dto.Ent_security_group_xref
	ent_security_group_xref.Security_group_id1 = id

	stmt, err := db.Prepare("delete from ent_security_group_xref where security_group_id1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ent_security_group_xref.Security_group_id1)
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


