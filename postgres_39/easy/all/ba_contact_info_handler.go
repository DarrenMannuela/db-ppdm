package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaContactInfo(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_contact_info")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_contact_info

	for rows.Next() {
		var ba_contact_info dto.Ba_contact_info
		if err := rows.Scan(&ba_contact_info.Business_associate_id, &ba_contact_info.Location_id, &ba_contact_info.Active_ind, &ba_contact_info.Address_obs_no, &ba_contact_info.Address_source, &ba_contact_info.Ba_organization_id, &ba_contact_info.Ba_organization_seq_no, &ba_contact_info.Contact_loc_type, &ba_contact_info.Effective_date, &ba_contact_info.Expiry_date, &ba_contact_info.Location_name, &ba_contact_info.Order_level, &ba_contact_info.Ppdm_guid, &ba_contact_info.Preferred_ind, &ba_contact_info.Remark, &ba_contact_info.Source, &ba_contact_info.Row_changed_by, &ba_contact_info.Row_changed_date, &ba_contact_info.Row_created_by, &ba_contact_info.Row_created_date, &ba_contact_info.Row_effective_date, &ba_contact_info.Row_expiry_date, &ba_contact_info.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_contact_info to result
		result = append(result, ba_contact_info)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaContactInfo(c *fiber.Ctx) error {
	var ba_contact_info dto.Ba_contact_info

	setDefaults(&ba_contact_info)

	if err := json.Unmarshal(c.Body(), &ba_contact_info); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_contact_info values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	ba_contact_info.Row_created_date = formatDate(ba_contact_info.Row_created_date)
	ba_contact_info.Row_changed_date = formatDate(ba_contact_info.Row_changed_date)
	ba_contact_info.Effective_date = formatDateString(ba_contact_info.Effective_date)
	ba_contact_info.Expiry_date = formatDateString(ba_contact_info.Expiry_date)
	ba_contact_info.Row_effective_date = formatDateString(ba_contact_info.Row_effective_date)
	ba_contact_info.Row_expiry_date = formatDateString(ba_contact_info.Row_expiry_date)






	rows, err := stmt.Exec(ba_contact_info.Business_associate_id, ba_contact_info.Location_id, ba_contact_info.Active_ind, ba_contact_info.Address_obs_no, ba_contact_info.Address_source, ba_contact_info.Ba_organization_id, ba_contact_info.Ba_organization_seq_no, ba_contact_info.Contact_loc_type, ba_contact_info.Effective_date, ba_contact_info.Expiry_date, ba_contact_info.Location_name, ba_contact_info.Order_level, ba_contact_info.Ppdm_guid, ba_contact_info.Preferred_ind, ba_contact_info.Remark, ba_contact_info.Source, ba_contact_info.Row_changed_by, ba_contact_info.Row_changed_date, ba_contact_info.Row_created_by, ba_contact_info.Row_created_date, ba_contact_info.Row_effective_date, ba_contact_info.Row_expiry_date, ba_contact_info.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaContactInfo(c *fiber.Ctx) error {
	var ba_contact_info dto.Ba_contact_info

	setDefaults(&ba_contact_info)

	if err := json.Unmarshal(c.Body(), &ba_contact_info); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_contact_info.Business_associate_id = id

    if ba_contact_info.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_contact_info set location_id = :1, active_ind = :2, address_obs_no = :3, address_source = :4, ba_organization_id = :5, ba_organization_seq_no = :6, contact_loc_type = :7, effective_date = :8, expiry_date = :9, location_name = :10, order_level = :11, ppdm_guid = :12, preferred_ind = :13, remark = :14, source = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where business_associate_id = :23")
	if err != nil {
		return err
	}

	ba_contact_info.Row_changed_date = formatDate(ba_contact_info.Row_changed_date)
	ba_contact_info.Effective_date = formatDateString(ba_contact_info.Effective_date)
	ba_contact_info.Expiry_date = formatDateString(ba_contact_info.Expiry_date)
	ba_contact_info.Row_effective_date = formatDateString(ba_contact_info.Row_effective_date)
	ba_contact_info.Row_expiry_date = formatDateString(ba_contact_info.Row_expiry_date)






	rows, err := stmt.Exec(ba_contact_info.Location_id, ba_contact_info.Active_ind, ba_contact_info.Address_obs_no, ba_contact_info.Address_source, ba_contact_info.Ba_organization_id, ba_contact_info.Ba_organization_seq_no, ba_contact_info.Contact_loc_type, ba_contact_info.Effective_date, ba_contact_info.Expiry_date, ba_contact_info.Location_name, ba_contact_info.Order_level, ba_contact_info.Ppdm_guid, ba_contact_info.Preferred_ind, ba_contact_info.Remark, ba_contact_info.Source, ba_contact_info.Row_changed_by, ba_contact_info.Row_changed_date, ba_contact_info.Row_created_by, ba_contact_info.Row_effective_date, ba_contact_info.Row_expiry_date, ba_contact_info.Row_quality, ba_contact_info.Business_associate_id)
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

func PatchBaContactInfo(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_contact_info set "
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

func DeleteBaContactInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_contact_info dto.Ba_contact_info
	ba_contact_info.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_contact_info where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_contact_info.Business_associate_id)
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


