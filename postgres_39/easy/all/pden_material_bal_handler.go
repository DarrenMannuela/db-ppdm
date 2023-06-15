package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenMaterialBal(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_material_bal")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_material_bal

	for rows.Next() {
		var pden_material_bal dto.Pden_material_bal
		if err := rows.Scan(&pden_material_bal.Pden_subtype, &pden_material_bal.Pden_id, &pden_material_bal.Pden_source, &pden_material_bal.Product_type, &pden_material_bal.Case_id, &pden_material_bal.Active_ind, &pden_material_bal.Case_name, &pden_material_bal.Close_month, &pden_material_bal.Close_year, &pden_material_bal.Co2_percent, &pden_material_bal.Critical_press, &pden_material_bal.Critical_press_ouom, &pden_material_bal.Critical_temp, &pden_material_bal.Critical_temp_ouom, &pden_material_bal.Cum_volume, &pden_material_bal.Cum_volume_date, &pden_material_bal.Cum_volume_date_desc, &pden_material_bal.Cum_volume_ouom, &pden_material_bal.Cum_volume_uom, &pden_material_bal.Curve_fit_error, &pden_material_bal.Curve_fit_type, &pden_material_bal.Curve_intercept, &pden_material_bal.Curve_name, &pden_material_bal.Curve_slope, &pden_material_bal.Effective_date, &pden_material_bal.End_date, &pden_material_bal.Expiry_date, &pden_material_bal.Final_cum_volume, &pden_material_bal.Final_cum_volume_ouom, &pden_material_bal.Final_cum_volume_uom, &pden_material_bal.Final_press, &pden_material_bal.Final_press_ouom, &pden_material_bal.Gas_abandon_press, &pden_material_bal.Gas_abandon_press_ouom, &pden_material_bal.Gas_abandon_recover, &pden_material_bal.H2s_percent, &pden_material_bal.Initial_cum_volume, &pden_material_bal.Initial_cum_volume_ouom, &pden_material_bal.Initial_cum_volume_uom, &pden_material_bal.Initial_press, &pden_material_bal.Initial_press_ouom, &pden_material_bal.Initial_temp, &pden_material_bal.Initial_temp_ouom, &pden_material_bal.Measured_press_ind, &pden_material_bal.N2_percent, &pden_material_bal.Orig_gas_in_place, &pden_material_bal.Pool_datum_depth, &pden_material_bal.Pool_datum_depth_ouom, &pden_material_bal.Ppdm_guid, &pden_material_bal.Project_id, &pden_material_bal.Recov_gas_in_place, &pden_material_bal.Remark, &pden_material_bal.Source, &pden_material_bal.Specific_gravity, &pden_material_bal.Start_date, &pden_material_bal.Surface_loss_percent, &pden_material_bal.Volume, &pden_material_bal.Volume_ouom, &pden_material_bal.Volume_uom, &pden_material_bal.Zero_press_ind, &pden_material_bal.Z_factor_method, &pden_material_bal.Row_changed_by, &pden_material_bal.Row_changed_date, &pden_material_bal.Row_created_by, &pden_material_bal.Row_created_date, &pden_material_bal.Row_effective_date, &pden_material_bal.Row_expiry_date, &pden_material_bal.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_material_bal to result
		result = append(result, pden_material_bal)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenMaterialBal(c *fiber.Ctx) error {
	var pden_material_bal dto.Pden_material_bal

	setDefaults(&pden_material_bal)

	if err := json.Unmarshal(c.Body(), &pden_material_bal); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_material_bal values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68)")
	if err != nil {
		return err
	}
	pden_material_bal.Row_created_date = formatDate(pden_material_bal.Row_created_date)
	pden_material_bal.Row_changed_date = formatDate(pden_material_bal.Row_changed_date)
	pden_material_bal.Cum_volume_date = formatDateString(pden_material_bal.Cum_volume_date)
	pden_material_bal.Effective_date = formatDateString(pden_material_bal.Effective_date)
	pden_material_bal.End_date = formatDateString(pden_material_bal.End_date)
	pden_material_bal.Expiry_date = formatDateString(pden_material_bal.Expiry_date)
	pden_material_bal.Start_date = formatDateString(pden_material_bal.Start_date)
	pden_material_bal.Row_effective_date = formatDateString(pden_material_bal.Row_effective_date)
	pden_material_bal.Row_expiry_date = formatDateString(pden_material_bal.Row_expiry_date)






	rows, err := stmt.Exec(pden_material_bal.Pden_subtype, pden_material_bal.Pden_id, pden_material_bal.Pden_source, pden_material_bal.Product_type, pden_material_bal.Case_id, pden_material_bal.Active_ind, pden_material_bal.Case_name, pden_material_bal.Close_month, pden_material_bal.Close_year, pden_material_bal.Co2_percent, pden_material_bal.Critical_press, pden_material_bal.Critical_press_ouom, pden_material_bal.Critical_temp, pden_material_bal.Critical_temp_ouom, pden_material_bal.Cum_volume, pden_material_bal.Cum_volume_date, pden_material_bal.Cum_volume_date_desc, pden_material_bal.Cum_volume_ouom, pden_material_bal.Cum_volume_uom, pden_material_bal.Curve_fit_error, pden_material_bal.Curve_fit_type, pden_material_bal.Curve_intercept, pden_material_bal.Curve_name, pden_material_bal.Curve_slope, pden_material_bal.Effective_date, pden_material_bal.End_date, pden_material_bal.Expiry_date, pden_material_bal.Final_cum_volume, pden_material_bal.Final_cum_volume_ouom, pden_material_bal.Final_cum_volume_uom, pden_material_bal.Final_press, pden_material_bal.Final_press_ouom, pden_material_bal.Gas_abandon_press, pden_material_bal.Gas_abandon_press_ouom, pden_material_bal.Gas_abandon_recover, pden_material_bal.H2s_percent, pden_material_bal.Initial_cum_volume, pden_material_bal.Initial_cum_volume_ouom, pden_material_bal.Initial_cum_volume_uom, pden_material_bal.Initial_press, pden_material_bal.Initial_press_ouom, pden_material_bal.Initial_temp, pden_material_bal.Initial_temp_ouom, pden_material_bal.Measured_press_ind, pden_material_bal.N2_percent, pden_material_bal.Orig_gas_in_place, pden_material_bal.Pool_datum_depth, pden_material_bal.Pool_datum_depth_ouom, pden_material_bal.Ppdm_guid, pden_material_bal.Project_id, pden_material_bal.Recov_gas_in_place, pden_material_bal.Remark, pden_material_bal.Source, pden_material_bal.Specific_gravity, pden_material_bal.Start_date, pden_material_bal.Surface_loss_percent, pden_material_bal.Volume, pden_material_bal.Volume_ouom, pden_material_bal.Volume_uom, pden_material_bal.Zero_press_ind, pden_material_bal.Z_factor_method, pden_material_bal.Row_changed_by, pden_material_bal.Row_changed_date, pden_material_bal.Row_created_by, pden_material_bal.Row_created_date, pden_material_bal.Row_effective_date, pden_material_bal.Row_expiry_date, pden_material_bal.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenMaterialBal(c *fiber.Ctx) error {
	var pden_material_bal dto.Pden_material_bal

	setDefaults(&pden_material_bal)

	if err := json.Unmarshal(c.Body(), &pden_material_bal); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_material_bal.Pden_subtype = id

    if pden_material_bal.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_material_bal set pden_id = :1, pden_source = :2, product_type = :3, case_id = :4, active_ind = :5, case_name = :6, close_month = :7, close_year = :8, co2_percent = :9, critical_press = :10, critical_press_ouom = :11, critical_temp = :12, critical_temp_ouom = :13, cum_volume = :14, cum_volume_date = :15, cum_volume_date_desc = :16, cum_volume_ouom = :17, cum_volume_uom = :18, curve_fit_error = :19, curve_fit_type = :20, curve_intercept = :21, curve_name = :22, curve_slope = :23, effective_date = :24, end_date = :25, expiry_date = :26, final_cum_volume = :27, final_cum_volume_ouom = :28, final_cum_volume_uom = :29, final_press = :30, final_press_ouom = :31, gas_abandon_press = :32, gas_abandon_press_ouom = :33, gas_abandon_recover = :34, h2s_percent = :35, initial_cum_volume = :36, initial_cum_volume_ouom = :37, initial_cum_volume_uom = :38, initial_press = :39, initial_press_ouom = :40, initial_temp = :41, initial_temp_ouom = :42, measured_press_ind = :43, n2_percent = :44, orig_gas_in_place = :45, pool_datum_depth = :46, pool_datum_depth_ouom = :47, ppdm_guid = :48, project_id = :49, recov_gas_in_place = :50, remark = :51, source = :52, specific_gravity = :53, start_date = :54, surface_loss_percent = :55, volume = :56, volume_ouom = :57, volume_uom = :58, zero_press_ind = :59, z_factor_method = :60, row_changed_by = :61, row_changed_date = :62, row_created_by = :63, row_effective_date = :64, row_expiry_date = :65, row_quality = :66 where pden_subtype = :68")
	if err != nil {
		return err
	}

	pden_material_bal.Row_changed_date = formatDate(pden_material_bal.Row_changed_date)
	pden_material_bal.Cum_volume_date = formatDateString(pden_material_bal.Cum_volume_date)
	pden_material_bal.Effective_date = formatDateString(pden_material_bal.Effective_date)
	pden_material_bal.End_date = formatDateString(pden_material_bal.End_date)
	pden_material_bal.Expiry_date = formatDateString(pden_material_bal.Expiry_date)
	pden_material_bal.Start_date = formatDateString(pden_material_bal.Start_date)
	pden_material_bal.Row_effective_date = formatDateString(pden_material_bal.Row_effective_date)
	pden_material_bal.Row_expiry_date = formatDateString(pden_material_bal.Row_expiry_date)






	rows, err := stmt.Exec(pden_material_bal.Pden_id, pden_material_bal.Pden_source, pden_material_bal.Product_type, pden_material_bal.Case_id, pden_material_bal.Active_ind, pden_material_bal.Case_name, pden_material_bal.Close_month, pden_material_bal.Close_year, pden_material_bal.Co2_percent, pden_material_bal.Critical_press, pden_material_bal.Critical_press_ouom, pden_material_bal.Critical_temp, pden_material_bal.Critical_temp_ouom, pden_material_bal.Cum_volume, pden_material_bal.Cum_volume_date, pden_material_bal.Cum_volume_date_desc, pden_material_bal.Cum_volume_ouom, pden_material_bal.Cum_volume_uom, pden_material_bal.Curve_fit_error, pden_material_bal.Curve_fit_type, pden_material_bal.Curve_intercept, pden_material_bal.Curve_name, pden_material_bal.Curve_slope, pden_material_bal.Effective_date, pden_material_bal.End_date, pden_material_bal.Expiry_date, pden_material_bal.Final_cum_volume, pden_material_bal.Final_cum_volume_ouom, pden_material_bal.Final_cum_volume_uom, pden_material_bal.Final_press, pden_material_bal.Final_press_ouom, pden_material_bal.Gas_abandon_press, pden_material_bal.Gas_abandon_press_ouom, pden_material_bal.Gas_abandon_recover, pden_material_bal.H2s_percent, pden_material_bal.Initial_cum_volume, pden_material_bal.Initial_cum_volume_ouom, pden_material_bal.Initial_cum_volume_uom, pden_material_bal.Initial_press, pden_material_bal.Initial_press_ouom, pden_material_bal.Initial_temp, pden_material_bal.Initial_temp_ouom, pden_material_bal.Measured_press_ind, pden_material_bal.N2_percent, pden_material_bal.Orig_gas_in_place, pden_material_bal.Pool_datum_depth, pden_material_bal.Pool_datum_depth_ouom, pden_material_bal.Ppdm_guid, pden_material_bal.Project_id, pden_material_bal.Recov_gas_in_place, pden_material_bal.Remark, pden_material_bal.Source, pden_material_bal.Specific_gravity, pden_material_bal.Start_date, pden_material_bal.Surface_loss_percent, pden_material_bal.Volume, pden_material_bal.Volume_ouom, pden_material_bal.Volume_uom, pden_material_bal.Zero_press_ind, pden_material_bal.Z_factor_method, pden_material_bal.Row_changed_by, pden_material_bal.Row_changed_date, pden_material_bal.Row_created_by, pden_material_bal.Row_effective_date, pden_material_bal.Row_expiry_date, pden_material_bal.Row_quality, pden_material_bal.Pden_subtype)
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

func PatchPdenMaterialBal(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_material_bal set "
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
		} else if key == "cum_volume_date" || key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePdenMaterialBal(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_material_bal dto.Pden_material_bal
	pden_material_bal.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_material_bal where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_material_bal.Pden_subtype)
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


