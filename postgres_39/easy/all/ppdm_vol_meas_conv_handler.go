package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmVolMeasConv(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_vol_meas_conv")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_vol_meas_conv

	for rows.Next() {
		var ppdm_vol_meas_conv dto.Ppdm_vol_meas_conv
		if err := rows.Scan(&ppdm_vol_meas_conv.Volume_regime_id, &ppdm_vol_meas_conv.Conversion_quantity, &ppdm_vol_meas_conv.Conversion_obs_no, &ppdm_vol_meas_conv.Active_ind, &ppdm_vol_meas_conv.Conversion_factor, &ppdm_vol_meas_conv.Conversion_formula, &ppdm_vol_meas_conv.Effective_date, &ppdm_vol_meas_conv.Expiry_date, &ppdm_vol_meas_conv.From_uom, &ppdm_vol_meas_conv.Ppdm_guid, &ppdm_vol_meas_conv.Preferrred_ind, &ppdm_vol_meas_conv.Pressure, &ppdm_vol_meas_conv.Pressure_uom, &ppdm_vol_meas_conv.Remark, &ppdm_vol_meas_conv.Source, &ppdm_vol_meas_conv.Temperature, &ppdm_vol_meas_conv.Temperature_uom, &ppdm_vol_meas_conv.To_uom, &ppdm_vol_meas_conv.Row_changed_by, &ppdm_vol_meas_conv.Row_changed_date, &ppdm_vol_meas_conv.Row_created_by, &ppdm_vol_meas_conv.Row_created_date, &ppdm_vol_meas_conv.Row_effective_date, &ppdm_vol_meas_conv.Row_expiry_date, &ppdm_vol_meas_conv.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_vol_meas_conv to result
		result = append(result, ppdm_vol_meas_conv)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmVolMeasConv(c *fiber.Ctx) error {
	var ppdm_vol_meas_conv dto.Ppdm_vol_meas_conv

	setDefaults(&ppdm_vol_meas_conv)

	if err := json.Unmarshal(c.Body(), &ppdm_vol_meas_conv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_vol_meas_conv values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	ppdm_vol_meas_conv.Row_created_date = formatDate(ppdm_vol_meas_conv.Row_created_date)
	ppdm_vol_meas_conv.Row_changed_date = formatDate(ppdm_vol_meas_conv.Row_changed_date)
	ppdm_vol_meas_conv.Effective_date = formatDateString(ppdm_vol_meas_conv.Effective_date)
	ppdm_vol_meas_conv.Expiry_date = formatDateString(ppdm_vol_meas_conv.Expiry_date)
	ppdm_vol_meas_conv.Row_effective_date = formatDateString(ppdm_vol_meas_conv.Row_effective_date)
	ppdm_vol_meas_conv.Row_expiry_date = formatDateString(ppdm_vol_meas_conv.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_vol_meas_conv.Volume_regime_id, ppdm_vol_meas_conv.Conversion_quantity, ppdm_vol_meas_conv.Conversion_obs_no, ppdm_vol_meas_conv.Active_ind, ppdm_vol_meas_conv.Conversion_factor, ppdm_vol_meas_conv.Conversion_formula, ppdm_vol_meas_conv.Effective_date, ppdm_vol_meas_conv.Expiry_date, ppdm_vol_meas_conv.From_uom, ppdm_vol_meas_conv.Ppdm_guid, ppdm_vol_meas_conv.Preferrred_ind, ppdm_vol_meas_conv.Pressure, ppdm_vol_meas_conv.Pressure_uom, ppdm_vol_meas_conv.Remark, ppdm_vol_meas_conv.Source, ppdm_vol_meas_conv.Temperature, ppdm_vol_meas_conv.Temperature_uom, ppdm_vol_meas_conv.To_uom, ppdm_vol_meas_conv.Row_changed_by, ppdm_vol_meas_conv.Row_changed_date, ppdm_vol_meas_conv.Row_created_by, ppdm_vol_meas_conv.Row_created_date, ppdm_vol_meas_conv.Row_effective_date, ppdm_vol_meas_conv.Row_expiry_date, ppdm_vol_meas_conv.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmVolMeasConv(c *fiber.Ctx) error {
	var ppdm_vol_meas_conv dto.Ppdm_vol_meas_conv

	setDefaults(&ppdm_vol_meas_conv)

	if err := json.Unmarshal(c.Body(), &ppdm_vol_meas_conv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_vol_meas_conv.Volume_regime_id = id

    if ppdm_vol_meas_conv.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_vol_meas_conv set conversion_quantity = :1, conversion_obs_no = :2, active_ind = :3, conversion_factor = :4, conversion_formula = :5, effective_date = :6, expiry_date = :7, from_uom = :8, ppdm_guid = :9, preferrred_ind = :10, pressure = :11, pressure_uom = :12, remark = :13, source = :14, temperature = :15, temperature_uom = :16, to_uom = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where volume_regime_id = :25")
	if err != nil {
		return err
	}

	ppdm_vol_meas_conv.Row_changed_date = formatDate(ppdm_vol_meas_conv.Row_changed_date)
	ppdm_vol_meas_conv.Effective_date = formatDateString(ppdm_vol_meas_conv.Effective_date)
	ppdm_vol_meas_conv.Expiry_date = formatDateString(ppdm_vol_meas_conv.Expiry_date)
	ppdm_vol_meas_conv.Row_effective_date = formatDateString(ppdm_vol_meas_conv.Row_effective_date)
	ppdm_vol_meas_conv.Row_expiry_date = formatDateString(ppdm_vol_meas_conv.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_vol_meas_conv.Conversion_quantity, ppdm_vol_meas_conv.Conversion_obs_no, ppdm_vol_meas_conv.Active_ind, ppdm_vol_meas_conv.Conversion_factor, ppdm_vol_meas_conv.Conversion_formula, ppdm_vol_meas_conv.Effective_date, ppdm_vol_meas_conv.Expiry_date, ppdm_vol_meas_conv.From_uom, ppdm_vol_meas_conv.Ppdm_guid, ppdm_vol_meas_conv.Preferrred_ind, ppdm_vol_meas_conv.Pressure, ppdm_vol_meas_conv.Pressure_uom, ppdm_vol_meas_conv.Remark, ppdm_vol_meas_conv.Source, ppdm_vol_meas_conv.Temperature, ppdm_vol_meas_conv.Temperature_uom, ppdm_vol_meas_conv.To_uom, ppdm_vol_meas_conv.Row_changed_by, ppdm_vol_meas_conv.Row_changed_date, ppdm_vol_meas_conv.Row_created_by, ppdm_vol_meas_conv.Row_effective_date, ppdm_vol_meas_conv.Row_expiry_date, ppdm_vol_meas_conv.Row_quality, ppdm_vol_meas_conv.Volume_regime_id)
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

func PatchPpdmVolMeasConv(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_vol_meas_conv set "
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
	query += " where volume_regime_id = :id"

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

func DeletePpdmVolMeasConv(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_vol_meas_conv dto.Ppdm_vol_meas_conv
	ppdm_vol_meas_conv.Volume_regime_id = id

	stmt, err := db.Prepare("delete from ppdm_vol_meas_conv where volume_regime_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_vol_meas_conv.Volume_regime_id)
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


