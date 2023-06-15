package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaAddress(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_address")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_address

	for rows.Next() {
		var ba_address dto.Ba_address
		if err := rows.Scan(&ba_address.Business_associate_id, &ba_address.Source, &ba_address.Address_obs_no, &ba_address.Active_ind, &ba_address.Addressee_text, &ba_address.Address_type, &ba_address.City_area_id, &ba_address.City_area_type, &ba_address.Country_area_id, &ba_address.Country_area_type, &ba_address.Effective_date, &ba_address.Expiry_date, &ba_address.First_address_line, &ba_address.Office_type, &ba_address.Postal_zip_code, &ba_address.Ppdm_guid, &ba_address.Preferred_ind, &ba_address.Province_state_area_id, &ba_address.Province_state_area_type, &ba_address.Remark, &ba_address.Second_address_line, &ba_address.Service_period, &ba_address.Service_quality, &ba_address.Third_address_line, &ba_address.Withholding_tax_ind, &ba_address.Row_changed_by, &ba_address.Row_changed_date, &ba_address.Row_created_by, &ba_address.Row_created_date, &ba_address.Row_effective_date, &ba_address.Row_expiry_date, &ba_address.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_address to result
		result = append(result, ba_address)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaAddress(c *fiber.Ctx) error {
	var ba_address dto.Ba_address

	setDefaults(&ba_address)

	if err := json.Unmarshal(c.Body(), &ba_address); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_address values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	ba_address.Row_created_date = formatDate(ba_address.Row_created_date)
	ba_address.Row_changed_date = formatDate(ba_address.Row_changed_date)
	ba_address.Effective_date = formatDateString(ba_address.Effective_date)
	ba_address.Expiry_date = formatDateString(ba_address.Expiry_date)
	ba_address.Row_effective_date = formatDateString(ba_address.Row_effective_date)
	ba_address.Row_expiry_date = formatDateString(ba_address.Row_expiry_date)






	rows, err := stmt.Exec(ba_address.Business_associate_id, ba_address.Source, ba_address.Address_obs_no, ba_address.Active_ind, ba_address.Addressee_text, ba_address.Address_type, ba_address.City_area_id, ba_address.City_area_type, ba_address.Country_area_id, ba_address.Country_area_type, ba_address.Effective_date, ba_address.Expiry_date, ba_address.First_address_line, ba_address.Office_type, ba_address.Postal_zip_code, ba_address.Ppdm_guid, ba_address.Preferred_ind, ba_address.Province_state_area_id, ba_address.Province_state_area_type, ba_address.Remark, ba_address.Second_address_line, ba_address.Service_period, ba_address.Service_quality, ba_address.Third_address_line, ba_address.Withholding_tax_ind, ba_address.Row_changed_by, ba_address.Row_changed_date, ba_address.Row_created_by, ba_address.Row_created_date, ba_address.Row_effective_date, ba_address.Row_expiry_date, ba_address.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaAddress(c *fiber.Ctx) error {
	var ba_address dto.Ba_address

	setDefaults(&ba_address)

	if err := json.Unmarshal(c.Body(), &ba_address); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_address.Business_associate_id = id

    if ba_address.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_address set source = :1, address_obs_no = :2, active_ind = :3, addressee_text = :4, address_type = :5, city_area_id = :6, city_area_type = :7, country_area_id = :8, country_area_type = :9, effective_date = :10, expiry_date = :11, first_address_line = :12, office_type = :13, postal_zip_code = :14, ppdm_guid = :15, preferred_ind = :16, province_state_area_id = :17, province_state_area_type = :18, remark = :19, second_address_line = :20, service_period = :21, service_quality = :22, third_address_line = :23, withholding_tax_ind = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where business_associate_id = :32")
	if err != nil {
		return err
	}

	ba_address.Row_changed_date = formatDate(ba_address.Row_changed_date)
	ba_address.Effective_date = formatDateString(ba_address.Effective_date)
	ba_address.Expiry_date = formatDateString(ba_address.Expiry_date)
	ba_address.Row_effective_date = formatDateString(ba_address.Row_effective_date)
	ba_address.Row_expiry_date = formatDateString(ba_address.Row_expiry_date)






	rows, err := stmt.Exec(ba_address.Source, ba_address.Address_obs_no, ba_address.Active_ind, ba_address.Addressee_text, ba_address.Address_type, ba_address.City_area_id, ba_address.City_area_type, ba_address.Country_area_id, ba_address.Country_area_type, ba_address.Effective_date, ba_address.Expiry_date, ba_address.First_address_line, ba_address.Office_type, ba_address.Postal_zip_code, ba_address.Ppdm_guid, ba_address.Preferred_ind, ba_address.Province_state_area_id, ba_address.Province_state_area_type, ba_address.Remark, ba_address.Second_address_line, ba_address.Service_period, ba_address.Service_quality, ba_address.Third_address_line, ba_address.Withholding_tax_ind, ba_address.Row_changed_by, ba_address.Row_changed_date, ba_address.Row_created_by, ba_address.Row_effective_date, ba_address.Row_expiry_date, ba_address.Row_quality, ba_address.Business_associate_id)
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

func PatchBaAddress(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_address set "
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

func DeleteBaAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_address dto.Ba_address
	ba_address.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_address where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_address.Business_associate_id)
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


