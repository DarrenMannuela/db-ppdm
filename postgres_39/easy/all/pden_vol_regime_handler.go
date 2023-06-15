package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenVolRegime(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_vol_regime")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_vol_regime

	for rows.Next() {
		var pden_vol_regime dto.Pden_vol_regime
		if err := rows.Scan(&pden_vol_regime.Pden_subtype, &pden_vol_regime.Pden_id, &pden_vol_regime.Pden_source, &pden_vol_regime.Volume_regime_id, &pden_vol_regime.Pden_regime_obs_no, &pden_vol_regime.Active_ind, &pden_vol_regime.Effective_date, &pden_vol_regime.Expiry_date, &pden_vol_regime.Ppdm_guid, &pden_vol_regime.Preferred_ind, &pden_vol_regime.Remark, &pden_vol_regime.Source, &pden_vol_regime.Row_changed_by, &pden_vol_regime.Row_changed_date, &pden_vol_regime.Row_created_by, &pden_vol_regime.Row_created_date, &pden_vol_regime.Row_effective_date, &pden_vol_regime.Row_expiry_date, &pden_vol_regime.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_vol_regime to result
		result = append(result, pden_vol_regime)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenVolRegime(c *fiber.Ctx) error {
	var pden_vol_regime dto.Pden_vol_regime

	setDefaults(&pden_vol_regime)

	if err := json.Unmarshal(c.Body(), &pden_vol_regime); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_vol_regime values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	pden_vol_regime.Row_created_date = formatDate(pden_vol_regime.Row_created_date)
	pden_vol_regime.Row_changed_date = formatDate(pden_vol_regime.Row_changed_date)
	pden_vol_regime.Effective_date = formatDateString(pden_vol_regime.Effective_date)
	pden_vol_regime.Expiry_date = formatDateString(pden_vol_regime.Expiry_date)
	pden_vol_regime.Row_effective_date = formatDateString(pden_vol_regime.Row_effective_date)
	pden_vol_regime.Row_expiry_date = formatDateString(pden_vol_regime.Row_expiry_date)






	rows, err := stmt.Exec(pden_vol_regime.Pden_subtype, pden_vol_regime.Pden_id, pden_vol_regime.Pden_source, pden_vol_regime.Volume_regime_id, pden_vol_regime.Pden_regime_obs_no, pden_vol_regime.Active_ind, pden_vol_regime.Effective_date, pden_vol_regime.Expiry_date, pden_vol_regime.Ppdm_guid, pden_vol_regime.Preferred_ind, pden_vol_regime.Remark, pden_vol_regime.Source, pden_vol_regime.Row_changed_by, pden_vol_regime.Row_changed_date, pden_vol_regime.Row_created_by, pden_vol_regime.Row_created_date, pden_vol_regime.Row_effective_date, pden_vol_regime.Row_expiry_date, pden_vol_regime.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenVolRegime(c *fiber.Ctx) error {
	var pden_vol_regime dto.Pden_vol_regime

	setDefaults(&pden_vol_regime)

	if err := json.Unmarshal(c.Body(), &pden_vol_regime); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_vol_regime.Pden_subtype = id

    if pden_vol_regime.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_vol_regime set pden_id = :1, pden_source = :2, volume_regime_id = :3, pden_regime_obs_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, preferred_ind = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where pden_subtype = :19")
	if err != nil {
		return err
	}

	pden_vol_regime.Row_changed_date = formatDate(pden_vol_regime.Row_changed_date)
	pden_vol_regime.Effective_date = formatDateString(pden_vol_regime.Effective_date)
	pden_vol_regime.Expiry_date = formatDateString(pden_vol_regime.Expiry_date)
	pden_vol_regime.Row_effective_date = formatDateString(pden_vol_regime.Row_effective_date)
	pden_vol_regime.Row_expiry_date = formatDateString(pden_vol_regime.Row_expiry_date)






	rows, err := stmt.Exec(pden_vol_regime.Pden_id, pden_vol_regime.Pden_source, pden_vol_regime.Volume_regime_id, pden_vol_regime.Pden_regime_obs_no, pden_vol_regime.Active_ind, pden_vol_regime.Effective_date, pden_vol_regime.Expiry_date, pden_vol_regime.Ppdm_guid, pden_vol_regime.Preferred_ind, pden_vol_regime.Remark, pden_vol_regime.Source, pden_vol_regime.Row_changed_by, pden_vol_regime.Row_changed_date, pden_vol_regime.Row_created_by, pden_vol_regime.Row_effective_date, pden_vol_regime.Row_expiry_date, pden_vol_regime.Row_quality, pden_vol_regime.Pden_subtype)
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

func PatchPdenVolRegime(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_vol_regime set "
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
	query += " where pden_subtype = :id"

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

func DeletePdenVolRegime(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_vol_regime dto.Pden_vol_regime
	pden_vol_regime.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_vol_regime where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_vol_regime.Pden_subtype)
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


