package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratColumnUnit(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_column_unit")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_column_unit

	for rows.Next() {
		var strat_column_unit dto.Strat_column_unit
		if err := rows.Scan(&strat_column_unit.Strat_column_id, &strat_column_unit.Strat_column_source, &strat_column_unit.Strat_name_set_id, &strat_column_unit.Strat_unit_id, &strat_column_unit.Interp_id, &strat_column_unit.Active_ind, &strat_column_unit.Business_associate_id, &strat_column_unit.Certified_ind, &strat_column_unit.Conformity_relationship, &strat_column_unit.Effective_date, &strat_column_unit.Expiry_date, &strat_column_unit.Interpretation_version, &strat_column_unit.Lithology, &strat_column_unit.Occurrence_type, &strat_column_unit.Ordinal_seq_no, &strat_column_unit.Overturned_ind, &strat_column_unit.Ppdm_guid, &strat_column_unit.Preferred_ind, &strat_column_unit.Remark, &strat_column_unit.Repeat_strat_occur_no, &strat_column_unit.Repeat_strat_type, &strat_column_unit.Source, &strat_column_unit.Source_document_id, &strat_column_unit.Sour_gas_ind, &strat_column_unit.Strat_interpret_method, &strat_column_unit.Strat_topo_relation_ind, &strat_column_unit.Sw_application_id, &strat_column_unit.Version_obs_no, &strat_column_unit.Row_changed_by, &strat_column_unit.Row_changed_date, &strat_column_unit.Row_created_by, &strat_column_unit.Row_created_date, &strat_column_unit.Row_effective_date, &strat_column_unit.Row_expiry_date, &strat_column_unit.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_column_unit to result
		result = append(result, strat_column_unit)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratColumnUnit(c *fiber.Ctx) error {
	var strat_column_unit dto.Strat_column_unit

	setDefaults(&strat_column_unit)

	if err := json.Unmarshal(c.Body(), &strat_column_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_column_unit values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35)")
	if err != nil {
		return err
	}
	strat_column_unit.Row_created_date = formatDate(strat_column_unit.Row_created_date)
	strat_column_unit.Row_changed_date = formatDate(strat_column_unit.Row_changed_date)
	strat_column_unit.Effective_date = formatDateString(strat_column_unit.Effective_date)
	strat_column_unit.Expiry_date = formatDateString(strat_column_unit.Expiry_date)
	strat_column_unit.Row_effective_date = formatDateString(strat_column_unit.Row_effective_date)
	strat_column_unit.Row_expiry_date = formatDateString(strat_column_unit.Row_expiry_date)






	rows, err := stmt.Exec(strat_column_unit.Strat_column_id, strat_column_unit.Strat_column_source, strat_column_unit.Strat_name_set_id, strat_column_unit.Strat_unit_id, strat_column_unit.Interp_id, strat_column_unit.Active_ind, strat_column_unit.Business_associate_id, strat_column_unit.Certified_ind, strat_column_unit.Conformity_relationship, strat_column_unit.Effective_date, strat_column_unit.Expiry_date, strat_column_unit.Interpretation_version, strat_column_unit.Lithology, strat_column_unit.Occurrence_type, strat_column_unit.Ordinal_seq_no, strat_column_unit.Overturned_ind, strat_column_unit.Ppdm_guid, strat_column_unit.Preferred_ind, strat_column_unit.Remark, strat_column_unit.Repeat_strat_occur_no, strat_column_unit.Repeat_strat_type, strat_column_unit.Source, strat_column_unit.Source_document_id, strat_column_unit.Sour_gas_ind, strat_column_unit.Strat_interpret_method, strat_column_unit.Strat_topo_relation_ind, strat_column_unit.Sw_application_id, strat_column_unit.Version_obs_no, strat_column_unit.Row_changed_by, strat_column_unit.Row_changed_date, strat_column_unit.Row_created_by, strat_column_unit.Row_created_date, strat_column_unit.Row_effective_date, strat_column_unit.Row_expiry_date, strat_column_unit.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratColumnUnit(c *fiber.Ctx) error {
	var strat_column_unit dto.Strat_column_unit

	setDefaults(&strat_column_unit)

	if err := json.Unmarshal(c.Body(), &strat_column_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_column_unit.Strat_column_id = id

    if strat_column_unit.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_column_unit set strat_column_source = :1, strat_name_set_id = :2, strat_unit_id = :3, interp_id = :4, active_ind = :5, business_associate_id = :6, certified_ind = :7, conformity_relationship = :8, effective_date = :9, expiry_date = :10, interpretation_version = :11, lithology = :12, occurrence_type = :13, ordinal_seq_no = :14, overturned_ind = :15, ppdm_guid = :16, preferred_ind = :17, remark = :18, repeat_strat_occur_no = :19, repeat_strat_type = :20, source = :21, source_document_id = :22, sour_gas_ind = :23, strat_interpret_method = :24, strat_topo_relation_ind = :25, sw_application_id = :26, version_obs_no = :27, row_changed_by = :28, row_changed_date = :29, row_created_by = :30, row_effective_date = :31, row_expiry_date = :32, row_quality = :33 where strat_column_id = :35")
	if err != nil {
		return err
	}

	strat_column_unit.Row_changed_date = formatDate(strat_column_unit.Row_changed_date)
	strat_column_unit.Effective_date = formatDateString(strat_column_unit.Effective_date)
	strat_column_unit.Expiry_date = formatDateString(strat_column_unit.Expiry_date)
	strat_column_unit.Row_effective_date = formatDateString(strat_column_unit.Row_effective_date)
	strat_column_unit.Row_expiry_date = formatDateString(strat_column_unit.Row_expiry_date)






	rows, err := stmt.Exec(strat_column_unit.Strat_column_source, strat_column_unit.Strat_name_set_id, strat_column_unit.Strat_unit_id, strat_column_unit.Interp_id, strat_column_unit.Active_ind, strat_column_unit.Business_associate_id, strat_column_unit.Certified_ind, strat_column_unit.Conformity_relationship, strat_column_unit.Effective_date, strat_column_unit.Expiry_date, strat_column_unit.Interpretation_version, strat_column_unit.Lithology, strat_column_unit.Occurrence_type, strat_column_unit.Ordinal_seq_no, strat_column_unit.Overturned_ind, strat_column_unit.Ppdm_guid, strat_column_unit.Preferred_ind, strat_column_unit.Remark, strat_column_unit.Repeat_strat_occur_no, strat_column_unit.Repeat_strat_type, strat_column_unit.Source, strat_column_unit.Source_document_id, strat_column_unit.Sour_gas_ind, strat_column_unit.Strat_interpret_method, strat_column_unit.Strat_topo_relation_ind, strat_column_unit.Sw_application_id, strat_column_unit.Version_obs_no, strat_column_unit.Row_changed_by, strat_column_unit.Row_changed_date, strat_column_unit.Row_created_by, strat_column_unit.Row_effective_date, strat_column_unit.Row_expiry_date, strat_column_unit.Row_quality, strat_column_unit.Strat_column_id)
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

func PatchStratColumnUnit(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_column_unit set "
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
	query += " where strat_column_id = :id"

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

func DeleteStratColumnUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_column_unit dto.Strat_column_unit
	strat_column_unit.Strat_column_id = id

	stmt, err := db.Prepare("delete from strat_column_unit where strat_column_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_column_unit.Strat_column_id)
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


