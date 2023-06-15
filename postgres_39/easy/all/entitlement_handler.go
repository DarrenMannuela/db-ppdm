package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEntitlement(c *fiber.Ctx) error {
	rows, err := db.Query("select * from entitlement")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Entitlement

	for rows.Next() {
		var entitlement dto.Entitlement
		if err := rows.Scan(&entitlement.Entitlement_id, &entitlement.Access_condition, &entitlement.Access_cond_code, &entitlement.Active_ind, &entitlement.Description, &entitlement.Effective_date, &entitlement.Entitlement_name, &entitlement.Entitlement_type, &entitlement.Expiry_action, &entitlement.Expiry_date, &entitlement.Ppdm_guid, &entitlement.Primary_term, &entitlement.Primary_term_ouom, &entitlement.Remark, &entitlement.Security_desc, &entitlement.Source, &entitlement.Use_condition, &entitlement.Row_changed_by, &entitlement.Row_changed_date, &entitlement.Row_created_by, &entitlement.Row_created_date, &entitlement.Row_effective_date, &entitlement.Row_expiry_date, &entitlement.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append entitlement to result
		result = append(result, entitlement)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEntitlement(c *fiber.Ctx) error {
	var entitlement dto.Entitlement

	setDefaults(&entitlement)

	if err := json.Unmarshal(c.Body(), &entitlement); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into entitlement values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	entitlement.Row_created_date = formatDate(entitlement.Row_created_date)
	entitlement.Row_changed_date = formatDate(entitlement.Row_changed_date)
	entitlement.Effective_date = formatDateString(entitlement.Effective_date)
	entitlement.Expiry_date = formatDateString(entitlement.Expiry_date)
	entitlement.Row_effective_date = formatDateString(entitlement.Row_effective_date)
	entitlement.Row_expiry_date = formatDateString(entitlement.Row_expiry_date)






	rows, err := stmt.Exec(entitlement.Entitlement_id, entitlement.Access_condition, entitlement.Access_cond_code, entitlement.Active_ind, entitlement.Description, entitlement.Effective_date, entitlement.Entitlement_name, entitlement.Entitlement_type, entitlement.Expiry_action, entitlement.Expiry_date, entitlement.Ppdm_guid, entitlement.Primary_term, entitlement.Primary_term_ouom, entitlement.Remark, entitlement.Security_desc, entitlement.Source, entitlement.Use_condition, entitlement.Row_changed_by, entitlement.Row_changed_date, entitlement.Row_created_by, entitlement.Row_created_date, entitlement.Row_effective_date, entitlement.Row_expiry_date, entitlement.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEntitlement(c *fiber.Ctx) error {
	var entitlement dto.Entitlement

	setDefaults(&entitlement)

	if err := json.Unmarshal(c.Body(), &entitlement); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	entitlement.Entitlement_id = id

    if entitlement.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update entitlement set access_condition = :1, access_cond_code = :2, active_ind = :3, description = :4, effective_date = :5, entitlement_name = :6, entitlement_type = :7, expiry_action = :8, expiry_date = :9, ppdm_guid = :10, primary_term = :11, primary_term_ouom = :12, remark = :13, security_desc = :14, source = :15, use_condition = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where entitlement_id = :24")
	if err != nil {
		return err
	}

	entitlement.Row_changed_date = formatDate(entitlement.Row_changed_date)
	entitlement.Effective_date = formatDateString(entitlement.Effective_date)
	entitlement.Expiry_date = formatDateString(entitlement.Expiry_date)
	entitlement.Row_effective_date = formatDateString(entitlement.Row_effective_date)
	entitlement.Row_expiry_date = formatDateString(entitlement.Row_expiry_date)






	rows, err := stmt.Exec(entitlement.Access_condition, entitlement.Access_cond_code, entitlement.Active_ind, entitlement.Description, entitlement.Effective_date, entitlement.Entitlement_name, entitlement.Entitlement_type, entitlement.Expiry_action, entitlement.Expiry_date, entitlement.Ppdm_guid, entitlement.Primary_term, entitlement.Primary_term_ouom, entitlement.Remark, entitlement.Security_desc, entitlement.Source, entitlement.Use_condition, entitlement.Row_changed_by, entitlement.Row_changed_date, entitlement.Row_created_by, entitlement.Row_effective_date, entitlement.Row_expiry_date, entitlement.Row_quality, entitlement.Entitlement_id)
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

func PatchEntitlement(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update entitlement set "
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

func DeleteEntitlement(c *fiber.Ctx) error {
	id := c.Params("id")
	var entitlement dto.Entitlement
	entitlement.Entitlement_id = id

	stmt, err := db.Prepare("delete from entitlement where entitlement_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(entitlement.Entitlement_id)
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


