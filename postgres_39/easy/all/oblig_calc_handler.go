package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligCalc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_calc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_calc

	for rows.Next() {
		var oblig_calc dto.Oblig_calc
		if err := rows.Scan(&oblig_calc.Obligation_id, &oblig_calc.Obligation_seq_no, &oblig_calc.Calculation_obs_no, &oblig_calc.Active_ind, &oblig_calc.Calculation_formula, &oblig_calc.Calculation_method, &oblig_calc.Calculation_point, &oblig_calc.Effective_date, &oblig_calc.Expiry_date, &oblig_calc.Ppdm_guid, &oblig_calc.Remark, &oblig_calc.Source, &oblig_calc.Substance_id, &oblig_calc.Row_changed_by, &oblig_calc.Row_changed_date, &oblig_calc.Row_created_by, &oblig_calc.Row_created_date, &oblig_calc.Row_effective_date, &oblig_calc.Row_expiry_date, &oblig_calc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_calc to result
		result = append(result, oblig_calc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligCalc(c *fiber.Ctx) error {
	var oblig_calc dto.Oblig_calc

	setDefaults(&oblig_calc)

	if err := json.Unmarshal(c.Body(), &oblig_calc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_calc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	oblig_calc.Row_created_date = formatDate(oblig_calc.Row_created_date)
	oblig_calc.Row_changed_date = formatDate(oblig_calc.Row_changed_date)
	oblig_calc.Effective_date = formatDateString(oblig_calc.Effective_date)
	oblig_calc.Expiry_date = formatDateString(oblig_calc.Expiry_date)
	oblig_calc.Row_effective_date = formatDateString(oblig_calc.Row_effective_date)
	oblig_calc.Row_expiry_date = formatDateString(oblig_calc.Row_expiry_date)






	rows, err := stmt.Exec(oblig_calc.Obligation_id, oblig_calc.Obligation_seq_no, oblig_calc.Calculation_obs_no, oblig_calc.Active_ind, oblig_calc.Calculation_formula, oblig_calc.Calculation_method, oblig_calc.Calculation_point, oblig_calc.Effective_date, oblig_calc.Expiry_date, oblig_calc.Ppdm_guid, oblig_calc.Remark, oblig_calc.Source, oblig_calc.Substance_id, oblig_calc.Row_changed_by, oblig_calc.Row_changed_date, oblig_calc.Row_created_by, oblig_calc.Row_created_date, oblig_calc.Row_effective_date, oblig_calc.Row_expiry_date, oblig_calc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligCalc(c *fiber.Ctx) error {
	var oblig_calc dto.Oblig_calc

	setDefaults(&oblig_calc)

	if err := json.Unmarshal(c.Body(), &oblig_calc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_calc.Obligation_id = id

    if oblig_calc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_calc set obligation_seq_no = :1, calculation_obs_no = :2, active_ind = :3, calculation_formula = :4, calculation_method = :5, calculation_point = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, substance_id = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where obligation_id = :20")
	if err != nil {
		return err
	}

	oblig_calc.Row_changed_date = formatDate(oblig_calc.Row_changed_date)
	oblig_calc.Effective_date = formatDateString(oblig_calc.Effective_date)
	oblig_calc.Expiry_date = formatDateString(oblig_calc.Expiry_date)
	oblig_calc.Row_effective_date = formatDateString(oblig_calc.Row_effective_date)
	oblig_calc.Row_expiry_date = formatDateString(oblig_calc.Row_expiry_date)






	rows, err := stmt.Exec(oblig_calc.Obligation_seq_no, oblig_calc.Calculation_obs_no, oblig_calc.Active_ind, oblig_calc.Calculation_formula, oblig_calc.Calculation_method, oblig_calc.Calculation_point, oblig_calc.Effective_date, oblig_calc.Expiry_date, oblig_calc.Ppdm_guid, oblig_calc.Remark, oblig_calc.Source, oblig_calc.Substance_id, oblig_calc.Row_changed_by, oblig_calc.Row_changed_date, oblig_calc.Row_created_by, oblig_calc.Row_effective_date, oblig_calc.Row_expiry_date, oblig_calc.Row_quality, oblig_calc.Obligation_id)
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

func PatchObligCalc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_calc set "
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
	query += " where obligation_id = :id"

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

func DeleteObligCalc(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_calc dto.Oblig_calc
	oblig_calc.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_calc where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_calc.Obligation_id)
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


