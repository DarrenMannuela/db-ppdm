package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWorkOrder(c *fiber.Ctx) error {
	rows, err := db.Query("select * from work_order")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Work_order

	for rows.Next() {
		var work_order dto.Work_order
		if err := rows.Scan(&work_order.Work_order_id, &work_order.Active_ind, &work_order.Complete_date, &work_order.Due_date, &work_order.Effective_date, &work_order.Expiry_date, &work_order.Final_billing_date, &work_order.Instructions, &work_order.Ppdm_guid, &work_order.Remark, &work_order.Request_date, &work_order.Rush_ind, &work_order.Source, &work_order.Work_order_number, &work_order.Work_order_type, &work_order.Row_changed_by, &work_order.Row_changed_date, &work_order.Row_created_by, &work_order.Row_created_date, &work_order.Row_effective_date, &work_order.Row_expiry_date, &work_order.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append work_order to result
		result = append(result, work_order)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWorkOrder(c *fiber.Ctx) error {
	var work_order dto.Work_order

	setDefaults(&work_order)

	if err := json.Unmarshal(c.Body(), &work_order); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into work_order values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	work_order.Row_created_date = formatDate(work_order.Row_created_date)
	work_order.Row_changed_date = formatDate(work_order.Row_changed_date)
	work_order.Complete_date = formatDateString(work_order.Complete_date)
	work_order.Due_date = formatDateString(work_order.Due_date)
	work_order.Effective_date = formatDateString(work_order.Effective_date)
	work_order.Expiry_date = formatDateString(work_order.Expiry_date)
	work_order.Final_billing_date = formatDateString(work_order.Final_billing_date)
	work_order.Request_date = formatDateString(work_order.Request_date)
	work_order.Row_effective_date = formatDateString(work_order.Row_effective_date)
	work_order.Row_expiry_date = formatDateString(work_order.Row_expiry_date)






	rows, err := stmt.Exec(work_order.Work_order_id, work_order.Active_ind, work_order.Complete_date, work_order.Due_date, work_order.Effective_date, work_order.Expiry_date, work_order.Final_billing_date, work_order.Instructions, work_order.Ppdm_guid, work_order.Remark, work_order.Request_date, work_order.Rush_ind, work_order.Source, work_order.Work_order_number, work_order.Work_order_type, work_order.Row_changed_by, work_order.Row_changed_date, work_order.Row_created_by, work_order.Row_created_date, work_order.Row_effective_date, work_order.Row_expiry_date, work_order.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWorkOrder(c *fiber.Ctx) error {
	var work_order dto.Work_order

	setDefaults(&work_order)

	if err := json.Unmarshal(c.Body(), &work_order); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	work_order.Work_order_id = id

    if work_order.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update work_order set active_ind = :1, complete_date = :2, due_date = :3, effective_date = :4, expiry_date = :5, final_billing_date = :6, instructions = :7, ppdm_guid = :8, remark = :9, request_date = :10, rush_ind = :11, source = :12, work_order_number = :13, work_order_type = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where work_order_id = :22")
	if err != nil {
		return err
	}

	work_order.Row_changed_date = formatDate(work_order.Row_changed_date)
	work_order.Complete_date = formatDateString(work_order.Complete_date)
	work_order.Due_date = formatDateString(work_order.Due_date)
	work_order.Effective_date = formatDateString(work_order.Effective_date)
	work_order.Expiry_date = formatDateString(work_order.Expiry_date)
	work_order.Final_billing_date = formatDateString(work_order.Final_billing_date)
	work_order.Request_date = formatDateString(work_order.Request_date)
	work_order.Row_effective_date = formatDateString(work_order.Row_effective_date)
	work_order.Row_expiry_date = formatDateString(work_order.Row_expiry_date)






	rows, err := stmt.Exec(work_order.Active_ind, work_order.Complete_date, work_order.Due_date, work_order.Effective_date, work_order.Expiry_date, work_order.Final_billing_date, work_order.Instructions, work_order.Ppdm_guid, work_order.Remark, work_order.Request_date, work_order.Rush_ind, work_order.Source, work_order.Work_order_number, work_order.Work_order_type, work_order.Row_changed_by, work_order.Row_changed_date, work_order.Row_created_by, work_order.Row_effective_date, work_order.Row_expiry_date, work_order.Row_quality, work_order.Work_order_id)
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

func PatchWorkOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update work_order set "
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
		} else if key == "complete_date" || key == "due_date" || key == "effective_date" || key == "expiry_date" || key == "final_billing_date" || key == "request_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWorkOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var work_order dto.Work_order
	work_order.Work_order_id = id

	stmt, err := db.Prepare("delete from work_order where work_order_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(work_order.Work_order_id)
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


