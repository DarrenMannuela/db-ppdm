package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlCalcFormula(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_calc_formula")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_calc_formula

	for rows.Next() {
		var anl_calc_formula dto.Anl_calc_formula
		if err := rows.Scan(&anl_calc_formula.Calculate_method_id, &anl_calc_formula.Formula_part_obs_no, &anl_calc_formula.Active_ind, &anl_calc_formula.Assigned_variable, &anl_calc_formula.Column_name, &anl_calc_formula.Effective_date, &anl_calc_formula.Expiry_date, &anl_calc_formula.Ppdm_guid, &anl_calc_formula.Remark, &anl_calc_formula.Source, &anl_calc_formula.Substance_id, &anl_calc_formula.System_id, &anl_calc_formula.Table_name, &anl_calc_formula.Row_changed_by, &anl_calc_formula.Row_changed_date, &anl_calc_formula.Row_created_by, &anl_calc_formula.Row_created_date, &anl_calc_formula.Row_effective_date, &anl_calc_formula.Row_expiry_date, &anl_calc_formula.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_calc_formula to result
		result = append(result, anl_calc_formula)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlCalcFormula(c *fiber.Ctx) error {
	var anl_calc_formula dto.Anl_calc_formula

	setDefaults(&anl_calc_formula)

	if err := json.Unmarshal(c.Body(), &anl_calc_formula); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_calc_formula values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	anl_calc_formula.Row_created_date = formatDate(anl_calc_formula.Row_created_date)
	anl_calc_formula.Row_changed_date = formatDate(anl_calc_formula.Row_changed_date)
	anl_calc_formula.Effective_date = formatDateString(anl_calc_formula.Effective_date)
	anl_calc_formula.Expiry_date = formatDateString(anl_calc_formula.Expiry_date)
	anl_calc_formula.Row_effective_date = formatDateString(anl_calc_formula.Row_effective_date)
	anl_calc_formula.Row_expiry_date = formatDateString(anl_calc_formula.Row_expiry_date)






	rows, err := stmt.Exec(anl_calc_formula.Calculate_method_id, anl_calc_formula.Formula_part_obs_no, anl_calc_formula.Active_ind, anl_calc_formula.Assigned_variable, anl_calc_formula.Column_name, anl_calc_formula.Effective_date, anl_calc_formula.Expiry_date, anl_calc_formula.Ppdm_guid, anl_calc_formula.Remark, anl_calc_formula.Source, anl_calc_formula.Substance_id, anl_calc_formula.System_id, anl_calc_formula.Table_name, anl_calc_formula.Row_changed_by, anl_calc_formula.Row_changed_date, anl_calc_formula.Row_created_by, anl_calc_formula.Row_created_date, anl_calc_formula.Row_effective_date, anl_calc_formula.Row_expiry_date, anl_calc_formula.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlCalcFormula(c *fiber.Ctx) error {
	var anl_calc_formula dto.Anl_calc_formula

	setDefaults(&anl_calc_formula)

	if err := json.Unmarshal(c.Body(), &anl_calc_formula); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_calc_formula.Calculate_method_id = id

    if anl_calc_formula.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_calc_formula set formula_part_obs_no = :1, active_ind = :2, assigned_variable = :3, column_name = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, substance_id = :10, system_id = :11, table_name = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where calculate_method_id = :20")
	if err != nil {
		return err
	}

	anl_calc_formula.Row_changed_date = formatDate(anl_calc_formula.Row_changed_date)
	anl_calc_formula.Effective_date = formatDateString(anl_calc_formula.Effective_date)
	anl_calc_formula.Expiry_date = formatDateString(anl_calc_formula.Expiry_date)
	anl_calc_formula.Row_effective_date = formatDateString(anl_calc_formula.Row_effective_date)
	anl_calc_formula.Row_expiry_date = formatDateString(anl_calc_formula.Row_expiry_date)






	rows, err := stmt.Exec(anl_calc_formula.Formula_part_obs_no, anl_calc_formula.Active_ind, anl_calc_formula.Assigned_variable, anl_calc_formula.Column_name, anl_calc_formula.Effective_date, anl_calc_formula.Expiry_date, anl_calc_formula.Ppdm_guid, anl_calc_formula.Remark, anl_calc_formula.Source, anl_calc_formula.Substance_id, anl_calc_formula.System_id, anl_calc_formula.Table_name, anl_calc_formula.Row_changed_by, anl_calc_formula.Row_changed_date, anl_calc_formula.Row_created_by, anl_calc_formula.Row_effective_date, anl_calc_formula.Row_expiry_date, anl_calc_formula.Row_quality, anl_calc_formula.Calculate_method_id)
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

func PatchAnlCalcFormula(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_calc_formula set "
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
	query += " where calculate_method_id = :id"

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

func DeleteAnlCalcFormula(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_calc_formula dto.Anl_calc_formula
	anl_calc_formula.Calculate_method_id = id

	stmt, err := db.Prepare("delete from anl_calc_formula where calculate_method_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_calc_formula.Calculate_method_id)
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


