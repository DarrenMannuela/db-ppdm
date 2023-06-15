package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContAccountProc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_account_proc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_account_proc

	for rows.Next() {
		var cont_account_proc dto.Cont_account_proc
		if err := rows.Scan(&cont_account_proc.Contract_id, &cont_account_proc.Account_proc_type, &cont_account_proc.Account_proc_obs_no, &cont_account_proc.Active_ind, &cont_account_proc.Advance_percent, &cont_account_proc.Effective_date, &cont_account_proc.Expiry_date, &cont_account_proc.Inventory_period, &cont_account_proc.Inventory_period_uom, &cont_account_proc.Material_sale_limit, &cont_account_proc.Ppdm_guid, &cont_account_proc.Quorum_count, &cont_account_proc.Quorum_percent, &cont_account_proc.Remark, &cont_account_proc.Source, &cont_account_proc.Source_document_id, &cont_account_proc.Row_changed_by, &cont_account_proc.Row_changed_date, &cont_account_proc.Row_created_by, &cont_account_proc.Row_created_date, &cont_account_proc.Row_effective_date, &cont_account_proc.Row_expiry_date, &cont_account_proc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_account_proc to result
		result = append(result, cont_account_proc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContAccountProc(c *fiber.Ctx) error {
	var cont_account_proc dto.Cont_account_proc

	setDefaults(&cont_account_proc)

	if err := json.Unmarshal(c.Body(), &cont_account_proc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_account_proc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	cont_account_proc.Row_created_date = formatDate(cont_account_proc.Row_created_date)
	cont_account_proc.Row_changed_date = formatDate(cont_account_proc.Row_changed_date)
	cont_account_proc.Effective_date = formatDateString(cont_account_proc.Effective_date)
	cont_account_proc.Expiry_date = formatDateString(cont_account_proc.Expiry_date)
	cont_account_proc.Row_effective_date = formatDateString(cont_account_proc.Row_effective_date)
	cont_account_proc.Row_expiry_date = formatDateString(cont_account_proc.Row_expiry_date)






	rows, err := stmt.Exec(cont_account_proc.Contract_id, cont_account_proc.Account_proc_type, cont_account_proc.Account_proc_obs_no, cont_account_proc.Active_ind, cont_account_proc.Advance_percent, cont_account_proc.Effective_date, cont_account_proc.Expiry_date, cont_account_proc.Inventory_period, cont_account_proc.Inventory_period_uom, cont_account_proc.Material_sale_limit, cont_account_proc.Ppdm_guid, cont_account_proc.Quorum_count, cont_account_proc.Quorum_percent, cont_account_proc.Remark, cont_account_proc.Source, cont_account_proc.Source_document_id, cont_account_proc.Row_changed_by, cont_account_proc.Row_changed_date, cont_account_proc.Row_created_by, cont_account_proc.Row_created_date, cont_account_proc.Row_effective_date, cont_account_proc.Row_expiry_date, cont_account_proc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContAccountProc(c *fiber.Ctx) error {
	var cont_account_proc dto.Cont_account_proc

	setDefaults(&cont_account_proc)

	if err := json.Unmarshal(c.Body(), &cont_account_proc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_account_proc.Contract_id = id

    if cont_account_proc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_account_proc set account_proc_type = :1, account_proc_obs_no = :2, active_ind = :3, advance_percent = :4, effective_date = :5, expiry_date = :6, inventory_period = :7, inventory_period_uom = :8, material_sale_limit = :9, ppdm_guid = :10, quorum_count = :11, quorum_percent = :12, remark = :13, source = :14, source_document_id = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where contract_id = :23")
	if err != nil {
		return err
	}

	cont_account_proc.Row_changed_date = formatDate(cont_account_proc.Row_changed_date)
	cont_account_proc.Effective_date = formatDateString(cont_account_proc.Effective_date)
	cont_account_proc.Expiry_date = formatDateString(cont_account_proc.Expiry_date)
	cont_account_proc.Row_effective_date = formatDateString(cont_account_proc.Row_effective_date)
	cont_account_proc.Row_expiry_date = formatDateString(cont_account_proc.Row_expiry_date)






	rows, err := stmt.Exec(cont_account_proc.Account_proc_type, cont_account_proc.Account_proc_obs_no, cont_account_proc.Active_ind, cont_account_proc.Advance_percent, cont_account_proc.Effective_date, cont_account_proc.Expiry_date, cont_account_proc.Inventory_period, cont_account_proc.Inventory_period_uom, cont_account_proc.Material_sale_limit, cont_account_proc.Ppdm_guid, cont_account_proc.Quorum_count, cont_account_proc.Quorum_percent, cont_account_proc.Remark, cont_account_proc.Source, cont_account_proc.Source_document_id, cont_account_proc.Row_changed_by, cont_account_proc.Row_changed_date, cont_account_proc.Row_created_by, cont_account_proc.Row_effective_date, cont_account_proc.Row_expiry_date, cont_account_proc.Row_quality, cont_account_proc.Contract_id)
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

func PatchContAccountProc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_account_proc set "
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

func DeleteContAccountProc(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_account_proc dto.Cont_account_proc
	cont_account_proc.Contract_id = id

	stmt, err := db.Prepare("delete from cont_account_proc where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_account_proc.Contract_id)
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


