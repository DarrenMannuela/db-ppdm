package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRConsultRole(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_consult_role")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_consult_role

	for rows.Next() {
		var r_consult_role dto.R_consult_role
		if err := rows.Scan(&r_consult_role.Consult_role, &r_consult_role.Abbreviation, &r_consult_role.Active_ind, &r_consult_role.Effective_date, &r_consult_role.Expiry_date, &r_consult_role.Long_name, &r_consult_role.Ppdm_guid, &r_consult_role.Remark, &r_consult_role.Short_name, &r_consult_role.Source, &r_consult_role.Row_changed_by, &r_consult_role.Row_changed_date, &r_consult_role.Row_created_by, &r_consult_role.Row_created_date, &r_consult_role.Row_effective_date, &r_consult_role.Row_expiry_date, &r_consult_role.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_consult_role to result
		result = append(result, r_consult_role)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRConsultRole(c *fiber.Ctx) error {
	var r_consult_role dto.R_consult_role

	setDefaults(&r_consult_role)

	if err := json.Unmarshal(c.Body(), &r_consult_role); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_consult_role values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	r_consult_role.Row_created_date = formatDate(r_consult_role.Row_created_date)
	r_consult_role.Row_changed_date = formatDate(r_consult_role.Row_changed_date)
	r_consult_role.Effective_date = formatDateString(r_consult_role.Effective_date)
	r_consult_role.Expiry_date = formatDateString(r_consult_role.Expiry_date)
	r_consult_role.Row_effective_date = formatDateString(r_consult_role.Row_effective_date)
	r_consult_role.Row_expiry_date = formatDateString(r_consult_role.Row_expiry_date)





	rows, err := stmt.Exec(r_consult_role.Consult_role, r_consult_role.Abbreviation, r_consult_role.Active_ind, r_consult_role.Effective_date, r_consult_role.Expiry_date, r_consult_role.Long_name, r_consult_role.Ppdm_guid, r_consult_role.Remark, r_consult_role.Short_name, r_consult_role.Source, r_consult_role.Row_changed_by, r_consult_role.Row_changed_date, r_consult_role.Row_created_by, r_consult_role.Row_created_date, r_consult_role.Row_effective_date, r_consult_role.Row_expiry_date, r_consult_role.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRConsultRole(c *fiber.Ctx) error {
	var r_consult_role dto.R_consult_role

	setDefaults(&r_consult_role)

	if err := json.Unmarshal(c.Body(), &r_consult_role); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_consult_role.Consult_role = id

    if r_consult_role.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_consult_role set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where consult_role = :17")
	if err != nil {
		return err
	}

	r_consult_role.Row_changed_date = formatDate(r_consult_role.Row_changed_date)
	r_consult_role.Effective_date = formatDateString(r_consult_role.Effective_date)
	r_consult_role.Expiry_date = formatDateString(r_consult_role.Expiry_date)
	r_consult_role.Row_effective_date = formatDateString(r_consult_role.Row_effective_date)
	r_consult_role.Row_expiry_date = formatDateString(r_consult_role.Row_expiry_date)





	rows, err := stmt.Exec(r_consult_role.Abbreviation, r_consult_role.Active_ind, r_consult_role.Effective_date, r_consult_role.Expiry_date, r_consult_role.Long_name, r_consult_role.Ppdm_guid, r_consult_role.Remark, r_consult_role.Short_name, r_consult_role.Source, r_consult_role.Row_changed_by, r_consult_role.Row_changed_date, r_consult_role.Row_created_by, r_consult_role.Row_effective_date, r_consult_role.Row_expiry_date, r_consult_role.Row_quality, r_consult_role.Consult_role)
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

func PatchRConsultRole(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_consult_role set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"     {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where consult_role = :id"

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

func DeleteRConsultRole(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_consult_role dto.R_consult_role
	r_consult_role.Consult_role = id

	stmt, err := db.Prepare("delete from r_consult_role where consult_role = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_consult_role.Consult_role)
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


