package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestStratUnit(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_strat_unit")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_strat_unit

	for rows.Next() {
		var well_test_strat_unit dto.Well_test_strat_unit
		if err := rows.Scan(&well_test_strat_unit.Uwi, &well_test_strat_unit.Source, &well_test_strat_unit.Test_type, &well_test_strat_unit.Run_num, &well_test_strat_unit.Test_num, &well_test_strat_unit.Strat_unit_id, &well_test_strat_unit.Strat_name_set_id, &well_test_strat_unit.Active_ind, &well_test_strat_unit.Effective_date, &well_test_strat_unit.Expiry_date, &well_test_strat_unit.Ppdm_guid, &well_test_strat_unit.Remark, &well_test_strat_unit.Strat_age, &well_test_strat_unit.Strat_age_name_set_id, &well_test_strat_unit.Row_changed_by, &well_test_strat_unit.Row_changed_date, &well_test_strat_unit.Row_created_by, &well_test_strat_unit.Row_created_date, &well_test_strat_unit.Row_effective_date, &well_test_strat_unit.Row_expiry_date, &well_test_strat_unit.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_strat_unit to result
		result = append(result, well_test_strat_unit)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestStratUnit(c *fiber.Ctx) error {
	var well_test_strat_unit dto.Well_test_strat_unit

	setDefaults(&well_test_strat_unit)

	if err := json.Unmarshal(c.Body(), &well_test_strat_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_strat_unit values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	well_test_strat_unit.Row_created_date = formatDate(well_test_strat_unit.Row_created_date)
	well_test_strat_unit.Row_changed_date = formatDate(well_test_strat_unit.Row_changed_date)
	well_test_strat_unit.Effective_date = formatDateString(well_test_strat_unit.Effective_date)
	well_test_strat_unit.Expiry_date = formatDateString(well_test_strat_unit.Expiry_date)
	well_test_strat_unit.Row_effective_date = formatDateString(well_test_strat_unit.Row_effective_date)
	well_test_strat_unit.Row_expiry_date = formatDateString(well_test_strat_unit.Row_expiry_date)






	rows, err := stmt.Exec(well_test_strat_unit.Uwi, well_test_strat_unit.Source, well_test_strat_unit.Test_type, well_test_strat_unit.Run_num, well_test_strat_unit.Test_num, well_test_strat_unit.Strat_unit_id, well_test_strat_unit.Strat_name_set_id, well_test_strat_unit.Active_ind, well_test_strat_unit.Effective_date, well_test_strat_unit.Expiry_date, well_test_strat_unit.Ppdm_guid, well_test_strat_unit.Remark, well_test_strat_unit.Strat_age, well_test_strat_unit.Strat_age_name_set_id, well_test_strat_unit.Row_changed_by, well_test_strat_unit.Row_changed_date, well_test_strat_unit.Row_created_by, well_test_strat_unit.Row_created_date, well_test_strat_unit.Row_effective_date, well_test_strat_unit.Row_expiry_date, well_test_strat_unit.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestStratUnit(c *fiber.Ctx) error {
	var well_test_strat_unit dto.Well_test_strat_unit

	setDefaults(&well_test_strat_unit)

	if err := json.Unmarshal(c.Body(), &well_test_strat_unit); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_strat_unit.Uwi = id

    if well_test_strat_unit.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_strat_unit set source = :1, test_type = :2, run_num = :3, test_num = :4, strat_unit_id = :5, strat_name_set_id = :6, active_ind = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, strat_age = :12, strat_age_name_set_id = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where uwi = :21")
	if err != nil {
		return err
	}

	well_test_strat_unit.Row_changed_date = formatDate(well_test_strat_unit.Row_changed_date)
	well_test_strat_unit.Effective_date = formatDateString(well_test_strat_unit.Effective_date)
	well_test_strat_unit.Expiry_date = formatDateString(well_test_strat_unit.Expiry_date)
	well_test_strat_unit.Row_effective_date = formatDateString(well_test_strat_unit.Row_effective_date)
	well_test_strat_unit.Row_expiry_date = formatDateString(well_test_strat_unit.Row_expiry_date)






	rows, err := stmt.Exec(well_test_strat_unit.Source, well_test_strat_unit.Test_type, well_test_strat_unit.Run_num, well_test_strat_unit.Test_num, well_test_strat_unit.Strat_unit_id, well_test_strat_unit.Strat_name_set_id, well_test_strat_unit.Active_ind, well_test_strat_unit.Effective_date, well_test_strat_unit.Expiry_date, well_test_strat_unit.Ppdm_guid, well_test_strat_unit.Remark, well_test_strat_unit.Strat_age, well_test_strat_unit.Strat_age_name_set_id, well_test_strat_unit.Row_changed_by, well_test_strat_unit.Row_changed_date, well_test_strat_unit.Row_created_by, well_test_strat_unit.Row_effective_date, well_test_strat_unit.Row_expiry_date, well_test_strat_unit.Row_quality, well_test_strat_unit.Uwi)
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

func PatchWellTestStratUnit(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_strat_unit set "
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
	query += " where uwi = :id"

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

func DeleteWellTestStratUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_strat_unit dto.Well_test_strat_unit
	well_test_strat_unit.Uwi = id

	stmt, err := db.Prepare("delete from well_test_strat_unit where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_strat_unit.Uwi)
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


