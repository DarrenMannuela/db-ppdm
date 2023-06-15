package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaService(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_service")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_service

	for rows.Next() {
		var ba_service dto.Ba_service
		if err := rows.Scan(&ba_service.Business_associate_id, &ba_service.Ba_service_type, &ba_service.Ba_service_seq_no, &ba_service.Active_ind, &ba_service.Effective_date, &ba_service.Expiry_date, &ba_service.Ppdm_guid, &ba_service.Remark, &ba_service.Service_quality, &ba_service.Source, &ba_service.Row_changed_by, &ba_service.Row_changed_date, &ba_service.Row_created_by, &ba_service.Row_created_date, &ba_service.Row_effective_date, &ba_service.Row_expiry_date, &ba_service.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_service to result
		result = append(result, ba_service)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaService(c *fiber.Ctx) error {
	var ba_service dto.Ba_service

	setDefaults(&ba_service)

	if err := json.Unmarshal(c.Body(), &ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_service values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	ba_service.Row_created_date = formatDate(ba_service.Row_created_date)
	ba_service.Row_changed_date = formatDate(ba_service.Row_changed_date)
	ba_service.Effective_date = formatDateString(ba_service.Effective_date)
	ba_service.Expiry_date = formatDateString(ba_service.Expiry_date)
	ba_service.Row_effective_date = formatDateString(ba_service.Row_effective_date)
	ba_service.Row_expiry_date = formatDateString(ba_service.Row_expiry_date)






	rows, err := stmt.Exec(ba_service.Business_associate_id, ba_service.Ba_service_type, ba_service.Ba_service_seq_no, ba_service.Active_ind, ba_service.Effective_date, ba_service.Expiry_date, ba_service.Ppdm_guid, ba_service.Remark, ba_service.Service_quality, ba_service.Source, ba_service.Row_changed_by, ba_service.Row_changed_date, ba_service.Row_created_by, ba_service.Row_created_date, ba_service.Row_effective_date, ba_service.Row_expiry_date, ba_service.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaService(c *fiber.Ctx) error {
	var ba_service dto.Ba_service

	setDefaults(&ba_service)

	if err := json.Unmarshal(c.Body(), &ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_service.Business_associate_id = id

    if ba_service.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_service set ba_service_type = :1, ba_service_seq_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, service_quality = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where business_associate_id = :17")
	if err != nil {
		return err
	}

	ba_service.Row_changed_date = formatDate(ba_service.Row_changed_date)
	ba_service.Effective_date = formatDateString(ba_service.Effective_date)
	ba_service.Expiry_date = formatDateString(ba_service.Expiry_date)
	ba_service.Row_effective_date = formatDateString(ba_service.Row_effective_date)
	ba_service.Row_expiry_date = formatDateString(ba_service.Row_expiry_date)






	rows, err := stmt.Exec(ba_service.Ba_service_type, ba_service.Ba_service_seq_no, ba_service.Active_ind, ba_service.Effective_date, ba_service.Expiry_date, ba_service.Ppdm_guid, ba_service.Remark, ba_service.Service_quality, ba_service.Source, ba_service.Row_changed_by, ba_service.Row_changed_date, ba_service.Row_created_by, ba_service.Row_effective_date, ba_service.Row_expiry_date, ba_service.Row_quality, ba_service.Business_associate_id)
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

func PatchBaService(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_service set "
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
	query += " where business_associate_id = :id"

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

func DeleteBaService(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_service dto.Ba_service
	ba_service.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_service where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_service.Business_associate_id)
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


