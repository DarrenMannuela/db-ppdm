package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoInterp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_interp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_interp

	for rows.Next() {
		var paleo_interp dto.Paleo_interp
		if err := rows.Scan(&paleo_interp.Paleo_summary_id, &paleo_interp.Detail_id, &paleo_interp.Active_ind, &paleo_interp.Analysis_id, &paleo_interp.Anl_source, &paleo_interp.Base_md, &paleo_interp.Base_md_ouom, &paleo_interp.Climate_id, &paleo_interp.Description, &paleo_interp.Ecozone_confidence_id, &paleo_interp.Ecozone_id, &paleo_interp.Ecozone_qualifier_id, &paleo_interp.Effective_date, &paleo_interp.Expiry_date, &paleo_interp.From_rel_strat_name_set_id, &paleo_interp.From_rel_strat_unit_id, &paleo_interp.From_strat_name_set_id, &paleo_interp.From_strat_unit_id, &paleo_interp.Interp_type, &paleo_interp.Lithology_type, &paleo_interp.Lith_structure_id, &paleo_interp.Maturation_obs_no, &paleo_interp.Paleo_confidence_id, &paleo_interp.Paleo_qualifier_id, &paleo_interp.Ppdm_guid, &paleo_interp.Preferred_ind, &paleo_interp.Remark, &paleo_interp.Source, &paleo_interp.Tai_color, &paleo_interp.Top_md, &paleo_interp.Top_md_ouom, &paleo_interp.To_rel_strat_name_set_id, &paleo_interp.To_rel_strat_unit_id, &paleo_interp.To_strat_name_set_id, &paleo_interp.To_strat_unit_id, &paleo_interp.Row_changed_by, &paleo_interp.Row_changed_date, &paleo_interp.Row_created_by, &paleo_interp.Row_created_date, &paleo_interp.Row_effective_date, &paleo_interp.Row_expiry_date, &paleo_interp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_interp to result
		result = append(result, paleo_interp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoInterp(c *fiber.Ctx) error {
	var paleo_interp dto.Paleo_interp

	setDefaults(&paleo_interp)

	if err := json.Unmarshal(c.Body(), &paleo_interp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_interp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42)")
	if err != nil {
		return err
	}
	paleo_interp.Row_created_date = formatDate(paleo_interp.Row_created_date)
	paleo_interp.Row_changed_date = formatDate(paleo_interp.Row_changed_date)
	paleo_interp.Effective_date = formatDateString(paleo_interp.Effective_date)
	paleo_interp.Expiry_date = formatDateString(paleo_interp.Expiry_date)
	paleo_interp.Row_effective_date = formatDateString(paleo_interp.Row_effective_date)
	paleo_interp.Row_expiry_date = formatDateString(paleo_interp.Row_expiry_date)






	rows, err := stmt.Exec(paleo_interp.Paleo_summary_id, paleo_interp.Detail_id, paleo_interp.Active_ind, paleo_interp.Analysis_id, paleo_interp.Anl_source, paleo_interp.Base_md, paleo_interp.Base_md_ouom, paleo_interp.Climate_id, paleo_interp.Description, paleo_interp.Ecozone_confidence_id, paleo_interp.Ecozone_id, paleo_interp.Ecozone_qualifier_id, paleo_interp.Effective_date, paleo_interp.Expiry_date, paleo_interp.From_rel_strat_name_set_id, paleo_interp.From_rel_strat_unit_id, paleo_interp.From_strat_name_set_id, paleo_interp.From_strat_unit_id, paleo_interp.Interp_type, paleo_interp.Lithology_type, paleo_interp.Lith_structure_id, paleo_interp.Maturation_obs_no, paleo_interp.Paleo_confidence_id, paleo_interp.Paleo_qualifier_id, paleo_interp.Ppdm_guid, paleo_interp.Preferred_ind, paleo_interp.Remark, paleo_interp.Source, paleo_interp.Tai_color, paleo_interp.Top_md, paleo_interp.Top_md_ouom, paleo_interp.To_rel_strat_name_set_id, paleo_interp.To_rel_strat_unit_id, paleo_interp.To_strat_name_set_id, paleo_interp.To_strat_unit_id, paleo_interp.Row_changed_by, paleo_interp.Row_changed_date, paleo_interp.Row_created_by, paleo_interp.Row_created_date, paleo_interp.Row_effective_date, paleo_interp.Row_expiry_date, paleo_interp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoInterp(c *fiber.Ctx) error {
	var paleo_interp dto.Paleo_interp

	setDefaults(&paleo_interp)

	if err := json.Unmarshal(c.Body(), &paleo_interp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_interp.Paleo_summary_id = id

    if paleo_interp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_interp set detail_id = :1, active_ind = :2, analysis_id = :3, anl_source = :4, base_md = :5, base_md_ouom = :6, climate_id = :7, description = :8, ecozone_confidence_id = :9, ecozone_id = :10, ecozone_qualifier_id = :11, effective_date = :12, expiry_date = :13, from_rel_strat_name_set_id = :14, from_rel_strat_unit_id = :15, from_strat_name_set_id = :16, from_strat_unit_id = :17, interp_type = :18, lithology_type = :19, lith_structure_id = :20, maturation_obs_no = :21, paleo_confidence_id = :22, paleo_qualifier_id = :23, ppdm_guid = :24, preferred_ind = :25, remark = :26, source = :27, tai_color = :28, top_md = :29, top_md_ouom = :30, to_rel_strat_name_set_id = :31, to_rel_strat_unit_id = :32, to_strat_name_set_id = :33, to_strat_unit_id = :34, row_changed_by = :35, row_changed_date = :36, row_created_by = :37, row_effective_date = :38, row_expiry_date = :39, row_quality = :40 where paleo_summary_id = :42")
	if err != nil {
		return err
	}

	paleo_interp.Row_changed_date = formatDate(paleo_interp.Row_changed_date)
	paleo_interp.Effective_date = formatDateString(paleo_interp.Effective_date)
	paleo_interp.Expiry_date = formatDateString(paleo_interp.Expiry_date)
	paleo_interp.Row_effective_date = formatDateString(paleo_interp.Row_effective_date)
	paleo_interp.Row_expiry_date = formatDateString(paleo_interp.Row_expiry_date)






	rows, err := stmt.Exec(paleo_interp.Detail_id, paleo_interp.Active_ind, paleo_interp.Analysis_id, paleo_interp.Anl_source, paleo_interp.Base_md, paleo_interp.Base_md_ouom, paleo_interp.Climate_id, paleo_interp.Description, paleo_interp.Ecozone_confidence_id, paleo_interp.Ecozone_id, paleo_interp.Ecozone_qualifier_id, paleo_interp.Effective_date, paleo_interp.Expiry_date, paleo_interp.From_rel_strat_name_set_id, paleo_interp.From_rel_strat_unit_id, paleo_interp.From_strat_name_set_id, paleo_interp.From_strat_unit_id, paleo_interp.Interp_type, paleo_interp.Lithology_type, paleo_interp.Lith_structure_id, paleo_interp.Maturation_obs_no, paleo_interp.Paleo_confidence_id, paleo_interp.Paleo_qualifier_id, paleo_interp.Ppdm_guid, paleo_interp.Preferred_ind, paleo_interp.Remark, paleo_interp.Source, paleo_interp.Tai_color, paleo_interp.Top_md, paleo_interp.Top_md_ouom, paleo_interp.To_rel_strat_name_set_id, paleo_interp.To_rel_strat_unit_id, paleo_interp.To_strat_name_set_id, paleo_interp.To_strat_unit_id, paleo_interp.Row_changed_by, paleo_interp.Row_changed_date, paleo_interp.Row_created_by, paleo_interp.Row_effective_date, paleo_interp.Row_expiry_date, paleo_interp.Row_quality, paleo_interp.Paleo_summary_id)
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

func PatchPaleoInterp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_interp set "
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
	query += " where paleo_summary_id = :id"

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

func DeletePaleoInterp(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_interp dto.Paleo_interp
	paleo_interp.Paleo_summary_id = id

	stmt, err := db.Prepare("delete from paleo_interp where paleo_summary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_interp.Paleo_summary_id)
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


