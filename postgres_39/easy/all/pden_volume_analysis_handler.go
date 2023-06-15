package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenVolumeAnalysis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_volume_analysis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_volume_analysis

	for rows.Next() {
		var pden_volume_analysis dto.Pden_volume_analysis
		if err := rows.Scan(&pden_volume_analysis.Pden_subtype, &pden_volume_analysis.Pden_id, &pden_volume_analysis.Pden_source, &pden_volume_analysis.Product_type, &pden_volume_analysis.Case_id, &pden_volume_analysis.Active_ind, &pden_volume_analysis.Area_size, &pden_volume_analysis.Area_size_ouom, &pden_volume_analysis.Case_name, &pden_volume_analysis.Date_format_desc, &pden_volume_analysis.Effective_date, &pden_volume_analysis.End_date, &pden_volume_analysis.Expiry_date, &pden_volume_analysis.Gas_abandon_compress, &pden_volume_analysis.Gas_abandon_press, &pden_volume_analysis.Gas_abandon_press_ouom, &pden_volume_analysis.Gas_init_compress, &pden_volume_analysis.Gas_init_pressure, &pden_volume_analysis.Gas_in_place, &pden_volume_analysis.Gas_in_place_ouom, &pden_volume_analysis.Gas_original_in_place, &pden_volume_analysis.Gas_original_in_place_ouom, &pden_volume_analysis.Gas_ratio_bgi, &pden_volume_analysis.Gas_recovery, &pden_volume_analysis.Gross_net_pay_ratio, &pden_volume_analysis.Gross_pay, &pden_volume_analysis.Gross_pay_ouom, &pden_volume_analysis.Init_res_temp, &pden_volume_analysis.Init_res_temp_ouom, &pden_volume_analysis.Net_pay, &pden_volume_analysis.Net_pay_ouom, &pden_volume_analysis.Oil_in_place, &pden_volume_analysis.Oil_in_place_ouom, &pden_volume_analysis.Oil_original_in_place, &pden_volume_analysis.Oil_original_in_place_ouom, &pden_volume_analysis.Oil_recovery_primary, &pden_volume_analysis.Oil_recovery_secondary, &pden_volume_analysis.Oil_recovery_total, &pden_volume_analysis.Oil_residual_sat, &pden_volume_analysis.Oil_shrinkage, &pden_volume_analysis.Original_gor, &pden_volume_analysis.Original_gor_ouom, &pden_volume_analysis.Orig_sol_gas_in_place, &pden_volume_analysis.Orig_sol_gas_in_place_ouom, &pden_volume_analysis.Permeability, &pden_volume_analysis.Permeability_ouom, &pden_volume_analysis.Porosity, &pden_volume_analysis.Ppdm_guid, &pden_volume_analysis.Project_id, &pden_volume_analysis.Recov_gor, &pden_volume_analysis.Recov_sol_gas_in_place, &pden_volume_analysis.Recov_sol_gas_in_place_ouom, &pden_volume_analysis.Remark, &pden_volume_analysis.Sol_gas_recovery, &pden_volume_analysis.Source, &pden_volume_analysis.Start_date, &pden_volume_analysis.Volume, &pden_volume_analysis.Volume_ouom, &pden_volume_analysis.Volume_uom, &pden_volume_analysis.Water_saturation, &pden_volume_analysis.Row_changed_by, &pden_volume_analysis.Row_changed_date, &pden_volume_analysis.Row_created_by, &pden_volume_analysis.Row_created_date, &pden_volume_analysis.Row_effective_date, &pden_volume_analysis.Row_expiry_date, &pden_volume_analysis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_volume_analysis to result
		result = append(result, pden_volume_analysis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenVolumeAnalysis(c *fiber.Ctx) error {
	var pden_volume_analysis dto.Pden_volume_analysis

	setDefaults(&pden_volume_analysis)

	if err := json.Unmarshal(c.Body(), &pden_volume_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_volume_analysis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67)")
	if err != nil {
		return err
	}
	pden_volume_analysis.Row_created_date = formatDate(pden_volume_analysis.Row_created_date)
	pden_volume_analysis.Row_changed_date = formatDate(pden_volume_analysis.Row_changed_date)
	pden_volume_analysis.Date_format_desc = formatDateString(pden_volume_analysis.Date_format_desc)
	pden_volume_analysis.Effective_date = formatDateString(pden_volume_analysis.Effective_date)
	pden_volume_analysis.End_date = formatDateString(pden_volume_analysis.End_date)
	pden_volume_analysis.Expiry_date = formatDateString(pden_volume_analysis.Expiry_date)
	pden_volume_analysis.Start_date = formatDateString(pden_volume_analysis.Start_date)
	pden_volume_analysis.Row_effective_date = formatDateString(pden_volume_analysis.Row_effective_date)
	pden_volume_analysis.Row_expiry_date = formatDateString(pden_volume_analysis.Row_expiry_date)






	rows, err := stmt.Exec(pden_volume_analysis.Pden_subtype, pden_volume_analysis.Pden_id, pden_volume_analysis.Pden_source, pden_volume_analysis.Product_type, pden_volume_analysis.Case_id, pden_volume_analysis.Active_ind, pden_volume_analysis.Area_size, pden_volume_analysis.Area_size_ouom, pden_volume_analysis.Case_name, pden_volume_analysis.Date_format_desc, pden_volume_analysis.Effective_date, pden_volume_analysis.End_date, pden_volume_analysis.Expiry_date, pden_volume_analysis.Gas_abandon_compress, pden_volume_analysis.Gas_abandon_press, pden_volume_analysis.Gas_abandon_press_ouom, pden_volume_analysis.Gas_init_compress, pden_volume_analysis.Gas_init_pressure, pden_volume_analysis.Gas_in_place, pden_volume_analysis.Gas_in_place_ouom, pden_volume_analysis.Gas_original_in_place, pden_volume_analysis.Gas_original_in_place_ouom, pden_volume_analysis.Gas_ratio_bgi, pden_volume_analysis.Gas_recovery, pden_volume_analysis.Gross_net_pay_ratio, pden_volume_analysis.Gross_pay, pden_volume_analysis.Gross_pay_ouom, pden_volume_analysis.Init_res_temp, pden_volume_analysis.Init_res_temp_ouom, pden_volume_analysis.Net_pay, pden_volume_analysis.Net_pay_ouom, pden_volume_analysis.Oil_in_place, pden_volume_analysis.Oil_in_place_ouom, pden_volume_analysis.Oil_original_in_place, pden_volume_analysis.Oil_original_in_place_ouom, pden_volume_analysis.Oil_recovery_primary, pden_volume_analysis.Oil_recovery_secondary, pden_volume_analysis.Oil_recovery_total, pden_volume_analysis.Oil_residual_sat, pden_volume_analysis.Oil_shrinkage, pden_volume_analysis.Original_gor, pden_volume_analysis.Original_gor_ouom, pden_volume_analysis.Orig_sol_gas_in_place, pden_volume_analysis.Orig_sol_gas_in_place_ouom, pden_volume_analysis.Permeability, pden_volume_analysis.Permeability_ouom, pden_volume_analysis.Porosity, pden_volume_analysis.Ppdm_guid, pden_volume_analysis.Project_id, pden_volume_analysis.Recov_gor, pden_volume_analysis.Recov_sol_gas_in_place, pden_volume_analysis.Recov_sol_gas_in_place_ouom, pden_volume_analysis.Remark, pden_volume_analysis.Sol_gas_recovery, pden_volume_analysis.Source, pden_volume_analysis.Start_date, pden_volume_analysis.Volume, pden_volume_analysis.Volume_ouom, pden_volume_analysis.Volume_uom, pden_volume_analysis.Water_saturation, pden_volume_analysis.Row_changed_by, pden_volume_analysis.Row_changed_date, pden_volume_analysis.Row_created_by, pden_volume_analysis.Row_created_date, pden_volume_analysis.Row_effective_date, pden_volume_analysis.Row_expiry_date, pden_volume_analysis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenVolumeAnalysis(c *fiber.Ctx) error {
	var pden_volume_analysis dto.Pden_volume_analysis

	setDefaults(&pden_volume_analysis)

	if err := json.Unmarshal(c.Body(), &pden_volume_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_volume_analysis.Pden_subtype = id

    if pden_volume_analysis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_volume_analysis set pden_id = :1, pden_source = :2, product_type = :3, case_id = :4, active_ind = :5, area_size = :6, area_size_ouom = :7, case_name = :8, date_format_desc = :9, effective_date = :10, end_date = :11, expiry_date = :12, gas_abandon_compress = :13, gas_abandon_press = :14, gas_abandon_press_ouom = :15, gas_init_compress = :16, gas_init_pressure = :17, gas_in_place = :18, gas_in_place_ouom = :19, gas_original_in_place = :20, gas_original_in_place_ouom = :21, gas_ratio_bgi = :22, gas_recovery = :23, gross_net_pay_ratio = :24, gross_pay = :25, gross_pay_ouom = :26, init_res_temp = :27, init_res_temp_ouom = :28, net_pay = :29, net_pay_ouom = :30, oil_in_place = :31, oil_in_place_ouom = :32, oil_original_in_place = :33, oil_original_in_place_ouom = :34, oil_recovery_primary = :35, oil_recovery_secondary = :36, oil_recovery_total = :37, oil_residual_sat = :38, oil_shrinkage = :39, original_gor = :40, original_gor_ouom = :41, orig_sol_gas_in_place = :42, orig_sol_gas_in_place_ouom = :43, permeability = :44, permeability_ouom = :45, porosity = :46, ppdm_guid = :47, project_id = :48, recov_gor = :49, recov_sol_gas_in_place = :50, recov_sol_gas_in_place_ouom = :51, remark = :52, sol_gas_recovery = :53, source = :54, start_date = :55, volume = :56, volume_ouom = :57, volume_uom = :58, water_saturation = :59, row_changed_by = :60, row_changed_date = :61, row_created_by = :62, row_effective_date = :63, row_expiry_date = :64, row_quality = :65 where pden_subtype = :67")
	if err != nil {
		return err
	}

	pden_volume_analysis.Row_changed_date = formatDate(pden_volume_analysis.Row_changed_date)
	pden_volume_analysis.Date_format_desc = formatDateString(pden_volume_analysis.Date_format_desc)
	pden_volume_analysis.Effective_date = formatDateString(pden_volume_analysis.Effective_date)
	pden_volume_analysis.End_date = formatDateString(pden_volume_analysis.End_date)
	pden_volume_analysis.Expiry_date = formatDateString(pden_volume_analysis.Expiry_date)
	pden_volume_analysis.Start_date = formatDateString(pden_volume_analysis.Start_date)
	pden_volume_analysis.Row_effective_date = formatDateString(pden_volume_analysis.Row_effective_date)
	pden_volume_analysis.Row_expiry_date = formatDateString(pden_volume_analysis.Row_expiry_date)






	rows, err := stmt.Exec(pden_volume_analysis.Pden_id, pden_volume_analysis.Pden_source, pden_volume_analysis.Product_type, pden_volume_analysis.Case_id, pden_volume_analysis.Active_ind, pden_volume_analysis.Area_size, pden_volume_analysis.Area_size_ouom, pden_volume_analysis.Case_name, pden_volume_analysis.Date_format_desc, pden_volume_analysis.Effective_date, pden_volume_analysis.End_date, pden_volume_analysis.Expiry_date, pden_volume_analysis.Gas_abandon_compress, pden_volume_analysis.Gas_abandon_press, pden_volume_analysis.Gas_abandon_press_ouom, pden_volume_analysis.Gas_init_compress, pden_volume_analysis.Gas_init_pressure, pden_volume_analysis.Gas_in_place, pden_volume_analysis.Gas_in_place_ouom, pden_volume_analysis.Gas_original_in_place, pden_volume_analysis.Gas_original_in_place_ouom, pden_volume_analysis.Gas_ratio_bgi, pden_volume_analysis.Gas_recovery, pden_volume_analysis.Gross_net_pay_ratio, pden_volume_analysis.Gross_pay, pden_volume_analysis.Gross_pay_ouom, pden_volume_analysis.Init_res_temp, pden_volume_analysis.Init_res_temp_ouom, pden_volume_analysis.Net_pay, pden_volume_analysis.Net_pay_ouom, pden_volume_analysis.Oil_in_place, pden_volume_analysis.Oil_in_place_ouom, pden_volume_analysis.Oil_original_in_place, pden_volume_analysis.Oil_original_in_place_ouom, pden_volume_analysis.Oil_recovery_primary, pden_volume_analysis.Oil_recovery_secondary, pden_volume_analysis.Oil_recovery_total, pden_volume_analysis.Oil_residual_sat, pden_volume_analysis.Oil_shrinkage, pden_volume_analysis.Original_gor, pden_volume_analysis.Original_gor_ouom, pden_volume_analysis.Orig_sol_gas_in_place, pden_volume_analysis.Orig_sol_gas_in_place_ouom, pden_volume_analysis.Permeability, pden_volume_analysis.Permeability_ouom, pden_volume_analysis.Porosity, pden_volume_analysis.Ppdm_guid, pden_volume_analysis.Project_id, pden_volume_analysis.Recov_gor, pden_volume_analysis.Recov_sol_gas_in_place, pden_volume_analysis.Recov_sol_gas_in_place_ouom, pden_volume_analysis.Remark, pden_volume_analysis.Sol_gas_recovery, pden_volume_analysis.Source, pden_volume_analysis.Start_date, pden_volume_analysis.Volume, pden_volume_analysis.Volume_ouom, pden_volume_analysis.Volume_uom, pden_volume_analysis.Water_saturation, pden_volume_analysis.Row_changed_by, pden_volume_analysis.Row_changed_date, pden_volume_analysis.Row_created_by, pden_volume_analysis.Row_effective_date, pden_volume_analysis.Row_expiry_date, pden_volume_analysis.Row_quality, pden_volume_analysis.Pden_subtype)
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

func PatchPdenVolumeAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_volume_analysis set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePdenVolumeAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_volume_analysis dto.Pden_volume_analysis
	pden_volume_analysis.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_volume_analysis where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_volume_analysis.Pden_subtype)
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


