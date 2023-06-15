package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisInterpLoadParm(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_interp_load_parm")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_interp_load_parm

	for rows.Next() {
		var seis_interp_load_parm dto.Seis_interp_load_parm
		if err := rows.Scan(&seis_interp_load_parm.Seis_set_subtype, &seis_interp_load_parm.Interp_set_id, &seis_interp_load_parm.Process_step_id, &seis_interp_load_parm.Parameter_id, &seis_interp_load_parm.Active_ind, &seis_interp_load_parm.Date_format_desc, &seis_interp_load_parm.Effective_date, &seis_interp_load_parm.Expiry_date, &seis_interp_load_parm.Parameter_desc, &seis_interp_load_parm.Parameter_origin, &seis_interp_load_parm.Parameter_type, &seis_interp_load_parm.Parameter_uom, &seis_interp_load_parm.Parameter_value_1, &seis_interp_load_parm.Parameter_value_2, &seis_interp_load_parm.Ppdm_guid, &seis_interp_load_parm.Remark, &seis_interp_load_parm.Source, &seis_interp_load_parm.Row_changed_by, &seis_interp_load_parm.Row_changed_date, &seis_interp_load_parm.Row_created_by, &seis_interp_load_parm.Row_created_date, &seis_interp_load_parm.Row_effective_date, &seis_interp_load_parm.Row_expiry_date, &seis_interp_load_parm.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_interp_load_parm to result
		result = append(result, seis_interp_load_parm)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisInterpLoadParm(c *fiber.Ctx) error {
	var seis_interp_load_parm dto.Seis_interp_load_parm

	setDefaults(&seis_interp_load_parm)

	if err := json.Unmarshal(c.Body(), &seis_interp_load_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_interp_load_parm values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	seis_interp_load_parm.Row_created_date = formatDate(seis_interp_load_parm.Row_created_date)
	seis_interp_load_parm.Row_changed_date = formatDate(seis_interp_load_parm.Row_changed_date)
	seis_interp_load_parm.Date_format_desc = formatDateString(seis_interp_load_parm.Date_format_desc)
	seis_interp_load_parm.Effective_date = formatDateString(seis_interp_load_parm.Effective_date)
	seis_interp_load_parm.Expiry_date = formatDateString(seis_interp_load_parm.Expiry_date)
	seis_interp_load_parm.Row_effective_date = formatDateString(seis_interp_load_parm.Row_effective_date)
	seis_interp_load_parm.Row_expiry_date = formatDateString(seis_interp_load_parm.Row_expiry_date)






	rows, err := stmt.Exec(seis_interp_load_parm.Seis_set_subtype, seis_interp_load_parm.Interp_set_id, seis_interp_load_parm.Process_step_id, seis_interp_load_parm.Parameter_id, seis_interp_load_parm.Active_ind, seis_interp_load_parm.Date_format_desc, seis_interp_load_parm.Effective_date, seis_interp_load_parm.Expiry_date, seis_interp_load_parm.Parameter_desc, seis_interp_load_parm.Parameter_origin, seis_interp_load_parm.Parameter_type, seis_interp_load_parm.Parameter_uom, seis_interp_load_parm.Parameter_value_1, seis_interp_load_parm.Parameter_value_2, seis_interp_load_parm.Ppdm_guid, seis_interp_load_parm.Remark, seis_interp_load_parm.Source, seis_interp_load_parm.Row_changed_by, seis_interp_load_parm.Row_changed_date, seis_interp_load_parm.Row_created_by, seis_interp_load_parm.Row_created_date, seis_interp_load_parm.Row_effective_date, seis_interp_load_parm.Row_expiry_date, seis_interp_load_parm.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisInterpLoadParm(c *fiber.Ctx) error {
	var seis_interp_load_parm dto.Seis_interp_load_parm

	setDefaults(&seis_interp_load_parm)

	if err := json.Unmarshal(c.Body(), &seis_interp_load_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_interp_load_parm.Seis_set_subtype = id

    if seis_interp_load_parm.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_interp_load_parm set interp_set_id = :1, process_step_id = :2, parameter_id = :3, active_ind = :4, date_format_desc = :5, effective_date = :6, expiry_date = :7, parameter_desc = :8, parameter_origin = :9, parameter_type = :10, parameter_uom = :11, parameter_value_1 = :12, parameter_value_2 = :13, ppdm_guid = :14, remark = :15, source = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where seis_set_subtype = :24")
	if err != nil {
		return err
	}

	seis_interp_load_parm.Row_changed_date = formatDate(seis_interp_load_parm.Row_changed_date)
	seis_interp_load_parm.Date_format_desc = formatDateString(seis_interp_load_parm.Date_format_desc)
	seis_interp_load_parm.Effective_date = formatDateString(seis_interp_load_parm.Effective_date)
	seis_interp_load_parm.Expiry_date = formatDateString(seis_interp_load_parm.Expiry_date)
	seis_interp_load_parm.Row_effective_date = formatDateString(seis_interp_load_parm.Row_effective_date)
	seis_interp_load_parm.Row_expiry_date = formatDateString(seis_interp_load_parm.Row_expiry_date)






	rows, err := stmt.Exec(seis_interp_load_parm.Interp_set_id, seis_interp_load_parm.Process_step_id, seis_interp_load_parm.Parameter_id, seis_interp_load_parm.Active_ind, seis_interp_load_parm.Date_format_desc, seis_interp_load_parm.Effective_date, seis_interp_load_parm.Expiry_date, seis_interp_load_parm.Parameter_desc, seis_interp_load_parm.Parameter_origin, seis_interp_load_parm.Parameter_type, seis_interp_load_parm.Parameter_uom, seis_interp_load_parm.Parameter_value_1, seis_interp_load_parm.Parameter_value_2, seis_interp_load_parm.Ppdm_guid, seis_interp_load_parm.Remark, seis_interp_load_parm.Source, seis_interp_load_parm.Row_changed_by, seis_interp_load_parm.Row_changed_date, seis_interp_load_parm.Row_created_by, seis_interp_load_parm.Row_effective_date, seis_interp_load_parm.Row_expiry_date, seis_interp_load_parm.Row_quality, seis_interp_load_parm.Seis_set_subtype)
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

func PatchSeisInterpLoadParm(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_interp_load_parm set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where seis_set_subtype = :id"

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

func DeleteSeisInterpLoadParm(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_interp_load_parm dto.Seis_interp_load_parm
	seis_interp_load_parm.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_interp_load_parm where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_interp_load_parm.Seis_set_subtype)
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


