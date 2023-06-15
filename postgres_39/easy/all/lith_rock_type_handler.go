package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithRockType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_rock_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_rock_type

	for rows.Next() {
		var lith_rock_type dto.Lith_rock_type
		if err := rows.Scan(&lith_rock_type.Lithology_log_id, &lith_rock_type.Lith_log_source, &lith_rock_type.Depth_obs_no, &lith_rock_type.Rock_type, &lith_rock_type.Rock_type_obs_no, &lith_rock_type.Active_ind, &lith_rock_type.Class, &lith_rock_type.Class_modifier, &lith_rock_type.Consolidation, &lith_rock_type.Cut_color, &lith_rock_type.Effective_date, &lith_rock_type.Expiry_date, &lith_rock_type.Fluorescence_color, &lith_rock_type.Fluorescence_distribution, &lith_rock_type.Fluorescence_rate, &lith_rock_type.Framework_percent, &lith_rock_type.Matrix_percent, &lith_rock_type.Oil_stain, &lith_rock_type.Permeability_quality, &lith_rock_type.Porosity_grade_percent, &lith_rock_type.Ppdm_guid, &lith_rock_type.Prim_porosity_type, &lith_rock_type.Remark, &lith_rock_type.Residue_color, &lith_rock_type.Residue_percent, &lith_rock_type.Rock_class_scheme, &lith_rock_type.Rock_matrix, &lith_rock_type.Rock_profile, &lith_rock_type.Rock_rel_abundance, &lith_rock_type.Rock_type_percent, &lith_rock_type.Rounding, &lith_rock_type.Sec_porosity_type, &lith_rock_type.Solid_hcarbon_percent, &lith_rock_type.Solid_hcarbon_type, &lith_rock_type.Sorting, &lith_rock_type.Row_changed_by, &lith_rock_type.Row_changed_date, &lith_rock_type.Row_created_by, &lith_rock_type.Row_created_date, &lith_rock_type.Row_effective_date, &lith_rock_type.Row_expiry_date, &lith_rock_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_rock_type to result
		result = append(result, lith_rock_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithRockType(c *fiber.Ctx) error {
	var lith_rock_type dto.Lith_rock_type

	setDefaults(&lith_rock_type)

	if err := json.Unmarshal(c.Body(), &lith_rock_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_rock_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42)")
	if err != nil {
		return err
	}
	lith_rock_type.Row_created_date = formatDate(lith_rock_type.Row_created_date)
	lith_rock_type.Row_changed_date = formatDate(lith_rock_type.Row_changed_date)
	lith_rock_type.Effective_date = formatDateString(lith_rock_type.Effective_date)
	lith_rock_type.Expiry_date = formatDateString(lith_rock_type.Expiry_date)
	lith_rock_type.Row_effective_date = formatDateString(lith_rock_type.Row_effective_date)
	lith_rock_type.Row_expiry_date = formatDateString(lith_rock_type.Row_expiry_date)






	rows, err := stmt.Exec(lith_rock_type.Lithology_log_id, lith_rock_type.Lith_log_source, lith_rock_type.Depth_obs_no, lith_rock_type.Rock_type, lith_rock_type.Rock_type_obs_no, lith_rock_type.Active_ind, lith_rock_type.Class, lith_rock_type.Class_modifier, lith_rock_type.Consolidation, lith_rock_type.Cut_color, lith_rock_type.Effective_date, lith_rock_type.Expiry_date, lith_rock_type.Fluorescence_color, lith_rock_type.Fluorescence_distribution, lith_rock_type.Fluorescence_rate, lith_rock_type.Framework_percent, lith_rock_type.Matrix_percent, lith_rock_type.Oil_stain, lith_rock_type.Permeability_quality, lith_rock_type.Porosity_grade_percent, lith_rock_type.Ppdm_guid, lith_rock_type.Prim_porosity_type, lith_rock_type.Remark, lith_rock_type.Residue_color, lith_rock_type.Residue_percent, lith_rock_type.Rock_class_scheme, lith_rock_type.Rock_matrix, lith_rock_type.Rock_profile, lith_rock_type.Rock_rel_abundance, lith_rock_type.Rock_type_percent, lith_rock_type.Rounding, lith_rock_type.Sec_porosity_type, lith_rock_type.Solid_hcarbon_percent, lith_rock_type.Solid_hcarbon_type, lith_rock_type.Sorting, lith_rock_type.Row_changed_by, lith_rock_type.Row_changed_date, lith_rock_type.Row_created_by, lith_rock_type.Row_created_date, lith_rock_type.Row_effective_date, lith_rock_type.Row_expiry_date, lith_rock_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithRockType(c *fiber.Ctx) error {
	var lith_rock_type dto.Lith_rock_type

	setDefaults(&lith_rock_type)

	if err := json.Unmarshal(c.Body(), &lith_rock_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_rock_type.Lithology_log_id = id

    if lith_rock_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_rock_type set lith_log_source = :1, depth_obs_no = :2, rock_type = :3, rock_type_obs_no = :4, active_ind = :5, class = :6, class_modifier = :7, consolidation = :8, cut_color = :9, effective_date = :10, expiry_date = :11, fluorescence_color = :12, fluorescence_distribution = :13, fluorescence_rate = :14, framework_percent = :15, matrix_percent = :16, oil_stain = :17, permeability_quality = :18, porosity_grade_percent = :19, ppdm_guid = :20, prim_porosity_type = :21, remark = :22, residue_color = :23, residue_percent = :24, rock_class_scheme = :25, rock_matrix = :26, rock_profile = :27, rock_rel_abundance = :28, rock_type_percent = :29, rounding = :30, sec_porosity_type = :31, solid_hcarbon_percent = :32, solid_hcarbon_type = :33, sorting = :34, row_changed_by = :35, row_changed_date = :36, row_created_by = :37, row_effective_date = :38, row_expiry_date = :39, row_quality = :40 where lithology_log_id = :42")
	if err != nil {
		return err
	}

	lith_rock_type.Row_changed_date = formatDate(lith_rock_type.Row_changed_date)
	lith_rock_type.Effective_date = formatDateString(lith_rock_type.Effective_date)
	lith_rock_type.Expiry_date = formatDateString(lith_rock_type.Expiry_date)
	lith_rock_type.Row_effective_date = formatDateString(lith_rock_type.Row_effective_date)
	lith_rock_type.Row_expiry_date = formatDateString(lith_rock_type.Row_expiry_date)






	rows, err := stmt.Exec(lith_rock_type.Lith_log_source, lith_rock_type.Depth_obs_no, lith_rock_type.Rock_type, lith_rock_type.Rock_type_obs_no, lith_rock_type.Active_ind, lith_rock_type.Class, lith_rock_type.Class_modifier, lith_rock_type.Consolidation, lith_rock_type.Cut_color, lith_rock_type.Effective_date, lith_rock_type.Expiry_date, lith_rock_type.Fluorescence_color, lith_rock_type.Fluorescence_distribution, lith_rock_type.Fluorescence_rate, lith_rock_type.Framework_percent, lith_rock_type.Matrix_percent, lith_rock_type.Oil_stain, lith_rock_type.Permeability_quality, lith_rock_type.Porosity_grade_percent, lith_rock_type.Ppdm_guid, lith_rock_type.Prim_porosity_type, lith_rock_type.Remark, lith_rock_type.Residue_color, lith_rock_type.Residue_percent, lith_rock_type.Rock_class_scheme, lith_rock_type.Rock_matrix, lith_rock_type.Rock_profile, lith_rock_type.Rock_rel_abundance, lith_rock_type.Rock_type_percent, lith_rock_type.Rounding, lith_rock_type.Sec_porosity_type, lith_rock_type.Solid_hcarbon_percent, lith_rock_type.Solid_hcarbon_type, lith_rock_type.Sorting, lith_rock_type.Row_changed_by, lith_rock_type.Row_changed_date, lith_rock_type.Row_created_by, lith_rock_type.Row_effective_date, lith_rock_type.Row_expiry_date, lith_rock_type.Row_quality, lith_rock_type.Lithology_log_id)
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

func PatchLithRockType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_rock_type set "
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
	query += " where lithology_log_id = :id"

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

func DeleteLithRockType(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_rock_type dto.Lith_rock_type
	lith_rock_type.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_rock_type where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_rock_type.Lithology_log_id)
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


