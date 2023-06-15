package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEntSecurityBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ent_security_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ent_security_ba

	for rows.Next() {
		var ent_security_ba dto.Ent_security_ba
		if err := rows.Scan(&ent_security_ba.Security_group_id, &ent_security_ba.Business_associate_id, &ent_security_ba.Ba_obs_no, &ent_security_ba.Active_ind, &ent_security_ba.Effective_date, &ent_security_ba.Expiry_date, &ent_security_ba.Ppdm_guid, &ent_security_ba.Remark, &ent_security_ba.Source, &ent_security_ba.Row_changed_by, &ent_security_ba.Row_changed_date, &ent_security_ba.Row_created_by, &ent_security_ba.Row_created_date, &ent_security_ba.Row_effective_date, &ent_security_ba.Row_expiry_date, &ent_security_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ent_security_ba to result
		result = append(result, ent_security_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEntSecurityBa(c *fiber.Ctx) error {
	var ent_security_ba dto.Ent_security_ba

	setDefaults(&ent_security_ba)

	if err := json.Unmarshal(c.Body(), &ent_security_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ent_security_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	ent_security_ba.Row_created_date = formatDate(ent_security_ba.Row_created_date)
	ent_security_ba.Row_changed_date = formatDate(ent_security_ba.Row_changed_date)
	ent_security_ba.Effective_date = formatDateString(ent_security_ba.Effective_date)
	ent_security_ba.Expiry_date = formatDateString(ent_security_ba.Expiry_date)
	ent_security_ba.Row_effective_date = formatDateString(ent_security_ba.Row_effective_date)
	ent_security_ba.Row_expiry_date = formatDateString(ent_security_ba.Row_expiry_date)






	rows, err := stmt.Exec(ent_security_ba.Security_group_id, ent_security_ba.Business_associate_id, ent_security_ba.Ba_obs_no, ent_security_ba.Active_ind, ent_security_ba.Effective_date, ent_security_ba.Expiry_date, ent_security_ba.Ppdm_guid, ent_security_ba.Remark, ent_security_ba.Source, ent_security_ba.Row_changed_by, ent_security_ba.Row_changed_date, ent_security_ba.Row_created_by, ent_security_ba.Row_created_date, ent_security_ba.Row_effective_date, ent_security_ba.Row_expiry_date, ent_security_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEntSecurityBa(c *fiber.Ctx) error {
	var ent_security_ba dto.Ent_security_ba

	setDefaults(&ent_security_ba)

	if err := json.Unmarshal(c.Body(), &ent_security_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ent_security_ba.Security_group_id = id

    if ent_security_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ent_security_ba set business_associate_id = :1, ba_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where security_group_id = :16")
	if err != nil {
		return err
	}

	ent_security_ba.Row_changed_date = formatDate(ent_security_ba.Row_changed_date)
	ent_security_ba.Effective_date = formatDateString(ent_security_ba.Effective_date)
	ent_security_ba.Expiry_date = formatDateString(ent_security_ba.Expiry_date)
	ent_security_ba.Row_effective_date = formatDateString(ent_security_ba.Row_effective_date)
	ent_security_ba.Row_expiry_date = formatDateString(ent_security_ba.Row_expiry_date)






	rows, err := stmt.Exec(ent_security_ba.Business_associate_id, ent_security_ba.Ba_obs_no, ent_security_ba.Active_ind, ent_security_ba.Effective_date, ent_security_ba.Expiry_date, ent_security_ba.Ppdm_guid, ent_security_ba.Remark, ent_security_ba.Source, ent_security_ba.Row_changed_by, ent_security_ba.Row_changed_date, ent_security_ba.Row_created_by, ent_security_ba.Row_effective_date, ent_security_ba.Row_expiry_date, ent_security_ba.Row_quality, ent_security_ba.Security_group_id)
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

func PatchEntSecurityBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ent_security_ba set "
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
	query += " where security_group_id = :id"

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

func DeleteEntSecurityBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var ent_security_ba dto.Ent_security_ba
	ent_security_ba.Security_group_id = id

	stmt, err := db.Prepare("delete from ent_security_ba where security_group_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ent_security_ba.Security_group_id)
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


