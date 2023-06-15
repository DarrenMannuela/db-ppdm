package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlElemental(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_elemental")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_elemental

	for rows.Next() {
		var anl_elemental dto.Anl_elemental
		if err := rows.Scan(&anl_elemental.Analysis_id, &anl_elemental.Anl_source, &anl_elemental.Elemental_anl_obs_no, &anl_elemental.Active_ind, &anl_elemental.Ash_element, &anl_elemental.Ash_element_ouom, &anl_elemental.C_element, &anl_elemental.C_element_ouom, &anl_elemental.Effective_date, &anl_elemental.Expiry_date, &anl_elemental.Fe_element, &anl_elemental.Fe_element_ouom, &anl_elemental.H_element, &anl_elemental.H_element_ouom, &anl_elemental.Measurement_type, &anl_elemental.Ni_element, &anl_elemental.Ni_element_ouom, &anl_elemental.N_element, &anl_elemental.N_element_ouom, &anl_elemental.O_element, &anl_elemental.O_element_ouom, &anl_elemental.Ppdm_guid, &anl_elemental.Problem_ind, &anl_elemental.Remark, &anl_elemental.Reported_h_c_ratio, &anl_elemental.Reported_h_c_ratio_ouom, &anl_elemental.Reported_ni_c_ratio, &anl_elemental.Reported_ni_c_ratio_ouom, &anl_elemental.Reported_n_c_ratio, &anl_elemental.Reported_n_c_ratio_ouom, &anl_elemental.Reported_o_c_ratio, &anl_elemental.Reported_o_c_ratio_ouom, &anl_elemental.Source, &anl_elemental.Step_seq_no, &anl_elemental.S_element, &anl_elemental.S_element_ouom, &anl_elemental.V_element, &anl_elemental.V_element_ouom, &anl_elemental.Row_changed_by, &anl_elemental.Row_changed_date, &anl_elemental.Row_created_by, &anl_elemental.Row_created_date, &anl_elemental.Row_effective_date, &anl_elemental.Row_expiry_date, &anl_elemental.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_elemental to result
		result = append(result, anl_elemental)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlElemental(c *fiber.Ctx) error {
	var anl_elemental dto.Anl_elemental

	setDefaults(&anl_elemental)

	if err := json.Unmarshal(c.Body(), &anl_elemental); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_elemental values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45)")
	if err != nil {
		return err
	}
	anl_elemental.Row_created_date = formatDate(anl_elemental.Row_created_date)
	anl_elemental.Row_changed_date = formatDate(anl_elemental.Row_changed_date)
	anl_elemental.Effective_date = formatDateString(anl_elemental.Effective_date)
	anl_elemental.Expiry_date = formatDateString(anl_elemental.Expiry_date)
	anl_elemental.Row_effective_date = formatDateString(anl_elemental.Row_effective_date)
	anl_elemental.Row_expiry_date = formatDateString(anl_elemental.Row_expiry_date)






	rows, err := stmt.Exec(anl_elemental.Analysis_id, anl_elemental.Anl_source, anl_elemental.Elemental_anl_obs_no, anl_elemental.Active_ind, anl_elemental.Ash_element, anl_elemental.Ash_element_ouom, anl_elemental.C_element, anl_elemental.C_element_ouom, anl_elemental.Effective_date, anl_elemental.Expiry_date, anl_elemental.Fe_element, anl_elemental.Fe_element_ouom, anl_elemental.H_element, anl_elemental.H_element_ouom, anl_elemental.Measurement_type, anl_elemental.Ni_element, anl_elemental.Ni_element_ouom, anl_elemental.N_element, anl_elemental.N_element_ouom, anl_elemental.O_element, anl_elemental.O_element_ouom, anl_elemental.Ppdm_guid, anl_elemental.Problem_ind, anl_elemental.Remark, anl_elemental.Reported_h_c_ratio, anl_elemental.Reported_h_c_ratio_ouom, anl_elemental.Reported_ni_c_ratio, anl_elemental.Reported_ni_c_ratio_ouom, anl_elemental.Reported_n_c_ratio, anl_elemental.Reported_n_c_ratio_ouom, anl_elemental.Reported_o_c_ratio, anl_elemental.Reported_o_c_ratio_ouom, anl_elemental.Source, anl_elemental.Step_seq_no, anl_elemental.S_element, anl_elemental.S_element_ouom, anl_elemental.V_element, anl_elemental.V_element_ouom, anl_elemental.Row_changed_by, anl_elemental.Row_changed_date, anl_elemental.Row_created_by, anl_elemental.Row_created_date, anl_elemental.Row_effective_date, anl_elemental.Row_expiry_date, anl_elemental.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlElemental(c *fiber.Ctx) error {
	var anl_elemental dto.Anl_elemental

	setDefaults(&anl_elemental)

	if err := json.Unmarshal(c.Body(), &anl_elemental); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_elemental.Analysis_id = id

    if anl_elemental.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_elemental set anl_source = :1, elemental_anl_obs_no = :2, active_ind = :3, ash_element = :4, ash_element_ouom = :5, c_element = :6, c_element_ouom = :7, effective_date = :8, expiry_date = :9, fe_element = :10, fe_element_ouom = :11, h_element = :12, h_element_ouom = :13, measurement_type = :14, ni_element = :15, ni_element_ouom = :16, n_element = :17, n_element_ouom = :18, o_element = :19, o_element_ouom = :20, ppdm_guid = :21, problem_ind = :22, remark = :23, reported_h_c_ratio = :24, reported_h_c_ratio_ouom = :25, reported_ni_c_ratio = :26, reported_ni_c_ratio_ouom = :27, reported_n_c_ratio = :28, reported_n_c_ratio_ouom = :29, reported_o_c_ratio = :30, reported_o_c_ratio_ouom = :31, source = :32, step_seq_no = :33, s_element = :34, s_element_ouom = :35, v_element = :36, v_element_ouom = :37, row_changed_by = :38, row_changed_date = :39, row_created_by = :40, row_effective_date = :41, row_expiry_date = :42, row_quality = :43 where analysis_id = :45")
	if err != nil {
		return err
	}

	anl_elemental.Row_changed_date = formatDate(anl_elemental.Row_changed_date)
	anl_elemental.Effective_date = formatDateString(anl_elemental.Effective_date)
	anl_elemental.Expiry_date = formatDateString(anl_elemental.Expiry_date)
	anl_elemental.Row_effective_date = formatDateString(anl_elemental.Row_effective_date)
	anl_elemental.Row_expiry_date = formatDateString(anl_elemental.Row_expiry_date)






	rows, err := stmt.Exec(anl_elemental.Anl_source, anl_elemental.Elemental_anl_obs_no, anl_elemental.Active_ind, anl_elemental.Ash_element, anl_elemental.Ash_element_ouom, anl_elemental.C_element, anl_elemental.C_element_ouom, anl_elemental.Effective_date, anl_elemental.Expiry_date, anl_elemental.Fe_element, anl_elemental.Fe_element_ouom, anl_elemental.H_element, anl_elemental.H_element_ouom, anl_elemental.Measurement_type, anl_elemental.Ni_element, anl_elemental.Ni_element_ouom, anl_elemental.N_element, anl_elemental.N_element_ouom, anl_elemental.O_element, anl_elemental.O_element_ouom, anl_elemental.Ppdm_guid, anl_elemental.Problem_ind, anl_elemental.Remark, anl_elemental.Reported_h_c_ratio, anl_elemental.Reported_h_c_ratio_ouom, anl_elemental.Reported_ni_c_ratio, anl_elemental.Reported_ni_c_ratio_ouom, anl_elemental.Reported_n_c_ratio, anl_elemental.Reported_n_c_ratio_ouom, anl_elemental.Reported_o_c_ratio, anl_elemental.Reported_o_c_ratio_ouom, anl_elemental.Source, anl_elemental.Step_seq_no, anl_elemental.S_element, anl_elemental.S_element_ouom, anl_elemental.V_element, anl_elemental.V_element_ouom, anl_elemental.Row_changed_by, anl_elemental.Row_changed_date, anl_elemental.Row_created_by, anl_elemental.Row_effective_date, anl_elemental.Row_expiry_date, anl_elemental.Row_quality, anl_elemental.Analysis_id)
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

func PatchAnlElemental(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_elemental set "
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
	query += " where analysis_id = :id"

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

func DeleteAnlElemental(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_elemental dto.Anl_elemental
	anl_elemental.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_elemental where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_elemental.Analysis_id)
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


