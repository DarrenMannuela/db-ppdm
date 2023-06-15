package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlCalcEquiv(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_calc_equiv")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_calc_equiv

	for rows.Next() {
		var anl_calc_equiv dto.Anl_calc_equiv
		if err := rows.Scan(&anl_calc_equiv.Calculate_method_id1, &anl_calc_equiv.Calculate_method_id2, &anl_calc_equiv.Active_ind, &anl_calc_equiv.Calculation_equiv_type, &anl_calc_equiv.Effective_date, &anl_calc_equiv.Expiry_date, &anl_calc_equiv.Ppdm_guid, &anl_calc_equiv.Remark, &anl_calc_equiv.Source, &anl_calc_equiv.Row_changed_by, &anl_calc_equiv.Row_changed_date, &anl_calc_equiv.Row_created_by, &anl_calc_equiv.Row_created_date, &anl_calc_equiv.Row_effective_date, &anl_calc_equiv.Row_expiry_date, &anl_calc_equiv.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_calc_equiv to result
		result = append(result, anl_calc_equiv)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlCalcEquiv(c *fiber.Ctx) error {
	var anl_calc_equiv dto.Anl_calc_equiv

	setDefaults(&anl_calc_equiv)

	if err := json.Unmarshal(c.Body(), &anl_calc_equiv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_calc_equiv values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	anl_calc_equiv.Row_created_date = formatDate(anl_calc_equiv.Row_created_date)
	anl_calc_equiv.Row_changed_date = formatDate(anl_calc_equiv.Row_changed_date)
	anl_calc_equiv.Effective_date = formatDateString(anl_calc_equiv.Effective_date)
	anl_calc_equiv.Expiry_date = formatDateString(anl_calc_equiv.Expiry_date)
	anl_calc_equiv.Row_effective_date = formatDateString(anl_calc_equiv.Row_effective_date)
	anl_calc_equiv.Row_expiry_date = formatDateString(anl_calc_equiv.Row_expiry_date)






	rows, err := stmt.Exec(anl_calc_equiv.Calculate_method_id1, anl_calc_equiv.Calculate_method_id2, anl_calc_equiv.Active_ind, anl_calc_equiv.Calculation_equiv_type, anl_calc_equiv.Effective_date, anl_calc_equiv.Expiry_date, anl_calc_equiv.Ppdm_guid, anl_calc_equiv.Remark, anl_calc_equiv.Source, anl_calc_equiv.Row_changed_by, anl_calc_equiv.Row_changed_date, anl_calc_equiv.Row_created_by, anl_calc_equiv.Row_created_date, anl_calc_equiv.Row_effective_date, anl_calc_equiv.Row_expiry_date, anl_calc_equiv.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlCalcEquiv(c *fiber.Ctx) error {
	var anl_calc_equiv dto.Anl_calc_equiv

	setDefaults(&anl_calc_equiv)

	if err := json.Unmarshal(c.Body(), &anl_calc_equiv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_calc_equiv.Calculate_method_id1 = id

    if anl_calc_equiv.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_calc_equiv set calculate_method_id2 = :1, active_ind = :2, calculation_equiv_type = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where calculate_method_id1 = :16")
	if err != nil {
		return err
	}

	anl_calc_equiv.Row_changed_date = formatDate(anl_calc_equiv.Row_changed_date)
	anl_calc_equiv.Effective_date = formatDateString(anl_calc_equiv.Effective_date)
	anl_calc_equiv.Expiry_date = formatDateString(anl_calc_equiv.Expiry_date)
	anl_calc_equiv.Row_effective_date = formatDateString(anl_calc_equiv.Row_effective_date)
	anl_calc_equiv.Row_expiry_date = formatDateString(anl_calc_equiv.Row_expiry_date)






	rows, err := stmt.Exec(anl_calc_equiv.Calculate_method_id2, anl_calc_equiv.Active_ind, anl_calc_equiv.Calculation_equiv_type, anl_calc_equiv.Effective_date, anl_calc_equiv.Expiry_date, anl_calc_equiv.Ppdm_guid, anl_calc_equiv.Remark, anl_calc_equiv.Source, anl_calc_equiv.Row_changed_by, anl_calc_equiv.Row_changed_date, anl_calc_equiv.Row_created_by, anl_calc_equiv.Row_effective_date, anl_calc_equiv.Row_expiry_date, anl_calc_equiv.Row_quality, anl_calc_equiv.Calculate_method_id1)
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

func PatchAnlCalcEquiv(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_calc_equiv set "
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
	query += " where calculate_method_id1 = :id"

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

func DeleteAnlCalcEquiv(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_calc_equiv dto.Anl_calc_equiv
	anl_calc_equiv.Calculate_method_id1 = id

	stmt, err := db.Prepare("delete from anl_calc_equiv where calculate_method_id1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_calc_equiv.Calculate_method_id1)
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


