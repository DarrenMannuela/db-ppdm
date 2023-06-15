package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratUnit(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_unit")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_unit

	for rows.Next() {
		var strat_unit dto.Strat_unit
		if err := rows.Scan(&strat_unit.Strat_name_set_id, &strat_unit.Strat_unit_id, &strat_unit.Abbreviation, &strat_unit.Active_ind, &strat_unit.Area_id, &strat_unit.Area_type, &strat_unit.Business_associate_id, &strat_unit.Confidence_id, &strat_unit.Current_status_date, &strat_unit.Description, &strat_unit.Effective_date, &strat_unit.Expiry_date, &strat_unit.Fault_type, &strat_unit.Form_code, &strat_unit.Group_code, &strat_unit.Long_name, &strat_unit.Ordinal_age_code, &strat_unit.Ppdm_guid, &strat_unit.Preferred_ind, &strat_unit.Remark, &strat_unit.Short_name, &strat_unit.Source, &strat_unit.Strat_interpret_method, &strat_unit.Strat_status, &strat_unit.Strat_type, &strat_unit.Strat_unit_type, &strat_unit.Row_changed_by, &strat_unit.Row_changed_date, &strat_unit.Row_created_by, &strat_unit.Row_created_date, &strat_unit.Row_effective_date, &strat_unit.Row_expiry_date, &strat_unit.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_unit to result
		result = append(result, strat_unit)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratUnit(c *fiber.Ctx) error {
	var strat_unit dto.Strat_unit

	setDefaults(&strat_unit)

	if err := json.Unmarshal(c.Body(), &strat_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_unit values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	strat_unit.Row_created_date = formatDate(strat_unit.Row_created_date)
	strat_unit.Row_changed_date = formatDate(strat_unit.Row_changed_date)
	strat_unit.Current_status_date = formatDateString(strat_unit.Current_status_date)
	strat_unit.Effective_date = formatDateString(strat_unit.Effective_date)
	strat_unit.Expiry_date = formatDateString(strat_unit.Expiry_date)
	strat_unit.Row_effective_date = formatDateString(strat_unit.Row_effective_date)
	strat_unit.Row_expiry_date = formatDateString(strat_unit.Row_expiry_date)






	rows, err := stmt.Exec(strat_unit.Strat_name_set_id, strat_unit.Strat_unit_id, strat_unit.Abbreviation, strat_unit.Active_ind, strat_unit.Area_id, strat_unit.Area_type, strat_unit.Business_associate_id, strat_unit.Confidence_id, strat_unit.Current_status_date, strat_unit.Description, strat_unit.Effective_date, strat_unit.Expiry_date, strat_unit.Fault_type, strat_unit.Form_code, strat_unit.Group_code, strat_unit.Long_name, strat_unit.Ordinal_age_code, strat_unit.Ppdm_guid, strat_unit.Preferred_ind, strat_unit.Remark, strat_unit.Short_name, strat_unit.Source, strat_unit.Strat_interpret_method, strat_unit.Strat_status, strat_unit.Strat_type, strat_unit.Strat_unit_type, strat_unit.Row_changed_by, strat_unit.Row_changed_date, strat_unit.Row_created_by, strat_unit.Row_created_date, strat_unit.Row_effective_date, strat_unit.Row_expiry_date, strat_unit.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratUnit(c *fiber.Ctx) error {
	var strat_unit dto.Strat_unit

	setDefaults(&strat_unit)

	if err := json.Unmarshal(c.Body(), &strat_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_unit.Strat_name_set_id = id

    if strat_unit.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_unit set strat_unit_id = :1, abbreviation = :2, active_ind = :3, area_id = :4, area_type = :5, business_associate_id = :6, confidence_id = :7, current_status_date = :8, description = :9, effective_date = :10, expiry_date = :11, fault_type = :12, form_code = :13, group_code = :14, long_name = :15, ordinal_age_code = :16, ppdm_guid = :17, preferred_ind = :18, remark = :19, short_name = :20, source = :21, strat_interpret_method = :22, strat_status = :23, strat_type = :24, strat_unit_type = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where strat_name_set_id = :33")
	if err != nil {
		return err
	}

	strat_unit.Row_changed_date = formatDate(strat_unit.Row_changed_date)
	strat_unit.Current_status_date = formatDateString(strat_unit.Current_status_date)
	strat_unit.Effective_date = formatDateString(strat_unit.Effective_date)
	strat_unit.Expiry_date = formatDateString(strat_unit.Expiry_date)
	strat_unit.Row_effective_date = formatDateString(strat_unit.Row_effective_date)
	strat_unit.Row_expiry_date = formatDateString(strat_unit.Row_expiry_date)






	rows, err := stmt.Exec(strat_unit.Strat_unit_id, strat_unit.Abbreviation, strat_unit.Active_ind, strat_unit.Area_id, strat_unit.Area_type, strat_unit.Business_associate_id, strat_unit.Confidence_id, strat_unit.Current_status_date, strat_unit.Description, strat_unit.Effective_date, strat_unit.Expiry_date, strat_unit.Fault_type, strat_unit.Form_code, strat_unit.Group_code, strat_unit.Long_name, strat_unit.Ordinal_age_code, strat_unit.Ppdm_guid, strat_unit.Preferred_ind, strat_unit.Remark, strat_unit.Short_name, strat_unit.Source, strat_unit.Strat_interpret_method, strat_unit.Strat_status, strat_unit.Strat_type, strat_unit.Strat_unit_type, strat_unit.Row_changed_by, strat_unit.Row_changed_date, strat_unit.Row_created_by, strat_unit.Row_effective_date, strat_unit.Row_expiry_date, strat_unit.Row_quality, strat_unit.Strat_name_set_id)
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

func PatchStratUnit(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_unit set "
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
		} else if key == "current_status_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where strat_name_set_id = :id"

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

func DeleteStratUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_unit dto.Strat_unit
	strat_unit.Strat_name_set_id = id

	stmt, err := db.Prepare("delete from strat_unit where strat_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_unit.Strat_name_set_id)
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


