package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWorkOrderBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from work_order_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Work_order_ba

	for rows.Next() {
		var work_order_ba dto.Work_order_ba
		if err := rows.Scan(&work_order_ba.Work_order_id, &work_order_ba.Business_associate_id, &work_order_ba.Ba_obs_no, &work_order_ba.Active_ind, &work_order_ba.Ba_role, &work_order_ba.Currency_conversion, &work_order_ba.Effective_date, &work_order_ba.Expiry_date, &work_order_ba.Payment_amount, &work_order_ba.Payment_amount_ouom, &work_order_ba.Payment_percent, &work_order_ba.Ppdm_guid, &work_order_ba.Remark, &work_order_ba.Source, &work_order_ba.Row_changed_by, &work_order_ba.Row_changed_date, &work_order_ba.Row_created_by, &work_order_ba.Row_created_date, &work_order_ba.Row_effective_date, &work_order_ba.Row_expiry_date, &work_order_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append work_order_ba to result
		result = append(result, work_order_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWorkOrderBa(c *fiber.Ctx) error {
	var work_order_ba dto.Work_order_ba

	setDefaults(&work_order_ba)

	if err := json.Unmarshal(c.Body(), &work_order_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into work_order_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	work_order_ba.Row_created_date = formatDate(work_order_ba.Row_created_date)
	work_order_ba.Row_changed_date = formatDate(work_order_ba.Row_changed_date)
	work_order_ba.Effective_date = formatDateString(work_order_ba.Effective_date)
	work_order_ba.Expiry_date = formatDateString(work_order_ba.Expiry_date)
	work_order_ba.Row_effective_date = formatDateString(work_order_ba.Row_effective_date)
	work_order_ba.Row_expiry_date = formatDateString(work_order_ba.Row_expiry_date)






	rows, err := stmt.Exec(work_order_ba.Work_order_id, work_order_ba.Business_associate_id, work_order_ba.Ba_obs_no, work_order_ba.Active_ind, work_order_ba.Ba_role, work_order_ba.Currency_conversion, work_order_ba.Effective_date, work_order_ba.Expiry_date, work_order_ba.Payment_amount, work_order_ba.Payment_amount_ouom, work_order_ba.Payment_percent, work_order_ba.Ppdm_guid, work_order_ba.Remark, work_order_ba.Source, work_order_ba.Row_changed_by, work_order_ba.Row_changed_date, work_order_ba.Row_created_by, work_order_ba.Row_created_date, work_order_ba.Row_effective_date, work_order_ba.Row_expiry_date, work_order_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWorkOrderBa(c *fiber.Ctx) error {
	var work_order_ba dto.Work_order_ba

	setDefaults(&work_order_ba)

	if err := json.Unmarshal(c.Body(), &work_order_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	work_order_ba.Work_order_id = id

    if work_order_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update work_order_ba set business_associate_id = :1, ba_obs_no = :2, active_ind = :3, ba_role = :4, currency_conversion = :5, effective_date = :6, expiry_date = :7, payment_amount = :8, payment_amount_ouom = :9, payment_percent = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where work_order_id = :21")
	if err != nil {
		return err
	}

	work_order_ba.Row_changed_date = formatDate(work_order_ba.Row_changed_date)
	work_order_ba.Effective_date = formatDateString(work_order_ba.Effective_date)
	work_order_ba.Expiry_date = formatDateString(work_order_ba.Expiry_date)
	work_order_ba.Row_effective_date = formatDateString(work_order_ba.Row_effective_date)
	work_order_ba.Row_expiry_date = formatDateString(work_order_ba.Row_expiry_date)






	rows, err := stmt.Exec(work_order_ba.Business_associate_id, work_order_ba.Ba_obs_no, work_order_ba.Active_ind, work_order_ba.Ba_role, work_order_ba.Currency_conversion, work_order_ba.Effective_date, work_order_ba.Expiry_date, work_order_ba.Payment_amount, work_order_ba.Payment_amount_ouom, work_order_ba.Payment_percent, work_order_ba.Ppdm_guid, work_order_ba.Remark, work_order_ba.Source, work_order_ba.Row_changed_by, work_order_ba.Row_changed_date, work_order_ba.Row_created_by, work_order_ba.Row_effective_date, work_order_ba.Row_expiry_date, work_order_ba.Row_quality, work_order_ba.Work_order_id)
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

func PatchWorkOrderBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update work_order_ba set "
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

func DeleteWorkOrderBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var work_order_ba dto.Work_order_ba
	work_order_ba.Work_order_id = id

	stmt, err := db.Prepare("delete from work_order_ba where work_order_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(work_order_ba.Work_order_id)
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


