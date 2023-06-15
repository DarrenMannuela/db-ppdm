package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBusinessAssociate(c *fiber.Ctx) error {
	rows, err := db.Query("select * from business_associate")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Business_associate

	for rows.Next() {
		var business_associate dto.Business_associate
		if err := rows.Scan(&business_associate.Business_associate_id, &business_associate.Active_ind, &business_associate.Ba_abbreviation, &business_associate.Ba_category, &business_associate.Ba_code, &business_associate.Ba_long_name, &business_associate.Ba_short_name, &business_associate.Ba_type, &business_associate.Credit_check_date, &business_associate.Credit_check_ind, &business_associate.Credit_check_source, &business_associate.Credit_rating, &business_associate.Credit_rating_source, &business_associate.Current_status, &business_associate.Effective_date, &business_associate.Expiry_date, &business_associate.First_name, &business_associate.Last_name, &business_associate.Middle_initial, &business_associate.Ppdm_guid, &business_associate.Remark, &business_associate.Service_period, &business_associate.Source, &business_associate.Row_changed_by, &business_associate.Row_changed_date, &business_associate.Row_created_by, &business_associate.Row_created_date, &business_associate.Row_effective_date, &business_associate.Row_expiry_date, &business_associate.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append business_associate to result
		result = append(result, business_associate)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBusinessAssociate(c *fiber.Ctx) error {
	var business_associate dto.Business_associate

	setDefaults(&business_associate)

	if err := json.Unmarshal(c.Body(), &business_associate); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into business_associate values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	business_associate.Row_created_date = formatDate(business_associate.Row_created_date)
	business_associate.Row_changed_date = formatDate(business_associate.Row_changed_date)
	business_associate.Credit_check_date = formatDateString(business_associate.Credit_check_date)
	business_associate.Effective_date = formatDateString(business_associate.Effective_date)
	business_associate.Expiry_date = formatDateString(business_associate.Expiry_date)
	business_associate.Row_effective_date = formatDateString(business_associate.Row_effective_date)
	business_associate.Row_expiry_date = formatDateString(business_associate.Row_expiry_date)






	rows, err := stmt.Exec(business_associate.Business_associate_id, business_associate.Active_ind, business_associate.Ba_abbreviation, business_associate.Ba_category, business_associate.Ba_code, business_associate.Ba_long_name, business_associate.Ba_short_name, business_associate.Ba_type, business_associate.Credit_check_date, business_associate.Credit_check_ind, business_associate.Credit_check_source, business_associate.Credit_rating, business_associate.Credit_rating_source, business_associate.Current_status, business_associate.Effective_date, business_associate.Expiry_date, business_associate.First_name, business_associate.Last_name, business_associate.Middle_initial, business_associate.Ppdm_guid, business_associate.Remark, business_associate.Service_period, business_associate.Source, business_associate.Row_changed_by, business_associate.Row_changed_date, business_associate.Row_created_by, business_associate.Row_created_date, business_associate.Row_effective_date, business_associate.Row_expiry_date, business_associate.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBusinessAssociate(c *fiber.Ctx) error {
	var business_associate dto.Business_associate

	setDefaults(&business_associate)

	if err := json.Unmarshal(c.Body(), &business_associate); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	business_associate.Business_associate_id = id

    if business_associate.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update business_associate set active_ind = :1, ba_abbreviation = :2, ba_category = :3, ba_code = :4, ba_long_name = :5, ba_short_name = :6, ba_type = :7, credit_check_date = :8, credit_check_ind = :9, credit_check_source = :10, credit_rating = :11, credit_rating_source = :12, current_status = :13, effective_date = :14, expiry_date = :15, first_name = :16, last_name = :17, middle_initial = :18, ppdm_guid = :19, remark = :20, service_period = :21, source = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where business_associate_id = :30")
	if err != nil {
		return err
	}

	business_associate.Row_changed_date = formatDate(business_associate.Row_changed_date)
	business_associate.Credit_check_date = formatDateString(business_associate.Credit_check_date)
	business_associate.Effective_date = formatDateString(business_associate.Effective_date)
	business_associate.Expiry_date = formatDateString(business_associate.Expiry_date)
	business_associate.Row_effective_date = formatDateString(business_associate.Row_effective_date)
	business_associate.Row_expiry_date = formatDateString(business_associate.Row_expiry_date)






	rows, err := stmt.Exec(business_associate.Active_ind, business_associate.Ba_abbreviation, business_associate.Ba_category, business_associate.Ba_code, business_associate.Ba_long_name, business_associate.Ba_short_name, business_associate.Ba_type, business_associate.Credit_check_date, business_associate.Credit_check_ind, business_associate.Credit_check_source, business_associate.Credit_rating, business_associate.Credit_rating_source, business_associate.Current_status, business_associate.Effective_date, business_associate.Expiry_date, business_associate.First_name, business_associate.Last_name, business_associate.Middle_initial, business_associate.Ppdm_guid, business_associate.Remark, business_associate.Service_period, business_associate.Source, business_associate.Row_changed_by, business_associate.Row_changed_date, business_associate.Row_created_by, business_associate.Row_effective_date, business_associate.Row_expiry_date, business_associate.Row_quality, business_associate.Business_associate_id)
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

func PatchBusinessAssociate(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update business_associate set "
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
		} else if key == "credit_check_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteBusinessAssociate(c *fiber.Ctx) error {
	id := c.Params("id")
	var business_associate dto.Business_associate
	business_associate.Business_associate_id = id

	stmt, err := db.Prepare("delete from business_associate where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(business_associate.Business_associate_id)
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


