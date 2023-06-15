package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmVolMeasUse(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_vol_meas_use")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_vol_meas_use

	for rows.Next() {
		var ppdm_vol_meas_use dto.Ppdm_vol_meas_use
		if err := rows.Scan(&ppdm_vol_meas_use.Volume_regime_id, &ppdm_vol_meas_use.Regime_use_obs_no, &ppdm_vol_meas_use.Active_ind, &ppdm_vol_meas_use.Area_id, &ppdm_vol_meas_use.Area_type, &ppdm_vol_meas_use.Effective_date, &ppdm_vol_meas_use.Expiry_date, &ppdm_vol_meas_use.Jurisdiction, &ppdm_vol_meas_use.Ppdm_guid, &ppdm_vol_meas_use.Preferred_ind, &ppdm_vol_meas_use.Remark, &ppdm_vol_meas_use.Source, &ppdm_vol_meas_use.Row_changed_by, &ppdm_vol_meas_use.Row_changed_date, &ppdm_vol_meas_use.Row_created_by, &ppdm_vol_meas_use.Row_created_date, &ppdm_vol_meas_use.Row_effective_date, &ppdm_vol_meas_use.Row_expiry_date, &ppdm_vol_meas_use.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_vol_meas_use to result
		result = append(result, ppdm_vol_meas_use)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmVolMeasUse(c *fiber.Ctx) error {
	var ppdm_vol_meas_use dto.Ppdm_vol_meas_use

	setDefaults(&ppdm_vol_meas_use)

	if err := json.Unmarshal(c.Body(), &ppdm_vol_meas_use); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_vol_meas_use values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	ppdm_vol_meas_use.Row_created_date = formatDate(ppdm_vol_meas_use.Row_created_date)
	ppdm_vol_meas_use.Row_changed_date = formatDate(ppdm_vol_meas_use.Row_changed_date)
	ppdm_vol_meas_use.Effective_date = formatDateString(ppdm_vol_meas_use.Effective_date)
	ppdm_vol_meas_use.Expiry_date = formatDateString(ppdm_vol_meas_use.Expiry_date)
	ppdm_vol_meas_use.Row_effective_date = formatDateString(ppdm_vol_meas_use.Row_effective_date)
	ppdm_vol_meas_use.Row_expiry_date = formatDateString(ppdm_vol_meas_use.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_vol_meas_use.Volume_regime_id, ppdm_vol_meas_use.Regime_use_obs_no, ppdm_vol_meas_use.Active_ind, ppdm_vol_meas_use.Area_id, ppdm_vol_meas_use.Area_type, ppdm_vol_meas_use.Effective_date, ppdm_vol_meas_use.Expiry_date, ppdm_vol_meas_use.Jurisdiction, ppdm_vol_meas_use.Ppdm_guid, ppdm_vol_meas_use.Preferred_ind, ppdm_vol_meas_use.Remark, ppdm_vol_meas_use.Source, ppdm_vol_meas_use.Row_changed_by, ppdm_vol_meas_use.Row_changed_date, ppdm_vol_meas_use.Row_created_by, ppdm_vol_meas_use.Row_created_date, ppdm_vol_meas_use.Row_effective_date, ppdm_vol_meas_use.Row_expiry_date, ppdm_vol_meas_use.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmVolMeasUse(c *fiber.Ctx) error {
	var ppdm_vol_meas_use dto.Ppdm_vol_meas_use

	setDefaults(&ppdm_vol_meas_use)

	if err := json.Unmarshal(c.Body(), &ppdm_vol_meas_use); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_vol_meas_use.Volume_regime_id = id

    if ppdm_vol_meas_use.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_vol_meas_use set regime_use_obs_no = :1, active_ind = :2, area_id = :3, area_type = :4, effective_date = :5, expiry_date = :6, jurisdiction = :7, ppdm_guid = :8, preferred_ind = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where volume_regime_id = :19")
	if err != nil {
		return err
	}

	ppdm_vol_meas_use.Row_changed_date = formatDate(ppdm_vol_meas_use.Row_changed_date)
	ppdm_vol_meas_use.Effective_date = formatDateString(ppdm_vol_meas_use.Effective_date)
	ppdm_vol_meas_use.Expiry_date = formatDateString(ppdm_vol_meas_use.Expiry_date)
	ppdm_vol_meas_use.Row_effective_date = formatDateString(ppdm_vol_meas_use.Row_effective_date)
	ppdm_vol_meas_use.Row_expiry_date = formatDateString(ppdm_vol_meas_use.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_vol_meas_use.Regime_use_obs_no, ppdm_vol_meas_use.Active_ind, ppdm_vol_meas_use.Area_id, ppdm_vol_meas_use.Area_type, ppdm_vol_meas_use.Effective_date, ppdm_vol_meas_use.Expiry_date, ppdm_vol_meas_use.Jurisdiction, ppdm_vol_meas_use.Ppdm_guid, ppdm_vol_meas_use.Preferred_ind, ppdm_vol_meas_use.Remark, ppdm_vol_meas_use.Source, ppdm_vol_meas_use.Row_changed_by, ppdm_vol_meas_use.Row_changed_date, ppdm_vol_meas_use.Row_created_by, ppdm_vol_meas_use.Row_effective_date, ppdm_vol_meas_use.Row_expiry_date, ppdm_vol_meas_use.Row_quality, ppdm_vol_meas_use.Volume_regime_id)
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

func PatchPpdmVolMeasUse(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_vol_meas_use set "
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

func DeletePpdmVolMeasUse(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_vol_meas_use dto.Ppdm_vol_meas_use
	ppdm_vol_meas_use.Volume_regime_id = id

	stmt, err := db.Prepare("delete from ppdm_vol_meas_use where volume_regime_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_vol_meas_use.Volume_regime_id)
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


