package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisSetAuthorize(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_set_authorize")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_set_authorize

	for rows.Next() {
		var seis_set_authorize dto.Seis_set_authorize
		if err := rows.Scan(&seis_set_authorize.Seis_set_subtype, &seis_set_authorize.Seis_set_id, &seis_set_authorize.Authorize_id, &seis_set_authorize.Active_ind, &seis_set_authorize.Authorized_by, &seis_set_authorize.Authorize_type, &seis_set_authorize.Currency_conversion, &seis_set_authorize.Currency_ouom, &seis_set_authorize.Effective_date, &seis_set_authorize.Expiry_date, &seis_set_authorize.Limit_desc, &seis_set_authorize.Limit_type, &seis_set_authorize.Partner_ba_id, &seis_set_authorize.Ppdm_guid, &seis_set_authorize.Price_per_length, &seis_set_authorize.Reason_type, &seis_set_authorize.Remark, &seis_set_authorize.Source, &seis_set_authorize.Row_changed_by, &seis_set_authorize.Row_changed_date, &seis_set_authorize.Row_created_by, &seis_set_authorize.Row_created_date, &seis_set_authorize.Row_effective_date, &seis_set_authorize.Row_expiry_date, &seis_set_authorize.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_set_authorize to result
		result = append(result, seis_set_authorize)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisSetAuthorize(c *fiber.Ctx) error {
	var seis_set_authorize dto.Seis_set_authorize

	setDefaults(&seis_set_authorize)

	if err := json.Unmarshal(c.Body(), &seis_set_authorize); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_set_authorize values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	seis_set_authorize.Row_created_date = formatDate(seis_set_authorize.Row_created_date)
	seis_set_authorize.Row_changed_date = formatDate(seis_set_authorize.Row_changed_date)
	seis_set_authorize.Effective_date = formatDateString(seis_set_authorize.Effective_date)
	seis_set_authorize.Expiry_date = formatDateString(seis_set_authorize.Expiry_date)
	seis_set_authorize.Row_effective_date = formatDateString(seis_set_authorize.Row_effective_date)
	seis_set_authorize.Row_expiry_date = formatDateString(seis_set_authorize.Row_expiry_date)






	rows, err := stmt.Exec(seis_set_authorize.Seis_set_subtype, seis_set_authorize.Seis_set_id, seis_set_authorize.Authorize_id, seis_set_authorize.Active_ind, seis_set_authorize.Authorized_by, seis_set_authorize.Authorize_type, seis_set_authorize.Currency_conversion, seis_set_authorize.Currency_ouom, seis_set_authorize.Effective_date, seis_set_authorize.Expiry_date, seis_set_authorize.Limit_desc, seis_set_authorize.Limit_type, seis_set_authorize.Partner_ba_id, seis_set_authorize.Ppdm_guid, seis_set_authorize.Price_per_length, seis_set_authorize.Reason_type, seis_set_authorize.Remark, seis_set_authorize.Source, seis_set_authorize.Row_changed_by, seis_set_authorize.Row_changed_date, seis_set_authorize.Row_created_by, seis_set_authorize.Row_created_date, seis_set_authorize.Row_effective_date, seis_set_authorize.Row_expiry_date, seis_set_authorize.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisSetAuthorize(c *fiber.Ctx) error {
	var seis_set_authorize dto.Seis_set_authorize

	setDefaults(&seis_set_authorize)

	if err := json.Unmarshal(c.Body(), &seis_set_authorize); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_set_authorize.Seis_set_subtype = id

    if seis_set_authorize.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_set_authorize set seis_set_id = :1, authorize_id = :2, active_ind = :3, authorized_by = :4, authorize_type = :5, currency_conversion = :6, currency_ouom = :7, effective_date = :8, expiry_date = :9, limit_desc = :10, limit_type = :11, partner_ba_id = :12, ppdm_guid = :13, price_per_length = :14, reason_type = :15, remark = :16, source = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where seis_set_subtype = :25")
	if err != nil {
		return err
	}

	seis_set_authorize.Row_changed_date = formatDate(seis_set_authorize.Row_changed_date)
	seis_set_authorize.Effective_date = formatDateString(seis_set_authorize.Effective_date)
	seis_set_authorize.Expiry_date = formatDateString(seis_set_authorize.Expiry_date)
	seis_set_authorize.Row_effective_date = formatDateString(seis_set_authorize.Row_effective_date)
	seis_set_authorize.Row_expiry_date = formatDateString(seis_set_authorize.Row_expiry_date)






	rows, err := stmt.Exec(seis_set_authorize.Seis_set_id, seis_set_authorize.Authorize_id, seis_set_authorize.Active_ind, seis_set_authorize.Authorized_by, seis_set_authorize.Authorize_type, seis_set_authorize.Currency_conversion, seis_set_authorize.Currency_ouom, seis_set_authorize.Effective_date, seis_set_authorize.Expiry_date, seis_set_authorize.Limit_desc, seis_set_authorize.Limit_type, seis_set_authorize.Partner_ba_id, seis_set_authorize.Ppdm_guid, seis_set_authorize.Price_per_length, seis_set_authorize.Reason_type, seis_set_authorize.Remark, seis_set_authorize.Source, seis_set_authorize.Row_changed_by, seis_set_authorize.Row_changed_date, seis_set_authorize.Row_created_by, seis_set_authorize.Row_effective_date, seis_set_authorize.Row_expiry_date, seis_set_authorize.Row_quality, seis_set_authorize.Seis_set_subtype)
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

func PatchSeisSetAuthorize(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_set_authorize set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeisSetAuthorize(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_set_authorize dto.Seis_set_authorize
	seis_set_authorize.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_set_authorize where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_set_authorize.Seis_set_subtype)
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


