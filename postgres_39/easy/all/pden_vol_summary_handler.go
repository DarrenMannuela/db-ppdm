package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenVolSummary(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_vol_summary")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_vol_summary

	for rows.Next() {
		var pden_vol_summary dto.Pden_vol_summary
		if err := rows.Scan(&pden_vol_summary.Pden_subtype, &pden_vol_summary.Pden_id, &pden_vol_summary.Period_id, &pden_vol_summary.Pden_source, &pden_vol_summary.Volume_method, &pden_vol_summary.Activity_type, &pden_vol_summary.Period_type, &pden_vol_summary.Amendment_seq_no, &pden_vol_summary.Active_ind, &pden_vol_summary.Amend_reason, &pden_vol_summary.Boe_cum_volume, &pden_vol_summary.Boe_volume, &pden_vol_summary.Boe_volume_ouom, &pden_vol_summary.Boe_ytd_volume, &pden_vol_summary.Co2_cum_volume, &pden_vol_summary.Co2_volume, &pden_vol_summary.Co2_volume_ouom, &pden_vol_summary.Co2_ytd_volume, &pden_vol_summary.Effective_date, &pden_vol_summary.Expiry_date, &pden_vol_summary.Gas_cum_volume, &pden_vol_summary.Gas_energy, &pden_vol_summary.Gas_energy_ouom, &pden_vol_summary.Gas_quality, &pden_vol_summary.Gas_quality_ouom, &pden_vol_summary.Gas_volume, &pden_vol_summary.Gas_volume_ouom, &pden_vol_summary.Gas_ytd_volume, &pden_vol_summary.Injection_cycle, &pden_vol_summary.Injection_pressure, &pden_vol_summary.Injection_pressure_ouom, &pden_vol_summary.Inventory_close_balance, &pden_vol_summary.Inventory_open_balance, &pden_vol_summary.Inventory_product, &pden_vol_summary.Invent_close_bal_ouom, &pden_vol_summary.Invent_open_bal_ouom, &pden_vol_summary.Ngl_cum_volume, &pden_vol_summary.Ngl_volume, &pden_vol_summary.Ngl_volume_ouom, &pden_vol_summary.Ngl_ytd_volume, &pden_vol_summary.Nitrogen_cum_volume, &pden_vol_summary.Nitrogen_volume, &pden_vol_summary.Nitrogen_volume_ouom, &pden_vol_summary.Nitrogen_ytd_volume, &pden_vol_summary.No_of_gas_wells, &pden_vol_summary.No_of_injection_wells, &pden_vol_summary.No_of_oil_wells, &pden_vol_summary.Oil_cum_volume, &pden_vol_summary.Oil_quality, &pden_vol_summary.Oil_quality_ouom, &pden_vol_summary.Oil_volume, &pden_vol_summary.Oil_volume_ouom, &pden_vol_summary.Oil_ytd_volume, &pden_vol_summary.Period_on_injection, &pden_vol_summary.Period_on_injection_ouom, &pden_vol_summary.Period_on_production, &pden_vol_summary.Period_on_production_ouom, &pden_vol_summary.Posted_date, &pden_vol_summary.Ppdm_guid, &pden_vol_summary.Primary_allowable, &pden_vol_summary.Primary_allowable_ouom, &pden_vol_summary.Primary_product, &pden_vol_summary.Project_id, &pden_vol_summary.Remark, &pden_vol_summary.Report_ind, &pden_vol_summary.Source, &pden_vol_summary.Sulphur_cum_volume, &pden_vol_summary.Sulphur_volume, &pden_vol_summary.Sulphur_volume_ouom, &pden_vol_summary.Sulphur_ytd_volume, &pden_vol_summary.Volume_date, &pden_vol_summary.Volume_date_desc, &pden_vol_summary.Volume_period, &pden_vol_summary.Volume_period_ouom, &pden_vol_summary.Water_cum_volume, &pden_vol_summary.Water_volume, &pden_vol_summary.Water_volume_ouom, &pden_vol_summary.Water_ytd_volume, &pden_vol_summary.Row_changed_by, &pden_vol_summary.Row_changed_date, &pden_vol_summary.Row_created_by, &pden_vol_summary.Row_created_date, &pden_vol_summary.Row_effective_date, &pden_vol_summary.Row_expiry_date, &pden_vol_summary.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_vol_summary to result
		result = append(result, pden_vol_summary)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenVolSummary(c *fiber.Ctx) error {
	var pden_vol_summary dto.Pden_vol_summary

	setDefaults(&pden_vol_summary)

	if err := json.Unmarshal(c.Body(), &pden_vol_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_vol_summary values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85)")
	if err != nil {
		return err
	}
	pden_vol_summary.Row_created_date = formatDate(pden_vol_summary.Row_created_date)
	pden_vol_summary.Row_changed_date = formatDate(pden_vol_summary.Row_changed_date)
	pden_vol_summary.Effective_date = formatDateString(pden_vol_summary.Effective_date)
	pden_vol_summary.Expiry_date = formatDateString(pden_vol_summary.Expiry_date)
	pden_vol_summary.Posted_date = formatDateString(pden_vol_summary.Posted_date)
	pden_vol_summary.Volume_date = formatDateString(pden_vol_summary.Volume_date)
	pden_vol_summary.Row_effective_date = formatDateString(pden_vol_summary.Row_effective_date)
	pden_vol_summary.Row_expiry_date = formatDateString(pden_vol_summary.Row_expiry_date)






	rows, err := stmt.Exec(pden_vol_summary.Pden_subtype, pden_vol_summary.Pden_id, pden_vol_summary.Period_id, pden_vol_summary.Pden_source, pden_vol_summary.Volume_method, pden_vol_summary.Activity_type, pden_vol_summary.Period_type, pden_vol_summary.Amendment_seq_no, pden_vol_summary.Active_ind, pden_vol_summary.Amend_reason, pden_vol_summary.Boe_cum_volume, pden_vol_summary.Boe_volume, pden_vol_summary.Boe_volume_ouom, pden_vol_summary.Boe_ytd_volume, pden_vol_summary.Co2_cum_volume, pden_vol_summary.Co2_volume, pden_vol_summary.Co2_volume_ouom, pden_vol_summary.Co2_ytd_volume, pden_vol_summary.Effective_date, pden_vol_summary.Expiry_date, pden_vol_summary.Gas_cum_volume, pden_vol_summary.Gas_energy, pden_vol_summary.Gas_energy_ouom, pden_vol_summary.Gas_quality, pden_vol_summary.Gas_quality_ouom, pden_vol_summary.Gas_volume, pden_vol_summary.Gas_volume_ouom, pden_vol_summary.Gas_ytd_volume, pden_vol_summary.Injection_cycle, pden_vol_summary.Injection_pressure, pden_vol_summary.Injection_pressure_ouom, pden_vol_summary.Inventory_close_balance, pden_vol_summary.Inventory_open_balance, pden_vol_summary.Inventory_product, pden_vol_summary.Invent_close_bal_ouom, pden_vol_summary.Invent_open_bal_ouom, pden_vol_summary.Ngl_cum_volume, pden_vol_summary.Ngl_volume, pden_vol_summary.Ngl_volume_ouom, pden_vol_summary.Ngl_ytd_volume, pden_vol_summary.Nitrogen_cum_volume, pden_vol_summary.Nitrogen_volume, pden_vol_summary.Nitrogen_volume_ouom, pden_vol_summary.Nitrogen_ytd_volume, pden_vol_summary.No_of_gas_wells, pden_vol_summary.No_of_injection_wells, pden_vol_summary.No_of_oil_wells, pden_vol_summary.Oil_cum_volume, pden_vol_summary.Oil_quality, pden_vol_summary.Oil_quality_ouom, pden_vol_summary.Oil_volume, pden_vol_summary.Oil_volume_ouom, pden_vol_summary.Oil_ytd_volume, pden_vol_summary.Period_on_injection, pden_vol_summary.Period_on_injection_ouom, pden_vol_summary.Period_on_production, pden_vol_summary.Period_on_production_ouom, pden_vol_summary.Posted_date, pden_vol_summary.Ppdm_guid, pden_vol_summary.Primary_allowable, pden_vol_summary.Primary_allowable_ouom, pden_vol_summary.Primary_product, pden_vol_summary.Project_id, pden_vol_summary.Remark, pden_vol_summary.Report_ind, pden_vol_summary.Source, pden_vol_summary.Sulphur_cum_volume, pden_vol_summary.Sulphur_volume, pden_vol_summary.Sulphur_volume_ouom, pden_vol_summary.Sulphur_ytd_volume, pden_vol_summary.Volume_date, pden_vol_summary.Volume_date_desc, pden_vol_summary.Volume_period, pden_vol_summary.Volume_period_ouom, pden_vol_summary.Water_cum_volume, pden_vol_summary.Water_volume, pden_vol_summary.Water_volume_ouom, pden_vol_summary.Water_ytd_volume, pden_vol_summary.Row_changed_by, pden_vol_summary.Row_changed_date, pden_vol_summary.Row_created_by, pden_vol_summary.Row_created_date, pden_vol_summary.Row_effective_date, pden_vol_summary.Row_expiry_date, pden_vol_summary.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenVolSummary(c *fiber.Ctx) error {
	var pden_vol_summary dto.Pden_vol_summary

	setDefaults(&pden_vol_summary)

	if err := json.Unmarshal(c.Body(), &pden_vol_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_vol_summary.Pden_subtype = id

    if pden_vol_summary.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_vol_summary set pden_id = :1, period_id = :2, pden_source = :3, volume_method = :4, activity_type = :5, period_type = :6, amendment_seq_no = :7, active_ind = :8, amend_reason = :9, boe_cum_volume = :10, boe_volume = :11, boe_volume_ouom = :12, boe_ytd_volume = :13, co2_cum_volume = :14, co2_volume = :15, co2_volume_ouom = :16, co2_ytd_volume = :17, effective_date = :18, expiry_date = :19, gas_cum_volume = :20, gas_energy = :21, gas_energy_ouom = :22, gas_quality = :23, gas_quality_ouom = :24, gas_volume = :25, gas_volume_ouom = :26, gas_ytd_volume = :27, injection_cycle = :28, injection_pressure = :29, injection_pressure_ouom = :30, inventory_close_balance = :31, inventory_open_balance = :32, inventory_product = :33, invent_close_bal_ouom = :34, invent_open_bal_ouom = :35, ngl_cum_volume = :36, ngl_volume = :37, ngl_volume_ouom = :38, ngl_ytd_volume = :39, nitrogen_cum_volume = :40, nitrogen_volume = :41, nitrogen_volume_ouom = :42, nitrogen_ytd_volume = :43, no_of_gas_wells = :44, no_of_injection_wells = :45, no_of_oil_wells = :46, oil_cum_volume = :47, oil_quality = :48, oil_quality_ouom = :49, oil_volume = :50, oil_volume_ouom = :51, oil_ytd_volume = :52, period_on_injection = :53, period_on_injection_ouom = :54, period_on_production = :55, period_on_production_ouom = :56, posted_date = :57, ppdm_guid = :58, primary_allowable = :59, primary_allowable_ouom = :60, primary_product = :61, project_id = :62, remark = :63, report_ind = :64, source = :65, sulphur_cum_volume = :66, sulphur_volume = :67, sulphur_volume_ouom = :68, sulphur_ytd_volume = :69, volume_date = :70, volume_date_desc = :71, volume_period = :72, volume_period_ouom = :73, water_cum_volume = :74, water_volume = :75, water_volume_ouom = :76, water_ytd_volume = :77, row_changed_by = :78, row_changed_date = :79, row_created_by = :80, row_effective_date = :81, row_expiry_date = :82, row_quality = :83 where pden_subtype = :85")
	if err != nil {
		return err
	}

	pden_vol_summary.Row_changed_date = formatDate(pden_vol_summary.Row_changed_date)
	pden_vol_summary.Effective_date = formatDateString(pden_vol_summary.Effective_date)
	pden_vol_summary.Expiry_date = formatDateString(pden_vol_summary.Expiry_date)
	pden_vol_summary.Posted_date = formatDateString(pden_vol_summary.Posted_date)
	pden_vol_summary.Volume_date = formatDateString(pden_vol_summary.Volume_date)
	pden_vol_summary.Row_effective_date = formatDateString(pden_vol_summary.Row_effective_date)
	pden_vol_summary.Row_expiry_date = formatDateString(pden_vol_summary.Row_expiry_date)






	rows, err := stmt.Exec(pden_vol_summary.Pden_id, pden_vol_summary.Period_id, pden_vol_summary.Pden_source, pden_vol_summary.Volume_method, pden_vol_summary.Activity_type, pden_vol_summary.Period_type, pden_vol_summary.Amendment_seq_no, pden_vol_summary.Active_ind, pden_vol_summary.Amend_reason, pden_vol_summary.Boe_cum_volume, pden_vol_summary.Boe_volume, pden_vol_summary.Boe_volume_ouom, pden_vol_summary.Boe_ytd_volume, pden_vol_summary.Co2_cum_volume, pden_vol_summary.Co2_volume, pden_vol_summary.Co2_volume_ouom, pden_vol_summary.Co2_ytd_volume, pden_vol_summary.Effective_date, pden_vol_summary.Expiry_date, pden_vol_summary.Gas_cum_volume, pden_vol_summary.Gas_energy, pden_vol_summary.Gas_energy_ouom, pden_vol_summary.Gas_quality, pden_vol_summary.Gas_quality_ouom, pden_vol_summary.Gas_volume, pden_vol_summary.Gas_volume_ouom, pden_vol_summary.Gas_ytd_volume, pden_vol_summary.Injection_cycle, pden_vol_summary.Injection_pressure, pden_vol_summary.Injection_pressure_ouom, pden_vol_summary.Inventory_close_balance, pden_vol_summary.Inventory_open_balance, pden_vol_summary.Inventory_product, pden_vol_summary.Invent_close_bal_ouom, pden_vol_summary.Invent_open_bal_ouom, pden_vol_summary.Ngl_cum_volume, pden_vol_summary.Ngl_volume, pden_vol_summary.Ngl_volume_ouom, pden_vol_summary.Ngl_ytd_volume, pden_vol_summary.Nitrogen_cum_volume, pden_vol_summary.Nitrogen_volume, pden_vol_summary.Nitrogen_volume_ouom, pden_vol_summary.Nitrogen_ytd_volume, pden_vol_summary.No_of_gas_wells, pden_vol_summary.No_of_injection_wells, pden_vol_summary.No_of_oil_wells, pden_vol_summary.Oil_cum_volume, pden_vol_summary.Oil_quality, pden_vol_summary.Oil_quality_ouom, pden_vol_summary.Oil_volume, pden_vol_summary.Oil_volume_ouom, pden_vol_summary.Oil_ytd_volume, pden_vol_summary.Period_on_injection, pden_vol_summary.Period_on_injection_ouom, pden_vol_summary.Period_on_production, pden_vol_summary.Period_on_production_ouom, pden_vol_summary.Posted_date, pden_vol_summary.Ppdm_guid, pden_vol_summary.Primary_allowable, pden_vol_summary.Primary_allowable_ouom, pden_vol_summary.Primary_product, pden_vol_summary.Project_id, pden_vol_summary.Remark, pden_vol_summary.Report_ind, pden_vol_summary.Source, pden_vol_summary.Sulphur_cum_volume, pden_vol_summary.Sulphur_volume, pden_vol_summary.Sulphur_volume_ouom, pden_vol_summary.Sulphur_ytd_volume, pden_vol_summary.Volume_date, pden_vol_summary.Volume_date_desc, pden_vol_summary.Volume_period, pden_vol_summary.Volume_period_ouom, pden_vol_summary.Water_cum_volume, pden_vol_summary.Water_volume, pden_vol_summary.Water_volume_ouom, pden_vol_summary.Water_ytd_volume, pden_vol_summary.Row_changed_by, pden_vol_summary.Row_changed_date, pden_vol_summary.Row_created_by, pden_vol_summary.Row_effective_date, pden_vol_summary.Row_expiry_date, pden_vol_summary.Row_quality, pden_vol_summary.Pden_subtype)
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

func PatchPdenVolSummary(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_vol_summary set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "posted_date" || key == "volume_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePdenVolSummary(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_vol_summary dto.Pden_vol_summary
	pden_vol_summary.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_vol_summary where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_vol_summary.Pden_subtype)
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


