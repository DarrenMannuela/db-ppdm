package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlMethodParm(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_method_parm")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_method_parm

	for rows.Next() {
		var anl_method_parm dto.Anl_method_parm
		if err := rows.Scan(&anl_method_parm.Method_set_id, &anl_method_parm.Method_id, &anl_method_parm.Parameter_type, &anl_method_parm.Method_parm_obs_no, &anl_method_parm.Abbreviation, &anl_method_parm.Active_ind, &anl_method_parm.Anl_value_remark, &anl_method_parm.Average_value, &anl_method_parm.Average_value_ouom, &anl_method_parm.Average_value_uom, &anl_method_parm.Effective_date, &anl_method_parm.Expiry_date, &anl_method_parm.Long_name, &anl_method_parm.Max_date, &anl_method_parm.Max_value, &anl_method_parm.Max_value_ouom, &anl_method_parm.Max_value_uom, &anl_method_parm.Measurement_type, &anl_method_parm.Min_date, &anl_method_parm.Min_value, &anl_method_parm.Min_value_ouom, &anl_method_parm.Min_value_uom, &anl_method_parm.Ppdm_guid, &anl_method_parm.Reference_value, &anl_method_parm.Reference_value_ouom, &anl_method_parm.Reference_value_type, &anl_method_parm.Reference_value_uom, &anl_method_parm.Remark, &anl_method_parm.Short_name, &anl_method_parm.Source, &anl_method_parm.Substance_id, &anl_method_parm.Row_changed_by, &anl_method_parm.Row_changed_date, &anl_method_parm.Row_created_by, &anl_method_parm.Row_created_date, &anl_method_parm.Row_effective_date, &anl_method_parm.Row_expiry_date, &anl_method_parm.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_method_parm to result
		result = append(result, anl_method_parm)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlMethodParm(c *fiber.Ctx) error {
	var anl_method_parm dto.Anl_method_parm

	setDefaults(&anl_method_parm)

	if err := json.Unmarshal(c.Body(), &anl_method_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_method_parm values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	anl_method_parm.Row_created_date = formatDate(anl_method_parm.Row_created_date)
	anl_method_parm.Row_changed_date = formatDate(anl_method_parm.Row_changed_date)
	anl_method_parm.Effective_date = formatDateString(anl_method_parm.Effective_date)
	anl_method_parm.Expiry_date = formatDateString(anl_method_parm.Expiry_date)
	anl_method_parm.Max_date = formatDateString(anl_method_parm.Max_date)
	anl_method_parm.Min_date = formatDateString(anl_method_parm.Min_date)
	anl_method_parm.Row_effective_date = formatDateString(anl_method_parm.Row_effective_date)
	anl_method_parm.Row_expiry_date = formatDateString(anl_method_parm.Row_expiry_date)






	rows, err := stmt.Exec(anl_method_parm.Method_set_id, anl_method_parm.Method_id, anl_method_parm.Parameter_type, anl_method_parm.Method_parm_obs_no, anl_method_parm.Abbreviation, anl_method_parm.Active_ind, anl_method_parm.Anl_value_remark, anl_method_parm.Average_value, anl_method_parm.Average_value_ouom, anl_method_parm.Average_value_uom, anl_method_parm.Effective_date, anl_method_parm.Expiry_date, anl_method_parm.Long_name, anl_method_parm.Max_date, anl_method_parm.Max_value, anl_method_parm.Max_value_ouom, anl_method_parm.Max_value_uom, anl_method_parm.Measurement_type, anl_method_parm.Min_date, anl_method_parm.Min_value, anl_method_parm.Min_value_ouom, anl_method_parm.Min_value_uom, anl_method_parm.Ppdm_guid, anl_method_parm.Reference_value, anl_method_parm.Reference_value_ouom, anl_method_parm.Reference_value_type, anl_method_parm.Reference_value_uom, anl_method_parm.Remark, anl_method_parm.Short_name, anl_method_parm.Source, anl_method_parm.Substance_id, anl_method_parm.Row_changed_by, anl_method_parm.Row_changed_date, anl_method_parm.Row_created_by, anl_method_parm.Row_created_date, anl_method_parm.Row_effective_date, anl_method_parm.Row_expiry_date, anl_method_parm.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlMethodParm(c *fiber.Ctx) error {
	var anl_method_parm dto.Anl_method_parm

	setDefaults(&anl_method_parm)

	if err := json.Unmarshal(c.Body(), &anl_method_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_method_parm.Method_set_id = id

    if anl_method_parm.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_method_parm set method_id = :1, parameter_type = :2, method_parm_obs_no = :3, abbreviation = :4, active_ind = :5, anl_value_remark = :6, average_value = :7, average_value_ouom = :8, average_value_uom = :9, effective_date = :10, expiry_date = :11, long_name = :12, max_date = :13, max_value = :14, max_value_ouom = :15, max_value_uom = :16, measurement_type = :17, min_date = :18, min_value = :19, min_value_ouom = :20, min_value_uom = :21, ppdm_guid = :22, reference_value = :23, reference_value_ouom = :24, reference_value_type = :25, reference_value_uom = :26, remark = :27, short_name = :28, source = :29, substance_id = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where method_set_id = :38")
	if err != nil {
		return err
	}

	anl_method_parm.Row_changed_date = formatDate(anl_method_parm.Row_changed_date)
	anl_method_parm.Effective_date = formatDateString(anl_method_parm.Effective_date)
	anl_method_parm.Expiry_date = formatDateString(anl_method_parm.Expiry_date)
	anl_method_parm.Max_date = formatDateString(anl_method_parm.Max_date)
	anl_method_parm.Min_date = formatDateString(anl_method_parm.Min_date)
	anl_method_parm.Row_effective_date = formatDateString(anl_method_parm.Row_effective_date)
	anl_method_parm.Row_expiry_date = formatDateString(anl_method_parm.Row_expiry_date)






	rows, err := stmt.Exec(anl_method_parm.Method_id, anl_method_parm.Parameter_type, anl_method_parm.Method_parm_obs_no, anl_method_parm.Abbreviation, anl_method_parm.Active_ind, anl_method_parm.Anl_value_remark, anl_method_parm.Average_value, anl_method_parm.Average_value_ouom, anl_method_parm.Average_value_uom, anl_method_parm.Effective_date, anl_method_parm.Expiry_date, anl_method_parm.Long_name, anl_method_parm.Max_date, anl_method_parm.Max_value, anl_method_parm.Max_value_ouom, anl_method_parm.Max_value_uom, anl_method_parm.Measurement_type, anl_method_parm.Min_date, anl_method_parm.Min_value, anl_method_parm.Min_value_ouom, anl_method_parm.Min_value_uom, anl_method_parm.Ppdm_guid, anl_method_parm.Reference_value, anl_method_parm.Reference_value_ouom, anl_method_parm.Reference_value_type, anl_method_parm.Reference_value_uom, anl_method_parm.Remark, anl_method_parm.Short_name, anl_method_parm.Source, anl_method_parm.Substance_id, anl_method_parm.Row_changed_by, anl_method_parm.Row_changed_date, anl_method_parm.Row_created_by, anl_method_parm.Row_effective_date, anl_method_parm.Row_expiry_date, anl_method_parm.Row_quality, anl_method_parm.Method_set_id)
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

func PatchAnlMethodParm(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_method_parm set "
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
	query += " where method_set_id = :id"

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

func DeleteAnlMethodParm(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_method_parm dto.Anl_method_parm
	anl_method_parm.Method_set_id = id

	stmt, err := db.Prepare("delete from anl_method_parm where method_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_method_parm.Method_set_id)
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


