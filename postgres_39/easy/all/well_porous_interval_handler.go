package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellPorousInterval(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_porous_interval")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_porous_interval

	for rows.Next() {
		var well_porous_interval dto.Well_porous_interval
		if err := rows.Scan(&well_porous_interval.Uwi, &well_porous_interval.Source, &well_porous_interval.Porous_interval_id, &well_porous_interval.Active_ind, &well_porous_interval.Average_porosity, &well_porous_interval.Base_depth, &well_porous_interval.Base_depth_ouom, &well_porous_interval.Base_tvd, &well_porous_interval.Base_tvd_ouom, &well_porous_interval.Core_sample_length, &well_porous_interval.Core_sample_length_ouom, &well_porous_interval.Effective_date, &well_porous_interval.Expiry_date, &well_porous_interval.Gas_show_ind, &well_porous_interval.Gross_thickness, &well_porous_interval.Gross_thickness_ouom, &well_porous_interval.Net_thickness, &well_porous_interval.Net_thickness_ouom, &well_porous_interval.Oil_show_ind, &well_porous_interval.Porosity_cutoff, &well_porous_interval.Porosity_quality, &well_porous_interval.Porosity_type, &well_porous_interval.Porous_form_age, &well_porous_interval.Porous_strat_unit_id, &well_porous_interval.Ppdm_guid, &well_porous_interval.Remark, &well_porous_interval.Strat_age, &well_porous_interval.Strat_age_name_set_id, &well_porous_interval.Strat_name_set_id, &well_porous_interval.Top_depth, &well_porous_interval.Top_depth_ouom, &well_porous_interval.Top_tvd, &well_porous_interval.Top_tvd_ouom, &well_porous_interval.Row_changed_by, &well_porous_interval.Row_changed_date, &well_porous_interval.Row_created_by, &well_porous_interval.Row_created_date, &well_porous_interval.Row_effective_date, &well_porous_interval.Row_expiry_date, &well_porous_interval.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_porous_interval to result
		result = append(result, well_porous_interval)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellPorousInterval(c *fiber.Ctx) error {
	var well_porous_interval dto.Well_porous_interval

	setDefaults(&well_porous_interval)

	if err := json.Unmarshal(c.Body(), &well_porous_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_porous_interval values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	well_porous_interval.Row_created_date = formatDate(well_porous_interval.Row_created_date)
	well_porous_interval.Row_changed_date = formatDate(well_porous_interval.Row_changed_date)
	well_porous_interval.Effective_date = formatDateString(well_porous_interval.Effective_date)
	well_porous_interval.Expiry_date = formatDateString(well_porous_interval.Expiry_date)
	well_porous_interval.Row_effective_date = formatDateString(well_porous_interval.Row_effective_date)
	well_porous_interval.Row_expiry_date = formatDateString(well_porous_interval.Row_expiry_date)






	rows, err := stmt.Exec(well_porous_interval.Uwi, well_porous_interval.Source, well_porous_interval.Porous_interval_id, well_porous_interval.Active_ind, well_porous_interval.Average_porosity, well_porous_interval.Base_depth, well_porous_interval.Base_depth_ouom, well_porous_interval.Base_tvd, well_porous_interval.Base_tvd_ouom, well_porous_interval.Core_sample_length, well_porous_interval.Core_sample_length_ouom, well_porous_interval.Effective_date, well_porous_interval.Expiry_date, well_porous_interval.Gas_show_ind, well_porous_interval.Gross_thickness, well_porous_interval.Gross_thickness_ouom, well_porous_interval.Net_thickness, well_porous_interval.Net_thickness_ouom, well_porous_interval.Oil_show_ind, well_porous_interval.Porosity_cutoff, well_porous_interval.Porosity_quality, well_porous_interval.Porosity_type, well_porous_interval.Porous_form_age, well_porous_interval.Porous_strat_unit_id, well_porous_interval.Ppdm_guid, well_porous_interval.Remark, well_porous_interval.Strat_age, well_porous_interval.Strat_age_name_set_id, well_porous_interval.Strat_name_set_id, well_porous_interval.Top_depth, well_porous_interval.Top_depth_ouom, well_porous_interval.Top_tvd, well_porous_interval.Top_tvd_ouom, well_porous_interval.Row_changed_by, well_porous_interval.Row_changed_date, well_porous_interval.Row_created_by, well_porous_interval.Row_created_date, well_porous_interval.Row_effective_date, well_porous_interval.Row_expiry_date, well_porous_interval.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellPorousInterval(c *fiber.Ctx) error {
	var well_porous_interval dto.Well_porous_interval

	setDefaults(&well_porous_interval)

	if err := json.Unmarshal(c.Body(), &well_porous_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_porous_interval.Uwi = id

    if well_porous_interval.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_porous_interval set source = :1, porous_interval_id = :2, active_ind = :3, average_porosity = :4, base_depth = :5, base_depth_ouom = :6, base_tvd = :7, base_tvd_ouom = :8, core_sample_length = :9, core_sample_length_ouom = :10, effective_date = :11, expiry_date = :12, gas_show_ind = :13, gross_thickness = :14, gross_thickness_ouom = :15, net_thickness = :16, net_thickness_ouom = :17, oil_show_ind = :18, porosity_cutoff = :19, porosity_quality = :20, porosity_type = :21, porous_form_age = :22, porous_strat_unit_id = :23, ppdm_guid = :24, remark = :25, strat_age = :26, strat_age_name_set_id = :27, strat_name_set_id = :28, top_depth = :29, top_depth_ouom = :30, top_tvd = :31, top_tvd_ouom = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where uwi = :40")
	if err != nil {
		return err
	}

	well_porous_interval.Row_changed_date = formatDate(well_porous_interval.Row_changed_date)
	well_porous_interval.Effective_date = formatDateString(well_porous_interval.Effective_date)
	well_porous_interval.Expiry_date = formatDateString(well_porous_interval.Expiry_date)
	well_porous_interval.Row_effective_date = formatDateString(well_porous_interval.Row_effective_date)
	well_porous_interval.Row_expiry_date = formatDateString(well_porous_interval.Row_expiry_date)






	rows, err := stmt.Exec(well_porous_interval.Source, well_porous_interval.Porous_interval_id, well_porous_interval.Active_ind, well_porous_interval.Average_porosity, well_porous_interval.Base_depth, well_porous_interval.Base_depth_ouom, well_porous_interval.Base_tvd, well_porous_interval.Base_tvd_ouom, well_porous_interval.Core_sample_length, well_porous_interval.Core_sample_length_ouom, well_porous_interval.Effective_date, well_porous_interval.Expiry_date, well_porous_interval.Gas_show_ind, well_porous_interval.Gross_thickness, well_porous_interval.Gross_thickness_ouom, well_porous_interval.Net_thickness, well_porous_interval.Net_thickness_ouom, well_porous_interval.Oil_show_ind, well_porous_interval.Porosity_cutoff, well_porous_interval.Porosity_quality, well_porous_interval.Porosity_type, well_porous_interval.Porous_form_age, well_porous_interval.Porous_strat_unit_id, well_porous_interval.Ppdm_guid, well_porous_interval.Remark, well_porous_interval.Strat_age, well_porous_interval.Strat_age_name_set_id, well_porous_interval.Strat_name_set_id, well_porous_interval.Top_depth, well_porous_interval.Top_depth_ouom, well_porous_interval.Top_tvd, well_porous_interval.Top_tvd_ouom, well_porous_interval.Row_changed_by, well_porous_interval.Row_changed_date, well_porous_interval.Row_created_by, well_porous_interval.Row_effective_date, well_porous_interval.Row_expiry_date, well_porous_interval.Row_quality, well_porous_interval.Uwi)
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

func PatchWellPorousInterval(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_porous_interval set "
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
	query += " where uwi = :id"

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

func DeleteWellPorousInterval(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_porous_interval dto.Well_porous_interval
	well_porous_interval.Uwi = id

	stmt, err := db.Prepare("delete from well_porous_interval where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_porous_interval.Uwi)
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


