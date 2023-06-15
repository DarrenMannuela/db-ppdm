package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisProcPlanParm(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_proc_plan_parm")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_proc_plan_parm

	for rows.Next() {
		var seis_proc_plan_parm dto.Seis_proc_plan_parm
		if err := rows.Scan(&seis_proc_plan_parm.Seis_proc_plan_id, &seis_proc_plan_parm.Proc_plan_step_id, &seis_proc_plan_parm.Parameter_id, &seis_proc_plan_parm.Active_ind, &seis_proc_plan_parm.Date_format_desc, &seis_proc_plan_parm.Effective_date, &seis_proc_plan_parm.Expiry_date, &seis_proc_plan_parm.Parameter_desc, &seis_proc_plan_parm.Parameter_origin, &seis_proc_plan_parm.Parameter_type, &seis_proc_plan_parm.Parameter_uom, &seis_proc_plan_parm.Parameter_value_1, &seis_proc_plan_parm.Parameter_value_2, &seis_proc_plan_parm.Ppdm_guid, &seis_proc_plan_parm.Remark, &seis_proc_plan_parm.Source, &seis_proc_plan_parm.Row_changed_by, &seis_proc_plan_parm.Row_changed_date, &seis_proc_plan_parm.Row_created_by, &seis_proc_plan_parm.Row_created_date, &seis_proc_plan_parm.Row_effective_date, &seis_proc_plan_parm.Row_expiry_date, &seis_proc_plan_parm.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_proc_plan_parm to result
		result = append(result, seis_proc_plan_parm)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisProcPlanParm(c *fiber.Ctx) error {
	var seis_proc_plan_parm dto.Seis_proc_plan_parm

	setDefaults(&seis_proc_plan_parm)

	if err := json.Unmarshal(c.Body(), &seis_proc_plan_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_proc_plan_parm values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	seis_proc_plan_parm.Row_created_date = formatDate(seis_proc_plan_parm.Row_created_date)
	seis_proc_plan_parm.Row_changed_date = formatDate(seis_proc_plan_parm.Row_changed_date)
	seis_proc_plan_parm.Date_format_desc = formatDateString(seis_proc_plan_parm.Date_format_desc)
	seis_proc_plan_parm.Effective_date = formatDateString(seis_proc_plan_parm.Effective_date)
	seis_proc_plan_parm.Expiry_date = formatDateString(seis_proc_plan_parm.Expiry_date)
	seis_proc_plan_parm.Row_effective_date = formatDateString(seis_proc_plan_parm.Row_effective_date)
	seis_proc_plan_parm.Row_expiry_date = formatDateString(seis_proc_plan_parm.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_plan_parm.Seis_proc_plan_id, seis_proc_plan_parm.Proc_plan_step_id, seis_proc_plan_parm.Parameter_id, seis_proc_plan_parm.Active_ind, seis_proc_plan_parm.Date_format_desc, seis_proc_plan_parm.Effective_date, seis_proc_plan_parm.Expiry_date, seis_proc_plan_parm.Parameter_desc, seis_proc_plan_parm.Parameter_origin, seis_proc_plan_parm.Parameter_type, seis_proc_plan_parm.Parameter_uom, seis_proc_plan_parm.Parameter_value_1, seis_proc_plan_parm.Parameter_value_2, seis_proc_plan_parm.Ppdm_guid, seis_proc_plan_parm.Remark, seis_proc_plan_parm.Source, seis_proc_plan_parm.Row_changed_by, seis_proc_plan_parm.Row_changed_date, seis_proc_plan_parm.Row_created_by, seis_proc_plan_parm.Row_created_date, seis_proc_plan_parm.Row_effective_date, seis_proc_plan_parm.Row_expiry_date, seis_proc_plan_parm.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisProcPlanParm(c *fiber.Ctx) error {
	var seis_proc_plan_parm dto.Seis_proc_plan_parm

	setDefaults(&seis_proc_plan_parm)

	if err := json.Unmarshal(c.Body(), &seis_proc_plan_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_proc_plan_parm.Seis_proc_plan_id = id

    if seis_proc_plan_parm.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_proc_plan_parm set proc_plan_step_id = :1, parameter_id = :2, active_ind = :3, date_format_desc = :4, effective_date = :5, expiry_date = :6, parameter_desc = :7, parameter_origin = :8, parameter_type = :9, parameter_uom = :10, parameter_value_1 = :11, parameter_value_2 = :12, ppdm_guid = :13, remark = :14, source = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where seis_proc_plan_id = :23")
	if err != nil {
		return err
	}

	seis_proc_plan_parm.Row_changed_date = formatDate(seis_proc_plan_parm.Row_changed_date)
	seis_proc_plan_parm.Date_format_desc = formatDateString(seis_proc_plan_parm.Date_format_desc)
	seis_proc_plan_parm.Effective_date = formatDateString(seis_proc_plan_parm.Effective_date)
	seis_proc_plan_parm.Expiry_date = formatDateString(seis_proc_plan_parm.Expiry_date)
	seis_proc_plan_parm.Row_effective_date = formatDateString(seis_proc_plan_parm.Row_effective_date)
	seis_proc_plan_parm.Row_expiry_date = formatDateString(seis_proc_plan_parm.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_plan_parm.Proc_plan_step_id, seis_proc_plan_parm.Parameter_id, seis_proc_plan_parm.Active_ind, seis_proc_plan_parm.Date_format_desc, seis_proc_plan_parm.Effective_date, seis_proc_plan_parm.Expiry_date, seis_proc_plan_parm.Parameter_desc, seis_proc_plan_parm.Parameter_origin, seis_proc_plan_parm.Parameter_type, seis_proc_plan_parm.Parameter_uom, seis_proc_plan_parm.Parameter_value_1, seis_proc_plan_parm.Parameter_value_2, seis_proc_plan_parm.Ppdm_guid, seis_proc_plan_parm.Remark, seis_proc_plan_parm.Source, seis_proc_plan_parm.Row_changed_by, seis_proc_plan_parm.Row_changed_date, seis_proc_plan_parm.Row_created_by, seis_proc_plan_parm.Row_effective_date, seis_proc_plan_parm.Row_expiry_date, seis_proc_plan_parm.Row_quality, seis_proc_plan_parm.Seis_proc_plan_id)
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

func PatchSeisProcPlanParm(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_proc_plan_parm set "
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
	query += " where seis_proc_plan_id = :id"

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

func DeleteSeisProcPlanParm(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_proc_plan_parm dto.Seis_proc_plan_parm
	seis_proc_plan_parm.Seis_proc_plan_id = id

	stmt, err := db.Prepare("delete from seis_proc_plan_parm where seis_proc_plan_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_proc_plan_parm.Seis_proc_plan_id)
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


