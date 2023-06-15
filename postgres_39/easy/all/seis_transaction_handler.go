package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisTransaction(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_transaction")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_transaction

	for rows.Next() {
		var seis_transaction dto.Seis_transaction
		if err := rows.Scan(&seis_transaction.Seis_transaction_id, &seis_transaction.Transaction_type, &seis_transaction.Active_ind, &seis_transaction.Currency_conversion, &seis_transaction.Currency_ouom, &seis_transaction.Effective_date, &seis_transaction.Expiry_date, &seis_transaction.Length, &seis_transaction.Length_ouom, &seis_transaction.Ppdm_guid, &seis_transaction.Price_per_length, &seis_transaction.Reference_num, &seis_transaction.Remark, &seis_transaction.Source, &seis_transaction.Total_price, &seis_transaction.Transaction_date, &seis_transaction.Transaction_status, &seis_transaction.Row_changed_by, &seis_transaction.Row_changed_date, &seis_transaction.Row_created_by, &seis_transaction.Row_created_date, &seis_transaction.Row_effective_date, &seis_transaction.Row_expiry_date, &seis_transaction.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_transaction to result
		result = append(result, seis_transaction)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisTransaction(c *fiber.Ctx) error {
	var seis_transaction dto.Seis_transaction

	setDefaults(&seis_transaction)

	if err := json.Unmarshal(c.Body(), &seis_transaction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_transaction values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	seis_transaction.Row_created_date = formatDate(seis_transaction.Row_created_date)
	seis_transaction.Row_changed_date = formatDate(seis_transaction.Row_changed_date)
	seis_transaction.Effective_date = formatDateString(seis_transaction.Effective_date)
	seis_transaction.Expiry_date = formatDateString(seis_transaction.Expiry_date)
	seis_transaction.Transaction_date = formatDateString(seis_transaction.Transaction_date)
	seis_transaction.Row_effective_date = formatDateString(seis_transaction.Row_effective_date)
	seis_transaction.Row_expiry_date = formatDateString(seis_transaction.Row_expiry_date)






	rows, err := stmt.Exec(seis_transaction.Seis_transaction_id, seis_transaction.Transaction_type, seis_transaction.Active_ind, seis_transaction.Currency_conversion, seis_transaction.Currency_ouom, seis_transaction.Effective_date, seis_transaction.Expiry_date, seis_transaction.Length, seis_transaction.Length_ouom, seis_transaction.Ppdm_guid, seis_transaction.Price_per_length, seis_transaction.Reference_num, seis_transaction.Remark, seis_transaction.Source, seis_transaction.Total_price, seis_transaction.Transaction_date, seis_transaction.Transaction_status, seis_transaction.Row_changed_by, seis_transaction.Row_changed_date, seis_transaction.Row_created_by, seis_transaction.Row_created_date, seis_transaction.Row_effective_date, seis_transaction.Row_expiry_date, seis_transaction.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisTransaction(c *fiber.Ctx) error {
	var seis_transaction dto.Seis_transaction

	setDefaults(&seis_transaction)

	if err := json.Unmarshal(c.Body(), &seis_transaction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_transaction.Seis_transaction_id = id

    if seis_transaction.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_transaction set transaction_type = :1, active_ind = :2, currency_conversion = :3, currency_ouom = :4, effective_date = :5, expiry_date = :6, length = :7, length_ouom = :8, ppdm_guid = :9, price_per_length = :10, reference_num = :11, remark = :12, source = :13, total_price = :14, transaction_date = :15, transaction_status = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where seis_transaction_id = :24")
	if err != nil {
		return err
	}

	seis_transaction.Row_changed_date = formatDate(seis_transaction.Row_changed_date)
	seis_transaction.Effective_date = formatDateString(seis_transaction.Effective_date)
	seis_transaction.Expiry_date = formatDateString(seis_transaction.Expiry_date)
	seis_transaction.Transaction_date = formatDateString(seis_transaction.Transaction_date)
	seis_transaction.Row_effective_date = formatDateString(seis_transaction.Row_effective_date)
	seis_transaction.Row_expiry_date = formatDateString(seis_transaction.Row_expiry_date)






	rows, err := stmt.Exec(seis_transaction.Transaction_type, seis_transaction.Active_ind, seis_transaction.Currency_conversion, seis_transaction.Currency_ouom, seis_transaction.Effective_date, seis_transaction.Expiry_date, seis_transaction.Length, seis_transaction.Length_ouom, seis_transaction.Ppdm_guid, seis_transaction.Price_per_length, seis_transaction.Reference_num, seis_transaction.Remark, seis_transaction.Source, seis_transaction.Total_price, seis_transaction.Transaction_date, seis_transaction.Transaction_status, seis_transaction.Row_changed_by, seis_transaction.Row_changed_date, seis_transaction.Row_created_by, seis_transaction.Row_effective_date, seis_transaction.Row_expiry_date, seis_transaction.Row_quality, seis_transaction.Seis_transaction_id)
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

func PatchSeisTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_transaction set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "transaction_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where seis_transaction_id = :id"

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

func DeleteSeisTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_transaction dto.Seis_transaction
	seis_transaction.Seis_transaction_id = id

	stmt, err := db.Prepare("delete from seis_transaction where seis_transaction_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_transaction.Seis_transaction_id)
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


