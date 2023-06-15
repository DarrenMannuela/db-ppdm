package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWorkOrderCondition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from work_order_condition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Work_order_condition

	for rows.Next() {
		var work_order_condition dto.Work_order_condition
		if err := rows.Scan(&work_order_condition.Work_order_id, &work_order_condition.Condition_obs_no, &work_order_condition.Active_ind, &work_order_condition.Ba_role, &work_order_condition.Business_associate_id, &work_order_condition.Condition_desc, &work_order_condition.Condition_type, &work_order_condition.Currency_conversion, &work_order_condition.Effective_date, &work_order_condition.Expiry_date, &work_order_condition.Payment_amount, &work_order_condition.Payment_amount_ouom, &work_order_condition.Payment_percent, &work_order_condition.Ppdm_guid, &work_order_condition.Remark, &work_order_condition.Source, &work_order_condition.Work_order_type, &work_order_condition.Row_changed_by, &work_order_condition.Row_changed_date, &work_order_condition.Row_created_by, &work_order_condition.Row_created_date, &work_order_condition.Row_effective_date, &work_order_condition.Row_expiry_date, &work_order_condition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append work_order_condition to result
		result = append(result, work_order_condition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWorkOrderCondition(c *fiber.Ctx) error {
	var work_order_condition dto.Work_order_condition

	setDefaults(&work_order_condition)

	if err := json.Unmarshal(c.Body(), &work_order_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into work_order_condition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	work_order_condition.Row_created_date = formatDate(work_order_condition.Row_created_date)
	work_order_condition.Row_changed_date = formatDate(work_order_condition.Row_changed_date)
	work_order_condition.Effective_date = formatDateString(work_order_condition.Effective_date)
	work_order_condition.Expiry_date = formatDateString(work_order_condition.Expiry_date)
	work_order_condition.Row_effective_date = formatDateString(work_order_condition.Row_effective_date)
	work_order_condition.Row_expiry_date = formatDateString(work_order_condition.Row_expiry_date)






	rows, err := stmt.Exec(work_order_condition.Work_order_id, work_order_condition.Condition_obs_no, work_order_condition.Active_ind, work_order_condition.Ba_role, work_order_condition.Business_associate_id, work_order_condition.Condition_desc, work_order_condition.Condition_type, work_order_condition.Currency_conversion, work_order_condition.Effective_date, work_order_condition.Expiry_date, work_order_condition.Payment_amount, work_order_condition.Payment_amount_ouom, work_order_condition.Payment_percent, work_order_condition.Ppdm_guid, work_order_condition.Remark, work_order_condition.Source, work_order_condition.Work_order_type, work_order_condition.Row_changed_by, work_order_condition.Row_changed_date, work_order_condition.Row_created_by, work_order_condition.Row_created_date, work_order_condition.Row_effective_date, work_order_condition.Row_expiry_date, work_order_condition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWorkOrderCondition(c *fiber.Ctx) error {
	var work_order_condition dto.Work_order_condition

	setDefaults(&work_order_condition)

	if err := json.Unmarshal(c.Body(), &work_order_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	work_order_condition.Work_order_id = id

    if work_order_condition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update work_order_condition set condition_obs_no = :1, active_ind = :2, ba_role = :3, business_associate_id = :4, condition_desc = :5, condition_type = :6, currency_conversion = :7, effective_date = :8, expiry_date = :9, payment_amount = :10, payment_amount_ouom = :11, payment_percent = :12, ppdm_guid = :13, remark = :14, source = :15, work_order_type = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where work_order_id = :24")
	if err != nil {
		return err
	}

	work_order_condition.Row_changed_date = formatDate(work_order_condition.Row_changed_date)
	work_order_condition.Effective_date = formatDateString(work_order_condition.Effective_date)
	work_order_condition.Expiry_date = formatDateString(work_order_condition.Expiry_date)
	work_order_condition.Row_effective_date = formatDateString(work_order_condition.Row_effective_date)
	work_order_condition.Row_expiry_date = formatDateString(work_order_condition.Row_expiry_date)






	rows, err := stmt.Exec(work_order_condition.Condition_obs_no, work_order_condition.Active_ind, work_order_condition.Ba_role, work_order_condition.Business_associate_id, work_order_condition.Condition_desc, work_order_condition.Condition_type, work_order_condition.Currency_conversion, work_order_condition.Effective_date, work_order_condition.Expiry_date, work_order_condition.Payment_amount, work_order_condition.Payment_amount_ouom, work_order_condition.Payment_percent, work_order_condition.Ppdm_guid, work_order_condition.Remark, work_order_condition.Source, work_order_condition.Work_order_type, work_order_condition.Row_changed_by, work_order_condition.Row_changed_date, work_order_condition.Row_created_by, work_order_condition.Row_effective_date, work_order_condition.Row_expiry_date, work_order_condition.Row_quality, work_order_condition.Work_order_id)
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

func PatchWorkOrderCondition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update work_order_condition set "
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
	query += " where work_order_id = :id"

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

func DeleteWorkOrderCondition(c *fiber.Ctx) error {
	id := c.Params("id")
	var work_order_condition dto.Work_order_condition
	work_order_condition.Work_order_id = id

	stmt, err := db.Prepare("delete from work_order_condition where work_order_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(work_order_condition.Work_order_id)
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


