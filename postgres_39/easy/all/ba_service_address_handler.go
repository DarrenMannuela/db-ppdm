package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaServiceAddress(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_service_address")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_service_address

	for rows.Next() {
		var ba_service_address dto.Ba_service_address
		if err := rows.Scan(&ba_service_address.Business_associate_id, &ba_service_address.Address_obs_no, &ba_service_address.Ba_service_type, &ba_service_address.Ba_service_seq_no, &ba_service_address.Active_ind, &ba_service_address.Address_source, &ba_service_address.Alias_source, &ba_service_address.Ba_alias_id, &ba_service_address.Ba_service_source, &ba_service_address.Effective_date, &ba_service_address.Expiry_date, &ba_service_address.Ppdm_guid, &ba_service_address.Remark, &ba_service_address.Service_quality, &ba_service_address.Source, &ba_service_address.Row_changed_by, &ba_service_address.Row_changed_date, &ba_service_address.Row_created_by, &ba_service_address.Row_created_date, &ba_service_address.Row_effective_date, &ba_service_address.Row_expiry_date, &ba_service_address.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_service_address to result
		result = append(result, ba_service_address)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaServiceAddress(c *fiber.Ctx) error {
	var ba_service_address dto.Ba_service_address

	setDefaults(&ba_service_address)

	if err := json.Unmarshal(c.Body(), &ba_service_address); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_service_address values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	ba_service_address.Row_created_date = formatDate(ba_service_address.Row_created_date)
	ba_service_address.Row_changed_date = formatDate(ba_service_address.Row_changed_date)
	ba_service_address.Effective_date = formatDateString(ba_service_address.Effective_date)
	ba_service_address.Expiry_date = formatDateString(ba_service_address.Expiry_date)
	ba_service_address.Row_effective_date = formatDateString(ba_service_address.Row_effective_date)
	ba_service_address.Row_expiry_date = formatDateString(ba_service_address.Row_expiry_date)






	rows, err := stmt.Exec(ba_service_address.Business_associate_id, ba_service_address.Address_obs_no, ba_service_address.Ba_service_type, ba_service_address.Ba_service_seq_no, ba_service_address.Active_ind, ba_service_address.Address_source, ba_service_address.Alias_source, ba_service_address.Ba_alias_id, ba_service_address.Ba_service_source, ba_service_address.Effective_date, ba_service_address.Expiry_date, ba_service_address.Ppdm_guid, ba_service_address.Remark, ba_service_address.Service_quality, ba_service_address.Source, ba_service_address.Row_changed_by, ba_service_address.Row_changed_date, ba_service_address.Row_created_by, ba_service_address.Row_created_date, ba_service_address.Row_effective_date, ba_service_address.Row_expiry_date, ba_service_address.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaServiceAddress(c *fiber.Ctx) error {
	var ba_service_address dto.Ba_service_address

	setDefaults(&ba_service_address)

	if err := json.Unmarshal(c.Body(), &ba_service_address); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_service_address.Business_associate_id = id

    if ba_service_address.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_service_address set address_obs_no = :1, ba_service_type = :2, ba_service_seq_no = :3, active_ind = :4, address_source = :5, alias_source = :6, ba_alias_id = :7, ba_service_source = :8, effective_date = :9, expiry_date = :10, ppdm_guid = :11, remark = :12, service_quality = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where business_associate_id = :22")
	if err != nil {
		return err
	}

	ba_service_address.Row_changed_date = formatDate(ba_service_address.Row_changed_date)
	ba_service_address.Effective_date = formatDateString(ba_service_address.Effective_date)
	ba_service_address.Expiry_date = formatDateString(ba_service_address.Expiry_date)
	ba_service_address.Row_effective_date = formatDateString(ba_service_address.Row_effective_date)
	ba_service_address.Row_expiry_date = formatDateString(ba_service_address.Row_expiry_date)






	rows, err := stmt.Exec(ba_service_address.Address_obs_no, ba_service_address.Ba_service_type, ba_service_address.Ba_service_seq_no, ba_service_address.Active_ind, ba_service_address.Address_source, ba_service_address.Alias_source, ba_service_address.Ba_alias_id, ba_service_address.Ba_service_source, ba_service_address.Effective_date, ba_service_address.Expiry_date, ba_service_address.Ppdm_guid, ba_service_address.Remark, ba_service_address.Service_quality, ba_service_address.Source, ba_service_address.Row_changed_by, ba_service_address.Row_changed_date, ba_service_address.Row_created_by, ba_service_address.Row_effective_date, ba_service_address.Row_expiry_date, ba_service_address.Row_quality, ba_service_address.Business_associate_id)
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

func PatchBaServiceAddress(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_service_address set "
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

func DeleteBaServiceAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_service_address dto.Ba_service_address
	ba_service_address.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_service_address where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_service_address.Business_associate_id)
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


