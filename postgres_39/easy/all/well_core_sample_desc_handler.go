package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCoreSampleDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_core_sample_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_core_sample_desc

	for rows.Next() {
		var well_core_sample_desc dto.Well_core_sample_desc
		if err := rows.Scan(&well_core_sample_desc.Uwi, &well_core_sample_desc.Source, &well_core_sample_desc.Core_id, &well_core_sample_desc.Analysis_obs_no, &well_core_sample_desc.Sample_num, &well_core_sample_desc.Sample_analysis_obs_no, &well_core_sample_desc.Sample_desc_obs_no, &well_core_sample_desc.Active_ind, &well_core_sample_desc.Base_depth, &well_core_sample_desc.Base_depth_ouom, &well_core_sample_desc.Description, &well_core_sample_desc.Desc_source, &well_core_sample_desc.Dip_angle, &well_core_sample_desc.Effective_date, &well_core_sample_desc.Expiry_date, &well_core_sample_desc.Lithology_desc, &well_core_sample_desc.Porosity_length, &well_core_sample_desc.Porosity_length_ouom, &well_core_sample_desc.Porosity_quality, &well_core_sample_desc.Porosity_type, &well_core_sample_desc.Ppdm_guid, &well_core_sample_desc.Recovered_amount, &well_core_sample_desc.Recovered_amount_ouom, &well_core_sample_desc.Remark, &well_core_sample_desc.Sample_type, &well_core_sample_desc.Show_length, &well_core_sample_desc.Show_length_ouom, &well_core_sample_desc.Show_quality, &well_core_sample_desc.Show_type, &well_core_sample_desc.Strat_name_set_id, &well_core_sample_desc.Strat_unit_id, &well_core_sample_desc.Swc_recovery_type, &well_core_sample_desc.Top_depth, &well_core_sample_desc.Top_depth_ouom, &well_core_sample_desc.Row_changed_by, &well_core_sample_desc.Row_changed_date, &well_core_sample_desc.Row_created_by, &well_core_sample_desc.Row_created_date, &well_core_sample_desc.Row_effective_date, &well_core_sample_desc.Row_expiry_date, &well_core_sample_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_core_sample_desc to result
		result = append(result, well_core_sample_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCoreSampleDesc(c *fiber.Ctx) error {
	var well_core_sample_desc dto.Well_core_sample_desc

	setDefaults(&well_core_sample_desc)

	if err := json.Unmarshal(c.Body(), &well_core_sample_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_core_sample_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41)")
	if err != nil {
		return err
	}
	well_core_sample_desc.Row_created_date = formatDate(well_core_sample_desc.Row_created_date)
	well_core_sample_desc.Row_changed_date = formatDate(well_core_sample_desc.Row_changed_date)
	well_core_sample_desc.Effective_date = formatDateString(well_core_sample_desc.Effective_date)
	well_core_sample_desc.Expiry_date = formatDateString(well_core_sample_desc.Expiry_date)
	well_core_sample_desc.Row_effective_date = formatDateString(well_core_sample_desc.Row_effective_date)
	well_core_sample_desc.Row_expiry_date = formatDateString(well_core_sample_desc.Row_expiry_date)






	rows, err := stmt.Exec(well_core_sample_desc.Uwi, well_core_sample_desc.Source, well_core_sample_desc.Core_id, well_core_sample_desc.Analysis_obs_no, well_core_sample_desc.Sample_num, well_core_sample_desc.Sample_analysis_obs_no, well_core_sample_desc.Sample_desc_obs_no, well_core_sample_desc.Active_ind, well_core_sample_desc.Base_depth, well_core_sample_desc.Base_depth_ouom, well_core_sample_desc.Description, well_core_sample_desc.Desc_source, well_core_sample_desc.Dip_angle, well_core_sample_desc.Effective_date, well_core_sample_desc.Expiry_date, well_core_sample_desc.Lithology_desc, well_core_sample_desc.Porosity_length, well_core_sample_desc.Porosity_length_ouom, well_core_sample_desc.Porosity_quality, well_core_sample_desc.Porosity_type, well_core_sample_desc.Ppdm_guid, well_core_sample_desc.Recovered_amount, well_core_sample_desc.Recovered_amount_ouom, well_core_sample_desc.Remark, well_core_sample_desc.Sample_type, well_core_sample_desc.Show_length, well_core_sample_desc.Show_length_ouom, well_core_sample_desc.Show_quality, well_core_sample_desc.Show_type, well_core_sample_desc.Strat_name_set_id, well_core_sample_desc.Strat_unit_id, well_core_sample_desc.Swc_recovery_type, well_core_sample_desc.Top_depth, well_core_sample_desc.Top_depth_ouom, well_core_sample_desc.Row_changed_by, well_core_sample_desc.Row_changed_date, well_core_sample_desc.Row_created_by, well_core_sample_desc.Row_created_date, well_core_sample_desc.Row_effective_date, well_core_sample_desc.Row_expiry_date, well_core_sample_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCoreSampleDesc(c *fiber.Ctx) error {
	var well_core_sample_desc dto.Well_core_sample_desc

	setDefaults(&well_core_sample_desc)

	if err := json.Unmarshal(c.Body(), &well_core_sample_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_core_sample_desc.Uwi = id

    if well_core_sample_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_core_sample_desc set source = :1, core_id = :2, analysis_obs_no = :3, sample_num = :4, sample_analysis_obs_no = :5, sample_desc_obs_no = :6, active_ind = :7, base_depth = :8, base_depth_ouom = :9, description = :10, desc_source = :11, dip_angle = :12, effective_date = :13, expiry_date = :14, lithology_desc = :15, porosity_length = :16, porosity_length_ouom = :17, porosity_quality = :18, porosity_type = :19, ppdm_guid = :20, recovered_amount = :21, recovered_amount_ouom = :22, remark = :23, sample_type = :24, show_length = :25, show_length_ouom = :26, show_quality = :27, show_type = :28, strat_name_set_id = :29, strat_unit_id = :30, swc_recovery_type = :31, top_depth = :32, top_depth_ouom = :33, row_changed_by = :34, row_changed_date = :35, row_created_by = :36, row_effective_date = :37, row_expiry_date = :38, row_quality = :39 where uwi = :41")
	if err != nil {
		return err
	}

	well_core_sample_desc.Row_changed_date = formatDate(well_core_sample_desc.Row_changed_date)
	well_core_sample_desc.Effective_date = formatDateString(well_core_sample_desc.Effective_date)
	well_core_sample_desc.Expiry_date = formatDateString(well_core_sample_desc.Expiry_date)
	well_core_sample_desc.Row_effective_date = formatDateString(well_core_sample_desc.Row_effective_date)
	well_core_sample_desc.Row_expiry_date = formatDateString(well_core_sample_desc.Row_expiry_date)






	rows, err := stmt.Exec(well_core_sample_desc.Source, well_core_sample_desc.Core_id, well_core_sample_desc.Analysis_obs_no, well_core_sample_desc.Sample_num, well_core_sample_desc.Sample_analysis_obs_no, well_core_sample_desc.Sample_desc_obs_no, well_core_sample_desc.Active_ind, well_core_sample_desc.Base_depth, well_core_sample_desc.Base_depth_ouom, well_core_sample_desc.Description, well_core_sample_desc.Desc_source, well_core_sample_desc.Dip_angle, well_core_sample_desc.Effective_date, well_core_sample_desc.Expiry_date, well_core_sample_desc.Lithology_desc, well_core_sample_desc.Porosity_length, well_core_sample_desc.Porosity_length_ouom, well_core_sample_desc.Porosity_quality, well_core_sample_desc.Porosity_type, well_core_sample_desc.Ppdm_guid, well_core_sample_desc.Recovered_amount, well_core_sample_desc.Recovered_amount_ouom, well_core_sample_desc.Remark, well_core_sample_desc.Sample_type, well_core_sample_desc.Show_length, well_core_sample_desc.Show_length_ouom, well_core_sample_desc.Show_quality, well_core_sample_desc.Show_type, well_core_sample_desc.Strat_name_set_id, well_core_sample_desc.Strat_unit_id, well_core_sample_desc.Swc_recovery_type, well_core_sample_desc.Top_depth, well_core_sample_desc.Top_depth_ouom, well_core_sample_desc.Row_changed_by, well_core_sample_desc.Row_changed_date, well_core_sample_desc.Row_created_by, well_core_sample_desc.Row_effective_date, well_core_sample_desc.Row_expiry_date, well_core_sample_desc.Row_quality, well_core_sample_desc.Uwi)
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

func PatchWellCoreSampleDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_core_sample_desc set "
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

func DeleteWellCoreSampleDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_core_sample_desc dto.Well_core_sample_desc
	well_core_sample_desc.Uwi = id

	stmt, err := db.Prepare("delete from well_core_sample_desc where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_core_sample_desc.Uwi)
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


