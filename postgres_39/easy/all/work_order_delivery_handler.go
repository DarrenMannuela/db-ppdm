package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWorkOrderDelivery(c *fiber.Ctx) error {
	rows, err := db.Query("select * from work_order_delivery")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Work_order_delivery

	for rows.Next() {
		var work_order_delivery dto.Work_order_delivery
		if err := rows.Scan(&work_order_delivery.Work_order_id, &work_order_delivery.Delivery_id, &work_order_delivery.Active_ind, &work_order_delivery.Ba_obs_no, &work_order_delivery.Business_associate_id, &work_order_delivery.Delivery_ba_address_obs_no, &work_order_delivery.Delivery_ba_address_source, &work_order_delivery.Delivery_ba_id, &work_order_delivery.Delivery_date, &work_order_delivery.Delivery_type, &work_order_delivery.Effective_date, &work_order_delivery.Expiry_date, &work_order_delivery.Ppdm_guid, &work_order_delivery.Remark, &work_order_delivery.Source, &work_order_delivery.Row_changed_by, &work_order_delivery.Row_changed_date, &work_order_delivery.Row_created_by, &work_order_delivery.Row_created_date, &work_order_delivery.Row_effective_date, &work_order_delivery.Row_expiry_date, &work_order_delivery.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append work_order_delivery to result
		result = append(result, work_order_delivery)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWorkOrderDelivery(c *fiber.Ctx) error {
	var work_order_delivery dto.Work_order_delivery

	setDefaults(&work_order_delivery)

	if err := json.Unmarshal(c.Body(), &work_order_delivery); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into work_order_delivery values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	work_order_delivery.Row_created_date = formatDate(work_order_delivery.Row_created_date)
	work_order_delivery.Row_changed_date = formatDate(work_order_delivery.Row_changed_date)
	work_order_delivery.Delivery_date = formatDateString(work_order_delivery.Delivery_date)
	work_order_delivery.Effective_date = formatDateString(work_order_delivery.Effective_date)
	work_order_delivery.Expiry_date = formatDateString(work_order_delivery.Expiry_date)
	work_order_delivery.Row_effective_date = formatDateString(work_order_delivery.Row_effective_date)
	work_order_delivery.Row_expiry_date = formatDateString(work_order_delivery.Row_expiry_date)






	rows, err := stmt.Exec(work_order_delivery.Work_order_id, work_order_delivery.Delivery_id, work_order_delivery.Active_ind, work_order_delivery.Ba_obs_no, work_order_delivery.Business_associate_id, work_order_delivery.Delivery_ba_address_obs_no, work_order_delivery.Delivery_ba_address_source, work_order_delivery.Delivery_ba_id, work_order_delivery.Delivery_date, work_order_delivery.Delivery_type, work_order_delivery.Effective_date, work_order_delivery.Expiry_date, work_order_delivery.Ppdm_guid, work_order_delivery.Remark, work_order_delivery.Source, work_order_delivery.Row_changed_by, work_order_delivery.Row_changed_date, work_order_delivery.Row_created_by, work_order_delivery.Row_created_date, work_order_delivery.Row_effective_date, work_order_delivery.Row_expiry_date, work_order_delivery.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWorkOrderDelivery(c *fiber.Ctx) error {
	var work_order_delivery dto.Work_order_delivery

	setDefaults(&work_order_delivery)

	if err := json.Unmarshal(c.Body(), &work_order_delivery); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	work_order_delivery.Work_order_id = id

    if work_order_delivery.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update work_order_delivery set delivery_id = :1, active_ind = :2, ba_obs_no = :3, business_associate_id = :4, delivery_ba_address_obs_no = :5, delivery_ba_address_source = :6, delivery_ba_id = :7, delivery_date = :8, delivery_type = :9, effective_date = :10, expiry_date = :11, ppdm_guid = :12, remark = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where work_order_id = :22")
	if err != nil {
		return err
	}

	work_order_delivery.Row_changed_date = formatDate(work_order_delivery.Row_changed_date)
	work_order_delivery.Delivery_date = formatDateString(work_order_delivery.Delivery_date)
	work_order_delivery.Effective_date = formatDateString(work_order_delivery.Effective_date)
	work_order_delivery.Expiry_date = formatDateString(work_order_delivery.Expiry_date)
	work_order_delivery.Row_effective_date = formatDateString(work_order_delivery.Row_effective_date)
	work_order_delivery.Row_expiry_date = formatDateString(work_order_delivery.Row_expiry_date)






	rows, err := stmt.Exec(work_order_delivery.Delivery_id, work_order_delivery.Active_ind, work_order_delivery.Ba_obs_no, work_order_delivery.Business_associate_id, work_order_delivery.Delivery_ba_address_obs_no, work_order_delivery.Delivery_ba_address_source, work_order_delivery.Delivery_ba_id, work_order_delivery.Delivery_date, work_order_delivery.Delivery_type, work_order_delivery.Effective_date, work_order_delivery.Expiry_date, work_order_delivery.Ppdm_guid, work_order_delivery.Remark, work_order_delivery.Source, work_order_delivery.Row_changed_by, work_order_delivery.Row_changed_date, work_order_delivery.Row_created_by, work_order_delivery.Row_effective_date, work_order_delivery.Row_expiry_date, work_order_delivery.Row_quality, work_order_delivery.Work_order_id)
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

func PatchWorkOrderDelivery(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update work_order_delivery set "
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
		} else if key == "delivery_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWorkOrderDelivery(c *fiber.Ctx) error {
	id := c.Params("id")
	var work_order_delivery dto.Work_order_delivery
	work_order_delivery.Work_order_id = id

	stmt, err := db.Prepare("delete from work_order_delivery where work_order_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(work_order_delivery.Work_order_id)
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


