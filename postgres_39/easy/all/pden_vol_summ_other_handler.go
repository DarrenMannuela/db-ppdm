package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenVolSummOther(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_vol_summ_other")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_vol_summ_other

	for rows.Next() {
		var pden_vol_summ_other dto.Pden_vol_summ_other
		if err := rows.Scan(&pden_vol_summ_other.Pden_subtype, &pden_vol_summ_other.Pden_id, &pden_vol_summ_other.Pden_source, &pden_vol_summ_other.Volume_method, &pden_vol_summ_other.Activity_type, &pden_vol_summ_other.Period_type, &pden_vol_summ_other.Pden_period_id, &pden_vol_summ_other.Amendment_seq_no, &pden_vol_summ_other.Product_type, &pden_vol_summ_other.Active_ind, &pden_vol_summ_other.Cum_volume, &pden_vol_summ_other.Effective_date, &pden_vol_summ_other.Expiry_date, &pden_vol_summ_other.Gas_energy, &pden_vol_summ_other.Gas_energy_ouom, &pden_vol_summ_other.Mass, &pden_vol_summ_other.Mass_ouom, &pden_vol_summ_other.Ppdm_guid, &pden_vol_summ_other.Remark, &pden_vol_summ_other.Report_ind, &pden_vol_summ_other.Source, &pden_vol_summ_other.Volume, &pden_vol_summ_other.Volume_date, &pden_vol_summ_other.Volume_date_desc, &pden_vol_summ_other.Volume_ouom, &pden_vol_summ_other.Volume_quality, &pden_vol_summ_other.Volume_quality_ouom, &pden_vol_summ_other.Volume_uom, &pden_vol_summ_other.Ytd_volume, &pden_vol_summ_other.Row_changed_by, &pden_vol_summ_other.Row_changed_date, &pden_vol_summ_other.Row_created_by, &pden_vol_summ_other.Row_created_date, &pden_vol_summ_other.Row_effective_date, &pden_vol_summ_other.Row_expiry_date, &pden_vol_summ_other.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_vol_summ_other to result
		result = append(result, pden_vol_summ_other)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenVolSummOther(c *fiber.Ctx) error {
	var pden_vol_summ_other dto.Pden_vol_summ_other

	setDefaults(&pden_vol_summ_other)

	if err := json.Unmarshal(c.Body(), &pden_vol_summ_other); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_vol_summ_other values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36)")
	if err != nil {
		return err
	}
	pden_vol_summ_other.Row_created_date = formatDate(pden_vol_summ_other.Row_created_date)
	pden_vol_summ_other.Row_changed_date = formatDate(pden_vol_summ_other.Row_changed_date)
	pden_vol_summ_other.Effective_date = formatDateString(pden_vol_summ_other.Effective_date)
	pden_vol_summ_other.Expiry_date = formatDateString(pden_vol_summ_other.Expiry_date)
	pden_vol_summ_other.Volume_date = formatDateString(pden_vol_summ_other.Volume_date)
	pden_vol_summ_other.Row_effective_date = formatDateString(pden_vol_summ_other.Row_effective_date)
	pden_vol_summ_other.Row_expiry_date = formatDateString(pden_vol_summ_other.Row_expiry_date)






	rows, err := stmt.Exec(pden_vol_summ_other.Pden_subtype, pden_vol_summ_other.Pden_id, pden_vol_summ_other.Pden_source, pden_vol_summ_other.Volume_method, pden_vol_summ_other.Activity_type, pden_vol_summ_other.Period_type, pden_vol_summ_other.Pden_period_id, pden_vol_summ_other.Amendment_seq_no, pden_vol_summ_other.Product_type, pden_vol_summ_other.Active_ind, pden_vol_summ_other.Cum_volume, pden_vol_summ_other.Effective_date, pden_vol_summ_other.Expiry_date, pden_vol_summ_other.Gas_energy, pden_vol_summ_other.Gas_energy_ouom, pden_vol_summ_other.Mass, pden_vol_summ_other.Mass_ouom, pden_vol_summ_other.Ppdm_guid, pden_vol_summ_other.Remark, pden_vol_summ_other.Report_ind, pden_vol_summ_other.Source, pden_vol_summ_other.Volume, pden_vol_summ_other.Volume_date, pden_vol_summ_other.Volume_date_desc, pden_vol_summ_other.Volume_ouom, pden_vol_summ_other.Volume_quality, pden_vol_summ_other.Volume_quality_ouom, pden_vol_summ_other.Volume_uom, pden_vol_summ_other.Ytd_volume, pden_vol_summ_other.Row_changed_by, pden_vol_summ_other.Row_changed_date, pden_vol_summ_other.Row_created_by, pden_vol_summ_other.Row_created_date, pden_vol_summ_other.Row_effective_date, pden_vol_summ_other.Row_expiry_date, pden_vol_summ_other.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenVolSummOther(c *fiber.Ctx) error {
	var pden_vol_summ_other dto.Pden_vol_summ_other

	setDefaults(&pden_vol_summ_other)

	if err := json.Unmarshal(c.Body(), &pden_vol_summ_other); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_vol_summ_other.Pden_subtype = id

    if pden_vol_summ_other.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_vol_summ_other set pden_id = :1, pden_source = :2, volume_method = :3, activity_type = :4, period_type = :5, pden_period_id = :6, amendment_seq_no = :7, product_type = :8, active_ind = :9, cum_volume = :10, effective_date = :11, expiry_date = :12, gas_energy = :13, gas_energy_ouom = :14, mass = :15, mass_ouom = :16, ppdm_guid = :17, remark = :18, report_ind = :19, source = :20, volume = :21, volume_date = :22, volume_date_desc = :23, volume_ouom = :24, volume_quality = :25, volume_quality_ouom = :26, volume_uom = :27, ytd_volume = :28, row_changed_by = :29, row_changed_date = :30, row_created_by = :31, row_effective_date = :32, row_expiry_date = :33, row_quality = :34 where pden_subtype = :36")
	if err != nil {
		return err
	}

	pden_vol_summ_other.Row_changed_date = formatDate(pden_vol_summ_other.Row_changed_date)
	pden_vol_summ_other.Effective_date = formatDateString(pden_vol_summ_other.Effective_date)
	pden_vol_summ_other.Expiry_date = formatDateString(pden_vol_summ_other.Expiry_date)
	pden_vol_summ_other.Volume_date = formatDateString(pden_vol_summ_other.Volume_date)
	pden_vol_summ_other.Row_effective_date = formatDateString(pden_vol_summ_other.Row_effective_date)
	pden_vol_summ_other.Row_expiry_date = formatDateString(pden_vol_summ_other.Row_expiry_date)






	rows, err := stmt.Exec(pden_vol_summ_other.Pden_id, pden_vol_summ_other.Pden_source, pden_vol_summ_other.Volume_method, pden_vol_summ_other.Activity_type, pden_vol_summ_other.Period_type, pden_vol_summ_other.Pden_period_id, pden_vol_summ_other.Amendment_seq_no, pden_vol_summ_other.Product_type, pden_vol_summ_other.Active_ind, pden_vol_summ_other.Cum_volume, pden_vol_summ_other.Effective_date, pden_vol_summ_other.Expiry_date, pden_vol_summ_other.Gas_energy, pden_vol_summ_other.Gas_energy_ouom, pden_vol_summ_other.Mass, pden_vol_summ_other.Mass_ouom, pden_vol_summ_other.Ppdm_guid, pden_vol_summ_other.Remark, pden_vol_summ_other.Report_ind, pden_vol_summ_other.Source, pden_vol_summ_other.Volume, pden_vol_summ_other.Volume_date, pden_vol_summ_other.Volume_date_desc, pden_vol_summ_other.Volume_ouom, pden_vol_summ_other.Volume_quality, pden_vol_summ_other.Volume_quality_ouom, pden_vol_summ_other.Volume_uom, pden_vol_summ_other.Ytd_volume, pden_vol_summ_other.Row_changed_by, pden_vol_summ_other.Row_changed_date, pden_vol_summ_other.Row_created_by, pden_vol_summ_other.Row_effective_date, pden_vol_summ_other.Row_expiry_date, pden_vol_summ_other.Row_quality, pden_vol_summ_other.Pden_subtype)
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

func PatchPdenVolSummOther(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_vol_summ_other set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "volume_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePdenVolSummOther(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_vol_summ_other dto.Pden_vol_summ_other
	pden_vol_summ_other.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_vol_summ_other where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_vol_summ_other.Pden_subtype)
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


