package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContExemption(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_exemption")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_exemption

	for rows.Next() {
		var cont_exemption dto.Cont_exemption
		if err := rows.Scan(&cont_exemption.Contract_id, &cont_exemption.Business_associate_id, &cont_exemption.Exemption_id, &cont_exemption.Active_ind, &cont_exemption.Effective_date, &cont_exemption.Exemption_desc, &cont_exemption.Expiry_date, &cont_exemption.Ppdm_guid, &cont_exemption.Provision_id, &cont_exemption.Remark, &cont_exemption.Source, &cont_exemption.Row_changed_by, &cont_exemption.Row_changed_date, &cont_exemption.Row_created_by, &cont_exemption.Row_created_date, &cont_exemption.Row_effective_date, &cont_exemption.Row_expiry_date, &cont_exemption.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_exemption to result
		result = append(result, cont_exemption)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContExemption(c *fiber.Ctx) error {
	var cont_exemption dto.Cont_exemption

	setDefaults(&cont_exemption)

	if err := json.Unmarshal(c.Body(), &cont_exemption); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_exemption values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	cont_exemption.Row_created_date = formatDate(cont_exemption.Row_created_date)
	cont_exemption.Row_changed_date = formatDate(cont_exemption.Row_changed_date)
	cont_exemption.Effective_date = formatDateString(cont_exemption.Effective_date)
	cont_exemption.Expiry_date = formatDateString(cont_exemption.Expiry_date)
	cont_exemption.Row_effective_date = formatDateString(cont_exemption.Row_effective_date)
	cont_exemption.Row_expiry_date = formatDateString(cont_exemption.Row_expiry_date)






	rows, err := stmt.Exec(cont_exemption.Contract_id, cont_exemption.Business_associate_id, cont_exemption.Exemption_id, cont_exemption.Active_ind, cont_exemption.Effective_date, cont_exemption.Exemption_desc, cont_exemption.Expiry_date, cont_exemption.Ppdm_guid, cont_exemption.Provision_id, cont_exemption.Remark, cont_exemption.Source, cont_exemption.Row_changed_by, cont_exemption.Row_changed_date, cont_exemption.Row_created_by, cont_exemption.Row_created_date, cont_exemption.Row_effective_date, cont_exemption.Row_expiry_date, cont_exemption.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContExemption(c *fiber.Ctx) error {
	var cont_exemption dto.Cont_exemption

	setDefaults(&cont_exemption)

	if err := json.Unmarshal(c.Body(), &cont_exemption); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_exemption.Contract_id = id

    if cont_exemption.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_exemption set business_associate_id = :1, exemption_id = :2, active_ind = :3, effective_date = :4, exemption_desc = :5, expiry_date = :6, ppdm_guid = :7, provision_id = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where contract_id = :18")
	if err != nil {
		return err
	}

	cont_exemption.Row_changed_date = formatDate(cont_exemption.Row_changed_date)
	cont_exemption.Effective_date = formatDateString(cont_exemption.Effective_date)
	cont_exemption.Expiry_date = formatDateString(cont_exemption.Expiry_date)
	cont_exemption.Row_effective_date = formatDateString(cont_exemption.Row_effective_date)
	cont_exemption.Row_expiry_date = formatDateString(cont_exemption.Row_expiry_date)






	rows, err := stmt.Exec(cont_exemption.Business_associate_id, cont_exemption.Exemption_id, cont_exemption.Active_ind, cont_exemption.Effective_date, cont_exemption.Exemption_desc, cont_exemption.Expiry_date, cont_exemption.Ppdm_guid, cont_exemption.Provision_id, cont_exemption.Remark, cont_exemption.Source, cont_exemption.Row_changed_by, cont_exemption.Row_changed_date, cont_exemption.Row_created_by, cont_exemption.Row_effective_date, cont_exemption.Row_expiry_date, cont_exemption.Row_quality, cont_exemption.Contract_id)
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

func PatchContExemption(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_exemption set "
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
	query += " where contract_id = :id"

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

func DeleteContExemption(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_exemption dto.Cont_exemption
	cont_exemption.Contract_id = id

	stmt, err := db.Prepare("delete from cont_exemption where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_exemption.Contract_id)
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


