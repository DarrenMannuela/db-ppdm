package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligPaymentRate(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_payment_rate")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_payment_rate

	for rows.Next() {
		var oblig_payment_rate dto.Oblig_payment_rate
		if err := rows.Scan(&oblig_payment_rate.Jurisdiction, &oblig_payment_rate.Pay_rate_type, &oblig_payment_rate.Payment_rate_id, &oblig_payment_rate.Active_ind, &oblig_payment_rate.Contract_id, &oblig_payment_rate.Effective_date, &oblig_payment_rate.Expiry_date, &oblig_payment_rate.Payment_rate, &oblig_payment_rate.Payment_rate_ouom, &oblig_payment_rate.Ppdm_guid, &oblig_payment_rate.Rate_percent, &oblig_payment_rate.Remark, &oblig_payment_rate.Source, &oblig_payment_rate.Row_changed_by, &oblig_payment_rate.Row_changed_date, &oblig_payment_rate.Row_created_by, &oblig_payment_rate.Row_created_date, &oblig_payment_rate.Row_effective_date, &oblig_payment_rate.Row_expiry_date, &oblig_payment_rate.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_payment_rate to result
		result = append(result, oblig_payment_rate)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligPaymentRate(c *fiber.Ctx) error {
	var oblig_payment_rate dto.Oblig_payment_rate

	setDefaults(&oblig_payment_rate)

	if err := json.Unmarshal(c.Body(), &oblig_payment_rate); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_payment_rate values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	oblig_payment_rate.Row_created_date = formatDate(oblig_payment_rate.Row_created_date)
	oblig_payment_rate.Row_changed_date = formatDate(oblig_payment_rate.Row_changed_date)
	oblig_payment_rate.Effective_date = formatDateString(oblig_payment_rate.Effective_date)
	oblig_payment_rate.Expiry_date = formatDateString(oblig_payment_rate.Expiry_date)
	oblig_payment_rate.Row_effective_date = formatDateString(oblig_payment_rate.Row_effective_date)
	oblig_payment_rate.Row_expiry_date = formatDateString(oblig_payment_rate.Row_expiry_date)






	rows, err := stmt.Exec(oblig_payment_rate.Jurisdiction, oblig_payment_rate.Pay_rate_type, oblig_payment_rate.Payment_rate_id, oblig_payment_rate.Active_ind, oblig_payment_rate.Contract_id, oblig_payment_rate.Effective_date, oblig_payment_rate.Expiry_date, oblig_payment_rate.Payment_rate, oblig_payment_rate.Payment_rate_ouom, oblig_payment_rate.Ppdm_guid, oblig_payment_rate.Rate_percent, oblig_payment_rate.Remark, oblig_payment_rate.Source, oblig_payment_rate.Row_changed_by, oblig_payment_rate.Row_changed_date, oblig_payment_rate.Row_created_by, oblig_payment_rate.Row_created_date, oblig_payment_rate.Row_effective_date, oblig_payment_rate.Row_expiry_date, oblig_payment_rate.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligPaymentRate(c *fiber.Ctx) error {
	var oblig_payment_rate dto.Oblig_payment_rate

	setDefaults(&oblig_payment_rate)

	if err := json.Unmarshal(c.Body(), &oblig_payment_rate); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_payment_rate.Jurisdiction = id

    if oblig_payment_rate.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_payment_rate set pay_rate_type = :1, payment_rate_id = :2, active_ind = :3, contract_id = :4, effective_date = :5, expiry_date = :6, payment_rate = :7, payment_rate_ouom = :8, ppdm_guid = :9, rate_percent = :10, remark = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where jurisdiction = :20")
	if err != nil {
		return err
	}

	oblig_payment_rate.Row_changed_date = formatDate(oblig_payment_rate.Row_changed_date)
	oblig_payment_rate.Effective_date = formatDateString(oblig_payment_rate.Effective_date)
	oblig_payment_rate.Expiry_date = formatDateString(oblig_payment_rate.Expiry_date)
	oblig_payment_rate.Row_effective_date = formatDateString(oblig_payment_rate.Row_effective_date)
	oblig_payment_rate.Row_expiry_date = formatDateString(oblig_payment_rate.Row_expiry_date)






	rows, err := stmt.Exec(oblig_payment_rate.Pay_rate_type, oblig_payment_rate.Payment_rate_id, oblig_payment_rate.Active_ind, oblig_payment_rate.Contract_id, oblig_payment_rate.Effective_date, oblig_payment_rate.Expiry_date, oblig_payment_rate.Payment_rate, oblig_payment_rate.Payment_rate_ouom, oblig_payment_rate.Ppdm_guid, oblig_payment_rate.Rate_percent, oblig_payment_rate.Remark, oblig_payment_rate.Source, oblig_payment_rate.Row_changed_by, oblig_payment_rate.Row_changed_date, oblig_payment_rate.Row_created_by, oblig_payment_rate.Row_effective_date, oblig_payment_rate.Row_expiry_date, oblig_payment_rate.Row_quality, oblig_payment_rate.Jurisdiction)
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

func PatchObligPaymentRate(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_payment_rate set "
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
	query += " where jurisdiction = :id"

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

func DeleteObligPaymentRate(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_payment_rate dto.Oblig_payment_rate
	oblig_payment_rate.Jurisdiction = id

	stmt, err := db.Prepare("delete from oblig_payment_rate where jurisdiction = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_payment_rate.Jurisdiction)
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


