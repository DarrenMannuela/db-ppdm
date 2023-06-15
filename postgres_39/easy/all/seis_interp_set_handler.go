package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisInterpSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_interp_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_interp_set

	for rows.Next() {
		var seis_interp_set dto.Seis_interp_set
		if err := rows.Scan(&seis_interp_set.Seis_set_subtype, &seis_interp_set.Interp_set_id, &seis_interp_set.Active_ind, &seis_interp_set.Effective_date, &seis_interp_set.Expiry_date, &seis_interp_set.Interpreter, &seis_interp_set.Interp_date, &seis_interp_set.Interp_objective, &seis_interp_set.Interp_set_name, &seis_interp_set.Interp_type, &seis_interp_set.Pick_method, &seis_interp_set.Ppdm_guid, &seis_interp_set.Remark, &seis_interp_set.Source, &seis_interp_set.Sw_application_id, &seis_interp_set.Trace_position, &seis_interp_set.Row_changed_by, &seis_interp_set.Row_changed_date, &seis_interp_set.Row_created_by, &seis_interp_set.Row_created_date, &seis_interp_set.Row_effective_date, &seis_interp_set.Row_expiry_date, &seis_interp_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_interp_set to result
		result = append(result, seis_interp_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisInterpSet(c *fiber.Ctx) error {
	var seis_interp_set dto.Seis_interp_set

	setDefaults(&seis_interp_set)

	if err := json.Unmarshal(c.Body(), &seis_interp_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_interp_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	seis_interp_set.Row_created_date = formatDate(seis_interp_set.Row_created_date)
	seis_interp_set.Row_changed_date = formatDate(seis_interp_set.Row_changed_date)
	seis_interp_set.Effective_date = formatDateString(seis_interp_set.Effective_date)
	seis_interp_set.Expiry_date = formatDateString(seis_interp_set.Expiry_date)
	seis_interp_set.Interp_date = formatDateString(seis_interp_set.Interp_date)
	seis_interp_set.Row_effective_date = formatDateString(seis_interp_set.Row_effective_date)
	seis_interp_set.Row_expiry_date = formatDateString(seis_interp_set.Row_expiry_date)






	rows, err := stmt.Exec(seis_interp_set.Seis_set_subtype, seis_interp_set.Interp_set_id, seis_interp_set.Active_ind, seis_interp_set.Effective_date, seis_interp_set.Expiry_date, seis_interp_set.Interpreter, seis_interp_set.Interp_date, seis_interp_set.Interp_objective, seis_interp_set.Interp_set_name, seis_interp_set.Interp_type, seis_interp_set.Pick_method, seis_interp_set.Ppdm_guid, seis_interp_set.Remark, seis_interp_set.Source, seis_interp_set.Sw_application_id, seis_interp_set.Trace_position, seis_interp_set.Row_changed_by, seis_interp_set.Row_changed_date, seis_interp_set.Row_created_by, seis_interp_set.Row_created_date, seis_interp_set.Row_effective_date, seis_interp_set.Row_expiry_date, seis_interp_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisInterpSet(c *fiber.Ctx) error {
	var seis_interp_set dto.Seis_interp_set

	setDefaults(&seis_interp_set)

	if err := json.Unmarshal(c.Body(), &seis_interp_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_interp_set.Seis_set_subtype = id

    if seis_interp_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_interp_set set interp_set_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, interpreter = :5, interp_date = :6, interp_objective = :7, interp_set_name = :8, interp_type = :9, pick_method = :10, ppdm_guid = :11, remark = :12, source = :13, sw_application_id = :14, trace_position = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where seis_set_subtype = :23")
	if err != nil {
		return err
	}

	seis_interp_set.Row_changed_date = formatDate(seis_interp_set.Row_changed_date)
	seis_interp_set.Effective_date = formatDateString(seis_interp_set.Effective_date)
	seis_interp_set.Expiry_date = formatDateString(seis_interp_set.Expiry_date)
	seis_interp_set.Interp_date = formatDateString(seis_interp_set.Interp_date)
	seis_interp_set.Row_effective_date = formatDateString(seis_interp_set.Row_effective_date)
	seis_interp_set.Row_expiry_date = formatDateString(seis_interp_set.Row_expiry_date)






	rows, err := stmt.Exec(seis_interp_set.Interp_set_id, seis_interp_set.Active_ind, seis_interp_set.Effective_date, seis_interp_set.Expiry_date, seis_interp_set.Interpreter, seis_interp_set.Interp_date, seis_interp_set.Interp_objective, seis_interp_set.Interp_set_name, seis_interp_set.Interp_type, seis_interp_set.Pick_method, seis_interp_set.Ppdm_guid, seis_interp_set.Remark, seis_interp_set.Source, seis_interp_set.Sw_application_id, seis_interp_set.Trace_position, seis_interp_set.Row_changed_by, seis_interp_set.Row_changed_date, seis_interp_set.Row_created_by, seis_interp_set.Row_effective_date, seis_interp_set.Row_expiry_date, seis_interp_set.Row_quality, seis_interp_set.Seis_set_subtype)
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

func PatchSeisInterpSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_interp_set set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "interp_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSeisInterpSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_interp_set dto.Seis_interp_set
	seis_interp_set.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_interp_set where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_interp_set.Seis_set_subtype)
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


