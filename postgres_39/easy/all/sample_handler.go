package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSample(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sample")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sample

	for rows.Next() {
		var sample dto.Sample
		if err := rows.Scan(&sample.Sample_id, &sample.Active_ind, &sample.Base_md, &sample.Base_md_ouom, &sample.Base_strat_unit_id, &sample.Collected_by_ba_id, &sample.Collected_date, &sample.Collected_equip_id, &sample.Collected_for_ba_id, &sample.Collected_time, &sample.Collection_phase_type, &sample.Collection_press, &sample.Collection_press_ouom, &sample.Collection_temp, &sample.Collection_temp_ouom, &sample.Collection_test_type, &sample.Collection_type, &sample.Composite_sample_count, &sample.Composite_sample_ind, &sample.Condensed_section_ind, &sample.Current_length, &sample.Current_length_ouom, &sample.Current_length_percent, &sample.Current_volume, &sample.Current_volume_ouom, &sample.Current_volume_percent, &sample.Current_weight, &sample.Current_weight_ouom, &sample.Current_weight_percent, &sample.Cylinder_press, &sample.Cylinder_press_ouom, &sample.Cylinder_temp, &sample.Cylinder_temp_ouom, &sample.Destroyed_date, &sample.Effective_date, &sample.Expiry_date, &sample.Initial_length, &sample.Initial_length_ouom, &sample.Initial_length_percent, &sample.Initial_volume, &sample.Initial_volume_ouom, &sample.Initial_volume_percent, &sample.Initial_weight, &sample.Initial_weight_ouom, &sample.Initial_weight_percent, &sample.Interval_percent, &sample.Oil_stained_ind, &sample.Physical_item_id, &sample.Ppdm_guid, &sample.Remark, &sample.Reported_mud_desc, &sample.Reported_res_press, &sample.Reported_res_press_ouom, &sample.Sample_count, &sample.Sample_diameter, &sample.Sample_diameter_ouom, &sample.Sample_interval, &sample.Sample_interval_ouom, &sample.Sample_quality, &sample.Sample_quality_desc, &sample.Sample_ref_num, &sample.Sample_shape_type, &sample.Sample_type, &sample.Sample_width, &sample.Sample_width_ouom, &sample.Separator_press, &sample.Separator_press_ouom, &sample.Separator_temp, &sample.Separator_temp_ouom, &sample.Source, &sample.Stored_equip_id, &sample.Stored_phase_type, &sample.Strat_name_set_id, &sample.Sub_sample_ind, &sample.Top_md, &sample.Top_md_ouom, &sample.Top_strat_unit_id, &sample.Total_length, &sample.Total_length_ouom, &sample.Total_volume, &sample.Total_volume_ouom, &sample.Total_weight, &sample.Total_weight_ouom, &sample.Well_datum_type, &sample.Row_changed_by, &sample.Row_changed_date, &sample.Row_created_by, &sample.Row_created_date, &sample.Row_effective_date, &sample.Row_expiry_date, &sample.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sample to result
		result = append(result, sample)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSample(c *fiber.Ctx) error {
	var sample dto.Sample

	setDefaults(&sample)

	if err := json.Unmarshal(c.Body(), &sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sample values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91)")
	if err != nil {
		return err
	}
	sample.Row_created_date = formatDate(sample.Row_created_date)
	sample.Row_changed_date = formatDate(sample.Row_changed_date)
	sample.Collected_date = formatDateString(sample.Collected_date)
	sample.Collected_time = formatDateString(sample.Collected_time)
	sample.Destroyed_date = formatDateString(sample.Destroyed_date)
	sample.Effective_date = formatDateString(sample.Effective_date)
	sample.Expiry_date = formatDateString(sample.Expiry_date)
	sample.Row_effective_date = formatDateString(sample.Row_effective_date)
	sample.Row_expiry_date = formatDateString(sample.Row_expiry_date)






	rows, err := stmt.Exec(sample.Sample_id, sample.Active_ind, sample.Base_md, sample.Base_md_ouom, sample.Base_strat_unit_id, sample.Collected_by_ba_id, sample.Collected_date, sample.Collected_equip_id, sample.Collected_for_ba_id, sample.Collected_time, sample.Collection_phase_type, sample.Collection_press, sample.Collection_press_ouom, sample.Collection_temp, sample.Collection_temp_ouom, sample.Collection_test_type, sample.Collection_type, sample.Composite_sample_count, sample.Composite_sample_ind, sample.Condensed_section_ind, sample.Current_length, sample.Current_length_ouom, sample.Current_length_percent, sample.Current_volume, sample.Current_volume_ouom, sample.Current_volume_percent, sample.Current_weight, sample.Current_weight_ouom, sample.Current_weight_percent, sample.Cylinder_press, sample.Cylinder_press_ouom, sample.Cylinder_temp, sample.Cylinder_temp_ouom, sample.Destroyed_date, sample.Effective_date, sample.Expiry_date, sample.Initial_length, sample.Initial_length_ouom, sample.Initial_length_percent, sample.Initial_volume, sample.Initial_volume_ouom, sample.Initial_volume_percent, sample.Initial_weight, sample.Initial_weight_ouom, sample.Initial_weight_percent, sample.Interval_percent, sample.Oil_stained_ind, sample.Physical_item_id, sample.Ppdm_guid, sample.Remark, sample.Reported_mud_desc, sample.Reported_res_press, sample.Reported_res_press_ouom, sample.Sample_count, sample.Sample_diameter, sample.Sample_diameter_ouom, sample.Sample_interval, sample.Sample_interval_ouom, sample.Sample_quality, sample.Sample_quality_desc, sample.Sample_ref_num, sample.Sample_shape_type, sample.Sample_type, sample.Sample_width, sample.Sample_width_ouom, sample.Separator_press, sample.Separator_press_ouom, sample.Separator_temp, sample.Separator_temp_ouom, sample.Source, sample.Stored_equip_id, sample.Stored_phase_type, sample.Strat_name_set_id, sample.Sub_sample_ind, sample.Top_md, sample.Top_md_ouom, sample.Top_strat_unit_id, sample.Total_length, sample.Total_length_ouom, sample.Total_volume, sample.Total_volume_ouom, sample.Total_weight, sample.Total_weight_ouom, sample.Well_datum_type, sample.Row_changed_by, sample.Row_changed_date, sample.Row_created_by, sample.Row_created_date, sample.Row_effective_date, sample.Row_expiry_date, sample.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSample(c *fiber.Ctx) error {
	var sample dto.Sample

	setDefaults(&sample)

	if err := json.Unmarshal(c.Body(), &sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sample.Sample_id = id

    if sample.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sample set active_ind = :1, base_md = :2, base_md_ouom = :3, base_strat_unit_id = :4, collected_by_ba_id = :5, collected_date = :6, collected_equip_id = :7, collected_for_ba_id = :8, collected_time = :9, collection_phase_type = :10, collection_press = :11, collection_press_ouom = :12, collection_temp = :13, collection_temp_ouom = :14, collection_test_type = :15, collection_type = :16, composite_sample_count = :17, composite_sample_ind = :18, condensed_section_ind = :19, current_length = :20, current_length_ouom = :21, current_length_percent = :22, current_volume = :23, current_volume_ouom = :24, current_volume_percent = :25, current_weight = :26, current_weight_ouom = :27, current_weight_percent = :28, cylinder_press = :29, cylinder_press_ouom = :30, cylinder_temp = :31, cylinder_temp_ouom = :32, destroyed_date = :33, effective_date = :34, expiry_date = :35, initial_length = :36, initial_length_ouom = :37, initial_length_percent = :38, initial_volume = :39, initial_volume_ouom = :40, initial_volume_percent = :41, initial_weight = :42, initial_weight_ouom = :43, initial_weight_percent = :44, interval_percent = :45, oil_stained_ind = :46, physical_item_id = :47, ppdm_guid = :48, remark = :49, reported_mud_desc = :50, reported_res_press = :51, reported_res_press_ouom = :52, sample_count = :53, sample_diameter = :54, sample_diameter_ouom = :55, sample_interval = :56, sample_interval_ouom = :57, sample_quality = :58, sample_quality_desc = :59, sample_ref_num = :60, sample_shape_type = :61, sample_type = :62, sample_width = :63, sample_width_ouom = :64, separator_press = :65, separator_press_ouom = :66, separator_temp = :67, separator_temp_ouom = :68, source = :69, stored_equip_id = :70, stored_phase_type = :71, strat_name_set_id = :72, sub_sample_ind = :73, top_md = :74, top_md_ouom = :75, top_strat_unit_id = :76, total_length = :77, total_length_ouom = :78, total_volume = :79, total_volume_ouom = :80, total_weight = :81, total_weight_ouom = :82, well_datum_type = :83, row_changed_by = :84, row_changed_date = :85, row_created_by = :86, row_effective_date = :87, row_expiry_date = :88, row_quality = :89 where sample_id = :91")
	if err != nil {
		return err
	}

	sample.Row_changed_date = formatDate(sample.Row_changed_date)
	sample.Collected_date = formatDateString(sample.Collected_date)
	sample.Collected_time = formatDateString(sample.Collected_time)
	sample.Destroyed_date = formatDateString(sample.Destroyed_date)
	sample.Effective_date = formatDateString(sample.Effective_date)
	sample.Expiry_date = formatDateString(sample.Expiry_date)
	sample.Row_effective_date = formatDateString(sample.Row_effective_date)
	sample.Row_expiry_date = formatDateString(sample.Row_expiry_date)






	rows, err := stmt.Exec(sample.Active_ind, sample.Base_md, sample.Base_md_ouom, sample.Base_strat_unit_id, sample.Collected_by_ba_id, sample.Collected_date, sample.Collected_equip_id, sample.Collected_for_ba_id, sample.Collected_time, sample.Collection_phase_type, sample.Collection_press, sample.Collection_press_ouom, sample.Collection_temp, sample.Collection_temp_ouom, sample.Collection_test_type, sample.Collection_type, sample.Composite_sample_count, sample.Composite_sample_ind, sample.Condensed_section_ind, sample.Current_length, sample.Current_length_ouom, sample.Current_length_percent, sample.Current_volume, sample.Current_volume_ouom, sample.Current_volume_percent, sample.Current_weight, sample.Current_weight_ouom, sample.Current_weight_percent, sample.Cylinder_press, sample.Cylinder_press_ouom, sample.Cylinder_temp, sample.Cylinder_temp_ouom, sample.Destroyed_date, sample.Effective_date, sample.Expiry_date, sample.Initial_length, sample.Initial_length_ouom, sample.Initial_length_percent, sample.Initial_volume, sample.Initial_volume_ouom, sample.Initial_volume_percent, sample.Initial_weight, sample.Initial_weight_ouom, sample.Initial_weight_percent, sample.Interval_percent, sample.Oil_stained_ind, sample.Physical_item_id, sample.Ppdm_guid, sample.Remark, sample.Reported_mud_desc, sample.Reported_res_press, sample.Reported_res_press_ouom, sample.Sample_count, sample.Sample_diameter, sample.Sample_diameter_ouom, sample.Sample_interval, sample.Sample_interval_ouom, sample.Sample_quality, sample.Sample_quality_desc, sample.Sample_ref_num, sample.Sample_shape_type, sample.Sample_type, sample.Sample_width, sample.Sample_width_ouom, sample.Separator_press, sample.Separator_press_ouom, sample.Separator_temp, sample.Separator_temp_ouom, sample.Source, sample.Stored_equip_id, sample.Stored_phase_type, sample.Strat_name_set_id, sample.Sub_sample_ind, sample.Top_md, sample.Top_md_ouom, sample.Top_strat_unit_id, sample.Total_length, sample.Total_length_ouom, sample.Total_volume, sample.Total_volume_ouom, sample.Total_weight, sample.Total_weight_ouom, sample.Well_datum_type, sample.Row_changed_by, sample.Row_changed_date, sample.Row_created_by, sample.Row_effective_date, sample.Row_expiry_date, sample.Row_quality, sample.Sample_id)
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

func PatchSample(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sample set "
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
		} else if key == "collected_date" || key == "collected_time" || key == "destroyed_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where sample_id = :id"

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

func DeleteSample(c *fiber.Ctx) error {
	id := c.Params("id")
	var sample dto.Sample
	sample.Sample_id = id

	stmt, err := db.Prepare("delete from sample where sample_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sample.Sample_id)
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


