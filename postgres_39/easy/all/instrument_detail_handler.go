package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetInstrumentDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from instrument_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Instrument_detail

	for rows.Next() {
		var instrument_detail dto.Instrument_detail
		if err := rows.Scan(&instrument_detail.Instrument_id, &instrument_detail.Detail_id, &instrument_detail.Active_ind, &instrument_detail.Average_value, &instrument_detail.Average_value_ouom, &instrument_detail.Average_value_uom, &instrument_detail.Cost_value, &instrument_detail.Currency_conversion, &instrument_detail.Currency_ouom, &instrument_detail.Currency_uom, &instrument_detail.Detail_description, &instrument_detail.Effective_date, &instrument_detail.Expiry_date, &instrument_detail.Instrument_detail_code, &instrument_detail.Instrument_detail_type, &instrument_detail.Max_date, &instrument_detail.Max_value, &instrument_detail.Max_value_ouom, &instrument_detail.Max_value_uom, &instrument_detail.Min_date, &instrument_detail.Min_value, &instrument_detail.Min_value_ouom, &instrument_detail.Min_value_uom, &instrument_detail.Ppdm_guid, &instrument_detail.Reference_value, &instrument_detail.Reference_value_ouom, &instrument_detail.Reference_value_type, &instrument_detail.Reference_value_uom, &instrument_detail.Remark, &instrument_detail.Source, &instrument_detail.Row_changed_by, &instrument_detail.Row_changed_date, &instrument_detail.Row_created_by, &instrument_detail.Row_created_date, &instrument_detail.Row_effective_date, &instrument_detail.Row_expiry_date, &instrument_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append instrument_detail to result
		result = append(result, instrument_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetInstrumentDetail(c *fiber.Ctx) error {
	var instrument_detail dto.Instrument_detail

	setDefaults(&instrument_detail)

	if err := json.Unmarshal(c.Body(), &instrument_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into instrument_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	instrument_detail.Row_created_date = formatDate(instrument_detail.Row_created_date)
	instrument_detail.Row_changed_date = formatDate(instrument_detail.Row_changed_date)
	instrument_detail.Effective_date = formatDateString(instrument_detail.Effective_date)
	instrument_detail.Expiry_date = formatDateString(instrument_detail.Expiry_date)
	instrument_detail.Max_date = formatDateString(instrument_detail.Max_date)
	instrument_detail.Min_date = formatDateString(instrument_detail.Min_date)
	instrument_detail.Row_effective_date = formatDateString(instrument_detail.Row_effective_date)
	instrument_detail.Row_expiry_date = formatDateString(instrument_detail.Row_expiry_date)






	rows, err := stmt.Exec(instrument_detail.Instrument_id, instrument_detail.Detail_id, instrument_detail.Active_ind, instrument_detail.Average_value, instrument_detail.Average_value_ouom, instrument_detail.Average_value_uom, instrument_detail.Cost_value, instrument_detail.Currency_conversion, instrument_detail.Currency_ouom, instrument_detail.Currency_uom, instrument_detail.Detail_description, instrument_detail.Effective_date, instrument_detail.Expiry_date, instrument_detail.Instrument_detail_code, instrument_detail.Instrument_detail_type, instrument_detail.Max_date, instrument_detail.Max_value, instrument_detail.Max_value_ouom, instrument_detail.Max_value_uom, instrument_detail.Min_date, instrument_detail.Min_value, instrument_detail.Min_value_ouom, instrument_detail.Min_value_uom, instrument_detail.Ppdm_guid, instrument_detail.Reference_value, instrument_detail.Reference_value_ouom, instrument_detail.Reference_value_type, instrument_detail.Reference_value_uom, instrument_detail.Remark, instrument_detail.Source, instrument_detail.Row_changed_by, instrument_detail.Row_changed_date, instrument_detail.Row_created_by, instrument_detail.Row_created_date, instrument_detail.Row_effective_date, instrument_detail.Row_expiry_date, instrument_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateInstrumentDetail(c *fiber.Ctx) error {
	var instrument_detail dto.Instrument_detail

	setDefaults(&instrument_detail)

	if err := json.Unmarshal(c.Body(), &instrument_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	instrument_detail.Instrument_id = id

    if instrument_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update instrument_detail set detail_id = :1, active_ind = :2, average_value = :3, average_value_ouom = :4, average_value_uom = :5, cost_value = :6, currency_conversion = :7, currency_ouom = :8, currency_uom = :9, detail_description = :10, effective_date = :11, expiry_date = :12, instrument_detail_code = :13, instrument_detail_type = :14, max_date = :15, max_value = :16, max_value_ouom = :17, max_value_uom = :18, min_date = :19, min_value = :20, min_value_ouom = :21, min_value_uom = :22, ppdm_guid = :23, reference_value = :24, reference_value_ouom = :25, reference_value_type = :26, reference_value_uom = :27, remark = :28, source = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where instrument_id = :37")
	if err != nil {
		return err
	}

	instrument_detail.Row_changed_date = formatDate(instrument_detail.Row_changed_date)
	instrument_detail.Effective_date = formatDateString(instrument_detail.Effective_date)
	instrument_detail.Expiry_date = formatDateString(instrument_detail.Expiry_date)
	instrument_detail.Max_date = formatDateString(instrument_detail.Max_date)
	instrument_detail.Min_date = formatDateString(instrument_detail.Min_date)
	instrument_detail.Row_effective_date = formatDateString(instrument_detail.Row_effective_date)
	instrument_detail.Row_expiry_date = formatDateString(instrument_detail.Row_expiry_date)






	rows, err := stmt.Exec(instrument_detail.Detail_id, instrument_detail.Active_ind, instrument_detail.Average_value, instrument_detail.Average_value_ouom, instrument_detail.Average_value_uom, instrument_detail.Cost_value, instrument_detail.Currency_conversion, instrument_detail.Currency_ouom, instrument_detail.Currency_uom, instrument_detail.Detail_description, instrument_detail.Effective_date, instrument_detail.Expiry_date, instrument_detail.Instrument_detail_code, instrument_detail.Instrument_detail_type, instrument_detail.Max_date, instrument_detail.Max_value, instrument_detail.Max_value_ouom, instrument_detail.Max_value_uom, instrument_detail.Min_date, instrument_detail.Min_value, instrument_detail.Min_value_ouom, instrument_detail.Min_value_uom, instrument_detail.Ppdm_guid, instrument_detail.Reference_value, instrument_detail.Reference_value_ouom, instrument_detail.Reference_value_type, instrument_detail.Reference_value_uom, instrument_detail.Remark, instrument_detail.Source, instrument_detail.Row_changed_by, instrument_detail.Row_changed_date, instrument_detail.Row_created_by, instrument_detail.Row_effective_date, instrument_detail.Row_expiry_date, instrument_detail.Row_quality, instrument_detail.Instrument_id)
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

func PatchInstrumentDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update instrument_detail set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where instrument_id = :id"

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

func DeleteInstrumentDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var instrument_detail dto.Instrument_detail
	instrument_detail.Instrument_id = id

	stmt, err := db.Prepare("delete from instrument_detail where instrument_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(instrument_detail.Instrument_id)
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


