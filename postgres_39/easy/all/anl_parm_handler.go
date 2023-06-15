package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlParm(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_parm")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_parm

	for rows.Next() {
		var anl_parm dto.Anl_parm
		if err := rows.Scan(&anl_parm.Analysis_id, &anl_parm.Anl_source, &anl_parm.Parm_obs_no, &anl_parm.Active_ind, &anl_parm.Anl_value_remark, &anl_parm.Average_value, &anl_parm.Average_value_ouom, &anl_parm.Average_value_uom, &anl_parm.Cat_equip_id, &anl_parm.Effective_date, &anl_parm.Equip_id, &anl_parm.Equip_obs_no, &anl_parm.Expiry_date, &anl_parm.Max_date, &anl_parm.Max_value, &anl_parm.Max_value_ouom, &anl_parm.Max_value_uom, &anl_parm.Method_id, &anl_parm.Method_parm_obs_no, &anl_parm.Method_set_id, &anl_parm.Min_date, &anl_parm.Min_value, &anl_parm.Min_value_ouom, &anl_parm.Min_value_uom, &anl_parm.Parameter_ba_id, &anl_parm.Parameter_type, &anl_parm.Ppdm_guid, &anl_parm.Reference_value, &anl_parm.Reference_value_ouom, &anl_parm.Reference_value_type, &anl_parm.Reference_value_uom, &anl_parm.Remark, &anl_parm.Source, &anl_parm.Step_seq_no, &anl_parm.Substance_id, &anl_parm.Row_changed_by, &anl_parm.Row_changed_date, &anl_parm.Row_created_by, &anl_parm.Row_created_date, &anl_parm.Row_effective_date, &anl_parm.Row_expiry_date, &anl_parm.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_parm to result
		result = append(result, anl_parm)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlParm(c *fiber.Ctx) error {
	var anl_parm dto.Anl_parm

	setDefaults(&anl_parm)

	if err := json.Unmarshal(c.Body(), &anl_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_parm values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42)")
	if err != nil {
		return err
	}
	anl_parm.Row_created_date = formatDate(anl_parm.Row_created_date)
	anl_parm.Row_changed_date = formatDate(anl_parm.Row_changed_date)
	anl_parm.Effective_date = formatDateString(anl_parm.Effective_date)
	anl_parm.Expiry_date = formatDateString(anl_parm.Expiry_date)
	anl_parm.Max_date = formatDateString(anl_parm.Max_date)
	anl_parm.Min_date = formatDateString(anl_parm.Min_date)
	anl_parm.Row_effective_date = formatDateString(anl_parm.Row_effective_date)
	anl_parm.Row_expiry_date = formatDateString(anl_parm.Row_expiry_date)






	rows, err := stmt.Exec(anl_parm.Analysis_id, anl_parm.Anl_source, anl_parm.Parm_obs_no, anl_parm.Active_ind, anl_parm.Anl_value_remark, anl_parm.Average_value, anl_parm.Average_value_ouom, anl_parm.Average_value_uom, anl_parm.Cat_equip_id, anl_parm.Effective_date, anl_parm.Equip_id, anl_parm.Equip_obs_no, anl_parm.Expiry_date, anl_parm.Max_date, anl_parm.Max_value, anl_parm.Max_value_ouom, anl_parm.Max_value_uom, anl_parm.Method_id, anl_parm.Method_parm_obs_no, anl_parm.Method_set_id, anl_parm.Min_date, anl_parm.Min_value, anl_parm.Min_value_ouom, anl_parm.Min_value_uom, anl_parm.Parameter_ba_id, anl_parm.Parameter_type, anl_parm.Ppdm_guid, anl_parm.Reference_value, anl_parm.Reference_value_ouom, anl_parm.Reference_value_type, anl_parm.Reference_value_uom, anl_parm.Remark, anl_parm.Source, anl_parm.Step_seq_no, anl_parm.Substance_id, anl_parm.Row_changed_by, anl_parm.Row_changed_date, anl_parm.Row_created_by, anl_parm.Row_created_date, anl_parm.Row_effective_date, anl_parm.Row_expiry_date, anl_parm.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlParm(c *fiber.Ctx) error {
	var anl_parm dto.Anl_parm

	setDefaults(&anl_parm)

	if err := json.Unmarshal(c.Body(), &anl_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_parm.Analysis_id = id

    if anl_parm.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_parm set anl_source = :1, parm_obs_no = :2, active_ind = :3, anl_value_remark = :4, average_value = :5, average_value_ouom = :6, average_value_uom = :7, cat_equip_id = :8, effective_date = :9, equip_id = :10, equip_obs_no = :11, expiry_date = :12, max_date = :13, max_value = :14, max_value_ouom = :15, max_value_uom = :16, method_id = :17, method_parm_obs_no = :18, method_set_id = :19, min_date = :20, min_value = :21, min_value_ouom = :22, min_value_uom = :23, parameter_ba_id = :24, parameter_type = :25, ppdm_guid = :26, reference_value = :27, reference_value_ouom = :28, reference_value_type = :29, reference_value_uom = :30, remark = :31, source = :32, step_seq_no = :33, substance_id = :34, row_changed_by = :35, row_changed_date = :36, row_created_by = :37, row_effective_date = :38, row_expiry_date = :39, row_quality = :40 where analysis_id = :42")
	if err != nil {
		return err
	}

	anl_parm.Row_changed_date = formatDate(anl_parm.Row_changed_date)
	anl_parm.Effective_date = formatDateString(anl_parm.Effective_date)
	anl_parm.Expiry_date = formatDateString(anl_parm.Expiry_date)
	anl_parm.Max_date = formatDateString(anl_parm.Max_date)
	anl_parm.Min_date = formatDateString(anl_parm.Min_date)
	anl_parm.Row_effective_date = formatDateString(anl_parm.Row_effective_date)
	anl_parm.Row_expiry_date = formatDateString(anl_parm.Row_expiry_date)






	rows, err := stmt.Exec(anl_parm.Anl_source, anl_parm.Parm_obs_no, anl_parm.Active_ind, anl_parm.Anl_value_remark, anl_parm.Average_value, anl_parm.Average_value_ouom, anl_parm.Average_value_uom, anl_parm.Cat_equip_id, anl_parm.Effective_date, anl_parm.Equip_id, anl_parm.Equip_obs_no, anl_parm.Expiry_date, anl_parm.Max_date, anl_parm.Max_value, anl_parm.Max_value_ouom, anl_parm.Max_value_uom, anl_parm.Method_id, anl_parm.Method_parm_obs_no, anl_parm.Method_set_id, anl_parm.Min_date, anl_parm.Min_value, anl_parm.Min_value_ouom, anl_parm.Min_value_uom, anl_parm.Parameter_ba_id, anl_parm.Parameter_type, anl_parm.Ppdm_guid, anl_parm.Reference_value, anl_parm.Reference_value_ouom, anl_parm.Reference_value_type, anl_parm.Reference_value_uom, anl_parm.Remark, anl_parm.Source, anl_parm.Step_seq_no, anl_parm.Substance_id, anl_parm.Row_changed_by, anl_parm.Row_changed_date, anl_parm.Row_created_by, anl_parm.Row_effective_date, anl_parm.Row_expiry_date, anl_parm.Row_quality, anl_parm.Analysis_id)
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

func PatchAnlParm(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_parm set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteAnlParm(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_parm dto.Anl_parm
	anl_parm.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_parm where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_parm.Analysis_id)
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


