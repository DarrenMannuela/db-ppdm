package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaAuthority(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_authority")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_authority

	for rows.Next() {
		var ba_authority dto.Ba_authority
		if err := rows.Scan(&ba_authority.Business_associate_id, &ba_authority.Authority_id, &ba_authority.Active_ind, &ba_authority.Authority_limit, &ba_authority.Authority_type, &ba_authority.Authorized_by, &ba_authority.Currency_conversion, &ba_authority.Currency_ouom, &ba_authority.Effective_date, &ba_authority.Expiry_date, &ba_authority.Ppdm_guid, &ba_authority.Remark, &ba_authority.Represented_ba_id, &ba_authority.Source, &ba_authority.Row_changed_by, &ba_authority.Row_changed_date, &ba_authority.Row_created_by, &ba_authority.Row_created_date, &ba_authority.Row_effective_date, &ba_authority.Row_expiry_date, &ba_authority.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_authority to result
		result = append(result, ba_authority)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaAuthority(c *fiber.Ctx) error {
	var ba_authority dto.Ba_authority

	setDefaults(&ba_authority)

	if err := json.Unmarshal(c.Body(), &ba_authority); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_authority values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	ba_authority.Row_created_date = formatDate(ba_authority.Row_created_date)
	ba_authority.Row_changed_date = formatDate(ba_authority.Row_changed_date)
	ba_authority.Effective_date = formatDateString(ba_authority.Effective_date)
	ba_authority.Expiry_date = formatDateString(ba_authority.Expiry_date)
	ba_authority.Row_effective_date = formatDateString(ba_authority.Row_effective_date)
	ba_authority.Row_expiry_date = formatDateString(ba_authority.Row_expiry_date)






	rows, err := stmt.Exec(ba_authority.Business_associate_id, ba_authority.Authority_id, ba_authority.Active_ind, ba_authority.Authority_limit, ba_authority.Authority_type, ba_authority.Authorized_by, ba_authority.Currency_conversion, ba_authority.Currency_ouom, ba_authority.Effective_date, ba_authority.Expiry_date, ba_authority.Ppdm_guid, ba_authority.Remark, ba_authority.Represented_ba_id, ba_authority.Source, ba_authority.Row_changed_by, ba_authority.Row_changed_date, ba_authority.Row_created_by, ba_authority.Row_created_date, ba_authority.Row_effective_date, ba_authority.Row_expiry_date, ba_authority.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaAuthority(c *fiber.Ctx) error {
	var ba_authority dto.Ba_authority

	setDefaults(&ba_authority)

	if err := json.Unmarshal(c.Body(), &ba_authority); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_authority.Business_associate_id = id

    if ba_authority.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_authority set authority_id = :1, active_ind = :2, authority_limit = :3, authority_type = :4, authorized_by = :5, currency_conversion = :6, currency_ouom = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, represented_ba_id = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where business_associate_id = :21")
	if err != nil {
		return err
	}

	ba_authority.Row_changed_date = formatDate(ba_authority.Row_changed_date)
	ba_authority.Effective_date = formatDateString(ba_authority.Effective_date)
	ba_authority.Expiry_date = formatDateString(ba_authority.Expiry_date)
	ba_authority.Row_effective_date = formatDateString(ba_authority.Row_effective_date)
	ba_authority.Row_expiry_date = formatDateString(ba_authority.Row_expiry_date)






	rows, err := stmt.Exec(ba_authority.Authority_id, ba_authority.Active_ind, ba_authority.Authority_limit, ba_authority.Authority_type, ba_authority.Authorized_by, ba_authority.Currency_conversion, ba_authority.Currency_ouom, ba_authority.Effective_date, ba_authority.Expiry_date, ba_authority.Ppdm_guid, ba_authority.Remark, ba_authority.Represented_ba_id, ba_authority.Source, ba_authority.Row_changed_by, ba_authority.Row_changed_date, ba_authority.Row_created_by, ba_authority.Row_effective_date, ba_authority.Row_expiry_date, ba_authority.Row_quality, ba_authority.Business_associate_id)
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

func PatchBaAuthority(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_authority set "
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

func DeleteBaAuthority(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_authority dto.Ba_authority
	ba_authority.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_authority where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_authority.Business_associate_id)
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


